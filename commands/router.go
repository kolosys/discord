package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kolosys/discord/models"
	"github.com/kolosys/discord/types"
)

// CommandSyncer is an interface for syncing commands with Discord.
type CommandSyncer interface {
	SyncCommands(ctx context.Context, appID string, commands []ApplicationCommandCreate) error
	SyncGuildCommands(ctx context.Context, appID, guildID string, commands []ApplicationCommandCreate) error
}

// Router handles routing interactions to commands and components.
type Router struct {
	mu                  sync.RWMutex
	commands            map[string]Command
	groups              map[string]*CommandGroup
	components          map[string]ComponentHandlerFunc
	middleware          []Middleware
	componentMiddleware []ComponentMiddleware

	// Syncer for syncing commands with Discord
	syncer CommandSyncer
	appID  string

	// Responder for creating contexts
	responder Responder

	// Services for dependency injection
	services *ServiceRegistry
}

// ComponentMiddleware wraps a component handler function.
type ComponentMiddleware func(next ComponentHandlerFunc) ComponentHandlerFunc

// CommandGroup represents a group of related commands.
type CommandGroup struct {
	Name        string
	Description string
	Commands    map[string]Command
	Middleware  []Middleware
}

// NewRouter creates a new command router.
func NewRouter() *Router {
	return &Router{
		commands:   make(map[string]Command),
		groups:     make(map[string]*CommandGroup),
		components: make(map[string]ComponentHandlerFunc),
		services:   newServiceRegistry(),
	}
}

// RegisterService registers a service for dependency injection.
// Services can be retrieved in handlers using commands.Service[T](ctx).
func (r *Router) RegisterService(service any) {
	r.services.Register(service)
}

// SetSyncer sets the command syncer.
func (r *Router) SetSyncer(syncer CommandSyncer, appID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.syncer = syncer
	r.appID = appID
}

// SetResponder sets the responder for creating contexts.
func (r *Router) SetResponder(responder Responder) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.responder = responder
}

// Responder returns the current responder, or nil if not set.
func (r *Router) Responder() Responder {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.responder
}

// Command registers a command.
func (r *Router) Command(cmd Command) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.commands[cmd.Name()] = cmd
}

// Commands returns all registered commands.
func (r *Router) Commands() []Command {
	r.mu.RLock()
	defer r.mu.RUnlock()

	cmds := make([]Command, 0, len(r.commands))
	for _, cmd := range r.commands {
		cmds = append(cmds, cmd)
	}
	return cmds
}

// GetCommand returns a command by name.
func (r *Router) GetCommand(name string) (Command, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	cmd, ok := r.commands[name]
	return cmd, ok
}

// Slash registers a slash command with a handler.
func (r *Router) Slash(name, description string, handler HandlerFunc, options ...Option) {
	r.Command(&SlashCommand{
		CommandName:        name,
		CommandDescription: description,
		CommandOptions:     options,
		Handler:            handler,
	})
}

// UserContext registers a user context menu command.
func (r *Router) UserContext(name string, handler HandlerFunc) {
	r.Command(&UserCommand{
		CommandName: name,
		Handler:     handler,
	})
}

// MessageContext registers a message context menu command.
func (r *Router) MessageContext(name string, handler HandlerFunc) {
	r.Command(&MessageCommand{
		CommandName: name,
		Handler:     handler,
	})
}

// Group registers a command group (subcommand group).
func (r *Router) Group(name, description string, cmds ...Command) {
	r.mu.Lock()
	defer r.mu.Unlock()

	group := &CommandGroup{
		Name:        name,
		Description: description,
		Commands:    make(map[string]Command),
	}

	for _, cmd := range cmds {
		group.Commands[cmd.Name()] = cmd
	}

	r.groups[name] = group
}

// Use adds middleware to the router for command handlers.
func (r *Router) Use(mw ...Middleware) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.middleware = append(r.middleware, mw...)
}

// UseComponent adds middleware to the router for component handlers.
func (r *Router) UseComponent(mw ...ComponentMiddleware) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.componentMiddleware = append(r.componentMiddleware, mw...)
}

// Component registers a component handler.
func (r *Router) Component(customID string, handler ComponentHandlerFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.components[customID] = handler
}

// ComponentPrefix registers a component handler that matches by prefix.
func (r *Router) ComponentPrefix(prefix string, handler ComponentHandlerFunc) {
	r.Component(prefix+"*", handler)
}

// HandleInteraction routes an interaction to the appropriate handler.
func (r *Router) HandleInteraction(ctx context.Context, interaction *Interaction) error {
	r.mu.RLock()
	responder := r.responder
	appID := r.appID
	r.mu.RUnlock()

	switch interaction.Type {
	case InteractionTypeApplicationCommand:
		return r.handleCommand(ctx, interaction, responder, appID)
	case InteractionTypeAutocomplete:
		return r.handleAutocomplete(ctx, interaction, responder, appID)
	case InteractionTypeMessageComponent:
		return r.handleComponent(ctx, interaction, responder, appID)
	case InteractionTypeModalSubmit:
		return r.handleModalSubmit(ctx, interaction, responder, appID)
	default:
		return fmt.Errorf("commands: unknown interaction type %d", interaction.Type)
	}
}

func (r *Router) handleCommand(ctx context.Context, interaction *Interaction, responder Responder, appID string) error {
	if interaction.Data == nil {
		return fmt.Errorf("commands: command interaction missing data")
	}
	cmdName := interaction.Data.Name

	r.mu.RLock()
	cmd, ok := r.commands[cmdName]
	middleware := r.middleware
	r.mu.RUnlock()

	if !ok {
		// Check groups for subcommand
		if len(interaction.Data.Options) > 0 {
			for _, opt := range interaction.Data.Options {
				if opt.Type == int32(OptionTypeSubCommandGroup) {
					r.mu.RLock()
					group, groupOk := r.groups[cmdName]
					r.mu.RUnlock()
					if groupOk {
						if subCmd, subOk := group.Commands[opt.Name]; subOk {
							cmd = subCmd
							middleware = append(middleware, group.Middleware...)
							ok = true
							break
						}
					}
				} else if opt.Type == int32(OptionTypeSubCommand) {
					r.mu.RLock()
					group, groupOk := r.groups[cmdName]
					r.mu.RUnlock()
					if groupOk {
						if subCmd, subOk := group.Commands[opt.Name]; subOk {
							cmd = subCmd
							ok = true
							break
						}
					}
				}
			}
		}
	}

	if !ok {
		log.Printf("commands: unknown command %q", cmdName)
		return fmt.Errorf("commands: unknown command %q", cmdName)
	}

	// Build middleware chain
	if mp, ok := cmd.(MiddlewareProvider); ok {
		middleware = append(middleware, mp.Middleware()...)
	}

	// Create command context
	cmdCtx := r.createCommandContext(ctx, interaction, responder, appID)

	// Execute with middleware
	handler := cmd.Execute
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	return handler(cmdCtx)
}

func (r *Router) handleAutocomplete(ctx context.Context, interaction *Interaction, responder Responder, appID string) error {
	if interaction.Data == nil {
		return fmt.Errorf("commands: autocomplete interaction missing data")
	}
	cmdName := interaction.Data.Name

	r.mu.RLock()
	cmd, ok := r.commands[cmdName]
	r.mu.RUnlock()

	if !ok {
		return fmt.Errorf("commands: unknown command for autocomplete %q", cmdName)
	}

	ac, ok := cmd.(AutocompleteHandler)
	if !ok {
		return fmt.Errorf("commands: command %q does not support autocomplete", cmdName)
	}

	// Find focused option
	focused := ""
	for _, opt := range interaction.Data.Options {
		if opt.Focused {
			focused = opt.Name
			break
		}
	}

	cmdCtx := r.createCommandContext(ctx, interaction, responder, appID)
	cmdCtx.FocusedOption = focused

	choices, err := ac.Autocomplete(cmdCtx, focused)
	if err != nil {
		return err
	}

	return cmdCtx.RespondAutocomplete(choices)
}

func (r *Router) handleComponent(ctx context.Context, interaction *Interaction, responder Responder, appID string) error {
	if interaction.Data == nil {
		return fmt.Errorf("commands: component interaction missing data")
	}
	customID := interaction.Data.CustomID

	r.mu.RLock()
	handler, ok := r.components[customID]
	if !ok {
		// Try prefix matching
		for pattern, h := range r.components {
			if len(pattern) > 0 && pattern[len(pattern)-1] == '*' {
				prefix := pattern[:len(pattern)-1]
				if len(customID) >= len(prefix) && customID[:len(prefix)] == prefix {
					handler = h
					ok = true
					break
				}
			}
		}
	}
	middleware := r.componentMiddleware
	r.mu.RUnlock()

	if !ok {
		log.Printf("commands: unknown component %q", customID)
		return fmt.Errorf("commands: unknown component %q", customID)
	}

	// Apply component middleware chain
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	compCtx := r.createComponentContext(ctx, interaction, responder, appID)
	return handler(compCtx)
}

func (r *Router) handleModalSubmit(ctx context.Context, interaction *Interaction, responder Responder, appID string) error {
	if interaction.Data == nil {
		return fmt.Errorf("commands: modal interaction missing data")
	}
	customID := interaction.Data.CustomID

	r.mu.RLock()
	handler, ok := r.components[customID]
	if !ok {
		// Try prefix matching
		for pattern, h := range r.components {
			if len(pattern) > 0 && pattern[len(pattern)-1] == '*' {
				prefix := pattern[:len(pattern)-1]
				if len(customID) >= len(prefix) && customID[:len(prefix)] == prefix {
					handler = h
					ok = true
					break
				}
			}
		}
	}
	middleware := r.componentMiddleware
	r.mu.RUnlock()

	if !ok {
		log.Printf("commands: unknown modal %q", customID)
		return fmt.Errorf("commands: unknown modal %q", customID)
	}

	// Apply component middleware chain
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	compCtx := r.createComponentContext(ctx, interaction, responder, appID)
	return handler(compCtx)
}

func (r *Router) createCommandContext(ctx context.Context, interaction *Interaction, responder Responder, appID string) *Context {
	cmdCtx := &Context{
		Context:          ctx,
		InteractionID:    interaction.ID,
		InteractionToken: interaction.Token,
		InteractionType:  interaction.Type,
		CommandName:      interaction.Data.Name,
		CommandID:        interaction.Data.ID,
		Options:          parseOptions(interaction),
		User:             interaction.User,
		Member:           interaction.Member,
		GuildID:          interaction.GuildID,
		ChannelID:        interaction.ChannelID,
		Locale:           interaction.Locale,
		GuildLocale:      interaction.GuildLocale,
		services:         r.services,
	}

	// Set user from member if not directly set
	if cmdCtx.User == nil && cmdCtx.Member != nil && cmdCtx.Member.User != nil {
		if user, ok := cmdCtx.Member.User.(*models.User); ok {
			cmdCtx.User = user
		}
	}

	// Parse member permissions
	if cmdCtx.Member != nil && cmdCtx.Member.Permissions != "" {
		cmdCtx.permissions = parsePermissions(cmdCtx.Member.Permissions)
	}

	// Parse bot's app permissions
	if interaction.AppPermissions != "" {
		cmdCtx.AppPermissions = parsePermissions(interaction.AppPermissions)
	}

	// Set target for context menu commands
	if interaction.Data.TargetID != "" {
		if interaction.Data.Resolved != nil {
			if user, ok := interaction.Data.Resolved.Users[interaction.Data.TargetID]; ok {
				cmdCtx.TargetUser = user
			}
			if msg, ok := interaction.Data.Resolved.Messages[interaction.Data.TargetID]; ok {
				cmdCtx.TargetMessage = msg
			}
		}
	}

	cmdCtx.SetResponder(responder, appID)
	return cmdCtx
}

// parsePermissions converts a permission string to types.Permissions.
func parsePermissions(perms string) types.Permissions {
	var p uint64
	for _, c := range perms {
		if c >= '0' && c <= '9' {
			p = p*10 + uint64(c-'0')
		}
	}
	return types.Permissions(p)
}

func (r *Router) createComponentContext(ctx context.Context, interaction *Interaction, responder Responder, appID string) *ComponentContext {
	compCtx := &ComponentContext{
		Context:          ctx,
		InteractionID:    interaction.ID,
		InteractionToken: interaction.Token,
		CustomID:         interaction.Data.CustomID,
		ComponentType:    types.ComponentType(interaction.Data.ComponentType),
		Values:           interaction.Data.Values,
		User:             interaction.User,
		Member:           interaction.Member,
		GuildID:          interaction.GuildID,
		ChannelID:        interaction.ChannelID,
		Locale:           interaction.Locale,
		GuildLocale:      interaction.GuildLocale,
		Message:          interaction.Message,
		services:         r.services,
	}

	// Set user from member if not directly set
	if compCtx.User == nil && compCtx.Member != nil && compCtx.Member.User != nil {
		if user, ok := compCtx.Member.User.(*models.User); ok {
			compCtx.User = user
		}
	}

	// Parse resolved data for select menu interactions
	if interaction.Data.Resolved != nil {
		compCtx.Resolved = &ResolvedData{
			Users:       interaction.Data.Resolved.Users,
			Members:     interaction.Data.Resolved.Members,
			Roles:       interaction.Data.Resolved.Roles,
			Channels:    interaction.Data.Resolved.Channels,
			Messages:    interaction.Data.Resolved.Messages,
			Attachments: interaction.Data.Resolved.Attachments,
		}
	}

	// Parse modal values for modal submit interactions
	if interaction.Type == InteractionTypeModalSubmit && len(interaction.Data.Components) > 0 {
		compCtx.ModalValues = parseModalComponents(interaction.Data.Components)
	}

	// Parse bot's app permissions
	if interaction.AppPermissions != "" {
		compCtx.AppPermissions = parsePermissions(interaction.AppPermissions)
	}

	compCtx.SetResponder(responder, appID)
	return compCtx
}

// parseModalComponents extracts text input values from modal submission components.
func parseModalComponents(components []json.RawMessage) map[string]string {
	values := make(map[string]string)

	for _, row := range components {
		var actionRow struct {
			Type       int `json:"type"`
			Components []struct {
				Type     int    `json:"type"`
				CustomID string `json:"custom_id"`
				Value    string `json:"value"`
			} `json:"components"`
		}
		if err := json.Unmarshal(row, &actionRow); err != nil {
			continue
		}
		for _, comp := range actionRow.Components {
			if comp.CustomID != "" {
				values[comp.CustomID] = comp.Value
			}
		}
	}

	return values
}

func parseOptions(interaction *Interaction) *OptionMap {
	optMap := NewOptionMap()

	if interaction.Data.Options == nil {
		return optMap
	}

	// Recursively parse options (handling subcommands)
	var parseOpts func(opts []InteractionOption)
	parseOpts = func(opts []InteractionOption) {
		for _, opt := range opts {
			switch OptionType(opt.Type) {
			case OptionTypeSubCommand, OptionTypeSubCommandGroup:
				if opt.Options != nil {
					parseOpts(opt.Options)
				}
			default:
				optMap.Set(opt.Name, opt.Value)
			}
		}
	}

	parseOpts(interaction.Data.Options)

	// Set resolved data
	if interaction.Data.Resolved != nil {
		optMap.SetResolved(&ResolvedData{
			Users:       interaction.Data.Resolved.Users,
			Members:     interaction.Data.Resolved.Members,
			Roles:       interaction.Data.Resolved.Roles,
			Channels:    interaction.Data.Resolved.Channels,
			Messages:    interaction.Data.Resolved.Messages,
			Attachments: interaction.Data.Resolved.Attachments,
		})
	}

	return optMap
}

// Sync syncs all registered commands with Discord.
func (r *Router) Sync(ctx context.Context) error {
	r.mu.RLock()
	syncer := r.syncer
	appID := r.appID
	commands := r.commands
	groups := r.groups
	r.mu.RUnlock()

	if syncer == nil {
		return fmt.Errorf("commands: no syncer configured")
	}

	// Build application commands
	appCmds := make([]ApplicationCommandCreate, 0, len(commands)+len(groups))

	for _, cmd := range commands {
		appCmd := commandToApplicationCommand(cmd)
		appCmds = append(appCmds, appCmd)
	}

	for _, group := range groups {
		appCmd := groupToApplicationCommand(group)
		appCmds = append(appCmds, appCmd)
	}

	return syncer.SyncCommands(ctx, appID, appCmds)
}

// SyncGuild syncs all registered commands with a specific guild.
func (r *Router) SyncGuild(ctx context.Context, guildID string) error {
	r.mu.RLock()
	syncer := r.syncer
	appID := r.appID
	commands := r.commands
	groups := r.groups
	r.mu.RUnlock()

	if syncer == nil {
		return fmt.Errorf("commands: no syncer configured")
	}

	appCmds := make([]ApplicationCommandCreate, 0, len(commands)+len(groups))

	for _, cmd := range commands {
		appCmd := commandToApplicationCommand(cmd)
		appCmds = append(appCmds, appCmd)
	}

	for _, group := range groups {
		appCmd := groupToApplicationCommand(group)
		appCmds = append(appCmds, appCmd)
	}

	return syncer.SyncGuildCommands(ctx, appID, guildID, appCmds)
}

// Interaction represents a Discord interaction.
type Interaction struct {
	ID             string             `json:"id"`
	ApplicationID  string             `json:"application_id"`
	Type           InteractionType    `json:"type"`
	Data           *InteractionData   `json:"data,omitempty"`
	GuildID        string             `json:"guild_id,omitempty"`
	ChannelID      string             `json:"channel_id,omitempty"`
	Member         *InteractionMember `json:"member,omitempty"`
	User           *models.User       `json:"user,omitempty"`
	Token          string             `json:"token"`
	Version        int                `json:"version"`
	Message        *models.Message    `json:"message,omitempty"`
	AppPermissions string             `json:"app_permissions,omitempty"`
	Locale         string             `json:"locale,omitempty"`
	GuildLocale    string             `json:"guild_locale,omitempty"`
}

// InteractionMember represents a guild member in an interaction context.
// This differs from models.GuildMember as it includes computed permissions.
type InteractionMember struct {
	Avatar                     *string    `json:"avatar,omitempty"`
	Banner                     *string    `json:"banner,omitempty"`
	CommunicationDisabledUntil *string    `json:"communication_disabled_until,omitempty"`
	Deaf                       bool       `json:"deaf"`
	Flags                      int32      `json:"flags"`
	JoinedAt                   string     `json:"joined_at"`
	Mute                       bool       `json:"mute"`
	Nick                       *string    `json:"nick,omitempty"`
	Pending                    bool       `json:"pending"`
	PremiumSince               *string    `json:"premium_since,omitempty"`
	Roles                      []string   `json:"roles"`
	User                       any        `json:"user"`
	Permissions                string     `json:"permissions,omitempty"`
}

// InteractionData contains the data for an interaction.
type InteractionData struct {
	ID            string                   `json:"id,omitempty"`
	Name          string                   `json:"name,omitempty"`
	Type          int32                    `json:"type,omitempty"`
	Resolved      *ResolvedInteractionData `json:"resolved,omitempty"`
	Options       []InteractionOption      `json:"options,omitempty"`
	GuildID       string                   `json:"guild_id,omitempty"`
	TargetID      string                   `json:"target_id,omitempty"`
	CustomID      string                   `json:"custom_id,omitempty"`
	ComponentType int32                    `json:"component_type,omitempty"`
	Values        []string                 `json:"values,omitempty"`
	Components    []json.RawMessage        `json:"components,omitempty"`
}

// InteractionOption represents an option in an interaction.
type InteractionOption struct {
	Name    string              `json:"name"`
	Type    int32               `json:"type"`
	Value   any                 `json:"value,omitempty"`
	Options []InteractionOption `json:"options,omitempty"`
	Focused bool                `json:"focused,omitempty"`
}

// ResolvedInteractionData contains resolved entities.
type ResolvedInteractionData struct {
	Users       map[string]*models.User         `json:"users,omitempty"`
	Members     map[string]*models.GuildMember  `json:"members,omitempty"`
	Roles       map[string]*models.GuildRole    `json:"roles,omitempty"`
	Channels    map[string]*models.GuildChannel `json:"channels,omitempty"`
	Messages    map[string]*models.Message      `json:"messages,omitempty"`
	Attachments map[string]*models.Attachment   `json:"attachments,omitempty"`
}

// ApplicationCommandCreate represents a command to create with Discord.
type ApplicationCommandCreate struct {
	Name                     string                     `json:"name"`
	Description              string                     `json:"description,omitempty"`
	Type                     int32                      `json:"type,omitempty"`
	Options                  []ApplicationCommandOption `json:"options,omitempty"`
	DefaultMemberPermissions *string                    `json:"default_member_permissions,omitempty"`
	DMPermission             *bool                      `json:"dm_permission,omitempty"`
	Contexts                 []int32                    `json:"contexts,omitempty"`
	IntegrationTypes         []int32                    `json:"integration_types,omitempty"`
	NSFW                     bool                       `json:"nsfw,omitempty"`
}

// ApplicationCommandOption represents an option for an application command.
type ApplicationCommandOption struct {
	Name         string                     `json:"name"`
	Description  string                     `json:"description"`
	Type         int32                      `json:"type"`
	Required     bool                       `json:"required,omitempty"`
	Choices      []ApplicationCommandChoice `json:"choices,omitempty"`
	Options      []ApplicationCommandOption `json:"options,omitempty"`
	ChannelTypes []int32                    `json:"channel_types,omitempty"`
	MinValue     *float64                   `json:"min_value,omitempty"`
	MaxValue     *float64                   `json:"max_value,omitempty"`
	MinLength    *int32                     `json:"min_length,omitempty"`
	MaxLength    *int32                     `json:"max_length,omitempty"`
	Autocomplete bool                       `json:"autocomplete,omitempty"`
}

// ApplicationCommandChoice represents a choice for an option.
type ApplicationCommandChoice struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

func commandToApplicationCommand(cmd Command) ApplicationCommandCreate {
	appCmd := ApplicationCommandCreate{
		Name:        cmd.Name(),
		Description: cmd.Description(),
		Type:        int32(cmd.Type()),
	}

	if op, ok := cmd.(OptionProvider); ok {
		opts := op.Options()
		appCmd.Options = make([]ApplicationCommandOption, len(opts))
		for i, opt := range opts {
			appCmd.Options[i] = optionToApplicationOption(opt)
		}
	}

	if sp, ok := cmd.(SubcommandProvider); ok {
		subs := sp.Subcommands()
		for _, sub := range subs {
			subOpt := ApplicationCommandOption{
				Name:        sub.Name(),
				Description: sub.Description(),
				Type:        int32(OptionTypeSubCommand),
			}
			if op, ok := sub.(OptionProvider); ok {
				opts := op.Options()
				subOpt.Options = make([]ApplicationCommandOption, len(opts))
				for i, opt := range opts {
					subOpt.Options[i] = optionToApplicationOption(opt)
				}
			}
			appCmd.Options = append(appCmd.Options, subOpt)
		}
	}

	return appCmd
}

func groupToApplicationCommand(group *CommandGroup) ApplicationCommandCreate {
	appCmd := ApplicationCommandCreate{
		Name:        group.Name,
		Description: group.Description,
		Type:        int32(CommandTypeChatInput),
	}

	for _, cmd := range group.Commands {
		subOpt := ApplicationCommandOption{
			Name:        cmd.Name(),
			Description: cmd.Description(),
			Type:        int32(OptionTypeSubCommand),
		}
		if op, ok := cmd.(OptionProvider); ok {
			opts := op.Options()
			subOpt.Options = make([]ApplicationCommandOption, len(opts))
			for i, opt := range opts {
				subOpt.Options[i] = optionToApplicationOption(opt)
			}
		}
		appCmd.Options = append(appCmd.Options, subOpt)
	}

	return appCmd
}

func optionToApplicationOption(opt Option) ApplicationCommandOption {
	appOpt := ApplicationCommandOption{
		Name:         opt.Name,
		Description:  opt.Description,
		Type:         int32(opt.Type),
		Required:     opt.Required,
		ChannelTypes: opt.ChannelTypes,
		MinValue:     opt.MinValue,
		MaxValue:     opt.MaxValue,
		MinLength:    opt.MinLength,
		MaxLength:    opt.MaxLength,
		Autocomplete: opt.Autocomplete,
	}

	if len(opt.Choices) > 0 {
		appOpt.Choices = make([]ApplicationCommandChoice, len(opt.Choices))
		for i, c := range opt.Choices {
			appOpt.Choices[i] = ApplicationCommandChoice{Name: c.Name, Value: c.Value}
		}
	}

	if len(opt.Options) > 0 {
		appOpt.Options = make([]ApplicationCommandOption, len(opt.Options))
		for i, o := range opt.Options {
			appOpt.Options[i] = optionToApplicationOption(o)
		}
	}

	return appOpt
}

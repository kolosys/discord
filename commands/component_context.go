package commands

import (
	"context"
	"fmt"

	"github.com/kolosys/discord/models"
	"github.com/kolosys/discord/types"
)

// ComponentContext represents the execution context for a component interaction
// (button clicks, select menu selections, and modal submissions).
//
// # Response Deadlines
//
// Discord requires all interactions to be acknowledged within 3 seconds.
// For component interactions, you can:
//   - Reply() or ReplyEphemeral() - send a new message
//   - Update() - edit the message containing the component
//   - DeferUpdate() - acknowledge without visible loading (for background processing)
//   - Defer() - show loading indicator for longer operations
//
// # Token Lifetime
//
// The interaction token expires after 15 minutes. After expiration,
// Followup(), Edit(), and Delete() calls will fail.
//
// # Modal Submissions
//
// For modal submissions, use ModalValue(customID) to retrieve text input values.
type ComponentContext struct {
	context.Context

	// Interaction data
	InteractionID    string
	InteractionToken string

	// Component data
	CustomID      string
	ComponentType types.ComponentType
	Values        []string // For select menus

	// User and member data
	User   *models.User
	Member *InteractionMember

	// Location data
	GuildID   string
	ChannelID string

	// Locale data
	Locale      string
	GuildLocale string

	// AppPermissions contains the bot's permissions in the channel where the interaction occurred.
	AppPermissions types.Permissions

	// The message the component is attached to
	Message *models.Message

	// Resolved data from select menu interactions
	Resolved *ResolvedData

	// Modal input values (for modal submissions)
	ModalValues map[string]string

	// Internal
	responder Responder
	responded bool
	deferred  bool
	appID     string
	services  *ServiceRegistry
}

// ComponentHandlerFunc is the signature for component handler functions.
type ComponentHandlerFunc func(ctx *ComponentContext) error

// SetResponder sets the responder for this context.
func (c *ComponentContext) SetResponder(r Responder, appID string) {
	c.responder = r
	c.appID = appID
}

// Reply sends a new message response.
func (c *ComponentContext) Reply(content string) error {
	return c.ReplyComplex(&InteractionResponseData{Content: content})
}

// ReplyEphemeral sends an ephemeral response.
func (c *ComponentContext) ReplyEphemeral(content string) error {
	return c.ReplyComplex(&InteractionResponseData{
		Content: content,
		Flags:   MessageFlagEphemeral,
	})
}

// ReplyEmbed sends an embed response.
func (c *ComponentContext) ReplyEmbed(embed *Embed) error {
	return c.ReplyComplex(&InteractionResponseData{Embeds: []*Embed{embed}})
}

// ReplyComponents sends a message with components.
func (c *ComponentContext) ReplyComponents(content string, comps ...any) error {
	return c.ReplyComplex(&InteractionResponseData{
		Content:    content,
		Components: comps,
	})
}

// ReplyEphemeralComponents sends an ephemeral message with components.
func (c *ComponentContext) ReplyEphemeralComponents(content string, comps ...any) error {
	return c.ReplyComplex(&InteractionResponseData{
		Content:    content,
		Components: comps,
		Flags:      MessageFlagEphemeral,
	})
}

// ReplyComplex sends a complex response.
func (c *ComponentContext) ReplyComplex(data *InteractionResponseData) error {
	if c.responded {
		return fmt.Errorf("commands: interaction already responded to")
	}
	if c.responder == nil {
		return fmt.Errorf("commands: no responder set")
	}

	callbackType := CallbackTypeChannelMessage
	if c.deferred {
		edit := &MessageEdit{
			Content: &data.Content,
			Embeds:  data.Embeds,
		}
		c.responded = true
		return c.responder.EditInteractionResponse(c.Context, c.InteractionToken, edit)
	}

	resp := &InteractionResponse{
		Type: callbackType,
		Data: data,
	}

	c.responded = true
	return c.responder.RespondToInteraction(c.Context, c.InteractionID, c.InteractionToken, resp)
}

// Update updates the original message (the one with the component).
func (c *ComponentContext) Update(content string) error {
	return c.UpdateComplex(&InteractionResponseData{Content: content})
}

// UpdateEmbed updates the original message with an embed.
func (c *ComponentContext) UpdateEmbed(embed *Embed) error {
	return c.UpdateComplex(&InteractionResponseData{Embeds: []*Embed{embed}})
}

// UpdateComponents updates the original message with new content and components.
func (c *ComponentContext) UpdateComponents(content string, comps ...any) error {
	return c.UpdateComplex(&InteractionResponseData{
		Content:    content,
		Components: comps,
	})
}

// UpdateComplex updates the original message with full control.
func (c *ComponentContext) UpdateComplex(data *InteractionResponseData) error {
	if c.responded {
		return fmt.Errorf("commands: interaction already responded to")
	}
	if c.responder == nil {
		return fmt.Errorf("commands: no responder set")
	}

	resp := &InteractionResponse{
		Type: CallbackTypeUpdateMessage,
		Data: data,
	}

	c.responded = true
	return c.responder.RespondToInteraction(c.Context, c.InteractionID, c.InteractionToken, resp)
}

// Defer defers the response (shows loading).
func (c *ComponentContext) Defer() error {
	return c.deferWithFlags(0, false)
}

// DeferEphemeral defers with an ephemeral response.
func (c *ComponentContext) DeferEphemeral() error {
	return c.deferWithFlags(MessageFlagEphemeral, false)
}

// DeferUpdate defers and acknowledges (no loading state shown).
func (c *ComponentContext) DeferUpdate() error {
	return c.deferWithFlags(0, true)
}

func (c *ComponentContext) deferWithFlags(flags MessageFlags, update bool) error {
	if c.responded || c.deferred {
		return fmt.Errorf("commands: interaction already responded to or deferred")
	}
	if c.responder == nil {
		return fmt.Errorf("commands: no responder set")
	}

	callbackType := CallbackTypeDeferredChannelMessage
	if update {
		callbackType = CallbackTypeDeferredUpdateMessage
	}

	resp := &InteractionResponse{
		Type: callbackType,
		Data: &InteractionResponseData{Flags: flags},
	}

	c.deferred = true
	return c.responder.RespondToInteraction(c.Context, c.InteractionID, c.InteractionToken, resp)
}

// Followup sends a followup message.
func (c *ComponentContext) Followup(content string) (*models.Message, error) {
	return c.FollowupComplex(&MessageCreate{Content: content})
}

// FollowupEphemeral sends an ephemeral followup message.
func (c *ComponentContext) FollowupEphemeral(content string) (*models.Message, error) {
	return c.FollowupComplex(&MessageCreate{
		Content: content,
		Flags:   MessageFlagEphemeral,
	})
}

// FollowupComplex sends a complex followup message.
func (c *ComponentContext) FollowupComplex(msg *MessageCreate) (*models.Message, error) {
	if c.responder == nil {
		return nil, fmt.Errorf("commands: no responder set")
	}
	return c.responder.FollowupMessage(c.Context, c.InteractionToken, msg)
}

// Edit edits the deferred response.
func (c *ComponentContext) Edit(content string) error {
	return c.EditComplex(&MessageEdit{Content: &content})
}

// EditComplex edits the deferred response with full control.
func (c *ComponentContext) EditComplex(edit *MessageEdit) error {
	if c.responder == nil {
		return fmt.Errorf("commands: no responder set")
	}
	return c.responder.EditInteractionResponse(c.Context, c.InteractionToken, edit)
}

// Delete deletes the original response.
func (c *ComponentContext) Delete() error {
	if c.responder == nil {
		return fmt.Errorf("commands: no responder set")
	}
	return c.responder.DeleteInteractionResponse(c.Context, c.InteractionToken)
}

// Modal shows a modal dialog.
func (c *ComponentContext) Modal(customID, title string, comps ...any) error {
	if c.responded {
		return fmt.Errorf("commands: interaction already responded to")
	}
	if c.responder == nil {
		return fmt.Errorf("commands: no responder set")
	}

	resp := &InteractionResponse{
		Type: CallbackTypeModal,
		Data: &InteractionResponseData{
			CustomID:   customID,
			Title:      title,
			Components: comps,
		},
	}

	c.responded = true
	return c.responder.RespondToInteraction(c.Context, c.InteractionID, c.InteractionToken, resp)
}

// InGuild returns true if the interaction occurred in a guild.
func (c *ComponentContext) InGuild() bool {
	return c.GuildID != ""
}

// InDM returns true if the interaction occurred in a DM.
func (c *ComponentContext) InDM() bool {
	return c.GuildID == ""
}

// BotHasPermission checks if the bot has a specific permission in the current channel.
func (c *ComponentContext) BotHasPermission(perm types.Permissions) bool {
	return c.AppPermissions.Has(perm)
}

// SelectedValue returns the first selected value (for single-select).
func (c *ComponentContext) SelectedValue() string {
	if len(c.Values) > 0 {
		return c.Values[0]
	}
	return ""
}

// SelectedUser returns the first selected user from a user select menu.
func (c *ComponentContext) SelectedUser() *models.User {
	if c.Resolved == nil || len(c.Values) == 0 {
		return nil
	}
	return c.Resolved.Users[c.Values[0]]
}

// SelectedUsers returns all selected users from a user select menu.
func (c *ComponentContext) SelectedUsers() []*models.User {
	if c.Resolved == nil || len(c.Values) == 0 {
		return nil
	}
	users := make([]*models.User, 0, len(c.Values))
	for _, id := range c.Values {
		if u := c.Resolved.Users[id]; u != nil {
			users = append(users, u)
		}
	}
	return users
}

// SelectedMember returns the first selected member from a user select menu.
func (c *ComponentContext) SelectedMember() *models.GuildMember {
	if c.Resolved == nil || len(c.Values) == 0 {
		return nil
	}
	return c.Resolved.Members[c.Values[0]]
}

// SelectedRole returns the first selected role from a role select menu.
func (c *ComponentContext) SelectedRole() *models.GuildRole {
	if c.Resolved == nil || len(c.Values) == 0 {
		return nil
	}
	return c.Resolved.Roles[c.Values[0]]
}

// SelectedRoles returns all selected roles from a role select menu.
func (c *ComponentContext) SelectedRoles() []*models.GuildRole {
	if c.Resolved == nil || len(c.Values) == 0 {
		return nil
	}
	roles := make([]*models.GuildRole, 0, len(c.Values))
	for _, id := range c.Values {
		if r := c.Resolved.Roles[id]; r != nil {
			roles = append(roles, r)
		}
	}
	return roles
}

// SelectedChannel returns the first selected channel from a channel select menu.
func (c *ComponentContext) SelectedChannel() *models.GuildChannel {
	if c.Resolved == nil || len(c.Values) == 0 {
		return nil
	}
	return c.Resolved.Channels[c.Values[0]]
}

// SelectedChannels returns all selected channels from a channel select menu.
func (c *ComponentContext) SelectedChannels() []*models.GuildChannel {
	if c.Resolved == nil || len(c.Values) == 0 {
		return nil
	}
	channels := make([]*models.GuildChannel, 0, len(c.Values))
	for _, id := range c.Values {
		if ch := c.Resolved.Channels[id]; ch != nil {
			channels = append(channels, ch)
		}
	}
	return channels
}

// ModalValue returns a value from a modal text input.
func (c *ComponentContext) ModalValue(customID string) string {
	if c.ModalValues == nil {
		return ""
	}
	return c.ModalValues[customID]
}

// Responded returns true if a response has been sent.
func (c *ComponentContext) Responded() bool {
	return c.responded
}

// Deferred returns true if the interaction has been deferred.
func (c *ComponentContext) Deferred() bool {
	return c.deferred
}

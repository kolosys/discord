package commands

import (
	"context"
	"fmt"

	"github.com/kolosys/discord/models"
	"github.com/kolosys/discord/types"
)

// Responder is an interface for responding to interactions.
// This is implemented by the Bot to avoid circular imports.
type Responder interface {
	RespondToInteraction(ctx context.Context, interactionID, token string, response *InteractionResponse) error
	EditInteractionResponse(ctx context.Context, token string, edit *MessageEdit) error
	DeleteInteractionResponse(ctx context.Context, token string) error
	FollowupMessage(ctx context.Context, token string, message *MessageCreate) (*models.Message, error)
}

// Context represents the execution context for a command.
//
// # Response Deadlines
//
// Discord requires all interactions to be acknowledged within 3 seconds.
// If your command handler takes longer than 3 seconds, you MUST call Defer()
// or DeferEphemeral() first to acknowledge the interaction.
//
// After deferring, you have up to 15 minutes to send a response using
// Reply(), Edit(), or other response methods.
//
// # Token Lifetime
//
// The interaction token (used for followups and edits) expires after 15 minutes.
// After expiration, Followup(), Edit(), and Delete() calls will fail.
//
// # Example Usage
//
//	func handleLongTask(ctx *commands.Context) error {
//	    // Acknowledge immediately for long-running tasks
//	    if err := ctx.Defer(); err != nil {
//	        return err
//	    }
//
//	    // Perform long operation (up to 15 minutes)
//	    result := performLongOperation()
//
//	    // Edit the deferred response
//	    return ctx.Reply(result)
//	}
type Context struct {
	context.Context

	// Interaction data
	InteractionID    string
	InteractionToken string
	InteractionType  InteractionType

	// Command data
	CommandName string
	CommandID   string
	Options     *OptionMap

	// User and member data
	User   *models.User
	Member *InteractionMember

	// Location data
	GuildID   string
	ChannelID string

	// Locale data
	Locale      string
	GuildLocale string

	// Permissions (computed for the invoking user in the channel)
	permissions types.Permissions

	// AppPermissions contains the bot's permissions in the channel where the interaction occurred.
	AppPermissions types.Permissions

	// Target data (for context menu commands)
	TargetUser    *models.User
	TargetMessage *models.Message

	// Focused option (for autocomplete)
	FocusedOption string
	FocusedValue  string

	// Internal
	responder Responder
	responded bool
	deferred  bool
	appID     string
}

// InteractionType represents the type of interaction.
type InteractionType int32

const (
	InteractionTypePing               InteractionType = 1
	InteractionTypeApplicationCommand InteractionType = 2
	InteractionTypeMessageComponent   InteractionType = 3
	InteractionTypeAutocomplete       InteractionType = 4
	InteractionTypeModalSubmit        InteractionType = 5
)

// InteractionCallbackType represents the type of interaction response.
type InteractionCallbackType int32

const (
	CallbackTypePong                   InteractionCallbackType = 1
	CallbackTypeChannelMessage         InteractionCallbackType = 4
	CallbackTypeDeferredChannelMessage InteractionCallbackType = 5
	CallbackTypeDeferredUpdateMessage  InteractionCallbackType = 6
	CallbackTypeUpdateMessage          InteractionCallbackType = 7
	CallbackTypeAutocompleteResult     InteractionCallbackType = 8
	CallbackTypeModal                  InteractionCallbackType = 9
	CallbackTypePremiumRequired        InteractionCallbackType = 10
	CallbackTypeLaunchActivity         InteractionCallbackType = 12
)

// InteractionResponse represents a response to an interaction.
type InteractionResponse struct {
	Type InteractionCallbackType  `json:"type"`
	Data *InteractionResponseData `json:"data,omitempty"`
}

// InteractionResponseData contains the data for an interaction response.
type InteractionResponseData struct {
	TTS             bool                 `json:"tts,omitempty"`
	Content         string               `json:"content,omitempty"`
	Embeds          []*Embed             `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions     `json:"allowed_mentions,omitempty"`
	Flags           MessageFlags         `json:"flags,omitempty"`
	Components      []any                `json:"components,omitempty"`
	Attachments     []*models.Attachment `json:"attachments,omitempty"`
	Choices         []Choice             `json:"choices,omitempty"`   // For autocomplete
	CustomID        string               `json:"custom_id,omitempty"` // For modals
	Title           string               `json:"title,omitempty"`     // For modals
}

// MessageFlags represents message flags.
type MessageFlags int32

const (
	MessageFlagEphemeral             MessageFlags = 1 << 6
	MessageFlagSuppressEmbeds        MessageFlags = 1 << 2
	MessageFlagSuppressNotifications MessageFlags = 1 << 12
)

// AllowedMentions configures which mentions are parsed.
type AllowedMentions struct {
	Parse       []string `json:"parse,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	Users       []string `json:"users,omitempty"`
	RepliedUser bool     `json:"replied_user,omitempty"`
}

// Embed represents a Discord embed.
type Embed struct {
	Title       string          `json:"title,omitempty"`
	Type        string          `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	URL         string          `json:"url,omitempty"`
	Timestamp   string          `json:"timestamp,omitempty"`
	Color       int             `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       *EmbedVideo     `json:"video,omitempty"`
	Provider    *EmbedProvider  `json:"provider,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []*EmbedField   `json:"fields,omitempty"`
}

type EmbedFooter struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url,omitempty"`
}

type EmbedImage struct {
	URL string `json:"url"`
}

type EmbedThumbnail struct {
	URL string `json:"url"`
}

type EmbedVideo struct {
	URL string `json:"url,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

// MessageCreate represents a new message to send.
type MessageCreate struct {
	Content         string               `json:"content,omitempty"`
	TTS             bool                 `json:"tts,omitempty"`
	Embeds          []*Embed             `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions     `json:"allowed_mentions,omitempty"`
	Flags           MessageFlags         `json:"flags,omitempty"`
	Components      []any                `json:"components,omitempty"`
	Attachments     []*models.Attachment `json:"attachments,omitempty"`
}

// MessageEdit represents edits to a message.
type MessageEdit struct {
	Content         *string              `json:"content,omitempty"`
	Embeds          []*Embed             `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions     `json:"allowed_mentions,omitempty"`
	Components      []any                `json:"components,omitempty"`
	Attachments     []*models.Attachment `json:"attachments,omitempty"`
}

// SetResponder sets the responder for this context.
func (c *Context) SetResponder(r Responder, appID string) {
	c.responder = r
	c.appID = appID
}

// Reply sends a response to the interaction.
func (c *Context) Reply(content string) error {
	return c.ReplyComplex(&InteractionResponseData{Content: content})
}

// ReplyEmbed sends an embed response.
func (c *Context) ReplyEmbed(embed *Embed) error {
	return c.ReplyComplex(&InteractionResponseData{Embeds: []*Embed{embed}})
}

// ReplyEmbeds sends multiple embeds.
func (c *Context) ReplyEmbeds(embeds ...*Embed) error {
	return c.ReplyComplex(&InteractionResponseData{Embeds: embeds})
}

// ReplyEphemeral sends an ephemeral response (only visible to the user).
func (c *Context) ReplyEphemeral(content string) error {
	return c.ReplyComplex(&InteractionResponseData{
		Content: content,
		Flags:   MessageFlagEphemeral,
	})
}

// ReplyComponents sends a message with components (buttons, select menus, etc.).
func (c *Context) ReplyComponents(content string, components ...any) error {
	return c.ReplyComplex(&InteractionResponseData{
		Content:    content,
		Components: components,
	})
}

// ReplyEphemeralComponents sends an ephemeral message with components.
func (c *Context) ReplyEphemeralComponents(content string, components ...any) error {
	return c.ReplyComplex(&InteractionResponseData{
		Content:    content,
		Components: components,
		Flags:      MessageFlagEphemeral,
	})
}

// ReplyEmbedComponents sends an embed with components.
func (c *Context) ReplyEmbedComponents(embed *Embed, components ...any) error {
	return c.ReplyComplex(&InteractionResponseData{
		Embeds:     []*Embed{embed},
		Components: components,
	})
}

// ReplyComplex sends a complex response with full control.
func (c *Context) ReplyComplex(data *InteractionResponseData) error {
	if c.responded {
		return fmt.Errorf("commands: interaction already responded to")
	}
	if c.responder == nil {
		return fmt.Errorf("commands: no responder set")
	}

	callbackType := CallbackTypeChannelMessage
	if c.deferred {
		// If deferred, we need to edit the original response
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

// Defer acknowledges the interaction and shows a loading state.
// This MUST be called within 3 seconds if your handler needs more time.
// After deferring, use Reply() to send the actual response (within 15 minutes).
func (c *Context) Defer() error {
	return c.deferWithFlags(0)
}

// DeferEphemeral acknowledges with an ephemeral loading state.
// This MUST be called within 3 seconds if your handler needs more time.
// After deferring, use Reply() to send the actual response (within 15 minutes).
func (c *Context) DeferEphemeral() error {
	return c.deferWithFlags(MessageFlagEphemeral)
}

func (c *Context) deferWithFlags(flags MessageFlags) error {
	if c.responded || c.deferred {
		return fmt.Errorf("commands: interaction already responded to or deferred")
	}
	if c.responder == nil {
		return fmt.Errorf("commands: no responder set")
	}

	resp := &InteractionResponse{
		Type: CallbackTypeDeferredChannelMessage,
		Data: &InteractionResponseData{Flags: flags},
	}

	c.deferred = true
	return c.responder.RespondToInteraction(c.Context, c.InteractionID, c.InteractionToken, resp)
}

// Followup sends a followup message.
// Note: The interaction token expires after 15 minutes. Followup calls after
// expiration will fail.
func (c *Context) Followup(content string) (*models.Message, error) {
	return c.FollowupComplex(&MessageCreate{Content: content})
}

// FollowupEphemeral sends an ephemeral followup message.
func (c *Context) FollowupEphemeral(content string) (*models.Message, error) {
	return c.FollowupComplex(&MessageCreate{
		Content: content,
		Flags:   MessageFlagEphemeral,
	})
}

// FollowupComplex sends a complex followup message.
func (c *Context) FollowupComplex(msg *MessageCreate) (*models.Message, error) {
	if c.responder == nil {
		return nil, fmt.Errorf("commands: no responder set")
	}
	return c.responder.FollowupMessage(c.Context, c.InteractionToken, msg)
}

// Edit edits the original response.
func (c *Context) Edit(content string) error {
	return c.EditComplex(&MessageEdit{Content: &content})
}

// EditComplex edits the original response with full control.
func (c *Context) EditComplex(edit *MessageEdit) error {
	if c.responder == nil {
		return fmt.Errorf("commands: no responder set")
	}
	return c.responder.EditInteractionResponse(c.Context, c.InteractionToken, edit)
}

// Delete deletes the original response.
func (c *Context) Delete() error {
	if c.responder == nil {
		return fmt.Errorf("commands: no responder set")
	}
	return c.responder.DeleteInteractionResponse(c.Context, c.InteractionToken)
}

// RespondAutocomplete sends autocomplete choices.
func (c *Context) RespondAutocomplete(choices []Choice) error {
	if c.responded {
		return fmt.Errorf("commands: interaction already responded to")
	}
	if c.responder == nil {
		return fmt.Errorf("commands: no responder set")
	}

	resp := &InteractionResponse{
		Type: CallbackTypeAutocompleteResult,
		Data: &InteractionResponseData{Choices: choices},
	}

	c.responded = true
	return c.responder.RespondToInteraction(c.Context, c.InteractionID, c.InteractionToken, resp)
}

// Modal shows a modal dialog.
func (c *Context) Modal(customID, title string, components ...any) error {
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
			Components: components,
		},
	}

	c.responded = true
	return c.responder.RespondToInteraction(c.Context, c.InteractionID, c.InteractionToken, resp)
}

// InGuild returns true if the interaction occurred in a guild.
func (c *Context) InGuild() bool {
	return c.GuildID != ""
}

// InDM returns true if the interaction occurred in a DM.
func (c *Context) InDM() bool {
	return c.GuildID == ""
}

// MemberPermissions returns the member's permissions in the channel.
// Note: Permissions are provided in the interaction, not on the member object.
func (c *Context) MemberPermissions() types.Permissions {
	return c.permissions
}

// HasPermission checks if the member has a specific permission.
func (c *Context) HasPermission(perm types.Permissions) bool {
	return c.MemberPermissions().Has(perm)
}

// BotHasPermission checks if the bot has a specific permission in the current channel.
func (c *Context) BotHasPermission(perm types.Permissions) bool {
	return c.AppPermissions.Has(perm)
}

// String is a convenience method to get a string option.
func (c *Context) String(name string) string {
	return c.Options.String(name)
}

// Int is a convenience method to get an integer option.
func (c *Context) Int(name string) int64 {
	return c.Options.Int(name)
}

// Bool is a convenience method to get a boolean option.
func (c *Context) Bool(name string) bool {
	return c.Options.Bool(name)
}

// Float is a convenience method to get a float option.
func (c *Context) Float(name string) float64 {
	return c.Options.Float(name)
}

// GetUser is a convenience method to get a user option.
func (c *Context) GetUser(name string) *models.User {
	return c.Options.User(name)
}

// GetMember is a convenience method to get a member option.
func (c *Context) GetMember(name string) *models.GuildMember {
	return c.Options.Member(name)
}

// GetRole is a convenience method to get a role option.
func (c *Context) GetRole(name string) *models.GuildRole {
	return c.Options.Role(name)
}

// GetChannel is a convenience method to get a channel option.
func (c *Context) GetChannel(name string) *models.GuildChannel {
	return c.Options.Channel(name)
}

// GetAttachment is a convenience method to get an attachment option.
func (c *Context) GetAttachment(name string) *models.Attachment {
	return c.Options.Attachment(name)
}

// Responded returns true if a response has been sent.
func (c *Context) Responded() bool {
	return c.responded
}

// Deferred returns true if the interaction has been deferred.
func (c *Context) Deferred() bool {
	return c.deferred
}

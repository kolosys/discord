package events

import (
	"time"

	"github.com/kolosys/discord/models"
)

// Base contains common fields for all events.
type Base struct {
	ShardID int `json:"-"` // Shard that received this event
}

func (b *Base) setShardID(id int) { b.ShardID = id }

// ReadyEvent is dispatched when the client has completed the initial handshake.
type ReadyEvent struct {
	Base
	Version          int                 `json:"v"`
	User             *models.User        `json:"user"`
	Guilds           []UnavailableGuild  `json:"guilds"`
	SessionID        string              `json:"session_id"`
	ResumeGatewayURL string              `json:"resume_gateway_url"`
	Shard            *[2]int             `json:"shard,omitempty"`
	Application      *PartialApplication `json:"application"`
}

// UnavailableGuild represents a guild that is initially unavailable.
type UnavailableGuild struct {
	ID          string `json:"id"`
	Unavailable bool   `json:"unavailable"`
}

// PartialApplication contains partial application info.
type PartialApplication struct {
	ID    string `json:"id"`
	Flags int    `json:"flags"`
}

// GuildCreateEvent is dispatched when a guild becomes available or the bot joins a new guild.
type GuildCreateEvent struct {
	Base
	models.Guild
	JoinedAt    time.Time             `json:"joined_at,omitempty"`
	Large       bool                  `json:"large,omitempty"`
	Unavailable bool                  `json:"unavailable,omitempty"`
	MemberCount int                   `json:"member_count,omitempty"`
	Members     []models.GuildMember  `json:"members,omitempty"`
	Channels    []models.GuildChannel `json:"channels,omitempty"`
	Threads     []any                 `json:"threads,omitempty"`
	Presences   []any                 `json:"presences,omitempty"`
	VoiceStates []any                 `json:"voice_states,omitempty"`
}

// GuildUpdateEvent is dispatched when a guild is updated.
type GuildUpdateEvent struct {
	Base
	models.Guild
}

// GuildDeleteEvent is dispatched when a guild becomes unavailable or the bot leaves/is removed.
type GuildDeleteEvent struct {
	Base
	ID          string `json:"id"`
	Unavailable bool   `json:"unavailable,omitempty"`
}

// MessageCreateEvent is dispatched when a message is created.
type MessageCreateEvent struct {
	Base
	models.Message
	GuildID string              `json:"guild_id,omitempty"`
	Member  *models.GuildMember `json:"member,omitempty"`
}

// AuthorUser returns the message author as a typed User, or nil if unavailable.
func (e *MessageCreateEvent) AuthorUser() *models.User {
	if e.Author == nil {
		return nil
	}
	if u, ok := e.Author.(*models.User); ok {
		return u
	}
	// Handle case where Author is a map (from JSON unmarshaling)
	if m, ok := e.Author.(map[string]any); ok {
		user := &models.User{}
		if id, ok := m["id"].(string); ok {
			user.ID = id
		}
		if username, ok := m["username"].(string); ok {
			user.Username = username
		}
		if discriminator, ok := m["discriminator"].(string); ok {
			user.Discriminator = discriminator
		}
		if bot, ok := m["bot"].(bool); ok {
			user.Bot = bot
		}
		if globalName, ok := m["global_name"].(string); ok {
			user.GlobalName = &globalName
		}
		if avatar, ok := m["avatar"].(string); ok {
			user.Avatar = &avatar
		}
		return user
	}
	return nil
}

// IsBot returns true if the message was sent by a bot.
func (e *MessageCreateEvent) IsBot() bool {
	user := e.AuthorUser()
	return user != nil && user.Bot
}

// IsSystem returns true if the message was sent by the Discord system.
func (e *MessageCreateEvent) IsSystem() bool {
	user := e.AuthorUser()
	return user != nil && user.System
}

// IsHuman returns true if the message was sent by a human user (not a bot or system).
func (e *MessageCreateEvent) IsHuman() bool {
	user := e.AuthorUser()
	return user != nil && !user.Bot && !user.System
}

// MessageUpdateEvent is dispatched when a message is updated.
type MessageUpdateEvent struct {
	Base
	ID        string     `json:"id"`
	ChannelID string     `json:"channel_id"`
	GuildID   string     `json:"guild_id,omitempty"`
	Content   *string    `json:"content,omitempty"`
	EditedAt  *time.Time `json:"edited_timestamp,omitempty"`
	// Partial message fields - may not all be present
	Author      *models.User `json:"author,omitempty"`
	Embeds      []any        `json:"embeds,omitempty"`
	Attachments []any        `json:"attachments,omitempty"`
}

// MessageDeleteEvent is dispatched when a message is deleted.
type MessageDeleteEvent struct {
	Base
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id,omitempty"`
}

// MessageDeleteBulkEvent is dispatched when multiple messages are deleted at once.
type MessageDeleteBulkEvent struct {
	Base
	IDs       []string `json:"ids"`
	ChannelID string   `json:"channel_id"`
	GuildID   string   `json:"guild_id,omitempty"`
}

// MessageReactionAddEvent is dispatched when a reaction is added to a message.
type MessageReactionAddEvent struct {
	Base
	UserID          string              `json:"user_id"`
	ChannelID       string              `json:"channel_id"`
	MessageID       string              `json:"message_id"`
	GuildID         string              `json:"guild_id,omitempty"`
	Member          *models.GuildMember `json:"member,omitempty"`
	Emoji           ReactionEmoji       `json:"emoji"`
	MessageAuthorID string              `json:"message_author_id,omitempty"`
}

// MessageReactionRemoveEvent is dispatched when a reaction is removed from a message.
type MessageReactionRemoveEvent struct {
	Base
	UserID    string        `json:"user_id"`
	ChannelID string        `json:"channel_id"`
	MessageID string        `json:"message_id"`
	GuildID   string        `json:"guild_id,omitempty"`
	Emoji     ReactionEmoji `json:"emoji"`
}

// ReactionEmoji represents an emoji in a reaction.
type ReactionEmoji struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Animated bool   `json:"animated,omitempty"`
}

// GuildMemberAddEvent is dispatched when a user joins a guild.
type GuildMemberAddEvent struct {
	Base
	models.GuildMember
	GuildID string `json:"guild_id"`
}

// GuildMemberUpdateEvent is dispatched when a guild member is updated.
type GuildMemberUpdateEvent struct {
	Base
	GuildID      string       `json:"guild_id"`
	Roles        []string     `json:"roles"`
	User         *models.User `json:"user"`
	Nick         *string      `json:"nick,omitempty"`
	JoinedAt     time.Time    `json:"joined_at,omitempty"`
	PremiumSince *time.Time   `json:"premium_since,omitempty"`
	Deaf         bool         `json:"deaf,omitempty"`
	Mute         bool         `json:"mute,omitempty"`
	Pending      bool         `json:"pending,omitempty"`
}

// GuildMemberRemoveEvent is dispatched when a user leaves or is removed from a guild.
type GuildMemberRemoveEvent struct {
	Base
	GuildID string       `json:"guild_id"`
	User    *models.User `json:"user"`
}

// GuildMembersChunkEvent is dispatched in response to RequestGuildMembers.
type GuildMembersChunkEvent struct {
	Base
	GuildID    string               `json:"guild_id"`
	Members    []models.GuildMember `json:"members"`
	ChunkIndex int                  `json:"chunk_index"`
	ChunkCount int                  `json:"chunk_count"`
	NotFound   []string             `json:"not_found,omitempty"`
	Presences  []any                `json:"presences,omitempty"`
	Nonce      string               `json:"nonce,omitempty"`
}

// ChannelCreateEvent is dispatched when a channel is created.
type ChannelCreateEvent struct {
	Base
	models.GuildChannel
}

// ChannelUpdateEvent is dispatched when a channel is updated.
type ChannelUpdateEvent struct {
	Base
	models.GuildChannel
}

// ChannelDeleteEvent is dispatched when a channel is deleted.
type ChannelDeleteEvent struct {
	Base
	models.GuildChannel
}

// InteractionCreateEvent is dispatched when a user interacts with the application.
type InteractionCreateEvent struct {
	Base
	models.Interaction
}

// TypingStartEvent is dispatched when a user starts typing.
type TypingStartEvent struct {
	Base
	ChannelID string              `json:"channel_id"`
	GuildID   string              `json:"guild_id,omitempty"`
	UserID    string              `json:"user_id"`
	Timestamp int64               `json:"timestamp"`
	Member    *models.GuildMember `json:"member,omitempty"`
}

// VoiceStateUpdateEvent is dispatched when a user's voice state changes.
type VoiceStateUpdateEvent struct {
	Base
	GuildID    string              `json:"guild_id,omitempty"`
	ChannelID  *string             `json:"channel_id"`
	UserID     string              `json:"user_id"`
	Member     *models.GuildMember `json:"member,omitempty"`
	SessionID  string              `json:"session_id"`
	Deaf       bool                `json:"deaf"`
	Mute       bool                `json:"mute"`
	SelfDeaf   bool                `json:"self_deaf"`
	SelfMute   bool                `json:"self_mute"`
	SelfStream bool                `json:"self_stream,omitempty"`
	SelfVideo  bool                `json:"self_video"`
	Suppress   bool                `json:"suppress"`
}

// VoiceServerUpdateEvent is dispatched when a guild's voice server is updated.
type VoiceServerUpdateEvent struct {
	Base
	Token    string  `json:"token"`
	GuildID  string  `json:"guild_id"`
	Endpoint *string `json:"endpoint"`
}

// PresenceUpdateEvent is dispatched when a user's presence is updated.
type PresenceUpdateEvent struct {
	Base
	User         *models.User `json:"user"`
	GuildID      string       `json:"guild_id"`
	Status       string       `json:"status"`
	Activities   []Activity   `json:"activities"`
	ClientStatus ClientStatus `json:"client_status"`
}

// Activity represents a user activity.
type Activity struct {
	Name      string         `json:"name"`
	Type      int            `json:"type"`
	URL       *string        `json:"url,omitempty"`
	CreatedAt int64          `json:"created_at,omitempty"`
	State     *string        `json:"state,omitempty"`
	Details   *string        `json:"details,omitempty"`
	Emoji     *ReactionEmoji `json:"emoji,omitempty"`
}

// ClientStatus represents a user's client status across platforms.
type ClientStatus struct {
	Desktop string `json:"desktop,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
	Web     string `json:"web,omitempty"`
}

// InviteCreateEvent is dispatched when an invite is created.
type InviteCreateEvent struct {
	Base
	ChannelID string       `json:"channel_id"`
	Code      string       `json:"code"`
	CreatedAt time.Time    `json:"created_at"`
	GuildID   string       `json:"guild_id,omitempty"`
	Inviter   *models.User `json:"inviter,omitempty"`
	MaxAge    int          `json:"max_age"`
	MaxUses   int          `json:"max_uses"`
	Temporary bool         `json:"temporary"`
	Uses      int          `json:"uses"`
}

// InviteDeleteEvent is dispatched when an invite is deleted.
type InviteDeleteEvent struct {
	Base
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id,omitempty"`
	Code      string `json:"code"`
}

// GuildBanAddEvent is dispatched when a user is banned from a guild.
type GuildBanAddEvent struct {
	Base
	GuildID string       `json:"guild_id"`
	User    *models.User `json:"user"`
}

// GuildBanRemoveEvent is dispatched when a user is unbanned from a guild.
type GuildBanRemoveEvent struct {
	Base
	GuildID string       `json:"guild_id"`
	User    *models.User `json:"user"`
}

// GuildRoleCreateEvent is dispatched when a role is created.
type GuildRoleCreateEvent struct {
	Base
	GuildID string           `json:"guild_id"`
	Role    models.GuildRole `json:"role"`
}

// GuildRoleUpdateEvent is dispatched when a role is updated.
type GuildRoleUpdateEvent struct {
	Base
	GuildID string           `json:"guild_id"`
	Role    models.GuildRole `json:"role"`
}

// GuildRoleDeleteEvent is dispatched when a role is deleted.
type GuildRoleDeleteEvent struct {
	Base
	GuildID string `json:"guild_id"`
	RoleID  string `json:"role_id"`
}

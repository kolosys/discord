package events

// Type represents a Discord gateway event type.
type Type string

// Discord Gateway Event Types
const (
	// Connection Events
	Ready   Type = "READY"
	Resumed Type = "RESUMED"

	// Guild Events
	GuildCreate Type = "GUILD_CREATE"
	GuildUpdate Type = "GUILD_UPDATE"
	GuildDelete Type = "GUILD_DELETE"

	// Guild Member Events
	GuildMemberAdd    Type = "GUILD_MEMBER_ADD"
	GuildMemberUpdate Type = "GUILD_MEMBER_UPDATE"
	GuildMemberRemove Type = "GUILD_MEMBER_REMOVE"
	GuildMembersChunk Type = "GUILD_MEMBERS_CHUNK"

	// Guild Ban Events
	GuildBanAdd    Type = "GUILD_BAN_ADD"
	GuildBanRemove Type = "GUILD_BAN_REMOVE"

	// Guild Role Events
	GuildRoleCreate Type = "GUILD_ROLE_CREATE"
	GuildRoleUpdate Type = "GUILD_ROLE_UPDATE"
	GuildRoleDelete Type = "GUILD_ROLE_DELETE"

	// Guild Emoji/Sticker Events
	GuildEmojisUpdate   Type = "GUILD_EMOJIS_UPDATE"
	GuildStickersUpdate Type = "GUILD_STICKERS_UPDATE"

	// Guild Scheduled Event Events
	GuildScheduledEventCreate     Type = "GUILD_SCHEDULED_EVENT_CREATE"
	GuildScheduledEventUpdate     Type = "GUILD_SCHEDULED_EVENT_UPDATE"
	GuildScheduledEventDelete     Type = "GUILD_SCHEDULED_EVENT_DELETE"
	GuildScheduledEventUserAdd    Type = "GUILD_SCHEDULED_EVENT_USER_ADD"
	GuildScheduledEventUserRemove Type = "GUILD_SCHEDULED_EVENT_USER_REMOVE"

	// Channel Events
	ChannelCreate Type = "CHANNEL_CREATE"
	ChannelUpdate Type = "CHANNEL_UPDATE"
	ChannelDelete Type = "CHANNEL_DELETE"
	ChannelPinsUpdate Type = "CHANNEL_PINS_UPDATE"

	// Thread Events
	ThreadCreate      Type = "THREAD_CREATE"
	ThreadUpdate      Type = "THREAD_UPDATE"
	ThreadDelete      Type = "THREAD_DELETE"
	ThreadListSync    Type = "THREAD_LIST_SYNC"
	ThreadMemberUpdate Type = "THREAD_MEMBER_UPDATE"
	ThreadMembersUpdate Type = "THREAD_MEMBERS_UPDATE"

	// Message Events
	MessageCreate            Type = "MESSAGE_CREATE"
	MessageUpdate            Type = "MESSAGE_UPDATE"
	MessageDelete            Type = "MESSAGE_DELETE"
	MessageDeleteBulk        Type = "MESSAGE_DELETE_BULK"
	MessageReactionAdd       Type = "MESSAGE_REACTION_ADD"
	MessageReactionRemove    Type = "MESSAGE_REACTION_REMOVE"
	MessageReactionRemoveAll Type = "MESSAGE_REACTION_REMOVE_ALL"
	MessageReactionRemoveEmoji Type = "MESSAGE_REACTION_REMOVE_EMOJI"

	// Presence Events
	PresenceUpdate Type = "PRESENCE_UPDATE"
	TypingStart    Type = "TYPING_START"
	UserUpdate     Type = "USER_UPDATE"

	// Voice Events
	VoiceStateUpdate  Type = "VOICE_STATE_UPDATE"
	VoiceServerUpdate Type = "VOICE_SERVER_UPDATE"

	// Interaction Events
	InteractionCreate Type = "INTERACTION_CREATE"

	// Integration Events
	IntegrationCreate Type = "INTEGRATION_CREATE"
	IntegrationUpdate Type = "INTEGRATION_UPDATE"
	IntegrationDelete Type = "INTEGRATION_DELETE"

	// Webhook Events
	WebhooksUpdate Type = "WEBHOOKS_UPDATE"

	// Invite Events
	InviteCreate Type = "INVITE_CREATE"
	InviteDelete Type = "INVITE_DELETE"

	// Stage Events
	StageInstanceCreate Type = "STAGE_INSTANCE_CREATE"
	StageInstanceUpdate Type = "STAGE_INSTANCE_UPDATE"
	StageInstanceDelete Type = "STAGE_INSTANCE_DELETE"

	// Auto Moderation Events
	AutoModerationRuleCreate      Type = "AUTO_MODERATION_RULE_CREATE"
	AutoModerationRuleUpdate      Type = "AUTO_MODERATION_RULE_UPDATE"
	AutoModerationRuleDelete      Type = "AUTO_MODERATION_RULE_DELETE"
	AutoModerationActionExecution Type = "AUTO_MODERATION_ACTION_EXECUTION"

	// Entitlement Events
	EntitlementCreate Type = "ENTITLEMENT_CREATE"
	EntitlementUpdate Type = "ENTITLEMENT_UPDATE"
	EntitlementDelete Type = "ENTITLEMENT_DELETE"

	// Poll Events
	MessagePollVoteAdd    Type = "MESSAGE_POLL_VOTE_ADD"
	MessagePollVoteRemove Type = "MESSAGE_POLL_VOTE_REMOVE"
)

// String returns the string representation of the event type.
func (t Type) String() string {
	return string(t)
}

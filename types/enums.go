package types

// ChannelType represents the type of a Discord channel.
type ChannelType int

const (
	ChannelTypeGuildText          ChannelType = 0  // Text channel within a server
	ChannelTypeDM                 ChannelType = 1  // Direct message between users
	ChannelTypeGuildVoice         ChannelType = 2  // Voice channel within a server
	ChannelTypeGroupDM            ChannelType = 3  // Group direct message
	ChannelTypeGuildCategory      ChannelType = 4  // Category that contains channels
	ChannelTypeGuildAnnouncement  ChannelType = 5  // Announcement channel (news)
	ChannelTypeAnnouncementThread ChannelType = 10 // Thread in an announcement channel
	ChannelTypePublicThread       ChannelType = 11 // Public thread
	ChannelTypePrivateThread      ChannelType = 12 // Private thread
	ChannelTypeGuildStageVoice    ChannelType = 13 // Stage voice channel
	ChannelTypeGuildDirectory     ChannelType = 14 // Directory channel in a hub
	ChannelTypeGuildForum         ChannelType = 15 // Forum channel
	ChannelTypeGuildMedia         ChannelType = 16 // Media channel
)

func (c ChannelType) IsThread() bool {
	return c == ChannelTypeAnnouncementThread || c == ChannelTypePublicThread || c == ChannelTypePrivateThread
}

func (c ChannelType) IsVoice() bool {
	return c == ChannelTypeGuildVoice || c == ChannelTypeGuildStageVoice
}

func (c ChannelType) IsGuild() bool {
	return c != ChannelTypeDM && c != ChannelTypeGroupDM
}

// MessageType represents the type of a Discord message.
type MessageType int

const (
	MessageTypeDefault                                 MessageType = 0
	MessageTypeRecipientAdd                            MessageType = 1
	MessageTypeRecipientRemove                         MessageType = 2
	MessageTypeCall                                    MessageType = 3
	MessageTypeChannelNameChange                       MessageType = 4
	MessageTypeChannelIconChange                       MessageType = 5
	MessageTypeChannelPinnedMessage                    MessageType = 6
	MessageTypeUserJoin                                MessageType = 7
	MessageTypeGuildBoost                              MessageType = 8
	MessageTypeGuildBoostTier1                         MessageType = 9
	MessageTypeGuildBoostTier2                         MessageType = 10
	MessageTypeGuildBoostTier3                         MessageType = 11
	MessageTypeChannelFollowAdd                        MessageType = 12
	MessageTypeGuildDiscoveryDisqualified              MessageType = 14
	MessageTypeGuildDiscoveryRequalified               MessageType = 15
	MessageTypeGuildDiscoveryGracePeriodInitialWarning MessageType = 16
	MessageTypeGuildDiscoveryGracePeriodFinalWarning   MessageType = 17
	MessageTypeThreadCreated                           MessageType = 18
	MessageTypeReply                                   MessageType = 19
	MessageTypeChatInputCommand                        MessageType = 20
	MessageTypeThreadStarterMessage                    MessageType = 21
	MessageTypeGuildInviteReminder                     MessageType = 22
	MessageTypeContextMenuCommand                      MessageType = 23
	MessageTypeAutoModerationAction                    MessageType = 24
	MessageTypeRoleSubscriptionPurchase                MessageType = 25
	MessageTypeInteractionPremiumUpsell                MessageType = 26
	MessageTypeStageStart                              MessageType = 27
	MessageTypeStageEnd                                MessageType = 28
	MessageTypeStageSpeaker                            MessageType = 29
	MessageTypeStageTopic                              MessageType = 31
	MessageTypeGuildApplicationPremiumSubscription     MessageType = 32
	MessageTypePurchaseNotification                    MessageType = 44
	MessageTypePollResult                              MessageType = 46
)

// ComponentType represents the type of a message component.
type ComponentType int

const (
	ComponentTypeActionRow         ComponentType = 1
	ComponentTypeButton            ComponentType = 2
	ComponentTypeStringSelect      ComponentType = 3
	ComponentTypeTextInput         ComponentType = 4
	ComponentTypeUserSelect        ComponentType = 5
	ComponentTypeRoleSelect        ComponentType = 6
	ComponentTypeMentionableSelect ComponentType = 7
	ComponentTypeChannelSelect     ComponentType = 8
	ComponentTypeSection           ComponentType = 9
	ComponentTypeTextDisplay       ComponentType = 10
	ComponentTypeThumbnail         ComponentType = 11
	ComponentTypeMediaGallery      ComponentType = 12
	ComponentTypeFile              ComponentType = 13
	ComponentTypeSeparator         ComponentType = 14
	ComponentTypeContainer         ComponentType = 17
)

// ButtonStyle represents the visual style of a button.
type ButtonStyle int

const (
	ButtonStylePrimary   ButtonStyle = 1 // Blurple
	ButtonStyleSecondary ButtonStyle = 2 // Grey
	ButtonStyleSuccess   ButtonStyle = 3 // Green
	ButtonStyleDanger    ButtonStyle = 4 // Red
	ButtonStyleLink      ButtonStyle = 5 // Grey, navigates to URL
	ButtonStylePremium   ButtonStyle = 6 // Blurple, premium upsell
)

// TextInputStyle represents the visual style of a text input.
type TextInputStyle int

const (
	TextInputStyleShort     TextInputStyle = 1 // Single-line input
	TextInputStyleParagraph TextInputStyle = 2 // Multi-line input
)

// InteractionType represents the type of an interaction.
type InteractionType int

const (
	InteractionTypePing                           InteractionType = 1
	InteractionTypeApplicationCommand             InteractionType = 2
	InteractionTypeMessageComponent               InteractionType = 3
	InteractionTypeApplicationCommandAutocomplete InteractionType = 4
	InteractionTypeModalSubmit                    InteractionType = 5
)

// InteractionCallbackType represents the type of an interaction response.
type InteractionCallbackType int

const (
	InteractionCallbackTypePong                                 InteractionCallbackType = 1
	InteractionCallbackTypeChannelMessageWithSource             InteractionCallbackType = 4
	InteractionCallbackTypeDeferredChannelMessageWithSource     InteractionCallbackType = 5
	InteractionCallbackTypeDeferredUpdateMessage                InteractionCallbackType = 6
	InteractionCallbackTypeUpdateMessage                        InteractionCallbackType = 7
	InteractionCallbackTypeApplicationCommandAutocompleteResult InteractionCallbackType = 8
	InteractionCallbackTypeModal                                InteractionCallbackType = 9
	InteractionCallbackTypePremiumRequired                      InteractionCallbackType = 10
	InteractionCallbackTypeLaunchActivity                       InteractionCallbackType = 12
)

// ApplicationCommandType represents the type of an application command.
type ApplicationCommandType int

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = 1 // Slash command
	ApplicationCommandTypeUser      ApplicationCommandType = 2 // Context menu on user
	ApplicationCommandTypeMessage   ApplicationCommandType = 3 // Context menu on message
)

// ApplicationCommandOptionType represents the type of a command option.
type ApplicationCommandOptionType int

const (
	ApplicationCommandOptionTypeSubCommand      ApplicationCommandOptionType = 1
	ApplicationCommandOptionTypeSubCommandGroup ApplicationCommandOptionType = 2
	ApplicationCommandOptionTypeString          ApplicationCommandOptionType = 3
	ApplicationCommandOptionTypeInteger         ApplicationCommandOptionType = 4
	ApplicationCommandOptionTypeBoolean         ApplicationCommandOptionType = 5
	ApplicationCommandOptionTypeUser            ApplicationCommandOptionType = 6
	ApplicationCommandOptionTypeChannel         ApplicationCommandOptionType = 7
	ApplicationCommandOptionTypeRole            ApplicationCommandOptionType = 8
	ApplicationCommandOptionTypeMentionable     ApplicationCommandOptionType = 9
	ApplicationCommandOptionTypeNumber          ApplicationCommandOptionType = 10
	ApplicationCommandOptionTypeAttachment      ApplicationCommandOptionType = 11
)

// ActivityType represents the type of a user activity.
type ActivityType int

const (
	ActivityTypePlaying   ActivityType = 0 // Playing {name}
	ActivityTypeStreaming ActivityType = 1 // Streaming {details}
	ActivityTypeListening ActivityType = 2 // Listening to {name}
	ActivityTypeWatching  ActivityType = 3 // Watching {name}
	ActivityTypeCustom    ActivityType = 4 // {emoji} {name}
	ActivityTypeCompeting ActivityType = 5 // Competing in {name}
)

// StatusType represents a user's online status.
type StatusType string

const (
	StatusTypeOnline    StatusType = "online"
	StatusTypeDND       StatusType = "dnd"
	StatusTypeIdle      StatusType = "idle"
	StatusTypeInvisible StatusType = "invisible"
	StatusTypeOffline   StatusType = "offline"
)

// PremiumType represents a user's Nitro subscription type.
type PremiumType int

const (
	PremiumTypeNone         PremiumType = 0
	PremiumTypeNitroClassic PremiumType = 1
	PremiumTypeNitro        PremiumType = 2
	PremiumTypeNitroBasic   PremiumType = 3
)

// VerificationLevel represents a guild's verification requirements.
type VerificationLevel int

const (
	VerificationLevelNone     VerificationLevel = 0 // Unrestricted
	VerificationLevelLow      VerificationLevel = 1 // Must have verified email
	VerificationLevelMedium   VerificationLevel = 2 // Registered for 5+ minutes
	VerificationLevelHigh     VerificationLevel = 3 // Member for 10+ minutes
	VerificationLevelVeryHigh VerificationLevel = 4 // Must have verified phone
)

// ExplicitContentFilterLevel represents a guild's media content filter.
type ExplicitContentFilterLevel int

const (
	ExplicitContentFilterDisabled            ExplicitContentFilterLevel = 0 // No scanning
	ExplicitContentFilterMembersWithoutRoles ExplicitContentFilterLevel = 1 // Scan members without roles
	ExplicitContentFilterAllMembers          ExplicitContentFilterLevel = 2 // Scan all members
)

// MFALevel represents a guild's MFA requirement for moderators.
type MFALevel int

const (
	MFALevelNone     MFALevel = 0 // No MFA required
	MFALevelElevated MFALevel = 1 // MFA required for moderator actions
)

// NSFWLevel represents a guild's NSFW classification.
type NSFWLevel int

const (
	NSFWLevelDefault       NSFWLevel = 0
	NSFWLevelExplicit      NSFWLevel = 1
	NSFWLevelSafe          NSFWLevel = 2
	NSFWLevelAgeRestricted NSFWLevel = 3
)

// PremiumTier represents a guild's boost level.
type PremiumTier int

const (
	PremiumTierNone  PremiumTier = 0 // No boosts
	PremiumTierTier1 PremiumTier = 1 // Level 1 (2 boosts)
	PremiumTierTier2 PremiumTier = 2 // Level 2 (7 boosts)
	PremiumTierTier3 PremiumTier = 3 // Level 3 (14 boosts)
)

// DefaultMessageNotificationLevel represents guild default notification settings.
type DefaultMessageNotificationLevel int

const (
	DefaultMessageNotificationLevelAllMessages  DefaultMessageNotificationLevel = 0
	DefaultMessageNotificationLevelOnlyMentions DefaultMessageNotificationLevel = 1
)

// StickerFormatType represents the format of a sticker.
type StickerFormatType int

const (
	StickerFormatTypePNG    StickerFormatType = 1
	StickerFormatTypeAPNG   StickerFormatType = 2
	StickerFormatTypeLottie StickerFormatType = 3
	StickerFormatTypeGIF    StickerFormatType = 4
)

// StickerType represents the type of a sticker.
type StickerType int

const (
	StickerTypeStandard StickerType = 1 // Official Discord sticker
	StickerTypeGuild    StickerType = 2 // Guild sticker
)

// WebhookType represents the type of a webhook.
type WebhookType int

const (
	WebhookTypeIncoming        WebhookType = 1 // Incoming webhooks
	WebhookTypeChannelFollower WebhookType = 2 // Channel follower webhooks
	WebhookTypeApplication     WebhookType = 3 // Application webhooks
)

// AuditLogEvent represents the type of audit log event.
type AuditLogEvent int

const (
	AuditLogEventGuildUpdate                             AuditLogEvent = 1
	AuditLogEventChannelCreate                           AuditLogEvent = 10
	AuditLogEventChannelUpdate                           AuditLogEvent = 11
	AuditLogEventChannelDelete                           AuditLogEvent = 12
	AuditLogEventChannelOverwriteCreate                  AuditLogEvent = 13
	AuditLogEventChannelOverwriteUpdate                  AuditLogEvent = 14
	AuditLogEventChannelOverwriteDelete                  AuditLogEvent = 15
	AuditLogEventMemberKick                              AuditLogEvent = 20
	AuditLogEventMemberPrune                             AuditLogEvent = 21
	AuditLogEventMemberBanAdd                            AuditLogEvent = 22
	AuditLogEventMemberBanRemove                         AuditLogEvent = 23
	AuditLogEventMemberUpdate                            AuditLogEvent = 24
	AuditLogEventMemberRoleUpdate                        AuditLogEvent = 25
	AuditLogEventMemberMove                              AuditLogEvent = 26
	AuditLogEventMemberDisconnect                        AuditLogEvent = 27
	AuditLogEventBotAdd                                  AuditLogEvent = 28
	AuditLogEventRoleCreate                              AuditLogEvent = 30
	AuditLogEventRoleUpdate                              AuditLogEvent = 31
	AuditLogEventRoleDelete                              AuditLogEvent = 32
	AuditLogEventInviteCreate                            AuditLogEvent = 40
	AuditLogEventInviteUpdate                            AuditLogEvent = 41
	AuditLogEventInviteDelete                            AuditLogEvent = 42
	AuditLogEventWebhookCreate                           AuditLogEvent = 50
	AuditLogEventWebhookUpdate                           AuditLogEvent = 51
	AuditLogEventWebhookDelete                           AuditLogEvent = 52
	AuditLogEventEmojiCreate                             AuditLogEvent = 60
	AuditLogEventEmojiUpdate                             AuditLogEvent = 61
	AuditLogEventEmojiDelete                             AuditLogEvent = 62
	AuditLogEventMessageDelete                           AuditLogEvent = 72
	AuditLogEventMessageBulkDelete                       AuditLogEvent = 73
	AuditLogEventMessagePin                              AuditLogEvent = 74
	AuditLogEventMessageUnpin                            AuditLogEvent = 75
	AuditLogEventIntegrationCreate                       AuditLogEvent = 80
	AuditLogEventIntegrationUpdate                       AuditLogEvent = 81
	AuditLogEventIntegrationDelete                       AuditLogEvent = 82
	AuditLogEventStageInstanceCreate                     AuditLogEvent = 83
	AuditLogEventStageInstanceUpdate                     AuditLogEvent = 84
	AuditLogEventStageInstanceDelete                     AuditLogEvent = 85
	AuditLogEventStickerCreate                           AuditLogEvent = 90
	AuditLogEventStickerUpdate                           AuditLogEvent = 91
	AuditLogEventStickerDelete                           AuditLogEvent = 92
	AuditLogEventGuildScheduledEventCreate               AuditLogEvent = 100
	AuditLogEventGuildScheduledEventUpdate               AuditLogEvent = 101
	AuditLogEventGuildScheduledEventDelete               AuditLogEvent = 102
	AuditLogEventThreadCreate                            AuditLogEvent = 110
	AuditLogEventThreadUpdate                            AuditLogEvent = 111
	AuditLogEventThreadDelete                            AuditLogEvent = 112
	AuditLogEventApplicationCommandPermissionUpdate      AuditLogEvent = 121
	AuditLogEventAutoModerationRuleCreate                AuditLogEvent = 140
	AuditLogEventAutoModerationRuleUpdate                AuditLogEvent = 141
	AuditLogEventAutoModerationRuleDelete                AuditLogEvent = 142
	AuditLogEventAutoModerationBlockMessage              AuditLogEvent = 143
	AuditLogEventAutoModerationFlagToChannel             AuditLogEvent = 144
	AuditLogEventAutoModerationUserCommunicationDisabled AuditLogEvent = 145
	AuditLogEventCreatorMonetizationRequestCreated       AuditLogEvent = 150
	AuditLogEventCreatorMonetizationTermsAccepted        AuditLogEvent = 151
	AuditLogEventOnboardingPromptCreate                  AuditLogEvent = 163
	AuditLogEventOnboardingPromptUpdate                  AuditLogEvent = 164
	AuditLogEventOnboardingPromptDelete                  AuditLogEvent = 165
	AuditLogEventOnboardingCreate                        AuditLogEvent = 166
	AuditLogEventOnboardingUpdate                        AuditLogEvent = 167
	AuditLogEventHomeSettingsCreate                      AuditLogEvent = 190
	AuditLogEventHomeSettingsUpdate                      AuditLogEvent = 191
)

// OverwriteType represents the type of a permission overwrite.
type OverwriteType int

const (
	OverwriteTypeRole   OverwriteType = 0
	OverwriteTypeMember OverwriteType = 1
)

// InviteTargetType represents the target type for an invite.
type InviteTargetType int

const (
	InviteTargetTypeStream              InviteTargetType = 1
	InviteTargetTypeEmbeddedApplication InviteTargetType = 2
)

// GuildScheduledEventPrivacyLevel represents the privacy level of a scheduled event.
type GuildScheduledEventPrivacyLevel int

const (
	GuildScheduledEventPrivacyLevelGuildOnly GuildScheduledEventPrivacyLevel = 2
)

// GuildScheduledEventStatus represents the status of a scheduled event.
type GuildScheduledEventStatus int

const (
	GuildScheduledEventStatusScheduled GuildScheduledEventStatus = 1
	GuildScheduledEventStatusActive    GuildScheduledEventStatus = 2
	GuildScheduledEventStatusCompleted GuildScheduledEventStatus = 3
	GuildScheduledEventStatusCanceled  GuildScheduledEventStatus = 4
)

// GuildScheduledEventEntityType represents the entity type of a scheduled event.
type GuildScheduledEventEntityType int

const (
	GuildScheduledEventEntityTypeStageInstance GuildScheduledEventEntityType = 1
	GuildScheduledEventEntityTypeVoice         GuildScheduledEventEntityType = 2
	GuildScheduledEventEntityTypeExternal      GuildScheduledEventEntityType = 3
)

// AllowedMentionType represents types allowed in message mentions.
type AllowedMentionType string

const (
	AllowedMentionTypeRoles    AllowedMentionType = "roles"
	AllowedMentionTypeUsers    AllowedMentionType = "users"
	AllowedMentionTypeEveryone AllowedMentionType = "everyone"
)

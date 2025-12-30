package gateway

// Intent represents a Discord gateway intent bitfield.
type Intent int

const (
	IntentGuilds                      Intent = 1 << 0
	IntentGuildMembers                Intent = 1 << 1 // Privileged
	IntentGuildModeration             Intent = 1 << 2
	IntentGuildEmojisAndStickers      Intent = 1 << 3
	IntentGuildIntegrations           Intent = 1 << 4
	IntentGuildWebhooks               Intent = 1 << 5
	IntentGuildInvites                Intent = 1 << 6
	IntentGuildVoiceStates            Intent = 1 << 7
	IntentGuildPresences              Intent = 1 << 8 // Privileged
	IntentGuildMessages               Intent = 1 << 9
	IntentGuildMessageReactions       Intent = 1 << 10
	IntentGuildMessageTyping          Intent = 1 << 11
	IntentDirectMessages              Intent = 1 << 12
	IntentDirectMessageReactions      Intent = 1 << 13
	IntentDirectMessageTyping         Intent = 1 << 14
	IntentMessageContent              Intent = 1 << 15 // Privileged
	IntentGuildScheduledEvents        Intent = 1 << 16
	IntentAutoModerationConfiguration Intent = 1 << 20
	IntentAutoModerationExecution     Intent = 1 << 21
	IntentGuildMessagePolls           Intent = 1 << 24
	IntentDirectMessagePolls          Intent = 1 << 25

	IntentsGuildAll = IntentGuilds |
		IntentGuildMembers |
		IntentGuildModeration |
		IntentGuildEmojisAndStickers |
		IntentGuildIntegrations |
		IntentGuildWebhooks |
		IntentGuildInvites |
		IntentGuildVoiceStates |
		IntentGuildPresences |
		IntentGuildMessages |
		IntentGuildMessageReactions |
		IntentGuildMessageTyping |
		IntentGuildScheduledEvents |
		IntentGuildMessagePolls

	IntentsDirectAll = IntentDirectMessages |
		IntentDirectMessageReactions |
		IntentDirectMessageTyping |
		IntentDirectMessagePolls

	IntentsAll = IntentsGuildAll | IntentsDirectAll | IntentMessageContent | IntentAutoModerationConfiguration | IntentAutoModerationExecution

	IntentsNonPrivileged = IntentGuilds |
		IntentGuildModeration |
		IntentGuildEmojisAndStickers |
		IntentGuildIntegrations |
		IntentGuildWebhooks |
		IntentGuildInvites |
		IntentGuildVoiceStates |
		IntentGuildMessages |
		IntentGuildMessageReactions |
		IntentGuildMessageTyping |
		IntentDirectMessages |
		IntentDirectMessageReactions |
		IntentDirectMessageTyping |
		IntentGuildScheduledEvents |
		IntentAutoModerationConfiguration |
		IntentAutoModerationExecution |
		IntentGuildMessagePolls |
		IntentDirectMessagePolls

	IntentsPrivileged = IntentGuildMembers | IntentGuildPresences | IntentMessageContent
)

func (i Intent) Has(intent Intent) bool {
	return (i & intent) == intent
}

func (i Intent) Add(intent Intent) Intent {
	return i | intent
}

func (i Intent) Remove(intent Intent) Intent {
	return i &^ intent
}

// PrivilegedIntents returns which privileged intents are included.
func (i Intent) PrivilegedIntents() []string {
	var privileged []string
	if i.Has(IntentGuildMembers) {
		privileged = append(privileged, "GUILD_MEMBERS")
	}
	if i.Has(IntentGuildPresences) {
		privileged = append(privileged, "GUILD_PRESENCES")
	}
	if i.Has(IntentMessageContent) {
		privileged = append(privileged, "MESSAGE_CONTENT")
	}
	return privileged
}

// Warnings returns helpful warnings about the intent configuration.
func (i Intent) Warnings() []string {
	var warnings []string

	// Check for message-related intents without MESSAGE_CONTENT
	hasMessageIntent := i.Has(IntentGuildMessages) || i.Has(IntentDirectMessages)
	if hasMessageIntent && !i.Has(IntentMessageContent) {
		warnings = append(warnings,
			"MESSAGE_CONTENT intent not enabled - message content, embeds, attachments, and components will be empty. "+
				"Enable in Discord Developer Portal if needed: https://discord.com/developers/applications")
	}

	// Check for reaction intents without message intents
	if i.Has(IntentGuildMessageReactions) && !i.Has(IntentGuildMessages) {
		warnings = append(warnings,
			"GUILD_MESSAGE_REACTIONS enabled without GUILD_MESSAGES - you won't receive the original message context")
	}

	return warnings
}

// String returns a human-readable list of enabled intents.
func (i Intent) String() string {
	var intents []string

	intentMap := map[Intent]string{
		IntentGuilds:                      "GUILDS",
		IntentGuildMembers:                "GUILD_MEMBERS",
		IntentGuildModeration:             "GUILD_MODERATION",
		IntentGuildEmojisAndStickers:      "GUILD_EMOJIS_AND_STICKERS",
		IntentGuildIntegrations:           "GUILD_INTEGRATIONS",
		IntentGuildWebhooks:               "GUILD_WEBHOOKS",
		IntentGuildInvites:                "GUILD_INVITES",
		IntentGuildVoiceStates:            "GUILD_VOICE_STATES",
		IntentGuildPresences:              "GUILD_PRESENCES",
		IntentGuildMessages:               "GUILD_MESSAGES",
		IntentGuildMessageReactions:       "GUILD_MESSAGE_REACTIONS",
		IntentGuildMessageTyping:          "GUILD_MESSAGE_TYPING",
		IntentDirectMessages:              "DIRECT_MESSAGES",
		IntentDirectMessageReactions:      "DIRECT_MESSAGE_REACTIONS",
		IntentDirectMessageTyping:         "DIRECT_MESSAGE_TYPING",
		IntentMessageContent:              "MESSAGE_CONTENT",
		IntentGuildScheduledEvents:        "GUILD_SCHEDULED_EVENTS",
		IntentAutoModerationConfiguration: "AUTO_MODERATION_CONFIGURATION",
		IntentAutoModerationExecution:     "AUTO_MODERATION_EXECUTION",
		IntentGuildMessagePolls:           "GUILD_MESSAGE_POLLS",
		IntentDirectMessagePolls:          "DIRECT_MESSAGE_POLLS",
	}

	for intent, name := range intentMap {
		if i.Has(intent) {
			intents = append(intents, name)
		}
	}

	if len(intents) == 0 {
		return "none"
	}
	return join(intents, ", ")
}

func join(s []string, sep string) string {
	if len(s) == 0 {
		return ""
	}
	result := s[0]
	for _, v := range s[1:] {
		result += sep + v
	}
	return result
}

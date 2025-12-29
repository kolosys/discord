package testutil

import (
	"encoding/json"
	"time"

	"github.com/kolosys/discord/events"
	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/discord/models"
)

// UserFixture creates a test User with sensible defaults.
func UserFixture(id, username string, opts ...func(*models.User)) *models.User {
	user := &models.User{
		ID:            id,
		Username:      username,
		Discriminator: "0000",
		Bot:           false,
		Flags:         0,
		PublicFlags:   0,
	}
	for _, opt := range opts {
		opt(user)
	}
	return user
}

// BotFixture creates a test User that is a bot.
func BotFixture(id, username string, opts ...func(*models.User)) *models.User {
	user := &models.User{
		ID:            id,
		Username:      username,
		Discriminator: "0000",
		Bot:           true,
		Flags:         0,
		PublicFlags:   0,
	}
	for _, opt := range opts {
		opt(user)
	}
	return user
}

// GuildFixture creates a test Guild with sensible defaults.
func GuildFixture(id, name string, opts ...func(*models.Guild)) *models.Guild {
	guild := &models.Guild{
		ID:                          id,
		Name:                        name,
		OwnerID:                     "1",
		DefaultMessageNotifications: 0,
		ExplicitContentFilter:       0,
		VerificationLevel:           0,
		MfaLevel:                    0,
		AfkTimeout:                  300,
		MaxMembers:                  0,
		MaxVideoChannelUsers:        25,
		MaxStageVideoChannelUsers:   50,
		NsfwLevel:                   0,
		PremiumTier:                 0,
		PremiumSubscriptionCount:    0,
		PremiumProgressBarEnabled:   false,
		WidgetEnabled:               false,
		Nsfw:                        false,
		Region:                      "us-west",
		PreferredLocale:             "en-US",
		Features:                    []string{},
		Emojis:                      []any{},
		Roles:                       []any{},
		Stickers:                    []any{},
		SystemChannelFlags:          0,
	}
	for _, opt := range opts {
		opt(guild)
	}
	return guild
}

// MessageFixture creates a test Message with sensible defaults.
func MessageFixture(id, content string, opts ...func(*models.Message)) *models.Message {
	now := time.Now()
	message := &models.Message{
		ID:              id,
		ChannelID:       "1",
		Content:         content,
		Author:          UserFixture("1", "test-user"),
		Timestamp:       now,
		EditedTimestamp: nil,
		Tts:             false,
		MentionEveryone: false,
		Mentions:        []any{},
		MentionRoles:    []string{},
		MentionChannels: []any{},
		Attachments:     []any{},
		Embeds:          []any{},
		Reactions:       []any{},
		Pinned:          false,
		Type:            0,
		Flags:           0,
	}
	for _, opt := range opts {
		opt(message)
	}
	return message
}

// GuildMemberFixture creates a test GuildMember with sensible defaults.
func GuildMemberFixture(userID string, opts ...func(*models.GuildMember)) *models.GuildMember {
	member := &models.GuildMember{
		User:                       UserFixture(userID, "test-user"),
		Nick:                       nil,
		Roles:                      []string{},
		JoinedAt:                   time.Now().Add(-24 * time.Hour),
		PremiumSince:               nil,
		Deaf:                       false,
		Mute:                       false,
		Pending:                    false,
		CommunicationDisabledUntil: nil,
	}
	for _, opt := range opts {
		opt(member)
	}
	return member
}

// HelloPayload creates a Hello gateway payload.
func HelloPayload(heartbeatInterval int) *gateway.GatewayPayload {
	data := gateway.HelloData{
		HeartbeatInterval: heartbeatInterval,
	}
	raw, _ := json.Marshal(data)
	opcode := gateway.OpcodeHello
	return &gateway.GatewayPayload{
		Op: opcode,
		D:  raw,
	}
}

// ReadyPayload creates a Ready event gateway payload.
func ReadyPayload(version int, user *models.User, sessionID string, opts ...func(*events.ReadyEvent)) *gateway.GatewayPayload {
	event := &events.ReadyEvent{
		Version:          version,
		User:             user,
		Guilds:           []events.UnavailableGuild{},
		SessionID:        sessionID,
		ResumeGatewayURL: "wss://gateway.discord.gg",
		Application: &events.PartialApplication{
			ID:    "1",
			Flags: 0,
		},
	}
	for _, opt := range opts {
		opt(event)
	}

	raw, _ := json.Marshal(event)
	eventType := string(events.Ready)
	seq := 1
	return &gateway.GatewayPayload{
		Op: gateway.OpcodeDispatch,
		D:  raw,
		S:  &seq,
		T:  &eventType,
	}
}

// GuildCreatePayload creates a GUILD_CREATE event gateway payload.
func GuildCreatePayload(guild *models.Guild, opts ...func(*events.GuildCreateEvent)) *gateway.GatewayPayload {
	event := &events.GuildCreateEvent{
		Guild:       *guild,
		JoinedAt:    time.Now(),
		Large:       false,
		Unavailable: false,
		MemberCount: 1,
		Members:     []models.GuildMember{},
		Channels:    []models.GuildChannel{},
		Threads:     []any{},
		Presences:   []any{},
		VoiceStates: []any{},
	}
	for _, opt := range opts {
		opt(event)
	}

	raw, _ := json.Marshal(event)
	eventType := string(events.GuildCreate)
	seq := 2
	return &gateway.GatewayPayload{
		Op: gateway.OpcodeDispatch,
		D:  raw,
		S:  &seq,
		T:  &eventType,
	}
}

// MessageCreatePayload creates a MESSAGE_CREATE event gateway payload.
func MessageCreatePayload(message *models.Message, guildID string, userID string, opts ...func(*events.MessageCreateEvent)) *gateway.GatewayPayload {
	event := &events.MessageCreateEvent{
		Message: *message,
		GuildID: guildID,
		Member:  GuildMemberFixture(userID),
	}
	for _, opt := range opts {
		opt(event)
	}

	raw, _ := json.Marshal(event)
	eventType := string(events.MessageCreate)
	seq := 3
	return &gateway.GatewayPayload{
		Op: gateway.OpcodeDispatch,
		D:  raw,
		S:  &seq,
		T:  &eventType,
	}
}

// DispatchPayload creates a generic dispatch payload with custom event data.
func DispatchPayload(eventType string, data any, seq int) *gateway.GatewayPayload {
	raw, _ := json.Marshal(data)
	return &gateway.GatewayPayload{
		Op: gateway.OpcodeDispatch,
		D:  raw,
		S:  &seq,
		T:  &eventType,
	}
}

// WithUsername sets a custom username on a User.
func WithUsername(username string) func(*models.User) {
	return func(u *models.User) {
		u.Username = username
	}
}

// WithBot marks a User as a bot.
func WithBot(bot bool) func(*models.User) {
	return func(u *models.User) {
		u.Bot = bot
	}
}

// WithOwnerID sets a custom owner ID on a Guild.
func WithOwnerID(ownerID string) func(*models.Guild) {
	return func(g *models.Guild) {
		g.OwnerID = ownerID
	}
}

// WithMaxMembers sets a custom max members on a Guild.
func WithMaxMembers(count int) func(*models.Guild) {
	return func(g *models.Guild) {
		g.MaxMembers = int32(count)
	}
}

// WithAuthor sets a custom author on a Message.
func WithAuthor(user *models.User) func(*models.Message) {
	return func(m *models.Message) {
		m.Author = user
	}
}

// WithGuildID sets a custom guild ID on a Message.
func WithGuildID(guildID string) func(*models.Message) {
	return func(m *models.Message) {
		m.ChannelID = guildID
	}
}

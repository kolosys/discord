package events_test

import (
	"testing"

	"github.com/kolosys/discord/events"
	"github.com/kolosys/discord/models"
)

// ============================================================================
// MessageCreateEvent Tests
// ============================================================================

func TestMessageCreateEvent_AuthorUser_Direct(t *testing.T) {
	user := &models.User{
		ID:       "123",
		Username: "testuser",
		Bot:      false,
		System:   false,
	}

	event := &events.MessageCreateEvent{
		Message: models.Message{
			Author: user,
		},
	}

	got := event.AuthorUser()
	if got == nil {
		t.Fatal("AuthorUser() returned nil")
	}
	if got.ID != "123" {
		t.Errorf("ID = %q, want %q", got.ID, "123")
	}
	if got.Username != "testuser" {
		t.Errorf("Username = %q, want %q", got.Username, "testuser")
	}
}

func TestMessageCreateEvent_AuthorUser_FromMap(t *testing.T) {
	authorMap := map[string]any{
		"id":        "456",
		"username":  "mapuser",
		"bot":       true,
		"system":    false,
		"avatar":    "abcdef",
		"global_name": "Map User",
		"discriminator": "0001",
	}

	event := &events.MessageCreateEvent{
		Message: models.Message{
			Author: authorMap,
		},
	}

	got := event.AuthorUser()
	if got == nil {
		t.Fatal("AuthorUser() returned nil for map author")
	}
	if got.ID != "456" {
		t.Errorf("ID = %q, want %q", got.ID, "456")
	}
	if got.Username != "mapuser" {
		t.Errorf("Username = %q, want %q", got.Username, "mapuser")
	}
	if !got.Bot {
		t.Error("Bot should be true")
	}
	if got.GlobalName == nil || *got.GlobalName != "Map User" {
		t.Error("GlobalName not set correctly")
	}
}

func TestMessageCreateEvent_AuthorUser_Nil(t *testing.T) {
	event := &events.MessageCreateEvent{
		Message: models.Message{
			Author: nil,
		},
	}

	got := event.AuthorUser()
	if got != nil {
		t.Error("AuthorUser() should return nil when Author is nil")
	}
}

func TestMessageCreateEvent_AuthorUser_InvalidType(t *testing.T) {
	event := &events.MessageCreateEvent{
		Message: models.Message{
			Author: "invalid",
		},
	}

	got := event.AuthorUser()
	if got != nil {
		t.Error("AuthorUser() should return nil for invalid author type")
	}
}

func TestMessageCreateEvent_IsBot_True(t *testing.T) {
	user := &models.User{ID: "123", Bot: true}
	event := &events.MessageCreateEvent{
		Message: models.Message{Author: user},
	}

	if !event.IsBot() {
		t.Error("IsBot() should return true for bot message")
	}
}

func TestMessageCreateEvent_IsBot_False(t *testing.T) {
	user := &models.User{ID: "123", Bot: false}
	event := &events.MessageCreateEvent{
		Message: models.Message{Author: user},
	}

	if event.IsBot() {
		t.Error("IsBot() should return false for human message")
	}
}

func TestMessageCreateEvent_IsBot_NoAuthor(t *testing.T) {
	event := &events.MessageCreateEvent{
		Message: models.Message{Author: nil},
	}

	if event.IsBot() {
		t.Error("IsBot() should return false when author is nil")
	}
}

func TestMessageCreateEvent_IsSystem_True(t *testing.T) {
	user := &models.User{ID: "123", System: true}
	event := &events.MessageCreateEvent{
		Message: models.Message{Author: user},
	}

	if !event.IsSystem() {
		t.Error("IsSystem() should return true for system message")
	}
}

func TestMessageCreateEvent_IsSystem_False(t *testing.T) {
	user := &models.User{ID: "123", System: false}
	event := &events.MessageCreateEvent{
		Message: models.Message{Author: user},
	}

	if event.IsSystem() {
		t.Error("IsSystem() should return false for non-system message")
	}
}

func TestMessageCreateEvent_IsSystem_NoAuthor(t *testing.T) {
	event := &events.MessageCreateEvent{
		Message: models.Message{Author: nil},
	}

	if event.IsSystem() {
		t.Error("IsSystem() should return false when author is nil")
	}
}

func TestMessageCreateEvent_IsHuman_True(t *testing.T) {
	user := &models.User{
		ID:     "123",
		Bot:    false,
		System: false,
	}
	event := &events.MessageCreateEvent{
		Message: models.Message{Author: user},
	}

	if !event.IsHuman() {
		t.Error("IsHuman() should return true for human message")
	}
}

func TestMessageCreateEvent_IsHuman_False_Bot(t *testing.T) {
	user := &models.User{
		ID:     "123",
		Bot:    true,
		System: false,
	}
	event := &events.MessageCreateEvent{
		Message: models.Message{Author: user},
	}

	if event.IsHuman() {
		t.Error("IsHuman() should return false for bot message")
	}
}

func TestMessageCreateEvent_IsHuman_False_System(t *testing.T) {
	user := &models.User{
		ID:     "123",
		Bot:    false,
		System: true,
	}
	event := &events.MessageCreateEvent{
		Message: models.Message{Author: user},
	}

	if event.IsHuman() {
		t.Error("IsHuman() should return false for system message")
	}
}

func TestMessageCreateEvent_IsHuman_NoAuthor(t *testing.T) {
	event := &events.MessageCreateEvent{
		Message: models.Message{Author: nil},
	}

	if event.IsHuman() {
		t.Error("IsHuman() should return false when author is nil")
	}
}

func TestMessageCreateEvent_BaseFields(t *testing.T) {
	event := &events.MessageCreateEvent{
		GuildID: "guild_123",
		Member: &models.GuildMember{
			Roles: []string{"role1", "role2"},
		},
		Message: models.Message{
			ID:      "msg_123",
			Content: "test message",
		},
	}

	if event.GuildID != "guild_123" {
		t.Errorf("GuildID = %q, want %q", event.GuildID, "guild_123")
	}
	if event.Message.ID != "msg_123" {
		t.Errorf("Message.ID = %q, want %q", event.Message.ID, "msg_123")
	}
	if event.Message.Content != "test message" {
		t.Errorf("Message.Content = %q, want %q", event.Message.Content, "test message")
	}
}

func TestMessageCreateEvent_MapAuthor_PartialFields(t *testing.T) {
	authorMap := map[string]any{
		"id":       "789",
		"username": "partial",
		// Omitting other fields
	}

	event := &events.MessageCreateEvent{
		Message: models.Message{
			Author: authorMap,
		},
	}

	user := event.AuthorUser()
	if user == nil {
		t.Fatal("AuthorUser() returned nil")
	}
	if user.ID != "789" {
		t.Errorf("ID = %q, want %q", user.ID, "789")
	}
	if user.Username != "partial" {
		t.Errorf("Username = %q, want %q", user.Username, "partial")
	}
	if user.Bot {
		t.Error("Bot should be false (default)")
	}
}

func TestMessageCreateEvent_Combinations(t *testing.T) {
	tests := []struct {
		name     string
		user     *models.User
		wantBot  bool
		wantSys  bool
		wantHuman bool
	}{
		{
			"Regular user",
			&models.User{Bot: false, System: false},
			false, false, true,
		},
		{
			"Bot user",
			&models.User{Bot: true, System: false},
			true, false, false,
		},
		{
			"System message",
			&models.User{Bot: false, System: true},
			false, true, false,
		},
		{
			"Bot with system flag",
			&models.User{Bot: true, System: true},
			true, true, false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := &events.MessageCreateEvent{
				Message: models.Message{Author: tt.user},
			}

			if event.IsBot() != tt.wantBot {
				t.Errorf("IsBot() = %v, want %v", event.IsBot(), tt.wantBot)
			}
			if event.IsSystem() != tt.wantSys {
				t.Errorf("IsSystem() = %v, want %v", event.IsSystem(), tt.wantSys)
			}
			if event.IsHuman() != tt.wantHuman {
				t.Errorf("IsHuman() = %v, want %v", event.IsHuman(), tt.wantHuman)
			}
		})
	}
}

// ============================================================================
// Base Event Tests
// ============================================================================

func TestBase_Creation(t *testing.T) {
	base := &events.Base{}
	// ShardID is set by internal setShardID() method
	if base.ShardID != 0 {
		t.Errorf("ShardID = %d, want 0", base.ShardID)
	}
}

func TestReadyEvent_Creation(t *testing.T) {
	event := &events.ReadyEvent{
		Version:   10,
		SessionID: "session_123",
		User: &models.User{
			ID:       "user_123",
			Username: "botname",
		},
	}

	if event.Version != 10 {
		t.Errorf("Version = %d, want 10", event.Version)
	}
	if event.SessionID != "session_123" {
		t.Errorf("SessionID = %q, want %q", event.SessionID, "session_123")
	}
	if event.User.Username != "botname" {
		t.Errorf("User.Username = %q, want %q", event.User.Username, "botname")
	}
}

func TestGuildCreateEvent_Creation(t *testing.T) {
	event := &events.GuildCreateEvent{
		Large:       true,
		MemberCount: 1000,
	}

	if !event.Large {
		t.Error("Large should be true")
	}
	if event.MemberCount != 1000 {
		t.Errorf("MemberCount = %d, want 1000", event.MemberCount)
	}
}

func TestMessageDeleteEvent_Creation(t *testing.T) {
	event := &events.MessageDeleteEvent{
		ID:        "msg_123",
		ChannelID: "channel_456",
		GuildID:   "guild_789",
	}

	if event.ID != "msg_123" {
		t.Errorf("ID = %q, want %q", event.ID, "msg_123")
	}
	if event.ChannelID != "channel_456" {
		t.Errorf("ChannelID = %q, want %q", event.ChannelID, "channel_456")
	}
	if event.GuildID != "guild_789" {
		t.Errorf("GuildID = %q, want %q", event.GuildID, "guild_789")
	}
}

func TestMessageReactionAddEvent_Creation(t *testing.T) {
	event := &events.MessageReactionAddEvent{
		UserID:    "user_123",
		MessageID: "msg_456",
		Emoji: events.ReactionEmoji{
			Name: "👍",
		},
	}

	if event.UserID != "user_123" {
		t.Errorf("UserID = %q, want %q", event.UserID, "user_123")
	}
	if event.MessageID != "msg_456" {
		t.Errorf("MessageID = %q, want %q", event.MessageID, "msg_456")
	}
	if event.Emoji.Name != "👍" {
		t.Errorf("Emoji.Name = %q, want %q", event.Emoji.Name, "👍")
	}
}

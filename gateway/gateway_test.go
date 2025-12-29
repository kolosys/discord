package gateway_test

import (
	"testing"

	"github.com/kolosys/discord/events"
	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/discord/internal/testutil"
)

func TestFixtures(t *testing.T) {
	assert := testutil.New(t)

	// Test user fixture
	user := testutil.UserFixture("123", "testuser")
	assert.NotNil(user, "user fixture should be created")
	assert.Equal("123", user.ID, "user should have correct ID")
	assert.Equal("testuser", user.Username, "user should have correct username")
	assert.False(user.Bot, "user should not be a bot by default")

	// Test with options
	botUser := testutil.BotFixture("456", "mybot")
	assert.True(botUser.Bot, "bot fixture should be marked as bot")

	// Test guild fixture
	guild := testutil.GuildFixture("789", "Test Guild")
	assert.NotNil(guild, "guild fixture should be created")
	assert.Equal("789", guild.ID, "guild should have correct ID")
	assert.Equal("Test Guild", guild.Name, "guild should have correct name")

	// Test guild with options
	customGuild := testutil.GuildFixture("999", "Custom Guild", testutil.WithOwnerID("123"))
	assert.Equal("123", customGuild.OwnerID, "guild should have custom owner ID")

	// Test message fixture
	message := testutil.MessageFixture("msg1", "Hello, world!")
	assert.NotNil(message, "message fixture should be created")
	assert.Equal("msg1", message.ID, "message should have correct ID")
	assert.Equal("Hello, world!", message.Content, "message should have correct content")
	assert.NotNil(message.Author, "message should have author")

	// Test guild member fixture
	member := testutil.GuildMemberFixture("user123")
	assert.NotNil(member, "guild member fixture should be created")
	assert.NotNil(member.User, "member should have user")
}

func TestPayloadFixtures(t *testing.T) {
	assert := testutil.New(t)

	// Test hello payload
	hello := testutil.HelloPayload(41250)
	assert.GatewayPayloadOpcode(hello, gateway.OpcodeHello, "hello should have correct opcode")
	assert.Nil(hello.T, "hello should not have event type")

	// Test ready payload
	user := testutil.UserFixture("bot123", "MyBot")
	ready := testutil.ReadyPayload(10, user, "session-id-123")
	assert.GatewayPayloadOpcode(ready, gateway.OpcodeDispatch, "ready should be dispatch opcode")
	assert.GatewayPayloadEventType(ready, string(events.Ready), "ready should have READY event type")
	assert.GatewayPayloadHasSequence(ready, "ready should have sequence")
	assert.GatewayPayloadSequence(ready, 1, "ready should have sequence 1")

	// Test guild create payload
	guild := testutil.GuildFixture("guild123", "My Guild")
	guildCreate := testutil.GuildCreatePayload(guild)
	assert.GatewayPayloadOpcode(guildCreate, gateway.OpcodeDispatch, "guild_create should be dispatch opcode")
	assert.GatewayPayloadEventType(guildCreate, string(events.GuildCreate), "guild_create should have GUILD_CREATE event type")
	assert.GatewayPayloadHasSequence(guildCreate, "guild_create should have sequence")

	// Test message create payload
	message := testutil.MessageFixture("msg123", "Test message")
	msgCreate := testutil.MessageCreatePayload(message, "guild123", "user123")
	assert.GatewayPayloadOpcode(msgCreate, gateway.OpcodeDispatch, "message_create should be dispatch opcode")
	assert.GatewayPayloadEventType(msgCreate, string(events.MessageCreate), "message_create should have MESSAGE_CREATE event type")

	// Test generic dispatch payload
	generic := testutil.DispatchPayload("CUSTOM_EVENT", map[string]string{"key": "value"}, 5)
	assert.GatewayPayloadOpcode(generic, gateway.OpcodeDispatch, "generic should be dispatch opcode")
	assert.GatewayPayloadEventType(generic, "CUSTOM_EVENT", "generic should have custom event type")
	assert.GatewayPayloadSequence(generic, 5, "generic should have correct sequence")
}

func TestMockDiscord(t *testing.T) {
	assert := testutil.New(t)

	// Set up mock Discord API
	mockAPI, err := testutil.NewMockDiscord()
	if err != nil {
		t.Fatalf("failed to create mock discord: %v", err)
	}
	defer mockAPI.Close()

	// Register some endpoints
	mockAPI.HandleJSON("GET", "/api/v10/users/@me", 200, map[string]string{
		"id":       "bot123",
		"username": "MyBot",
	})

	mockAPI.HandleError("GET", "/api/v10/guilds/invalid", 404, "Guild not found")

	assert.Equal(0, mockAPI.RequestCount(), "should start with no requests")
	assert.Nil(mockAPI.LastRequest(), "should have no last request")

	// Verify endpoints are available (in real test, would use http client)
	assert.NotNil(mockAPI, "mock discord should be created")
	assert.Contains(mockAPI.URL(), "http://", "mock discord should have valid URL")

	// Clear requests
	mockAPI.ClearRequests()
	assert.Equal(0, mockAPI.RequestCount(), "requests should be cleared")
}

func TestAssertions(t *testing.T) {
	assert := testutil.New(t)

	// Test basic assertions
	assert.True(true, "true should pass")
	assert.False(false, "false should pass")
	assert.Equal(42, 42, "equal values should pass")
	assert.NotEqual(1, 2, "unequal values should pass")

	// Test nil assertions
	var nilVal *int
	assert.Nil(nilVal, "nil should pass")
	assert.NotNil(&nilVal, "not nil should pass")

	// Test string assertions
	assert.Contains("hello world", "world", "contains should find substring")
	assert.NotContains("hello world", "goodbye", "not contains should work")

	// Test slice assertions
	slice := []string{"a", "b", "c"}
	assert.SliceLen(slice, 3, "slice should have length 3")
	assert.SliceContains(slice, "b", "slice should contain element")

	// Test error assertions
	var err error
	assert.NoError(err, "nil error should pass")

	mockGW, errCreate := testutil.NewMockGateway()
	assert.NoError(errCreate, "creating mock gateway should not error")
	if mockGW != nil {
		mockGW.Close()
	}
}

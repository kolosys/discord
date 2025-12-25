package discord

import (
	"context"
	"encoding/json"

	"github.com/kolosys/discord/events"
)

// On registers a raw event handler for a specific event type.
func (c *Client) On(eventType events.Type, handler events.RawHandler) error {
	return c.dispatcher.OnRaw(eventType, handler)
}

// Listen is a convenience helper that registers a typed handler for an event.
// It returns a function that can be called to unsubscribe.
//
// Example:
//
//	discord.Listen(client, events.MessageCreate, func(ctx context.Context, e *events.MessageCreateEvent) {
//	    fmt.Printf("Message: %s\n", e.Content)
//	})
func Listen[T any](c *Client, eventType events.Type, handler events.Handler[T]) error {
	return events.On(c.dispatcher, eventType, handler)
}

// ListenRaw registers a handler that receives raw JSON data for an event.
func ListenRaw(c *Client, eventType events.Type, handler func(ctx context.Context, data json.RawMessage)) error {
	return c.dispatcher.OnRaw(eventType, func(ctx context.Context, _ events.Type, data json.RawMessage) {
		handler(ctx, data)
	})
}

// OnReady registers a handler for the Ready event.
func (c *Client) OnReady(handler events.Handler[events.ReadyEvent]) error {
	return events.On(c.dispatcher, events.Ready, handler)
}

// OnGuildCreate registers a handler for the GuildCreate event.
func (c *Client) OnGuildCreate(handler events.Handler[events.GuildCreateEvent]) error {
	return events.On(c.dispatcher, events.GuildCreate, handler)
}

// OnGuildUpdate registers a handler for the GuildUpdate event.
func (c *Client) OnGuildUpdate(handler events.Handler[events.GuildUpdateEvent]) error {
	return events.On(c.dispatcher, events.GuildUpdate, handler)
}

// OnGuildDelete registers a handler for the GuildDelete event.
func (c *Client) OnGuildDelete(handler events.Handler[events.GuildDeleteEvent]) error {
	return events.On(c.dispatcher, events.GuildDelete, handler)
}

// OnMessageCreate registers a handler for the MessageCreate event.
func (c *Client) OnMessageCreate(handler events.Handler[events.MessageCreateEvent]) error {
	return events.On(c.dispatcher, events.MessageCreate, handler)
}

// OnMessageUpdate registers a handler for the MessageUpdate event.
func (c *Client) OnMessageUpdate(handler events.Handler[events.MessageUpdateEvent]) error {
	return events.On(c.dispatcher, events.MessageUpdate, handler)
}

// OnMessageDelete registers a handler for the MessageDelete event.
func (c *Client) OnMessageDelete(handler events.Handler[events.MessageDeleteEvent]) error {
	return events.On(c.dispatcher, events.MessageDelete, handler)
}

// OnMessageDeleteBulk registers a handler for the MessageDeleteBulk event.
func (c *Client) OnMessageDeleteBulk(handler events.Handler[events.MessageDeleteBulkEvent]) error {
	return events.On(c.dispatcher, events.MessageDeleteBulk, handler)
}

// OnMessageReactionAdd registers a handler for the MessageReactionAdd event.
func (c *Client) OnMessageReactionAdd(handler events.Handler[events.MessageReactionAddEvent]) error {
	return events.On(c.dispatcher, events.MessageReactionAdd, handler)
}

// OnMessageReactionRemove registers a handler for the MessageReactionRemove event.
func (c *Client) OnMessageReactionRemove(handler events.Handler[events.MessageReactionRemoveEvent]) error {
	return events.On(c.dispatcher, events.MessageReactionRemove, handler)
}

// OnGuildMemberAdd registers a handler for the GuildMemberAdd event.
func (c *Client) OnGuildMemberAdd(handler events.Handler[events.GuildMemberAddEvent]) error {
	return events.On(c.dispatcher, events.GuildMemberAdd, handler)
}

// OnGuildMemberUpdate registers a handler for the GuildMemberUpdate event.
func (c *Client) OnGuildMemberUpdate(handler events.Handler[events.GuildMemberUpdateEvent]) error {
	return events.On(c.dispatcher, events.GuildMemberUpdate, handler)
}

// OnGuildMemberRemove registers a handler for the GuildMemberRemove event.
func (c *Client) OnGuildMemberRemove(handler events.Handler[events.GuildMemberRemoveEvent]) error {
	return events.On(c.dispatcher, events.GuildMemberRemove, handler)
}

// OnGuildMembersChunk registers a handler for the GuildMembersChunk event.
// This event is received in response to RequestGuildMembers.
func (c *Client) OnGuildMembersChunk(handler events.Handler[events.GuildMembersChunkEvent]) error {
	return events.On(c.dispatcher, events.GuildMembersChunk, handler)
}

// OnGuildBanAdd registers a handler for the GuildBanAdd event.
func (c *Client) OnGuildBanAdd(handler events.Handler[events.GuildBanAddEvent]) error {
	return events.On(c.dispatcher, events.GuildBanAdd, handler)
}

// OnGuildBanRemove registers a handler for the GuildBanRemove event.
func (c *Client) OnGuildBanRemove(handler events.Handler[events.GuildBanRemoveEvent]) error {
	return events.On(c.dispatcher, events.GuildBanRemove, handler)
}

// OnGuildRoleCreate registers a handler for the GuildRoleCreate event.
func (c *Client) OnGuildRoleCreate(handler events.Handler[events.GuildRoleCreateEvent]) error {
	return events.On(c.dispatcher, events.GuildRoleCreate, handler)
}

// OnGuildRoleUpdate registers a handler for the GuildRoleUpdate event.
func (c *Client) OnGuildRoleUpdate(handler events.Handler[events.GuildRoleUpdateEvent]) error {
	return events.On(c.dispatcher, events.GuildRoleUpdate, handler)
}

// OnGuildRoleDelete registers a handler for the GuildRoleDelete event.
func (c *Client) OnGuildRoleDelete(handler events.Handler[events.GuildRoleDeleteEvent]) error {
	return events.On(c.dispatcher, events.GuildRoleDelete, handler)
}

// OnChannelCreate registers a handler for the ChannelCreate event.
func (c *Client) OnChannelCreate(handler events.Handler[events.ChannelCreateEvent]) error {
	return events.On(c.dispatcher, events.ChannelCreate, handler)
}

// OnChannelUpdate registers a handler for the ChannelUpdate event.
func (c *Client) OnChannelUpdate(handler events.Handler[events.ChannelUpdateEvent]) error {
	return events.On(c.dispatcher, events.ChannelUpdate, handler)
}

// OnChannelDelete registers a handler for the ChannelDelete event.
func (c *Client) OnChannelDelete(handler events.Handler[events.ChannelDeleteEvent]) error {
	return events.On(c.dispatcher, events.ChannelDelete, handler)
}

// OnInteractionCreate registers a handler for the InteractionCreate event.
func (c *Client) OnInteractionCreate(handler events.Handler[events.InteractionCreateEvent]) error {
	return events.On(c.dispatcher, events.InteractionCreate, handler)
}

// OnTypingStart registers a handler for the TypingStart event.
func (c *Client) OnTypingStart(handler events.Handler[events.TypingStartEvent]) error {
	return events.On(c.dispatcher, events.TypingStart, handler)
}

// OnVoiceStateUpdate registers a handler for the VoiceStateUpdate event.
func (c *Client) OnVoiceStateUpdate(handler events.Handler[events.VoiceStateUpdateEvent]) error {
	return events.On(c.dispatcher, events.VoiceStateUpdate, handler)
}

// OnVoiceServerUpdate registers a handler for the VoiceServerUpdate event.
func (c *Client) OnVoiceServerUpdate(handler events.Handler[events.VoiceServerUpdateEvent]) error {
	return events.On(c.dispatcher, events.VoiceServerUpdate, handler)
}

// OnPresenceUpdate registers a handler for the PresenceUpdate event.
func (c *Client) OnPresenceUpdate(handler events.Handler[events.PresenceUpdateEvent]) error {
	return events.On(c.dispatcher, events.PresenceUpdate, handler)
}

// OnInviteCreate registers a handler for the InviteCreate event.
func (c *Client) OnInviteCreate(handler events.Handler[events.InviteCreateEvent]) error {
	return events.On(c.dispatcher, events.InviteCreate, handler)
}

// OnInviteDelete registers a handler for the InviteDelete event.
func (c *Client) OnInviteDelete(handler events.Handler[events.InviteDeleteEvent]) error {
	return events.On(c.dispatcher, events.InviteDelete, handler)
}

// OnEvent registers a typed event handler using generics.
// Example: discord.OnEvent[events.MessageCreateEvent](client, events.MessageCreate, handler)
func OnEvent[T any](c *Client, eventType events.Type, handler events.Handler[T]) error {
	return events.On(c.dispatcher, eventType, handler)
}

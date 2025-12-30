package gateway

import (
	"context"
	"encoding/json"

	"github.com/kolosys/discord/events"
)

// On registers a raw event handler for a specific event type.
func (g *Gateway) On(eventType events.Type, handler events.RawHandler) error {
	return g.dispatcher.OnRaw(eventType, handler)
}

// Listen is a convenience helper that registers a typed handler for an event.
//
// Example:
//
//	gateway.Listen(gw, events.MessageCreate, func(ctx context.Context, e *events.MessageCreateEvent) {
//	    fmt.Printf("Message: %s\n", e.Content)
//	})
func Listen[T any](g *Gateway, eventType events.Type, handler events.Handler[T]) error {
	return events.On(g.dispatcher, eventType, handler)
}

// ListenRaw registers a handler that receives raw JSON data for an event.
func ListenRaw(g *Gateway, eventType events.Type, handler func(ctx context.Context, data json.RawMessage)) error {
	return g.dispatcher.OnRaw(eventType, func(ctx context.Context, _ events.Type, data json.RawMessage) {
		handler(ctx, data)
	})
}

// OnReady registers a handler for the Ready event.
func (g *Gateway) OnReady(handler events.Handler[events.ReadyEvent]) error {
	return events.On(g.dispatcher, events.Ready, handler)
}

// OnGuildCreate registers a handler for the GuildCreate event.
func (g *Gateway) OnGuildCreate(handler events.Handler[events.GuildCreateEvent]) error {
	return events.On(g.dispatcher, events.GuildCreate, handler)
}

// OnGuildUpdate registers a handler for the GuildUpdate event.
func (g *Gateway) OnGuildUpdate(handler events.Handler[events.GuildUpdateEvent]) error {
	return events.On(g.dispatcher, events.GuildUpdate, handler)
}

// OnGuildDelete registers a handler for the GuildDelete event.
func (g *Gateway) OnGuildDelete(handler events.Handler[events.GuildDeleteEvent]) error {
	return events.On(g.dispatcher, events.GuildDelete, handler)
}

// OnMessageCreate registers a handler for the MessageCreate event.
func (g *Gateway) OnMessageCreate(handler events.Handler[events.MessageCreateEvent]) error {
	return events.On(g.dispatcher, events.MessageCreate, handler)
}

// OnMessageUpdate registers a handler for the MessageUpdate event.
func (g *Gateway) OnMessageUpdate(handler events.Handler[events.MessageUpdateEvent]) error {
	return events.On(g.dispatcher, events.MessageUpdate, handler)
}

// OnMessageDelete registers a handler for the MessageDelete event.
func (g *Gateway) OnMessageDelete(handler events.Handler[events.MessageDeleteEvent]) error {
	return events.On(g.dispatcher, events.MessageDelete, handler)
}

// OnMessageDeleteBulk registers a handler for the MessageDeleteBulk event.
func (g *Gateway) OnMessageDeleteBulk(handler events.Handler[events.MessageDeleteBulkEvent]) error {
	return events.On(g.dispatcher, events.MessageDeleteBulk, handler)
}

// OnMessageReactionAdd registers a handler for the MessageReactionAdd event.
func (g *Gateway) OnMessageReactionAdd(handler events.Handler[events.MessageReactionAddEvent]) error {
	return events.On(g.dispatcher, events.MessageReactionAdd, handler)
}

// OnMessageReactionRemove registers a handler for the MessageReactionRemove event.
func (g *Gateway) OnMessageReactionRemove(handler events.Handler[events.MessageReactionRemoveEvent]) error {
	return events.On(g.dispatcher, events.MessageReactionRemove, handler)
}

// OnGuildMemberAdd registers a handler for the GuildMemberAdd event.
func (g *Gateway) OnGuildMemberAdd(handler events.Handler[events.GuildMemberAddEvent]) error {
	return events.On(g.dispatcher, events.GuildMemberAdd, handler)
}

// OnGuildMemberUpdate registers a handler for the GuildMemberUpdate event.
func (g *Gateway) OnGuildMemberUpdate(handler events.Handler[events.GuildMemberUpdateEvent]) error {
	return events.On(g.dispatcher, events.GuildMemberUpdate, handler)
}

// OnGuildMemberRemove registers a handler for the GuildMemberRemove event.
func (g *Gateway) OnGuildMemberRemove(handler events.Handler[events.GuildMemberRemoveEvent]) error {
	return events.On(g.dispatcher, events.GuildMemberRemove, handler)
}

// OnGuildMembersChunk registers a handler for the GuildMembersChunk event.
func (g *Gateway) OnGuildMembersChunk(handler events.Handler[events.GuildMembersChunkEvent]) error {
	return events.On(g.dispatcher, events.GuildMembersChunk, handler)
}

// OnGuildBanAdd registers a handler for the GuildBanAdd event.
func (g *Gateway) OnGuildBanAdd(handler events.Handler[events.GuildBanAddEvent]) error {
	return events.On(g.dispatcher, events.GuildBanAdd, handler)
}

// OnGuildBanRemove registers a handler for the GuildBanRemove event.
func (g *Gateway) OnGuildBanRemove(handler events.Handler[events.GuildBanRemoveEvent]) error {
	return events.On(g.dispatcher, events.GuildBanRemove, handler)
}

// OnGuildRoleCreate registers a handler for the GuildRoleCreate event.
func (g *Gateway) OnGuildRoleCreate(handler events.Handler[events.GuildRoleCreateEvent]) error {
	return events.On(g.dispatcher, events.GuildRoleCreate, handler)
}

// OnGuildRoleUpdate registers a handler for the GuildRoleUpdate event.
func (g *Gateway) OnGuildRoleUpdate(handler events.Handler[events.GuildRoleUpdateEvent]) error {
	return events.On(g.dispatcher, events.GuildRoleUpdate, handler)
}

// OnGuildRoleDelete registers a handler for the GuildRoleDelete event.
func (g *Gateway) OnGuildRoleDelete(handler events.Handler[events.GuildRoleDeleteEvent]) error {
	return events.On(g.dispatcher, events.GuildRoleDelete, handler)
}

// OnChannelCreate registers a handler for the ChannelCreate event.
func (g *Gateway) OnChannelCreate(handler events.Handler[events.ChannelCreateEvent]) error {
	return events.On(g.dispatcher, events.ChannelCreate, handler)
}

// OnChannelUpdate registers a handler for the ChannelUpdate event.
func (g *Gateway) OnChannelUpdate(handler events.Handler[events.ChannelUpdateEvent]) error {
	return events.On(g.dispatcher, events.ChannelUpdate, handler)
}

// OnChannelDelete registers a handler for the ChannelDelete event.
func (g *Gateway) OnChannelDelete(handler events.Handler[events.ChannelDeleteEvent]) error {
	return events.On(g.dispatcher, events.ChannelDelete, handler)
}

// OnInteractionCreate registers a handler for the InteractionCreate event.
func (g *Gateway) OnInteractionCreate(handler events.Handler[events.InteractionCreateEvent]) error {
	return events.On(g.dispatcher, events.InteractionCreate, handler)
}

// OnTypingStart registers a handler for the TypingStart event.
func (g *Gateway) OnTypingStart(handler events.Handler[events.TypingStartEvent]) error {
	return events.On(g.dispatcher, events.TypingStart, handler)
}

// OnVoiceStateUpdate registers a handler for the VoiceStateUpdate event.
func (g *Gateway) OnVoiceStateUpdate(handler events.Handler[events.VoiceStateUpdateEvent]) error {
	return events.On(g.dispatcher, events.VoiceStateUpdate, handler)
}

// OnVoiceServerUpdate registers a handler for the VoiceServerUpdate event.
func (g *Gateway) OnVoiceServerUpdate(handler events.Handler[events.VoiceServerUpdateEvent]) error {
	return events.On(g.dispatcher, events.VoiceServerUpdate, handler)
}

// OnPresenceUpdate registers a handler for the PresenceUpdate event.
func (g *Gateway) OnPresenceUpdate(handler events.Handler[events.PresenceUpdateEvent]) error {
	return events.On(g.dispatcher, events.PresenceUpdate, handler)
}

// OnInviteCreate registers a handler for the InviteCreate event.
func (g *Gateway) OnInviteCreate(handler events.Handler[events.InviteCreateEvent]) error {
	return events.On(g.dispatcher, events.InviteCreate, handler)
}

// OnInviteDelete registers a handler for the InviteDelete event.
func (g *Gateway) OnInviteDelete(handler events.Handler[events.InviteDeleteEvent]) error {
	return events.On(g.dispatcher, events.InviteDelete, handler)
}

// OnEvent registers a typed event handler using generics.
func OnEvent[T any](g *Gateway, eventType events.Type, handler events.Handler[T]) error {
	return events.On(g.dispatcher, eventType, handler)
}

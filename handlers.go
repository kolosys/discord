package discord

import (
	"context"
	"encoding/json"

	"github.com/kolosys/discord/events"
)

// On registers a raw event handler for a specific event type.
func (b *Bot) On(eventType events.Type, handler events.RawHandler) error {
	return b.dispatcher.OnRaw(eventType, handler)
}

// Listen is a convenience helper that registers a typed handler for an event.
//
// Example:
//
//	discord.Listen(bot, events.MessageCreate, func(ctx context.Context, e *events.MessageCreateEvent) {
//	    fmt.Printf("Message: %s\n", e.Content)
//	})
func Listen[T any](b *Bot, eventType events.Type, handler events.Handler[T]) error {
	return events.On(b.dispatcher, eventType, handler)
}

// ListenRaw registers a handler that receives raw JSON data for an event.
func ListenRaw(b *Bot, eventType events.Type, handler func(ctx context.Context, data json.RawMessage)) error {
	return b.dispatcher.OnRaw(eventType, func(ctx context.Context, _ events.Type, data json.RawMessage) {
		handler(ctx, data)
	})
}

// OnReady registers a handler for the Ready event.
func (b *Bot) OnReady(handler events.Handler[events.ReadyEvent]) error {
	return events.On(b.dispatcher, events.Ready, handler)
}

// OnGuildCreate registers a handler for the GuildCreate event.
func (b *Bot) OnGuildCreate(handler events.Handler[events.GuildCreateEvent]) error {
	return events.On(b.dispatcher, events.GuildCreate, handler)
}

// OnGuildUpdate registers a handler for the GuildUpdate event.
func (b *Bot) OnGuildUpdate(handler events.Handler[events.GuildUpdateEvent]) error {
	return events.On(b.dispatcher, events.GuildUpdate, handler)
}

// OnGuildDelete registers a handler for the GuildDelete event.
func (b *Bot) OnGuildDelete(handler events.Handler[events.GuildDeleteEvent]) error {
	return events.On(b.dispatcher, events.GuildDelete, handler)
}

// OnMessageCreate registers a handler for the MessageCreate event.
func (b *Bot) OnMessageCreate(handler events.Handler[events.MessageCreateEvent]) error {
	return events.On(b.dispatcher, events.MessageCreate, handler)
}

// OnMessageUpdate registers a handler for the MessageUpdate event.
func (b *Bot) OnMessageUpdate(handler events.Handler[events.MessageUpdateEvent]) error {
	return events.On(b.dispatcher, events.MessageUpdate, handler)
}

// OnMessageDelete registers a handler for the MessageDelete event.
func (b *Bot) OnMessageDelete(handler events.Handler[events.MessageDeleteEvent]) error {
	return events.On(b.dispatcher, events.MessageDelete, handler)
}

// OnMessageDeleteBulk registers a handler for the MessageDeleteBulk event.
func (b *Bot) OnMessageDeleteBulk(handler events.Handler[events.MessageDeleteBulkEvent]) error {
	return events.On(b.dispatcher, events.MessageDeleteBulk, handler)
}

// OnMessageReactionAdd registers a handler for the MessageReactionAdd event.
func (b *Bot) OnMessageReactionAdd(handler events.Handler[events.MessageReactionAddEvent]) error {
	return events.On(b.dispatcher, events.MessageReactionAdd, handler)
}

// OnMessageReactionRemove registers a handler for the MessageReactionRemove event.
func (b *Bot) OnMessageReactionRemove(handler events.Handler[events.MessageReactionRemoveEvent]) error {
	return events.On(b.dispatcher, events.MessageReactionRemove, handler)
}

// OnGuildMemberAdd registers a handler for the GuildMemberAdd event.
func (b *Bot) OnGuildMemberAdd(handler events.Handler[events.GuildMemberAddEvent]) error {
	return events.On(b.dispatcher, events.GuildMemberAdd, handler)
}

// OnGuildMemberUpdate registers a handler for the GuildMemberUpdate event.
func (b *Bot) OnGuildMemberUpdate(handler events.Handler[events.GuildMemberUpdateEvent]) error {
	return events.On(b.dispatcher, events.GuildMemberUpdate, handler)
}

// OnGuildMemberRemove registers a handler for the GuildMemberRemove event.
func (b *Bot) OnGuildMemberRemove(handler events.Handler[events.GuildMemberRemoveEvent]) error {
	return events.On(b.dispatcher, events.GuildMemberRemove, handler)
}

// OnGuildMembersChunk registers a handler for the GuildMembersChunk event.
func (b *Bot) OnGuildMembersChunk(handler events.Handler[events.GuildMembersChunkEvent]) error {
	return events.On(b.dispatcher, events.GuildMembersChunk, handler)
}

// OnGuildBanAdd registers a handler for the GuildBanAdd event.
func (b *Bot) OnGuildBanAdd(handler events.Handler[events.GuildBanAddEvent]) error {
	return events.On(b.dispatcher, events.GuildBanAdd, handler)
}

// OnGuildBanRemove registers a handler for the GuildBanRemove event.
func (b *Bot) OnGuildBanRemove(handler events.Handler[events.GuildBanRemoveEvent]) error {
	return events.On(b.dispatcher, events.GuildBanRemove, handler)
}

// OnGuildRoleCreate registers a handler for the GuildRoleCreate event.
func (b *Bot) OnGuildRoleCreate(handler events.Handler[events.GuildRoleCreateEvent]) error {
	return events.On(b.dispatcher, events.GuildRoleCreate, handler)
}

// OnGuildRoleUpdate registers a handler for the GuildRoleUpdate event.
func (b *Bot) OnGuildRoleUpdate(handler events.Handler[events.GuildRoleUpdateEvent]) error {
	return events.On(b.dispatcher, events.GuildRoleUpdate, handler)
}

// OnGuildRoleDelete registers a handler for the GuildRoleDelete event.
func (b *Bot) OnGuildRoleDelete(handler events.Handler[events.GuildRoleDeleteEvent]) error {
	return events.On(b.dispatcher, events.GuildRoleDelete, handler)
}

// OnChannelCreate registers a handler for the ChannelCreate event.
func (b *Bot) OnChannelCreate(handler events.Handler[events.ChannelCreateEvent]) error {
	return events.On(b.dispatcher, events.ChannelCreate, handler)
}

// OnChannelUpdate registers a handler for the ChannelUpdate event.
func (b *Bot) OnChannelUpdate(handler events.Handler[events.ChannelUpdateEvent]) error {
	return events.On(b.dispatcher, events.ChannelUpdate, handler)
}

// OnChannelDelete registers a handler for the ChannelDelete event.
func (b *Bot) OnChannelDelete(handler events.Handler[events.ChannelDeleteEvent]) error {
	return events.On(b.dispatcher, events.ChannelDelete, handler)
}

// OnInteractionCreate registers a handler for the InteractionCreate event.
func (b *Bot) OnInteractionCreate(handler events.Handler[events.InteractionCreateEvent]) error {
	return events.On(b.dispatcher, events.InteractionCreate, handler)
}

// OnTypingStart registers a handler for the TypingStart event.
func (b *Bot) OnTypingStart(handler events.Handler[events.TypingStartEvent]) error {
	return events.On(b.dispatcher, events.TypingStart, handler)
}

// OnVoiceStateUpdate registers a handler for the VoiceStateUpdate event.
func (b *Bot) OnVoiceStateUpdate(handler events.Handler[events.VoiceStateUpdateEvent]) error {
	return events.On(b.dispatcher, events.VoiceStateUpdate, handler)
}

// OnVoiceServerUpdate registers a handler for the VoiceServerUpdate event.
func (b *Bot) OnVoiceServerUpdate(handler events.Handler[events.VoiceServerUpdateEvent]) error {
	return events.On(b.dispatcher, events.VoiceServerUpdate, handler)
}

// OnPresenceUpdate registers a handler for the PresenceUpdate event.
func (b *Bot) OnPresenceUpdate(handler events.Handler[events.PresenceUpdateEvent]) error {
	return events.On(b.dispatcher, events.PresenceUpdate, handler)
}

// OnInviteCreate registers a handler for the InviteCreate event.
func (b *Bot) OnInviteCreate(handler events.Handler[events.InviteCreateEvent]) error {
	return events.On(b.dispatcher, events.InviteCreate, handler)
}

// OnInviteDelete registers a handler for the InviteDelete event.
func (b *Bot) OnInviteDelete(handler events.Handler[events.InviteDeleteEvent]) error {
	return events.On(b.dispatcher, events.InviteDelete, handler)
}

// OnEvent registers a typed event handler using generics.
func OnEvent[T any](b *Bot, eventType events.Type, handler events.Handler[T]) error {
	return events.On(b.dispatcher, eventType, handler)
}

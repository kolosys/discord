package discord

import (
	"context"

	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/discord/models"
)

// SendMessage sends a text message to a channel.
func (b *Bot) SendMessage(ctx context.Context, channelID, content string) (*models.Message, error) {
	return b.REST.CreateMessage(ctx, channelID, models.MessageCreateOptions{
		Content: &content,
	})
}

// SendReply sends a text message as a reply to another message.
func (b *Bot) SendReply(ctx context.Context, channelID, messageID, content string) (*models.Message, error) {
	return b.REST.CreateMessage(ctx, channelID, models.MessageCreateOptions{
		Content: &content,
		MessageReference: map[string]any{
			"message_id": messageID,
		},
	})
}

// SendEmbed sends an embed message to a channel.
func (b *Bot) SendEmbed(ctx context.Context, channelID string, embed models.RichEmbed) (*models.Message, error) {
	return b.REST.CreateMessage(ctx, channelID, models.MessageCreateOptions{
		Embeds: []any{embed},
	})
}

// SendMessageComplex sends a message with full options control.
func (b *Bot) SendMessageComplex(ctx context.Context, channelID string, opts models.MessageCreateOptions) (*models.Message, error) {
	return b.REST.CreateMessage(ctx, channelID, opts)
}

// SetStatus updates the bot's status on all shards.
func (b *Bot) SetStatus(ctx context.Context, status string) error {
	return b.Gateway.UpdatePresence(ctx, &gateway.PresenceUpdate{
		Status:     status,
		Activities: nil,
		AFK:        false,
	})
}

// SetActivity updates the bot's activity (e.g., "Playing a game").
func (b *Bot) SetActivity(ctx context.Context, activityType gateway.ActivityType, name string) error {
	return b.Gateway.UpdatePresence(ctx, &gateway.PresenceUpdate{
		Status: "online",
		Activities: []*gateway.Activity{
			{Name: name, Type: activityType},
		},
		AFK: false,
	})
}

// SetPresence updates the bot's full presence (status + activity).
func (b *Bot) SetPresence(ctx context.Context, presence *gateway.PresenceUpdate) error {
	return b.Gateway.UpdatePresence(ctx, presence)
}

// JoinVoiceChannel joins a voice channel.
func (b *Bot) JoinVoiceChannel(ctx context.Context, guildID, channelID string) error {
	return b.Gateway.UpdateVoiceState(ctx, &gateway.VoiceStateUpdateData{
		GuildID:   guildID,
		ChannelID: &channelID,
		SelfMute:  false,
		SelfDeaf:  false,
	})
}

// LeaveVoiceChannel leaves the voice channel in a guild.
func (b *Bot) LeaveVoiceChannel(ctx context.Context, guildID string) error {
	return b.Gateway.UpdateVoiceState(ctx, &gateway.VoiceStateUpdateData{
		GuildID:   guildID,
		ChannelID: nil,
		SelfMute:  false,
		SelfDeaf:  false,
	})
}

// UpdateVoiceState updates the bot's voice state with full control.
func (b *Bot) UpdateVoiceState(ctx context.Context, opts *gateway.VoiceStateUpdateData) error {
	return b.Gateway.UpdateVoiceState(ctx, opts)
}

// RequestGuildMembers requests guild members from Discord.
// Results are received via GUILD_MEMBERS_CHUNK events.
func (b *Bot) RequestGuildMembers(ctx context.Context, guildID string, query string, limit int) error {
	return b.Gateway.RequestGuildMembers(ctx, &gateway.RequestGuildMembersData{
		GuildID: guildID,
		Query:   &query,
		Limit:   limit,
	})
}

// RequestGuildMembersByID requests specific guild members by user ID.
// Results are received via GUILD_MEMBERS_CHUNK events.
func (b *Bot) RequestGuildMembersByID(ctx context.Context, guildID string, userIDs []string) error {
	return b.Gateway.RequestGuildMembers(ctx, &gateway.RequestGuildMembersData{
		GuildID: guildID,
		UserIDs: userIDs,
		Limit:   0,
	})
}

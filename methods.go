package discord

import (
	"context"

	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/discord/models"
)

// SendMessage sends a text message to a channel.
func (c *Client) SendMessage(ctx context.Context, channelID, content string) (*models.Message, error) {
	return c.rest.CreateMessage(ctx, channelID, models.MessageCreateOptions{
		Content: &content,
	})
}

// SendReply sends a text message as a reply to another message.
func (c *Client) SendReply(ctx context.Context, channelID, messageID, content string) (*models.Message, error) {
	return c.rest.CreateMessage(ctx, channelID, models.MessageCreateOptions{
		Content: &content,
		MessageReference: map[string]any{
			"message_id": messageID,
		},
	})
}

// SendEmbed sends an embed message to a channel.
func (c *Client) SendEmbed(ctx context.Context, channelID string, embed models.RichEmbed) (*models.Message, error) {
	return c.rest.CreateMessage(ctx, channelID, models.MessageCreateOptions{
		Embeds: []any{embed},
	})
}

// SendMessageComplex sends a message with full options control.
func (c *Client) SendMessageComplex(ctx context.Context, channelID string, opts models.MessageCreateOptions) (*models.Message, error) {
	return c.rest.CreateMessage(ctx, channelID, opts)
}

// SetStatus updates the bot's status on all shards.
func (c *Client) SetStatus(ctx context.Context, status string) error {
	return c.gateway.UpdatePresence(ctx, &gateway.PresenceUpdate{
		Status:     status,
		Activities: nil,
		AFK:        false,
	})
}

// SetActivity updates the bot's activity (e.g., "Playing a game").
func (c *Client) SetActivity(ctx context.Context, activityType gateway.ActivityType, name string) error {
	return c.gateway.UpdatePresence(ctx, &gateway.PresenceUpdate{
		Status: "online",
		Activities: []*gateway.Activity{
			{
				Name: name,
				Type: activityType,
			},
		},
		AFK: false,
	})
}

// SetPresence updates the bot's full presence (status + activity).
func (c *Client) SetPresence(ctx context.Context, presence *gateway.PresenceUpdate) error {
	return c.gateway.UpdatePresence(ctx, presence)
}

// JoinVoiceChannel joins a voice channel.
func (c *Client) JoinVoiceChannel(ctx context.Context, guildID, channelID string) error {
	return c.gateway.UpdateVoiceState(ctx, &gateway.VoiceStateUpdateData{
		GuildID:   guildID,
		ChannelID: &channelID,
		SelfMute:  false,
		SelfDeaf:  false,
	})
}

// LeaveVoiceChannel leaves the voice channel in a guild.
func (c *Client) LeaveVoiceChannel(ctx context.Context, guildID string) error {
	return c.gateway.UpdateVoiceState(ctx, &gateway.VoiceStateUpdateData{
		GuildID:   guildID,
		ChannelID: nil,
		SelfMute:  false,
		SelfDeaf:  false,
	})
}

// UpdateVoiceState updates the bot's voice state with full control.
func (c *Client) UpdateVoiceState(ctx context.Context, opts *gateway.VoiceStateUpdateData) error {
	return c.gateway.UpdateVoiceState(ctx, opts)
}

// RequestGuildMembers requests guild members from Discord.
// Results are received via GUILD_MEMBERS_CHUNK events.
// Requires GUILD_MEMBERS intent.
func (c *Client) RequestGuildMembers(ctx context.Context, guildID string, query string, limit int) error {
	return c.gateway.RequestGuildMembers(ctx, &gateway.RequestGuildMembersData{
		GuildID: guildID,
		Query:   &query,
		Limit:   limit,
	})
}

// RequestGuildMembersByID requests specific guild members by user ID.
// Results are received via GUILD_MEMBERS_CHUNK events.
// Requires GUILD_MEMBERS intent.
func (c *Client) RequestGuildMembersByID(ctx context.Context, guildID string, userIDs []string) error {
	return c.gateway.RequestGuildMembers(ctx, &gateway.RequestGuildMembersData{
		GuildID: guildID,
		UserIDs: userIDs,
		Limit:   0,
	})
}

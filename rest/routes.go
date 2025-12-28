package rest

import (
	"context"
	"fmt"

	"github.com/kolosys/discord/models"
	"github.com/kolosys/neuron"
)

// GetGateway retrieves the gateway URL for connecting to the Discord gateway.
func (r *REST) GetGateway(ctx context.Context) (*models.Gateway, error) {
	if err := r.wait(ctx, "GET", "/gateway"); err != nil {
		return nil, err
	}
	resp, err := r.client.GetGateway(ctx)
	if err != nil {
		return nil, fmt.Errorf("get gateway: %w", err)
	}
	return &resp.Data, nil
}

// GetBotGateway retrieves bot-specific gateway information including sharding.
func (r *REST) GetBotGateway(ctx context.Context) (*models.GatewayBot, error) {
	if err := r.wait(ctx, "GET", "/gateway/bot"); err != nil {
		return nil, err
	}
	resp, err := r.client.GetBotGateway(ctx)
	if err != nil {
		return nil, fmt.Errorf("get bot gateway: %w", err)
	}
	return &resp.Data, nil
}

// GetMyApplication retrieves the current application's information.
func (r *REST) GetMyApplication(ctx context.Context) (*models.PrivateApplication, error) {
	if err := r.wait(ctx, "GET", "/applications/@me"); err != nil {
		return nil, err
	}
	resp, err := r.client.GetMyApplication(ctx)
	if err != nil {
		return nil, fmt.Errorf("get my application: %w", err)
	}
	return &resp.Data, nil
}

// GetUser retrieves a user by ID.
func (r *REST) GetUser(ctx context.Context, userID string) (*models.User, error) {
	if err := r.wait(ctx, "GET", "/users/"+userID); err != nil {
		return nil, err
	}
	resp, err := r.client.GetUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	return &resp.Data, nil
}

// GetCurrentUser retrieves the current bot user.
func (r *REST) GetCurrentUser(ctx context.Context) (*models.UserPII, error) {
	if err := r.wait(ctx, "GET", "/users/@me"); err != nil {
		return nil, err
	}
	resp, err := r.client.GetMyUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("get current user: %w", err)
	}
	return &resp.Data, nil
}

// GetGuild retrieves a guild by ID.
func (r *REST) GetGuild(ctx context.Context, guildID string) (*models.GuildWithCounts, error) {
	if err := r.wait(ctx, "GET", "/guilds/"+guildID); err != nil {
		return nil, err
	}
	resp, err := r.client.GetGuild(ctx, guildID)
	if err != nil {
		return nil, fmt.Errorf("get guild: %w", err)
	}
	return &resp.Data, nil
}

// ListGuildChannels retrieves a list of channels for a guild.
func (r *REST) ListGuildChannels(ctx context.Context, guildID string) ([]models.GuildChannel, error) {
	if err := r.wait(ctx, "GET", "/guilds/"+guildID+"/channels"); err != nil {
		return nil, err
	}
	resp, err := r.client.ListGuildChannels(ctx, guildID)
	if err != nil {
		return nil, fmt.Errorf("list guild channels: %w", err)
	}
	return resp.Data, nil
}

// GetChannel retrieves a channel by ID.
func (r *REST) GetChannel(ctx context.Context, channelID string) (any, error) {
	if err := r.wait(ctx, "GET", "/channels/"+channelID); err != nil {
		return nil, err
	}
	resp, err := r.client.GetChannel(ctx, channelID)
	if err != nil {
		return nil, fmt.Errorf("get channel: %w", err)
	}
	return resp.Data, nil
}

// CreateMessage sends a message to a channel.
func (r *REST) CreateMessage(ctx context.Context, channelID string, msg models.MessageCreateOptions) (*models.Message, error) {
	if err := r.wait(ctx, "POST", "/channels/"+channelID+"/messages"); err != nil {
		return nil, err
	}
	resp, err := r.client.CreateMessage(ctx, channelID, msg)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}
	return &resp.Data, nil
}

// GetMessage retrieves a message by ID.
func (r *REST) GetMessage(ctx context.Context, channelID, messageID string) (*models.Message, error) {
	if err := r.wait(ctx, "GET", "/channels/"+channelID+"/messages/"+messageID); err != nil {
		return nil, err
	}
	resp, err := r.client.GetMessage(ctx, channelID, messageID)
	if err != nil {
		return nil, fmt.Errorf("get message: %w", err)
	}
	return &resp.Data, nil
}

// ListMessagesParams configures ListMessages requests.
type ListMessagesParams struct {
	Around *string
	Before *string
	After  *string
	Limit  int
}

// ListMessages retrieves messages from a channel.
func (r *REST) ListMessages(ctx context.Context, channelID string, params *ListMessagesParams) ([]models.Message, error) {
	if err := r.wait(ctx, "GET", "/channels/"+channelID+"/messages"); err != nil {
		return nil, err
	}

	var reqOpts *neuron.RequestOptions
	if params != nil {
		query := make(map[string]any)
		if params.Around != nil {
			query["around"] = params.Around
		}
		if params.Before != nil {
			query["before"] = params.Before
		}
		if params.After != nil {
			query["after"] = params.After
		}
		if params.Limit > 0 {
			limit := params.Limit
			if limit > 100 {
				limit = 100
			}
			query["limit"] = limit
		}
		if len(query) > 0 {
			reqOpts = &neuron.RequestOptions{Query: query}
		}
	}

	resp, err := r.client.ListMessages(ctx, channelID, reqOpts)
	if err != nil {
		return nil, fmt.Errorf("list messages: %w", err)
	}
	return resp.Data, nil
}

// GetGuildMember retrieves a guild member by user ID.
func (r *REST) GetGuildMember(ctx context.Context, guildID, userID string) (*models.GuildMember, error) {
	if err := r.wait(ctx, "GET", "/guilds/"+guildID+"/members/"+userID); err != nil {
		return nil, err
	}
	resp, err := r.client.GetGuildMember(ctx, guildID, userID)
	if err != nil {
		return nil, fmt.Errorf("get guild member: %w", err)
	}
	return &resp.Data, nil
}

// ListGuildMembers retrieves members from a guild.
func (r *REST) ListGuildMembers(ctx context.Context, guildID string) ([]models.GuildMember, error) {
	if err := r.wait(ctx, "GET", "/guilds/"+guildID+"/members"); err != nil {
		return nil, err
	}
	resp, err := r.client.ListGuildMembers(ctx, guildID)
	if err != nil {
		return nil, fmt.Errorf("list guild members: %w", err)
	}
	return resp.Data, nil
}

// Interaction endpoints

// CreateInteractionResponse responds to an interaction.
func (r *REST) CreateInteractionResponse(ctx context.Context, interactionID, token string, response any) error {
	if err := r.wait(ctx, "POST", "/interactions/"+interactionID+"/callback"); err != nil {
		return err
	}
	_, err := r.client.CreateInteractionResponse(ctx, models.SnowflakeType(interactionID), token, response)
	if err != nil {
		return fmt.Errorf("create interaction response: %w", err)
	}
	return nil
}

// EditOriginalInteractionResponse edits the original interaction response.
func (r *REST) EditOriginalInteractionResponse(ctx context.Context, appID, token string, edit *models.IncomingWebhookUpdateOptionsPartial) (*models.Message, error) {
	if err := r.wait(ctx, "PATCH", "/webhooks/"+appID+"/"+token+"/messages/@original"); err != nil {
		return nil, err
	}
	if edit == nil {
		edit = &models.IncomingWebhookUpdateOptionsPartial{}
	}
	resp, err := r.client.UpdateOriginalWebhookMessage(ctx, models.SnowflakeType(appID), token, *edit)
	if err != nil {
		return nil, fmt.Errorf("edit original interaction response: %w", err)
	}
	return &resp.Data, nil
}

// DeleteOriginalInteractionResponse deletes the original interaction response.
func (r *REST) DeleteOriginalInteractionResponse(ctx context.Context, appID, token string) error {
	if err := r.wait(ctx, "DELETE", "/webhooks/"+appID+"/"+token+"/messages/@original"); err != nil {
		return err
	}
	_, err := r.client.DeleteOriginalWebhookMessage(ctx, models.SnowflakeType(appID), token)
	if err != nil {
		return fmt.Errorf("delete original interaction response: %w", err)
	}
	return nil
}

// CreateFollowupMessage creates a followup message for an interaction.
func (r *REST) CreateFollowupMessage(ctx context.Context, appID, token string, message any) (*models.Message, error) {
	if err := r.wait(ctx, "POST", "/webhooks/"+appID+"/"+token); err != nil {
		return nil, err
	}
	resp, err := r.client.ExecuteWebhook(ctx, models.SnowflakeType(appID), token, message)
	if err != nil {
		return nil, fmt.Errorf("create followup message: %w", err)
	}
	return &resp.Data, nil
}

// Application command endpoints

// BulkOverwriteGlobalCommands replaces all global application commands.
func (r *REST) BulkOverwriteGlobalCommands(ctx context.Context, appID string, commands any) ([]models.ApplicationCommand, error) {
	if err := r.wait(ctx, "PUT", "/applications/"+appID+"/commands"); err != nil {
		return nil, err
	}
	resp, err := r.client.BulkSetApplicationCommands(ctx, models.SnowflakeType(appID), commands)
	if err != nil {
		return nil, fmt.Errorf("bulk overwrite global commands: %w", err)
	}
	return resp.Data, nil
}

// BulkOverwriteGuildCommands replaces all guild application commands.
func (r *REST) BulkOverwriteGuildCommands(ctx context.Context, appID, guildID string, commands any) ([]models.ApplicationCommand, error) {
	if err := r.wait(ctx, "PUT", "/applications/"+appID+"/guilds/"+guildID+"/commands"); err != nil {
		return nil, err
	}
	resp, err := r.client.BulkSetGuildApplicationCommands(ctx, models.SnowflakeType(appID), models.SnowflakeType(guildID), commands)
	if err != nil {
		return nil, fmt.Errorf("bulk overwrite guild commands: %w", err)
	}
	return resp.Data, nil
}

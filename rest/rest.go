package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kolosys/discord/models"
	client "github.com/kolosys/discord/rest/internal"
	"github.com/kolosys/discord/types"
	"github.com/kolosys/neuron"
)

const (
	DefaultBaseURL   = "https://discord.com/api/v10"
	DefaultUserAgent = "DiscordBot (kolosys/discord, v0.1.0)"
	DefaultTimeout   = 30 * time.Second
)

// REST provides access to the Discord REST API with rate limiting and authentication.
type REST struct {
	client  *client.RESTClient
	token   string
	limiter *DiscordRateLimiter
}

// Options holds REST client configuration.
type Options struct {
	BaseURL          string
	UserAgent        string
	Timeout          time.Duration
	RateLimiter      *RateLimiterConfig
	DisableRateLimit bool
}

// New creates a new REST client with the given token and options.
func New(token string, opts *Options) *REST {
	if opts == nil {
		opts = &Options{}
	}
	if opts.BaseURL == "" {
		opts.BaseURL = DefaultBaseURL
	}
	if opts.UserAgent == "" {
		opts.UserAgent = DefaultUserAgent
	}
	if opts.Timeout == 0 {
		opts.Timeout = DefaultTimeout
	}

	headers := make(http.Header)
	headers.Set("Authorization", "Bot "+token)

	c := client.NewRESTClient(client.ClientOptions{
		BaseURL:   opts.BaseURL,
		UserAgent: opts.UserAgent,
		Timeout:   opts.Timeout,
		Headers:   headers,
	})

	var limiter *DiscordRateLimiter
	if !opts.DisableRateLimit {
		limiter = NewRateLimiter(opts.RateLimiter)
	}

	return &REST{
		client:  c,
		token:   token,
		limiter: limiter,
	}
}

// Token returns the bot token (masked for security).
func (r *REST) Token() string {
	if len(r.token) < 20 {
		return "***"
	}
	return r.token[:10] + "..." + r.token[len(r.token)-5:]
}

// RateLimiter returns the rate limiter for this client.
func (r *REST) RateLimiter() *DiscordRateLimiter {
	return r.limiter
}

func (r *REST) wait(ctx context.Context, method, endpoint string) error {
	if r.limiter == nil {
		return nil
	}
	return r.limiter.Wait(ctx, method, endpoint)
}

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
func (r *REST) GetUser(ctx context.Context, userID types.Snowflake) (*models.User, error) {
	if err := r.wait(ctx, "GET", "/users/"+userID.String()); err != nil {
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
func (r *REST) GetGuild(ctx context.Context, guildID types.Snowflake) (*models.GuildWithCounts, error) {
	if err := r.wait(ctx, "GET", "/guilds/"+guildID.String()); err != nil {
		return nil, err
	}
	resp, err := r.client.GetGuild(ctx, guildID)
	if err != nil {
		return nil, fmt.Errorf("get guild: %w", err)
	}
	return &resp.Data, nil
}

// GetChannel retrieves a channel by ID.
func (r *REST) GetChannel(ctx context.Context, channelID types.Snowflake) (any, error) {
	if err := r.wait(ctx, "GET", "/channels/"+channelID.String()); err != nil {
		return nil, err
	}
	resp, err := r.client.GetChannel(ctx, channelID)
	if err != nil {
		return nil, fmt.Errorf("get channel: %w", err)
	}
	return resp.Data, nil
}

// CreateMessage sends a message to a channel.
func (r *REST) CreateMessage(ctx context.Context, channelID types.Snowflake, msg models.MessageCreateOptions) (*models.Message, error) {
	if err := r.wait(ctx, "POST", "/channels/"+channelID.String()+"/messages"); err != nil {
		return nil, err
	}
	resp, err := r.client.CreateMessage(ctx, channelID, msg)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}
	return &resp.Data, nil
}

// GetMessage retrieves a message by ID.
func (r *REST) GetMessage(ctx context.Context, channelID, messageID types.Snowflake) (*models.Message, error) {
	if err := r.wait(ctx, "GET", "/channels/"+channelID.String()+"/messages/"+messageID.String()); err != nil {
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
	Around *types.Snowflake
	Before *types.Snowflake
	After  *types.Snowflake
	Limit  int
}

// ListMessages retrieves messages from a channel.
func (r *REST) ListMessages(ctx context.Context, channelID types.Snowflake, params *ListMessagesParams) ([]models.Message, error) {
	if err := r.wait(ctx, "GET", "/channels/"+channelID.String()+"/messages"); err != nil {
		return nil, err
	}

	var reqOpts *neuron.RequestOptions
	if params != nil {
		query := make(map[string]any)
		if params.Around != nil {
			query["around"] = params.Around.String()
		}
		if params.Before != nil {
			query["before"] = params.Before.String()
		}
		if params.After != nil {
			query["after"] = params.After.String()
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
func (r *REST) GetGuildMember(ctx context.Context, guildID, userID types.Snowflake) (*models.GuildMember, error) {
	if err := r.wait(ctx, "GET", "/guilds/"+guildID.String()+"/members/"+userID.String()); err != nil {
		return nil, err
	}
	resp, err := r.client.GetGuildMember(ctx, guildID, userID)
	if err != nil {
		return nil, fmt.Errorf("get guild member: %w", err)
	}
	return &resp.Data, nil
}

// ListGuildMembers retrieves members from a guild.
func (r *REST) ListGuildMembers(ctx context.Context, guildID types.Snowflake) ([]models.GuildMember, error) {
	if err := r.wait(ctx, "GET", "/guilds/"+guildID.String()+"/members"); err != nil {
		return nil, err
	}
	resp, err := r.client.ListGuildMembers(ctx, guildID)
	if err != nil {
		return nil, fmt.Errorf("list guild members: %w", err)
	}
	return resp.Data, nil
}

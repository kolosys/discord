package rest

import (
	"context"
	"net/http"
	"time"

	client "github.com/kolosys/discord/rest/internal"
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

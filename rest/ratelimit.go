package rest

import (
	"context"

	"github.com/kolosys/ion/ratelimit"
	"github.com/kolosys/neuron"
)

// DiscordRateLimiter handles Discord's complex rate limiting including bucket-based limits.
type DiscordRateLimiter struct {
	limiter *ratelimit.MultiTierLimiter
}

// NewRateLimiter creates a new Discord rate limiter.
func NewRateLimiter() *DiscordRateLimiter {
	return &DiscordRateLimiter{}
}

// Allow checks if a request can be made without blocking.
func (d *DiscordRateLimiter) Allow(ctx context.Context, method, endpoint string) bool {
	if d.limiter == nil {
		return true
	}
	// TODO: implement bucket-based rate limiting
	return true
}

// Wait blocks until the request can be made or context is cancelled.
func (d *DiscordRateLimiter) Wait(ctx context.Context, method, endpoint string) error {
	if d.limiter == nil {
		return nil
	}
	// TODO: implement bucket-based rate limiting
	return nil
}

// UpdateFromHeaders updates rate limit state from response headers.
func (d *DiscordRateLimiter) UpdateFromHeaders(method, endpoint string, info *neuron.RateLimitInfo) error {
	if d.limiter == nil || info == nil {
		return nil
	}
	// TODO: implement bucket tracking from headers
	return nil
}

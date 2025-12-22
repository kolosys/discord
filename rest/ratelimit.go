package rest

import (
	"context"

	"github.com/kolosys/ion/ratelimit"
	"github.com/kolosys/neuron"
)

type DiscordRateLimiter struct {
	limiter *ratelimit.MultiTierLimiter
}

func (d *DiscordRateLimiter) Allow(ctx context.Context, method, endpoint string) bool
func (d *DiscordRateLimiter) Wait(ctx context.Context, method, endpoint string) error
func (d *DiscordRateLimiter) UpdateFromHeaders(method, endpoint string, info *neuron.RateLimitInfo) error

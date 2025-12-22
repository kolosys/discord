package rest

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/kolosys/ion/ratelimit"
	"github.com/kolosys/neuron"
)

// DiscordRateLimiter adapts Ion's MultiTierLimiter to Neuron's RateLimitHandler interface
// for Discord's bucket-based rate limiting system.
type DiscordRateLimiter struct {
	limiter *ratelimit.MultiTierLimiter

	mu           sync.RWMutex
	bucketRoutes map[string]string // endpoint -> bucket
	globalPaused time.Time
}

// RateLimiterConfig configures the Discord rate limiter.
type RateLimiterConfig struct {
	// GlobalRate is the overall requests per second limit (Discord: 50/s)
	GlobalRate int
	// GlobalBurst is the burst capacity for the global limiter
	GlobalBurst int
	// DefaultRouteRate is the default per-route rate
	DefaultRouteRate int
	// DefaultRouteBurst is the default per-route burst
	DefaultRouteBurst int
}

// DefaultRateLimiterConfig returns Discord-appropriate defaults.
func DefaultRateLimiterConfig() *RateLimiterConfig {
	return &RateLimiterConfig{
		GlobalRate:        50,
		GlobalBurst:       50,
		DefaultRouteRate:  5,
		DefaultRouteBurst: 5,
	}
}

// NewRateLimiter creates a new Discord rate limiter.
func NewRateLimiter(cfg *RateLimiterConfig) *DiscordRateLimiter {
	if cfg == nil {
		cfg = DefaultRateLimiterConfig()
	}

	multiTierCfg := &ratelimit.MultiTierConfig{
		GlobalRate:           ratelimit.PerSecond(cfg.GlobalRate),
		GlobalBurst:          cfg.GlobalBurst,
		DefaultRouteRate:     ratelimit.PerSecond(cfg.DefaultRouteRate),
		DefaultRouteBurst:    cfg.DefaultRouteBurst,
		DefaultResourceRate:  ratelimit.PerSecond(cfg.DefaultRouteRate),
		DefaultResourceBurst: cfg.DefaultRouteBurst,
		EnableBucketMapping:  true,
		BucketTTL:            time.Hour,
		RoutePatterns:        discordRoutePatterns(),
	}

	return &DiscordRateLimiter{
		limiter:      ratelimit.NewMultiTierLimiter(multiTierCfg),
		bucketRoutes: make(map[string]string),
	}
}

// Allow checks if a request can proceed without blocking.
func (d *DiscordRateLimiter) Allow(ctx context.Context, method, endpoint string) bool {
	if d.isGloballyPaused() {
		return false
	}

	req := d.buildRequest(ctx, method, endpoint)
	return d.limiter.Allow(req)
}

// Wait blocks until the request is allowed or context is cancelled.
func (d *DiscordRateLimiter) Wait(ctx context.Context, method, endpoint string) error {
	if err := d.waitForGlobalPause(ctx); err != nil {
		return err
	}

	req := d.buildRequest(ctx, method, endpoint)
	return d.limiter.Wait(req)
}

// UpdateFromHeaders updates rate limit state from Discord response headers.
func (d *DiscordRateLimiter) UpdateFromHeaders(method, endpoint string, info *neuron.RateLimitInfo) error {
	if info == nil {
		return nil
	}

	// Handle global rate limit
	if info.Global {
		d.pauseGlobal(info.WaitDuration())
		return nil
	}

	// Track bucket mapping
	if info.Bucket != "" {
		d.mu.Lock()
		d.bucketRoutes[method+":"+endpoint] = info.Bucket
		d.mu.Unlock()
	}

	// Convert to header map for Ion
	headers := map[string]string{}
	if info.Limit > 0 {
		headers["X-RateLimit-Limit"] = itoa(info.Limit)
	}
	if info.Remaining >= 0 {
		headers["X-RateLimit-Remaining"] = itoa(info.Remaining)
	}
	if info.ResetAfter > 0 {
		headers["X-RateLimit-Reset-After"] = ftoa(info.ResetAfter.Seconds())
	}
	if info.Bucket != "" {
		headers["X-RateLimit-Bucket"] = info.Bucket
	}

	req := d.buildRequest(context.Background(), method, endpoint)
	return d.limiter.UpdateRateLimitFromHeaders(req, headers)
}

// Metrics returns rate limiting metrics.
func (d *DiscordRateLimiter) Metrics() *ratelimit.MultiTierMetrics {
	return d.limiter.GetMetrics()
}

// Reset clears all rate limit state.
func (d *DiscordRateLimiter) Reset() {
	d.limiter.Reset()

	d.mu.Lock()
	d.bucketRoutes = make(map[string]string)
	d.globalPaused = time.Time{}
	d.mu.Unlock()
}

func (d *DiscordRateLimiter) buildRequest(ctx context.Context, method, endpoint string) *ratelimit.Request {
	return &ratelimit.Request{
		Method:   method,
		Endpoint: endpoint,
		Context:  ctx,
	}
}

func (d *DiscordRateLimiter) isGloballyPaused() bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return !d.globalPaused.IsZero() && time.Now().Before(d.globalPaused)
}

func (d *DiscordRateLimiter) pauseGlobal(duration time.Duration) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.globalPaused = time.Now().Add(duration)
	d.limiter.PauseFor(duration)
}

func (d *DiscordRateLimiter) waitForGlobalPause(ctx context.Context) error {
	d.mu.RLock()
	pausedUntil := d.globalPaused
	d.mu.RUnlock()

	if pausedUntil.IsZero() || time.Now().After(pausedUntil) {
		return nil
	}

	wait := time.Until(pausedUntil)
	if wait <= 0 {
		return nil
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(wait):
		return nil
	}
}

// discordRoutePatterns returns rate limit configurations for known Discord routes.
func discordRoutePatterns() map[string]ratelimit.RouteConfig {
	return map[string]ratelimit.RouteConfig{
		// Channels - messages have stricter limits
		"POST:/channels/{id}/messages": {
			Rate:  ratelimit.PerSecond(5),
			Burst: 5,
		},
		"PATCH:/channels/{id}/messages/{id}": {
			Rate:  ratelimit.PerSecond(5),
			Burst: 5,
		},
		"DELETE:/channels/{id}/messages/{id}": {
			Rate:  ratelimit.PerSecond(5),
			Burst: 5,
		},

		// Reactions
		"PUT:/channels/{id}/messages/{id}/reactions/{id}/@me": {
			Rate:  ratelimit.Per(1, time.Millisecond*250),
			Burst: 1,
		},

		// Typing indicator
		"POST:/channels/{id}/typing": {
			Rate:  ratelimit.PerSecond(5),
			Burst: 5,
		},

		// Webhooks
		"POST:/webhooks/{id}/{id}": {
			Rate:  ratelimit.PerSecond(5),
			Burst: 5,
		},

		// Gateway
		"GET:/gateway": {
			Rate:  ratelimit.Per(1, time.Second),
			Burst: 1,
		},
		"GET:/gateway/bot": {
			Rate:  ratelimit.Per(1, time.Second),
			Burst: 1,
		},
	}
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func ftoa(f float64) string {
	return strconv.FormatFloat(f, 'f', 3, 64)
}

package discord

import (
	"context"
	"fmt"
	"sync"

	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/discord/rest"
	"github.com/kolosys/ion/workerpool"
	"github.com/kolosys/nova/bus"
)

// Client is the main Discord client that manages REST and Gateway connections.
type Client struct {
	token   string
	rest    *rest.REST
	gateway *gateway.Gateway
	events  bus.EventBus
	worker  *workerpool.Pool

	mu      sync.RWMutex
	running bool
}

// Options configures the Discord client.
type Options struct {
	Token   string
	Intents gateway.Intent
	REST    *rest.Options
}

// New creates a new Discord client.
func New(opts *Options) (*Client, error) {
	if opts == nil || opts.Token == "" {
		return nil, fmt.Errorf("discord: token is required")
	}

	// Create REST client
	r := rest.New(opts.Token, opts.REST)

	// Create worker pool for event processing
	worker := workerpool.New(10, 100, workerpool.WithName("discord-events"))

	// Create event bus
	eventBus := bus.New(bus.Config{
		WorkerPool:          worker,
		DefaultBufferSize:   1000,
		DefaultPartitions:   4,
		DefaultDeliveryMode: bus.AtLeastOnce,
		Name:                "discord",
	})

	// Create gateway
	gw := gateway.NewGateway(opts.Token, opts.Intents, eventBus)

	return &Client{
		token:   opts.Token,
		rest:    r,
		gateway: gw,
		events:  eventBus,
		worker:  worker,
	}, nil
}

// REST returns the REST client for making API requests.
func (c *Client) REST() *rest.REST {
	return c.rest
}

// Gateway returns the Gateway client for real-time events.
func (c *Client) Gateway() *gateway.Gateway {
	return c.gateway
}

// Events returns the event bus for subscribing to Discord events.
func (c *Client) Events() bus.EventBus {
	return c.events
}

// Start connects to the Discord gateway and begins receiving events.
func (c *Client) Start(ctx context.Context) error {
	c.mu.Lock()
	if c.running {
		c.mu.Unlock()
		return fmt.Errorf("discord: client already running")
	}
	c.running = true
	c.mu.Unlock()

	// Get gateway URL and shard count from Discord
	gatewayInfo, err := c.rest.GetBotGateway(ctx)
	if err != nil {
		c.mu.Lock()
		c.running = false
		c.mu.Unlock()
		return fmt.Errorf("discord: failed to get gateway info: %w", err)
	}

	// Connect to gateway with recommended shard count
	shardCount := int(gatewayInfo.Shards)
	if shardCount <= 0 {
		shardCount = 1
	}

	if err := c.gateway.Connect(ctx, gatewayInfo.URL, shardCount); err != nil {
		c.mu.Lock()
		c.running = false
		c.mu.Unlock()
		return fmt.Errorf("discord: failed to connect to gateway: %w", err)
	}

	return nil
}

// Stop gracefully disconnects from the Discord gateway.
func (c *Client) Stop(ctx context.Context) error {
	c.mu.Lock()
	if !c.running {
		c.mu.Unlock()
		return nil
	}
	c.mu.Unlock()

	// Close gateway connection
	if err := c.gateway.Close(); err != nil {
		return fmt.Errorf("discord: failed to close gateway: %w", err)
	}

	// Shutdown event bus
	if err := c.events.Shutdown(ctx); err != nil {
		return fmt.Errorf("discord: failed to shutdown event bus: %w", err)
	}

	// Close worker pool
	if err := c.worker.Close(ctx); err != nil {
		return fmt.Errorf("discord: failed to close worker pool: %w", err)
	}

	c.mu.Lock()
	c.running = false
	c.mu.Unlock()

	return nil
}

// IsRunning returns whether the client is currently connected.
func (c *Client) IsRunning() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.running
}

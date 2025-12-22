package discord

import (
	"context"
	"fmt"
	"sync"

	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/discord/rest"
)

// Client is the main Discord client that manages REST and Gateway connections.
type Client struct {
	token   string
	rest    *rest.REST
	gateway *gateway.Gateway

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

	r := rest.New(opts.Token, opts.REST)

	return &Client{
		token: opts.Token,
		rest:  r,
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

// Start connects to the Discord gateway and begins receiving events.
func (c *Client) Start(ctx context.Context) error {
	c.mu.Lock()
	if c.running {
		c.mu.Unlock()
		return fmt.Errorf("discord: client already running")
	}
	c.running = true
	c.mu.Unlock()

	// Get gateway URL from Discord
	gatewayInfo, err := c.rest.GetBotGateway(ctx)
	if err != nil {
		c.mu.Lock()
		c.running = false
		c.mu.Unlock()
		return fmt.Errorf("discord: failed to get gateway info: %w", err)
	}

	_ = gatewayInfo // TODO: use for gateway connection

	return nil
}

// Stop gracefully disconnects from the Discord gateway.
func (c *Client) Stop(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.running {
		return nil
	}

	c.running = false

	// TODO: close gateway connection

	return nil
}

// IsRunning returns whether the client is currently connected.
func (c *Client) IsRunning() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.running
}

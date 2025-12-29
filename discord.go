package discord

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/kolosys/discord/commands"
	"github.com/kolosys/discord/events"
	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/discord/rest"
	"github.com/kolosys/discord/state"
	"github.com/kolosys/helix"
	"github.com/kolosys/ion/workerpool"
	"github.com/kolosys/nova/bus"
)

const discord_banner = `
    ____  _                          __
   / __ \(_)_____________  _________/ /
  / / / / / ___/ ___/ __ \/ ___/ __  / 
 / /_/ / (__  ) /__/ /_/ / /  / /_/ /  
/_____/_/____/\___/\____/_/   \__,_/   
Discord for Go (kolosys/discord, v0.1.0)
Developer friendly Discord SDK using (helix v{version})
_________________________________________  
`

// Bot is a unified Discord bot that combines:
// - Discord Gateway (real-time events via WebSocket)
// - Discord REST client (API calls)
// - Helix HTTP server (webhooks, interactions, admin API) [optional]
// - State cache (automatic entity caching) [enabled by default]
type Bot struct {
	*helix.Server // Embedded Helix server (nil if HTTP disabled)

	REST    *rest.REST       // Discord REST client
	Gateway *gateway.Gateway // Discord Gateway client
	Bus     bus.EventBus     // Event bus for advanced usage
	State   *state.State     // State cache (nil if disabled)

	token         string
	applicationID string
	dispatcher    *events.Dispatcher
	worker        *workerpool.Pool
	router        *commands.Router

	httpEnabled bool
	mu          sync.RWMutex
	running     bool
}

// Options configures the Discord bot.
type Options struct {
	// Discord configuration (required)
	Token   string
	Intents gateway.Intent

	// HTTP server configuration (optional)
	// If Addr is set, an HTTP server will be started alongside the gateway
	Addr     string // HTTP server address (e.g., ":8080")
	BasePath string // Base path for routes (default: "")

	// REST client options
	REST *rest.Options

	// Server options (optional, for advanced HTTP configuration)
	Server *helix.Options

	// State options (optional, for cache configuration)
	State        *state.Options
	DisableState bool // Set to true to disable state caching entirely
}

// New creates a new Discord bot.
// If Options.Addr is set, the bot will also run an HTTP server.
func New(opts *Options) (*Bot, error) {
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

	// Create event dispatcher
	dispatcher := events.NewDispatcher(eventBus)

	// Create gateway
	gw := gateway.NewGateway(opts.Token, opts.Intents, eventBus)

	// Create state cache if not disabled
	var st *state.State
	if !opts.DisableState {
		st = state.New(r, opts.State)
		if err := st.RegisterHandlers(dispatcher); err != nil {
			return nil, fmt.Errorf("discord: failed to register state handlers: %w", err)
		}
	}

	bot := &Bot{
		REST:       r,
		Gateway:    gw,
		Bus:        eventBus,
		State:      st,
		token:      opts.Token,
		dispatcher: dispatcher,
		worker:     worker,
		router:     commands.NewRouter(),
	}

	// Create HTTP server if address is specified
	if opts.Addr != "" {
		helixOpts := opts.Server
		if helixOpts == nil {
			helixOpts = &helix.Options{}
		}
		helixOpts.Addr = opts.Addr
		if opts.BasePath != "" {
			helixOpts.BasePath = opts.BasePath
		}

		// Setup Discord SDK banner
		helixOpts.Banner = discord_banner
		helixOpts.HideBanner = false

		// Setup auto port
		helixOpts.AutoPort = true
		helixOpts.MaxPortAttempts = 10

		bot.Server = helix.Default(helixOpts)
		bot.httpEnabled = true
	}

	return bot, nil
}

// Start connects to the Discord gateway and optionally starts the HTTP server.
func (b *Bot) Start(ctx context.Context) error {
	b.mu.Lock()
	if b.running {
		b.mu.Unlock()
		return fmt.Errorf("discord: bot already running")
	}
	b.running = true
	b.mu.Unlock()

	// Get current user to set application ID (for bots, user ID = application ID)
	user, err := b.REST.GetCurrentUser(ctx)
	if err != nil {
		b.mu.Lock()
		b.running = false
		b.mu.Unlock()
		return fmt.Errorf("discord: failed to get current user: %w", err)
	}
	b.SetApplicationID(user.ID)

	// Get gateway URL and shard count from Discord
	gatewayInfo, err := b.REST.GetBotGateway(ctx)
	if err != nil {
		b.mu.Lock()
		b.running = false
		b.mu.Unlock()
		return fmt.Errorf("discord: failed to get gateway info: %w", err)
	}

	// Connect to gateway
	shardCount := int(gatewayInfo.Shards)
	if shardCount <= 0 {
		shardCount = 1
	}

	if err := b.Gateway.Connect(ctx, gatewayInfo.URL, shardCount); err != nil {
		b.mu.Lock()
		b.running = false
		b.mu.Unlock()
		return fmt.Errorf("discord: failed to connect to gateway: %w", err)
	}

	// Start HTTP server if enabled
	if b.httpEnabled {
		go func() {
			if err := b.Server.Start(); err != nil && err != http.ErrServerClosed {
				fmt.Printf("discord: HTTP server error: %v\n", err)
			}
		}()
	}

	return nil
}

// Stop gracefully stops the bot.
func (b *Bot) Stop(ctx context.Context) error {
	b.mu.Lock()
	if !b.running {
		b.mu.Unlock()
		return nil
	}
	b.mu.Unlock()

	// Close event dispatcher
	if err := b.dispatcher.Close(); err != nil {
		return fmt.Errorf("discord: failed to close event dispatcher: %w", err)
	}

	// Close gateway connection
	if err := b.Gateway.Close(); err != nil {
		return fmt.Errorf("discord: failed to close gateway: %w", err)
	}

	// Close state cache if enabled
	if b.State != nil {
		b.State.Close()
	}

	// Shutdown HTTP server if enabled
	if b.httpEnabled {
		if err := b.Server.Shutdown(ctx); err != nil {
			return fmt.Errorf("discord: failed to shutdown HTTP server: %w", err)
		}
	}

	// Shutdown event bus
	if err := b.Bus.Shutdown(ctx); err != nil {
		return fmt.Errorf("discord: failed to shutdown event bus: %w", err)
	}

	// Close worker pool
	if err := b.worker.Close(ctx); err != nil {
		return fmt.Errorf("discord: failed to close worker pool: %w", err)
	}

	b.mu.Lock()
	b.running = false
	b.mu.Unlock()

	return nil
}

// IsRunning returns whether the bot is running.
func (b *Bot) IsRunning() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.running
}

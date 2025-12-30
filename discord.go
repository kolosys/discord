package discord

// Package discord provides a unified Discord bot framework that combines:
//   - Discord Gateway (real-time events via WebSocket)
//   - Discord REST client (API calls)
//   - Helix HTTP server (webhooks, interactions, admin API) [optional]
//   - State cache (automatic entity caching) [enabled by default]
//
// # Modular Architecture
//
// Large bots can organize features into modules that implement the Module interface.
// This enables clean separation of concerns, independent testing, and easy composition.
//
// Example module:
//
//	type ModerationModule struct {
//	    service *moderation.Service
//	}
//
//	func (m *ModerationModule) Register(bot *Bot) error {
//	    m.service = moderation.NewService()
//	    bot.Commands().RegisterService(m.service)
//	    bot.Slash("warn", "Warn a user", m.handleWarn)
//	    bot.OnMessageCreate(m.handleAutomod)
//	    return nil
//	}
//
// Usage:
//
//	bot, _ := discord.New(&discord.Options{
//	    Token:   token,
//	    Intents: gateway.IntentsPrivileged,
//	})
//
//	discord.MustRegisterModules(bot,
//	    &utils.Module{},
//	    &moderation.Module{},
//	    &leveling.Module{},
//	)
//
//	bot.Start(context.Background())
//
// See [Module] and examples/modular for more details.

import (
	"context"
	"fmt"
	"sync"

	"github.com/kolosys/discord/commands"
	"github.com/kolosys/discord/events"
	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/discord/rest"
	"github.com/kolosys/discord/server"
	"github.com/kolosys/discord/state"
	"github.com/kolosys/helix"
	"github.com/kolosys/ion/workerpool"
	"github.com/kolosys/nova/bus"
)

const Version = "0.1.0"

const discord_banner = `
    ____  _                          __
   / __ \(_)_____________  _________/ /
  / / / / / ___/ ___/ __ \/ ___/ __  /
 / /_/ / (__  ) /__/ /_/ / /  / /_/ /
/_____/_/____/\___/\____/_/   \__,_/
Discord for Go (kolosys/discord, v%s)
Developer friendly Discord SDK
_________________________________________
`

// Bot is a unified Discord bot that combines:
// - Discord Gateway (real-time events via WebSocket)
// - Discord REST client (API calls)
// - Helix HTTP server (webhooks, interactions, admin API) [optional]
// - State cache (automatic entity caching) [enabled by default]
type Bot struct {
	REST    *rest.REST       // Discord REST client
	Gateway *gateway.Gateway // Discord Gateway client
	Bus     bus.EventBus     // Event bus for advanced usage
	State   *state.State     // State cache (nil if disabled)

	token         string
	applicationID string
	dispatcher    *events.Dispatcher
	worker        *workerpool.Pool
	router        *commands.Router
	serverMgr     *server.Manager // Server lifecycle manager

	mu      sync.RWMutex
	running bool
}

// Options configures the Discord bot.
type Options struct {
	// Discord configuration (required)
	Token   string
	Intents gateway.Intent

	// REST client options
	REST *rest.Options

	// Server configuration (optional, for HTTP server and lifecycle management)
	Server *server.Options

	// State options (optional, for cache configuration)
	State        *state.Options
	DisableState bool // Set to true to disable state caching entirely
}

// New creates a new Discord bot.
// If Options.Server.EnableHTTP is true, the bot will run an embedded HTTP server.
// The HTTP server defaults to :8080 unless Options.Server.HelixOptions specifies a different address.
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

	// Create gateway with bot's dependencies
	gw := gateway.NewGateway(opts.Token, opts.Intents, &gateway.Options{
		WorkerPool: worker,
		EventBus:   eventBus,
	})

	// Get event dispatcher from gateway
	dispatcher := gw.Dispatcher()

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

	// Initialize command router responder and interaction handler
	bot.router.SetResponder(bot)
	bot.setupInteractionHandler()

	// Initialize server configuration if not provided
	if opts.Server == nil {
		opts.Server = &server.Options{}
	}

	// Set Discord-specific branding and callbacks
	opts.Server.Banner = fmt.Sprintf(discord_banner, Version)
	opts.Server.ConnectFn = bot.connectBot
	opts.Server.StopFn = bot.stopBot

	// Create server manager with configuration
	serverMgr, err := server.New(opts.Server)
	if err != nil {
		return nil, fmt.Errorf("discord: failed to create server manager: %w", err)
	}
	bot.serverMgr = serverMgr

	return bot, nil
}

// Start connects to the Discord gateway and starts the bot.
// This method blocks until the bot is shut down via signal (SIGINT/SIGTERM).
// If HTTP is enabled, the Helix HTTP server handles signal management.
// If HTTP is disabled, this method handles signals directly.
func (b *Bot) Start(ctx context.Context) error {
	b.mu.Lock()
	if b.running {
		b.mu.Unlock()
		return fmt.Errorf("discord: bot already running")
	}
	b.running = true
	b.mu.Unlock()

	return b.serverMgr.Start(ctx)
}

// connectBot connects to the Discord gateway.
// This is called by server.Manager during bot startup.
func (b *Bot) connectBot(ctx context.Context) error {
	// Get current user to set application ID (for bots, user ID = application ID)
	user, err := b.REST.GetCurrentUser(ctx)
	if err != nil {
		return fmt.Errorf("discord: failed to get current user: %w", err)
	}
	b.SetApplicationID(user.ID)

	// Get gateway URL and shard count from Discord
	gatewayInfo, err := b.REST.GetBotGateway(ctx)
	if err != nil {
		return fmt.Errorf("discord: failed to get gateway info: %w", err)
	}

	// Connect to gateway
	shardCount := int(gatewayInfo.Shards)
	if shardCount <= 0 {
		shardCount = 1
	}

	if err := b.Gateway.Connect(ctx, gatewayInfo.URL, shardCount); err != nil {
		return fmt.Errorf("discord: failed to connect to gateway: %w", err)
	}

	return nil
}

// stopBot gracefully stops bot components.
// This is called by server.Manager during shutdown.
func (b *Bot) stopBot(ctx context.Context) error {
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

// Stop gracefully stops the bot.
// If HTTP is enabled, this triggers server shutdown which invokes lifecycle hooks.
// If HTTP is disabled, this directly stops bot components.
func (b *Bot) Stop(ctx context.Context) error {
	b.mu.Lock()
	if !b.running {
		b.mu.Unlock()
		return nil
	}
	b.mu.Unlock()

	return b.serverMgr.Stop(ctx)
}

// IsRunning returns whether the bot is running.
func (b *Bot) IsRunning() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.running
}

// Server returns the embedded Helix HTTP server (nil if HTTP disabled).
//
// For convenience, common HTTP methods are available directly on Bot:
//
//	bot.GET("/health", handler)
//	bot.POST("/api/users", handler)
//	bot.Group("/api").GET("/users", handler)
//
// Use this method when you need direct access to advanced Helix features.
func (b *Bot) Server() *helix.Server {
	if b.serverMgr == nil {
		return nil
	}
	return b.serverMgr.HTTPServer()
}

package server

import (
	"context"

	"github.com/kolosys/helix"
)

// Options configures the Discord bot's HTTP server and lifecycle management.
// All fields are optional and will use sensible defaults if not specified.
//
// For HTTP-disabled bots, only ConnectFn and StopFn are required.
// For HTTP-enabled bots, the server defaults to listening on :8080.
type Options struct {
	// Enabled enables the embedded Helix HTTP server for webhooks and interactions.
	// If false, the bot runs in gateway-only mode.
	// Defaults to false.
	Enabled bool

	// HideBanner suppresses the startup banner for clean production logs.
	// Defaults to false.
	HideBanner bool

	// HelixOptions provides fine-grained control over the HTTP server configuration.
	// If EnableHTTP is true and HelixOptions is nil, the server defaults to ":8080".
	// See helix.Options for available configuration options.
	// Only used if EnableHTTP is true.
	HelixOptions *helix.Options

	// Banner is the startup banner text displayed when the bot starts.
	// Set by the bot itself - users should not set this field directly.
	// The banner is not displayed if HideBanner is true.
	Banner string

	// ConnectFn is the callback invoked during bot startup to initialize bot components.
	// Must be set by the caller (typically discord.New sets this to bot.connectBot).
	// Called before the HTTP server starts if EnableHTTP is true.
	// Must not be nil.
	ConnectFn func(context.Context) error

	// StopFn is the callback invoked during graceful shutdown to clean up bot resources.
	// Must be set by the caller (typically discord.New sets this to bot.stopBot).
	// Called before exiting after signal or Stop() call.
	// Must not be nil.
	StopFn func(context.Context) error
}

func (o *Options) applyDefaults() {
	if o.HelixOptions == nil {
		o.HelixOptions = &helix.Options{}
	}

	if o.HelixOptions.Addr == "" {
		o.HelixOptions.Addr = ":8080"
	}

	if o.Banner != "" && !o.HideBanner {
		o.HelixOptions.Banner = o.Banner
		o.HelixOptions.HideBanner = false
	} else if o.HideBanner {
		o.HelixOptions.HideBanner = true
	}

	o.HelixOptions.AutoPort = true
	o.HelixOptions.MaxPortAttempts = 10
}

# Discord for Go 📡

![GoVersion](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-blue.svg)
![Zero Dependencies](https://img.shields.io/badge/Zero-Dependencies-green.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/kolosys/helix.svg)](https://pkg.go.dev/github.com/kolosys/helix)
[![Go Report Card](https://goreportcard.com/badge/github.com/kolosys/helix)](https://goreportcard.com/report/github.com/kolosys/helix)

```
     ____  _                          __
    / __ \(_)_____________  _________/ /
   / / / / / ___/ ___/ __ \/ ___/ __  /
  / /_/ / (__  ) /__/ /_/ / /  / /_/ /
 /_____/_/____/\___/\____/_/   \__,_/
 Developer friendly Discord SDK
```

A high-performance, production-ready Discord API library built with the Kolosys ecosystem. Provides a unified bot interface combining real-time gateway events, REST API access, optional HTTP webhooks, and automatic state caching.

## Features

- **Unified Bot Interface**: Single `Bot` struct combining gateway, REST, state cache, and optional HTTP server
- **Real-Time Events**: WebSocket gateway with automatic reconnection and heartbeat handling
- **REST Client**: Full Discord API support with automatic rate limiting and error handling
- **Event Handlers**: Typed, generic event handlers with middleware support
- **Slash Commands**: Built-in command router and interaction handling
- **State Caching**: Automatic caching of guilds, channels, users, messages, and more
- **Message Components**: Buttons, select menus, text inputs with structured responses
- **Voice Support**: Audio streaming and voice channel integration
- **Zero Dependencies**: Standard library only (except Kolosys libraries)
- **Production-Ready**: Graceful shutdown, context cancellation, thread-safe operations

## Installation

```bash
go get github.com/kolosys/discord
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kolosys/discord"
	"github.com/kolosys/discord/events"
	"github.com/kolosys/discord/gateway"
)

func main() {
	// Create bot
	bot, err := discord.New(&discord.Options{
		Token:   os.Getenv("DISCORD_TOKEN"),
		Intents: gateway.IntentGuildMessages | gateway.IntentDirectMessages,
	})
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Register event handlers
	bot.OnReady(func(ctx context.Context, e *events.ReadyEvent) {
		fmt.Printf("Bot ready as %s\n", e.User.Username)
	})

	bot.OnMessageCreate(func(ctx context.Context, e *events.MessageCreateEvent) {
		if e.Author.ID == e.Author.ID { // Skip bot messages
			return
		}
		fmt.Printf("Message from %s: %s\n", e.Author.Username, e.Content)
	})

	// Start bot
	ctx := context.Background()
	if err := bot.Start(ctx); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Graceful shutdown
	if err := bot.Stop(ctx); err != nil {
		log.Fatalf("Failed to stop bot: %v", err)
	}
}
```

## Core Concepts

### Bot Structure

The `Bot` is the main entry point. It provides:

- **Gateway**: WebSocket connection to Discord for real-time events
- **REST**: HTTP client for API calls
- **State**: Automatic caching of Discord entities
- **Bus**: Event bus for advanced event routing
- **Server** (optional): Helix HTTP server for webhooks and interactions

### Event Handlers

Register handlers for Discord events using typed generics:

```go
// Typed handler with automatic unmarshaling
discord.Listen(bot, events.MessageCreate, func(ctx context.Context, e *events.MessageCreateEvent) {
	// Type-safe access to event data
	fmt.Println(e.Content)
})

// Raw handler for advanced use cases
bot.On(events.MessageCreate, func(ctx context.Context, t events.Type, data json.RawMessage) {
	// Raw JSON handling
})
```

### REST API

Use the REST client for synchronous API operations:

```go
// Create a message
msg, err := bot.REST.CreateMessage(ctx, channelID, &models.MessageCreate{
	Content: "Hello, World!",
})

// Get a user
user, err := bot.REST.GetUser(ctx, userID)

// Manage guilds
guilds, err := bot.REST.GetCurrentUserGuilds(ctx)
```

### State Cache

Automatically cache and retrieve Discord entities:

```go
// State is updated automatically from events
guild := bot.State.Guild(guildID)
if guild != nil {
	fmt.Printf("Guild: %s\n", guild.Name)
}

channel := bot.State.Channel(channelID)
user := bot.State.User(userID)
```

### Slash Commands

Register and handle slash commands:

```go
// Register a command
bot.Command("ping", "Ping the bot", func(ctx context.Context, i *discord.InteractionCreateEvent) {
	bot.REST.InteractionRespond(ctx, i.ID, i.Token, &models.InteractionResponse{
		Type: models.InteractionResponseTypePong,
	})
})
```

### HTTP Webhooks (Optional)

Enable an HTTP server for Discord interactions:

```go
bot, err := discord.New(&discord.Options{
	Token:   token,
	Intents: intents,
	Addr:    ":8080", // Enable HTTP server
	BasePath: "/discord",
})

// HTTP server will handle Discord webhooks and interactions automatically
```

## Architecture

The library is built on top of Kolosys libraries:

- **helix**: HTTP web framework for webhook handling
- **ion**: Worker pool for event processing
- **nova**: Event bus for message routing
- **axon**: WebSocket support
- **neuron**: HTTP client for REST API
- **synapse**: Caching for state management
- **atomic**: Thread-safe primitives

## Package Organization

| Package      | Purpose                                     |
| ------------ | ------------------------------------------- |
| `gateway`    | WebSocket connection and real-time events   |
| `rest`       | REST API client                             |
| `events`     | Event types and dispatcher                  |
| `commands`   | Slash command routing                       |
| `state`      | Entity caching                              |
| `models`     | Discord data models                         |
| `components` | Message components (buttons, selects, etc.) |
| `types`      | Utility types (Snowflake, etc.)             |

## Examples

Full examples are available in the `examples/` directory:

- `examples/gateway/`: Basic event handling
- `examples/rest/`: REST API usage
- `examples/commands/`: Slash command handling
- `examples/server/`: HTTP webhook server

## Performance

- Zero allocations in hot paths
- Efficient event dispatching with worker pools
- Connection pooling for REST requests
- Automatic rate limiting and request batching

## Configuration

### Bot Options

```go
&discord.Options{
	// Required
	Token:   "discord-token",
	Intents: gateway.IntentGuildMessages,

	// Optional HTTP server
	Addr:     ":8080",
	BasePath: "/discord",

	// Advanced configuration
	REST: &rest.Options{...},
	Server: &helix.Options{...},
	State: &state.Options{...},

	// Disable features
	DisableState: false,
}
```

### Gateway Intents

Specify which events to receive:

```go
intents := gateway.IntentGuildMessages |
	gateway.IntentDirectMessages |
	gateway.IntentGuildMembers |
	gateway.IntentMessageContent
```

## Thread Safety

The `Bot` is fully thread-safe. Event handlers are processed concurrently by the worker pool. Use context cancellation for graceful shutdown.

## Error Handling

All operations return errors following Go conventions:

```go
if err := bot.Start(ctx); err != nil {
	// Handle error
}

if err := bot.REST.CreateMessage(ctx, channelID, &msg); err != nil {
	// Handle API error
}
```

## License

See LICENSE file.

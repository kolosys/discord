# main

This example demonstrates basic usage of the library.

## Source Code

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
	"github.com/kolosys/discord/examples/internal"
	"github.com/kolosys/discord/gateway"
)

func main() {
	internal.LoadEnv(".env")

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN environment variable is required")
	}

	// Create Discord bot (gateway only, no HTTP server)
	bot, err := discord.New(&discord.Options{
		Token: token,
		Intents: gateway.IntentGuilds |
			gateway.IntentGuildMessages |
			gateway.IntentMessageContentPrivileged,
	})
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Subscribe to READY event
	bot.OnReady(func(ctx context.Context, e *events.ReadyEvent) {
		fmt.Println("\n=== BOT READY ===")
		fmt.Printf("User: %s#%s\n", e.User.Username, e.User.Discriminator)
		fmt.Printf("Session ID: %s\n", e.SessionID)
		fmt.Printf("Guilds: %d\n", len(e.Guilds))
		fmt.Printf("Shard ID: %d\n", e.ShardID)
		fmt.Println("================")

		// Set bot activity
		bot.SetActivity(ctx, gateway.ActivityTypeWatching, "for commands")
	})

	// Subscribe to MESSAGE_CREATE event
	bot.OnMessageCreate(func(ctx context.Context, e *events.MessageCreateEvent) {
		if !e.IsHuman() {
			return
		}

		author := e.AuthorUser()
		fmt.Printf("[MESSAGE] %s: %s\n", author.Username, e.Content)

		// Echo command
		if len(e.Content) > 6 && e.Content[:6] == "!echo " {
			reply := e.Content[6:]
			if _, err := bot.SendReply(ctx, e.ChannelID, e.ID, reply); err != nil {
				log.Printf("Failed to send reply: %v", err)
			}
		}
	})

	// Subscribe to GUILD_CREATE event
	bot.OnGuildCreate(func(ctx context.Context, e *events.GuildCreateEvent) {
		fmt.Printf("[GUILD] %s (%d members)\n", e.Name, e.MemberCount)
	})

	// Start the bot
	ctx := context.Background()
	if err := bot.Start(ctx); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	fmt.Println("Bot is running. Press Ctrl+C to exit.")

	// Wait for interrupt
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("\nShutting down...")
	if err := bot.Stop(ctx); err != nil {
		log.Printf("Error stopping bot: %v", err)
	}

	fmt.Println("Goodbye!")
}

```

## Running the Example

To run this example:

```bash
cd gateway
go run main.go
```

## Expected Output

```
Hello from Proton examples!
```

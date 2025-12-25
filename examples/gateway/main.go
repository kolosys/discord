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

	// Create Discord client with intents
	// Note: IntentMessageContent is a privileged intent that requires enabling in the Discord Developer Portal.
	// If you need message content, enable it at: https://discord.com/developers/applications → Bot → Privileged Gateway Intents
	client, err := discord.New(&discord.Options{
		Token: token,
		Intents: gateway.IntentGuilds |
			gateway.IntentGuildMessages,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Subscribe to READY event with typed handler
	client.OnReady(func(ctx context.Context, e *events.ReadyEvent) {
		fmt.Println("\n=== BOT READY ===")
		fmt.Printf("User: %s#%s\n", e.User.Username, e.User.Discriminator)
		fmt.Printf("Session ID: %s\n", e.SessionID)
		fmt.Printf("Guilds: %d\n", len(e.Guilds))
		fmt.Printf("Shard ID: %d\n", e.ShardID)
		fmt.Println("================")
	})

	// Subscribe to MESSAGE_CREATE event with typed handler
	client.OnMessageCreate(func(ctx context.Context, e *events.MessageCreateEvent) {
		// Skip bot messages
		if e.Author == nil {
			return
		}

		fmt.Printf("[MESSAGE] #%s: %s\n", e.ChannelID, e.Content)
	})

	// Subscribe to GUILD_CREATE event with typed handler
	client.OnGuildCreate(func(ctx context.Context, e *events.GuildCreateEvent) {
		fmt.Printf("[GUILD] %s (%d members)\n", e.Name, e.MemberCount)
	})

	// Subscribe to GUILD_MEMBER_ADD event with typed handler
	client.OnGuildMemberAdd(func(ctx context.Context, e *events.GuildMemberAddEvent) {
		fmt.Printf("[MEMBER+] New member joined guild %s\n", e.GuildID)
	})

	// Start the client
	ctx := context.Background()
	if err := client.Start(ctx); err != nil {
		log.Fatalf("Failed to start client: %v", err)
	}

	fmt.Println("Bot is running. Press Ctrl+C to exit.")

	// Wait for interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("\nShutting down...")
	if err := client.Stop(ctx); err != nil {
		log.Printf("Error stopping client: %v", err)
	}

	fmt.Println("Goodbye!")
}

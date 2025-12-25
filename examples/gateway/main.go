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
	// Privileged intents (GUILD_MEMBERS, GUILD_PRESENCES, MESSAGE_CONTENT) require enabling in Developer Portal:
	// https://discord.com/developers/applications → Bot → Privileged Gateway Intents
	client, err := discord.New(&discord.Options{
		Token: token,
		Intents: gateway.IntentGuilds |
			gateway.IntentGuildMessages |
			gateway.IntentMessageContentPrivileged, // Privileged: requires enabling in Developer Portal
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

		// Set bot activity - Note: Custom status (type 4) only works for user accounts, not bots
		// Valid bot activities: Playing, Streaming, Listening, Watching, Competing
		if err := client.SetActivity(ctx, gateway.ActivityTypeWatching, "for commands"); err != nil {
			log.Printf("Failed to set activity: %v", err)
		}
	})

	// Subscribe to MESSAGE_CREATE event with typed handler
	client.OnMessageCreate(func(ctx context.Context, e *events.MessageCreateEvent) {
		// Skip bot and system messages
		if !e.IsHuman() {
			return
		}

		author := e.AuthorUser()
		fmt.Printf("[MESSAGE] %s: %s\n", author.Username, e.Content)

		// Echo back messages that start with "!echo "
		if len(e.Content) > 6 && e.Content[:6] == "!echo " {
			reply := e.Content[6:]
			if _, err := client.SendReply(ctx, e.ChannelID, e.ID, reply); err != nil {
				log.Printf("Failed to send reply: %v", err)
			}
		}
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

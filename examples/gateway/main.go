package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kolosys/discord"
	"github.com/kolosys/discord/examples/internal"
	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/nova/shared"
)

func main() {
	// Load environment variables
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

	// Subscribe to READY event
	client.Events().Subscribe("READY", &EventListener{
		id: "ready-listener",
		handler: func(event shared.Event) error {
			fmt.Println("\n=== BOT READY ===")

			// Parse READY data
			var ready gateway.ReadyData
			if err := json.Unmarshal(event.Data().(json.RawMessage), &ready); err != nil {
				return fmt.Errorf("failed to unmarshal READY: %w", err)
			}

			fmt.Printf("User: %s#%s\n", ready.User.Username, ready.User.Discriminator)
			fmt.Printf("Session ID: %s\n", ready.SessionID)
			fmt.Printf("Guilds: %d\n", len(ready.Guilds))
			fmt.Println("================")

			return nil
		},
	})

	// Subscribe to MESSAGE_CREATE event
	client.Events().Subscribe("MESSAGE_CREATE", &EventListener{
		id: "message-listener",
		handler: func(event shared.Event) error {
			// You can unmarshal the event data to the appropriate Discord model
			// For now, just log that we received a message
			fmt.Printf("[MESSAGE_CREATE] Received message event\n")
			return nil
		},
	})

	// Subscribe to GUILD_CREATE event
	client.Events().Subscribe("GUILD_CREATE", &EventListener{
		id: "guild-listener",
		handler: func(event shared.Event) error {
			fmt.Printf("[GUILD_CREATE] Guild became available\n")
			return nil
		},
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

// EventListener implements the shared.Listener interface.
type EventListener struct {
	id      string
	handler func(shared.Event) error
}

func (l *EventListener) ID() string {
	return l.id
}

func (l *EventListener) Handle(event shared.Event) error {
	return l.handler(event)
}

func (l *EventListener) OnError(event shared.Event, err error) error {
	log.Printf("[%s] Error handling event %s: %v", l.id, event.Type(), err)
	return nil
}

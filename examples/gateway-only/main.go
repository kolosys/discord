package main

import (
	"context"
	"log"
	"os"

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

	// Create gateway with defaults
	// IntentGuildMessages to receive MESSAGE_CREATE events
	// IntentMessageContent to see message content (privileged, must be enabled in Discord Developer Portal)
	gw := gateway.NewGateway(token, gateway.IntentGuildMessages|gateway.IntentMessageContent, nil)

	// Register event handlers
	gw.OnReady(func(ctx context.Context, e *events.ReadyEvent) {
		log.Printf("Bot ready! Logged in as %s#%s", e.User.Username, e.User.Discriminator)
	})

	gw.OnMessageCreate(func(ctx context.Context, e *events.MessageCreateEvent) {
		if e.Content == "!ping" {
			author := e.AuthorUser()
			if author != nil {
				log.Printf("Received ping from %s", author.Username)
			}
		}
	})

	// Start and block until SIGINT/SIGTERM
	log.Println("Starting gateway...")
	if err := gw.Start(context.Background()); err != nil {
		log.Fatalf("Gateway error: %v", err)
	}
}

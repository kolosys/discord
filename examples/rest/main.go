package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kolosys/discord/examples/internal"
	"github.com/kolosys/discord/rest"
)

func main() {
	internal.LoadEnv(".env")

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN environment variable is required")
	}

	client := rest.New(token, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Testing Discord REST API...")
	fmt.Printf("Token: %s\n\n", client.Token())

	// Get current user (bot)
	user, err := client.GetCurrentUser(ctx)
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}
	fmt.Printf("Bot User: %s#%s (ID: %s)\n", user.Username, user.Discriminator, user.ID)

	// Get application info
	app, err := client.GetMyApplication(ctx)
	if err != nil {
		log.Fatalf("Failed to get application: %v", err)
	}
	fmt.Printf("Application: %s (ID: %s)\n", app.Name, app.ID)

	// Get gateway info
	gateway, err := client.GetBotGateway(ctx)
	if err != nil {
		log.Fatalf("Failed to get gateway: %v", err)
	}
	fmt.Printf("Gateway URL: %s\n", gateway.URL)
	fmt.Printf("Recommended Shards: %d\n", gateway.Shards)

	fmt.Println("\nREST API test complete!")
}

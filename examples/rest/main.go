package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/kolosys/discord/rest"
)

func main() {
	loadEnv(".env")

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

func loadEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		os.Setenv(strings.TrimSpace(key), strings.TrimSpace(value))
	}
}

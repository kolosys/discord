package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kolosys/discord/examples/internal"
	"github.com/kolosys/discord/models"
	"github.com/kolosys/discord/rest"
)

func main() {
	internal.LoadEnv(".env")

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN environment variable is required")
	}

	client := rest.New(token, nil)

	fmt.Println("Testing Discord REST API...")
	fmt.Printf("Token: %s\n\n", client.Token())

	guildID := os.Getenv("DISCORD_GUILD_ID")
	if guildID == "" {
		log.Fatal("DISCORD_GUILD_ID environment variable is required")
	}

	guild, err := client.GetGuild(context.Background(), guildID)
	if err != nil {
		log.Fatalf("Failed to get guild: %v", err)
	}

	channels, err := client.ListGuildChannels(context.Background(), guildID)
	if err != nil {
		log.Fatalf("Failed to get channels: %v", err)
	}

	var restTestChannel string
	for _, channel := range channels {
		if channel.Name == "rest-tests" {
			restTestChannel = channel.ID
			log.Printf("Rest test channel: %s (%s)", channel.Name, channel.ID)
			break
		}
	}

	log.Printf("Guild: %s", guild.Name)

	content := "Hello, world!"
	message, err := client.CreateMessage(context.Background(), restTestChannel, models.MessageCreateOptions{
		Content: &content,
	})
	if err != nil {
		log.Fatalf("Failed to create message: %v", err)
	}

	log.Printf("Message: %s", message.Content)
}

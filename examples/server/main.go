package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kolosys/discord"
	"github.com/kolosys/discord/events"
	"github.com/kolosys/discord/examples/internal"
	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/helix"
)

func main() {
	internal.LoadEnv(".env")

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN environment variable is required")
	}

	// Create Discord bot with HTTP server
	bot, err := discord.New(&discord.Options{
		Token: token,
		Intents: gateway.IntentGuilds |
			gateway.IntentGuildMessages |
			gateway.IntentMessageContentPrivileged,
		Addr: ":8080", // Enable HTTP server
	})
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// === HTTP Routes ===

	bot.GET("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"status":  "ok",
			"gateway": bot.Gateway.IsReady(),
			"shards":  bot.Gateway.ShardCount(),
		})
	})

	bot.GET("/api/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"running": bot.IsRunning(),
			"shards":  bot.Gateway.ShardCount(),
		})
	})

	bot.POST("/api/channels/{channelID}/messages", func(w http.ResponseWriter, r *http.Request) {
		channelID := r.PathValue("channelID")

		var req struct {
			Content string `json:"content"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			helix.WriteProblem(w, helix.ErrBadRequest)
			return
		}

		msg, err := bot.SendMessage(r.Context(), channelID, req.Content)
		if err != nil {
			helix.WriteProblem(w, helix.NewProblem(http.StatusBadGateway, "discord_error", err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
	})

	// === Discord Gateway Events ===

	bot.OnReady(func(ctx context.Context, e *events.ReadyEvent) {
		fmt.Println("\n=== BOT READY ===")
		fmt.Printf("User: %s#%s\n", e.User.Username, e.User.Discriminator)
		fmt.Printf("Guilds: %d\n", len(e.Guilds))
		fmt.Println("================")

		bot.SetActivity(ctx, gateway.ActivityTypeWatching, "for commands")
	})

	bot.OnMessageCreate(func(ctx context.Context, e *events.MessageCreateEvent) {
		if !e.IsHuman() {
			return
		}

		author := e.AuthorUser()
		fmt.Printf("[MESSAGE] %s: %s\n", author.Username, e.Content)

		if e.Content == "!ping" {
			bot.SendReply(ctx, e.ChannelID, e.ID, "Pong! 🏓")
		}
	})

	bot.OnGuildCreate(func(ctx context.Context, e *events.GuildCreateEvent) {
		fmt.Printf("[GUILD] %s (%d members)\n", e.Name, e.MemberCount)
	})

	// === Start Bot ===

	ctx := context.Background()
	if err := bot.Start(ctx); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	fmt.Println("\nBot running:")
	fmt.Println("  - Discord Gateway: connected")
	fmt.Println("  - HTTP Server: http://localhost:8080")
	fmt.Println("  - Health: http://localhost:8080/health")
	fmt.Println("\nPress Ctrl+C to exit.")

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

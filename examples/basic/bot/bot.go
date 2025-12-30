package bot

import (
	"os"

	"github.com/kolosys/discord"
	"github.com/kolosys/discord/examples/internal"
	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/discord/server"
)

func init() {
	internal.LoadEnv(".env")
}

func New() (*discord.Bot, error) {
	return discord.New(&discord.Options{
		Token:   os.Getenv("DISCORD_TOKEN"),
		Intents: gateway.IntentGuilds | gateway.IntentGuildMessages | gateway.IntentMessageContent,
		Server: &server.Options{
			Enabled: true,
		},
	})
}

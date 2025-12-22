package discord

import (
	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/discord/rest"
)

type DiscordClient struct {
	token   string
	gateway *gateway.Gateway
	rest    *rest.REST
}

func Bot(options BotOptions) *DiscordClient {
	return &DiscordClient{token: options.Token}
}

func (c *DiscordClient) Start() error {
	return nil
}

func (c *DiscordClient) Stop() error {
	return nil
}

package discord

import (
	"github.com/kolosys/discord/events"
	"github.com/kolosys/discord/types"
)

// Exports for developer convenience
type Snowflake = types.Snowflake

// Event type aliases for convenience
type ReadyEvent = events.ReadyEvent
type MessageCreateEvent = events.MessageCreateEvent
type GuildCreateEvent = events.GuildCreateEvent
type InteractionCreateEvent = events.InteractionCreateEvent

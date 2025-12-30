# events

This example demonstrates basic usage of the library.

## Source Code

```go
package events

import (
	"context"
	"log"

	"github.com/kolosys/discord"
	"github.com/kolosys/discord/events"
)

func Register(bot *discord.Bot) {
	bot.OnReady(onReady)
	bot.OnGuildCreate(onGuildCreate)
}

func onReady(ctx context.Context, e *events.ReadyEvent) {
	log.Printf("Bot ready! Logged in as %s", e.User.Username)
	log.Printf("Serving %d guilds", len(e.Guilds))
}

func onGuildCreate(ctx context.Context, e *events.GuildCreateEvent) {
	log.Printf("Joined guild: %s (%s)", e.Name, e.ID)
}

```

## Running the Example

To run this example:

```bash
cd events
go run events.go
```

## Expected Output

```
Hello from Proton examples!
```

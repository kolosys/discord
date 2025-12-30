# commands

This example demonstrates basic usage of the library.

## Source Code

```go
package commands

import (
	"github.com/kolosys/discord"
	"github.com/kolosys/discord/commands"
)

func Register(bot *discord.Bot) {
	bot.Slash("ping", "Check if the bot is alive", handlePing)
	bot.Slash("hello", "Say hello", handleHello)
}

func handlePing(ctx *commands.Context) error {
	return ctx.Reply("Pong!")
}

func handleHello(ctx *commands.Context) error {
	return ctx.Reply("Hello, " + ctx.User.Username + "!")
}

```

## Running the Example

To run this example:

```bash
cd commands
go run commands.go
```

## Expected Output

```
Hello from Proton examples!
```

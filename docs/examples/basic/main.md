# main

This example demonstrates basic usage of the library.

## Source Code

```go
package main

import (
	"context"
	"log"

	"github.com/kolosys/discord/examples/basic/bot"
	"github.com/kolosys/discord/examples/basic/commands"
	"github.com/kolosys/discord/examples/basic/events"
	"github.com/kolosys/discord/examples/basic/routes"
)

func main() {
	b, err := bot.New()
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	commands.Register(b)
	events.Register(b)
	routes.Register(b)

	if err := b.Start(context.Background()); err != nil {
		log.Fatalf("Bot error: %v", err)
	}
}

```

## Running the Example

To run this example:

```bash
cd basic
go run main.go
```

## Expected Output

```
Hello from Proton examples!
```

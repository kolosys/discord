# main

This example demonstrates basic usage of the library.

## Source Code

```go
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

	user, err := client.GetCurrentUser(ctx)
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}
	fmt.Printf("Current user: %s#%s (ID: %s)\n", user.Username, user.Discriminator, user.ID)
}

```

## Running the Example

To run this example:

```bash
cd rest-only
go run main.go
```

## Expected Output

```
Hello from Proton examples!
```

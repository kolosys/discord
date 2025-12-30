# routes

This example demonstrates basic usage of the library.

## Source Code

```go
package routes

import (
	"context"

	"github.com/kolosys/discord"
	"github.com/kolosys/helix"
)

func Register(bot *discord.Bot) {
	bot.GET("/health", helix.Handle(handleHealth))
	bot.GET("/api/status", helix.Handle(handleStatus(bot)))
}

type HealthResponse struct {
	Status string `json:"status"`
}

func handleHealth(ctx context.Context, req struct{}) (HealthResponse, error) {
	return HealthResponse{Status: "ok"}, nil
}

type StatusResponse struct {
	Gateway GatewayStatus `json:"gateway"`
}

type GatewayStatus struct {
	Connected bool `json:"connected"`
	Shards    int  `json:"shards"`
}

func handleStatus(bot *discord.Bot) func(context.Context, struct{}) (StatusResponse, error) {
	return func(ctx context.Context, req struct{}) (StatusResponse, error) {
		return StatusResponse{
			Gateway: GatewayStatus{
				Connected: bot.Gateway.IsRunning(),
				Shards:    bot.Gateway.ShardCount(),
			},
		}, nil
	}
}

```

## Running the Example

To run this example:

```bash
cd routes
go run routes.go
```

## Expected Output

```
Hello from Proton examples!
```

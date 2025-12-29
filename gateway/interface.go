package gateway

import (
	"context"

	"github.com/kolosys/axon"
)

// GatewayClient defines the interface for WebSocket client operations required by Shard.
// It enables dependency injection for testing with axon/mock while maintaining type safety.
type GatewayClient[T any] interface {
	// Core connection operations
	ConnectWithReadLoop(ctx context.Context) error
	Write(ctx context.Context, msg T) error
	Close() error

	// Callback registration
	OnMessage(fn func(T))
	OnConnect(fn func(*axon.Client[T]))
	OnDisconnect(fn func(*axon.Client[T], error))

	// State inspection
	State() axon.ConnectionState
	IsConnected() bool
}

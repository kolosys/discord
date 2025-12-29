package testutil

import (
	"github.com/kolosys/axon"
	"github.com/kolosys/axon/mock"
	"github.com/kolosys/discord/gateway"
)

// MockShardClient adapts axon/mock.MockClient for Shard testing.
// It implements gateway.GatewayClient[gateway.GatewayPayload] by adapting
// callback signatures to match those expected by the Shard.
type MockShardClient struct {
	*mock.MockClient[gateway.GatewayPayload]
}

// NewMockShardClient creates a mock client for shard testing.
func NewMockShardClient(opts *mock.MockClientOptions) *MockShardClient {
	if opts == nil {
		opts = &mock.MockClientOptions{
			BufferSize:         100,
			InitialState:       axon.StateDisconnected,
			RecordMessages:     true,
			RecordStateChanges: true,
		}
	}

	return &MockShardClient{
		MockClient: mock.NewMockClient[gateway.GatewayPayload](opts),
	}
}

// OnConnect adapts the callback signature to match Shard's expectations.
// The Shard passes *axon.Client[T] to the callback, but the mock only
// supports a no-arg callback. We pass nil for the client parameter.
func (msc *MockShardClient) OnConnect(fn func(*axon.Client[gateway.GatewayPayload])) {
	msc.MockClient.OnConnect(func() {
		if fn != nil {
			fn(nil)
		}
	})
}

// OnDisconnect adapts the callback signature to match Shard's expectations.
// The Shard passes *axon.Client[T] and error to the callback.
func (msc *MockShardClient) OnDisconnect(fn func(*axon.Client[gateway.GatewayPayload], error)) {
	msc.MockClient.OnDisconnect(func(err error) {
		if fn != nil {
			fn(nil, err)
		}
	})
}

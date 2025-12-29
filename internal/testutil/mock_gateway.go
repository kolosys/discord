package testutil

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/kolosys/axon"
	"github.com/kolosys/discord/gateway"
)

// MockGateway simulates a Discord gateway WebSocket server for testing.
type MockGateway struct {
	listener  net.Listener
	addr      string
	server    *http.Server
	mu        sync.RWMutex
	conns     map[*axon.Conn[gateway.GatewayPayload]]bool
	closeOnce sync.Once
}

// NewMockGateway creates a new mock gateway server listening on a random port.
func NewMockGateway() (*MockGateway, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("listen: %w", err)
	}

	mg := &MockGateway{
		listener: listener,
		addr:     listener.Addr().String(),
		conns:    make(map[*axon.Conn[gateway.GatewayPayload]]bool),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", mg.handleWebSocket)

	mg.server = &http.Server{
		Addr:    mg.addr,
		Handler: mux,
	}

	go func() {
		if err := mg.server.Serve(listener); err != nil && err != http.ErrServerClosed {
			fmt.Printf("mock gateway error: %v\n", err)
		}
	}()

	return mg, nil
}

// URL returns the WebSocket URL for connecting to this gateway.
func (mg *MockGateway) URL() string {
	return fmt.Sprintf("ws://%s", mg.addr)
}

// Close shuts down the mock gateway.
func (mg *MockGateway) Close() error {
	var err error
	mg.closeOnce.Do(func() {
		mg.mu.Lock()
		defer mg.mu.Unlock()

		for conn := range mg.conns {
			conn.Close(1000, "server closing")
		}

		err = mg.server.Close()
	})
	return err
}

// SendToAll sends a payload to all connected clients.
func (mg *MockGateway) SendToAll(ctx context.Context, payload *gateway.GatewayPayload) error {
	mg.mu.RLock()
	defer mg.mu.RUnlock()

	for conn := range mg.conns {
		if err := conn.Write(ctx, *payload); err != nil {
			return fmt.Errorf("send to connection: %w", err)
		}
	}
	return nil
}

// ConnectionCount returns the number of active connections.
func (mg *MockGateway) ConnectionCount() int {
	mg.mu.RLock()
	defer mg.mu.RUnlock()
	return len(mg.conns)
}

// GetConnections returns a list of all active connections for manual testing.
func (mg *MockGateway) GetConnections() []*axon.Conn[gateway.GatewayPayload] {
	mg.mu.RLock()
	defer mg.mu.RUnlock()

	conns := make([]*axon.Conn[gateway.GatewayPayload], 0, len(mg.conns))
	for conn := range mg.conns {
		conns = append(conns, conn)
	}
	return conns
}

func (mg *MockGateway) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := axon.Upgrade[gateway.GatewayPayload](w, r, nil)
	if err != nil {
		fmt.Printf("upgrade error: %v\n", err)
		return
	}

	mg.mu.Lock()
	mg.conns[conn] = true
	mg.mu.Unlock()

	// Read payloads from connection until it closes
	ctx := context.Background()
	for {
		payload, err := conn.Read(ctx)
		if err != nil {
			break
		}
		_ = payload // Payload received but not processed in base mock
	}

	mg.mu.Lock()
	delete(mg.conns, conn)
	mg.mu.Unlock()
}

// GatewayBuilder provides a fluent API for setting up a mock gateway.
type GatewayBuilder struct {
	mg   *MockGateway
	errs []error
}

// NewGatewayBuilder creates a new gateway builder.
func NewGatewayBuilder() (*GatewayBuilder, error) {
	mg, err := NewMockGateway()
	if err != nil {
		return nil, err
	}
	return &GatewayBuilder{mg: mg}, nil
}

// Build returns the configured mock gateway.
func (b *GatewayBuilder) Build() (*MockGateway, error) {
	if len(b.errs) > 0 {
		return nil, b.errs[0]
	}
	return b.mg, nil
}

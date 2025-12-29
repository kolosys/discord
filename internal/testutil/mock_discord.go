package testutil

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
)

// MockDiscord provides a mock HTTP server for Discord REST API testing.
type MockDiscord struct {
	server      *http.Server
	listener    net.Listener
	addr        string
	mu          sync.RWMutex
	handlers    map[string]Handler
	requests    []*ReceivedRequest
	requestsMu  sync.Mutex
	closed      atomic.Bool
}

// Handler processes a mock HTTP request.
type Handler func(w http.ResponseWriter, r *http.Request)

// ReceivedRequest records details about a request received by the mock.
type ReceivedRequest struct {
	Method string
	Path   string
	Header http.Header
	Body   []byte
}

// NewMockDiscord creates a new mock Discord API server.
func NewMockDiscord() (*MockDiscord, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("listen: %w", err)
	}

	md := &MockDiscord{
		listener: listener,
		addr:     listener.Addr().String(),
		handlers: make(map[string]Handler),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", md.handleRequest)

	md.server = &http.Server{
		Addr:    md.addr,
		Handler: mux,
	}

	go func() {
		if err := md.server.Serve(listener); err != nil && err != http.ErrServerClosed {
			fmt.Printf("mock discord error: %v\n", err)
		}
	}()

	return md, nil
}

// URL returns the base URL for the mock Discord API.
func (md *MockDiscord) URL() string {
	return fmt.Sprintf("http://%s", md.addr)
}

// Close shuts down the mock server.
func (md *MockDiscord) Close() error {
	md.closed.Store(true)
	return md.server.Close()
}

// Handle registers a handler for a specific method+path combination.
func (md *MockDiscord) Handle(method, path string, handler Handler) {
	md.mu.Lock()
	defer md.mu.Unlock()
	md.handlers[methodPath(method, path)] = handler
}

// HandleJSON registers a handler that responds with JSON.
func (md *MockDiscord) HandleJSON(method, path string, status int, response any) {
	data, _ := json.Marshal(response)
	md.Handle(method, path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(data)
	})
}

// HandleError registers a handler that responds with an error.
func (md *MockDiscord) HandleError(method, path string, status int, message string) {
	md.HandleJSON(method, path, status, map[string]any{
		"message": message,
		"code":    status,
	})
}

// ReceivedRequests returns all requests received by the mock.
func (md *MockDiscord) ReceivedRequests() []*ReceivedRequest {
	md.requestsMu.Lock()
	defer md.requestsMu.Unlock()

	reqs := make([]*ReceivedRequest, len(md.requests))
	copy(reqs, md.requests)
	return reqs
}

// LastRequest returns the most recently received request, or nil if none.
func (md *MockDiscord) LastRequest() *ReceivedRequest {
	md.requestsMu.Lock()
	defer md.requestsMu.Unlock()

	if len(md.requests) == 0 {
		return nil
	}
	return md.requests[len(md.requests)-1]
}

// RequestCount returns the number of requests received.
func (md *MockDiscord) RequestCount() int {
	md.requestsMu.Lock()
	defer md.requestsMu.Unlock()
	return len(md.requests)
}

// ClearRequests resets the request history.
func (md *MockDiscord) ClearRequests() {
	md.requestsMu.Lock()
	defer md.requestsMu.Unlock()
	md.requests = nil
}

func (md *MockDiscord) handleRequest(w http.ResponseWriter, r *http.Request) {
	// Record the request
	body := make([]byte, 0)
	if r.Body != nil {
		defer r.Body.Close()
		body = make([]byte, r.ContentLength)
		if r.ContentLength > 0 {
			r.Body.Read(body)
		}
	}

	md.requestsMu.Lock()
	md.requests = append(md.requests, &ReceivedRequest{
		Method: r.Method,
		Path:   r.URL.Path,
		Header: r.Header,
		Body:   body,
	})
	md.requestsMu.Unlock()

	// Check for registered handler
	md.mu.RLock()
	handler := md.handlers[methodPath(r.Method, r.URL.Path)]
	md.mu.RUnlock()

	if handler != nil {
		handler(w, r)
		return
	}

	// Default 404
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Not Found",
	})
}

// DiscordBuilder provides a fluent API for setting up a mock Discord server.
type DiscordBuilder struct {
	md    *MockDiscord
	errs  []error
}

// NewDiscordBuilder creates a new Discord builder.
func NewDiscordBuilder() (*DiscordBuilder, error) {
	md, err := NewMockDiscord()
	if err != nil {
		return nil, err
	}
	return &DiscordBuilder{md: md}, nil
}

// Handle registers a custom handler.
func (b *DiscordBuilder) Handle(method, path string, handler Handler) *DiscordBuilder {
	b.md.Handle(method, path, handler)
	return b
}

// HandleJSON registers a JSON response.
func (b *DiscordBuilder) HandleJSON(method, path string, status int, response any) *DiscordBuilder {
	b.md.HandleJSON(method, path, status, response)
	return b
}

// HandleError registers an error response.
func (b *DiscordBuilder) HandleError(method, path string, status int, message string) *DiscordBuilder {
	b.md.HandleError(method, path, status, message)
	return b
}

// Build returns the configured mock Discord server.
func (b *DiscordBuilder) Build() (*MockDiscord, error) {
	if len(b.errs) > 0 {
		return nil, b.errs[0]
	}
	return b.md, nil
}

// Helper to create method+path key
func methodPath(method, path string) string {
	return method + " " + path
}

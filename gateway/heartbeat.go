package gateway

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// HeartbeatManager manages the heartbeat loop for a Discord gateway connection.
type HeartbeatManager struct {
	interval      time.Duration
	lastHeartbeat atomic.Int64 // Unix timestamp in milliseconds
	lastACK       atomic.Int64 // Unix timestamp in milliseconds
	ackReceived   atomic.Bool
	session       *Session

	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.RWMutex
	wg     sync.WaitGroup
}

// NewHeartbeatManager creates a new heartbeat manager.
func NewHeartbeatManager(session *Session) *HeartbeatManager {
	return &HeartbeatManager{
		session: session,
	}
}

// Start begins the heartbeat loop.
// interval is in milliseconds.
// sendFunc is called to send the heartbeat payload.
func (h *HeartbeatManager) Start(interval time.Duration, sendFunc func(seq int) error) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	// Cancel existing heartbeat if running
	if h.cancel != nil {
		h.cancel()
		h.wg.Wait()
	}

	h.interval = interval
	h.ackReceived.Store(true) // Initialize as acknowledged
	now := time.Now().UnixMilli()
	h.lastACK.Store(now)

	h.ctx, h.cancel = context.WithCancel(context.Background())

	h.wg.Add(1)
	go h.heartbeatLoop(sendFunc)

	return nil
}

// Stop stops the heartbeat loop.
func (h *HeartbeatManager) Stop() {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.cancel != nil {
		h.cancel()
		h.wg.Wait()
		h.cancel = nil
	}
}

// ACK marks a heartbeat as acknowledged.
func (h *HeartbeatManager) ACK() {
	now := time.Now().UnixMilli()
	h.lastACK.Store(now)
	h.ackReceived.Store(true)
}

// IsHealthy checks if the last heartbeat was acknowledged.
// Returns false if a heartbeat was sent but not acknowledged, indicating a zombie connection.
func (h *HeartbeatManager) IsHealthy() bool {
	return h.ackReceived.Load()
}

// LastHeartbeat returns the timestamp of the last sent heartbeat.
func (h *HeartbeatManager) LastHeartbeat() time.Time {
	ms := h.lastHeartbeat.Load()
	if ms == 0 {
		return time.Time{}
	}
	return time.UnixMilli(ms)
}

// LastACK returns the timestamp of the last received ACK.
func (h *HeartbeatManager) LastACK() time.Time {
	ms := h.lastACK.Load()
	if ms == 0 {
		return time.Time{}
	}
	return time.UnixMilli(ms)
}

// heartbeatLoop runs the heartbeat sending loop.
func (h *HeartbeatManager) heartbeatLoop(sendFunc func(seq int) error) {
	defer h.wg.Done()

	// Send first heartbeat after interval * jitter (per Discord docs)
	jitter := rand.Float64()
	firstDelay := time.Duration(float64(h.interval) * jitter)

	select {
	case <-h.ctx.Done():
		return
	case <-time.After(firstDelay):
		h.sendHeartbeat(sendFunc)
	}

	ticker := time.NewTicker(h.interval)
	defer ticker.Stop()

	for {
		select {
		case <-h.ctx.Done():
			return
		case <-ticker.C:
			h.sendHeartbeat(sendFunc)
		}
	}
}

// sendHeartbeat sends a single heartbeat and updates state.
func (h *HeartbeatManager) sendHeartbeat(sendFunc func(seq int) error) {
	// Check if last heartbeat was acknowledged
	if !h.ackReceived.Load() {
		// Zombie connection detected - should trigger reconnect
		// The shard will handle this by checking IsHealthy()
		return
	}

	// Mark as waiting for ACK
	h.ackReceived.Store(false)

	// Get current sequence
	seq := h.session.Sequence()

	// Send heartbeat
	if err := sendFunc(seq); err != nil {
		// Error sending heartbeat - log but continue
		// The connection might be dead, but we'll let the shard handle it
		return
	}

	// Record heartbeat timestamp
	now := time.Now().UnixMilli()
	h.lastHeartbeat.Store(now)
}

// String returns a string representation of the heartbeat manager state.
func (h *HeartbeatManager) String() string {
	lastHB := h.LastHeartbeat()
	lastACK := h.LastACK()
	healthy := h.IsHealthy()

	return fmt.Sprintf(
		"HeartbeatManager{interval=%v, lastHeartbeat=%v, lastACK=%v, healthy=%v}",
		h.interval,
		lastHB,
		lastACK,
		healthy,
	)
}

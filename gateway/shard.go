package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/kolosys/axon"
	"github.com/kolosys/nova/bus"
	"github.com/kolosys/nova/shared"
)

// Shard represents a single Discord gateway shard connection.
type Shard struct {
	id      int
	total   int
	client  *axon.Client[GatewayPayload]
	session *Session
	hb      *HeartbeatManager
	events  bus.EventBus

	token   string
	intents Intent

	mu         sync.RWMutex
	ready      bool
	fatalError error // Non-nil if a fatal error occurred
}

// NewShard creates a new shard.
func NewShard(id, total int, token string, intents Intent, events bus.EventBus) *Shard {
	session := NewSession()
	hb := NewHeartbeatManager(session)

	return &Shard{
		id:      id,
		total:   total,
		session: session,
		hb:      hb,
		events:  events,
		token:   token,
		intents: intents,
	}
}

// Connect establishes a WebSocket connection to the gateway.
func (s *Shard) Connect(ctx context.Context, url string) error {
	// Add gateway version and encoding parameters
	gatewayURL := fmt.Sprintf("%s?v=10&encoding=json", url)

	// Create axon client options
	opts := axon.DefaultClientOptions()
	opts.Reconnect.Enabled = true
	opts.Reconnect.MaxAttempts = 0 // Unlimited reconnection attempts
	opts.Reconnect.InitialDelay = 1 * time.Second
	opts.Reconnect.MaxDelay = 120 * time.Second
	opts.Reconnect.OnReconnecting = func(attempt int, delay time.Duration) {
		log.Printf("[Shard %d] Reconnecting (attempt %d, delay %s)...", s.id, attempt, delay)
	}
	opts.Reconnect.OnReconnected = func(attempt int) {
		log.Printf("[Shard %d] Reconnected after %d attempt(s)", s.id, attempt)
	}
	opts.Reconnect.OnReconnectFailed = func(attempt int, err error) {
		log.Printf("[Shard %d] Reconnect attempt %d failed: %v", s.id, attempt, err)
	}
	// Custom reconnection logic for Discord gateway
	// Discord close code 1000 should trigger reconnection (unlike standard WebSocket behavior)
	opts.Reconnect.ShouldReconnect = func(err error, attempt int) bool {
		closeErr := axon.AsCloseError(err)
		if closeErr == nil {
			return true // Network errors should always reconnect
		}
		// Check Discord-specific close codes
		code := CloseCode(closeErr.Code)
		if code.IsFatal() {
			return false // Don't reconnect for fatal Discord errors
		}
		return true // Reconnect for everything else including 1000
	}

	// Discord heartbeat interval is ~41.25 seconds, set read deadline longer to avoid timeout
	opts.ReadDeadline = 90 * time.Second

	// Disable WebSocket-level pings - Discord uses its own heartbeat mechanism (opcode 1)
	opts.PingInterval = 0

	// Create WebSocket client
	client := axon.NewClient[GatewayPayload](gatewayURL, opts)

	// Set up callbacks
	client.OnMessage(s.handlePayload)
	client.OnConnect(func(c *axon.Client[GatewayPayload]) {
		log.Printf("[Shard %d] Connected to gateway", s.id)
	})
	client.OnDisconnect(func(c *axon.Client[GatewayPayload], err error) {
		s.handleDisconnect(c, err)
	})

	s.mu.Lock()
	s.client = client
	s.mu.Unlock()

	// Connect with read loop
	if err := client.ConnectWithReadLoop(ctx); err != nil {
		return fmt.Errorf("failed to connect to gateway: %w", err)
	}

	return nil
}

// handlePayload processes incoming gateway payloads.
func (s *Shard) handlePayload(payload GatewayPayload) {
	switch payload.Op {
	case OpcodeHello:
		s.handleHello(payload)
	case OpcodeHeartbeat:
		// Server requested an immediate heartbeat - respond without waiting
		s.sendHeartbeat(s.session.Sequence())
	case OpcodeHeartbeatACK:
		s.hb.ACK()
	case OpcodeDispatch:
		s.handleDispatch(payload)
	case OpcodeReconnect:
		log.Printf("[Shard %d] Received RECONNECT, reconnecting...", s.id)
		// Axon will handle reconnection automatically
	case OpcodeInvalidSession:
		s.handleInvalidSession(payload)
	default:
		log.Printf("[Shard %d] Received opcode %s", s.id, payload.Op.String())
	}
}

// handleHello processes the Hello opcode and starts the heartbeat.
func (s *Shard) handleHello(payload GatewayPayload) {
	var hello HelloData
	if err := json.Unmarshal(payload.D, &hello); err != nil {
		log.Printf("[Shard %d] Failed to unmarshal Hello: %v", s.id, err)
		return
	}

	log.Printf("[Shard %d] Received Hello, heartbeat interval: %dms", s.id, hello.HeartbeatInterval)

	// Start heartbeat
	interval := time.Duration(hello.HeartbeatInterval) * time.Millisecond
	err := s.hb.Start(interval, s.sendHeartbeat)
	if err != nil {
		log.Printf("[Shard %d] Failed to start heartbeat: %v", s.id, err)
		return
	}

	// Identify or Resume
	if s.session.CanResume() {
		log.Printf("[Shard %d] Resuming session %s", s.id, s.session.SessionID())
		s.sendResume()
	} else {
		s.logIntentInfo()
		s.sendIdentify()
	}
}

// handleInvalidSession processes the Invalid Session opcode.
func (s *Shard) handleInvalidSession(payload GatewayPayload) {
	var resumable InvalidSessionData
	if err := json.Unmarshal(payload.D, &resumable); err != nil {
		log.Printf("[Shard %d] Failed to unmarshal Invalid Session: %v", s.id, err)
		return
	}

	if resumable {
		log.Printf("[Shard %d] Session invalidated but resumable, resuming...", s.id)
		// Wait a bit before resuming as recommended by Discord
		time.Sleep(time.Second * 5)
		s.sendResume()
	} else {
		log.Printf("[Shard %d] Session invalidated and not resumable, re-identifying...", s.id)
		s.session.Reset()
		// Wait a bit before re-identifying
		time.Sleep(time.Second * 5)
		s.sendIdentify()
	}
}

// handleDispatch processes dispatch events.
func (s *Shard) handleDispatch(payload GatewayPayload) {
	// Update sequence
	if payload.S != nil {
		s.session.SetSequence(*payload.S)
	}

	// Get event name
	if payload.T == nil {
		log.Printf("[Shard %d] Dispatch event missing event name", s.id)
		return
	}
	eventName := *payload.T

	// Handle READY specially
	if eventName == "READY" {
		s.handleReady(payload)
		return
	}

	// Publish event to event bus
	event := shared.NewBaseEvent(
		fmt.Sprintf("shard-%d-%s-%d", s.id, eventName, time.Now().UnixNano()),
		eventName,
		payload.D,
	)
	event.SetMetadata("shard_id", fmt.Sprintf("%d", s.id))

	if eventName == "INTERACTION_CREATE" {
		log.Printf("[Shard %d] Publishing INTERACTION_CREATE event", s.id)
	}

	ctx := context.Background()
	if err := s.events.Publish(ctx, eventName, event); err != nil {
		log.Printf("[Shard %d] Failed to publish event %s: %v", s.id, eventName, err)
	}
}

// handleReady processes the Ready event.
func (s *Shard) handleReady(payload GatewayPayload) {
	var ready ReadyData
	if err := json.Unmarshal(payload.D, &ready); err != nil {
		log.Printf("[Shard %d] Failed to unmarshal Ready: %v", s.id, err)
		return
	}

	log.Printf("[Shard %d] Ready! Session ID: %s, User: %s#%s",
		s.id,
		ready.SessionID,
		ready.User.Username,
		ready.User.Discriminator,
	)

	// Store session info
	s.session.SetSessionID(ready.SessionID)
	s.session.SetResumeURL(ready.ResumeGatewayURL)
	s.session.MarkIdentified()

	// Mark as ready
	s.mu.Lock()
	s.ready = true
	s.mu.Unlock()

	// Publish READY event
	event := shared.NewBaseEvent(
		fmt.Sprintf("shard-%d-READY-%d", s.id, time.Now().UnixNano()),
		"READY",
		payload.D,
	)
	event.SetMetadata("shard_id", fmt.Sprintf("%d", s.id))

	ctx := context.Background()
	if err := s.events.Publish(ctx, "READY", event); err != nil {
		log.Printf("[Shard %d] Failed to publish READY event: %v", s.id, err)
	}
}

// sendIdentify sends an Identify payload.
func (s *Shard) sendIdentify() error {
	var shard *[2]int
	if s.total > 1 {
		shard = &[2]int{s.id, s.total}
	}

	payload, err := MarshalIdentify(s.token, s.intents, shard, nil)
	if err != nil {
		return fmt.Errorf("failed to marshal identify: %w", err)
	}

	ctx := context.Background()
	if err := s.client.Write(ctx, *payload); err != nil {
		return fmt.Errorf("failed to send identify: %w", err)
	}

	return nil
}

// sendResume sends a Resume payload.
func (s *Shard) sendResume() error {
	payload, err := MarshalResume(
		s.token,
		s.session.SessionID(),
		s.session.Sequence(),
	)
	if err != nil {
		return fmt.Errorf("failed to marshal resume: %w", err)
	}

	ctx := context.Background()
	if err := s.client.Write(ctx, *payload); err != nil {
		return fmt.Errorf("failed to send resume: %w", err)
	}

	return nil
}

// sendHeartbeat sends a Heartbeat payload.
func (s *Shard) sendHeartbeat(seq int) error {
	seqPtr := &seq
	payload, err := MarshalHeartbeat(seqPtr)
	if err != nil {
		return fmt.Errorf("failed to marshal heartbeat: %w", err)
	}

	ctx := context.Background()
	if err := s.client.Write(ctx, *payload); err != nil {
		return fmt.Errorf("failed to send heartbeat: %w", err)
	}

	return nil
}

// IsReady returns whether the shard has received the READY event.
func (s *Shard) IsReady() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.ready
}

// handleDisconnect processes gateway disconnections with meaningful error context.
func (s *Shard) handleDisconnect(c *axon.Client[GatewayPayload], err error) {
	s.mu.Lock()
	s.ready = false
	s.mu.Unlock()

	// Check if error is an axon CloseError (contains close code and reason)
	closeErr := axon.AsCloseError(err)
	if closeErr == nil {
		// No close code means network error or unexpected disconnect
		log.Printf("[Shard %d] Disconnected from gateway: %v", s.id, err)
		return
	}

	// Map axon close code to Discord close code for detailed descriptions
	code := CloseCode(closeErr.Code)

	// Log with appropriate severity based on error type
	if code.IsFatal() {
		log.Printf("[Shard %d] FATAL: Disconnected from gateway (code: %d - %s)", s.id, closeErr.Code, code.String())
		if closeErr.Reason != "" {
			log.Printf("[Shard %d] Reason: %s", s.id, closeErr.Reason)
		}
		log.Printf("[Shard %d] %s", s.id, code.Description())
		log.Printf("[Shard %d] Reconnection disabled due to non-recoverable error", s.id)

		// Store fatal error and close the client to prevent reconnection
		s.mu.Lock()
		s.fatalError = fmt.Errorf("gateway close code %d: %s", closeErr.Code, code.String())
		s.mu.Unlock()

		// Close the client to stop reconnection attempts
		go func() {
			if err := c.Close(); err != nil {
				log.Printf("[Shard %d] Error closing client after fatal error: %v", s.id, err)
			}
		}()
	} else {
		log.Printf("[Shard %d] Disconnected from gateway (code: %d - %s)", s.id, closeErr.Code, code.String())
		if closeErr.Reason != "" {
			log.Printf("[Shard %d] Reason: %s", s.id, closeErr.Reason)
		}
		log.Printf("[Shard %d] %s", s.id, code.Description())
		log.Printf("[Shard %d] Reconnection will be attempted...", s.id)
	}
}

// FatalError returns the fatal error if one occurred, or nil.
func (s *Shard) FatalError() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.fatalError
}

// UpdatePresence updates the bot's presence/status.
func (s *Shard) UpdatePresence(ctx context.Context, presence *PresenceUpdate) error {
	payload, err := MarshalPresenceUpdate(presence)
	if err != nil {
		return fmt.Errorf("failed to marshal presence update: %w", err)
	}

	s.mu.RLock()
	client := s.client
	ready := s.ready
	s.mu.RUnlock()

	if client == nil {
		return fmt.Errorf("shard not connected")
	}

	if !ready {
		return fmt.Errorf("shard not ready")
	}

	if err := client.Write(ctx, *payload); err != nil {
		return fmt.Errorf("failed to send presence update: %w", err)
	}

	return nil
}

// UpdateVoiceState updates the bot's voice state (join/leave/move voice channels).
func (s *Shard) UpdateVoiceState(ctx context.Context, data *VoiceStateUpdateData) error {
	payload, err := MarshalVoiceStateUpdate(data)
	if err != nil {
		return fmt.Errorf("failed to marshal voice state update: %w", err)
	}

	s.mu.RLock()
	client := s.client
	s.mu.RUnlock()

	if client == nil {
		return fmt.Errorf("shard not connected")
	}

	if err := client.Write(ctx, *payload); err != nil {
		return fmt.Errorf("failed to send voice state update: %w", err)
	}

	return nil
}

// RequestGuildMembers requests guild members from Discord.
// Results are received via GUILD_MEMBERS_CHUNK events.
func (s *Shard) RequestGuildMembers(ctx context.Context, data *RequestGuildMembersData) error {
	payload, err := MarshalRequestGuildMembers(data)
	if err != nil {
		return fmt.Errorf("failed to marshal request guild members: %w", err)
	}

	s.mu.RLock()
	client := s.client
	s.mu.RUnlock()

	if client == nil {
		return fmt.Errorf("shard not connected")
	}

	if err := client.Write(ctx, *payload); err != nil {
		return fmt.Errorf("failed to send request guild members: %w", err)
	}

	return nil
}

// logIntentInfo logs information about configured intents and any warnings.
func (s *Shard) logIntentInfo() {
	log.Printf("[Shard %d] Identifying with intents: %s", s.id, s.intents.String())

	// Log privileged intents if any are requested
	if privileged := s.intents.PrivilegedIntents(); len(privileged) > 0 {
		log.Printf("[Shard %d] Privileged intents requested: %v", s.id, privileged)
		log.Printf("[Shard %d] Ensure these are enabled in Discord Developer Portal", s.id)
	}

	// Log any configuration warnings
	for _, warning := range s.intents.Warnings() {
		log.Printf("[Shard %d] WARNING: %s", s.id, warning)
	}
}

// Close gracefully closes the shard connection.
func (s *Shard) Close() error {
	log.Printf("[Shard %d] Closing...", s.id)

	// Stop heartbeat
	s.hb.Stop()

	// Close WebSocket client
	s.mu.RLock()
	client := s.client
	s.mu.RUnlock()

	if client != nil {
		return client.Close()
	}

	return nil
}

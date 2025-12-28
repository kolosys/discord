package events

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/kolosys/nova/bus"
	"github.com/kolosys/nova/shared"
)

// Handler is a type-safe event handler function.
type Handler[T any] func(ctx context.Context, event *T)

// RawHandler handles raw event data without type conversion.
type RawHandler func(ctx context.Context, eventType Type, data json.RawMessage)

// Dispatcher manages event subscriptions and dispatching.
type Dispatcher struct {
	bus           bus.EventBus
	subscriptions []shared.Subscription
	mu            sync.Mutex
}

// NewDispatcher creates a new event dispatcher.
func NewDispatcher(b bus.EventBus) *Dispatcher {
	return &Dispatcher{
		bus:           b,
		subscriptions: make([]shared.Subscription, 0),
	}
}

// listenerID is a counter for generating unique listener IDs.
var listenerID atomic.Uint64

// eventListener implements shared.Listener for event handling.
type eventListener struct {
	id      string
	handler func(ctx context.Context, event shared.Event) error
}

func newEventListener(handler func(ctx context.Context, event shared.Event) error) *eventListener {
	id := listenerID.Add(1)
	return &eventListener{
		id:      fmt.Sprintf("discord-listener-%d", id),
		handler: handler,
	}
}

func (l *eventListener) ID() string { return l.id }
func (l *eventListener) Handle(ctx context.Context, event shared.Event) error {
	return l.handler(ctx, event)
}
func (l *eventListener) OnError(_ context.Context, _ shared.Event, _ error) error { return nil }

// On registers a type-safe handler for a specific event type.
// The handler will receive the event data already deserialized into the correct type.
// Context is propagated from the event bus for cancellation and deadline support.
func On[T any](d *Dispatcher, eventType Type, handler Handler[T]) error {
	listener := newEventListener(func(ctx context.Context, e shared.Event) error {
		raw, ok := e.Data().(json.RawMessage)
		if !ok {
			return fmt.Errorf("invalid event data type: expected json.RawMessage, got %T", e.Data())
		}

		var event T
		if err := json.Unmarshal(raw, &event); err != nil {
			return fmt.Errorf("failed to unmarshal %s event: %w", eventType, err)
		}

		if shardID := e.Metadata()["shard_id"]; shardID != "" {
			setShardID(&event, shardID)
		}

		handler(ctx, &event)
		return nil
	})

	sub := d.bus.Subscribe(string(eventType), listener)

	d.mu.Lock()
	d.subscriptions = append(d.subscriptions, sub)
	d.mu.Unlock()

	return nil
}

// OnRaw registers a handler that receives raw event data.
// Useful for custom event processing or debugging.
// Context is propagated from the event bus for cancellation and deadline support.
func (d *Dispatcher) OnRaw(eventType Type, handler RawHandler) error {
	listener := newEventListener(func(ctx context.Context, e shared.Event) error {
		raw, ok := e.Data().(json.RawMessage)
		if !ok {
			return fmt.Errorf("invalid event data type: expected json.RawMessage, got %T", e.Data())
		}
		handler(ctx, eventType, raw)
		return nil
	})

	sub := d.bus.Subscribe(string(eventType), listener)

	d.mu.Lock()
	d.subscriptions = append(d.subscriptions, sub)
	d.mu.Unlock()

	return nil
}

// OnAny registers a handler for all events using pattern matching.
// Context is propagated from the event bus for cancellation and deadline support.
func (d *Dispatcher) OnAny(handler RawHandler) error {
	listener := newEventListener(func(ctx context.Context, e shared.Event) error {
		raw, ok := e.Data().(json.RawMessage)
		if !ok {
			return nil // Skip non-JSON events
		}
		eventType := Type(e.Type())
		handler(ctx, eventType, raw)
		return nil
	})

	sub := d.bus.SubscribePattern(".*", listener)

	d.mu.Lock()
	d.subscriptions = append(d.subscriptions, sub)
	d.mu.Unlock()

	return nil
}

// Close unsubscribes all registered handlers.
func (d *Dispatcher) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	for _, sub := range d.subscriptions {
		if sub.Active() {
			if err := sub.Unsubscribe(); err != nil {
				return fmt.Errorf("failed to unsubscribe: %w", err)
			}
		}
	}
	d.subscriptions = nil
	return nil
}

// shardAware is implemented by events that track which shard they came from.
type shardAware interface {
	setShardID(id int)
}

// setShardID sets the shard ID on events that implement shardAware.
func setShardID(event any, shardID string) {
	if sa, ok := event.(shardAware); ok {
		id, _ := strconv.Atoi(shardID)
		sa.setShardID(id)
	}
}

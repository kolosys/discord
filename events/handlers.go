package events

import (
	"context"
	"encoding/json"
	"fmt"
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
	handler func(shared.Event) error
}

func newEventListener(handler func(shared.Event) error) *eventListener {
	id := listenerID.Add(1)
	return &eventListener{
		id:      fmt.Sprintf("discord-listener-%d", id),
		handler: handler,
	}
}

func (l *eventListener) ID() string                                  { return l.id }
func (l *eventListener) Handle(event shared.Event) error             { return l.handler(event) }
func (l *eventListener) OnError(event shared.Event, err error) error { return nil }

// On registers a type-safe handler for a specific event type.
// The handler will receive the event data already deserialized into the correct type.
func On[T any](d *Dispatcher, eventType Type, handler Handler[T]) error {
	listener := newEventListener(func(e shared.Event) error {
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

		handler(context.Background(), &event)
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
func (d *Dispatcher) OnRaw(eventType Type, handler RawHandler) error {
	listener := newEventListener(func(e shared.Event) error {
		raw, ok := e.Data().(json.RawMessage)
		if !ok {
			return fmt.Errorf("invalid event data type: expected json.RawMessage, got %T", e.Data())
		}
		handler(context.Background(), eventType, raw)
		return nil
	})

	sub := d.bus.Subscribe(string(eventType), listener)

	d.mu.Lock()
	d.subscriptions = append(d.subscriptions, sub)
	d.mu.Unlock()

	return nil
}

// OnAny registers a handler for all events using pattern matching.
func (d *Dispatcher) OnAny(handler RawHandler) error {
	listener := newEventListener(func(e shared.Event) error {
		raw, ok := e.Data().(json.RawMessage)
		if !ok {
			return nil // Skip non-JSON events
		}
		eventType := Type(e.Type())
		handler(context.Background(), eventType, raw)
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
		var id int
		_, _ = fmt.Sscanf(shardID, "%d", &id)
		sa.setShardID(id)
	}
}

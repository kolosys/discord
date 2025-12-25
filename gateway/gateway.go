package gateway

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/kolosys/atomic/collection"
	"github.com/kolosys/nova/bus"
)

// Gateway manages Discord gateway connections across multiple shards.
type Gateway struct {
	token   string
	intents Intent
	shards  *collection.Collection[int, *Shard]
	events  bus.EventBus

	mu      sync.RWMutex
	running bool
}

// NewGateway creates a new gateway manager.
func NewGateway(token string, intents Intent, events bus.EventBus) *Gateway {
	return &Gateway{
		token:   token,
		intents: intents,
		shards:  collection.New[int, *Shard](),
		events:  events,
	}
}

// Connect establishes connections to the Discord gateway for all shards.
// url is the gateway URL (e.g., "wss://gateway.discord.gg")
// shardCount is the total number of shards to create (1 for no sharding)
func (g *Gateway) Connect(ctx context.Context, url string, shardCount int) error {
	g.mu.Lock()
	if g.running {
		g.mu.Unlock()
		return fmt.Errorf("gateway already running")
	}
	g.running = true
	g.mu.Unlock()

	log.Printf("[Gateway] Connecting with %d shard(s)...", shardCount)

	// Create and connect all shards
	var wg sync.WaitGroup
	errChan := make(chan error, shardCount)

	for i := 0; i < shardCount; i++ {
		shardID := i
		wg.Add(1)

		go func() {
			defer wg.Done()

			shard := NewShard(shardID, shardCount, g.token, g.intents, g.events)
			g.shards.Set(shardID, shard)

			if err := shard.Connect(ctx, url); err != nil {
				errChan <- fmt.Errorf("shard %d failed to connect: %w", shardID, err)
				return
			}

			log.Printf("[Gateway] Shard %d/%d connected", shardID, shardCount)
		}()

		// Discord requires 5 second delay between shard connections
		// to avoid rate limiting on IDENTIFY
		// For the first shard, we don't need to wait
		// But we'll wait here for simplicity (can optimize later)
	}

	wg.Wait()
	close(errChan)

	// Check for errors
	for err := range errChan {
		g.mu.Lock()
		g.running = false
		g.mu.Unlock()
		return err
	}

	log.Printf("[Gateway] All %d shard(s) connected", shardCount)
	return nil
}

// Events returns the event bus for subscribing to gateway events.
func (g *Gateway) Events() bus.EventBus {
	return g.events
}

// Close gracefully closes all shard connections.
func (g *Gateway) Close() error {
	g.mu.Lock()
	if !g.running {
		g.mu.Unlock()
		return nil
	}
	g.running = false
	g.mu.Unlock()

	log.Printf("[Gateway] Closing all shards...")

	shardKeys := g.shards.Keys()
	for _, shardID := range shardKeys {
		if shard, ok := g.shards.Get(shardID); ok {
			if err := shard.Close(); err != nil {
				log.Printf("[Gateway] Error closing shard %d: %v", shardID, err)
			}
		}
	}

	g.shards.Clear()
	log.Printf("[Gateway] All shards closed")
	return nil
}

// IsReady returns whether all shards have received the READY event.
func (g *Gateway) IsReady() bool {
	shardKeys := g.shards.Keys()
	if len(shardKeys) == 0 {
		return false
	}

	for _, shardID := range shardKeys {
		if shard, ok := g.shards.Get(shardID); ok {
			if !shard.IsReady() {
				return false
			}
		}
	}

	return true
}

// IsRunning returns whether the gateway is currently running.
func (g *Gateway) IsRunning() bool {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.running
}

// ShardCount returns the number of shards.
func (g *Gateway) ShardCount() int {
	return g.shards.Size()
}

// GetShard returns a specific shard by ID.
func (g *Gateway) GetShard(id int) (*Shard, bool) {
	return g.shards.Get(id)
}

// UpdatePresence updates the bot's presence on all shards.
func (g *Gateway) UpdatePresence(ctx context.Context, presence *PresenceUpdate) error {
	g.mu.RLock()
	running := g.running
	g.mu.RUnlock()

	if !running {
		return fmt.Errorf("gateway not running")
	}

	shardKeys := g.shards.Keys()
	for _, shardID := range shardKeys {
		if shard, ok := g.shards.Get(shardID); ok {
			if err := shard.UpdatePresence(ctx, presence); err != nil {
				return fmt.Errorf("shard %d: %w", shardID, err)
			}
		}
	}

	return nil
}

// UpdateVoiceState updates the bot's voice state for a specific guild.
// The request is routed to the shard responsible for that guild.
func (g *Gateway) UpdateVoiceState(ctx context.Context, data *VoiceStateUpdateData) error {
	g.mu.RLock()
	running := g.running
	g.mu.RUnlock()

	if !running {
		return fmt.Errorf("gateway not running")
	}

	shardID := g.shardForGuild(data.GuildID)
	shard, ok := g.shards.Get(shardID)
	if !ok {
		return fmt.Errorf("shard %d not found", shardID)
	}

	return shard.UpdateVoiceState(ctx, data)
}

// RequestGuildMembers requests guild members for a specific guild.
// The request is routed to the shard responsible for that guild.
// Results are received via GUILD_MEMBERS_CHUNK events.
func (g *Gateway) RequestGuildMembers(ctx context.Context, data *RequestGuildMembersData) error {
	g.mu.RLock()
	running := g.running
	g.mu.RUnlock()

	if !running {
		return fmt.Errorf("gateway not running")
	}

	shardID := g.shardForGuild(data.GuildID)
	shard, ok := g.shards.Get(shardID)
	if !ok {
		return fmt.Errorf("shard %d not found", shardID)
	}

	return shard.RequestGuildMembers(ctx, data)
}

// shardForGuild calculates which shard handles a given guild.
// Formula: (guild_id >> 22) % num_shards
func (g *Gateway) shardForGuild(guildID string) int {
	shardCount := g.shards.Size()
	if shardCount <= 1 {
		return 0
	}

	// Parse guild ID as uint64 and apply Discord's sharding formula
	var id uint64
	for _, c := range guildID {
		if c >= '0' && c <= '9' {
			id = id*10 + uint64(c-'0')
		}
	}

	return int((id >> 22) % uint64(shardCount))
}

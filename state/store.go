package state

import (
	"context"
	"time"

	"github.com/kolosys/synapse"
)

type cacheConfig struct {
	disabled bool
	maxSize  int
	ttl      time.Duration
}

// Store is a generic entity cache with standard CRUD operations.
type Store[K comparable, V any] struct {
	cache *synapse.Cache[K, V]
}

func newStore[K comparable, V any](cfg cacheConfig, shards int) *Store[K, V] {
	if cfg.disabled {
		return nil
	}
	opts := []synapse.Option{synapse.WithShards(shards)}
	if cfg.maxSize > 0 {
		opts = append(opts, synapse.WithMaxSize(cfg.maxSize))
	}
	if cfg.ttl > 0 {
		opts = append(opts, synapse.WithTTL(cfg.ttl))
	}
	return &Store[K, V]{cache: synapse.New[K, V](opts...)}
}

func (s *Store[K, V]) Get(ctx context.Context, key K) (V, bool) {
	if s == nil || s.cache == nil {
		var zero V
		return zero, false
	}
	return s.cache.Get(ctx, key)
}

func (s *Store[K, V]) Set(ctx context.Context, key K, value V) {
	if s == nil || s.cache == nil {
		return
	}
	s.cache.Set(ctx, key, value)
}

func (s *Store[K, V]) Delete(ctx context.Context, key K) bool {
	if s == nil || s.cache == nil {
		return false
	}
	_, existed := s.cache.Get(ctx, key)
	if existed {
		s.cache.Delete(ctx, key)
	}
	return existed
}

func (s *Store[K, V]) Count() int {
	if s == nil || s.cache == nil {
		return 0
	}
	return s.cache.Len()
}

func (s *Store[K, V]) Enabled() bool {
	return s != nil && s.cache != nil
}

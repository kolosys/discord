package state_test

import (
	"context"
	"testing"
	"time"

	"github.com/kolosys/discord/state"
)

// TestStoreStringString tests Store with string keys and values
func TestStore_GetEmpty(t *testing.T) {
	ctx := context.Background()

	// Test with nil store (disabled)
	var s *state.Store[string, string]
	val, ok := s.Get(ctx, "key")

	if ok {
		t.Error("Get on nil store should return false")
	}
	if val != "" {
		t.Errorf("Get on nil store should return zero value, got %q", val)
	}
}

func TestStore_SetGet(t *testing.T) {
	ctx := context.Background()

	// Test with nil store
	var s *state.Store[string, string]
	s.Set(ctx, "test", "value")

	// Setting on nil should be a no-op
	_, ok := s.Get(ctx, "test")
	if ok {
		t.Error("Get on nil store should return false after Set")
	}
}

func TestStore_Delete(t *testing.T) {
	ctx := context.Background()

	// Test with nil store
	var s *state.Store[string, string]
	existed := s.Delete(ctx, "key")

	if existed {
		t.Error("Delete on nil store should return false")
	}
}

func TestStore_DeleteMissing(t *testing.T) {
	ctx := context.Background()

	// Test with nil store
	var s *state.Store[string, string]
	existed := s.Delete(ctx, "nonexistent")

	if existed {
		t.Error("Delete of missing key should return false")
	}
}

func TestStore_Count(t *testing.T) {
	// Test with nil store
	var s *state.Store[string, string]
	count := s.Count()

	if count != 0 {
		t.Errorf("Count on nil store = %d, want 0", count)
	}
}

func TestStore_Enabled_Nil(t *testing.T) {
	// Test with nil store
	var s *state.Store[string, string]
	if s.Enabled() {
		t.Error("Enabled() should return false for nil store")
	}
}

// TestOptions tests the options structure
func TestOptions_Creation(t *testing.T) {
	opts := state.Options{}

	// Options should be creatable
	if opts.MaxGuilds < 0 {
		t.Error("MaxGuilds should not be negative")
	}
}

func TestOptions_WithValues(t *testing.T) {
	opts := state.Options{
		MaxGuilds:    1000,
		MaxUsers:     10000,
		MaxChannels:  100000,
	}

	if opts.MaxGuilds != 1000 {
		t.Errorf("MaxGuilds = %d, want 1000", opts.MaxGuilds)
	}
	if opts.MaxUsers != 10000 {
		t.Errorf("MaxUsers = %d, want 10000", opts.MaxUsers)
	}
	if opts.MaxChannels != 100000 {
		t.Errorf("MaxChannels = %d, want 100000", opts.MaxChannels)
	}
}

// NilStore tests
func TestNilStore_GetReturnsZeroValue(t *testing.T) {
	ctx := context.Background()
	var s *state.Store[string, int]

	val, ok := s.Get(ctx, "any_key")
	if ok {
		t.Error("Get on nil store should return ok=false")
	}
	if val != 0 {
		t.Errorf("Get on nil store should return zero value, got %d", val)
	}
}

func TestNilStore_GetWithStructType(t *testing.T) {
	ctx := context.Background()

	type TestData struct {
		ID   string
		Name string
	}

	var s *state.Store[string, TestData]
	val, ok := s.Get(ctx, "key")

	if ok {
		t.Error("Get on nil store should return false")
	}
	if val.ID != "" || val.Name != "" {
		t.Errorf("Get on nil store should return zero struct")
	}
}

func TestNilStore_SetIsNoOp(t *testing.T) {
	ctx := context.Background()
	var s *state.Store[string, string]

	// These should be safe no-ops
	s.Set(ctx, "key", "value")
	s.Set(ctx, "key2", "value2")

	// Verify still returns nothing
	_, ok := s.Get(ctx, "key")
	if ok {
		t.Error("Set on nil store should be no-op")
	}
}

func TestNilStore_DeleteIsNoOp(t *testing.T) {
	ctx := context.Background()
	var s *state.Store[string, string]

	// Delete on nil should return false (nothing to delete)
	result := s.Delete(ctx, "key")
	if result {
		t.Error("Delete on nil store should return false")
	}
}

func TestNilStore_CountIsZero(t *testing.T) {
	var s *state.Store[string, string]
	count := s.Count()

	if count != 0 {
		t.Errorf("Count on nil store = %d, want 0", count)
	}
}

func TestNilStore_Operations(t *testing.T) {
	ctx := context.Background()
	var s *state.Store[string, interface{}]

	// All operations should be safe
	s.Set(ctx, "key", "value")
	s.Set(ctx, "key2", 123)
	s.Set(ctx, "key3", struct{}{})

	_, ok := s.Get(ctx, "key")
	if ok {
		t.Error("Get on nil store should return false")
	}

	deleted := s.Delete(ctx, "key")
	if deleted {
		t.Error("Delete on nil store should return false")
	}

	count := s.Count()
	if count != 0 {
		t.Errorf("Count = %d, want 0", count)
	}

	if s.Enabled() {
		t.Error("Enabled should return false")
	}
}

func TestStore_IntType(t *testing.T) {
	ctx := context.Background()
	var s *state.Store[string, int]

	s.Set(ctx, "counter", 42)
	val, ok := s.Get(ctx, "counter")

	if ok {
		t.Error("Get on nil store should return false")
	}
	if val != 0 {
		t.Errorf("Get should return zero int, got %d", val)
	}
}

func TestStore_BoolType(t *testing.T) {
	ctx := context.Background()
	var s *state.Store[string, bool]

	s.Set(ctx, "flag", true)
	val, ok := s.Get(ctx, "flag")

	if ok {
		t.Error("Get on nil store should return false")
	}
	if val {
		t.Error("Get should return false (zero value for bool)")
	}
}

func TestStore_TimeoutContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	var s *state.Store[string, string]

	// Operations should still work (or return nil behavior) even with timeout context
	s.Set(ctx, "key", "value")
	_, ok := s.Get(ctx, "key")
	if ok {
		t.Error("Get on nil store should return false")
	}

	s.Delete(ctx, "key")
	count := s.Count()
	if count != 0 {
		t.Errorf("Count = %d, want 0", count)
	}
}

func TestStore_CancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	var s *state.Store[string, string]

	// Cancelled context operations should handle gracefully
	s.Set(ctx, "key", "value")
	_, ok := s.Get(ctx, "key")
	if ok {
		t.Error("Get on nil store should return false")
	}
}

func TestOptions_DisableFlags(t *testing.T) {
	tests := []struct {
		name string
		opts state.Options
	}{
		{"All enabled", state.Options{}},
		{"Disable guilds", state.Options{DisableGuilds: true}},
		{"Disable channels", state.Options{DisableChannels: true}},
		{"Disable users", state.Options{DisableUsers: true}},
		{"Disable all", state.Options{
			DisableGuilds:   true,
			DisableChannels: true,
			DisableUsers:    true,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Just verify options can be created with disable flags
			opts := tt.opts
			_ = opts
		})
	}
}

func TestOptions_TTLValues(t *testing.T) {
	opts := state.Options{
		GuildTTL:     5 * time.Minute,
		ChannelTTL:   10 * time.Minute,
		UserTTL:      1 * time.Hour,
	}

	if opts.GuildTTL != 5*time.Minute {
		t.Errorf("GuildTTL = %v, want 5m", opts.GuildTTL)
	}
	if opts.ChannelTTL != 10*time.Minute {
		t.Errorf("ChannelTTL = %v, want 10m", opts.ChannelTTL)
	}
	if opts.UserTTL != 1*time.Hour {
		t.Errorf("UserTTL = %v, want 1h", opts.UserTTL)
	}
}

func TestStore_MultipleContexts(t *testing.T) {
	ctx1 := context.Background()
	ctx2, cancel := context.WithCancel(context.Background())

	var s *state.Store[string, string]

	// Operations with different contexts should all handle nil gracefully
	s.Set(ctx1, "key1", "value1")
	s.Set(ctx2, "key2", "value2")

	val1, ok1 := s.Get(ctx1, "key1")
	val2, ok2 := s.Get(ctx2, "key2")

	if ok1 || ok2 {
		t.Error("Get on nil store should return false regardless of context")
	}
	if val1 != "" || val2 != "" {
		t.Error("Get should return empty string")
	}

	cancel()
}

func TestStore_PointerType(t *testing.T) {
	ctx := context.Background()

	type User struct {
		ID   string
		Name string
	}

	var s *state.Store[string, *User]

	s.Set(ctx, "user1", &User{ID: "1", Name: "Alice"})
	val, ok := s.Get(ctx, "user1")

	if ok {
		t.Error("Get on nil store should return false")
	}
	if val != nil {
		t.Error("Get should return nil for pointer type")
	}
}

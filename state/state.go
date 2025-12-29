package state

import (
	"sync"

	"github.com/kolosys/discord/models"
	"github.com/kolosys/discord/rest"
)

// State manages Discord entity caches with automatic population from gateway events.
type State struct {
	Guilds      *GuildStore
	Channels    *ChannelStore
	Users       *UserStore
	Members     *MemberStore
	Roles       *RoleStore
	Presences   *PresenceStore
	VoiceStates *VoiceStateStore

	rest        *rest.REST
	opts        *Options
	currentUser *models.User

	mu    sync.RWMutex
	ready bool
}

// New creates a new State with the given options.
func New(r *rest.REST, opts *Options) *State {
	if opts == nil {
		opts = &Options{}
	}
	opts.applyDefaults()
	shards := opts.NumShards

	channels := &ChannelStore{
		Store: newStore[string, *models.GuildChannel](
			cacheConfig{opts.DisableChannels, opts.MaxChannels, opts.ChannelTTL}, shards),
		rest: r,
	}
	roles := &RoleStore{
		Store: newStore[string, *models.GuildRole](
			cacheConfig{opts.DisableRoles, opts.MaxRoles, opts.RoleTTL}, shards),
		rest: r,
	}
	members := &MemberStore{
		Store: newStore[string, *CachedMember](
			cacheConfig{opts.DisableMembers, opts.MaxMembers, opts.MemberTTL}, shards),
		rest: r,
	}

	return &State{
		Guilds: &GuildStore{
			Store: newStore[string, *CachedGuild](
				cacheConfig{opts.DisableGuilds, opts.MaxGuilds, opts.GuildTTL}, shards),
			rest:     r,
			channels: channels,
			roles:    roles,
			members:  members,
		},
		Channels: channels,
		Users: &UserStore{
			Store: newStore[string, *models.User](
				cacheConfig{opts.DisableUsers, opts.MaxUsers, opts.UserTTL}, shards),
			rest: r,
		},
		Members: members,
		Roles:   roles,
		Presences: &PresenceStore{
			Store: newStore[string, *Presence](
				cacheConfig{opts.DisablePresences, opts.MaxPresences, opts.PresenceTTL}, shards),
		},
		VoiceStates: &VoiceStateStore{
			Store: newStore[string, *VoiceState](
				cacheConfig{opts.DisableVoiceStates, opts.MaxVoiceStates, opts.VoiceStateTTL}, shards),
		},
		rest: r,
		opts: opts,
	}
}

// Ready returns whether the state has been initialized from the READY event.
func (s *State) Ready() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.ready
}

// setReady marks the state as ready.
func (s *State) setReady(ready bool) {
	s.mu.Lock()
	s.ready = ready
	s.mu.Unlock()
}

// CurrentUser returns the bot's user from the READY event.
func (s *State) CurrentUser() *models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.currentUser
}

// setCurrentUser sets the bot's user.
func (s *State) setCurrentUser(u *models.User) {
	s.mu.Lock()
	s.currentUser = u
	s.mu.Unlock()
}

// Stats returns statistics for all caches.
func (s *State) Stats() StateStats {
	return StateStats{
		GuildCount:      s.Guilds.Count(),
		ChannelCount:    s.Channels.Count(),
		UserCount:       s.Users.Count(),
		MemberCount:     s.Members.Count(),
		RoleCount:       s.Roles.Count(),
		PresenceCount:   s.Presences.Count(),
		VoiceStateCount: s.VoiceStates.Count(),
	}
}

// InvalidateAll reinitializes all caches, clearing all data.
func (s *State) InvalidateAll() {
	s.mu.Lock()
	defer s.mu.Unlock()

	shards := s.opts.NumShards

	if s.Channels.Enabled() {
		s.Channels.Store = newStore[string, *models.GuildChannel](
			cacheConfig{s.opts.DisableChannels, s.opts.MaxChannels, s.opts.ChannelTTL}, shards)
	}
	if s.Roles.Enabled() {
		s.Roles.Store = newStore[string, *models.GuildRole](
			cacheConfig{s.opts.DisableRoles, s.opts.MaxRoles, s.opts.RoleTTL}, shards)
	}
	if s.Members.Enabled() {
		s.Members.Store = newStore[string, *CachedMember](
			cacheConfig{s.opts.DisableMembers, s.opts.MaxMembers, s.opts.MemberTTL}, shards)
	}
	if s.Guilds.Enabled() {
		s.Guilds.Store = newStore[string, *CachedGuild](
			cacheConfig{s.opts.DisableGuilds, s.opts.MaxGuilds, s.opts.GuildTTL}, shards)
	}
	if s.Users.Enabled() {
		s.Users.Store = newStore[string, *models.User](
			cacheConfig{s.opts.DisableUsers, s.opts.MaxUsers, s.opts.UserTTL}, shards)
	}
	if s.Presences.Enabled() {
		s.Presences.Store = newStore[string, *Presence](
			cacheConfig{s.opts.DisablePresences, s.opts.MaxPresences, s.opts.PresenceTTL}, shards)
	}
	if s.VoiceStates.Enabled() {
		s.VoiceStates.Store = newStore[string, *VoiceState](
			cacheConfig{s.opts.DisableVoiceStates, s.opts.MaxVoiceStates, s.opts.VoiceStateTTL}, shards)
	}
}

// Close cleans up the state store.
func (s *State) Close() error {
	s.InvalidateAll()
	return nil
}

// memberKey generates a composite key for member cache.
func memberKey(guildID, userID string) string {
	return guildID + ":" + userID
}

// roleKey generates a composite key for role cache.
func roleKey(guildID, roleID string) string {
	return guildID + ":" + roleID
}

// presenceKey generates a composite key for presence cache.
func presenceKey(guildID, userID string) string {
	return guildID + ":" + userID
}

// voiceStateKey generates a composite key for voice state cache.
func voiceStateKey(guildID, userID string) string {
	return guildID + ":" + userID
}

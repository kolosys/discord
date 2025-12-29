package state

import "time"

// Options configures state caching behavior.
type Options struct {
	// Cache disable toggles
	DisableGuilds      bool
	DisableChannels    bool
	DisableUsers       bool
	DisableMembers     bool
	DisableRoles       bool
	DisablePresences   bool
	DisableVoiceStates bool

	// Size limits (0 = unlimited)
	MaxGuilds      int
	MaxChannels    int
	MaxUsers       int
	MaxMembers     int
	MaxRoles       int
	MaxPresences   int
	MaxVoiceStates int

	// TTL settings (0 = no expiration)
	GuildTTL      time.Duration
	ChannelTTL    time.Duration
	UserTTL       time.Duration
	MemberTTL     time.Duration
	RoleTTL       time.Duration
	PresenceTTL   time.Duration
	VoiceStateTTL time.Duration

	// Synapse settings
	NumShards   int
	EnableStats bool
}

func (o *Options) applyDefaults() {
	if o.NumShards <= 0 {
		o.NumShards = 16
	}
}

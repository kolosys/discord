package state

import (
	"time"

	"github.com/kolosys/discord/events"
	"github.com/kolosys/discord/models"
)

// CachedGuild extends Guild with additional cached data from GUILD_CREATE.
type CachedGuild struct {
	models.Guild

	JoinedAt    time.Time
	Large       bool
	Unavailable bool
	MemberCount int
	ChannelIDs  []string
	RoleIDs     []string
	MemberIDs   []string
}

// CachedMember wraps GuildMember with guild and user references.
type CachedMember struct {
	models.GuildMember

	GuildID string
	UserID  string
}

// Presence represents a user's presence state in a guild.
type Presence struct {
	UserID       string
	GuildID      string
	Status       string
	Activities   []events.Activity
	ClientStatus events.ClientStatus
}

// VoiceState represents a user's voice connection state in a guild.
type VoiceState struct {
	UserID     string
	GuildID    string
	ChannelID  *string
	SessionID  string
	Deaf       bool
	Mute       bool
	SelfDeaf   bool
	SelfMute   bool
	SelfStream bool
	SelfVideo  bool
	Suppress   bool
}

// StateStats contains aggregated statistics from all caches.
type StateStats struct {
	GuildCount      int
	ChannelCount    int
	UserCount       int
	MemberCount     int
	RoleCount       int
	PresenceCount   int
	VoiceStateCount int
}

package gateway

import (
	"encoding/json"

	"github.com/kolosys/discord/models"
	"github.com/kolosys/discord/types"
)

// GatewayPayload represents a Discord gateway message.
type GatewayPayload struct {
	Op Opcode          `json:"op"`           // Opcode
	D  json.RawMessage `json:"d,omitempty"`  // Data (varies by opcode)
	S  *int            `json:"s,omitempty"`  // Sequence number (only for Dispatch)
	T  *string         `json:"t,omitempty"`  // Event name (only for Dispatch)
}

// HelloData is the data sent in the Hello event (Opcode 10).
type HelloData struct {
	HeartbeatInterval int `json:"heartbeat_interval"` // Interval in milliseconds
}

// ReadyData is the data sent in the Ready event after successful identification.
type ReadyData struct {
	V                int                 `json:"v"`                  // Gateway version
	User             *models.User        `json:"user"`               // Bot user
	Guilds           []*UnavailableGuild `json:"guilds"`             // Guilds (initially unavailable)
	SessionID        string              `json:"session_id"`         // Session ID
	ResumeGatewayURL string              `json:"resume_gateway_url"` // URL for resuming
	Shard            *[2]int             `json:"shard,omitempty"`    // [shard_id, num_shards]
	Application      *PartialApplication `json:"application"`        // Partial application object
}

// UnavailableGuild represents a guild that is initially unavailable in the Ready event.
type UnavailableGuild struct {
	ID          types.Snowflake `json:"id"`
	Unavailable bool            `json:"unavailable"`
}

// PartialApplication contains partial application info from Ready.
type PartialApplication struct {
	ID    types.Snowflake `json:"id"`
	Flags int             `json:"flags"`
}

// IdentifyData is sent to identify a new session.
type IdentifyData struct {
	Token      string               `json:"token"`       // Bot token
	Intents    Intent               `json:"intents"`     // Gateway intents
	Properties IdentifyProperties   `json:"properties"`  // Connection properties
	Compress   bool                 `json:"compress"`    // Whether to use compression
	Shard      *[2]int              `json:"shard,omitempty"` // [shard_id, num_shards]
	Presence   *PresenceUpdate      `json:"presence,omitempty"` // Initial presence
}

// IdentifyProperties contains connection properties for identification.
type IdentifyProperties struct {
	OS      string `json:"os"`      // Operating system
	Browser string `json:"browser"` // Browser/library name
	Device  string `json:"device"`  // Device name
}

// ResumeData is sent to resume a disconnected session.
type ResumeData struct {
	Token     string `json:"token"`      // Bot token
	SessionID string `json:"session_id"` // Session ID to resume
	Seq       int    `json:"seq"`        // Last sequence number received
}

// PresenceUpdate represents a presence update payload.
type PresenceUpdate struct {
	Since      *int              `json:"since"`      // Unix time (ms) of when client went idle, or null
	Activities []*Activity       `json:"activities"` // User's activities
	Status     string            `json:"status"`     // Status (online, dnd, idle, invisible, offline)
	AFK        bool              `json:"afk"`        // Whether client is AFK
}

// Activity represents a user activity.
type Activity struct {
	Name string       `json:"name"`           // Activity name
	Type ActivityType `json:"type"`           // Activity type
	URL  *string      `json:"url,omitempty"`  // Stream URL (only for Streaming type)
}

// ActivityType represents the type of activity.
type ActivityType int

const (
	ActivityTypeGame      ActivityType = 0 // Playing {name}
	ActivityTypeStreaming ActivityType = 1 // Streaming {name}
	ActivityTypeListening ActivityType = 2 // Listening to {name}
	ActivityTypeWatching  ActivityType = 3 // Watching {name}
	ActivityTypeCustom    ActivityType = 4 // Custom status
	ActivityTypeCompeting ActivityType = 5 // Competing in {name}
)

// RequestGuildMembersData is sent to request guild members.
type RequestGuildMembersData struct {
	GuildID   types.Snowflake   `json:"guild_id"`             // Guild ID
	Query     *string           `json:"query,omitempty"`      // String to match usernames (empty for all)
	Limit     int               `json:"limit"`                // Max members to return (0 for all)
	Presences bool              `json:"presences,omitempty"`  // Whether to include presences
	UserIDs   []types.Snowflake `json:"user_ids,omitempty"`   // Specific user IDs to request
	Nonce     *string           `json:"nonce,omitempty"`      // Nonce for response matching
}

// VoiceStateUpdateData is sent to join/leave/move voice channels.
type VoiceStateUpdateData struct {
	GuildID   types.Snowflake  `json:"guild_id"`             // Guild ID
	ChannelID *types.Snowflake `json:"channel_id"`           // Channel ID (null to disconnect)
	SelfMute  bool             `json:"self_mute"`            // Whether to mute
	SelfDeaf  bool             `json:"self_deaf"`            // Whether to deafen
}

// InvalidSessionData indicates whether the session can be resumed.
type InvalidSessionData bool

// MarshalIdentify creates an Identify payload.
func MarshalIdentify(token string, intents Intent, shard *[2]int, presence *PresenceUpdate) (*GatewayPayload, error) {
	data := IdentifyData{
		Token:   token,
		Intents: intents,
		Properties: IdentifyProperties{
			OS:      "linux",
			Browser: "kolosys-discord",
			Device:  "kolosys-discord",
		},
		Compress: false,
		Shard:    shard,
		Presence: presence,
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &GatewayPayload{
		Op: OpcodeIdentify,
		D:  dataJSON,
	}, nil
}

// MarshalResume creates a Resume payload.
func MarshalResume(token, sessionID string, seq int) (*GatewayPayload, error) {
	data := ResumeData{
		Token:     token,
		SessionID: sessionID,
		Seq:       seq,
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &GatewayPayload{
		Op: OpcodeResume,
		D:  dataJSON,
	}, nil
}

// MarshalHeartbeat creates a Heartbeat payload.
func MarshalHeartbeat(seq *int) (*GatewayPayload, error) {
	dataJSON, err := json.Marshal(seq)
	if err != nil {
		return nil, err
	}

	return &GatewayPayload{
		Op: OpcodeHeartbeat,
		D:  dataJSON,
	}, nil
}

// MarshalPresenceUpdate creates a Presence Update payload.
func MarshalPresenceUpdate(presence *PresenceUpdate) (*GatewayPayload, error) {
	dataJSON, err := json.Marshal(presence)
	if err != nil {
		return nil, err
	}

	return &GatewayPayload{
		Op: OpcodePresenceUpdate,
		D:  dataJSON,
	}, nil
}

// MarshalRequestGuildMembers creates a Request Guild Members payload.
func MarshalRequestGuildMembers(data *RequestGuildMembersData) (*GatewayPayload, error) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &GatewayPayload{
		Op: OpcodeRequestGuildMembers,
		D:  dataJSON,
	}, nil
}

// MarshalVoiceStateUpdate creates a Voice State Update payload.
func MarshalVoiceStateUpdate(data *VoiceStateUpdateData) (*GatewayPayload, error) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &GatewayPayload{
		Op: OpcodeVoiceStateUpdate,
		D:  dataJSON,
	}, nil
}

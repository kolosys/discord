package gateway

// Opcode represents a Discord gateway opcode.
type Opcode int

const (
	OpcodeDispatch            Opcode = 0  // Receive: Event dispatched
	OpcodeHeartbeat           Opcode = 1  // Send/Receive: Heartbeat ping
	OpcodeIdentify            Opcode = 2  // Send: Identify connection
	OpcodePresenceUpdate      Opcode = 3  // Send: Update presence
	OpcodeVoiceStateUpdate    Opcode = 4  // Send: Join/leave/move voice channels
	OpcodeResume              Opcode = 6  // Send: Resume a disconnected session
	OpcodeReconnect           Opcode = 7  // Receive: Reconnect to gateway
	OpcodeRequestGuildMembers Opcode = 8  // Send: Request guild members
	OpcodeInvalidSession      Opcode = 9  // Receive: Session invalidated
	OpcodeHello               Opcode = 10 // Receive: Hello with heartbeat interval
	OpcodeHeartbeatACK        Opcode = 11 // Receive: Heartbeat acknowledged
	OpcodeRequestSoundboard   Opcode = 31 // Send: Request soundboard sounds
)

func (o Opcode) String() string {
	switch o {
	case OpcodeDispatch:
		return "DISPATCH"
	case OpcodeHeartbeat:
		return "HEARTBEAT"
	case OpcodeIdentify:
		return "IDENTIFY"
	case OpcodePresenceUpdate:
		return "PRESENCE_UPDATE"
	case OpcodeVoiceStateUpdate:
		return "VOICE_STATE_UPDATE"
	case OpcodeResume:
		return "RESUME"
	case OpcodeReconnect:
		return "RECONNECT"
	case OpcodeRequestGuildMembers:
		return "REQUEST_GUILD_MEMBERS"
	case OpcodeInvalidSession:
		return "INVALID_SESSION"
	case OpcodeHello:
		return "HELLO"
	case OpcodeHeartbeatACK:
		return "HEARTBEAT_ACK"
	case OpcodeRequestSoundboard:
		return "REQUEST_SOUNDBOARD_SOUNDS"
	default:
		return "UNKNOWN"
	}
}

// CloseCode represents a Discord gateway close code.
type CloseCode int

const (
	CloseCodeUnknownError         CloseCode = 4000 // Unknown error
	CloseCodeUnknownOpcode        CloseCode = 4001 // Invalid opcode sent
	CloseCodeDecodeError          CloseCode = 4002 // Invalid payload sent
	CloseCodeNotAuthenticated     CloseCode = 4003 // Payload sent before identifying
	CloseCodeAuthenticationFailed CloseCode = 4004 // Invalid token
	CloseCodeAlreadyAuthenticated CloseCode = 4005 // Already identified
	CloseCodeInvalidSeq           CloseCode = 4007 // Invalid sequence number
	CloseCodeRateLimited          CloseCode = 4008 // Too many payloads sent
	CloseCodeSessionTimeout       CloseCode = 4009 // Session timed out
	CloseCodeInvalidShard         CloseCode = 4010 // Invalid shard
	CloseCodeShardingRequired     CloseCode = 4011 // Sharding required
	CloseCodeInvalidAPIVersion    CloseCode = 4012 // Invalid API version
	CloseCodeInvalidIntents       CloseCode = 4013 // Invalid intents
	CloseCodeDisallowedIntents    CloseCode = 4014 // Disallowed intents
)

func (c CloseCode) String() string {
	switch c {
	case CloseCodeUnknownError:
		return "Unknown Error"
	case CloseCodeUnknownOpcode:
		return "Unknown Opcode"
	case CloseCodeDecodeError:
		return "Decode Error"
	case CloseCodeNotAuthenticated:
		return "Not Authenticated"
	case CloseCodeAuthenticationFailed:
		return "Authentication Failed"
	case CloseCodeAlreadyAuthenticated:
		return "Already Authenticated"
	case CloseCodeInvalidSeq:
		return "Invalid Sequence"
	case CloseCodeRateLimited:
		return "Rate Limited"
	case CloseCodeSessionTimeout:
		return "Session Timeout"
	case CloseCodeInvalidShard:
		return "Invalid Shard"
	case CloseCodeShardingRequired:
		return "Sharding Required"
	case CloseCodeInvalidAPIVersion:
		return "Invalid API Version"
	case CloseCodeInvalidIntents:
		return "Invalid Intents"
	case CloseCodeDisallowedIntents:
		return "Disallowed Intents"
	default:
		return "Unknown"
	}
}

// Description returns a detailed, actionable description of the close code.
func (c CloseCode) Description() string {
	switch c {
	case CloseCodeUnknownError:
		return "An unknown error occurred. This is typically transient - reconnection will be attempted."
	case CloseCodeUnknownOpcode:
		return "An invalid opcode was sent. This indicates a bug in the gateway client."
	case CloseCodeDecodeError:
		return "An invalid payload was sent. This indicates malformed JSON or incorrect payload structure."
	case CloseCodeNotAuthenticated:
		return "A payload was sent before identifying. Ensure Identify is sent immediately after Hello."
	case CloseCodeAuthenticationFailed:
		return "The provided token is invalid. Verify DISCORD_TOKEN is correct and the bot has not been deleted."
	case CloseCodeAlreadyAuthenticated:
		return "An Identify payload was sent after already authenticating. This indicates a client bug."
	case CloseCodeInvalidSeq:
		return "An invalid sequence number was sent when resuming. The session will be invalidated."
	case CloseCodeRateLimited:
		return "Too many payloads were sent. Reduce the frequency of gateway commands."
	case CloseCodeSessionTimeout:
		return "The session timed out due to inactivity. Ensure heartbeats are being sent correctly."
	case CloseCodeInvalidShard:
		return "An invalid shard configuration was sent. Verify shard ID is less than total shard count."
	case CloseCodeShardingRequired:
		return "The bot is in too many guilds and must use sharding. Enable sharding in your client configuration."
	case CloseCodeInvalidAPIVersion:
		return "An invalid API version was specified. Update the gateway URL to use a supported version (v10)."
	case CloseCodeInvalidIntents:
		return "Invalid intents were specified. Verify intent values are valid Discord gateway intents."
	case CloseCodeDisallowedIntents:
		return "Privileged intents were requested but not enabled. " +
			"Enable the required intents in the Discord Developer Portal: " +
			"https://discord.com/developers/applications → Bot → Privileged Gateway Intents. " +
			"Privileged intents include: GUILD_MEMBERS, GUILD_PRESENCES, and MESSAGE_CONTENT."
	default:
		return "An unrecognized close code was received."
	}
}

// CanReconnect returns true if the error is recoverable and reconnection should be attempted.
func (c CloseCode) CanReconnect() bool {
	switch c {
	case CloseCodeAuthenticationFailed,
		CloseCodeInvalidShard,
		CloseCodeShardingRequired,
		CloseCodeInvalidAPIVersion,
		CloseCodeInvalidIntents,
		CloseCodeDisallowedIntents:
		return false
	default:
		return true
	}
}

// IsFatal returns true if the error is non-recoverable and the client should stop.
func (c CloseCode) IsFatal() bool {
	return !c.CanReconnect()
}

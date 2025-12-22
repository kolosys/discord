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

# gateway API

Complete API documentation for the gateway package.

**Import Path:** `github.com/kolosys/discord/gateway`

## Package Documentation



## Types

### Activity
Activity represents a user activity.

#### Example Usage

```go
// Create a new Activity
activity := Activity{
    Name: "example",
    Type: ActivityType{},
    URL: &"example"{},
    State: &"example"{},
}
```

#### Type Definition

```go
type Activity struct {
    Name string `json:"name"`
    Type ActivityType `json:"type"`
    URL *string `json:"url,omitempty"`
    State *string `json:"state,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` | Activity name (required) |
| Type | `ActivityType` | Activity type |
| URL | `*string` | Stream URL (only for Streaming type) |
| State | `*string` | Custom status text (only for Custom type, user accounts only) |

### ActivityType
ActivityType represents the type of activity.

#### Example Usage

```go
// Example usage of ActivityType
var value ActivityType
// Initialize with appropriate value
```

#### Type Definition

```go
type ActivityType int
```

### CloseCode
CloseCode represents a Discord gateway close code. For standard WebSocket codes (1000-1015), use axon.CloseCode directly. This type focuses on Discord-specific codes (4000+) with detailed descriptions.

#### Example Usage

```go
// Example usage of CloseCode
var value CloseCode
// Initialize with appropriate value
```

#### Type Definition

```go
type CloseCode int
```

## Methods

### CanReconnect

CanReconnect returns true if the error is recoverable and reconnection should be attempted.

```go
func (CloseCode) CanReconnect() bool
```

**Parameters:**
  None

**Returns:**
- bool

### Description

Description returns a detailed, actionable description of the close code. For Discord-specific codes, provides actionable guidance. For standard WebSocket codes, provides context relevant to Discord gateway usage.

```go
func (CloseCode) Description() string
```

**Parameters:**
  None

**Returns:**
- string

### IsFatal

IsFatal returns true if the error is non-recoverable and the client should stop.

```go
func (CloseCode) IsFatal() bool
```

**Parameters:**
  None

**Returns:**
- bool

### String



```go
func (CloseCode) String() string
```

**Parameters:**
  None

**Returns:**
- string

### Gateway
Gateway manages Discord gateway connections across multiple shards.

#### Example Usage

```go
// Create a new Gateway
gateway := Gateway{

}
```

#### Type Definition

```go
type Gateway struct {
}
```

### Constructor Functions

### NewGateway

NewGateway creates a new gateway manager.

```go
func NewGateway(token string, intents Intent, events bus.EventBus) *Gateway
```

**Parameters:**
- `token` (string)
- `intents` (Intent)
- `events` (bus.EventBus)

**Returns:**
- *Gateway

## Methods

### Close

Close gracefully closes all shard connections.

```go
func (*Shard) Close() error
```

**Parameters:**
  None

**Returns:**
- error

### Connect

Connect establishes connections to the Discord gateway for all shards. url is the gateway URL (e.g., "wss://gateway.discord.gg") shardCount is the total number of shards to create (1 for no sharding)

```go
func (*Shard) Connect(ctx context.Context, url string) error
```

**Parameters:**
- `ctx` (context.Context)
- `url` (string)

**Returns:**
- error

### Events

Events returns the event bus for subscribing to gateway events.

```go
func (*Gateway) Events() bus.EventBus
```

**Parameters:**
  None

**Returns:**
- bus.EventBus

### GetShard

GetShard returns a specific shard by ID.

```go
func (*Gateway) GetShard(id int) (*Shard, bool)
```

**Parameters:**
- `id` (int)

**Returns:**
- *Shard
- bool

### IsReady

IsReady returns whether all shards have received the READY event.

```go
func (*Shard) IsReady() bool
```

**Parameters:**
  None

**Returns:**
- bool

### IsRunning

IsRunning returns whether the gateway is currently running.

```go
func (*Gateway) IsRunning() bool
```

**Parameters:**
  None

**Returns:**
- bool

### RequestGuildMembers

RequestGuildMembers requests guild members for a specific guild. The request is routed to the shard responsible for that guild. Results are received via GUILD_MEMBERS_CHUNK events.

```go
func (*Shard) RequestGuildMembers(ctx context.Context, data *RequestGuildMembersData) error
```

**Parameters:**
- `ctx` (context.Context)
- `data` (*RequestGuildMembersData)

**Returns:**
- error

### ShardCount

ShardCount returns the number of shards.

```go
func (*Gateway) ShardCount() int
```

**Parameters:**
  None

**Returns:**
- int

### UpdatePresence

UpdatePresence updates the bot's presence on all shards.

```go
func (*Shard) UpdatePresence(ctx context.Context, presence *PresenceUpdate) error
```

**Parameters:**
- `ctx` (context.Context)
- `presence` (*PresenceUpdate)

**Returns:**
- error

### UpdateVoiceState

UpdateVoiceState updates the bot's voice state for a specific guild. The request is routed to the shard responsible for that guild.

```go
func (*Shard) UpdateVoiceState(ctx context.Context, data *VoiceStateUpdateData) error
```

**Parameters:**
- `ctx` (context.Context)
- `data` (*VoiceStateUpdateData)

**Returns:**
- error

### GatewayClient
GatewayClient defines the interface for WebSocket client operations required by Shard. It enables dependency injection for testing with axon/mock while maintaining type safety.

#### Example Usage

```go
// Example implementation of GatewayClient
type MyGatewayClient struct {
    // Add your fields here
}

func (m MyGatewayClient) ConnectWithReadLoop(param1 context.Context) error {
    // Implement your logic here
    return
}

func (m MyGatewayClient) Write(param1 context.Context, param2 T) error {
    // Implement your logic here
    return
}

func (m MyGatewayClient) Close() error {
    // Implement your logic here
    return
}

func (m MyGatewayClient) OnMessage(param1 func(T))  {
    // Implement your logic here
    return
}

func (m MyGatewayClient) OnConnect(param1 func(**ast.IndexExpr))  {
    // Implement your logic here
    return
}

func (m MyGatewayClient) OnDisconnect(param1 func(**ast.IndexExpr, error))  {
    // Implement your logic here
    return
}

func (m MyGatewayClient) State() axon.ConnectionState {
    // Implement your logic here
    return
}

func (m MyGatewayClient) IsConnected() bool {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type GatewayClient interface {
    ConnectWithReadLoop(ctx context.Context) error
    Write(ctx context.Context, msg T) error
    Close() error
    OnMessage(fn func(T))
    OnConnect(fn func(**ast.IndexExpr))
    OnDisconnect(fn func(**ast.IndexExpr, error))
    State() axon.ConnectionState
    IsConnected() bool
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### GatewayPayload
GatewayPayload represents a Discord gateway message.

#### Example Usage

```go
// Create a new GatewayPayload
gatewaypayload := GatewayPayload{
    Op: Opcode{},
    D: /* value */,
    S: &42{},
    T: &"example"{},
}
```

#### Type Definition

```go
type GatewayPayload struct {
    Op Opcode `json:"op"`
    D json.RawMessage `json:"d,omitempty"`
    S *int `json:"s,omitempty"`
    T *string `json:"t,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Op | `Opcode` | Opcode |
| D | `json.RawMessage` | Data (varies by opcode) |
| S | `*int` | Sequence number (only for Dispatch) |
| T | `*string` | Event name (only for Dispatch) |

### Constructor Functions

### MarshalHeartbeat

MarshalHeartbeat creates a Heartbeat payload.

```go
func MarshalHeartbeat(seq *int) (*GatewayPayload, error)
```

**Parameters:**
- `seq` (*int)

**Returns:**
- *GatewayPayload
- error

### MarshalIdentify

MarshalIdentify creates an Identify payload.

```go
func MarshalIdentify(token string, intents Intent, shard *[*ast.BasicLit]int, presence *PresenceUpdate) (*GatewayPayload, error)
```

**Parameters:**
- `token` (string)
- `intents` (Intent)
- `shard` (*[*ast.BasicLit]int)
- `presence` (*PresenceUpdate)

**Returns:**
- *GatewayPayload
- error

### MarshalPresenceUpdate

MarshalPresenceUpdate creates a Presence Update payload.

```go
func MarshalPresenceUpdate(presence *PresenceUpdate) (*GatewayPayload, error)
```

**Parameters:**
- `presence` (*PresenceUpdate)

**Returns:**
- *GatewayPayload
- error

### MarshalRequestGuildMembers

MarshalRequestGuildMembers creates a Request Guild Members payload.

```go
func MarshalRequestGuildMembers(data *RequestGuildMembersData) (*GatewayPayload, error)
```

**Parameters:**
- `data` (*RequestGuildMembersData)

**Returns:**
- *GatewayPayload
- error

### MarshalResume

MarshalResume creates a Resume payload.

```go
func MarshalResume(token, sessionID string, seq int) (*GatewayPayload, error)
```

**Parameters:**
- `token` (string)
- `sessionID` (string)
- `seq` (int)

**Returns:**
- *GatewayPayload
- error

### MarshalVoiceStateUpdate

MarshalVoiceStateUpdate creates a Voice State Update payload.

```go
func MarshalVoiceStateUpdate(data *VoiceStateUpdateData) (*GatewayPayload, error)
```

**Parameters:**
- `data` (*VoiceStateUpdateData)

**Returns:**
- *GatewayPayload
- error

### HeartbeatManager
HeartbeatManager manages the heartbeat loop for a Discord gateway connection.

#### Example Usage

```go
// Create a new HeartbeatManager
heartbeatmanager := HeartbeatManager{

}
```

#### Type Definition

```go
type HeartbeatManager struct {
}
```

### Constructor Functions

### NewHeartbeatManager

NewHeartbeatManager creates a new heartbeat manager.

```go
func NewHeartbeatManager(session *Session) *HeartbeatManager
```

**Parameters:**
- `session` (*Session)

**Returns:**
- *HeartbeatManager

## Methods

### ACK

ACK marks a heartbeat as acknowledged.

```go
func (*HeartbeatManager) ACK()
```

**Parameters:**
  None

**Returns:**
  None

### IsHealthy

IsHealthy checks if the last heartbeat was acknowledged. Returns false if a heartbeat was sent but not acknowledged, indicating a zombie connection.

```go
func (*HeartbeatManager) IsHealthy() bool
```

**Parameters:**
  None

**Returns:**
- bool

### LastACK

LastACK returns the timestamp of the last received ACK.

```go
func (*HeartbeatManager) LastACK() time.Time
```

**Parameters:**
  None

**Returns:**
- time.Time

### LastHeartbeat

LastHeartbeat returns the timestamp of the last sent heartbeat.

```go
func (*HeartbeatManager) LastHeartbeat() time.Time
```

**Parameters:**
  None

**Returns:**
- time.Time

### Start

Start begins the heartbeat loop. parent is the parent context for cancellation propagation. interval is the heartbeat interval. sendFunc is called to send the heartbeat payload.

```go
func (*HeartbeatManager) Start(parent context.Context, interval time.Duration, sendFunc func(seq int) error) error
```

**Parameters:**
- `parent` (context.Context)
- `interval` (time.Duration)
- `sendFunc` (func(seq int) error)

**Returns:**
- error

### Stop

Stop stops the heartbeat loop.

```go
func (*HeartbeatManager) Stop()
```

**Parameters:**
  None

**Returns:**
  None

### String

String returns a string representation of the heartbeat manager state.

```go
func (Intent) String() string
```

**Parameters:**
  None

**Returns:**
- string

### HelloData
HelloData is the data sent in the Hello event (Opcode 10).

#### Example Usage

```go
// Create a new HelloData
hellodata := HelloData{
    HeartbeatInterval: 42,
}
```

#### Type Definition

```go
type HelloData struct {
    HeartbeatInterval int `json:"heartbeat_interval"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| HeartbeatInterval | `int` | Interval in milliseconds |

### IdentifyData
IdentifyData is sent to identify a new session.

#### Example Usage

```go
// Create a new IdentifyData
identifydata := IdentifyData{
    Token: "example",
    Intents: Intent{},
    Properties: IdentifyProperties{},
    Compress: true,
    Shard: &[]{},
    Presence: &PresenceUpdate{}{},
}
```

#### Type Definition

```go
type IdentifyData struct {
    Token string `json:"token"`
    Intents Intent `json:"intents"`
    Properties IdentifyProperties `json:"properties"`
    Compress bool `json:"compress"`
    Shard *[*ast.BasicLit]int `json:"shard,omitempty"`
    Presence *PresenceUpdate `json:"presence,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Token | `string` | Bot token |
| Intents | `Intent` | Gateway intents |
| Properties | `IdentifyProperties` | Connection properties |
| Compress | `bool` | Whether to use compression |
| Shard | `*[*ast.BasicLit]int` | [shard_id, num_shards] |
| Presence | `*PresenceUpdate` | Initial presence |

### IdentifyProperties
IdentifyProperties contains connection properties for identification.

#### Example Usage

```go
// Create a new IdentifyProperties
identifyproperties := IdentifyProperties{
    OS: "example",
    Browser: "example",
    Device: "example",
}
```

#### Type Definition

```go
type IdentifyProperties struct {
    OS string `json:"os"`
    Browser string `json:"browser"`
    Device string `json:"device"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| OS | `string` | Operating system |
| Browser | `string` | Browser/library name |
| Device | `string` | Device name |

### Intent
Intent represents a Discord gateway intent bitfield.

#### Example Usage

```go
// Example usage of Intent
var value Intent
// Initialize with appropriate value
```

#### Type Definition

```go
type Intent int
```

## Methods

### Add



```go
func (Intent) Add(intent Intent) Intent
```

**Parameters:**
- `intent` (Intent)

**Returns:**
- Intent

### Has



```go
func (Intent) Has(intent Intent) bool
```

**Parameters:**
- `intent` (Intent)

**Returns:**
- bool

### PrivilegedIntents

PrivilegedIntents returns which privileged intents are included.

```go
func (Intent) PrivilegedIntents() []string
```

**Parameters:**
  None

**Returns:**
- []string

### Remove



```go
func (Intent) Remove(intent Intent) Intent
```

**Parameters:**
- `intent` (Intent)

**Returns:**
- Intent

### String

String returns a human-readable list of enabled intents.

```go
func (*HeartbeatManager) String() string
```

**Parameters:**
  None

**Returns:**
- string

### Warnings

Warnings returns helpful warnings about the intent configuration.

```go
func (Intent) Warnings() []string
```

**Parameters:**
  None

**Returns:**
- []string

### InvalidSessionData
InvalidSessionData indicates whether the session can be resumed.

#### Example Usage

```go
// Example usage of InvalidSessionData
var value InvalidSessionData
// Initialize with appropriate value
```

#### Type Definition

```go
type InvalidSessionData bool
```

### Opcode
Opcode represents a Discord gateway opcode.

#### Example Usage

```go
// Example usage of Opcode
var value Opcode
// Initialize with appropriate value
```

#### Type Definition

```go
type Opcode int
```

## Methods

### String



```go
func (*HeartbeatManager) String() string
```

**Parameters:**
  None

**Returns:**
- string

### PartialApplication
PartialApplication contains partial application info from Ready.

#### Example Usage

```go
// Create a new PartialApplication
partialapplication := PartialApplication{
    ID: /* value */,
    Flags: 42,
}
```

#### Type Definition

```go
type PartialApplication struct {
    ID types.Snowflake `json:"id"`
    Flags int `json:"flags"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `types.Snowflake` |  |
| Flags | `int` |  |

### PresenceUpdate
PresenceUpdate represents a presence update payload.

#### Example Usage

```go
// Create a new PresenceUpdate
presenceupdate := PresenceUpdate{
    Since: &42{},
    Activities: [],
    Status: "example",
    AFK: true,
}
```

#### Type Definition

```go
type PresenceUpdate struct {
    Since *int `json:"since"`
    Activities []*Activity `json:"activities"`
    Status string `json:"status"`
    AFK bool `json:"afk"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Since | `*int` | Unix time (ms) of when client went idle, or null |
| Activities | `[]*Activity` | User's activities |
| Status | `string` | Status (online, dnd, idle, invisible, offline) |
| AFK | `bool` | Whether client is AFK |

### ReadyData
ReadyData is the data sent in the Ready event after successful identification.

#### Example Usage

```go
// Create a new ReadyData
readydata := ReadyData{
    V: 42,
    User: &/* value */{},
    Guilds: [],
    SessionID: "example",
    ResumeGatewayURL: "example",
    Shard: &[]{},
    Application: &PartialApplication{}{},
}
```

#### Type Definition

```go
type ReadyData struct {
    V int `json:"v"`
    User *models.User `json:"user"`
    Guilds []*UnavailableGuild `json:"guilds"`
    SessionID string `json:"session_id"`
    ResumeGatewayURL string `json:"resume_gateway_url"`
    Shard *[*ast.BasicLit]int `json:"shard,omitempty"`
    Application *PartialApplication `json:"application"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| V | `int` | Gateway version |
| User | `*models.User` | Bot user |
| Guilds | `[]*UnavailableGuild` | Guilds (initially unavailable) |
| SessionID | `string` | Session ID |
| ResumeGatewayURL | `string` | URL for resuming |
| Shard | `*[*ast.BasicLit]int` | [shard_id, num_shards] |
| Application | `*PartialApplication` | Partial application object |

### RequestGuildMembersData
RequestGuildMembersData is sent to request guild members.

#### Example Usage

```go
// Create a new RequestGuildMembersData
requestguildmembersdata := RequestGuildMembersData{
    GuildID: /* value */,
    Query: &"example"{},
    Limit: 42,
    Presences: true,
    UserIDs: [],
    Nonce: &"example"{},
}
```

#### Type Definition

```go
type RequestGuildMembersData struct {
    GuildID types.Snowflake `json:"guild_id"`
    Query *string `json:"query,omitempty"`
    Limit int `json:"limit"`
    Presences bool `json:"presences,omitempty"`
    UserIDs []types.Snowflake `json:"user_ids,omitempty"`
    Nonce *string `json:"nonce,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| GuildID | `types.Snowflake` | Guild ID |
| Query | `*string` | String to match usernames (empty for all) |
| Limit | `int` | Max members to return (0 for all) |
| Presences | `bool` | Whether to include presences |
| UserIDs | `[]types.Snowflake` | Specific user IDs to request |
| Nonce | `*string` | Nonce for response matching |

### ResumeData
ResumeData is sent to resume a disconnected session.

#### Example Usage

```go
// Create a new ResumeData
resumedata := ResumeData{
    Token: "example",
    SessionID: "example",
    Seq: 42,
}
```

#### Type Definition

```go
type ResumeData struct {
    Token string `json:"token"`
    SessionID string `json:"session_id"`
    Seq int `json:"seq"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Token | `string` | Bot token |
| SessionID | `string` | Session ID to resume |
| Seq | `int` | Last sequence number received |

### Session
Session manages the state of a Discord gateway session.

#### Example Usage

```go
// Create a new Session
session := Session{

}
```

#### Type Definition

```go
type Session struct {
}
```

### Constructor Functions

### NewSession

NewSession creates a new session.

```go
func NewSession() *Session
```

**Parameters:**
  None

**Returns:**
- *Session

## Methods

### CanResume

CanResume returns whether the session can be resumed. A session can be resumed if it has a session ID, a resume URL, and has been identified.

```go
func (*Session) CanResume() bool
```

**Parameters:**
  None

**Returns:**
- bool

### IncrementSequence

IncrementSequence atomically increments the sequence number and returns the new value.

```go
func (*Session) IncrementSequence() int
```

**Parameters:**
  None

**Returns:**
- int

### IsIdentified

IsIdentified returns whether the session has been identified.

```go
func (*Session) IsIdentified() bool
```

**Parameters:**
  None

**Returns:**
- bool

### MarkIdentified

MarkIdentified marks the session as identified.

```go
func (*Session) MarkIdentified()
```

**Parameters:**
  None

**Returns:**
  None

### Reset

Reset clears the session state.

```go
func (*Session) Reset()
```

**Parameters:**
  None

**Returns:**
  None

### ResumeURL

ResumeURL returns the resume gateway URL.

```go
func (*Session) ResumeURL() string
```

**Parameters:**
  None

**Returns:**
- string

### Sequence

Sequence returns the current sequence number.

```go
func (*Session) Sequence() int
```

**Parameters:**
  None

**Returns:**
- int

### SessionID

SessionID returns the current session ID.

```go
func (*Session) SessionID() string
```

**Parameters:**
  None

**Returns:**
- string

### SetResumeURL

SetResumeURL stores the resume gateway URL from the READY event.

```go
func (*Session) SetResumeURL(url string)
```

**Parameters:**
- `url` (string)

**Returns:**
  None

### SetSequence

SetSequence updates the sequence number.

```go
func (*Session) SetSequence(seq int)
```

**Parameters:**
- `seq` (int)

**Returns:**
  None

### SetSessionID

SetSessionID stores the session ID from the READY event.

```go
func (*Session) SetSessionID(id string)
```

**Parameters:**
- `id` (string)

**Returns:**
  None

### Shard
Shard represents a single Discord gateway shard connection.

#### Example Usage

```go
// Create a new Shard
shard := Shard{

}
```

#### Type Definition

```go
type Shard struct {
}
```

### Constructor Functions

### NewShard

NewShard creates a new shard.

```go
func NewShard(id, total int, token string, intents Intent, events bus.EventBus) *Shard
```

**Parameters:**
- `id` (int)
- `total` (int)
- `token` (string)
- `intents` (Intent)
- `events` (bus.EventBus)

**Returns:**
- *Shard

### NewShardWithClient

NewShardWithClient creates a new shard with a custom client (for testing).

```go
func NewShardWithClient(id, total int, token string, intents Intent, events bus.EventBus, client *ast.IndexExpr) *Shard
```

**Parameters:**
- `id` (int)
- `total` (int)
- `token` (string)
- `intents` (Intent)
- `events` (bus.EventBus)
- `client` (*ast.IndexExpr)

**Returns:**
- *Shard

## Methods

### Close

Close gracefully closes the shard connection.

```go
func (*Shard) Close() error
```

**Parameters:**
  None

**Returns:**
- error

### Connect

Connect establishes a WebSocket connection to the gateway.

```go
func (*Gateway) Connect(ctx context.Context, url string, shardCount int) error
```

**Parameters:**
- `ctx` (context.Context)
- `url` (string)
- `shardCount` (int)

**Returns:**
- error

### FatalError

FatalError returns the fatal error if one occurred, or nil.

```go
func (*Shard) FatalError() error
```

**Parameters:**
  None

**Returns:**
- error

### IsReady

IsReady returns whether the shard has received the READY event.

```go
func (*Shard) IsReady() bool
```

**Parameters:**
  None

**Returns:**
- bool

### RequestGuildMembers

RequestGuildMembers requests guild members from Discord. Results are received via GUILD_MEMBERS_CHUNK events.

```go
func (*Shard) RequestGuildMembers(ctx context.Context, data *RequestGuildMembersData) error
```

**Parameters:**
- `ctx` (context.Context)
- `data` (*RequestGuildMembersData)

**Returns:**
- error

### UpdatePresence

UpdatePresence updates the bot's presence/status.

```go
func (*Shard) UpdatePresence(ctx context.Context, presence *PresenceUpdate) error
```

**Parameters:**
- `ctx` (context.Context)
- `presence` (*PresenceUpdate)

**Returns:**
- error

### UpdateVoiceState

UpdateVoiceState updates the bot's voice state (join/leave/move voice channels).

```go
func (*Shard) UpdateVoiceState(ctx context.Context, data *VoiceStateUpdateData) error
```

**Parameters:**
- `ctx` (context.Context)
- `data` (*VoiceStateUpdateData)

**Returns:**
- error

### UnavailableGuild
UnavailableGuild represents a guild that is initially unavailable in the Ready event.

#### Example Usage

```go
// Create a new UnavailableGuild
unavailableguild := UnavailableGuild{
    ID: /* value */,
    Unavailable: true,
}
```

#### Type Definition

```go
type UnavailableGuild struct {
    ID types.Snowflake `json:"id"`
    Unavailable bool `json:"unavailable"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `types.Snowflake` |  |
| Unavailable | `bool` |  |

### VoiceStateUpdateData
VoiceStateUpdateData is sent to join/leave/move voice channels.

#### Example Usage

```go
// Create a new VoiceStateUpdateData
voicestateupdatedata := VoiceStateUpdateData{
    GuildID: /* value */,
    ChannelID: &/* value */{},
    SelfMute: true,
    SelfDeaf: true,
}
```

#### Type Definition

```go
type VoiceStateUpdateData struct {
    GuildID types.Snowflake `json:"guild_id"`
    ChannelID *types.Snowflake `json:"channel_id"`
    SelfMute bool `json:"self_mute"`
    SelfDeaf bool `json:"self_deaf"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| GuildID | `types.Snowflake` | Guild ID |
| ChannelID | `*types.Snowflake` | Channel ID (null to disconnect) |
| SelfMute | `bool` | Whether to mute |
| SelfDeaf | `bool` | Whether to deafen |

## External Links

- [Package Overview](../packages/gateway.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/discord/gateway)
- [Source Code](https://github.com/kolosys/discord/tree/main/gateway)

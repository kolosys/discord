# state API

Complete API documentation for the state package.

**Import Path:** `github.com/kolosys/discord/state`

## Package Documentation



## Variables

**ErrNotFound, ErrDisabled, ErrRESTNotAvailable, ErrNotReady**



```go
var ErrNotFound = errors.New("state: entity not found")
var ErrDisabled = errors.New("state: cache disabled")
var ErrRESTNotAvailable = errors.New("state: REST client not available")
var ErrNotReady = errors.New("state: not ready")
```

## Types

### CachedGuild
CachedGuild extends Guild with additional cached data from GUILD_CREATE.

#### Example Usage

```go
// Create a new CachedGuild
cachedguild := CachedGuild{
    JoinedAt: /* value */,
    Large: true,
    Unavailable: true,
    MemberCount: 42,
    ChannelIDs: [],
    RoleIDs: [],
    MemberIDs: [],
}
```

#### Type Definition

```go
type CachedGuild struct {
    models.Guild
    JoinedAt time.Time
    Large bool
    Unavailable bool
    MemberCount int
    ChannelIDs []string
    RoleIDs []string
    MemberIDs []string
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *models.Guild | `models.Guild` |  |
| JoinedAt | `time.Time` |  |
| Large | `bool` |  |
| Unavailable | `bool` |  |
| MemberCount | `int` |  |
| ChannelIDs | `[]string` |  |
| RoleIDs | `[]string` |  |
| MemberIDs | `[]string` |  |

### CachedMember
CachedMember wraps GuildMember with guild and user references.

#### Example Usage

```go
// Create a new CachedMember
cachedmember := CachedMember{
    GuildID: "example",
    UserID: "example",
}
```

#### Type Definition

```go
type CachedMember struct {
    models.GuildMember
    GuildID string
    UserID string
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *models.GuildMember | `models.GuildMember` |  |
| GuildID | `string` |  |
| UserID | `string` |  |

### ChannelStore
ChannelStore manages channel cache.

#### Example Usage

```go
// Create a new ChannelStore
channelstore := ChannelStore{

}
```

#### Type Definition

```go
type ChannelStore struct {
    **ast.IndexListExpr
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ***ast.IndexListExpr | `**ast.IndexListExpr` |  |

## Methods

### Refresh



```go
func (*MemberStore) Refresh(ctx context.Context, guildID, userID string) (*CachedMember, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)
- `userID` (string)

**Returns:**
- *CachedMember
- error

### GuildStore
GuildStore manages guild cache with relationship methods.

#### Example Usage

```go
// Create a new GuildStore
guildstore := GuildStore{

}
```

#### Type Definition

```go
type GuildStore struct {
    **ast.IndexListExpr
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ***ast.IndexListExpr | `**ast.IndexListExpr` |  |

## Methods

### Channels



```go
func (*GuildStore) Channels(ctx context.Context, guildID string) ([]*models.GuildChannel, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)

**Returns:**
- []*models.GuildChannel
- error

### Members



```go
func (*GuildStore) Members(ctx context.Context, guildID string) ([]*CachedMember, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)

**Returns:**
- []*CachedMember
- error

### Refresh



```go
func (*MemberStore) Refresh(ctx context.Context, guildID, userID string) (*CachedMember, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)
- `userID` (string)

**Returns:**
- *CachedMember
- error

### Roles



```go
func (*GuildStore) Roles(ctx context.Context, guildID string) ([]*models.GuildRole, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)

**Returns:**
- []*models.GuildRole
- error

### MemberStore
MemberStore manages member cache with composite keys.

#### Example Usage

```go
// Create a new MemberStore
memberstore := MemberStore{

}
```

#### Type Definition

```go
type MemberStore struct {
    **ast.IndexListExpr
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ***ast.IndexListExpr | `**ast.IndexListExpr` |  |

## Methods

### Delete



```go
func (**ast.IndexListExpr) Delete(ctx context.Context, key K) bool
```

**Parameters:**
- `ctx` (context.Context)
- `key` (K)

**Returns:**
- bool

### Get



```go
func (**ast.IndexListExpr) Get(ctx context.Context, key K) (V, bool)
```

**Parameters:**
- `ctx` (context.Context)
- `key` (K)

**Returns:**
- V
- bool

### Refresh



```go
func (*MemberStore) Refresh(ctx context.Context, guildID, userID string) (*CachedMember, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)
- `userID` (string)

**Returns:**
- *CachedMember
- error

### Set



```go
func (**ast.IndexListExpr) Set(ctx context.Context, key K, value V)
```

**Parameters:**
- `ctx` (context.Context)
- `key` (K)
- `value` (V)

**Returns:**
  None

### Options
Options configures state caching behavior.

#### Example Usage

```go
// Create a new Options
options := Options{
    DisableGuilds: true,
    DisableChannels: true,
    DisableUsers: true,
    DisableMembers: true,
    DisableRoles: true,
    DisablePresences: true,
    DisableVoiceStates: true,
    MaxGuilds: 42,
    MaxChannels: 42,
    MaxUsers: 42,
    MaxMembers: 42,
    MaxRoles: 42,
    MaxPresences: 42,
    MaxVoiceStates: 42,
    GuildTTL: /* value */,
    ChannelTTL: /* value */,
    UserTTL: /* value */,
    MemberTTL: /* value */,
    RoleTTL: /* value */,
    PresenceTTL: /* value */,
    VoiceStateTTL: /* value */,
    NumShards: 42,
    EnableStats: true,
}
```

#### Type Definition

```go
type Options struct {
    DisableGuilds bool
    DisableChannels bool
    DisableUsers bool
    DisableMembers bool
    DisableRoles bool
    DisablePresences bool
    DisableVoiceStates bool
    MaxGuilds int
    MaxChannels int
    MaxUsers int
    MaxMembers int
    MaxRoles int
    MaxPresences int
    MaxVoiceStates int
    GuildTTL time.Duration
    ChannelTTL time.Duration
    UserTTL time.Duration
    MemberTTL time.Duration
    RoleTTL time.Duration
    PresenceTTL time.Duration
    VoiceStateTTL time.Duration
    NumShards int
    EnableStats bool
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| DisableGuilds | `bool` | Cache disable toggles |
| DisableChannels | `bool` |  |
| DisableUsers | `bool` |  |
| DisableMembers | `bool` |  |
| DisableRoles | `bool` |  |
| DisablePresences | `bool` |  |
| DisableVoiceStates | `bool` |  |
| MaxGuilds | `int` | Size limits (0 = unlimited) |
| MaxChannels | `int` |  |
| MaxUsers | `int` |  |
| MaxMembers | `int` |  |
| MaxRoles | `int` |  |
| MaxPresences | `int` |  |
| MaxVoiceStates | `int` |  |
| GuildTTL | `time.Duration` | TTL settings (0 = no expiration) |
| ChannelTTL | `time.Duration` |  |
| UserTTL | `time.Duration` |  |
| MemberTTL | `time.Duration` |  |
| RoleTTL | `time.Duration` |  |
| PresenceTTL | `time.Duration` |  |
| VoiceStateTTL | `time.Duration` |  |
| NumShards | `int` | Synapse settings |
| EnableStats | `bool` |  |

### Presence
Presence represents a user's presence state in a guild.

#### Example Usage

```go
// Create a new Presence
presence := Presence{
    UserID: "example",
    GuildID: "example",
    Status: "example",
    Activities: [],
    ClientStatus: /* value */,
}
```

#### Type Definition

```go
type Presence struct {
    UserID string
    GuildID string
    Status string
    Activities []events.Activity
    ClientStatus events.ClientStatus
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| UserID | `string` |  |
| GuildID | `string` |  |
| Status | `string` |  |
| Activities | `[]events.Activity` |  |
| ClientStatus | `events.ClientStatus` |  |

### PresenceStore
PresenceStore manages presence cache.

#### Example Usage

```go
// Create a new PresenceStore
presencestore := PresenceStore{

}
```

#### Type Definition

```go
type PresenceStore struct {
    **ast.IndexListExpr
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ***ast.IndexListExpr | `**ast.IndexListExpr` |  |

## Methods

### Delete



```go
func (**ast.IndexListExpr) Delete(ctx context.Context, key K) bool
```

**Parameters:**
- `ctx` (context.Context)
- `key` (K)

**Returns:**
- bool

### Get



```go
func (*VoiceStateStore) Get(ctx context.Context, guildID, userID string) (*VoiceState, bool)
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)
- `userID` (string)

**Returns:**
- *VoiceState
- bool

### Set



```go
func (**ast.IndexListExpr) Set(ctx context.Context, key K, value V)
```

**Parameters:**
- `ctx` (context.Context)
- `key` (K)
- `value` (V)

**Returns:**
  None

### RoleStore
RoleStore manages role cache with composite keys.

#### Example Usage

```go
// Create a new RoleStore
rolestore := RoleStore{

}
```

#### Type Definition

```go
type RoleStore struct {
    **ast.IndexListExpr
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ***ast.IndexListExpr | `**ast.IndexListExpr` |  |

## Methods

### Delete



```go
func (*VoiceStateStore) Delete(ctx context.Context, guildID, userID string) bool
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)
- `userID` (string)

**Returns:**
- bool

### Get



```go
func (**ast.IndexListExpr) Get(ctx context.Context, key K) (V, bool)
```

**Parameters:**
- `ctx` (context.Context)
- `key` (K)

**Returns:**
- V
- bool

### Set



```go
func (*VoiceStateStore) Set(ctx context.Context, vs *VoiceState)
```

**Parameters:**
- `ctx` (context.Context)
- `vs` (*VoiceState)

**Returns:**
  None

### State
State manages Discord entity caches with automatic population from gateway events.

#### Example Usage

```go
// Create a new State
state := State{
    Guilds: &GuildStore{}{},
    Channels: &ChannelStore{}{},
    Users: &UserStore{}{},
    Members: &MemberStore{}{},
    Roles: &RoleStore{}{},
    Presences: &PresenceStore{}{},
    VoiceStates: &VoiceStateStore{}{},
}
```

#### Type Definition

```go
type State struct {
    Guilds *GuildStore
    Channels *ChannelStore
    Users *UserStore
    Members *MemberStore
    Roles *RoleStore
    Presences *PresenceStore
    VoiceStates *VoiceStateStore
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Guilds | `*GuildStore` |  |
| Channels | `*ChannelStore` |  |
| Users | `*UserStore` |  |
| Members | `*MemberStore` |  |
| Roles | `*RoleStore` |  |
| Presences | `*PresenceStore` |  |
| VoiceStates | `*VoiceStateStore` |  |

### Constructor Functions

### New

New creates a new State with the given options.

```go
func New(r *rest.REST, opts *Options) *State
```

**Parameters:**
- `r` (*rest.REST)
- `opts` (*Options)

**Returns:**
- *State

## Methods

### Close

Close cleans up the state store.

```go
func (*State) Close() error
```

**Parameters:**
  None

**Returns:**
- error

### CurrentUser

CurrentUser returns the bot's user from the READY event.

```go
func (*State) CurrentUser() *models.User
```

**Parameters:**
  None

**Returns:**
- *models.User

### InvalidateAll

InvalidateAll reinitializes all caches, clearing all data.

```go
func (*State) InvalidateAll()
```

**Parameters:**
  None

**Returns:**
  None

### Ready

Ready returns whether the state has been initialized from the READY event.

```go
func (*State) Ready() bool
```

**Parameters:**
  None

**Returns:**
- bool

### RegisterHandlers

RegisterHandlers sets up event handlers for automatic state population.

```go
func (*State) RegisterHandlers(dispatcher *events.Dispatcher) error
```

**Parameters:**
- `dispatcher` (*events.Dispatcher)

**Returns:**
- error

### Stats

Stats returns statistics for all caches.

```go
func (*State) Stats() StateStats
```

**Parameters:**
  None

**Returns:**
- StateStats

### StateStats
StateStats contains aggregated statistics from all caches.

#### Example Usage

```go
// Create a new StateStats
statestats := StateStats{
    GuildCount: 42,
    ChannelCount: 42,
    UserCount: 42,
    MemberCount: 42,
    RoleCount: 42,
    PresenceCount: 42,
    VoiceStateCount: 42,
}
```

#### Type Definition

```go
type StateStats struct {
    GuildCount int
    ChannelCount int
    UserCount int
    MemberCount int
    RoleCount int
    PresenceCount int
    VoiceStateCount int
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| GuildCount | `int` |  |
| ChannelCount | `int` |  |
| UserCount | `int` |  |
| MemberCount | `int` |  |
| RoleCount | `int` |  |
| PresenceCount | `int` |  |
| VoiceStateCount | `int` |  |

### Store
Store is a generic entity cache with standard CRUD operations.

#### Example Usage

```go
// Create a new Store
store := Store{

}
```

#### Type Definition

```go
type Store struct {
}
```

## Methods

### Count



```go
func (**ast.IndexListExpr) Count() int
```

**Parameters:**
  None

**Returns:**
- int

### Delete



```go
func (**ast.IndexListExpr) Delete(ctx context.Context, key K) bool
```

**Parameters:**
- `ctx` (context.Context)
- `key` (K)

**Returns:**
- bool

### Enabled



```go
func (**ast.IndexListExpr) Enabled() bool
```

**Parameters:**
  None

**Returns:**
- bool

### Get



```go
func (**ast.IndexListExpr) Get(ctx context.Context, key K) (V, bool)
```

**Parameters:**
- `ctx` (context.Context)
- `key` (K)

**Returns:**
- V
- bool

### Set



```go
func (**ast.IndexListExpr) Set(ctx context.Context, key K, value V)
```

**Parameters:**
- `ctx` (context.Context)
- `key` (K)
- `value` (V)

**Returns:**
  None

### UserStore
UserStore manages user cache.

#### Example Usage

```go
// Create a new UserStore
userstore := UserStore{

}
```

#### Type Definition

```go
type UserStore struct {
    **ast.IndexListExpr
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ***ast.IndexListExpr | `**ast.IndexListExpr` |  |

## Methods

### Refresh



```go
func (*MemberStore) Refresh(ctx context.Context, guildID, userID string) (*CachedMember, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)
- `userID` (string)

**Returns:**
- *CachedMember
- error

### VoiceState
VoiceState represents a user's voice connection state in a guild.

#### Example Usage

```go
// Create a new VoiceState
voicestate := VoiceState{
    UserID: "example",
    GuildID: "example",
    ChannelID: &"example"{},
    SessionID: "example",
    Deaf: true,
    Mute: true,
    SelfDeaf: true,
    SelfMute: true,
    SelfStream: true,
    SelfVideo: true,
    Suppress: true,
}
```

#### Type Definition

```go
type VoiceState struct {
    UserID string
    GuildID string
    ChannelID *string
    SessionID string
    Deaf bool
    Mute bool
    SelfDeaf bool
    SelfMute bool
    SelfStream bool
    SelfVideo bool
    Suppress bool
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| UserID | `string` |  |
| GuildID | `string` |  |
| ChannelID | `*string` |  |
| SessionID | `string` |  |
| Deaf | `bool` |  |
| Mute | `bool` |  |
| SelfDeaf | `bool` |  |
| SelfMute | `bool` |  |
| SelfStream | `bool` |  |
| SelfVideo | `bool` |  |
| Suppress | `bool` |  |

### VoiceStateStore
VoiceStateStore manages voice state cache.

#### Example Usage

```go
// Create a new VoiceStateStore
voicestatestore := VoiceStateStore{

}
```

#### Type Definition

```go
type VoiceStateStore struct {
    **ast.IndexListExpr
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ***ast.IndexListExpr | `**ast.IndexListExpr` |  |

## Methods

### Delete



```go
func (**ast.IndexListExpr) Delete(ctx context.Context, key K) bool
```

**Parameters:**
- `ctx` (context.Context)
- `key` (K)

**Returns:**
- bool

### Get



```go
func (*VoiceStateStore) Get(ctx context.Context, guildID, userID string) (*VoiceState, bool)
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)
- `userID` (string)

**Returns:**
- *VoiceState
- bool

### Set



```go
func (**ast.IndexListExpr) Set(ctx context.Context, key K, value V)
```

**Parameters:**
- `ctx` (context.Context)
- `key` (K)
- `value` (V)

**Returns:**
  None

## External Links

- [Package Overview](../packages/state.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/discord/state)
- [Source Code](https://github.com/kolosys/discord/tree/main/state)

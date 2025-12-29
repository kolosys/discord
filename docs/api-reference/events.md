# events API

Complete API documentation for the events package.

**Import Path:** `github.com/kolosys/discord/events`

## Package Documentation



## Types

### Activity
Activity represents a user activity.

#### Example Usage

```go
// Create a new Activity
activity := Activity{
    Name: "example",
    Type: 42,
    URL: &"example"{},
    CreatedAt: 42,
    State: &"example"{},
    Details: &"example"{},
    Emoji: &ReactionEmoji{}{},
}
```

#### Type Definition

```go
type Activity struct {
    Name string `json:"name"`
    Type int `json:"type"`
    URL *string `json:"url,omitempty"`
    CreatedAt int64 `json:"created_at,omitempty"`
    State *string `json:"state,omitempty"`
    Details *string `json:"details,omitempty"`
    Emoji *ReactionEmoji `json:"emoji,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Type | `int` |  |
| URL | `*string` |  |
| CreatedAt | `int64` |  |
| State | `*string` |  |
| Details | `*string` |  |
| Emoji | `*ReactionEmoji` |  |

### Base
Base contains common fields for all events.

#### Example Usage

```go
// Create a new Base
base := Base{
    ShardID: 42,
}
```

#### Type Definition

```go
type Base struct {
    ShardID int `json:"-"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ShardID | `int` | Shard that received this event |

### ChannelCreateEvent
ChannelCreateEvent is dispatched when a channel is created.

#### Example Usage

```go
// Create a new ChannelCreateEvent
channelcreateevent := ChannelCreateEvent{

}
```

#### Type Definition

```go
type ChannelCreateEvent struct {
    Base
    models.GuildChannel
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| *models.GuildChannel | `models.GuildChannel` |  |

### ChannelDeleteEvent
ChannelDeleteEvent is dispatched when a channel is deleted.

#### Example Usage

```go
// Create a new ChannelDeleteEvent
channeldeleteevent := ChannelDeleteEvent{

}
```

#### Type Definition

```go
type ChannelDeleteEvent struct {
    Base
    models.GuildChannel
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| *models.GuildChannel | `models.GuildChannel` |  |

### ChannelUpdateEvent
ChannelUpdateEvent is dispatched when a channel is updated.

#### Example Usage

```go
// Create a new ChannelUpdateEvent
channelupdateevent := ChannelUpdateEvent{

}
```

#### Type Definition

```go
type ChannelUpdateEvent struct {
    Base
    models.GuildChannel
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| *models.GuildChannel | `models.GuildChannel` |  |

### ClientStatus
ClientStatus represents a user's client status across platforms.

#### Example Usage

```go
// Create a new ClientStatus
clientstatus := ClientStatus{
    Desktop: "example",
    Mobile: "example",
    Web: "example",
}
```

#### Type Definition

```go
type ClientStatus struct {
    Desktop string `json:"desktop,omitempty"`
    Mobile string `json:"mobile,omitempty"`
    Web string `json:"web,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Desktop | `string` |  |
| Mobile | `string` |  |
| Web | `string` |  |

### Dispatcher
Dispatcher manages event subscriptions and dispatching.

#### Example Usage

```go
// Create a new Dispatcher
dispatcher := Dispatcher{

}
```

#### Type Definition

```go
type Dispatcher struct {
}
```

### Constructor Functions

### NewDispatcher

NewDispatcher creates a new event dispatcher.

```go
func NewDispatcher(b bus.EventBus) *Dispatcher
```

**Parameters:**
- `b` (bus.EventBus)

**Returns:**
- *Dispatcher

## Methods

### Close

Close unsubscribes all registered handlers.

```go
func (*Dispatcher) Close() error
```

**Parameters:**
  None

**Returns:**
- error

### OnAny

OnAny registers a handler for all events using pattern matching. Context is propagated from the event bus for cancellation and deadline support.

```go
func (*Dispatcher) OnAny(handler RawHandler) error
```

**Parameters:**
- `handler` (RawHandler)

**Returns:**
- error

### OnRaw

OnRaw registers a handler that receives raw event data. Useful for custom event processing or debugging. Context is propagated from the event bus for cancellation and deadline support.

```go
func (*Dispatcher) OnRaw(eventType Type, handler RawHandler) error
```

**Parameters:**
- `eventType` (Type)
- `handler` (RawHandler)

**Returns:**
- error

### GuildBanAddEvent
GuildBanAddEvent is dispatched when a user is banned from a guild.

#### Example Usage

```go
// Create a new GuildBanAddEvent
guildbanaddevent := GuildBanAddEvent{
    GuildID: "example",
    User: &/* value */{},
}
```

#### Type Definition

```go
type GuildBanAddEvent struct {
    Base
    GuildID string `json:"guild_id"`
    User *models.User `json:"user"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| GuildID | `string` |  |
| User | `*models.User` |  |

### GuildBanRemoveEvent
GuildBanRemoveEvent is dispatched when a user is unbanned from a guild.

#### Example Usage

```go
// Create a new GuildBanRemoveEvent
guildbanremoveevent := GuildBanRemoveEvent{
    GuildID: "example",
    User: &/* value */{},
}
```

#### Type Definition

```go
type GuildBanRemoveEvent struct {
    Base
    GuildID string `json:"guild_id"`
    User *models.User `json:"user"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| GuildID | `string` |  |
| User | `*models.User` |  |

### GuildCreateEvent
GuildCreateEvent is dispatched when a guild becomes available or the bot joins a new guild.

#### Example Usage

```go
// Create a new GuildCreateEvent
guildcreateevent := GuildCreateEvent{
    JoinedAt: /* value */,
    Large: true,
    Unavailable: true,
    MemberCount: 42,
    Members: [],
    Channels: [],
    Threads: [],
    Presences: [],
    VoiceStates: [],
}
```

#### Type Definition

```go
type GuildCreateEvent struct {
    Base
    models.Guild
    JoinedAt time.Time `json:"joined_at,omitempty"`
    Large bool `json:"large,omitempty"`
    Unavailable bool `json:"unavailable,omitempty"`
    MemberCount int `json:"member_count,omitempty"`
    Members []models.GuildMember `json:"members,omitempty"`
    Channels []models.GuildChannel `json:"channels,omitempty"`
    Threads []any `json:"threads,omitempty"`
    Presences []any `json:"presences,omitempty"`
    VoiceStates []any `json:"voice_states,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| *models.Guild | `models.Guild` |  |
| JoinedAt | `time.Time` |  |
| Large | `bool` |  |
| Unavailable | `bool` |  |
| MemberCount | `int` |  |
| Members | `[]models.GuildMember` |  |
| Channels | `[]models.GuildChannel` |  |
| Threads | `[]any` |  |
| Presences | `[]any` |  |
| VoiceStates | `[]any` |  |

### GuildDeleteEvent
GuildDeleteEvent is dispatched when a guild becomes unavailable or the bot leaves/is removed.

#### Example Usage

```go
// Create a new GuildDeleteEvent
guilddeleteevent := GuildDeleteEvent{
    ID: "example",
    Unavailable: true,
}
```

#### Type Definition

```go
type GuildDeleteEvent struct {
    Base
    ID string `json:"id"`
    Unavailable bool `json:"unavailable,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| ID | `string` |  |
| Unavailable | `bool` |  |

### GuildMemberAddEvent
GuildMemberAddEvent is dispatched when a user joins a guild.

#### Example Usage

```go
// Create a new GuildMemberAddEvent
guildmemberaddevent := GuildMemberAddEvent{
    GuildID: "example",
}
```

#### Type Definition

```go
type GuildMemberAddEvent struct {
    Base
    models.GuildMember
    GuildID string `json:"guild_id"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| *models.GuildMember | `models.GuildMember` |  |
| GuildID | `string` |  |

### GuildMemberRemoveEvent
GuildMemberRemoveEvent is dispatched when a user leaves or is removed from a guild.

#### Example Usage

```go
// Create a new GuildMemberRemoveEvent
guildmemberremoveevent := GuildMemberRemoveEvent{
    GuildID: "example",
    User: &/* value */{},
}
```

#### Type Definition

```go
type GuildMemberRemoveEvent struct {
    Base
    GuildID string `json:"guild_id"`
    User *models.User `json:"user"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| GuildID | `string` |  |
| User | `*models.User` |  |

### GuildMemberUpdateEvent
GuildMemberUpdateEvent is dispatched when a guild member is updated.

#### Example Usage

```go
// Create a new GuildMemberUpdateEvent
guildmemberupdateevent := GuildMemberUpdateEvent{
    GuildID: "example",
    Roles: [],
    User: &/* value */{},
    Nick: &"example"{},
    JoinedAt: /* value */,
    PremiumSince: &/* value */{},
    Deaf: true,
    Mute: true,
    Pending: true,
}
```

#### Type Definition

```go
type GuildMemberUpdateEvent struct {
    Base
    GuildID string `json:"guild_id"`
    Roles []string `json:"roles"`
    User *models.User `json:"user"`
    Nick *string `json:"nick,omitempty"`
    JoinedAt time.Time `json:"joined_at,omitempty"`
    PremiumSince *time.Time `json:"premium_since,omitempty"`
    Deaf bool `json:"deaf,omitempty"`
    Mute bool `json:"mute,omitempty"`
    Pending bool `json:"pending,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| GuildID | `string` |  |
| Roles | `[]string` |  |
| User | `*models.User` |  |
| Nick | `*string` |  |
| JoinedAt | `time.Time` |  |
| PremiumSince | `*time.Time` |  |
| Deaf | `bool` |  |
| Mute | `bool` |  |
| Pending | `bool` |  |

### GuildMembersChunkEvent
GuildMembersChunkEvent is dispatched in response to RequestGuildMembers.

#### Example Usage

```go
// Create a new GuildMembersChunkEvent
guildmemberschunkevent := GuildMembersChunkEvent{
    GuildID: "example",
    Members: [],
    ChunkIndex: 42,
    ChunkCount: 42,
    NotFound: [],
    Presences: [],
    Nonce: "example",
}
```

#### Type Definition

```go
type GuildMembersChunkEvent struct {
    Base
    GuildID string `json:"guild_id"`
    Members []models.GuildMember `json:"members"`
    ChunkIndex int `json:"chunk_index"`
    ChunkCount int `json:"chunk_count"`
    NotFound []string `json:"not_found,omitempty"`
    Presences []any `json:"presences,omitempty"`
    Nonce string `json:"nonce,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| GuildID | `string` |  |
| Members | `[]models.GuildMember` |  |
| ChunkIndex | `int` |  |
| ChunkCount | `int` |  |
| NotFound | `[]string` |  |
| Presences | `[]any` |  |
| Nonce | `string` |  |

### GuildRoleCreateEvent
GuildRoleCreateEvent is dispatched when a role is created.

#### Example Usage

```go
// Create a new GuildRoleCreateEvent
guildrolecreateevent := GuildRoleCreateEvent{
    GuildID: "example",
    Role: /* value */,
}
```

#### Type Definition

```go
type GuildRoleCreateEvent struct {
    Base
    GuildID string `json:"guild_id"`
    Role models.GuildRole `json:"role"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| GuildID | `string` |  |
| Role | `models.GuildRole` |  |

### GuildRoleDeleteEvent
GuildRoleDeleteEvent is dispatched when a role is deleted.

#### Example Usage

```go
// Create a new GuildRoleDeleteEvent
guildroledeleteevent := GuildRoleDeleteEvent{
    GuildID: "example",
    RoleID: "example",
}
```

#### Type Definition

```go
type GuildRoleDeleteEvent struct {
    Base
    GuildID string `json:"guild_id"`
    RoleID string `json:"role_id"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| GuildID | `string` |  |
| RoleID | `string` |  |

### GuildRoleUpdateEvent
GuildRoleUpdateEvent is dispatched when a role is updated.

#### Example Usage

```go
// Create a new GuildRoleUpdateEvent
guildroleupdateevent := GuildRoleUpdateEvent{
    GuildID: "example",
    Role: /* value */,
}
```

#### Type Definition

```go
type GuildRoleUpdateEvent struct {
    Base
    GuildID string `json:"guild_id"`
    Role models.GuildRole `json:"role"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| GuildID | `string` |  |
| Role | `models.GuildRole` |  |

### GuildUpdateEvent
GuildUpdateEvent is dispatched when a guild is updated.

#### Example Usage

```go
// Create a new GuildUpdateEvent
guildupdateevent := GuildUpdateEvent{

}
```

#### Type Definition

```go
type GuildUpdateEvent struct {
    Base
    models.Guild
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| *models.Guild | `models.Guild` |  |

### Handler
Handler is a type-safe event handler function.

#### Example Usage

```go
// Example usage of Handler
var value Handler
// Initialize with appropriate value
```

#### Type Definition

```go
type Handler func(ctx context.Context, event *T)
```

### InteractionCreateEvent
InteractionCreateEvent is dispatched when a user interacts with the application.

#### Example Usage

```go
// Create a new InteractionCreateEvent
interactioncreateevent := InteractionCreateEvent{

}
```

#### Type Definition

```go
type InteractionCreateEvent struct {
    Base
    models.Interaction
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| *models.Interaction | `models.Interaction` |  |

### InviteCreateEvent
InviteCreateEvent is dispatched when an invite is created.

#### Example Usage

```go
// Create a new InviteCreateEvent
invitecreateevent := InviteCreateEvent{
    ChannelID: "example",
    Code: "example",
    CreatedAt: /* value */,
    GuildID: "example",
    Inviter: &/* value */{},
    MaxAge: 42,
    MaxUses: 42,
    Temporary: true,
    Uses: 42,
}
```

#### Type Definition

```go
type InviteCreateEvent struct {
    Base
    ChannelID string `json:"channel_id"`
    Code string `json:"code"`
    CreatedAt time.Time `json:"created_at"`
    GuildID string `json:"guild_id,omitempty"`
    Inviter *models.User `json:"inviter,omitempty"`
    MaxAge int `json:"max_age"`
    MaxUses int `json:"max_uses"`
    Temporary bool `json:"temporary"`
    Uses int `json:"uses"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| ChannelID | `string` |  |
| Code | `string` |  |
| CreatedAt | `time.Time` |  |
| GuildID | `string` |  |
| Inviter | `*models.User` |  |
| MaxAge | `int` |  |
| MaxUses | `int` |  |
| Temporary | `bool` |  |
| Uses | `int` |  |

### InviteDeleteEvent
InviteDeleteEvent is dispatched when an invite is deleted.

#### Example Usage

```go
// Create a new InviteDeleteEvent
invitedeleteevent := InviteDeleteEvent{
    ChannelID: "example",
    GuildID: "example",
    Code: "example",
}
```

#### Type Definition

```go
type InviteDeleteEvent struct {
    Base
    ChannelID string `json:"channel_id"`
    GuildID string `json:"guild_id,omitempty"`
    Code string `json:"code"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| ChannelID | `string` |  |
| GuildID | `string` |  |
| Code | `string` |  |

### MessageCreateEvent
MessageCreateEvent is dispatched when a message is created.

#### Example Usage

```go
// Create a new MessageCreateEvent
messagecreateevent := MessageCreateEvent{
    GuildID: "example",
    Member: &/* value */{},
}
```

#### Type Definition

```go
type MessageCreateEvent struct {
    Base
    models.Message
    GuildID string `json:"guild_id,omitempty"`
    Member *models.GuildMember `json:"member,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| *models.Message | `models.Message` |  |
| GuildID | `string` |  |
| Member | `*models.GuildMember` |  |

## Methods

### AuthorUser

AuthorUser returns the message author as a typed User, or nil if unavailable.

```go
func (*MessageCreateEvent) AuthorUser() *models.User
```

**Parameters:**
  None

**Returns:**
- *models.User

### IsBot

IsBot returns true if the message was sent by a bot.

```go
func (*MessageCreateEvent) IsBot() bool
```

**Parameters:**
  None

**Returns:**
- bool

### IsHuman

IsHuman returns true if the message was sent by a human user (not a bot or system).

```go
func (*MessageCreateEvent) IsHuman() bool
```

**Parameters:**
  None

**Returns:**
- bool

### IsSystem

IsSystem returns true if the message was sent by the Discord system.

```go
func (*MessageCreateEvent) IsSystem() bool
```

**Parameters:**
  None

**Returns:**
- bool

### MessageDeleteBulkEvent
MessageDeleteBulkEvent is dispatched when multiple messages are deleted at once.

#### Example Usage

```go
// Create a new MessageDeleteBulkEvent
messagedeletebulkevent := MessageDeleteBulkEvent{
    IDs: [],
    ChannelID: "example",
    GuildID: "example",
}
```

#### Type Definition

```go
type MessageDeleteBulkEvent struct {
    Base
    IDs []string `json:"ids"`
    ChannelID string `json:"channel_id"`
    GuildID string `json:"guild_id,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| IDs | `[]string` |  |
| ChannelID | `string` |  |
| GuildID | `string` |  |

### MessageDeleteEvent
MessageDeleteEvent is dispatched when a message is deleted.

#### Example Usage

```go
// Create a new MessageDeleteEvent
messagedeleteevent := MessageDeleteEvent{
    ID: "example",
    ChannelID: "example",
    GuildID: "example",
}
```

#### Type Definition

```go
type MessageDeleteEvent struct {
    Base
    ID string `json:"id"`
    ChannelID string `json:"channel_id"`
    GuildID string `json:"guild_id,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| ID | `string` |  |
| ChannelID | `string` |  |
| GuildID | `string` |  |

### MessageReactionAddEvent
MessageReactionAddEvent is dispatched when a reaction is added to a message.

#### Example Usage

```go
// Create a new MessageReactionAddEvent
messagereactionaddevent := MessageReactionAddEvent{
    UserID: "example",
    ChannelID: "example",
    MessageID: "example",
    GuildID: "example",
    Member: &/* value */{},
    Emoji: ReactionEmoji{},
    MessageAuthorID: "example",
}
```

#### Type Definition

```go
type MessageReactionAddEvent struct {
    Base
    UserID string `json:"user_id"`
    ChannelID string `json:"channel_id"`
    MessageID string `json:"message_id"`
    GuildID string `json:"guild_id,omitempty"`
    Member *models.GuildMember `json:"member,omitempty"`
    Emoji ReactionEmoji `json:"emoji"`
    MessageAuthorID string `json:"message_author_id,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| UserID | `string` |  |
| ChannelID | `string` |  |
| MessageID | `string` |  |
| GuildID | `string` |  |
| Member | `*models.GuildMember` |  |
| Emoji | `ReactionEmoji` |  |
| MessageAuthorID | `string` |  |

### MessageReactionRemoveEvent
MessageReactionRemoveEvent is dispatched when a reaction is removed from a message.

#### Example Usage

```go
// Create a new MessageReactionRemoveEvent
messagereactionremoveevent := MessageReactionRemoveEvent{
    UserID: "example",
    ChannelID: "example",
    MessageID: "example",
    GuildID: "example",
    Emoji: ReactionEmoji{},
}
```

#### Type Definition

```go
type MessageReactionRemoveEvent struct {
    Base
    UserID string `json:"user_id"`
    ChannelID string `json:"channel_id"`
    MessageID string `json:"message_id"`
    GuildID string `json:"guild_id,omitempty"`
    Emoji ReactionEmoji `json:"emoji"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| UserID | `string` |  |
| ChannelID | `string` |  |
| MessageID | `string` |  |
| GuildID | `string` |  |
| Emoji | `ReactionEmoji` |  |

### MessageUpdateEvent
MessageUpdateEvent is dispatched when a message is updated.

#### Example Usage

```go
// Create a new MessageUpdateEvent
messageupdateevent := MessageUpdateEvent{
    ID: "example",
    ChannelID: "example",
    GuildID: "example",
    Content: &"example"{},
    EditedAt: &/* value */{},
    Author: &/* value */{},
    Embeds: [],
    Attachments: [],
}
```

#### Type Definition

```go
type MessageUpdateEvent struct {
    Base
    ID string `json:"id"`
    ChannelID string `json:"channel_id"`
    GuildID string `json:"guild_id,omitempty"`
    Content *string `json:"content,omitempty"`
    EditedAt *time.Time `json:"edited_timestamp,omitempty"`
    Author *models.User `json:"author,omitempty"`
    Embeds []any `json:"embeds,omitempty"`
    Attachments []any `json:"attachments,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| ID | `string` |  |
| ChannelID | `string` |  |
| GuildID | `string` |  |
| Content | `*string` |  |
| EditedAt | `*time.Time` |  |
| Author | `*models.User` | Partial message fields - may not all be present |
| Embeds | `[]any` |  |
| Attachments | `[]any` |  |

### PartialApplication
PartialApplication contains partial application info.

#### Example Usage

```go
// Create a new PartialApplication
partialapplication := PartialApplication{
    ID: "example",
    Flags: 42,
}
```

#### Type Definition

```go
type PartialApplication struct {
    ID string `json:"id"`
    Flags int `json:"flags"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Flags | `int` |  |

### PresenceUpdateEvent
PresenceUpdateEvent is dispatched when a user's presence is updated.

#### Example Usage

```go
// Create a new PresenceUpdateEvent
presenceupdateevent := PresenceUpdateEvent{
    User: &/* value */{},
    GuildID: "example",
    Status: "example",
    Activities: [],
    ClientStatus: ClientStatus{},
}
```

#### Type Definition

```go
type PresenceUpdateEvent struct {
    Base
    User *models.User `json:"user"`
    GuildID string `json:"guild_id"`
    Status string `json:"status"`
    Activities []Activity `json:"activities"`
    ClientStatus ClientStatus `json:"client_status"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| User | `*models.User` |  |
| GuildID | `string` |  |
| Status | `string` |  |
| Activities | `[]Activity` |  |
| ClientStatus | `ClientStatus` |  |

### RawHandler
RawHandler handles raw event data without type conversion.

#### Example Usage

```go
// Example usage of RawHandler
var value RawHandler
// Initialize with appropriate value
```

#### Type Definition

```go
type RawHandler func(ctx context.Context, eventType Type, data json.RawMessage)
```

### ReactionEmoji
ReactionEmoji represents an emoji in a reaction.

#### Example Usage

```go
// Create a new ReactionEmoji
reactionemoji := ReactionEmoji{
    ID: "example",
    Name: "example",
    Animated: true,
}
```

#### Type Definition

```go
type ReactionEmoji struct {
    ID string `json:"id,omitempty"`
    Name string `json:"name"`
    Animated bool `json:"animated,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Name | `string` |  |
| Animated | `bool` |  |

### ReadyEvent
ReadyEvent is dispatched when the client has completed the initial handshake.

#### Example Usage

```go
// Create a new ReadyEvent
readyevent := ReadyEvent{
    Version: 42,
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
type ReadyEvent struct {
    Base
    Version int `json:"v"`
    User *models.User `json:"user"`
    Guilds []UnavailableGuild `json:"guilds"`
    SessionID string `json:"session_id"`
    ResumeGatewayURL string `json:"resume_gateway_url"`
    Shard *[*ast.BasicLit]int `json:"shard,omitempty"`
    Application *PartialApplication `json:"application"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| Version | `int` |  |
| User | `*models.User` |  |
| Guilds | `[]UnavailableGuild` |  |
| SessionID | `string` |  |
| ResumeGatewayURL | `string` |  |
| Shard | `*[*ast.BasicLit]int` |  |
| Application | `*PartialApplication` |  |

### Type
Type represents a Discord gateway event type.

#### Example Usage

```go
// Example usage of Type
var value Type
// Initialize with appropriate value
```

#### Type Definition

```go
type Type string
```

## Methods

### String

String returns the string representation of the event type.

```go
func (Type) String() string
```

**Parameters:**
  None

**Returns:**
- string

### TypingStartEvent
TypingStartEvent is dispatched when a user starts typing.

#### Example Usage

```go
// Create a new TypingStartEvent
typingstartevent := TypingStartEvent{
    ChannelID: "example",
    GuildID: "example",
    UserID: "example",
    Timestamp: 42,
    Member: &/* value */{},
}
```

#### Type Definition

```go
type TypingStartEvent struct {
    Base
    ChannelID string `json:"channel_id"`
    GuildID string `json:"guild_id,omitempty"`
    UserID string `json:"user_id"`
    Timestamp int64 `json:"timestamp"`
    Member *models.GuildMember `json:"member,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| ChannelID | `string` |  |
| GuildID | `string` |  |
| UserID | `string` |  |
| Timestamp | `int64` |  |
| Member | `*models.GuildMember` |  |

### UnavailableGuild
UnavailableGuild represents a guild that is initially unavailable.

#### Example Usage

```go
// Create a new UnavailableGuild
unavailableguild := UnavailableGuild{
    ID: "example",
    Unavailable: true,
}
```

#### Type Definition

```go
type UnavailableGuild struct {
    ID string `json:"id"`
    Unavailable bool `json:"unavailable"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Unavailable | `bool` |  |

### VoiceServerUpdateEvent
VoiceServerUpdateEvent is dispatched when a guild's voice server is updated.

#### Example Usage

```go
// Create a new VoiceServerUpdateEvent
voiceserverupdateevent := VoiceServerUpdateEvent{
    Token: "example",
    GuildID: "example",
    Endpoint: &"example"{},
}
```

#### Type Definition

```go
type VoiceServerUpdateEvent struct {
    Base
    Token string `json:"token"`
    GuildID string `json:"guild_id"`
    Endpoint *string `json:"endpoint"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| Token | `string` |  |
| GuildID | `string` |  |
| Endpoint | `*string` |  |

### VoiceStateUpdateEvent
VoiceStateUpdateEvent is dispatched when a user's voice state changes.

#### Example Usage

```go
// Create a new VoiceStateUpdateEvent
voicestateupdateevent := VoiceStateUpdateEvent{
    GuildID: "example",
    ChannelID: &"example"{},
    UserID: "example",
    Member: &/* value */{},
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
type VoiceStateUpdateEvent struct {
    Base
    GuildID string `json:"guild_id,omitempty"`
    ChannelID *string `json:"channel_id"`
    UserID string `json:"user_id"`
    Member *models.GuildMember `json:"member,omitempty"`
    SessionID string `json:"session_id"`
    Deaf bool `json:"deaf"`
    Mute bool `json:"mute"`
    SelfDeaf bool `json:"self_deaf"`
    SelfMute bool `json:"self_mute"`
    SelfStream bool `json:"self_stream,omitempty"`
    SelfVideo bool `json:"self_video"`
    Suppress bool `json:"suppress"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *Base | `Base` |  |
| GuildID | `string` |  |
| ChannelID | `*string` |  |
| UserID | `string` |  |
| Member | `*models.GuildMember` |  |
| SessionID | `string` |  |
| Deaf | `bool` |  |
| Mute | `bool` |  |
| SelfDeaf | `bool` |  |
| SelfMute | `bool` |  |
| SelfStream | `bool` |  |
| SelfVideo | `bool` |  |
| Suppress | `bool` |  |

## Functions

### On
On registers a type-safe handler for a specific event type. The handler will receive the event data already deserialized into the correct type. Context is propagated from the event bus for cancellation and deadline support.

```go
func On(d *Dispatcher, eventType Type, handler *ast.IndexExpr) error
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `d` | `*Dispatcher` | |
| `eventType` | `Type` | |
| `handler` | `*ast.IndexExpr` | |

**Returns:**
| Type | Description |
|------|-------------|
| `error` | |

**Example:**

```go
// Example usage of On
result := On(/* parameters */)
```

## External Links

- [Package Overview](../packages/events.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/discord/events)
- [Source Code](https://github.com/kolosys/discord/tree/main/events)

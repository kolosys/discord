# discord API

Complete API documentation for the discord package.

**Import Path:** `github.com/kolosys/discord`

## Package Documentation



## Constants

**Version**



```go
const Version = "0.1.0"
```

## Types

### Bot
Bot is a unified Discord bot that combines: - Discord Gateway (real-time events via WebSocket) - Discord REST client (API calls) - Helix HTTP server (webhooks, interactions, admin API) [optional] - State cache (automatic entity caching) [enabled by default]

#### Example Usage

```go
// Create a new Bot
bot := Bot{
    REST: &/* value */{},
    Gateway: &/* value */{},
    Bus: /* value */,
    State: &/* value */{},
}
```

#### Type Definition

```go
type Bot struct {
    REST *rest.REST
    Gateway *gateway.Gateway
    Bus bus.EventBus
    State *state.State
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| REST | `*rest.REST` | Discord REST client |
| Gateway | `*gateway.Gateway` | Discord Gateway client |
| Bus | `bus.EventBus` | Event bus for advanced usage |
| State | `*state.State` | State cache (nil if disabled) |

### Constructor Functions

### New

New creates a new Discord bot. If Options.Server.EnableHTTP is true, the bot will run an embedded HTTP server. The HTTP server defaults to :8080 unless Options.Server.HelixOptions specifies a different address.

```go
func New(opts *Options) (*Bot, error)
```

**Parameters:**
- `opts` (*Options)

**Returns:**
- *Bot
- error

## Methods

### Any

Any registers a route that matches any HTTP method. Panics if HTTP is not enabled (Server.EnableHTTP must be true).

```go
func (*Bot) Any(pattern string, handler http.HandlerFunc)
```

**Parameters:**
- `pattern` (string)
- `handler` (http.HandlerFunc)

**Returns:**
  None

### Commands

Commands returns the command router for registering commands and interactions. The router is initialized during bot creation and ready to use immediately.

```go
func (*Bot) Commands() *commands.Router
```

**Parameters:**
  None

**Returns:**
- *commands.Router

### Component

Component registers a component handler.

```go
func (*Bot) Component(customID string, handler commands.ComponentHandlerFunc)
```

**Parameters:**
- `customID` (string)
- `handler` (commands.ComponentHandlerFunc)

**Returns:**
  None

### ComponentPrefix

ComponentPrefix registers a component handler that matches by prefix.

```go
func (*Bot) ComponentPrefix(prefix string, handler commands.ComponentHandlerFunc)
```

**Parameters:**
- `prefix` (string)
- `handler` (commands.ComponentHandlerFunc)

**Returns:**
  None

### DELETE

DELETE registers a DELETE route on the HTTP server. Panics if HTTP is not enabled (Server.EnableHTTP must be true).

```go
func (*Bot) DELETE(pattern string, handler http.HandlerFunc)
```

**Parameters:**
- `pattern` (string)
- `handler` (http.HandlerFunc)

**Returns:**
  None

### DeleteInteractionResponse

DeleteInteractionResponse deletes the original interaction response.

```go
func (*httpInteractionResponder) DeleteInteractionResponse(ctx context.Context, token string) error
```

**Parameters:**
- `ctx` (context.Context)
- `token` (string)

**Returns:**
- error

### EditInteractionResponse

EditInteractionResponse edits the original interaction response.

```go
func (*httpInteractionResponder) EditInteractionResponse(ctx context.Context, token string, edit *commands.MessageEdit) error
```

**Parameters:**
- `ctx` (context.Context)
- `token` (string)
- `edit` (*commands.MessageEdit)

**Returns:**
- error

### FollowupMessage

FollowupMessage sends a followup message to an interaction.

```go
func (*httpInteractionResponder) FollowupMessage(ctx context.Context, token string, message *commands.MessageCreate) (*models.Message, error)
```

**Parameters:**
- `ctx` (context.Context)
- `token` (string)
- `message` (*commands.MessageCreate)

**Returns:**
- *models.Message
- error

### GET

GET registers a GET route on the HTTP server. Panics if HTTP is not enabled (Server.EnableHTTP must be true).

```go
func (*Bot) GET(pattern string, handler http.HandlerFunc)
```

**Parameters:**
- `pattern` (string)
- `handler` (http.HandlerFunc)

**Returns:**
  None

### Group

Group creates a route group with the specified prefix and optional middleware. Panics if HTTP is not enabled (Server.EnableHTTP must be true).

```go
func (*Bot) Group(prefix string, middleware ...any) *helix.Group
```

**Parameters:**
- `prefix` (string)
- `middleware` (...any)

**Returns:**
- *helix.Group

### HEAD

HEAD registers a HEAD route on the HTTP server. Panics if HTTP is not enabled (Server.EnableHTTP must be true).

```go
func (*Bot) HEAD(pattern string, handler http.HandlerFunc)
```

**Parameters:**
- `pattern` (string)
- `handler` (http.HandlerFunc)

**Returns:**
  None

### Handle

Handle registers a route with a custom HTTP method. Panics if HTTP is not enabled (Server.EnableHTTP must be true).

```go
func (*Bot) Handle(method, pattern string, handler http.HandlerFunc)
```

**Parameters:**
- `method` (string)
- `pattern` (string)
- `handler` (http.HandlerFunc)

**Returns:**
  None

### InteractionEndpoint

InteractionEndpoint returns an HTTP handler for Discord's interactions endpoint. This can be used with the Helix server or any standard HTTP server.

```go
func (*Bot) InteractionEndpoint(publicKey string) http.HandlerFunc
```

**Parameters:**
- `publicKey` (string)

**Returns:**
- http.HandlerFunc

### IsRunning

IsRunning returns whether the bot is running.

```go
func (*Bot) IsRunning() bool
```

**Parameters:**
  None

**Returns:**
- bool

### JoinVoiceChannel

JoinVoiceChannel joins a voice channel.

```go
func (*Bot) JoinVoiceChannel(ctx context.Context, guildID, channelID string) error
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)
- `channelID` (string)

**Returns:**
- error

### LeaveVoiceChannel

LeaveVoiceChannel leaves the voice channel in a guild.

```go
func (*Bot) LeaveVoiceChannel(ctx context.Context, guildID string) error
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)

**Returns:**
- error

### MessageContext

MessageContext is a convenience method to register a message context menu command.

```go
func (*Bot) MessageContext(name string, handler commands.HandlerFunc)
```

**Parameters:**
- `name` (string)
- `handler` (commands.HandlerFunc)

**Returns:**
  None

### OPTIONS

OPTIONS registers an OPTIONS route on the HTTP server. Panics if HTTP is not enabled (Server.EnableHTTP must be true).

```go
func (*Bot) OPTIONS(pattern string, handler http.HandlerFunc)
```

**Parameters:**
- `pattern` (string)
- `handler` (http.HandlerFunc)

**Returns:**
  None

### On

On registers a raw event handler for a specific event type.

```go
func (*Bot) On(eventType events.Type, handler events.RawHandler) error
```

**Parameters:**
- `eventType` (events.Type)
- `handler` (events.RawHandler)

**Returns:**
- error

### OnChannelCreate

OnChannelCreate registers a handler for the ChannelCreate event.

```go
func (*Bot) OnChannelCreate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnChannelDelete

OnChannelDelete registers a handler for the ChannelDelete event.

```go
func (*Bot) OnChannelDelete(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnChannelUpdate

OnChannelUpdate registers a handler for the ChannelUpdate event.

```go
func (*Bot) OnChannelUpdate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnGuildBanAdd

OnGuildBanAdd registers a handler for the GuildBanAdd event.

```go
func (*Bot) OnGuildBanAdd(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnGuildBanRemove

OnGuildBanRemove registers a handler for the GuildBanRemove event.

```go
func (*Bot) OnGuildBanRemove(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnGuildCreate

OnGuildCreate registers a handler for the GuildCreate event.

```go
func (*Bot) OnGuildCreate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnGuildDelete

OnGuildDelete registers a handler for the GuildDelete event.

```go
func (*Bot) OnGuildDelete(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnGuildMemberAdd

OnGuildMemberAdd registers a handler for the GuildMemberAdd event.

```go
func (*Bot) OnGuildMemberAdd(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnGuildMemberRemove

OnGuildMemberRemove registers a handler for the GuildMemberRemove event.

```go
func (*Bot) OnGuildMemberRemove(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnGuildMemberUpdate

OnGuildMemberUpdate registers a handler for the GuildMemberUpdate event.

```go
func (*Bot) OnGuildMemberUpdate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnGuildMembersChunk

OnGuildMembersChunk registers a handler for the GuildMembersChunk event.

```go
func (*Bot) OnGuildMembersChunk(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnGuildRoleCreate

OnGuildRoleCreate registers a handler for the GuildRoleCreate event.

```go
func (*Bot) OnGuildRoleCreate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnGuildRoleDelete

OnGuildRoleDelete registers a handler for the GuildRoleDelete event.

```go
func (*Bot) OnGuildRoleDelete(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnGuildRoleUpdate

OnGuildRoleUpdate registers a handler for the GuildRoleUpdate event.

```go
func (*Bot) OnGuildRoleUpdate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnGuildUpdate

OnGuildUpdate registers a handler for the GuildUpdate event.

```go
func (*Bot) OnGuildUpdate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnInteractionCreate

OnInteractionCreate registers a handler for the InteractionCreate event.

```go
func (*Bot) OnInteractionCreate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnInviteCreate

OnInviteCreate registers a handler for the InviteCreate event.

```go
func (*Bot) OnInviteCreate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnInviteDelete

OnInviteDelete registers a handler for the InviteDelete event.

```go
func (*Bot) OnInviteDelete(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnMessageCreate

OnMessageCreate registers a handler for the MessageCreate event.

```go
func (*Bot) OnMessageCreate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnMessageDelete

OnMessageDelete registers a handler for the MessageDelete event.

```go
func (*Bot) OnMessageDelete(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnMessageDeleteBulk

OnMessageDeleteBulk registers a handler for the MessageDeleteBulk event.

```go
func (*Bot) OnMessageDeleteBulk(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnMessageReactionAdd

OnMessageReactionAdd registers a handler for the MessageReactionAdd event.

```go
func (*Bot) OnMessageReactionAdd(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnMessageReactionRemove

OnMessageReactionRemove registers a handler for the MessageReactionRemove event.

```go
func (*Bot) OnMessageReactionRemove(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnMessageUpdate

OnMessageUpdate registers a handler for the MessageUpdate event.

```go
func (*Bot) OnMessageUpdate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnPresenceUpdate

OnPresenceUpdate registers a handler for the PresenceUpdate event.

```go
func (*Bot) OnPresenceUpdate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnReady

OnReady registers a handler for the Ready event.

```go
func (*Bot) OnReady(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnTypingStart

OnTypingStart registers a handler for the TypingStart event.

```go
func (*Bot) OnTypingStart(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnVoiceServerUpdate

OnVoiceServerUpdate registers a handler for the VoiceServerUpdate event.

```go
func (*Bot) OnVoiceServerUpdate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### OnVoiceStateUpdate

OnVoiceStateUpdate registers a handler for the VoiceStateUpdate event.

```go
func (*Bot) OnVoiceStateUpdate(handler *ast.IndexExpr) error
```

**Parameters:**
- `handler` (*ast.IndexExpr)

**Returns:**
- error

### PATCH

PATCH registers a PATCH route on the HTTP server. Panics if HTTP is not enabled (Server.EnableHTTP must be true).

```go
func (*Bot) PATCH(pattern string, handler http.HandlerFunc)
```

**Parameters:**
- `pattern` (string)
- `handler` (http.HandlerFunc)

**Returns:**
  None

### POST

POST registers a POST route on the HTTP server. Panics if HTTP is not enabled (Server.EnableHTTP must be true).

```go
func (*Bot) POST(pattern string, handler http.HandlerFunc)
```

**Parameters:**
- `pattern` (string)
- `handler` (http.HandlerFunc)

**Returns:**
  None

### PUT

PUT registers a PUT route on the HTTP server. Panics if HTTP is not enabled (Server.EnableHTTP must be true).

```go
func (*Bot) PUT(pattern string, handler http.HandlerFunc)
```

**Parameters:**
- `pattern` (string)
- `handler` (http.HandlerFunc)

**Returns:**
  None

### RegisterCommand

RegisterCommand is a convenience method to register a command.

```go
func (*Bot) RegisterCommand(cmd commands.Command)
```

**Parameters:**
- `cmd` (commands.Command)

**Returns:**
  None

### RequestGuildMembers

RequestGuildMembers requests guild members from Discord. Results are received via GUILD_MEMBERS_CHUNK events.

```go
func (*Bot) RequestGuildMembers(ctx context.Context, guildID string, query string, limit int) error
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)
- `query` (string)
- `limit` (int)

**Returns:**
- error

### RequestGuildMembersByID

RequestGuildMembersByID requests specific guild members by user ID. Results are received via GUILD_MEMBERS_CHUNK events.

```go
func (*Bot) RequestGuildMembersByID(ctx context.Context, guildID string, userIDs []string) error
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)
- `userIDs` ([]string)

**Returns:**
- error

### Resource

Resource creates a resource builder for RESTful CRUD routes. Panics if HTTP is not enabled (Server.EnableHTTP must be true).

```go
func (*Bot) Resource(pattern string) *helix.ResourceBuilder
```

**Parameters:**
- `pattern` (string)

**Returns:**
- *helix.ResourceBuilder

### RespondToInteraction

RespondToInteraction sends a response to an interaction.

```go
func (*httpInteractionResponder) RespondToInteraction(_ context.Context, _, token string, response *commands.InteractionResponse) error
```

**Parameters:**
- `_` (context.Context)
- `_` (string)
- `token` (string)
- `response` (*commands.InteractionResponse)

**Returns:**
- error

### SendEmbed

SendEmbed sends an embed message to a channel.

```go
func (*Bot) SendEmbed(ctx context.Context, channelID string, embed models.RichEmbed) (*models.Message, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelID` (string)
- `embed` (models.RichEmbed)

**Returns:**
- *models.Message
- error

### SendMessage

SendMessage sends a text message to a channel.

```go
func (*Bot) SendMessage(ctx context.Context, channelID, content string) (*models.Message, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelID` (string)
- `content` (string)

**Returns:**
- *models.Message
- error

### SendMessageComplex

SendMessageComplex sends a message with full options control.

```go
func (*Bot) SendMessageComplex(ctx context.Context, channelID string, opts models.MessageCreateOptions) (*models.Message, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelID` (string)
- `opts` (models.MessageCreateOptions)

**Returns:**
- *models.Message
- error

### SendReply

SendReply sends a text message as a reply to another message.

```go
func (*Bot) SendReply(ctx context.Context, channelID, messageID, content string) (*models.Message, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelID` (string)
- `messageID` (string)
- `content` (string)

**Returns:**
- *models.Message
- error

### Server

Server returns the embedded Helix HTTP server (nil if HTTP disabled). For convenience, common HTTP methods are available directly on Bot: bot.GET("/health", handler) bot.POST("/api/users", handler) bot.Group("/api").GET("/users", handler) Use this method when you need direct access to advanced Helix features.

```go
func (*Bot) Server() *helix.Server
```

**Parameters:**
  None

**Returns:**
- *helix.Server

### SetActivity

SetActivity updates the bot's activity (e.g., "Playing a game").

```go
func (*Bot) SetActivity(ctx context.Context, activityType gateway.ActivityType, name string) error
```

**Parameters:**
- `ctx` (context.Context)
- `activityType` (gateway.ActivityType)
- `name` (string)

**Returns:**
- error

### SetApplicationID

SetApplicationID sets the application ID for the bot. This is automatically called during bot startup from the Ready event. Manual calls are rarely needed unless managing application ID externally.

```go
func (*Bot) SetApplicationID(id string)
```

**Parameters:**
- `id` (string)

**Returns:**
  None

### SetPresence

SetPresence updates the bot's full presence (status + activity).

```go
func (*Bot) SetPresence(ctx context.Context, presence *gateway.PresenceUpdate) error
```

**Parameters:**
- `ctx` (context.Context)
- `presence` (*gateway.PresenceUpdate)

**Returns:**
- error

### SetStatus

SetStatus updates the bot's status on all shards.

```go
func (*Bot) SetStatus(ctx context.Context, status string) error
```

**Parameters:**
- `ctx` (context.Context)
- `status` (string)

**Returns:**
- error

### Slash

Slash is a convenience method to register a slash command.

```go
func (*Bot) Slash(name, description string, handler commands.HandlerFunc, options ...commands.Option)
```

**Parameters:**
- `name` (string)
- `description` (string)
- `handler` (commands.HandlerFunc)
- `options` (...commands.Option)

**Returns:**
  None

### Start

Start connects to the Discord gateway and starts the bot. This method blocks until the bot is shut down via signal (SIGINT/SIGTERM). If HTTP is enabled, the Helix HTTP server handles signal management. If HTTP is disabled, this method handles signals directly.

```go
func (*Bot) Start(ctx context.Context) error
```

**Parameters:**
- `ctx` (context.Context)

**Returns:**
- error

### Static

Static serves static files from the specified directory. Panics if HTTP is not enabled (Server.EnableHTTP must be true).

```go
func (*Bot) Static(pattern, root string)
```

**Parameters:**
- `pattern` (string)
- `root` (string)

**Returns:**
  None

### Stop

Stop gracefully stops the bot. If HTTP is enabled, this triggers server shutdown which invokes lifecycle hooks. If HTTP is disabled, this directly stops bot components.

```go
func (*Bot) Stop(ctx context.Context) error
```

**Parameters:**
- `ctx` (context.Context)

**Returns:**
- error

### SyncCommands

SyncCommands syncs commands globally.

```go
func (*Bot) SyncCommands(ctx context.Context, appID string, cmds []commands.ApplicationCommandCreate) error
```

**Parameters:**
- `ctx` (context.Context)
- `appID` (string)
- `cmds` ([]commands.ApplicationCommandCreate)

**Returns:**
- error

### SyncCommandsToDiscord

SyncCommandsToDiscord syncs all registered commands with Discord.

```go
func (*Bot) SyncCommandsToDiscord(ctx context.Context) error
```

**Parameters:**
- `ctx` (context.Context)

**Returns:**
- error

### SyncCommandsToGuild

SyncCommandsToGuild syncs all registered commands to a specific guild.

```go
func (*Bot) SyncCommandsToGuild(ctx context.Context, guildID string) error
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)

**Returns:**
- error

### SyncGuildCommands

SyncGuildCommands syncs commands to a specific guild.

```go
func (*Bot) SyncGuildCommands(ctx context.Context, appID, guildID string, cmds []commands.ApplicationCommandCreate) error
```

**Parameters:**
- `ctx` (context.Context)
- `appID` (string)
- `guildID` (string)
- `cmds` ([]commands.ApplicationCommandCreate)

**Returns:**
- error

### UpdateVoiceState

UpdateVoiceState updates the bot's voice state with full control.

```go
func (*Bot) UpdateVoiceState(ctx context.Context, opts *gateway.VoiceStateUpdateData) error
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (*gateway.VoiceStateUpdateData)

**Returns:**
- error

### Use

Use adds middleware to the HTTP server that applies to all routes. Panics if HTTP is not enabled (Server.EnableHTTP must be true).

```go
func (*Bot) Use(middleware ...any)
```

**Parameters:**
- `middleware` (...any)

**Returns:**
  None

### UserContext

UserContext is a convenience method to register a user context menu command.

```go
func (*Bot) UserContext(name string, handler commands.HandlerFunc)
```

**Parameters:**
- `name` (string)
- `handler` (commands.HandlerFunc)

**Returns:**
  None

### GuildCreateEvent
_No documentation available_

#### Example Usage

```go
// Example usage of GuildCreateEvent
var value GuildCreateEvent
// Initialize with appropriate value
```

#### Type Definition

```go
type GuildCreateEvent events.GuildCreateEvent
```

### InteractionCreateEvent
_No documentation available_

#### Example Usage

```go
// Example usage of InteractionCreateEvent
var value InteractionCreateEvent
// Initialize with appropriate value
```

#### Type Definition

```go
type InteractionCreateEvent events.InteractionCreateEvent
```

### MessageCreateEvent
_No documentation available_

#### Example Usage

```go
// Example usage of MessageCreateEvent
var value MessageCreateEvent
// Initialize with appropriate value
```

#### Type Definition

```go
type MessageCreateEvent events.MessageCreateEvent
```

### Module
// Register HTTP routes api := bot.Group("/api/moderation") api.GET("/warnings/{userID}", m.handleGetWarnings) api.POST("/warnings", m.handleCreateWarning) return nil } See examples/modular for a complete example.

#### Example Usage

```go
// Example implementation of Module
type MyModule struct {
    // Add your fields here
}

func (m MyModule) Register(param1 *Bot) error {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type Module interface {
    Register(bot *Bot) error
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### Options
Options configures the Discord bot.

#### Example Usage

```go
// Create a new Options
options := Options{
    Token: "example",
    Intents: /* value */,
    REST: &/* value */{},
    Server: &/* value */{},
    State: &/* value */{},
    DisableState: true,
}
```

#### Type Definition

```go
type Options struct {
    Token string
    Intents gateway.Intent
    REST *rest.Options
    Server *server.Options
    State *state.Options
    DisableState bool
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Token | `string` | Discord configuration (required) |
| Intents | `gateway.Intent` |  |
| REST | `*rest.Options` | REST client options |
| Server | `*server.Options` | Server configuration (optional, for HTTP server and lifecycle management) |
| State | `*state.Options` | State options (optional, for cache configuration) |
| DisableState | `bool` | Set to true to disable state caching entirely |

### ReadyEvent
Event type aliases for convenience

#### Example Usage

```go
// Example usage of ReadyEvent
var value ReadyEvent
// Initialize with appropriate value
```

#### Type Definition

```go
type ReadyEvent events.ReadyEvent
```

### Snowflake
Exports for developer convenience

#### Example Usage

```go
// Example usage of Snowflake
var value Snowflake
// Initialize with appropriate value
```

#### Type Definition

```go
type Snowflake types.Snowflake
```

## Functions

### Listen
Listen is a convenience helper that registers a typed handler for an event. Example: discord.Listen(bot, events.MessageCreate, func(ctx context.Context, e *events.MessageCreateEvent) { fmt.Printf("Message: %s\n", e.Content) })

```go
func Listen(b *Bot, eventType events.Type, handler *ast.IndexExpr) error
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `b` | `*Bot` | |
| `eventType` | `events.Type` | |
| `handler` | `*ast.IndexExpr` | |

**Returns:**
| Type | Description |
|------|-------------|
| `error` | |

**Example:**

```go
// Example usage of Listen
result := Listen(/* parameters */)
```

### ListenRaw
ListenRaw registers a handler that receives raw JSON data for an event.

```go
func ListenRaw(b *Bot, eventType events.Type, handler func(ctx context.Context, data json.RawMessage)) error
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `b` | `*Bot` | |
| `eventType` | `events.Type` | |
| `handler` | `func(ctx context.Context, data json.RawMessage)` | |

**Returns:**
| Type | Description |
|------|-------------|
| `error` | |

**Example:**

```go
// Example usage of ListenRaw
result := ListenRaw(/* parameters */)
```

### MustRegisterModules
MustRegisterModules is like RegisterModules but panics on error. Useful for application startup where registration errors are fatal. Example: discord.MustRegisterModules(bot, &utils.Module{}, &moderation.Module{}, &leveling.Module{}, )

```go
func MustRegisterModules(bot *Bot, modules ...Module)
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `bot` | `*Bot` | |
| `modules` | `...Module` | |

**Returns:**
None

**Example:**

```go
// Example usage of MustRegisterModules
result := MustRegisterModules(/* parameters */)
```

### OnEvent
OnEvent registers a typed event handler using generics.

```go
func OnEvent(b *Bot, eventType events.Type, handler *ast.IndexExpr) error
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `b` | `*Bot` | |
| `eventType` | `events.Type` | |
| `handler` | `*ast.IndexExpr` | |

**Returns:**
| Type | Description |
|------|-------------|
| `error` | |

**Example:**

```go
// Example usage of OnEvent
result := OnEvent(/* parameters */)
```

### RegisterModules
RegisterModules registers multiple modules with the bot. Modules are registered in the order they are provided. If any module returns an error, registration stops and that error is returned. Example: err := discord.RegisterModules(bot, &utils.Module{}, &moderation.Module{}, &leveling.Module{}, &admin.Module{}, ) if err != nil { return fmt.Errorf("failed to register modules: %w", err) }

```go
func RegisterModules(bot *Bot, modules ...Module) error
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `bot` | `*Bot` | |
| `modules` | `...Module` | |

**Returns:**
| Type | Description |
|------|-------------|
| `error` | |

**Example:**

```go
// Example usage of RegisterModules
result := RegisterModules(/* parameters */)
```

## External Links

- [Package Overview](../packages/discord.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/discord)
- [Source Code](https://github.com/kolosys/discord/tree/main/github.com/kolosys/discord)

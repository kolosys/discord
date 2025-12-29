# rest API

Complete API documentation for the rest package.

**Import Path:** `github.com/kolosys/discord/rest`

## Package Documentation



## Constants

**DefaultBaseURL, DefaultUserAgent, DefaultTimeout**



```go
const DefaultBaseURL = "https://discord.com/api/v10"
const DefaultUserAgent = "DiscordBot (kolosys/discord, v0.1.0)"
const DefaultTimeout = 30 * time.Second
```

## Types

### DiscordRateLimiter
DiscordRateLimiter adapts Ion's MultiTierLimiter to Neuron's RateLimitHandler interface for Discord's bucket-based rate limiting system.

#### Example Usage

```go
// Create a new DiscordRateLimiter
discordratelimiter := DiscordRateLimiter{

}
```

#### Type Definition

```go
type DiscordRateLimiter struct {
}
```

### Constructor Functions

### NewRateLimiter

NewRateLimiter creates a new Discord rate limiter.

```go
func NewRateLimiter(cfg *RateLimiterConfig) *DiscordRateLimiter
```

**Parameters:**
- `cfg` (*RateLimiterConfig)

**Returns:**
- *DiscordRateLimiter

## Methods

### Allow

Allow checks if a request can proceed without blocking.

```go
func (*DiscordRateLimiter) Allow(ctx context.Context, method, endpoint string) bool
```

**Parameters:**
- `ctx` (context.Context)
- `method` (string)
- `endpoint` (string)

**Returns:**
- bool

### Metrics

Metrics returns rate limiting metrics.

```go
func (*DiscordRateLimiter) Metrics() *ratelimit.MultiTierMetrics
```

**Parameters:**
  None

**Returns:**
- *ratelimit.MultiTierMetrics

### Reset

Reset clears all rate limit state.

```go
func (*DiscordRateLimiter) Reset()
```

**Parameters:**
  None

**Returns:**
  None

### UpdateFromHeaders

UpdateFromHeaders updates rate limit state from Discord response headers.

```go
func (*DiscordRateLimiter) UpdateFromHeaders(method, endpoint string, info *neuron.RateLimitInfo) error
```

**Parameters:**
- `method` (string)
- `endpoint` (string)
- `info` (*neuron.RateLimitInfo)

**Returns:**
- error

### Wait

Wait blocks until the request is allowed or context is cancelled.

```go
func (*DiscordRateLimiter) Wait(ctx context.Context, method, endpoint string) error
```

**Parameters:**
- `ctx` (context.Context)
- `method` (string)
- `endpoint` (string)

**Returns:**
- error

### ListMessagesParams
ListMessagesParams configures ListMessages requests.

#### Example Usage

```go
// Create a new ListMessagesParams
listmessagesparams := ListMessagesParams{
    Around: &"example"{},
    Before: &"example"{},
    After: &"example"{},
    Limit: 42,
}
```

#### Type Definition

```go
type ListMessagesParams struct {
    Around *string
    Before *string
    After *string
    Limit int
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Around | `*string` |  |
| Before | `*string` |  |
| After | `*string` |  |
| Limit | `int` |  |

### Options
Options holds REST client configuration.

#### Example Usage

```go
// Create a new Options
options := Options{
    BaseURL: "example",
    UserAgent: "example",
    Timeout: /* value */,
    RateLimiter: &RateLimiterConfig{}{},
    DisableRateLimit: true,
}
```

#### Type Definition

```go
type Options struct {
    BaseURL string
    UserAgent string
    Timeout time.Duration
    RateLimiter *RateLimiterConfig
    DisableRateLimit bool
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| BaseURL | `string` |  |
| UserAgent | `string` |  |
| Timeout | `time.Duration` |  |
| RateLimiter | `*RateLimiterConfig` |  |
| DisableRateLimit | `bool` |  |

### REST
REST provides access to the Discord REST API with rate limiting and authentication.

#### Example Usage

```go
// Create a new REST
rest := REST{

}
```

#### Type Definition

```go
type REST struct {
}
```

### Constructor Functions

### New

New creates a new REST client with the given token and options.

```go
func New(token string, opts *Options) *REST
```

**Parameters:**
- `token` (string)
- `opts` (*Options)

**Returns:**
- *REST

## Methods

### BulkOverwriteGlobalCommands

BulkOverwriteGlobalCommands replaces all global application commands.

```go
func (*REST) BulkOverwriteGlobalCommands(ctx context.Context, appID string, commands any) ([]models.ApplicationCommand, error)
```

**Parameters:**
- `ctx` (context.Context)
- `appID` (string)
- `commands` (any)

**Returns:**
- []models.ApplicationCommand
- error

### BulkOverwriteGuildCommands

BulkOverwriteGuildCommands replaces all guild application commands.

```go
func (*REST) BulkOverwriteGuildCommands(ctx context.Context, appID, guildID string, commands any) ([]models.ApplicationCommand, error)
```

**Parameters:**
- `ctx` (context.Context)
- `appID` (string)
- `guildID` (string)
- `commands` (any)

**Returns:**
- []models.ApplicationCommand
- error

### CreateFollowupMessage

CreateFollowupMessage creates a followup message for an interaction.

```go
func (*REST) CreateFollowupMessage(ctx context.Context, appID, token string, message any) (*models.Message, error)
```

**Parameters:**
- `ctx` (context.Context)
- `appID` (string)
- `token` (string)
- `message` (any)

**Returns:**
- *models.Message
- error

### CreateInteractionResponse

CreateInteractionResponse responds to an interaction.

```go
func (*REST) CreateInteractionResponse(ctx context.Context, interactionID, token string, response any) error
```

**Parameters:**
- `ctx` (context.Context)
- `interactionID` (string)
- `token` (string)
- `response` (any)

**Returns:**
- error

### CreateMessage

CreateMessage sends a message to a channel.

```go
func (*REST) CreateMessage(ctx context.Context, channelID string, msg models.MessageCreateOptions) (*models.Message, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelID` (string)
- `msg` (models.MessageCreateOptions)

**Returns:**
- *models.Message
- error

### DeleteOriginalInteractionResponse

DeleteOriginalInteractionResponse deletes the original interaction response.

```go
func (*REST) DeleteOriginalInteractionResponse(ctx context.Context, appID, token string) error
```

**Parameters:**
- `ctx` (context.Context)
- `appID` (string)
- `token` (string)

**Returns:**
- error

### EditOriginalInteractionResponse

EditOriginalInteractionResponse edits the original interaction response.

```go
func (*REST) EditOriginalInteractionResponse(ctx context.Context, appID, token string, edit *models.IncomingWebhookUpdateOptionsPartial) (*models.Message, error)
```

**Parameters:**
- `ctx` (context.Context)
- `appID` (string)
- `token` (string)
- `edit` (*models.IncomingWebhookUpdateOptionsPartial)

**Returns:**
- *models.Message
- error

### GetBotGateway

GetBotGateway retrieves bot-specific gateway information including sharding.

```go
func (*REST) GetBotGateway(ctx context.Context) (*models.GatewayBot, error)
```

**Parameters:**
- `ctx` (context.Context)

**Returns:**
- *models.GatewayBot
- error

### GetChannel

GetChannel retrieves a channel by ID.

```go
func (*REST) GetChannel(ctx context.Context, channelID string) (any, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelID` (string)

**Returns:**
- any
- error

### GetCurrentUser

GetCurrentUser retrieves the current bot user.

```go
func (*REST) GetCurrentUser(ctx context.Context) (*models.UserPII, error)
```

**Parameters:**
- `ctx` (context.Context)

**Returns:**
- *models.UserPII
- error

### GetGateway

GetGateway retrieves the gateway URL for connecting to the Discord gateway.

```go
func (*REST) GetGateway(ctx context.Context) (*models.Gateway, error)
```

**Parameters:**
- `ctx` (context.Context)

**Returns:**
- *models.Gateway
- error

### GetGuild

GetGuild retrieves a guild by ID.

```go
func (*REST) GetGuild(ctx context.Context, guildID string) (*models.GuildWithCounts, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)

**Returns:**
- *models.GuildWithCounts
- error

### GetGuildMember

GetGuildMember retrieves a guild member by user ID.

```go
func (*REST) GetGuildMember(ctx context.Context, guildID, userID string) (*models.GuildMember, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)
- `userID` (string)

**Returns:**
- *models.GuildMember
- error

### GetMessage

GetMessage retrieves a message by ID.

```go
func (*REST) GetMessage(ctx context.Context, channelID, messageID string) (*models.Message, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelID` (string)
- `messageID` (string)

**Returns:**
- *models.Message
- error

### GetMyApplication

GetMyApplication retrieves the current application's information.

```go
func (*REST) GetMyApplication(ctx context.Context) (*models.PrivateApplication, error)
```

**Parameters:**
- `ctx` (context.Context)

**Returns:**
- *models.PrivateApplication
- error

### GetUser

GetUser retrieves a user by ID.

```go
func (*REST) GetUser(ctx context.Context, userID string) (*models.User, error)
```

**Parameters:**
- `ctx` (context.Context)
- `userID` (string)

**Returns:**
- *models.User
- error

### ListGuildChannels

ListGuildChannels retrieves a list of channels for a guild.

```go
func (*REST) ListGuildChannels(ctx context.Context, guildID string) ([]models.GuildChannel, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)

**Returns:**
- []models.GuildChannel
- error

### ListGuildMembers

ListGuildMembers retrieves members from a guild.

```go
func (*REST) ListGuildMembers(ctx context.Context, guildID string) ([]models.GuildMember, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)

**Returns:**
- []models.GuildMember
- error

### ListMessages

ListMessages retrieves messages from a channel.

```go
func (*REST) ListMessages(ctx context.Context, channelID string, params *ListMessagesParams) ([]models.Message, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelID` (string)
- `params` (*ListMessagesParams)

**Returns:**
- []models.Message
- error

### RateLimiter

RateLimiter returns the rate limiter for this client.

```go
func (*REST) RateLimiter() *DiscordRateLimiter
```

**Parameters:**
  None

**Returns:**
- *DiscordRateLimiter

### Token

Token returns the bot token (masked for security).

```go
func (*REST) Token() string
```

**Parameters:**
  None

**Returns:**
- string

### RateLimiterConfig
RateLimiterConfig configures the Discord rate limiter.

#### Example Usage

```go
// Create a new RateLimiterConfig
ratelimiterconfig := RateLimiterConfig{
    GlobalRate: 42,
    GlobalBurst: 42,
    DefaultRouteRate: 42,
    DefaultRouteBurst: 42,
}
```

#### Type Definition

```go
type RateLimiterConfig struct {
    GlobalRate int
    GlobalBurst int
    DefaultRouteRate int
    DefaultRouteBurst int
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| GlobalRate | `int` | GlobalRate is the overall requests per second limit (Discord: 50/s) |
| GlobalBurst | `int` | GlobalBurst is the burst capacity for the global limiter |
| DefaultRouteRate | `int` | DefaultRouteRate is the default per-route rate |
| DefaultRouteBurst | `int` | DefaultRouteBurst is the default per-route burst |

### Constructor Functions

### DefaultRateLimiterConfig

DefaultRateLimiterConfig returns Discord-appropriate defaults.

```go
func DefaultRateLimiterConfig() *RateLimiterConfig
```

**Parameters:**
  None

**Returns:**
- *RateLimiterConfig

## External Links

- [Package Overview](../packages/rest.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/discord/rest)
- [Source Code](https://github.com/kolosys/discord/tree/main/rest)

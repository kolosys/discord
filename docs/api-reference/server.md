# server API

Complete API documentation for the server package.

**Import Path:** `github.com/kolosys/discord/server`

## Package Documentation



## Variables

**ErrAlreadyRunning, ErrNotRunning, ErrNoCallbacks**



```go
var ErrAlreadyRunning = errors.New("server: manager already running")
var ErrNotRunning = errors.New("server: manager not running")
var ErrNoCallbacks = errors.New("server: connect and stop callbacks required")
```

## Types

### Manager
_No documentation available_

#### Example Usage

```go
// Create a new Manager
manager := Manager{

}
```

#### Type Definition

```go
type Manager struct {
}
```

### Constructor Functions

### New



```go
func New(opts *Options) (*Manager, error)
```

**Parameters:**
- `opts` (*Options)

**Returns:**
- *Manager
- error

## Methods

### HTTPServer



```go
func (*Manager) HTTPServer() *helix.Server
```

**Parameters:**
  None

**Returns:**
- *helix.Server

### IsRunning



```go
func (*Manager) IsRunning() bool
```

**Parameters:**
  None

**Returns:**
- bool

### Start



```go
func (*Manager) Start(ctx context.Context) error
```

**Parameters:**
- `ctx` (context.Context)

**Returns:**
- error

### Stop



```go
func (*Manager) Stop(ctx context.Context) error
```

**Parameters:**
- `ctx` (context.Context)

**Returns:**
- error

### Options
Options configures the Discord bot's HTTP server and lifecycle management. All fields are optional and will use sensible defaults if not specified. For HTTP-disabled bots, only ConnectFn and StopFn are required. For HTTP-enabled bots, the server defaults to listening on :8080.

#### Example Usage

```go
// Create a new Options
options := Options{
    Enabled: true,
    HideBanner: true,
    HelixOptions: &/* value */{},
    Banner: "example",
    ConnectFn: /* value */,
    StopFn: /* value */,
}
```

#### Type Definition

```go
type Options struct {
    Enabled bool
    HideBanner bool
    HelixOptions *helix.Options
    Banner string
    ConnectFn func(context.Context) error
    StopFn func(context.Context) error
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Enabled | `bool` | Enabled enables the embedded Helix HTTP server for webhooks and interactions. If false, the bot runs in gateway-only mode. Defaults to false. |
| HideBanner | `bool` | HideBanner suppresses the startup banner for clean production logs. Defaults to false. |
| HelixOptions | `*helix.Options` | HelixOptions provides fine-grained control over the HTTP server configuration. If EnableHTTP is true and HelixOptions is nil, the server defaults to ":8080". See helix.Options for available configuration options. Only used if EnableHTTP is true. |
| Banner | `string` | Banner is the startup banner text displayed when the bot starts. Set by the bot itself - users should not set this field directly. The banner is not displayed if HideBanner is true. |
| ConnectFn | `func(context.Context) error` | ConnectFn is the callback invoked during bot startup to initialize bot components. Must be set by the caller (typically discord.New sets this to bot.connectBot). Called before the HTTP server starts if EnableHTTP is true. Must not be nil. |
| StopFn | `func(context.Context) error` | StopFn is the callback invoked during graceful shutdown to clean up bot resources. Must be set by the caller (typically discord.New sets this to bot.stopBot). Called before exiting after signal or Stop() call. Must not be nil. |

## External Links

- [Package Overview](../packages/server.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/discord/server)
- [Source Code](https://github.com/kolosys/discord/tree/main/server)

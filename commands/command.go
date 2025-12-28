package commands

// CommandType represents the type of application command.
type CommandType int32

const (
	CommandTypeChatInput CommandType = 1 // Slash command
	CommandTypeUser      CommandType = 2 // User context menu
	CommandTypeMessage   CommandType = 3 // Message context menu
)

// Command is the interface all commands implement.
type Command interface {
	Name() string
	Description() string
	Type() CommandType
	Execute(ctx *Context) error
}

// AutocompleteHandler is an optional interface for commands that support autocomplete.
type AutocompleteHandler interface {
	Autocomplete(ctx *Context, focused string) ([]Choice, error)
}

// MiddlewareProvider is an optional interface for commands that have their own middleware.
type MiddlewareProvider interface {
	Middleware() []Middleware
}

// OptionProvider is an optional interface for commands that have options.
type OptionProvider interface {
	Options() []Option
}

// SubcommandProvider is an optional interface for commands that have subcommands.
type SubcommandProvider interface {
	Subcommands() []Command
}

// Choice represents an autocomplete choice.
type Choice struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

// SlashCommand represents a slash command with typed options.
type SlashCommand struct {
	CommandName        string
	CommandDescription string
	CommandOptions     []Option
	CommandMiddleware  []Middleware
	Handler            HandlerFunc
	AutocompleteFunc   func(ctx *Context, focused string) ([]Choice, error)
}

func (c *SlashCommand) Name() string               { return c.CommandName }
func (c *SlashCommand) Description() string        { return c.CommandDescription }
func (c *SlashCommand) Type() CommandType          { return CommandTypeChatInput }
func (c *SlashCommand) Execute(ctx *Context) error { return c.Handler(ctx) }
func (c *SlashCommand) Options() []Option          { return c.CommandOptions }
func (c *SlashCommand) Middleware() []Middleware   { return c.CommandMiddleware }
func (c *SlashCommand) Autocomplete(ctx *Context, focused string) ([]Choice, error) {
	if c.AutocompleteFunc != nil {
		return c.AutocompleteFunc(ctx, focused)
	}
	return nil, nil
}

// UserCommand represents a user context menu command.
type UserCommand struct {
	CommandName       string
	CommandMiddleware []Middleware
	Handler           HandlerFunc
}

func (c *UserCommand) Name() string               { return c.CommandName }
func (c *UserCommand) Description() string        { return "" } // User commands have no description
func (c *UserCommand) Type() CommandType          { return CommandTypeUser }
func (c *UserCommand) Execute(ctx *Context) error { return c.Handler(ctx) }
func (c *UserCommand) Middleware() []Middleware   { return c.CommandMiddleware }

// MessageCommand represents a message context menu command.
type MessageCommand struct {
	CommandName       string
	CommandMiddleware []Middleware
	Handler           HandlerFunc
}

func (c *MessageCommand) Name() string               { return c.CommandName }
func (c *MessageCommand) Description() string        { return "" } // Message commands have no description
func (c *MessageCommand) Type() CommandType          { return CommandTypeMessage }
func (c *MessageCommand) Execute(ctx *Context) error { return c.Handler(ctx) }
func (c *MessageCommand) Middleware() []Middleware   { return c.CommandMiddleware }

// SubcommandGroup represents a group of subcommands.
type SubcommandGroup struct {
	GroupName        string
	GroupDescription string
	GroupSubcommands []Command
}

func (g *SubcommandGroup) Name() string           { return g.GroupName }
func (g *SubcommandGroup) Description() string    { return g.GroupDescription }
func (g *SubcommandGroup) Type() CommandType      { return CommandTypeChatInput }
func (g *SubcommandGroup) Execute(*Context) error { return nil } // Groups don't execute directly
func (g *SubcommandGroup) Subcommands() []Command { return g.GroupSubcommands }

// HandlerFunc is the signature for command handler functions.
type HandlerFunc func(ctx *Context) error

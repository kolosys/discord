# commands API

Complete API documentation for the commands package.

**Import Path:** `github.com/kolosys/discord/commands`

## Package Documentation



## Types

### AllowedMentions
AllowedMentions configures which mentions are parsed.

#### Example Usage

```go
// Create a new AllowedMentions
allowedmentions := AllowedMentions{
    Parse: [],
    Roles: [],
    Users: [],
    RepliedUser: true,
}
```

#### Type Definition

```go
type AllowedMentions struct {
    Parse []string `json:"parse,omitempty"`
    Roles []string `json:"roles,omitempty"`
    Users []string `json:"users,omitempty"`
    RepliedUser bool `json:"replied_user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Parse | `[]string` |  |
| Roles | `[]string` |  |
| Users | `[]string` |  |
| RepliedUser | `bool` |  |

### ApplicationCommandChoice
ApplicationCommandChoice represents a choice for an option.

#### Example Usage

```go
// Create a new ApplicationCommandChoice
applicationcommandchoice := ApplicationCommandChoice{
    Name: "example",
    Value: any{},
}
```

#### Type Definition

```go
type ApplicationCommandChoice struct {
    Name string `json:"name"`
    Value any `json:"value"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Value | `any` |  |

### ApplicationCommandCreate
ApplicationCommandCreate represents a command to create with Discord.

#### Example Usage

```go
// Create a new ApplicationCommandCreate
applicationcommandcreate := ApplicationCommandCreate{
    Name: "example",
    Description: "example",
    Type: 42,
    Options: [],
    DefaultMemberPermissions: &"example"{},
    DMPermission: &true{},
    Contexts: [],
    IntegrationTypes: [],
    NSFW: true,
}
```

#### Type Definition

```go
type ApplicationCommandCreate struct {
    Name string `json:"name"`
    Description string `json:"description,omitempty"`
    Type int32 `json:"type,omitempty"`
    Options []ApplicationCommandOption `json:"options,omitempty"`
    DefaultMemberPermissions *string `json:"default_member_permissions,omitempty"`
    DMPermission *bool `json:"dm_permission,omitempty"`
    Contexts []int32 `json:"contexts,omitempty"`
    IntegrationTypes []int32 `json:"integration_types,omitempty"`
    NSFW bool `json:"nsfw,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Description | `string` |  |
| Type | `int32` |  |
| Options | `[]ApplicationCommandOption` |  |
| DefaultMemberPermissions | `*string` |  |
| DMPermission | `*bool` |  |
| Contexts | `[]int32` |  |
| IntegrationTypes | `[]int32` |  |
| NSFW | `bool` |  |

### ApplicationCommandOption
ApplicationCommandOption represents an option for an application command.

#### Example Usage

```go
// Create a new ApplicationCommandOption
applicationcommandoption := ApplicationCommandOption{
    Name: "example",
    Description: "example",
    Type: 42,
    Required: true,
    Choices: [],
    Options: [],
    ChannelTypes: [],
    MinValue: &3.14{},
    MaxValue: &3.14{},
    MinLength: &42{},
    MaxLength: &42{},
    Autocomplete: true,
}
```

#### Type Definition

```go
type ApplicationCommandOption struct {
    Name string `json:"name"`
    Description string `json:"description"`
    Type int32 `json:"type"`
    Required bool `json:"required,omitempty"`
    Choices []ApplicationCommandChoice `json:"choices,omitempty"`
    Options []ApplicationCommandOption `json:"options,omitempty"`
    ChannelTypes []int32 `json:"channel_types,omitempty"`
    MinValue *float64 `json:"min_value,omitempty"`
    MaxValue *float64 `json:"max_value,omitempty"`
    MinLength *int32 `json:"min_length,omitempty"`
    MaxLength *int32 `json:"max_length,omitempty"`
    Autocomplete bool `json:"autocomplete,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Description | `string` |  |
| Type | `int32` |  |
| Required | `bool` |  |
| Choices | `[]ApplicationCommandChoice` |  |
| Options | `[]ApplicationCommandOption` |  |
| ChannelTypes | `[]int32` |  |
| MinValue | `*float64` |  |
| MaxValue | `*float64` |  |
| MinLength | `*int32` |  |
| MaxLength | `*int32` |  |
| Autocomplete | `bool` |  |

### AutocompleteHandler
AutocompleteHandler is an optional interface for commands that support autocomplete.

#### Example Usage

```go
// Example implementation of AutocompleteHandler
type MyAutocompleteHandler struct {
    // Add your fields here
}

func (m MyAutocompleteHandler) Autocomplete(param1 *Context, param2 string) []Choice {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type AutocompleteHandler interface {
    Autocomplete(ctx *Context, focused string) ([]Choice, error)
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### Choice
Choice represents an autocomplete choice.

#### Example Usage

```go
// Create a new Choice
choice := Choice{
    Name: "example",
    Value: any{},
}
```

#### Type Definition

```go
type Choice struct {
    Name string `json:"name"`
    Value any `json:"value"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Value | `any` |  |

### Command
Command is the interface all commands implement.

#### Example Usage

```go
// Example implementation of Command
type MyCommand struct {
    // Add your fields here
}

func (m MyCommand) Name() string {
    // Implement your logic here
    return
}

func (m MyCommand) Description() string {
    // Implement your logic here
    return
}

func (m MyCommand) Type() CommandType {
    // Implement your logic here
    return
}

func (m MyCommand) Execute(param1 *Context) error {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type Command interface {
    Name() string
    Description() string
    Type() CommandType
    Execute(ctx *Context) error
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### CommandGroup
CommandGroup represents a group of related commands.

#### Example Usage

```go
// Create a new CommandGroup
commandgroup := CommandGroup{
    Name: "example",
    Description: "example",
    Commands: map[],
    Middleware: [],
}
```

#### Type Definition

```go
type CommandGroup struct {
    Name string
    Description string
    Commands map[string]Command
    Middleware []Middleware
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Description | `string` |  |
| Commands | `map[string]Command` |  |
| Middleware | `[]Middleware` |  |

### CommandSyncer
CommandSyncer is an interface for syncing commands with Discord.

#### Example Usage

```go
// Example implementation of CommandSyncer
type MyCommandSyncer struct {
    // Add your fields here
}

func (m MyCommandSyncer) SyncCommands(param1 context.Context, param2 string, param3 []ApplicationCommandCreate) error {
    // Implement your logic here
    return
}

func (m MyCommandSyncer) SyncGuildCommands(param1 context.Context, param2 string, param3 []ApplicationCommandCreate) error {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type CommandSyncer interface {
    SyncCommands(ctx context.Context, appID string, commands []ApplicationCommandCreate) error
    SyncGuildCommands(ctx context.Context, appID, guildID string, commands []ApplicationCommandCreate) error
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### CommandType
CommandType represents the type of application command.

#### Example Usage

```go
// Example usage of CommandType
var value CommandType
// Initialize with appropriate value
```

#### Type Definition

```go
type CommandType int32
```

### ComponentContext
- Defer() - show loading indicator for longer operations # Token Lifetime The interaction token expires after 15 minutes. After expiration, Followup(), Edit(), and Delete() calls will fail. # Modal Submissions For modal submissions, use ModalValue(customID) to retrieve text input values.

#### Example Usage

```go
// Create a new ComponentContext
componentcontext := ComponentContext{
    InteractionID: "example",
    InteractionToken: "example",
    CustomID: "example",
    ComponentType: /* value */,
    Values: [],
    User: &/* value */{},
    Member: &InteractionMember{}{},
    GuildID: "example",
    ChannelID: "example",
    Locale: "example",
    GuildLocale: "example",
    AppPermissions: /* value */,
    Message: &/* value */{},
    Resolved: &ResolvedData{}{},
    ModalValues: map[],
}
```

#### Type Definition

```go
type ComponentContext struct {
    context.Context
    InteractionID string
    InteractionToken string
    CustomID string
    ComponentType types.ComponentType
    Values []string
    User *models.User
    Member *InteractionMember
    GuildID string
    ChannelID string
    Locale string
    GuildLocale string
    AppPermissions types.Permissions
    Message *models.Message
    Resolved *ResolvedData
    ModalValues map[string]string
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *context.Context | `context.Context` |  |
| InteractionID | `string` | Interaction data |
| InteractionToken | `string` |  |
| CustomID | `string` | Component data |
| ComponentType | `types.ComponentType` |  |
| Values | `[]string` | For select menus |
| User | `*models.User` | User and member data |
| Member | `*InteractionMember` |  |
| GuildID | `string` | Location data |
| ChannelID | `string` |  |
| Locale | `string` | Locale data |
| GuildLocale | `string` |  |
| AppPermissions | `types.Permissions` | AppPermissions contains the bot's permissions in the channel where the interaction occurred. |
| Message | `*models.Message` | The message the component is attached to |
| Resolved | `*ResolvedData` | Resolved data from select menu interactions |
| ModalValues | `map[string]string` | Modal input values (for modal submissions) |

## Methods

### BotHasPermission

BotHasPermission checks if the bot has a specific permission in the current channel.

```go
func (*Context) BotHasPermission(perm types.Permissions) bool
```

**Parameters:**
- `perm` (types.Permissions)

**Returns:**
- bool

### Defer

Defer defers the response (shows loading).

```go
func (*Context) Defer() error
```

**Parameters:**
  None

**Returns:**
- error

### DeferEphemeral

DeferEphemeral defers with an ephemeral response.

```go
func (*ComponentContext) DeferEphemeral() error
```

**Parameters:**
  None

**Returns:**
- error

### DeferUpdate

DeferUpdate defers and acknowledges (no loading state shown).

```go
func (*ComponentContext) DeferUpdate() error
```

**Parameters:**
  None

**Returns:**
- error

### Deferred

Deferred returns true if the interaction has been deferred.

```go
func (*ComponentContext) Deferred() bool
```

**Parameters:**
  None

**Returns:**
- bool

### Delete

Delete deletes the original response.

```go
func (*Context) Delete() error
```

**Parameters:**
  None

**Returns:**
- error

### Edit

Edit edits the deferred response.

```go
func (*ComponentContext) Edit(content string) error
```

**Parameters:**
- `content` (string)

**Returns:**
- error

### EditComplex

EditComplex edits the deferred response with full control.

```go
func (*ComponentContext) EditComplex(edit *MessageEdit) error
```

**Parameters:**
- `edit` (*MessageEdit)

**Returns:**
- error

### Followup

Followup sends a followup message.

```go
func (*ComponentContext) Followup(content string) (*models.Message, error)
```

**Parameters:**
- `content` (string)

**Returns:**
- *models.Message
- error

### FollowupComplex

FollowupComplex sends a complex followup message.

```go
func (*Context) FollowupComplex(msg *MessageCreate) (*models.Message, error)
```

**Parameters:**
- `msg` (*MessageCreate)

**Returns:**
- *models.Message
- error

### FollowupEphemeral

FollowupEphemeral sends an ephemeral followup message.

```go
func (*Context) FollowupEphemeral(content string) (*models.Message, error)
```

**Parameters:**
- `content` (string)

**Returns:**
- *models.Message
- error

### InDM

InDM returns true if the interaction occurred in a DM.

```go
func (*Context) InDM() bool
```

**Parameters:**
  None

**Returns:**
- bool

### InGuild

InGuild returns true if the interaction occurred in a guild.

```go
func (*ComponentContext) InGuild() bool
```

**Parameters:**
  None

**Returns:**
- bool

### Modal

Modal shows a modal dialog.

```go
func (*ComponentContext) Modal(customID, title string, comps ...any) error
```

**Parameters:**
- `customID` (string)
- `title` (string)
- `comps` (...any)

**Returns:**
- error

### ModalValue

ModalValue returns a value from a modal text input.

```go
func (*ComponentContext) ModalValue(customID string) string
```

**Parameters:**
- `customID` (string)

**Returns:**
- string

### Reply

Reply sends a new message response.

```go
func (*ComponentContext) Reply(content string) error
```

**Parameters:**
- `content` (string)

**Returns:**
- error

### ReplyComplex

ReplyComplex sends a complex response.

```go
func (*Context) ReplyComplex(data *InteractionResponseData) error
```

**Parameters:**
- `data` (*InteractionResponseData)

**Returns:**
- error

### ReplyComponents

ReplyComponents sends a message with components.

```go
func (*ComponentContext) ReplyComponents(content string, comps ...any) error
```

**Parameters:**
- `content` (string)
- `comps` (...any)

**Returns:**
- error

### ReplyEmbed

ReplyEmbed sends an embed response.

```go
func (*ComponentContext) ReplyEmbed(embed *Embed) error
```

**Parameters:**
- `embed` (*Embed)

**Returns:**
- error

### ReplyEphemeral

ReplyEphemeral sends an ephemeral response.

```go
func (*ComponentContext) ReplyEphemeral(content string) error
```

**Parameters:**
- `content` (string)

**Returns:**
- error

### ReplyEphemeralComponents

ReplyEphemeralComponents sends an ephemeral message with components.

```go
func (*ComponentContext) ReplyEphemeralComponents(content string, comps ...any) error
```

**Parameters:**
- `content` (string)
- `comps` (...any)

**Returns:**
- error

### Responded

Responded returns true if a response has been sent.

```go
func (*ComponentContext) Responded() bool
```

**Parameters:**
  None

**Returns:**
- bool

### SelectedChannel

SelectedChannel returns the first selected channel from a channel select menu.

```go
func (*ComponentContext) SelectedChannel() *models.GuildChannel
```

**Parameters:**
  None

**Returns:**
- *models.GuildChannel

### SelectedChannels

SelectedChannels returns all selected channels from a channel select menu.

```go
func (*ComponentContext) SelectedChannels() []*models.GuildChannel
```

**Parameters:**
  None

**Returns:**
- []*models.GuildChannel

### SelectedMember

SelectedMember returns the first selected member from a user select menu.

```go
func (*ComponentContext) SelectedMember() *models.GuildMember
```

**Parameters:**
  None

**Returns:**
- *models.GuildMember

### SelectedRole

SelectedRole returns the first selected role from a role select menu.

```go
func (*ComponentContext) SelectedRole() *models.GuildRole
```

**Parameters:**
  None

**Returns:**
- *models.GuildRole

### SelectedRoles

SelectedRoles returns all selected roles from a role select menu.

```go
func (*ComponentContext) SelectedRoles() []*models.GuildRole
```

**Parameters:**
  None

**Returns:**
- []*models.GuildRole

### SelectedUser

SelectedUser returns the first selected user from a user select menu.

```go
func (*ComponentContext) SelectedUser() *models.User
```

**Parameters:**
  None

**Returns:**
- *models.User

### SelectedUsers

SelectedUsers returns all selected users from a user select menu.

```go
func (*ComponentContext) SelectedUsers() []*models.User
```

**Parameters:**
  None

**Returns:**
- []*models.User

### SelectedValue

SelectedValue returns the first selected value (for single-select).

```go
func (*ComponentContext) SelectedValue() string
```

**Parameters:**
  None

**Returns:**
- string

### SetResponder

SetResponder sets the responder for this context.

```go
func (*ComponentContext) SetResponder(r Responder, appID string)
```

**Parameters:**
- `r` (Responder)
- `appID` (string)

**Returns:**
  None

### Update

Update updates the original message (the one with the component).

```go
func (*ComponentContext) Update(content string) error
```

**Parameters:**
- `content` (string)

**Returns:**
- error

### UpdateComplex

UpdateComplex updates the original message with full control.

```go
func (*ComponentContext) UpdateComplex(data *InteractionResponseData) error
```

**Parameters:**
- `data` (*InteractionResponseData)

**Returns:**
- error

### UpdateComponents

UpdateComponents updates the original message with new content and components.

```go
func (*ComponentContext) UpdateComponents(content string, comps ...any) error
```

**Parameters:**
- `content` (string)
- `comps` (...any)

**Returns:**
- error

### UpdateEmbed

UpdateEmbed updates the original message with an embed.

```go
func (*ComponentContext) UpdateEmbed(embed *Embed) error
```

**Parameters:**
- `embed` (*Embed)

**Returns:**
- error

### ComponentHandlerFunc
ComponentHandlerFunc is the signature for component handler functions.

#### Example Usage

```go
// Example usage of ComponentHandlerFunc
var value ComponentHandlerFunc
// Initialize with appropriate value
```

#### Type Definition

```go
type ComponentHandlerFunc func(ctx *ComponentContext) error
```

### ComponentMiddleware
ComponentMiddleware wraps a component handler function.

#### Example Usage

```go
// Example usage of ComponentMiddleware
var value ComponentMiddleware
// Initialize with appropriate value
```

#### Type Definition

```go
type ComponentMiddleware func(next ComponentHandlerFunc) ComponentHandlerFunc
```

### Constructor Functions

### ComponentChain

ComponentChain combines multiple component middleware into one.

```go
func ComponentChain(middleware ...ComponentMiddleware) ComponentMiddleware
```

**Parameters:**
- `middleware` (...ComponentMiddleware)

**Returns:**
- ComponentMiddleware

### ComponentErrorHandler

ComponentErrorHandler wraps errors with a custom error handler for components.

```go
func ComponentErrorHandler(handler func(ctx *ComponentContext, err error) error) ComponentMiddleware
```

**Parameters:**
- `handler` (func(ctx *ComponentContext, err error) error)

**Returns:**
- ComponentMiddleware

### ComponentGuildOnly

ComponentGuildOnly ensures the component interaction is only usable in guilds.

```go
func ComponentGuildOnly() ComponentMiddleware
```

**Parameters:**
  None

**Returns:**
- ComponentMiddleware

### ComponentLogger

ComponentLogger logs component interaction execution.

```go
func ComponentLogger(logFunc func(format string, args ...any)) ComponentMiddleware
```

**Parameters:**
- `logFunc` (func(format string, args ...any))

**Returns:**
- ComponentMiddleware

### ComponentRecover

ComponentRecover is middleware that recovers from panics in component handlers.

```go
func ComponentRecover() ComponentMiddleware
```

**Parameters:**
  None

**Returns:**
- ComponentMiddleware

### Context
if err := ctx.Defer(); err != nil { return err } // Perform long operation (up to 15 minutes) result := performLongOperation() // Edit the deferred response return ctx.Reply(result) }

#### Example Usage

```go
// Create a new Context
context := Context{
    InteractionID: "example",
    InteractionToken: "example",
    InteractionType: InteractionType{},
    CommandName: "example",
    CommandID: "example",
    Options: &OptionMap{}{},
    User: &/* value */{},
    Member: &InteractionMember{}{},
    GuildID: "example",
    ChannelID: "example",
    Locale: "example",
    GuildLocale: "example",
    AppPermissions: /* value */,
    TargetUser: &/* value */{},
    TargetMessage: &/* value */{},
    FocusedOption: "example",
    FocusedValue: "example",
}
```

#### Type Definition

```go
type Context struct {
    context.Context
    InteractionID string
    InteractionToken string
    InteractionType InteractionType
    CommandName string
    CommandID string
    Options *OptionMap
    User *models.User
    Member *InteractionMember
    GuildID string
    ChannelID string
    Locale string
    GuildLocale string
    AppPermissions types.Permissions
    TargetUser *models.User
    TargetMessage *models.Message
    FocusedOption string
    FocusedValue string
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| *context.Context | `context.Context` |  |
| InteractionID | `string` | Interaction data |
| InteractionToken | `string` |  |
| InteractionType | `InteractionType` |  |
| CommandName | `string` | Command data |
| CommandID | `string` |  |
| Options | `*OptionMap` |  |
| User | `*models.User` | User and member data |
| Member | `*InteractionMember` |  |
| GuildID | `string` | Location data |
| ChannelID | `string` |  |
| Locale | `string` | Locale data |
| GuildLocale | `string` |  |
| AppPermissions | `types.Permissions` | AppPermissions contains the bot's permissions in the channel where the interaction occurred. |
| TargetUser | `*models.User` | Target data (for context menu commands) |
| TargetMessage | `*models.Message` |  |
| FocusedOption | `string` | Focused option (for autocomplete) |
| FocusedValue | `string` |  |

## Methods

### Bool

Bool is a convenience method to get a boolean option.

```go
func (*Context) Bool(name string) bool
```

**Parameters:**
- `name` (string)

**Returns:**
- bool

### BotHasPermission

BotHasPermission checks if the bot has a specific permission in the current channel.

```go
func (*ComponentContext) BotHasPermission(perm types.Permissions) bool
```

**Parameters:**
- `perm` (types.Permissions)

**Returns:**
- bool

### Defer

Defer acknowledges the interaction and shows a loading state. This MUST be called within 3 seconds if your handler needs more time. After deferring, use Reply() to send the actual response (within 15 minutes).

```go
func (*ComponentContext) Defer() error
```

**Parameters:**
  None

**Returns:**
- error

### DeferEphemeral

DeferEphemeral acknowledges with an ephemeral loading state. This MUST be called within 3 seconds if your handler needs more time. After deferring, use Reply() to send the actual response (within 15 minutes).

```go
func (*ComponentContext) DeferEphemeral() error
```

**Parameters:**
  None

**Returns:**
- error

### Deferred

Deferred returns true if the interaction has been deferred.

```go
func (*ComponentContext) Deferred() bool
```

**Parameters:**
  None

**Returns:**
- bool

### Delete

Delete deletes the original response.

```go
func (*ComponentContext) Delete() error
```

**Parameters:**
  None

**Returns:**
- error

### Edit

Edit edits the original response.

```go
func (*ComponentContext) Edit(content string) error
```

**Parameters:**
- `content` (string)

**Returns:**
- error

### EditComplex

EditComplex edits the original response with full control.

```go
func (*ComponentContext) EditComplex(edit *MessageEdit) error
```

**Parameters:**
- `edit` (*MessageEdit)

**Returns:**
- error

### Float

Float is a convenience method to get a float option.

```go
func (*OptionMap) Float(name string) float64
```

**Parameters:**
- `name` (string)

**Returns:**
- float64

### Followup

Followup sends a followup message. Note: The interaction token expires after 15 minutes. Followup calls after expiration will fail.

```go
func (*ComponentContext) Followup(content string) (*models.Message, error)
```

**Parameters:**
- `content` (string)

**Returns:**
- *models.Message
- error

### FollowupComplex

FollowupComplex sends a complex followup message.

```go
func (*ComponentContext) FollowupComplex(msg *MessageCreate) (*models.Message, error)
```

**Parameters:**
- `msg` (*MessageCreate)

**Returns:**
- *models.Message
- error

### FollowupEphemeral

FollowupEphemeral sends an ephemeral followup message.

```go
func (*ComponentContext) FollowupEphemeral(content string) (*models.Message, error)
```

**Parameters:**
- `content` (string)

**Returns:**
- *models.Message
- error

### GetAttachment

GetAttachment is a convenience method to get an attachment option.

```go
func (*Context) GetAttachment(name string) *models.Attachment
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.Attachment

### GetChannel

GetChannel is a convenience method to get a channel option.

```go
func (*Context) GetChannel(name string) *models.GuildChannel
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.GuildChannel

### GetMember

GetMember is a convenience method to get a member option.

```go
func (*Context) GetMember(name string) *models.GuildMember
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.GuildMember

### GetRole

GetRole is a convenience method to get a role option.

```go
func (*Context) GetRole(name string) *models.GuildRole
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.GuildRole

### GetUser

GetUser is a convenience method to get a user option.

```go
func (*Context) GetUser(name string) *models.User
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.User

### HasPermission

HasPermission checks if the member has a specific permission.

```go
func (*Context) HasPermission(perm types.Permissions) bool
```

**Parameters:**
- `perm` (types.Permissions)

**Returns:**
- bool

### InDM

InDM returns true if the interaction occurred in a DM.

```go
func (*ComponentContext) InDM() bool
```

**Parameters:**
  None

**Returns:**
- bool

### InGuild

InGuild returns true if the interaction occurred in a guild.

```go
func (*ComponentContext) InGuild() bool
```

**Parameters:**
  None

**Returns:**
- bool

### Int

Int is a convenience method to get an integer option.

```go
func (*Context) Int(name string) int64
```

**Parameters:**
- `name` (string)

**Returns:**
- int64

### MemberPermissions

MemberPermissions returns the member's permissions in the channel. Note: Permissions are provided in the interaction, not on the member object.

```go
func (*Context) MemberPermissions() types.Permissions
```

**Parameters:**
  None

**Returns:**
- types.Permissions

### Modal

Modal shows a modal dialog.

```go
func (*ComponentContext) Modal(customID, title string, comps ...any) error
```

**Parameters:**
- `customID` (string)
- `title` (string)
- `comps` (...any)

**Returns:**
- error

### Reply

Reply sends a response to the interaction.

```go
func (*ComponentContext) Reply(content string) error
```

**Parameters:**
- `content` (string)

**Returns:**
- error

### ReplyComplex

ReplyComplex sends a complex response with full control.

```go
func (*ComponentContext) ReplyComplex(data *InteractionResponseData) error
```

**Parameters:**
- `data` (*InteractionResponseData)

**Returns:**
- error

### ReplyComponents

ReplyComponents sends a message with components (buttons, select menus, etc.).

```go
func (*ComponentContext) ReplyComponents(content string, comps ...any) error
```

**Parameters:**
- `content` (string)
- `comps` (...any)

**Returns:**
- error

### ReplyEmbed

ReplyEmbed sends an embed response.

```go
func (*ComponentContext) ReplyEmbed(embed *Embed) error
```

**Parameters:**
- `embed` (*Embed)

**Returns:**
- error

### ReplyEmbedComponents

ReplyEmbedComponents sends an embed with components.

```go
func (*Context) ReplyEmbedComponents(embed *Embed, components ...any) error
```

**Parameters:**
- `embed` (*Embed)
- `components` (...any)

**Returns:**
- error

### ReplyEmbeds

ReplyEmbeds sends multiple embeds.

```go
func (*Context) ReplyEmbeds(embeds ...*Embed) error
```

**Parameters:**
- `embeds` (...*Embed)

**Returns:**
- error

### ReplyEphemeral

ReplyEphemeral sends an ephemeral response (only visible to the user).

```go
func (*Context) ReplyEphemeral(content string) error
```

**Parameters:**
- `content` (string)

**Returns:**
- error

### ReplyEphemeralComponents

ReplyEphemeralComponents sends an ephemeral message with components.

```go
func (*ComponentContext) ReplyEphemeralComponents(content string, comps ...any) error
```

**Parameters:**
- `content` (string)
- `comps` (...any)

**Returns:**
- error

### RespondAutocomplete

RespondAutocomplete sends autocomplete choices.

```go
func (*Context) RespondAutocomplete(choices []Choice) error
```

**Parameters:**
- `choices` ([]Choice)

**Returns:**
- error

### Responded

Responded returns true if a response has been sent.

```go
func (*ComponentContext) Responded() bool
```

**Parameters:**
  None

**Returns:**
- bool

### SetResponder

SetResponder sets the responder for this context.

```go
func (*Router) SetResponder(responder Responder)
```

**Parameters:**
- `responder` (Responder)

**Returns:**
  None

### String

String is a convenience method to get a string option.

```go
func (*Context) String(name string) string
```

**Parameters:**
- `name` (string)

**Returns:**
- string

### Embed
Embed represents a Discord embed.

#### Example Usage

```go
// Create a new Embed
embed := Embed{
    Title: "example",
    Type: "example",
    Description: "example",
    URL: "example",
    Timestamp: "example",
    Color: 42,
    Footer: &EmbedFooter{}{},
    Image: &EmbedImage{}{},
    Thumbnail: &EmbedThumbnail{}{},
    Video: &EmbedVideo{}{},
    Provider: &EmbedProvider{}{},
    Author: &EmbedAuthor{}{},
    Fields: [],
}
```

#### Type Definition

```go
type Embed struct {
    Title string `json:"title,omitempty"`
    Type string `json:"type,omitempty"`
    Description string `json:"description,omitempty"`
    URL string `json:"url,omitempty"`
    Timestamp string `json:"timestamp,omitempty"`
    Color int `json:"color,omitempty"`
    Footer *EmbedFooter `json:"footer,omitempty"`
    Image *EmbedImage `json:"image,omitempty"`
    Thumbnail *EmbedThumbnail `json:"thumbnail,omitempty"`
    Video *EmbedVideo `json:"video,omitempty"`
    Provider *EmbedProvider `json:"provider,omitempty"`
    Author *EmbedAuthor `json:"author,omitempty"`
    Fields []*EmbedField `json:"fields,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Title | `string` |  |
| Type | `string` |  |
| Description | `string` |  |
| URL | `string` |  |
| Timestamp | `string` |  |
| Color | `int` |  |
| Footer | `*EmbedFooter` |  |
| Image | `*EmbedImage` |  |
| Thumbnail | `*EmbedThumbnail` |  |
| Video | `*EmbedVideo` |  |
| Provider | `*EmbedProvider` |  |
| Author | `*EmbedAuthor` |  |
| Fields | `[]*EmbedField` |  |

### EmbedAuthor
_No documentation available_

#### Example Usage

```go
// Create a new EmbedAuthor
embedauthor := EmbedAuthor{
    Name: "example",
    URL: "example",
    IconURL: "example",
}
```

#### Type Definition

```go
type EmbedAuthor struct {
    Name string `json:"name"`
    URL string `json:"url,omitempty"`
    IconURL string `json:"icon_url,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| URL | `string` |  |
| IconURL | `string` |  |

### EmbedField
_No documentation available_

#### Example Usage

```go
// Create a new EmbedField
embedfield := EmbedField{
    Name: "example",
    Value: "example",
    Inline: true,
}
```

#### Type Definition

```go
type EmbedField struct {
    Name string `json:"name"`
    Value string `json:"value"`
    Inline bool `json:"inline,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Value | `string` |  |
| Inline | `bool` |  |

### EmbedFooter
_No documentation available_

#### Example Usage

```go
// Create a new EmbedFooter
embedfooter := EmbedFooter{
    Text: "example",
    IconURL: "example",
}
```

#### Type Definition

```go
type EmbedFooter struct {
    Text string `json:"text"`
    IconURL string `json:"icon_url,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Text | `string` |  |
| IconURL | `string` |  |

### EmbedImage
_No documentation available_

#### Example Usage

```go
// Create a new EmbedImage
embedimage := EmbedImage{
    URL: "example",
}
```

#### Type Definition

```go
type EmbedImage struct {
    URL string `json:"url"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| URL | `string` |  |

### EmbedProvider
_No documentation available_

#### Example Usage

```go
// Create a new EmbedProvider
embedprovider := EmbedProvider{
    Name: "example",
    URL: "example",
}
```

#### Type Definition

```go
type EmbedProvider struct {
    Name string `json:"name,omitempty"`
    URL string `json:"url,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| URL | `string` |  |

### EmbedThumbnail
_No documentation available_

#### Example Usage

```go
// Create a new EmbedThumbnail
embedthumbnail := EmbedThumbnail{
    URL: "example",
}
```

#### Type Definition

```go
type EmbedThumbnail struct {
    URL string `json:"url"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| URL | `string` |  |

### EmbedVideo
_No documentation available_

#### Example Usage

```go
// Create a new EmbedVideo
embedvideo := EmbedVideo{
    URL: "example",
}
```

#### Type Definition

```go
type EmbedVideo struct {
    URL string `json:"url,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| URL | `string` |  |

### HandlerFunc
HandlerFunc is the signature for command handler functions.

#### Example Usage

```go
// Example usage of HandlerFunc
var value HandlerFunc
// Initialize with appropriate value
```

#### Type Definition

```go
type HandlerFunc func(ctx *Context) error
```

### Interaction
Interaction represents a Discord interaction.

#### Example Usage

```go
// Create a new Interaction
interaction := Interaction{
    ID: "example",
    ApplicationID: "example",
    Type: InteractionType{},
    Data: &InteractionData{}{},
    GuildID: "example",
    ChannelID: "example",
    Member: &InteractionMember{}{},
    User: &/* value */{},
    Token: "example",
    Version: 42,
    Message: &/* value */{},
    AppPermissions: "example",
    Locale: "example",
    GuildLocale: "example",
}
```

#### Type Definition

```go
type Interaction struct {
    ID string `json:"id"`
    ApplicationID string `json:"application_id"`
    Type InteractionType `json:"type"`
    Data *InteractionData `json:"data,omitempty"`
    GuildID string `json:"guild_id,omitempty"`
    ChannelID string `json:"channel_id,omitempty"`
    Member *InteractionMember `json:"member,omitempty"`
    User *models.User `json:"user,omitempty"`
    Token string `json:"token"`
    Version int `json:"version"`
    Message *models.Message `json:"message,omitempty"`
    AppPermissions string `json:"app_permissions,omitempty"`
    Locale string `json:"locale,omitempty"`
    GuildLocale string `json:"guild_locale,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ApplicationID | `string` |  |
| Type | `InteractionType` |  |
| Data | `*InteractionData` |  |
| GuildID | `string` |  |
| ChannelID | `string` |  |
| Member | `*InteractionMember` |  |
| User | `*models.User` |  |
| Token | `string` |  |
| Version | `int` |  |
| Message | `*models.Message` |  |
| AppPermissions | `string` |  |
| Locale | `string` |  |
| GuildLocale | `string` |  |

### InteractionCallbackType
InteractionCallbackType represents the type of interaction response.

#### Example Usage

```go
// Example usage of InteractionCallbackType
var value InteractionCallbackType
// Initialize with appropriate value
```

#### Type Definition

```go
type InteractionCallbackType int32
```

### InteractionData
InteractionData contains the data for an interaction.

#### Example Usage

```go
// Create a new InteractionData
interactiondata := InteractionData{
    ID: "example",
    Name: "example",
    Type: 42,
    Resolved: &ResolvedInteractionData{}{},
    Options: [],
    GuildID: "example",
    TargetID: "example",
    CustomID: "example",
    ComponentType: 42,
    Values: [],
    Components: [],
}
```

#### Type Definition

```go
type InteractionData struct {
    ID string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
    Type int32 `json:"type,omitempty"`
    Resolved *ResolvedInteractionData `json:"resolved,omitempty"`
    Options []InteractionOption `json:"options,omitempty"`
    GuildID string `json:"guild_id,omitempty"`
    TargetID string `json:"target_id,omitempty"`
    CustomID string `json:"custom_id,omitempty"`
    ComponentType int32 `json:"component_type,omitempty"`
    Values []string `json:"values,omitempty"`
    Components []json.RawMessage `json:"components,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Name | `string` |  |
| Type | `int32` |  |
| Resolved | `*ResolvedInteractionData` |  |
| Options | `[]InteractionOption` |  |
| GuildID | `string` |  |
| TargetID | `string` |  |
| CustomID | `string` |  |
| ComponentType | `int32` |  |
| Values | `[]string` |  |
| Components | `[]json.RawMessage` |  |

### InteractionMember
InteractionMember represents a guild member in an interaction context. This differs from models.GuildMember as it includes computed permissions.

#### Example Usage

```go
// Create a new InteractionMember
interactionmember := InteractionMember{
    Avatar: &"example"{},
    Banner: &"example"{},
    CommunicationDisabledUntil: &"example"{},
    Deaf: true,
    Flags: 42,
    JoinedAt: "example",
    Mute: true,
    Nick: &"example"{},
    Pending: true,
    PremiumSince: &"example"{},
    Roles: [],
    User: any{},
    Permissions: "example",
}
```

#### Type Definition

```go
type InteractionMember struct {
    Avatar *string `json:"avatar,omitempty"`
    Banner *string `json:"banner,omitempty"`
    CommunicationDisabledUntil *string `json:"communication_disabled_until,omitempty"`
    Deaf bool `json:"deaf"`
    Flags int32 `json:"flags"`
    JoinedAt string `json:"joined_at"`
    Mute bool `json:"mute"`
    Nick *string `json:"nick,omitempty"`
    Pending bool `json:"pending"`
    PremiumSince *string `json:"premium_since,omitempty"`
    Roles []string `json:"roles"`
    User any `json:"user"`
    Permissions string `json:"permissions,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Avatar | `*string` |  |
| Banner | `*string` |  |
| CommunicationDisabledUntil | `*string` |  |
| Deaf | `bool` |  |
| Flags | `int32` |  |
| JoinedAt | `string` |  |
| Mute | `bool` |  |
| Nick | `*string` |  |
| Pending | `bool` |  |
| PremiumSince | `*string` |  |
| Roles | `[]string` |  |
| User | `any` |  |
| Permissions | `string` |  |

### InteractionOption
InteractionOption represents an option in an interaction.

#### Example Usage

```go
// Create a new InteractionOption
interactionoption := InteractionOption{
    Name: "example",
    Type: 42,
    Value: any{},
    Options: [],
    Focused: true,
}
```

#### Type Definition

```go
type InteractionOption struct {
    Name string `json:"name"`
    Type int32 `json:"type"`
    Value any `json:"value,omitempty"`
    Options []InteractionOption `json:"options,omitempty"`
    Focused bool `json:"focused,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Type | `int32` |  |
| Value | `any` |  |
| Options | `[]InteractionOption` |  |
| Focused | `bool` |  |

### InteractionResponse
InteractionResponse represents a response to an interaction.

#### Example Usage

```go
// Create a new InteractionResponse
interactionresponse := InteractionResponse{
    Type: InteractionCallbackType{},
    Data: &InteractionResponseData{}{},
}
```

#### Type Definition

```go
type InteractionResponse struct {
    Type InteractionCallbackType `json:"type"`
    Data *InteractionResponseData `json:"data,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Type | `InteractionCallbackType` |  |
| Data | `*InteractionResponseData` |  |

### InteractionResponseData
InteractionResponseData contains the data for an interaction response.

#### Example Usage

```go
// Create a new InteractionResponseData
interactionresponsedata := InteractionResponseData{
    TTS: true,
    Content: "example",
    Embeds: [],
    AllowedMentions: &AllowedMentions{}{},
    Flags: MessageFlags{},
    Components: [],
    Attachments: [],
    Choices: [],
    CustomID: "example",
    Title: "example",
}
```

#### Type Definition

```go
type InteractionResponseData struct {
    TTS bool `json:"tts,omitempty"`
    Content string `json:"content,omitempty"`
    Embeds []*Embed `json:"embeds,omitempty"`
    AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
    Flags MessageFlags `json:"flags,omitempty"`
    Components []any `json:"components,omitempty"`
    Attachments []*models.Attachment `json:"attachments,omitempty"`
    Choices []Choice `json:"choices,omitempty"`
    CustomID string `json:"custom_id,omitempty"`
    Title string `json:"title,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| TTS | `bool` |  |
| Content | `string` |  |
| Embeds | `[]*Embed` |  |
| AllowedMentions | `*AllowedMentions` |  |
| Flags | `MessageFlags` |  |
| Components | `[]any` |  |
| Attachments | `[]*models.Attachment` |  |
| Choices | `[]Choice` | For autocomplete |
| CustomID | `string` | For modals |
| Title | `string` | For modals |

### InteractionType
InteractionType represents the type of interaction.

#### Example Usage

```go
// Example usage of InteractionType
var value InteractionType
// Initialize with appropriate value
```

#### Type Definition

```go
type InteractionType int32
```

### MessageCommand
MessageCommand represents a message context menu command.

#### Example Usage

```go
// Create a new MessageCommand
messagecommand := MessageCommand{
    CommandName: "example",
    CommandMiddleware: [],
    Handler: HandlerFunc{},
}
```

#### Type Definition

```go
type MessageCommand struct {
    CommandName string
    CommandMiddleware []Middleware
    Handler HandlerFunc
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| CommandName | `string` |  |
| CommandMiddleware | `[]Middleware` |  |
| Handler | `HandlerFunc` |  |

## Methods

### Description



```go
func (*SubcommandGroup) Description() string
```

**Parameters:**
  None

**Returns:**
- string

### Execute



```go
func (*SubcommandGroup) Execute(*Context) error
```

**Parameters:**
- `` (*Context)

**Returns:**
- error

### Middleware



```go
func (*MessageCommand) Middleware() []Middleware
```

**Parameters:**
  None

**Returns:**
- []Middleware

### Name



```go
func (*SubcommandGroup) Name() string
```

**Parameters:**
  None

**Returns:**
- string

### Type



```go
func (*SubcommandGroup) Type() CommandType
```

**Parameters:**
  None

**Returns:**
- CommandType

### MessageCreate
MessageCreate represents a new message to send.

#### Example Usage

```go
// Create a new MessageCreate
messagecreate := MessageCreate{
    Content: "example",
    TTS: true,
    Embeds: [],
    AllowedMentions: &AllowedMentions{}{},
    Flags: MessageFlags{},
    Components: [],
    Attachments: [],
}
```

#### Type Definition

```go
type MessageCreate struct {
    Content string `json:"content,omitempty"`
    TTS bool `json:"tts,omitempty"`
    Embeds []*Embed `json:"embeds,omitempty"`
    AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
    Flags MessageFlags `json:"flags,omitempty"`
    Components []any `json:"components,omitempty"`
    Attachments []*models.Attachment `json:"attachments,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Content | `string` |  |
| TTS | `bool` |  |
| Embeds | `[]*Embed` |  |
| AllowedMentions | `*AllowedMentions` |  |
| Flags | `MessageFlags` |  |
| Components | `[]any` |  |
| Attachments | `[]*models.Attachment` |  |

### MessageEdit
MessageEdit represents edits to a message.

#### Example Usage

```go
// Create a new MessageEdit
messageedit := MessageEdit{
    Content: &"example"{},
    Embeds: [],
    AllowedMentions: &AllowedMentions{}{},
    Components: [],
    Attachments: [],
}
```

#### Type Definition

```go
type MessageEdit struct {
    Content *string `json:"content,omitempty"`
    Embeds []*Embed `json:"embeds,omitempty"`
    AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
    Components []any `json:"components,omitempty"`
    Attachments []*models.Attachment `json:"attachments,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Content | `*string` |  |
| Embeds | `[]*Embed` |  |
| AllowedMentions | `*AllowedMentions` |  |
| Components | `[]any` |  |
| Attachments | `[]*models.Attachment` |  |

### MessageFlags
MessageFlags represents message flags.

#### Example Usage

```go
// Example usage of MessageFlags
var value MessageFlags
// Initialize with appropriate value
```

#### Type Definition

```go
type MessageFlags int32
```

### Middleware
Middleware wraps a handler function to provide cross-cutting concerns.

#### Example Usage

```go
// Example usage of Middleware
var value Middleware
// Initialize with appropriate value
```

#### Type Definition

```go
type Middleware func(next HandlerFunc) HandlerFunc
```

### Constructor Functions

### Chain

Chain combines multiple middleware into one.

```go
func Chain(middleware ...Middleware) Middleware
```

**Parameters:**
- `middleware` (...Middleware)

**Returns:**
- Middleware

### Cooldown

Cooldown applies a simple cooldown per user.

```go
func Cooldown(duration time.Duration) Middleware
```

**Parameters:**
- `duration` (time.Duration)

**Returns:**
- Middleware

### DMOnly

DMOnly ensures the command is only usable in DMs.

```go
func DMOnly() Middleware
```

**Parameters:**
  None

**Returns:**
- Middleware

### DeferredEphemeral

DeferredEphemeral automatically defers with ephemeral response.

```go
func DeferredEphemeral() Middleware
```

**Parameters:**
  None

**Returns:**
- Middleware

### DeferredResponse

DeferredResponse automatically defers the response.

```go
func DeferredResponse() Middleware
```

**Parameters:**
  None

**Returns:**
- Middleware

### ErrorHandler

ErrorHandler wraps errors with a custom error handler.

```go
func ErrorHandler(handler func(ctx *Context, err error) error) Middleware
```

**Parameters:**
- `handler` (func(ctx *Context, err error) error)

**Returns:**
- Middleware

### GuildOnly

GuildOnly ensures the command is only usable in guilds.

```go
func GuildOnly() Middleware
```

**Parameters:**
  None

**Returns:**
- Middleware

### Logger

Logger logs command execution.

```go
func Logger(logFunc func(format string, args ...any)) Middleware
```

**Parameters:**
- `logFunc` (func(format string, args ...any))

**Returns:**
- Middleware

### NSFW

NSFW ensures the command is used in an NSFW channel.

```go
func NSFW() Middleware
```

**Parameters:**
  None

**Returns:**
- Middleware

### RateLimit

RateLimit applies rate limiting per user.

```go
func RateLimit(limit int, window time.Duration) Middleware
```

**Parameters:**
- `limit` (int)
- `window` (time.Duration)

**Returns:**
- Middleware

### Recover

Recover is middleware that recovers from panics.

```go
func Recover() Middleware
```

**Parameters:**
  None

**Returns:**
- Middleware

### RequireOwner

RequireOwner ensures only the bot owner can use the command.

```go
func RequireOwner(ownerIDs ...string) Middleware
```

**Parameters:**
- `ownerIDs` (...string)

**Returns:**
- Middleware

### RequirePermissions

RequirePermissions ensures the user has the required permissions.

```go
func RequirePermissions(perms ...types.Permissions) Middleware
```

**Parameters:**
- `perms` (...types.Permissions)

**Returns:**
- Middleware

### RequireRoles

RequireRoles ensures the user has one of the required roles.

```go
func RequireRoles(roleIDs ...string) Middleware
```

**Parameters:**
- `roleIDs` (...string)

**Returns:**
- Middleware

### MiddlewareProvider
MiddlewareProvider is an optional interface for commands that have their own middleware.

#### Example Usage

```go
// Example implementation of MiddlewareProvider
type MyMiddlewareProvider struct {
    // Add your fields here
}

func (m MyMiddlewareProvider) Middleware() []Middleware {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type MiddlewareProvider interface {
    Middleware() []Middleware
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### Option
Option represents a command option.

#### Example Usage

```go
// Create a new Option
option := Option{
    Name: "example",
    Description: "example",
    Type: OptionType{},
    Required: true,
    Choices: [],
    Options: [],
    ChannelTypes: [],
    MinValue: &3.14{},
    MaxValue: &3.14{},
    MinLength: &42{},
    MaxLength: &42{},
    Autocomplete: true,
}
```

#### Type Definition

```go
type Option struct {
    Name string
    Description string
    Type OptionType
    Required bool
    Choices []Choice
    Options []Option
    ChannelTypes []int32
    MinValue *float64
    MaxValue *float64
    MinLength *int32
    MaxLength *int32
    Autocomplete bool
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Description | `string` |  |
| Type | `OptionType` |  |
| Required | `bool` |  |
| Choices | `[]Choice` |  |
| Options | `[]Option` | For subcommands |
| ChannelTypes | `[]int32` |  |
| MinValue | `*float64` |  |
| MaxValue | `*float64` |  |
| MinLength | `*int32` |  |
| MaxLength | `*int32` |  |
| Autocomplete | `bool` |  |

### OptionBuilder
OptionBuilder builds command options with a fluent API.

#### Example Usage

```go
// Create a new OptionBuilder
optionbuilder := OptionBuilder{

}
```

#### Type Definition

```go
type OptionBuilder struct {
}
```

### Constructor Functions

### Attachment

Attachment creates an attachment option.

```go
func (*OptionMap) Attachment(name string) *models.Attachment
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.Attachment

### Boolean

Boolean creates a boolean option.

```go
func Boolean(name, description string) *OptionBuilder
```

**Parameters:**
- `name` (string)
- `description` (string)

**Returns:**
- *OptionBuilder

### Channel

Channel creates a channel option.

```go
func (*OptionMap) Channel(name string) *models.GuildChannel
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.GuildChannel

### Integer

Integer creates an integer option.

```go
func Integer(name, description string) *OptionBuilder
```

**Parameters:**
- `name` (string)
- `description` (string)

**Returns:**
- *OptionBuilder

### Mentionable

Mentionable creates a mentionable option.

```go
func Mentionable(name, description string) *OptionBuilder
```

**Parameters:**
- `name` (string)
- `description` (string)

**Returns:**
- *OptionBuilder

### Number

Number creates a number (float) option.

```go
func Number(name, description string) *OptionBuilder
```

**Parameters:**
- `name` (string)
- `description` (string)

**Returns:**
- *OptionBuilder

### Role

Role creates a role option.

```go
func (*OptionMap) Role(name string) *models.GuildRole
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.GuildRole

### String

String creates a string option.

```go
func (*OptionMap) String(name string) string
```

**Parameters:**
- `name` (string)

**Returns:**
- string

### User

User creates a user option.

```go
func (*OptionMap) User(name string) *models.User
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.User

## Methods

### Autocomplete



```go
func (*SlashCommand) Autocomplete(ctx *Context, focused string) ([]Choice, error)
```

**Parameters:**
- `ctx` (*Context)
- `focused` (string)

**Returns:**
- []Choice
- error

### Build



```go
func (*OptionBuilder) Build() Option
```

**Parameters:**
  None

**Returns:**
- Option

### ChannelTypes



```go
func (*OptionBuilder) ChannelTypes(types ...int32) *OptionBuilder
```

**Parameters:**
- `types` (...int32)

**Returns:**
- *OptionBuilder

### Choices



```go
func (*OptionBuilder) Choices(choices ...Choice) *OptionBuilder
```

**Parameters:**
- `choices` (...Choice)

**Returns:**
- *OptionBuilder

### MaxLength



```go
func (*OptionBuilder) MaxLength(max int32) *OptionBuilder
```

**Parameters:**
- `max` (int32)

**Returns:**
- *OptionBuilder

### MaxValue



```go
func (*OptionBuilder) MaxValue(max float64) *OptionBuilder
```

**Parameters:**
- `max` (float64)

**Returns:**
- *OptionBuilder

### MinLength



```go
func (*OptionBuilder) MinLength(min int32) *OptionBuilder
```

**Parameters:**
- `min` (int32)

**Returns:**
- *OptionBuilder

### MinValue



```go
func (*OptionBuilder) MinValue(min float64) *OptionBuilder
```

**Parameters:**
- `min` (float64)

**Returns:**
- *OptionBuilder

### Required



```go
func (*OptionBuilder) Required() *OptionBuilder
```

**Parameters:**
  None

**Returns:**
- *OptionBuilder

### OptionMap
OptionMap provides type-safe access to command options.

#### Example Usage

```go
// Create a new OptionMap
optionmap := OptionMap{

}
```

#### Type Definition

```go
type OptionMap struct {
}
```

### Constructor Functions

### NewOptionMap

NewOptionMap creates a new OptionMap.

```go
func NewOptionMap() *OptionMap
```

**Parameters:**
  None

**Returns:**
- *OptionMap

## Methods

### All

All returns all options as a map.

```go
func (*OptionMap) All() map[string]any
```

**Parameters:**
  None

**Returns:**
- map[string]any

### Attachment

Attachment returns a resolved attachment.

```go
func (*OptionMap) Attachment(name string) *models.Attachment
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.Attachment

### Bool

Bool returns a boolean option value.

```go
func (*Context) Bool(name string) bool
```

**Parameters:**
- `name` (string)

**Returns:**
- bool

### Channel

Channel returns a resolved channel.

```go
func (*OptionMap) Channel(name string) *models.GuildChannel
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.GuildChannel

### Float

Float returns a number option value.

```go
func (*Context) Float(name string) float64
```

**Parameters:**
- `name` (string)

**Returns:**
- float64

### FloatDefault

FloatDefault returns a number option value or a default.

```go
func (*OptionMap) FloatDefault(name string, def float64) float64
```

**Parameters:**
- `name` (string)
- `def` (float64)

**Returns:**
- float64

### Get

Get returns the raw option value.

```go
func (*OptionMap) Get(name string) (any, bool)
```

**Parameters:**
- `name` (string)

**Returns:**
- any
- bool

### Has

Has returns true if an option was provided.

```go
func (*OptionMap) Has(name string) bool
```

**Parameters:**
- `name` (string)

**Returns:**
- bool

### Int

Int returns an integer option value.

```go
func (*Context) Int(name string) int64
```

**Parameters:**
- `name` (string)

**Returns:**
- int64

### IntDefault

IntDefault returns an integer option value or a default.

```go
func (*OptionMap) IntDefault(name string, def int64) int64
```

**Parameters:**
- `name` (string)
- `def` (int64)

**Returns:**
- int64

### Member

Member returns a resolved guild member.

```go
func (*OptionMap) Member(name string) *models.GuildMember
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.GuildMember

### Message

Message returns a resolved message.

```go
func (*OptionMap) Message(name string) *models.Message
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.Message

### Resolved

Resolved returns the resolved data for direct access.

```go
func (*OptionMap) Resolved() *ResolvedData
```

**Parameters:**
  None

**Returns:**
- *ResolvedData

### Role

Role returns a resolved role.

```go
func (*OptionMap) Role(name string) *models.GuildRole
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.GuildRole

### Set

Set sets an option value.

```go
func (*OptionMap) Set(name string, value any)
```

**Parameters:**
- `name` (string)
- `value` (any)

**Returns:**
  None

### SetResolved

SetResolved sets the resolved data.

```go
func (*OptionMap) SetResolved(r *ResolvedData)
```

**Parameters:**
- `r` (*ResolvedData)

**Returns:**
  None

### String

String returns a string option value.

```go
func (*Context) String(name string) string
```

**Parameters:**
- `name` (string)

**Returns:**
- string

### StringDefault

StringDefault returns a string option value or a default.

```go
func (*OptionMap) StringDefault(name, def string) string
```

**Parameters:**
- `name` (string)
- `def` (string)

**Returns:**
- string

### User

User returns a resolved user.

```go
func (*OptionMap) User(name string) *models.User
```

**Parameters:**
- `name` (string)

**Returns:**
- *models.User

### OptionProvider
OptionProvider is an optional interface for commands that have options.

#### Example Usage

```go
// Example implementation of OptionProvider
type MyOptionProvider struct {
    // Add your fields here
}

func (m MyOptionProvider) Options() []Option {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type OptionProvider interface {
    Options() []Option
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### OptionType
OptionType represents the type of a command option.

#### Example Usage

```go
// Example usage of OptionType
var value OptionType
// Initialize with appropriate value
```

#### Type Definition

```go
type OptionType int32
```

### ResolvedData
ResolvedData contains resolved entities from interaction options.

#### Example Usage

```go
// Create a new ResolvedData
resolveddata := ResolvedData{
    Users: map[],
    Members: map[],
    Roles: map[],
    Channels: map[],
    Messages: map[],
    Attachments: map[],
}
```

#### Type Definition

```go
type ResolvedData struct {
    Users map[string]*models.User
    Members map[string]*models.GuildMember
    Roles map[string]*models.GuildRole
    Channels map[string]*models.GuildChannel
    Messages map[string]*models.Message
    Attachments map[string]*models.Attachment
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Users | `map[string]*models.User` |  |
| Members | `map[string]*models.GuildMember` |  |
| Roles | `map[string]*models.GuildRole` |  |
| Channels | `map[string]*models.GuildChannel` |  |
| Messages | `map[string]*models.Message` |  |
| Attachments | `map[string]*models.Attachment` |  |

### ResolvedInteractionData
ResolvedInteractionData contains resolved entities.

#### Example Usage

```go
// Create a new ResolvedInteractionData
resolvedinteractiondata := ResolvedInteractionData{
    Users: map[],
    Members: map[],
    Roles: map[],
    Channels: map[],
    Messages: map[],
    Attachments: map[],
}
```

#### Type Definition

```go
type ResolvedInteractionData struct {
    Users map[string]*models.User `json:"users,omitempty"`
    Members map[string]*models.GuildMember `json:"members,omitempty"`
    Roles map[string]*models.GuildRole `json:"roles,omitempty"`
    Channels map[string]*models.GuildChannel `json:"channels,omitempty"`
    Messages map[string]*models.Message `json:"messages,omitempty"`
    Attachments map[string]*models.Attachment `json:"attachments,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Users | `map[string]*models.User` |  |
| Members | `map[string]*models.GuildMember` |  |
| Roles | `map[string]*models.GuildRole` |  |
| Channels | `map[string]*models.GuildChannel` |  |
| Messages | `map[string]*models.Message` |  |
| Attachments | `map[string]*models.Attachment` |  |

### Responder
Responder is an interface for responding to interactions. This is implemented by the Bot to avoid circular imports.

#### Example Usage

```go
// Example implementation of Responder
type MyResponder struct {
    // Add your fields here
}

func (m MyResponder) RespondToInteraction(param1 context.Context, param2 string, param3 *InteractionResponse) error {
    // Implement your logic here
    return
}

func (m MyResponder) EditInteractionResponse(param1 context.Context, param2 string, param3 *MessageEdit) error {
    // Implement your logic here
    return
}

func (m MyResponder) DeleteInteractionResponse(param1 context.Context, param2 string) error {
    // Implement your logic here
    return
}

func (m MyResponder) FollowupMessage(param1 context.Context, param2 string, param3 *MessageCreate) *models.Message {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type Responder interface {
    RespondToInteraction(ctx context.Context, interactionID, token string, response *InteractionResponse) error
    EditInteractionResponse(ctx context.Context, token string, edit *MessageEdit) error
    DeleteInteractionResponse(ctx context.Context, token string) error
    FollowupMessage(ctx context.Context, token string, message *MessageCreate) (*models.Message, error)
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### Router
Router handles routing interactions to commands and components.

#### Example Usage

```go
// Create a new Router
router := Router{

}
```

#### Type Definition

```go
type Router struct {
}
```

### Constructor Functions

### NewRouter

NewRouter creates a new command router.

```go
func NewRouter() *Router
```

**Parameters:**
  None

**Returns:**
- *Router

## Methods

### Command

Command registers a command.

```go
func (*Router) Command(cmd Command)
```

**Parameters:**
- `cmd` (Command)

**Returns:**
  None

### Commands

Commands returns all registered commands.

```go
func (*Router) Commands() []Command
```

**Parameters:**
  None

**Returns:**
- []Command

### Component

Component registers a component handler.

```go
func (*Router) Component(customID string, handler ComponentHandlerFunc)
```

**Parameters:**
- `customID` (string)
- `handler` (ComponentHandlerFunc)

**Returns:**
  None

### ComponentPrefix

ComponentPrefix registers a component handler that matches by prefix.

```go
func (*Router) ComponentPrefix(prefix string, handler ComponentHandlerFunc)
```

**Parameters:**
- `prefix` (string)
- `handler` (ComponentHandlerFunc)

**Returns:**
  None

### GetCommand

GetCommand returns a command by name.

```go
func (*Router) GetCommand(name string) (Command, bool)
```

**Parameters:**
- `name` (string)

**Returns:**
- Command
- bool

### Group

Group registers a command group (subcommand group).

```go
func (*Router) Group(name, description string, cmds ...Command)
```

**Parameters:**
- `name` (string)
- `description` (string)
- `cmds` (...Command)

**Returns:**
  None

### HandleInteraction

HandleInteraction routes an interaction to the appropriate handler.

```go
func (*Router) HandleInteraction(ctx context.Context, interaction *Interaction) error
```

**Parameters:**
- `ctx` (context.Context)
- `interaction` (*Interaction)

**Returns:**
- error

### MessageContext

MessageContext registers a message context menu command.

```go
func (*Router) MessageContext(name string, handler HandlerFunc)
```

**Parameters:**
- `name` (string)
- `handler` (HandlerFunc)

**Returns:**
  None

### MustSync

MustSync syncs commands and panics on error.

```go
func (*Router) MustSync(ctx context.Context)
```

**Parameters:**
- `ctx` (context.Context)

**Returns:**
  None

### RegisterService

RegisterService registers a service for dependency injection. Services can be retrieved in handlers using commands.Service[T](ctx).

```go
func (*Router) RegisterService(service any)
```

**Parameters:**
- `service` (any)

**Returns:**
  None

### Responder

Responder returns the current responder, or nil if not set.

```go
func (*Router) Responder() Responder
```

**Parameters:**
  None

**Returns:**
- Responder

### SetResponder

SetResponder sets the responder for creating contexts.

```go
func (*Router) SetResponder(responder Responder)
```

**Parameters:**
- `responder` (Responder)

**Returns:**
  None

### SetSyncer

SetSyncer sets the command syncer.

```go
func (*Router) SetSyncer(syncer CommandSyncer, appID string)
```

**Parameters:**
- `syncer` (CommandSyncer)
- `appID` (string)

**Returns:**
  None

### Slash

Slash registers a slash command with a handler.

```go
func (*Router) Slash(name, description string, handler HandlerFunc, options ...Option)
```

**Parameters:**
- `name` (string)
- `description` (string)
- `handler` (HandlerFunc)
- `options` (...Option)

**Returns:**
  None

### Sync

Sync syncs all registered commands with Discord.

```go
func (*Router) Sync(ctx context.Context) error
```

**Parameters:**
- `ctx` (context.Context)

**Returns:**
- error

### SyncGuild

SyncGuild syncs all registered commands with a specific guild.

```go
func (*Router) SyncGuild(ctx context.Context, guildID string) error
```

**Parameters:**
- `ctx` (context.Context)
- `guildID` (string)

**Returns:**
- error

### SyncOnReady

SyncOnReady returns a handler that syncs commands when the bot is ready. This is useful for development to ensure commands are always up-to-date.

```go
func (*Router) SyncOnReady(ctx context.Context, guildIDs ...string) func()
```

**Parameters:**
- `ctx` (context.Context)
- `guildIDs` (...string)

**Returns:**
- func()

### SyncWithOptions

SyncWithOptions syncs commands with custom options.

```go
func (*Router) SyncWithOptions(ctx context.Context, opts *SyncOptions) (*SyncResult, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (*SyncOptions)

**Returns:**
- *SyncResult
- error

### Use

Use adds middleware to the router for command handlers.

```go
func (*Router) Use(mw ...Middleware)
```

**Parameters:**
- `mw` (...Middleware)

**Returns:**
  None

### UseComponent

UseComponent adds middleware to the router for component handlers.

```go
func (*Router) UseComponent(mw ...ComponentMiddleware)
```

**Parameters:**
- `mw` (...ComponentMiddleware)

**Returns:**
  None

### UserContext

UserContext registers a user context menu command.

```go
func (*Router) UserContext(name string, handler HandlerFunc)
```

**Parameters:**
- `name` (string)
- `handler` (HandlerFunc)

**Returns:**
  None

### ServiceRegistry
ServiceRegistry holds registered services for dependency injection.

#### Example Usage

```go
// Create a new ServiceRegistry
serviceregistry := ServiceRegistry{

}
```

#### Type Definition

```go
type ServiceRegistry struct {
}
```

## Methods

### Get

Get retrieves a service by type.

```go
func (*ServiceRegistry) Get(t reflect.Type) (any, bool)
```

**Parameters:**
- `t` (reflect.Type)

**Returns:**
- any
- bool

### Register

Register adds a service to the registry. The service is stored by its concrete type.

```go
func (*ServiceRegistry) Register(service any)
```

**Parameters:**
- `service` (any)

**Returns:**
  None

### SlashCommand
SlashCommand represents a slash command with typed options.

#### Example Usage

```go
// Create a new SlashCommand
slashcommand := SlashCommand{
    CommandName: "example",
    CommandDescription: "example",
    CommandOptions: [],
    CommandMiddleware: [],
    Handler: HandlerFunc{},
    AutocompleteFunc: /* value */,
}
```

#### Type Definition

```go
type SlashCommand struct {
    CommandName string
    CommandDescription string
    CommandOptions []Option
    CommandMiddleware []Middleware
    Handler HandlerFunc
    AutocompleteFunc func(ctx *Context, focused string) ([]Choice, error)
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| CommandName | `string` |  |
| CommandDescription | `string` |  |
| CommandOptions | `[]Option` |  |
| CommandMiddleware | `[]Middleware` |  |
| Handler | `HandlerFunc` |  |
| AutocompleteFunc | `func(ctx *Context, focused string) ([]Choice, error)` |  |

## Methods

### Autocomplete



```go
func (*OptionBuilder) Autocomplete() *OptionBuilder
```

**Parameters:**
  None

**Returns:**
- *OptionBuilder

### Description



```go
func (*SubcommandGroup) Description() string
```

**Parameters:**
  None

**Returns:**
- string

### Execute



```go
func (*SubcommandGroup) Execute(*Context) error
```

**Parameters:**
- `` (*Context)

**Returns:**
- error

### Middleware



```go
func (*MessageCommand) Middleware() []Middleware
```

**Parameters:**
  None

**Returns:**
- []Middleware

### Name



```go
func (*SubcommandGroup) Name() string
```

**Parameters:**
  None

**Returns:**
- string

### Options



```go
func (*SlashCommand) Options() []Option
```

**Parameters:**
  None

**Returns:**
- []Option

### Type



```go
func (*SubcommandGroup) Type() CommandType
```

**Parameters:**
  None

**Returns:**
- CommandType

### SubcommandGroup
SubcommandGroup represents a group of subcommands.

#### Example Usage

```go
// Create a new SubcommandGroup
subcommandgroup := SubcommandGroup{
    GroupName: "example",
    GroupDescription: "example",
    GroupSubcommands: [],
}
```

#### Type Definition

```go
type SubcommandGroup struct {
    GroupName string
    GroupDescription string
    GroupSubcommands []Command
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| GroupName | `string` |  |
| GroupDescription | `string` |  |
| GroupSubcommands | `[]Command` |  |

## Methods

### Description



```go
func (*SubcommandGroup) Description() string
```

**Parameters:**
  None

**Returns:**
- string

### Execute



```go
func (*SubcommandGroup) Execute(*Context) error
```

**Parameters:**
- `` (*Context)

**Returns:**
- error

### Name



```go
func (*SubcommandGroup) Name() string
```

**Parameters:**
  None

**Returns:**
- string

### Subcommands



```go
func (*SubcommandGroup) Subcommands() []Command
```

**Parameters:**
  None

**Returns:**
- []Command

### Type



```go
func (*SubcommandGroup) Type() CommandType
```

**Parameters:**
  None

**Returns:**
- CommandType

### SubcommandProvider
SubcommandProvider is an optional interface for commands that have subcommands.

#### Example Usage

```go
// Example implementation of SubcommandProvider
type MySubcommandProvider struct {
    // Add your fields here
}

func (m MySubcommandProvider) Subcommands() []Command {
    // Implement your logic here
    return
}


```

#### Type Definition

```go
type SubcommandProvider interface {
    Subcommands() []Command
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### SyncOptions
SyncOptions configures command synchronization.

#### Example Usage

```go
// Create a new SyncOptions
syncoptions := SyncOptions{
    GuildIDs: [],
    DeleteUnknown: true,
    DryRun: true,
}
```

#### Type Definition

```go
type SyncOptions struct {
    GuildIDs []string
    DeleteUnknown bool
    DryRun bool
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| GuildIDs | `[]string` | GuildIDs limits sync to specific guilds (for development). If empty, syncs globally. |
| DeleteUnknown | `bool` | DeleteUnknown removes commands not registered in the router. |
| DryRun | `bool` | DryRun logs what would happen without making changes. |

### SyncResult
SyncResult contains the result of a sync operation.

#### Example Usage

```go
// Create a new SyncResult
syncresult := SyncResult{
    Created: [],
    Updated: [],
    Deleted: [],
    Errors: [],
}
```

#### Type Definition

```go
type SyncResult struct {
    Created []string
    Updated []string
    Deleted []string
    Errors []error
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Created | `[]string` |  |
| Updated | `[]string` |  |
| Deleted | `[]string` |  |
| Errors | `[]error` |  |

### UserCommand
UserCommand represents a user context menu command.

#### Example Usage

```go
// Create a new UserCommand
usercommand := UserCommand{
    CommandName: "example",
    CommandMiddleware: [],
    Handler: HandlerFunc{},
}
```

#### Type Definition

```go
type UserCommand struct {
    CommandName string
    CommandMiddleware []Middleware
    Handler HandlerFunc
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| CommandName | `string` |  |
| CommandMiddleware | `[]Middleware` |  |
| Handler | `HandlerFunc` |  |

## Methods

### Description



```go
func (*SubcommandGroup) Description() string
```

**Parameters:**
  None

**Returns:**
- string

### Execute



```go
func (*SubcommandGroup) Execute(*Context) error
```

**Parameters:**
- `` (*Context)

**Returns:**
- error

### Middleware



```go
func (*MessageCommand) Middleware() []Middleware
```

**Parameters:**
  None

**Returns:**
- []Middleware

### Name



```go
func (*SubcommandGroup) Name() string
```

**Parameters:**
  None

**Returns:**
- string

### Type



```go
func (*SubcommandGroup) Type() CommandType
```

**Parameters:**
  None

**Returns:**
- CommandType

## Functions

### ComponentService
ComponentService retrieves a typed service from a ComponentContext. Returns the zero value if the service is not registered.

```go
func ComponentService(ctx *ComponentContext) T
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `ctx` | `*ComponentContext` | |

**Returns:**
| Type | Description |
|------|-------------|
| `T` | |

**Example:**

```go
// Example usage of ComponentService
result := ComponentService(/* parameters */)
```

### Service
Service retrieves a typed service from the context. Returns the zero value if the service is not registered. Example: users := commands.Service[*UserService](ctx) if users == nil { return ctx.ReplyEphemeral("Service unavailable") }

```go
func Service(ctx *Context) T
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `ctx` | `*Context` | |

**Returns:**
| Type | Description |
|------|-------------|
| `T` | |

**Example:**

```go
// Example usage of Service
result := Service(/* parameters */)
```

## External Links

- [Package Overview](../packages/commands.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/discord/commands)
- [Source Code](https://github.com/kolosys/discord/tree/main/commands)

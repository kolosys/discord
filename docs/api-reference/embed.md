# embed API

Complete API documentation for the embed package.

**Import Path:** `github.com/kolosys/discord/embed`

## Package Documentation



## Constants

**ColorSuccess, ColorError, ColorWarning, ColorInfo**

Preset colors for common embed types


```go
const ColorSuccess = int(types.ColorGreen)	// 0x57F287 - Green

const ColorError = int(types.ColorRed)	// 0xED4245 - Red

const ColorWarning = int(types.ColorYellow)	// 0xFEE75C - Yellow

const ColorInfo = int(types.ColorBlurple)	// 0x5865F2 - Discord Blurple

```

**MaxTitleLength, MaxDescriptionLength, MaxFields, MaxFieldNameLength, MaxFieldValueLength, MaxFooterLength, MaxAuthorNameLength, MaxTotalLength**

Discord embed limits


```go
const MaxTitleLength = 256
const MaxDescriptionLength = 4096
const MaxFields = 25
const MaxFieldNameLength = 256
const MaxFieldValueLength = 1024
const MaxFooterLength = 2048
const MaxAuthorNameLength = 256
const MaxTotalLength = 6000
```

## Variables

**ErrTitleTooLong, ErrDescriptionTooLong, ErrTooManyFields, ErrFieldNameTooLong, ErrFieldValueTooLong, ErrFooterTooLong, ErrAuthorNameTooLong, ErrEmbedTooLarge**

Sentinel errors for embed validation


```go
var ErrTitleTooLong = errors.New("embed: title exceeds 256 characters")
var ErrDescriptionTooLong = errors.New("embed: description exceeds 4096 characters")
var ErrTooManyFields = errors.New("embed: exceeds 25 fields maximum")
var ErrFieldNameTooLong = errors.New("embed: field name exceeds 256 characters")
var ErrFieldValueTooLong = errors.New("embed: field value exceeds 1024 characters")
var ErrFooterTooLong = errors.New("embed: footer text exceeds 2048 characters")
var ErrAuthorNameTooLong = errors.New("embed: author name exceeds 256 characters")
var ErrEmbedTooLarge = errors.New("embed: total length exceeds 6000 characters")
```

## Types

### Builder
Builder provides a fluent API for constructing Discord embeds.

#### Example Usage

```go
// Create a new Builder
builder := Builder{

}
```

#### Type Definition

```go
type Builder struct {
}
```

### Constructor Functions

### ErrorEmbed

ErrorEmbed creates a red embed with the given title and description.

```go
func ErrorEmbed(title, description string) *Builder
```

**Parameters:**
- `title` (string)
- `description` (string)

**Returns:**
- *Builder

### InfoEmbed

InfoEmbed creates a blurple embed with the given title and description.

```go
func InfoEmbed(title, description string) *Builder
```

**Parameters:**
- `title` (string)
- `description` (string)

**Returns:**
- *Builder

### New

New creates a new embed builder.

```go
func New() *Builder
```

**Parameters:**
  None

**Returns:**
- *Builder

### SimpleEmbed

SimpleEmbed creates a basic embed with just a description.

```go
func SimpleEmbed(description string) *Builder
```

**Parameters:**
- `description` (string)

**Returns:**
- *Builder

### SuccessEmbed

SuccessEmbed creates a green embed with the given title and description.

```go
func SuccessEmbed(title, description string) *Builder
```

**Parameters:**
- `title` (string)
- `description` (string)

**Returns:**
- *Builder

### TitleEmbed

TitleEmbed creates an embed with just a title.

```go
func TitleEmbed(title string) *Builder
```

**Parameters:**
- `title` (string)

**Returns:**
- *Builder

### WarningEmbed

WarningEmbed creates a yellow embed with the given title and description.

```go
func WarningEmbed(title, description string) *Builder
```

**Parameters:**
- `title` (string)
- `description` (string)

**Returns:**
- *Builder

## Methods

### Author

Author sets the embed author name.

```go
func (*Builder) Author(name string) *Builder
```

**Parameters:**
- `name` (string)

**Returns:**
- *Builder

### AuthorFull

AuthorFull sets all author fields at once.

```go
func (*Builder) AuthorFull(name, url, iconURL string) *Builder
```

**Parameters:**
- `name` (string)
- `url` (string)
- `iconURL` (string)

**Returns:**
- *Builder

### AuthorIcon

AuthorIcon sets the author icon URL.

```go
func (*Builder) AuthorIcon(iconURL string) *Builder
```

**Parameters:**
- `iconURL` (string)

**Returns:**
- *Builder

### AuthorURL

AuthorURL sets the author URL (makes author name a hyperlink).

```go
func (*Builder) AuthorURL(url string) *Builder
```

**Parameters:**
- `url` (string)

**Returns:**
- *Builder

### BlockField

BlockField adds a block (non-inline) field to the embed.

```go
func (*Builder) BlockField(name, value string) *Builder
```

**Parameters:**
- `name` (string)
- `value` (string)

**Returns:**
- *Builder

### Blurple

Blurple sets the embed color to Discord's brand color.

```go
func (*Builder) Blurple() *Builder
```

**Parameters:**
  None

**Returns:**
- *Builder

### Build

Build returns the constructed embed.

```go
func (*Builder) Build() *commands.Embed
```

**Parameters:**
  None

**Returns:**
- *commands.Embed

### BuildSafe

BuildSafe validates and returns the embed, or an error if invalid.

```go
func (*Builder) BuildSafe() (*commands.Embed, error)
```

**Parameters:**
  None

**Returns:**
- *commands.Embed
- error

### Color

Color sets the embed color.

```go
func (*Builder) Color(color int) *Builder
```

**Parameters:**
- `color` (int)

**Returns:**
- *Builder

### Copy

Copy returns a new builder with a copy of the current embed.

```go
func (*Builder) Copy() *Builder
```

**Parameters:**
  None

**Returns:**
- *Builder

### Description

Description sets the embed description.

```go
func (*Builder) Description(desc string) *Builder
```

**Parameters:**
- `desc` (string)

**Returns:**
- *Builder

### Error

Error sets the embed color to red (error).

```go
func (*Builder) Error() *Builder
```

**Parameters:**
  None

**Returns:**
- *Builder

### Field

Field adds a field to the embed.

```go
func (*Builder) Field(name, value string, inline bool) *Builder
```

**Parameters:**
- `name` (string)
- `value` (string)
- `inline` (bool)

**Returns:**
- *Builder

### Fields

Fields adds multiple fields at once.

```go
func (*Builder) Fields(fields ...*commands.EmbedField) *Builder
```

**Parameters:**
- `fields` (...*commands.EmbedField)

**Returns:**
- *Builder

### Footer

Footer sets the embed footer text.

```go
func (*Builder) Footer(text string) *Builder
```

**Parameters:**
- `text` (string)

**Returns:**
- *Builder

### FooterFull

FooterFull sets both footer text and icon.

```go
func (*Builder) FooterFull(text, iconURL string) *Builder
```

**Parameters:**
- `text` (string)
- `iconURL` (string)

**Returns:**
- *Builder

### FooterIcon

FooterIcon sets the footer icon URL.

```go
func (*Builder) FooterIcon(iconURL string) *Builder
```

**Parameters:**
- `iconURL` (string)

**Returns:**
- *Builder

### Image

Image sets the embed image URL.

```go
func (*Builder) Image(url string) *Builder
```

**Parameters:**
- `url` (string)

**Returns:**
- *Builder

### Info

Info sets the embed color to blurple (info).

```go
func (*Builder) Info() *Builder
```

**Parameters:**
  None

**Returns:**
- *Builder

### InlineField

InlineField adds an inline field to the embed.

```go
func (*Builder) InlineField(name, value string) *Builder
```

**Parameters:**
- `name` (string)
- `value` (string)

**Returns:**
- *Builder

### Success

Success sets the embed color to green (success).

```go
func (*Builder) Success() *Builder
```

**Parameters:**
  None

**Returns:**
- *Builder

### Thumbnail

Thumbnail sets the embed thumbnail URL.

```go
func (*Builder) Thumbnail(url string) *Builder
```

**Parameters:**
- `url` (string)

**Returns:**
- *Builder

### Timestamp

Timestamp sets the embed timestamp.

```go
func (*Builder) Timestamp(ts string) *Builder
```

**Parameters:**
- `ts` (string)

**Returns:**
- *Builder

### TimestampNow

TimestampNow sets the timestamp to the current time.

```go
func (*Builder) TimestampNow() *Builder
```

**Parameters:**
  None

**Returns:**
- *Builder

### TimestampTime

TimestampTime sets the timestamp from a time.Time value.

```go
func (*Builder) TimestampTime(t time.Time) *Builder
```

**Parameters:**
- `t` (time.Time)

**Returns:**
- *Builder

### Title

Title sets the embed title.

```go
func (*Builder) Title(title string) *Builder
```

**Parameters:**
- `title` (string)

**Returns:**
- *Builder

### URL

URL sets the embed URL (makes title a hyperlink).

```go
func (*Builder) URL(url string) *Builder
```

**Parameters:**
- `url` (string)

**Returns:**
- *Builder

### Validate

Validate checks if the embed conforms to Discord's limits.

```go
func (*Builder) Validate() error
```

**Parameters:**
  None

**Returns:**
- error

### Warning

Warning sets the embed color to yellow (warning).

```go
func (*Builder) Warning() *Builder
```

**Parameters:**
  None

**Returns:**
- *Builder

## External Links

- [Package Overview](../packages/embed.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/discord/embed)
- [Source Code](https://github.com/kolosys/discord/tree/main/embed)

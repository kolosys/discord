# components API

Complete API documentation for the components package.

**Import Path:** `github.com/kolosys/discord/components`

## Package Documentation



## Constants

**MaxCustomIDLength, MaxActionRows, MaxButtonsPerRow, MaxSelectOptionsPerMenu, MaxComponentsPerMessage, MaxTextInputsPerModal, MaxButtonLabelLength, MaxSelectPlaceholderLength, MaxSelectOptionLabelLength, MaxSelectOptionValueLength, MaxSelectOptionDescriptionLength, MaxModalTitleLength, MaxTextInputLabelLength, MaxTextInputPlaceholderLength, MaxTextInputValueLength**

Discord API limits for message components.


```go
const MaxCustomIDLength = 100	// MaxCustomIDLength is the maximum length of a custom_id field.

const MaxActionRows = 5	// MaxActionRows is the maximum number of action rows per message.

const MaxButtonsPerRow = 5	// MaxButtonsPerRow is the maximum number of buttons per action row.

const MaxSelectOptionsPerMenu = 25	// MaxSelectOptionsPerMenu is the maximum number of options in a string select menu.

const MaxComponentsPerMessage = 25	// MaxComponentsPerMessage is the maximum total components per message.

const MaxTextInputsPerModal = 5	// MaxTextInputsPerModal is the maximum number of text inputs per modal.

const MaxButtonLabelLength = 80	// MaxButtonLabelLength is the maximum length of a button label.

const MaxSelectPlaceholderLength = 150	// MaxSelectPlaceholderLength is the maximum length of a select menu placeholder.

const MaxSelectOptionLabelLength = 100	// MaxSelectOptionLabelLength is the maximum length of a select option label.

const MaxSelectOptionValueLength = 100	// MaxSelectOptionValueLength is the maximum length of a select option value.

const MaxSelectOptionDescriptionLength = 100	// MaxSelectOptionDescriptionLength is the maximum length of a select option description.

const MaxModalTitleLength = 45	// MaxModalTitleLength is the maximum length of a modal title.

const MaxTextInputLabelLength = 45	// MaxTextInputLabelLength is the maximum length of a text input label.

const MaxTextInputPlaceholderLength = 100	// MaxTextInputPlaceholderLength is the maximum length of a text input placeholder.

const MaxTextInputValueLength = 4000	// MaxTextInputValueLength is the maximum length of a text input value.

```

## Variables

**ErrCustomIDTooLong, ErrTooManyActionRows, ErrTooManyButtonsPerRow, ErrTooManySelectOptions, ErrTooManyComponents, ErrTooManyModalInputs, ErrButtonLabelTooLong, ErrSelectPlaceholderTooLong, ErrSelectOptionLabelTooLong, ErrSelectOptionValueTooLong, ErrSelectOptionDescTooLong, ErrModalTitleTooLong, ErrTextInputLabelTooLong, ErrTextInputPlaceholderTooLong, ErrNoSelectOptions, ErrButtonMissingLabelOrEmoji, ErrLinkButtonWithCustomID, ErrNonLinkButtonMissingCustomID**

Validation errors.


```go
var ErrCustomIDTooLong = errors.New("components: custom_id exceeds 100 characters")
var ErrTooManyActionRows = errors.New("components: message cannot have more than 5 action rows")
var ErrTooManyButtonsPerRow = errors.New("components: action row cannot have more than 5 buttons")
var ErrTooManySelectOptions = errors.New("components: select menu cannot have more than 25 options")
var ErrTooManyComponents = errors.New("components: message cannot have more than 25 components")
var ErrTooManyModalInputs = errors.New("components: modal cannot have more than 5 text inputs")
var ErrButtonLabelTooLong = errors.New("components: button label exceeds 80 characters")
var ErrSelectPlaceholderTooLong = errors.New("components: select placeholder exceeds 150 characters")
var ErrSelectOptionLabelTooLong = errors.New("components: select option label exceeds 100 characters")
var ErrSelectOptionValueTooLong = errors.New("components: select option value exceeds 100 characters")
var ErrSelectOptionDescTooLong = errors.New("components: select option description exceeds 100 characters")
var ErrModalTitleTooLong = errors.New("components: modal title exceeds 45 characters")
var ErrTextInputLabelTooLong = errors.New("components: text input label exceeds 45 characters")
var ErrTextInputPlaceholderTooLong = errors.New("components: text input placeholder exceeds 100 characters")
var ErrNoSelectOptions = errors.New("components: select menu must have at least 1 option")
var ErrButtonMissingLabelOrEmoji = errors.New("components: button must have a label or emoji")
var ErrLinkButtonWithCustomID = errors.New("components: link button cannot have a custom_id")
var ErrNonLinkButtonMissingCustomID = errors.New("components: non-link button must have a custom_id")
```

## Types

### ActionRowComponent
ActionRowComponent is a container for other components.

#### Example Usage

```go
// Create a new ActionRowComponent
actionrowcomponent := ActionRowComponent{
    Type: /* value */,
    Components: [],
}
```

#### Type Definition

```go
type ActionRowComponent struct {
    Type types.ComponentType `json:"type"`
    Components []Component `json:"components"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Type | `types.ComponentType` |  |
| Components | `[]Component` |  |

### Constructor Functions

### ActionRow

ActionRow creates a new action row with the given components.

```go
func ActionRow(components ...Component) *ActionRowComponent
```

**Parameters:**
- `components` (...Component)

**Returns:**
- *ActionRowComponent

### ButtonBuilder
ButtonBuilder builds buttons with a fluent API.

#### Example Usage

```go
// Create a new ButtonBuilder
buttonbuilder := ButtonBuilder{

}
```

#### Type Definition

```go
type ButtonBuilder struct {
}
```

### Constructor Functions

### Button

Button creates a new button builder.

```go
func Button(customID string) *ButtonBuilder
```

**Parameters:**
- `customID` (string)

**Returns:**
- *ButtonBuilder

### LinkButton

LinkButton creates a link button.

```go
func LinkButton(url string) *ButtonBuilder
```

**Parameters:**
- `url` (string)

**Returns:**
- *ButtonBuilder

### PremiumButton

PremiumButton creates a premium/SKU button.

```go
func PremiumButton(skuID string) *ButtonBuilder
```

**Parameters:**
- `skuID` (string)

**Returns:**
- *ButtonBuilder

## Methods

### Build

Build returns the built button component.

```go
func (*ButtonBuilder) Build() *ButtonComponent
```

**Parameters:**
  None

**Returns:**
- *ButtonComponent

### BuildSafe

BuildSafe validates and returns the built button component.

```go
func (*ButtonBuilder) BuildSafe() (*ButtonComponent, error)
```

**Parameters:**
  None

**Returns:**
- *ButtonComponent
- error

### CustomEmoji

CustomEmoji sets the button emoji by ID (for custom emoji).

```go
func (*ButtonBuilder) CustomEmoji(id, name string, animated bool) *ButtonBuilder
```

**Parameters:**
- `id` (string)
- `name` (string)
- `animated` (bool)

**Returns:**
- *ButtonBuilder

### Danger

Danger sets the button to danger style.

```go
func (*ButtonBuilder) Danger() *ButtonBuilder
```

**Parameters:**
  None

**Returns:**
- *ButtonBuilder

### Disabled

Disabled sets whether the button is disabled.

```go
func (*ButtonBuilder) Disabled(disabled bool) *ButtonBuilder
```

**Parameters:**
- `disabled` (bool)

**Returns:**
- *ButtonBuilder

### Emoji

Emoji sets the button emoji by name (for standard emoji).

```go
func (*ButtonBuilder) Emoji(name string) *ButtonBuilder
```

**Parameters:**
- `name` (string)

**Returns:**
- *ButtonBuilder

### Label

Label sets the button label.

```go
func (*ButtonBuilder) Label(label string) *ButtonBuilder
```

**Parameters:**
- `label` (string)

**Returns:**
- *ButtonBuilder

### Primary

Primary sets the button to primary style.

```go
func (*ButtonBuilder) Primary() *ButtonBuilder
```

**Parameters:**
  None

**Returns:**
- *ButtonBuilder

### Secondary

Secondary sets the button to secondary style.

```go
func (*ButtonBuilder) Secondary() *ButtonBuilder
```

**Parameters:**
  None

**Returns:**
- *ButtonBuilder

### Style

Style sets the button style.

```go
func (*ButtonBuilder) Style(style types.ButtonStyle) *ButtonBuilder
```

**Parameters:**
- `style` (types.ButtonStyle)

**Returns:**
- *ButtonBuilder

### Success

Success sets the button to success style.

```go
func (*ButtonBuilder) Success() *ButtonBuilder
```

**Parameters:**
  None

**Returns:**
- *ButtonBuilder

### Validate

Validate checks if the button configuration is valid according to Discord limits.

```go
func (*ButtonBuilder) Validate() error
```

**Parameters:**
  None

**Returns:**
- error

### ButtonComponent
ButtonComponent represents a clickable button.

#### Example Usage

```go
// Create a new ButtonComponent
buttoncomponent := ButtonComponent{
    Type: /* value */,
    Style: /* value */,
    Label: "example",
    Emoji: &ButtonEmoji{}{},
    CustomID: "example",
    URL: "example",
    Disabled: true,
    SkuID: "example",
}
```

#### Type Definition

```go
type ButtonComponent struct {
    Type types.ComponentType `json:"type"`
    Style types.ButtonStyle `json:"style"`
    Label string `json:"label,omitempty"`
    Emoji *ButtonEmoji `json:"emoji,omitempty"`
    CustomID string `json:"custom_id,omitempty"`
    URL string `json:"url,omitempty"`
    Disabled bool `json:"disabled,omitempty"`
    SkuID string `json:"sku_id,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Type | `types.ComponentType` |  |
| Style | `types.ButtonStyle` |  |
| Label | `string` |  |
| Emoji | `*ButtonEmoji` |  |
| CustomID | `string` |  |
| URL | `string` |  |
| Disabled | `bool` |  |
| SkuID | `string` |  |

### Constructor Functions

### DangerButton

DangerButton creates a danger style button.

```go
func DangerButton(customID, label string) *ButtonComponent
```

**Parameters:**
- `customID` (string)
- `label` (string)

**Returns:**
- *ButtonComponent

### Link

Link creates a link button.

```go
func Link(url, label string) *ButtonComponent
```

**Parameters:**
- `url` (string)
- `label` (string)

**Returns:**
- *ButtonComponent

### PrimaryButton

PrimaryButton creates a primary style button.

```go
func PrimaryButton(customID, label string) *ButtonComponent
```

**Parameters:**
- `customID` (string)
- `label` (string)

**Returns:**
- *ButtonComponent

### SecondaryButton

SecondaryButton creates a secondary style button.

```go
func SecondaryButton(customID, label string) *ButtonComponent
```

**Parameters:**
- `customID` (string)
- `label` (string)

**Returns:**
- *ButtonComponent

### SuccessButton

SuccessButton creates a success style button.

```go
func SuccessButton(customID, label string) *ButtonComponent
```

**Parameters:**
- `customID` (string)
- `label` (string)

**Returns:**
- *ButtonComponent

### ButtonEmoji
ButtonEmoji represents an emoji on a button.

#### Example Usage

```go
// Create a new ButtonEmoji
buttonemoji := ButtonEmoji{
    ID: "example",
    Name: "example",
    Animated: true,
}
```

#### Type Definition

```go
type ButtonEmoji struct {
    ID string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
    Animated bool `json:"animated,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Name | `string` |  |
| Animated | `bool` |  |

### ChannelSelectBuilder
ChannelSelectBuilder builds channel select menus.

#### Example Usage

```go
// Create a new ChannelSelectBuilder
channelselectbuilder := ChannelSelectBuilder{

}
```

#### Type Definition

```go
type ChannelSelectBuilder struct {
}
```

### Constructor Functions

### ChannelSelect

ChannelSelect creates a new channel select builder.

```go
func ChannelSelect(customID string) *ChannelSelectBuilder
```

**Parameters:**
- `customID` (string)

**Returns:**
- *ChannelSelectBuilder

## Methods

### Build

Build returns the built select component.

```go
func (*ButtonBuilder) Build() *ButtonComponent
```

**Parameters:**
  None

**Returns:**
- *ButtonComponent

### BuildSafe

BuildSafe validates and returns the built select component.

```go
func (*ModalBuilder) BuildSafe() (*ModalData, error)
```

**Parameters:**
  None

**Returns:**
- *ModalData
- error

### ChannelTypes

ChannelTypes filters the channels by type.

```go
func (*ChannelSelectBuilder) ChannelTypes(types ...int32) *ChannelSelectBuilder
```

**Parameters:**
- `types` (...int32)

**Returns:**
- *ChannelSelectBuilder

### Disabled

Disabled sets whether the select is disabled.

```go
func (*MentionableSelectBuilder) Disabled(disabled bool) *MentionableSelectBuilder
```

**Parameters:**
- `disabled` (bool)

**Returns:**
- *MentionableSelectBuilder

### MaxValues

MaxValues sets the maximum number of selections.

```go
func (*MentionableSelectBuilder) MaxValues(max int) *MentionableSelectBuilder
```

**Parameters:**
- `max` (int)

**Returns:**
- *MentionableSelectBuilder

### MinValues

MinValues sets the minimum number of selections.

```go
func (*MentionableSelectBuilder) MinValues(min int) *MentionableSelectBuilder
```

**Parameters:**
- `min` (int)

**Returns:**
- *MentionableSelectBuilder

### Placeholder

Placeholder sets the placeholder text.

```go
func (*ModalTextInputBuilder) Placeholder(text string) *ModalTextInputBuilder
```

**Parameters:**
- `text` (string)

**Returns:**
- *ModalTextInputBuilder

### Validate

Validate checks if the select menu configuration is valid according to Discord limits.

```go
func (*ButtonBuilder) Validate() error
```

**Parameters:**
  None

**Returns:**
- error

### ChannelSelectComponent
ChannelSelectComponent represents a select menu for channels.

#### Example Usage

```go
// Create a new ChannelSelectComponent
channelselectcomponent := ChannelSelectComponent{
    Type: /* value */,
    CustomID: "example",
    Placeholder: "example",
    MinValues: &42{},
    MaxValues: &42{},
    Disabled: true,
    ChannelTypes: [],
}
```

#### Type Definition

```go
type ChannelSelectComponent struct {
    Type types.ComponentType `json:"type"`
    CustomID string `json:"custom_id"`
    Placeholder string `json:"placeholder,omitempty"`
    MinValues *int `json:"min_values,omitempty"`
    MaxValues *int `json:"max_values,omitempty"`
    Disabled bool `json:"disabled,omitempty"`
    ChannelTypes []int32 `json:"channel_types,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Type | `types.ComponentType` |  |
| CustomID | `string` |  |
| Placeholder | `string` |  |
| MinValues | `*int` |  |
| MaxValues | `*int` |  |
| Disabled | `bool` |  |
| ChannelTypes | `[]int32` |  |

### Component
Component is the interface for all message components.

#### Example Usage

```go
// Example implementation of Component
type MyComponent struct {
    // Add your fields here
}


```

#### Type Definition

```go
type Component interface {
}
```

## Methods

| Method | Description |
| ------ | ----------- |

### MentionableSelectBuilder
MentionableSelectBuilder builds mentionable select menus.

#### Example Usage

```go
// Create a new MentionableSelectBuilder
mentionableselectbuilder := MentionableSelectBuilder{

}
```

#### Type Definition

```go
type MentionableSelectBuilder struct {
}
```

### Constructor Functions

### MentionableSelect

MentionableSelect creates a new mentionable select builder.

```go
func MentionableSelect(customID string) *MentionableSelectBuilder
```

**Parameters:**
- `customID` (string)

**Returns:**
- *MentionableSelectBuilder

## Methods

### Build

Build returns the built select component.

```go
func (*ButtonBuilder) Build() *ButtonComponent
```

**Parameters:**
  None

**Returns:**
- *ButtonComponent

### BuildSafe

BuildSafe validates and returns the built select component.

```go
func (*ButtonBuilder) BuildSafe() (*ButtonComponent, error)
```

**Parameters:**
  None

**Returns:**
- *ButtonComponent
- error

### Disabled

Disabled sets whether the select is disabled.

```go
func (*ButtonBuilder) Disabled(disabled bool) *ButtonBuilder
```

**Parameters:**
- `disabled` (bool)

**Returns:**
- *ButtonBuilder

### MaxValues

MaxValues sets the maximum number of selections.

```go
func (*MentionableSelectBuilder) MaxValues(max int) *MentionableSelectBuilder
```

**Parameters:**
- `max` (int)

**Returns:**
- *MentionableSelectBuilder

### MinValues

MinValues sets the minimum number of selections.

```go
func (*MentionableSelectBuilder) MinValues(min int) *MentionableSelectBuilder
```

**Parameters:**
- `min` (int)

**Returns:**
- *MentionableSelectBuilder

### Placeholder

Placeholder sets the placeholder text.

```go
func (*MentionableSelectBuilder) Placeholder(text string) *MentionableSelectBuilder
```

**Parameters:**
- `text` (string)

**Returns:**
- *MentionableSelectBuilder

### Validate

Validate checks if the select menu configuration is valid according to Discord limits.

```go
func (*ButtonBuilder) Validate() error
```

**Parameters:**
  None

**Returns:**
- error

### MentionableSelectComponent
MentionableSelectComponent represents a select menu for mentionable entities.

#### Example Usage

```go
// Create a new MentionableSelectComponent
mentionableselectcomponent := MentionableSelectComponent{
    Type: /* value */,
    CustomID: "example",
    Placeholder: "example",
    MinValues: &42{},
    MaxValues: &42{},
    Disabled: true,
}
```

#### Type Definition

```go
type MentionableSelectComponent struct {
    Type types.ComponentType `json:"type"`
    CustomID string `json:"custom_id"`
    Placeholder string `json:"placeholder,omitempty"`
    MinValues *int `json:"min_values,omitempty"`
    MaxValues *int `json:"max_values,omitempty"`
    Disabled bool `json:"disabled,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Type | `types.ComponentType` |  |
| CustomID | `string` |  |
| Placeholder | `string` |  |
| MinValues | `*int` |  |
| MaxValues | `*int` |  |
| Disabled | `bool` |  |

### ModalBuilder
ModalBuilder builds modals with a fluent API.

#### Example Usage

```go
// Create a new ModalBuilder
modalbuilder := ModalBuilder{

}
```

#### Type Definition

```go
type ModalBuilder struct {
}
```

### Constructor Functions

### Modal

Modal creates a new modal builder.

```go
func Modal(customID, title string) *ModalBuilder
```

**Parameters:**
- `customID` (string)
- `title` (string)

**Returns:**
- *ModalBuilder

## Methods

### Build

Build returns the modal data for responding.

```go
func (*ButtonBuilder) Build() *ButtonComponent
```

**Parameters:**
  None

**Returns:**
- *ButtonComponent

### BuildSafe

BuildSafe validates and returns the modal data.

```go
func (*ModalBuilder) BuildSafe() (*ModalData, error)
```

**Parameters:**
  None

**Returns:**
- *ModalData
- error

### Component

Component adds a raw component.

```go
func (*ModalBuilder) Component(comp any) *ModalBuilder
```

**Parameters:**
- `comp` (any)

**Returns:**
- *ModalBuilder

### ParagraphInput

ParagraphInput adds a paragraph text input.

```go
func (*ModalBuilder) ParagraphInput(customID, label string) *ModalBuilder
```

**Parameters:**
- `customID` (string)
- `label` (string)

**Returns:**
- *ModalBuilder

### ShortInput

ShortInput adds a short text input.

```go
func (*ModalBuilder) ShortInput(customID, label string) *ModalBuilder
```

**Parameters:**
- `customID` (string)
- `label` (string)

**Returns:**
- *ModalBuilder

### TextInput

TextInput adds a text input to the modal.

```go
func (*ModalBuilder) TextInput(customID, label string) *ModalTextInputBuilder
```

**Parameters:**
- `customID` (string)
- `label` (string)

**Returns:**
- *ModalTextInputBuilder

### Validate

Validate checks if the modal configuration is valid according to Discord limits.

```go
func (*ButtonBuilder) Validate() error
```

**Parameters:**
  None

**Returns:**
- error

### ModalData
ModalData represents the data for a modal response.

#### Example Usage

```go
// Create a new ModalData
modaldata := ModalData{
    CustomID: "example",
    Title: "example",
    Components: [],
}
```

#### Type Definition

```go
type ModalData struct {
    CustomID string `json:"custom_id"`
    Title string `json:"title"`
    Components []any `json:"components"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| CustomID | `string` |  |
| Title | `string` |  |
| Components | `[]any` |  |

### ModalTextInputBuilder
ModalTextInputBuilder is a builder for text inputs within a modal.

#### Example Usage

```go
// Create a new ModalTextInputBuilder
modaltextinputbuilder := ModalTextInputBuilder{

}
```

#### Type Definition

```go
type ModalTextInputBuilder struct {
}
```

## Methods

### Done

Done finishes the text input and returns to the modal builder.

```go
func (*ModalTextInputBuilder) Done() *ModalBuilder
```

**Parameters:**
  None

**Returns:**
- *ModalBuilder

### MaxLength

MaxLength sets the maximum length.

```go
func (*ModalTextInputBuilder) MaxLength(max int) *ModalTextInputBuilder
```

**Parameters:**
- `max` (int)

**Returns:**
- *ModalTextInputBuilder

### MinLength

MinLength sets the minimum length.

```go
func (*ModalTextInputBuilder) MinLength(min int) *ModalTextInputBuilder
```

**Parameters:**
- `min` (int)

**Returns:**
- *ModalTextInputBuilder

### Paragraph

Paragraph sets the input to multi-line style.

```go
func (*ModalTextInputBuilder) Paragraph() *ModalTextInputBuilder
```

**Parameters:**
  None

**Returns:**
- *ModalTextInputBuilder

### Placeholder

Placeholder sets the placeholder text.

```go
func (*MentionableSelectBuilder) Placeholder(text string) *MentionableSelectBuilder
```

**Parameters:**
- `text` (string)

**Returns:**
- *MentionableSelectBuilder

### Required

Required sets whether the input is required.

```go
func (*ModalTextInputBuilder) Required(required bool) *ModalTextInputBuilder
```

**Parameters:**
- `required` (bool)

**Returns:**
- *ModalTextInputBuilder

### Short

Short sets the input to single-line style.

```go
func (*ModalTextInputBuilder) Short() *ModalTextInputBuilder
```

**Parameters:**
  None

**Returns:**
- *ModalTextInputBuilder

### Value

Value sets the default value.

```go
func (*ModalTextInputBuilder) Value(value string) *ModalTextInputBuilder
```

**Parameters:**
- `value` (string)

**Returns:**
- *ModalTextInputBuilder

### RoleSelectBuilder
RoleSelectBuilder builds role select menus.

#### Example Usage

```go
// Create a new RoleSelectBuilder
roleselectbuilder := RoleSelectBuilder{

}
```

#### Type Definition

```go
type RoleSelectBuilder struct {
}
```

### Constructor Functions

### RoleSelect

RoleSelect creates a new role select builder.

```go
func RoleSelect(customID string) *RoleSelectBuilder
```

**Parameters:**
- `customID` (string)

**Returns:**
- *RoleSelectBuilder

## Methods

### Build

Build returns the built select component.

```go
func (*ButtonBuilder) Build() *ButtonComponent
```

**Parameters:**
  None

**Returns:**
- *ButtonComponent

### BuildSafe

BuildSafe validates and returns the built select component.

```go
func (*ButtonBuilder) BuildSafe() (*ButtonComponent, error)
```

**Parameters:**
  None

**Returns:**
- *ButtonComponent
- error

### Disabled

Disabled sets whether the select is disabled.

```go
func (*ButtonBuilder) Disabled(disabled bool) *ButtonBuilder
```

**Parameters:**
- `disabled` (bool)

**Returns:**
- *ButtonBuilder

### MaxValues

MaxValues sets the maximum number of selections.

```go
func (*MentionableSelectBuilder) MaxValues(max int) *MentionableSelectBuilder
```

**Parameters:**
- `max` (int)

**Returns:**
- *MentionableSelectBuilder

### MinValues

MinValues sets the minimum number of selections.

```go
func (*MentionableSelectBuilder) MinValues(min int) *MentionableSelectBuilder
```

**Parameters:**
- `min` (int)

**Returns:**
- *MentionableSelectBuilder

### Placeholder

Placeholder sets the placeholder text.

```go
func (*ModalTextInputBuilder) Placeholder(text string) *ModalTextInputBuilder
```

**Parameters:**
- `text` (string)

**Returns:**
- *ModalTextInputBuilder

### Validate

Validate checks if the select menu configuration is valid according to Discord limits.

```go
func (*ButtonBuilder) Validate() error
```

**Parameters:**
  None

**Returns:**
- error

### RoleSelectComponent
RoleSelectComponent represents a select menu for roles.

#### Example Usage

```go
// Create a new RoleSelectComponent
roleselectcomponent := RoleSelectComponent{
    Type: /* value */,
    CustomID: "example",
    Placeholder: "example",
    MinValues: &42{},
    MaxValues: &42{},
    Disabled: true,
}
```

#### Type Definition

```go
type RoleSelectComponent struct {
    Type types.ComponentType `json:"type"`
    CustomID string `json:"custom_id"`
    Placeholder string `json:"placeholder,omitempty"`
    MinValues *int `json:"min_values,omitempty"`
    MaxValues *int `json:"max_values,omitempty"`
    Disabled bool `json:"disabled,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Type | `types.ComponentType` |  |
| CustomID | `string` |  |
| Placeholder | `string` |  |
| MinValues | `*int` |  |
| MaxValues | `*int` |  |
| Disabled | `bool` |  |

### SelectOption
SelectOption represents an option in a string select menu.

#### Example Usage

```go
// Create a new SelectOption
selectoption := SelectOption{
    Label: "example",
    Value: "example",
    Description: "example",
    Emoji: &ButtonEmoji{}{},
    Default: true,
}
```

#### Type Definition

```go
type SelectOption struct {
    Label string `json:"label"`
    Value string `json:"value"`
    Description string `json:"description,omitempty"`
    Emoji *ButtonEmoji `json:"emoji,omitempty"`
    Default bool `json:"default,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Label | `string` |  |
| Value | `string` |  |
| Description | `string` |  |
| Emoji | `*ButtonEmoji` |  |
| Default | `bool` |  |

### Constructor Functions

### NewSelectOption

NewSelectOption creates a new select option.

```go
func NewSelectOption(label, value string) SelectOption
```

**Parameters:**
- `label` (string)
- `value` (string)

**Returns:**
- SelectOption

### NewSelectOptionWithEmoji

NewSelectOptionWithEmoji creates a select option with an emoji.

```go
func NewSelectOptionWithEmoji(label, value, emoji string) SelectOption
```

**Parameters:**
- `label` (string)
- `value` (string)
- `emoji` (string)

**Returns:**
- SelectOption

### StringSelectBuilder
StringSelectBuilder builds string select menus.

#### Example Usage

```go
// Create a new StringSelectBuilder
stringselectbuilder := StringSelectBuilder{

}
```

#### Type Definition

```go
type StringSelectBuilder struct {
}
```

### Constructor Functions

### StringSelect

StringSelect creates a new string select builder.

```go
func StringSelect(customID string) *StringSelectBuilder
```

**Parameters:**
- `customID` (string)

**Returns:**
- *StringSelectBuilder

## Methods

### Build

Build returns the built select component.

```go
func (*ButtonBuilder) Build() *ButtonComponent
```

**Parameters:**
  None

**Returns:**
- *ButtonComponent

### BuildSafe

BuildSafe validates and returns the built select component.

```go
func (*MentionableSelectBuilder) BuildSafe() (*MentionableSelectComponent, error)
```

**Parameters:**
  None

**Returns:**
- *MentionableSelectComponent
- error

### Disabled

Disabled sets whether the select is disabled.

```go
func (*MentionableSelectBuilder) Disabled(disabled bool) *MentionableSelectBuilder
```

**Parameters:**
- `disabled` (bool)

**Returns:**
- *MentionableSelectBuilder

### MaxValues

MaxValues sets the maximum number of selections.

```go
func (*MentionableSelectBuilder) MaxValues(max int) *MentionableSelectBuilder
```

**Parameters:**
- `max` (int)

**Returns:**
- *MentionableSelectBuilder

### MinValues

MinValues sets the minimum number of selections.

```go
func (*MentionableSelectBuilder) MinValues(min int) *MentionableSelectBuilder
```

**Parameters:**
- `min` (int)

**Returns:**
- *MentionableSelectBuilder

### Option

Option adds a single option.

```go
func (*StringSelectBuilder) Option(label, value string) *StringSelectBuilder
```

**Parameters:**
- `label` (string)
- `value` (string)

**Returns:**
- *StringSelectBuilder

### OptionWithDescription

OptionWithDescription adds an option with description.

```go
func (*StringSelectBuilder) OptionWithDescription(label, value, description string) *StringSelectBuilder
```

**Parameters:**
- `label` (string)
- `value` (string)
- `description` (string)

**Returns:**
- *StringSelectBuilder

### Options

Options adds options to the select menu.

```go
func (*StringSelectBuilder) Options(opts ...SelectOption) *StringSelectBuilder
```

**Parameters:**
- `opts` (...SelectOption)

**Returns:**
- *StringSelectBuilder

### Placeholder

Placeholder sets the placeholder text.

```go
func (*ModalTextInputBuilder) Placeholder(text string) *ModalTextInputBuilder
```

**Parameters:**
- `text` (string)

**Returns:**
- *ModalTextInputBuilder

### Validate

Validate checks if the select menu configuration is valid according to Discord limits.

```go
func (*ModalBuilder) Validate() error
```

**Parameters:**
  None

**Returns:**
- error

### StringSelectComponent
StringSelectComponent represents a dropdown with string options.

#### Example Usage

```go
// Create a new StringSelectComponent
stringselectcomponent := StringSelectComponent{
    Type: /* value */,
    CustomID: "example",
    Options: [],
    Placeholder: "example",
    MinValues: &42{},
    MaxValues: &42{},
    Disabled: true,
}
```

#### Type Definition

```go
type StringSelectComponent struct {
    Type types.ComponentType `json:"type"`
    CustomID string `json:"custom_id"`
    Options []SelectOption `json:"options"`
    Placeholder string `json:"placeholder,omitempty"`
    MinValues *int `json:"min_values,omitempty"`
    MaxValues *int `json:"max_values,omitempty"`
    Disabled bool `json:"disabled,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Type | `types.ComponentType` |  |
| CustomID | `string` |  |
| Options | `[]SelectOption` |  |
| Placeholder | `string` |  |
| MinValues | `*int` |  |
| MaxValues | `*int` |  |
| Disabled | `bool` |  |

### TextInputBuilder
TextInputBuilder builds text inputs with a fluent API.

#### Example Usage

```go
// Create a new TextInputBuilder
textinputbuilder := TextInputBuilder{

}
```

#### Type Definition

```go
type TextInputBuilder struct {
}
```

### Constructor Functions

### ParagraphTextInput

ParagraphTextInput creates a paragraph text input.

```go
func ParagraphTextInput(customID, label string) *TextInputBuilder
```

**Parameters:**
- `customID` (string)
- `label` (string)

**Returns:**
- *TextInputBuilder

### ShortTextInput

ShortTextInput creates a short text input.

```go
func ShortTextInput(customID, label string) *TextInputBuilder
```

**Parameters:**
- `customID` (string)
- `label` (string)

**Returns:**
- *TextInputBuilder

### TextInput

TextInput creates a new text input builder.

```go
func (*ModalBuilder) TextInput(customID, label string) *ModalTextInputBuilder
```

**Parameters:**
- `customID` (string)
- `label` (string)

**Returns:**
- *ModalTextInputBuilder

## Methods

### Build

Build returns the built text input wrapped in an action row.

```go
func (*ButtonBuilder) Build() *ButtonComponent
```

**Parameters:**
  None

**Returns:**
- *ButtonComponent

### BuildRaw

BuildRaw returns just the text input component.

```go
func (*TextInputBuilder) BuildRaw() *TextInputComponent
```

**Parameters:**
  None

**Returns:**
- *TextInputComponent

### BuildSafe

BuildSafe validates and returns the built text input wrapped in an action row.

```go
func (*ButtonBuilder) BuildSafe() (*ButtonComponent, error)
```

**Parameters:**
  None

**Returns:**
- *ButtonComponent
- error

### MaxLength

MaxLength sets the maximum length.

```go
func (*ModalTextInputBuilder) MaxLength(max int) *ModalTextInputBuilder
```

**Parameters:**
- `max` (int)

**Returns:**
- *ModalTextInputBuilder

### MinLength

MinLength sets the minimum length.

```go
func (*ModalTextInputBuilder) MinLength(min int) *ModalTextInputBuilder
```

**Parameters:**
- `min` (int)

**Returns:**
- *ModalTextInputBuilder

### Paragraph

Paragraph sets the input to multi-line style.

```go
func (*ModalTextInputBuilder) Paragraph() *ModalTextInputBuilder
```

**Parameters:**
  None

**Returns:**
- *ModalTextInputBuilder

### Placeholder

Placeholder sets the placeholder text.

```go
func (*MentionableSelectBuilder) Placeholder(text string) *MentionableSelectBuilder
```

**Parameters:**
- `text` (string)

**Returns:**
- *MentionableSelectBuilder

### Required

Required sets whether the input is required.

```go
func (*ModalTextInputBuilder) Required(required bool) *ModalTextInputBuilder
```

**Parameters:**
- `required` (bool)

**Returns:**
- *ModalTextInputBuilder

### Short

Short sets the input to single-line style.

```go
func (*ModalTextInputBuilder) Short() *ModalTextInputBuilder
```

**Parameters:**
  None

**Returns:**
- *ModalTextInputBuilder

### Style

Style sets the input style.

```go
func (*ButtonBuilder) Style(style types.ButtonStyle) *ButtonBuilder
```

**Parameters:**
- `style` (types.ButtonStyle)

**Returns:**
- *ButtonBuilder

### Validate

Validate checks if the text input configuration is valid according to Discord limits.

```go
func (*ModalBuilder) Validate() error
```

**Parameters:**
  None

**Returns:**
- error

### Value

Value sets the default value.

```go
func (*ModalTextInputBuilder) Value(value string) *ModalTextInputBuilder
```

**Parameters:**
- `value` (string)

**Returns:**
- *ModalTextInputBuilder

### TextInputComponent
TextInputComponent represents a text input in a modal.

#### Example Usage

```go
// Create a new TextInputComponent
textinputcomponent := TextInputComponent{
    Type: /* value */,
    CustomID: "example",
    Style: /* value */,
    Label: "example",
    MinLength: &42{},
    MaxLength: &42{},
    Required: &true{},
    Value: "example",
    Placeholder: "example",
}
```

#### Type Definition

```go
type TextInputComponent struct {
    Type types.ComponentType `json:"type"`
    CustomID string `json:"custom_id"`
    Style types.TextInputStyle `json:"style"`
    Label string `json:"label"`
    MinLength *int `json:"min_length,omitempty"`
    MaxLength *int `json:"max_length,omitempty"`
    Required *bool `json:"required,omitempty"`
    Value string `json:"value,omitempty"`
    Placeholder string `json:"placeholder,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Type | `types.ComponentType` |  |
| CustomID | `string` |  |
| Style | `types.TextInputStyle` |  |
| Label | `string` |  |
| MinLength | `*int` |  |
| MaxLength | `*int` |  |
| Required | `*bool` |  |
| Value | `string` |  |
| Placeholder | `string` |  |

### UserSelectBuilder
UserSelectBuilder builds user select menus.

#### Example Usage

```go
// Create a new UserSelectBuilder
userselectbuilder := UserSelectBuilder{

}
```

#### Type Definition

```go
type UserSelectBuilder struct {
}
```

### Constructor Functions

### UserSelect

UserSelect creates a new user select builder.

```go
func UserSelect(customID string) *UserSelectBuilder
```

**Parameters:**
- `customID` (string)

**Returns:**
- *UserSelectBuilder

## Methods

### Build

Build returns the built select component.

```go
func (*ButtonBuilder) Build() *ButtonComponent
```

**Parameters:**
  None

**Returns:**
- *ButtonComponent

### BuildSafe

BuildSafe validates and returns the built select component.

```go
func (*ButtonBuilder) BuildSafe() (*ButtonComponent, error)
```

**Parameters:**
  None

**Returns:**
- *ButtonComponent
- error

### Disabled

Disabled sets whether the select is disabled.

```go
func (*ButtonBuilder) Disabled(disabled bool) *ButtonBuilder
```

**Parameters:**
- `disabled` (bool)

**Returns:**
- *ButtonBuilder

### MaxValues

MaxValues sets the maximum number of selections.

```go
func (*MentionableSelectBuilder) MaxValues(max int) *MentionableSelectBuilder
```

**Parameters:**
- `max` (int)

**Returns:**
- *MentionableSelectBuilder

### MinValues

MinValues sets the minimum number of selections.

```go
func (*MentionableSelectBuilder) MinValues(min int) *MentionableSelectBuilder
```

**Parameters:**
- `min` (int)

**Returns:**
- *MentionableSelectBuilder

### Placeholder

Placeholder sets the placeholder text.

```go
func (*ModalTextInputBuilder) Placeholder(text string) *ModalTextInputBuilder
```

**Parameters:**
- `text` (string)

**Returns:**
- *ModalTextInputBuilder

### Validate

Validate checks if the select menu configuration is valid according to Discord limits.

```go
func (*MentionableSelectBuilder) Validate() error
```

**Parameters:**
  None

**Returns:**
- error

### UserSelectComponent
UserSelectComponent represents a select menu for users.

#### Example Usage

```go
// Create a new UserSelectComponent
userselectcomponent := UserSelectComponent{
    Type: /* value */,
    CustomID: "example",
    Placeholder: "example",
    MinValues: &42{},
    MaxValues: &42{},
    Disabled: true,
    DefaultUsers: [],
}
```

#### Type Definition

```go
type UserSelectComponent struct {
    Type types.ComponentType `json:"type"`
    CustomID string `json:"custom_id"`
    Placeholder string `json:"placeholder,omitempty"`
    MinValues *int `json:"min_values,omitempty"`
    MaxValues *int `json:"max_values,omitempty"`
    Disabled bool `json:"disabled,omitempty"`
    DefaultUsers []string `json:"default_values,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Type | `types.ComponentType` |  |
| CustomID | `string` |  |
| Placeholder | `string` |  |
| MinValues | `*int` |  |
| MaxValues | `*int` |  |
| Disabled | `bool` |  |
| DefaultUsers | `[]string` |  |

## Functions

### BuildComponents
BuildComponents takes individual components and wraps them in action rows. Buttons are grouped together (max 5 per row), select menus get their own rows.

```go
func BuildComponents(components ...Component) []any
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `components` | `...Component` | |

**Returns:**
| Type | Description |
|------|-------------|
| `[]any` | |

**Example:**

```go
// Example usage of BuildComponents
result := BuildComponents(/* parameters */)
```

### BuildComponentsSafe
BuildComponentsSafe takes individual components and wraps them in action rows, returning an error if Discord limits are exceeded.

```go
func BuildComponentsSafe(components ...Component) ([]any, error)
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `components` | `...Component` | |

**Returns:**
| Type | Description |
|------|-------------|
| `[]any` | |
| `error` | |

**Example:**

```go
// Example usage of BuildComponentsSafe
result := BuildComponentsSafe(/* parameters */)
```

### ValidateActionRowCount
ValidateActionRowCount checks if the number of action rows is within limits.

```go
func ValidateActionRowCount(count int) error
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `count` | `int` | |

**Returns:**
| Type | Description |
|------|-------------|
| `error` | |

**Example:**

```go
// Example usage of ValidateActionRowCount
result := ValidateActionRowCount(/* parameters */)
```

### ValidateCustomID
ValidateCustomID checks if a custom_id is within the allowed length.

```go
func ValidateCustomID(customID string) error
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `customID` | `string` | |

**Returns:**
| Type | Description |
|------|-------------|
| `error` | |

**Example:**

```go
// Example usage of ValidateCustomID
result := ValidateCustomID(/* parameters */)
```

### ValidateModalInputCount
ValidateModalInputCount checks if the number of modal inputs is within limits.

```go
func ValidateModalInputCount(count int) error
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `count` | `int` | |

**Returns:**
| Type | Description |
|------|-------------|
| `error` | |

**Example:**

```go
// Example usage of ValidateModalInputCount
result := ValidateModalInputCount(/* parameters */)
```

### ValidateSelectOptionCount
ValidateSelectOptionCount checks if the number of select options is within limits.

```go
func ValidateSelectOptionCount(count int) error
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `count` | `int` | |

**Returns:**
| Type | Description |
|------|-------------|
| `error` | |

**Example:**

```go
// Example usage of ValidateSelectOptionCount
result := ValidateSelectOptionCount(/* parameters */)
```

## External Links

- [Package Overview](../packages/components.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/discord/components)
- [Source Code](https://github.com/kolosys/discord/tree/main/components)

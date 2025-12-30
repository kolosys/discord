# types API

Complete API documentation for the types package.

**Import Path:** `github.com/kolosys/discord/types`

## Package Documentation



## Types

### ActivityType
ActivityType represents the type of a user activity.

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

### AllowedMentionType
AllowedMentionType represents types allowed in message mentions.

#### Example Usage

```go
// Example usage of AllowedMentionType
var value AllowedMentionType
// Initialize with appropriate value
```

#### Type Definition

```go
type AllowedMentionType string
```

### ApplicationCommandOptionType
ApplicationCommandOptionType represents the type of a command option.

#### Example Usage

```go
// Example usage of ApplicationCommandOptionType
var value ApplicationCommandOptionType
// Initialize with appropriate value
```

#### Type Definition

```go
type ApplicationCommandOptionType int
```

### ApplicationCommandType
ApplicationCommandType represents the type of an application command.

#### Example Usage

```go
// Example usage of ApplicationCommandType
var value ApplicationCommandType
// Initialize with appropriate value
```

#### Type Definition

```go
type ApplicationCommandType int
```

### AuditLogEvent
AuditLogEvent represents the type of audit log event.

#### Example Usage

```go
// Example usage of AuditLogEvent
var value AuditLogEvent
// Initialize with appropriate value
```

#### Type Definition

```go
type AuditLogEvent int
```

### ButtonStyle
ButtonStyle represents the visual style of a button.

#### Example Usage

```go
// Example usage of ButtonStyle
var value ButtonStyle
// Initialize with appropriate value
```

#### Type Definition

```go
type ButtonStyle int
```

### ChannelType
ChannelType represents the type of a Discord channel.

#### Example Usage

```go
// Example usage of ChannelType
var value ChannelType
// Initialize with appropriate value
```

#### Type Definition

```go
type ChannelType int
```

## Methods

### IsGuild



```go
func (ChannelType) IsGuild() bool
```

**Parameters:**
  None

**Returns:**
- bool

### IsThread



```go
func (ChannelType) IsThread() bool
```

**Parameters:**
  None

**Returns:**
- bool

### IsVoice



```go
func (ChannelType) IsVoice() bool
```

**Parameters:**
  None

**Returns:**
- bool

### Color
Color represents a 24-bit RGB color value.

#### Example Usage

```go
// Example usage of Color
var value Color
// Initialize with appropriate value
```

#### Type Definition

```go
type Color int32
```

### Constructor Functions

### ColorFromHex



```go
func ColorFromHex(hex string) Color
```

**Parameters:**
- `hex` (string)

**Returns:**
- Color

### NewColor



```go
func NewColor(r, g, b uint8) Color
```

**Parameters:**
- `r` (uint8)
- `g` (uint8)
- `b` (uint8)

**Returns:**
- Color

## Methods

### B



```go
func (Color) B() uint8
```

**Parameters:**
  None

**Returns:**
- uint8

### G



```go
func (Color) G() uint8
```

**Parameters:**
  None

**Returns:**
- uint8

### R



```go
func (Color) R() uint8
```

**Parameters:**
  None

**Returns:**
- uint8

### ComponentType
ComponentType represents the type of a message component.

#### Example Usage

```go
// Example usage of ComponentType
var value ComponentType
// Initialize with appropriate value
```

#### Type Definition

```go
type ComponentType int
```

### DefaultMessageNotificationLevel
DefaultMessageNotificationLevel represents guild default notification settings.

#### Example Usage

```go
// Example usage of DefaultMessageNotificationLevel
var value DefaultMessageNotificationLevel
// Initialize with appropriate value
```

#### Type Definition

```go
type DefaultMessageNotificationLevel int
```

### ExplicitContentFilterLevel
ExplicitContentFilterLevel represents a guild's media content filter.

#### Example Usage

```go
// Example usage of ExplicitContentFilterLevel
var value ExplicitContentFilterLevel
// Initialize with appropriate value
```

#### Type Definition

```go
type ExplicitContentFilterLevel int
```

### GuildScheduledEventEntityType
GuildScheduledEventEntityType represents the entity type of a scheduled event.

#### Example Usage

```go
// Example usage of GuildScheduledEventEntityType
var value GuildScheduledEventEntityType
// Initialize with appropriate value
```

#### Type Definition

```go
type GuildScheduledEventEntityType int
```

### GuildScheduledEventPrivacyLevel
GuildScheduledEventPrivacyLevel represents the privacy level of a scheduled event.

#### Example Usage

```go
// Example usage of GuildScheduledEventPrivacyLevel
var value GuildScheduledEventPrivacyLevel
// Initialize with appropriate value
```

#### Type Definition

```go
type GuildScheduledEventPrivacyLevel int
```

### GuildScheduledEventStatus
GuildScheduledEventStatus represents the status of a scheduled event.

#### Example Usage

```go
// Example usage of GuildScheduledEventStatus
var value GuildScheduledEventStatus
// Initialize with appropriate value
```

#### Type Definition

```go
type GuildScheduledEventStatus int
```

### InteractionCallbackType
InteractionCallbackType represents the type of an interaction response.

#### Example Usage

```go
// Example usage of InteractionCallbackType
var value InteractionCallbackType
// Initialize with appropriate value
```

#### Type Definition

```go
type InteractionCallbackType int
```

### InteractionType
InteractionType represents the type of an interaction.

#### Example Usage

```go
// Example usage of InteractionType
var value InteractionType
// Initialize with appropriate value
```

#### Type Definition

```go
type InteractionType int
```

### InviteTargetType
InviteTargetType represents the target type for an invite.

#### Example Usage

```go
// Example usage of InviteTargetType
var value InviteTargetType
// Initialize with appropriate value
```

#### Type Definition

```go
type InviteTargetType int
```

### MFALevel
MFALevel represents a guild's MFA requirement for moderators.

#### Example Usage

```go
// Example usage of MFALevel
var value MFALevel
// Initialize with appropriate value
```

#### Type Definition

```go
type MFALevel int
```

### MessageType
MessageType represents the type of a Discord message.

#### Example Usage

```go
// Example usage of MessageType
var value MessageType
// Initialize with appropriate value
```

#### Type Definition

```go
type MessageType int
```

### NSFWLevel
NSFWLevel represents a guild's NSFW classification.

#### Example Usage

```go
// Example usage of NSFWLevel
var value NSFWLevel
// Initialize with appropriate value
```

#### Type Definition

```go
type NSFWLevel int
```

### NullableTimestamp
NullableTimestamp is a timestamp that can be null in JSON.

#### Example Usage

```go
// Example usage of NullableTimestamp
var value NullableTimestamp
// Initialize with appropriate value
```

#### Type Definition

```go
type NullableTimestamp timestamp.Nullable
```

### OverwriteType
OverwriteType represents the type of a permission overwrite.

#### Example Usage

```go
// Example usage of OverwriteType
var value OverwriteType
// Initialize with appropriate value
```

#### Type Definition

```go
type OverwriteType int
```

### Permissions
Permissions represents a Discord permission bitfield.

#### Example Usage

```go
// Example usage of Permissions
var value Permissions
// Initialize with appropriate value
```

#### Type Definition

```go
type Permissions uint64
```

## Methods

### Add



```go
func (Permissions) Add(perm Permissions) Permissions
```

**Parameters:**
- `perm` (Permissions)

**Returns:**
- Permissions

### Has



```go
func (Permissions) Has(perm Permissions) bool
```

**Parameters:**
- `perm` (Permissions)

**Returns:**
- bool

### MarshalJSON



```go
func (Permissions) MarshalJSON() ([]byte, error)
```

**Parameters:**
  None

**Returns:**
- []byte
- error

### Remove



```go
func (Permissions) Remove(perm Permissions) Permissions
```

**Parameters:**
- `perm` (Permissions)

**Returns:**
- Permissions

### Toggle



```go
func (Permissions) Toggle(perm Permissions) Permissions
```

**Parameters:**
- `perm` (Permissions)

**Returns:**
- Permissions

### UnmarshalJSON



```go
func (*Permissions) UnmarshalJSON(data []byte) error
```

**Parameters:**
- `data` ([]byte)

**Returns:**
- error

### PremiumTier
PremiumTier represents a guild's boost level.

#### Example Usage

```go
// Example usage of PremiumTier
var value PremiumTier
// Initialize with appropriate value
```

#### Type Definition

```go
type PremiumTier int
```

### PremiumType
PremiumType represents a user's Nitro subscription type.

#### Example Usage

```go
// Example usage of PremiumType
var value PremiumType
// Initialize with appropriate value
```

#### Type Definition

```go
type PremiumType int
```

### Snowflake
Snowflake is a Discord snowflake ID.

#### Example Usage

```go
// Example usage of Snowflake
var value Snowflake
// Initialize with appropriate value
```

#### Type Definition

```go
type Snowflake string
```

### StatusType
StatusType represents a user's online status.

#### Example Usage

```go
// Example usage of StatusType
var value StatusType
// Initialize with appropriate value
```

#### Type Definition

```go
type StatusType string
```

### StickerFormatType
StickerFormatType represents the format of a sticker.

#### Example Usage

```go
// Example usage of StickerFormatType
var value StickerFormatType
// Initialize with appropriate value
```

#### Type Definition

```go
type StickerFormatType int
```

### StickerType
StickerType represents the type of a sticker.

#### Example Usage

```go
// Example usage of StickerType
var value StickerType
// Initialize with appropriate value
```

#### Type Definition

```go
type StickerType int
```

### TextInputStyle
TextInputStyle represents the visual style of a text input.

#### Example Usage

```go
// Example usage of TextInputStyle
var value TextInputStyle
// Initialize with appropriate value
```

#### Type Definition

```go
type TextInputStyle int
```

### Timestamp
Timestamp is an ISO8601/RFC3339 timestamp.

#### Example Usage

```go
// Example usage of Timestamp
var value Timestamp
// Initialize with appropriate value
```

#### Type Definition

```go
type Timestamp timestamp.Timestamp
```

### VerificationLevel
VerificationLevel represents a guild's verification requirements.

#### Example Usage

```go
// Example usage of VerificationLevel
var value VerificationLevel
// Initialize with appropriate value
```

#### Type Definition

```go
type VerificationLevel int
```

### WebhookType
WebhookType represents the type of a webhook.

#### Example Usage

```go
// Example usage of WebhookType
var value WebhookType
// Initialize with appropriate value
```

#### Type Definition

```go
type WebhookType int
```

## External Links

- [Package Overview](../packages/types.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/discord/types)
- [Source Code](https://github.com/kolosys/discord/tree/main/types)

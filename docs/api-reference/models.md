# models API

Complete API documentation for the models package.

**Import Path:** `github.com/kolosys/discord/models`

## Package Documentation



## Types

### Account
Account

#### Example Usage

```go
// Create a new Account
account := Account{
    ID: "example",
    Name: &"example"{},
}
```

#### Type Definition

```go
type Account struct {
    ID string `json:"id"`
    Name *string `json:"name,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Name | `*string` |  |

### ActionRowComponent
ActionRowComponent

#### Example Usage

```go
// Create a new ActionRowComponent
actionrowcomponent := ActionRowComponent{
    ID: 42,
    Components: [],
    Type: 42,
}
```

#### Type Definition

```go
type ActionRowComponent struct {
    ID int32 `json:"id"`
    Components []any `json:"components"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| Components | `[]any` |  |
| Type | `int32` |  |

### ActionRowComponentForMessageOptions
ActionRowComponentForMessageOptions

#### Example Usage

```go
// Create a new ActionRowComponentForMessageOptions
actionrowcomponentformessageoptions := ActionRowComponentForMessageOptions{
    ID: &42{},
    Components: [],
    Type: 42,
}
```

#### Type Definition

```go
type ActionRowComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    Components []any `json:"components"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| Components | `[]any` |  |
| Type | `int32` |  |

### ActionRowComponentForModalOptions
ActionRowComponentForModalOptions

#### Example Usage

```go
// Create a new ActionRowComponentForModalOptions
actionrowcomponentformodaloptions := ActionRowComponentForModalOptions{
    ID: &42{},
    Components: [],
    Type: 42,
}
```

#### Type Definition

```go
type ActionRowComponentForModalOptions struct {
    ID *int32 `json:"id,omitempty"`
    Components []any `json:"components"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| Components | `[]any` |  |
| Type | `int32` |  |

### ActivitiesAttachment
ActivitiesAttachment

#### Example Usage

```go
// Create a new ActivitiesAttachment
activitiesattachment := ActivitiesAttachment{
    Attachment: any{},
}
```

#### Type Definition

```go
type ActivitiesAttachment struct {
    Attachment any `json:"attachment"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Attachment | `any` |  |

### AfkTimeouts
AfkTimeouts

#### Example Usage

```go
// Example usage of AfkTimeouts
var value AfkTimeouts
// Initialize with appropriate value
```

#### Type Definition

```go
type AfkTimeouts int
```

### AllowedMentionTypes
AllowedMentionTypes

#### Example Usage

```go
// Example usage of AllowedMentionTypes
var value AllowedMentionTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type AllowedMentionTypes string
```

### Application
Application

#### Example Usage

```go
// Create a new Application
application := Application{
    ID: "example",
    Bot: any{},
    BotPublic: true,
    BotRequireCodeGrant: true,
    CoverImage: "example",
    CustomInstallURL: "example",
    Description: "example",
    Flags: 42,
    GuildID: "example",
    Icon: &"example"{},
    InstallParams: any{},
    IntegrationTypesConfig: map[],
    MaxParticipants: &42{},
    Name: "example",
    PrimarySkuID: "example",
    PrivacyPolicyURL: "example",
    RpcOrigins: [],
    Slug: "example",
    Tags: [],
    TermsOfServiceURL: "example",
    Type: any{},
    VerifyKey: "example",
}
```

#### Type Definition

```go
type Application struct {
    ID string `json:"id"`
    Bot any `json:"bot,omitempty"`
    BotPublic bool `json:"bot_public,omitempty"`
    BotRequireCodeGrant bool `json:"bot_require_code_grant,omitempty"`
    CoverImage string `json:"cover_image,omitempty"`
    CustomInstallURL string `json:"custom_install_url,omitempty"`
    Description string `json:"description"`
    Flags int32 `json:"flags"`
    GuildID string `json:"guild_id,omitempty"`
    Icon *string `json:"icon,omitempty"`
    InstallParams any `json:"install_params,omitempty"`
    IntegrationTypesConfig map[string]any `json:"integration_types_config,omitempty"`
    MaxParticipants *int32 `json:"max_participants,omitempty"`
    Name string `json:"name"`
    PrimarySkuID string `json:"primary_sku_id,omitempty"`
    PrivacyPolicyURL string `json:"privacy_policy_url,omitempty"`
    RpcOrigins []*string `json:"rpc_origins,omitempty"`
    Slug string `json:"slug,omitempty"`
    Tags []string `json:"tags,omitempty"`
    TermsOfServiceURL string `json:"terms_of_service_url,omitempty"`
    Type any `json:"type,omitempty"`
    VerifyKey string `json:"verify_key"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Bot | `any` |  |
| BotPublic | `bool` |  |
| BotRequireCodeGrant | `bool` |  |
| CoverImage | `string` |  |
| CustomInstallURL | `string` |  |
| Description | `string` |  |
| Flags | `int32` |  |
| GuildID | `string` |  |
| Icon | `*string` |  |
| InstallParams | `any` |  |
| IntegrationTypesConfig | `map[string]any` |  |
| MaxParticipants | `*int32` |  |
| Name | `string` |  |
| PrimarySkuID | `string` |  |
| PrivacyPolicyURL | `string` |  |
| RpcOrigins | `[]*string` |  |
| Slug | `string` |  |
| Tags | `[]string` |  |
| TermsOfServiceURL | `string` |  |
| Type | `any` |  |
| VerifyKey | `string` |  |

### ApplicationCommand
ApplicationCommand

#### Example Usage

```go
// Create a new ApplicationCommand
applicationcommand := ApplicationCommand{
    ID: "example",
    ApplicationID: "example",
    Contexts: [],
    DefaultMemberPermissions: &"example"{},
    Description: "example",
    DescriptionLocalizations: map[],
    DescriptionLocalized: "example",
    DmPermission: true,
    GuildID: "example",
    IntegrationTypes: [],
    Name: "example",
    NameLocalizations: map[],
    NameLocalized: "example",
    Nsfw: true,
    Options: [],
    Type: 42,
    Version: "example",
}
```

#### Type Definition

```go
type ApplicationCommand struct {
    ID string `json:"id"`
    ApplicationID string `json:"application_id"`
    Contexts []int32 `json:"contexts,omitempty"`
    DefaultMemberPermissions *string `json:"default_member_permissions,omitempty"`
    Description string `json:"description"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    DescriptionLocalized string `json:"description_localized,omitempty"`
    DmPermission bool `json:"dm_permission,omitempty"`
    GuildID string `json:"guild_id,omitempty"`
    IntegrationTypes []int32 `json:"integration_types,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    NameLocalized string `json:"name_localized,omitempty"`
    Nsfw bool `json:"nsfw,omitempty"`
    Options []any `json:"options,omitempty"`
    Type int32 `json:"type"`
    Version string `json:"version"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ApplicationID | `string` |  |
| Contexts | `[]int32` |  |
| DefaultMemberPermissions | `*string` |  |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| DescriptionLocalized | `string` |  |
| DmPermission | `bool` |  |
| GuildID | `string` |  |
| IntegrationTypes | `[]int32` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| NameLocalized | `string` |  |
| Nsfw | `bool` |  |
| Options | `[]any` |  |
| Type | `int32` |  |
| Version | `string` |  |

### ApplicationCommandAttachmentOption
ApplicationCommandAttachmentOption

#### Example Usage

```go
// Create a new ApplicationCommandAttachmentOption
applicationcommandattachmentoption := ApplicationCommandAttachmentOption{
    Description: "example",
    DescriptionLocalizations: map[],
    Name: "example",
    NameLocalizations: map[],
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationCommandAttachmentOption struct {
    Description string `json:"description"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### ApplicationCommandAutocompleteCallbackOptions
ApplicationCommandAutocompleteCallbackOptions

#### Example Usage

```go
// Create a new ApplicationCommandAutocompleteCallbackOptions
applicationcommandautocompletecallbackoptions := ApplicationCommandAutocompleteCallbackOptions{
    Data: any{},
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationCommandAutocompleteCallbackOptions struct {
    Data any `json:"data"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Data | `any` |  |
| Type | `int32` |  |

### ApplicationCommandBooleanOption
ApplicationCommandBooleanOption

#### Example Usage

```go
// Create a new ApplicationCommandBooleanOption
applicationcommandbooleanoption := ApplicationCommandBooleanOption{
    Description: "example",
    DescriptionLocalizations: map[],
    DescriptionLocalized: "example",
    Name: "example",
    NameLocalizations: map[],
    NameLocalized: "example",
    Required: true,
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationCommandBooleanOption struct {
    Description string `json:"description"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    DescriptionLocalized string `json:"description_localized,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    NameLocalized string `json:"name_localized,omitempty"`
    Required bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| DescriptionLocalized | `string` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| NameLocalized | `string` |  |
| Required | `bool` |  |
| Type | `int32` |  |

### ApplicationCommandChannelOption
ApplicationCommandChannelOption

#### Example Usage

```go
// Create a new ApplicationCommandChannelOption
applicationcommandchanneloption := ApplicationCommandChannelOption{
    ChannelTypes: [],
    Description: "example",
    DescriptionLocalizations: map[],
    Name: "example",
    NameLocalizations: map[],
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationCommandChannelOption struct {
    ChannelTypes []int32 `json:"channel_types,omitempty"`
    Description string `json:"description"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelTypes | `[]int32` |  |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### ApplicationCommandCreateOptions
ApplicationCommandCreateOptions

#### Example Usage

```go
// Create a new ApplicationCommandCreateOptions
applicationcommandcreateoptions := ApplicationCommandCreateOptions{
    Contexts: [],
    DefaultMemberPermissions: &42{},
    Description: &"example"{},
    DescriptionLocalizations: map[],
    DmPermission: &true{},
    Handler: any{},
    IntegrationTypes: [],
    Name: "example",
    NameLocalizations: map[],
    Options: [],
    Type: any{},
}
```

#### Type Definition

```go
type ApplicationCommandCreateOptions struct {
    Contexts []int32 `json:"contexts,omitempty"`
    DefaultMemberPermissions *int64 `json:"default_member_permissions,omitempty"`
    Description *string `json:"description,omitempty"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    DmPermission *bool `json:"dm_permission,omitempty"`
    Handler any `json:"handler,omitempty"`
    IntegrationTypes []int32 `json:"integration_types,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    Options []any `json:"options,omitempty"`
    Type any `json:"type,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Contexts | `[]int32` |  |
| DefaultMemberPermissions | `*int64` |  |
| Description | `*string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| DmPermission | `*bool` |  |
| Handler | `any` |  |
| IntegrationTypes | `[]int32` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| Options | `[]any` |  |
| Type | `any` |  |

### ApplicationCommandHandler
ApplicationCommandHandler

#### Example Usage

```go
// Create a new ApplicationCommandHandler
applicationcommandhandler := ApplicationCommandHandler{

}
```

#### Type Definition

```go
type ApplicationCommandHandler struct {
}
```

### ApplicationCommandIntegerOption
ApplicationCommandIntegerOption

#### Example Usage

```go
// Create a new ApplicationCommandIntegerOption
applicationcommandintegeroption := ApplicationCommandIntegerOption{
    Autocomplete: true,
    Choices: [],
    Description: "example",
    DescriptionLocalizations: map[],
    DescriptionLocalized: "example",
    MaxValue: 42,
    MinValue: 42,
    Name: "example",
    NameLocalizations: map[],
    NameLocalized: "example",
    Required: true,
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationCommandIntegerOption struct {
    Autocomplete bool `json:"autocomplete,omitempty"`
    Choices []any `json:"choices,omitempty"`
    Description string `json:"description"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    DescriptionLocalized string `json:"description_localized,omitempty"`
    MaxValue int64 `json:"max_value,omitempty"`
    MinValue int64 `json:"min_value,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    NameLocalized string `json:"name_localized,omitempty"`
    Required bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Autocomplete | `bool` |  |
| Choices | `[]any` |  |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| DescriptionLocalized | `string` |  |
| MaxValue | `int64` |  |
| MinValue | `int64` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| NameLocalized | `string` |  |
| Required | `bool` |  |
| Type | `int32` |  |

### ApplicationCommandInteractionMetadata
ApplicationCommandInteractionMetadata

#### Example Usage

```go
// Create a new ApplicationCommandInteractionMetadata
applicationcommandinteractionmetadata := ApplicationCommandInteractionMetadata{
    ID: "example",
    AuthorizingIntegrationOwners: map[],
    OriginalResponseMessageID: "example",
    TargetMessageID: "example",
    TargetUser: any{},
    Type: 42,
    User: any{},
}
```

#### Type Definition

```go
type ApplicationCommandInteractionMetadata struct {
    ID string `json:"id"`
    AuthorizingIntegrationOwners map[string]string `json:"authorizing_integration_owners"`
    OriginalResponseMessageID string `json:"original_response_message_id,omitempty"`
    TargetMessageID string `json:"target_message_id,omitempty"`
    TargetUser any `json:"target_user,omitempty"`
    Type int32 `json:"type"`
    User any `json:"user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| AuthorizingIntegrationOwners | `map[string]string` |  |
| OriginalResponseMessageID | `string` |  |
| TargetMessageID | `string` |  |
| TargetUser | `any` |  |
| Type | `int32` |  |
| User | `any` |  |

### ApplicationCommandMentionableOption
ApplicationCommandMentionableOption

#### Example Usage

```go
// Create a new ApplicationCommandMentionableOption
applicationcommandmentionableoption := ApplicationCommandMentionableOption{
    Description: "example",
    DescriptionLocalizations: map[],
    Name: "example",
    NameLocalizations: map[],
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationCommandMentionableOption struct {
    Description string `json:"description"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### ApplicationCommandNumberOption
ApplicationCommandNumberOption

#### Example Usage

```go
// Create a new ApplicationCommandNumberOption
applicationcommandnumberoption := ApplicationCommandNumberOption{
    Autocomplete: true,
    Choices: [],
    Description: "example",
    DescriptionLocalizations: map[],
    DescriptionLocalized: "example",
    MaxValue: 3.14,
    MinValue: 3.14,
    Name: "example",
    NameLocalizations: map[],
    NameLocalized: "example",
    Required: true,
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationCommandNumberOption struct {
    Autocomplete bool `json:"autocomplete,omitempty"`
    Choices []any `json:"choices,omitempty"`
    Description string `json:"description"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    DescriptionLocalized string `json:"description_localized,omitempty"`
    MaxValue float64 `json:"max_value,omitempty"`
    MinValue float64 `json:"min_value,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    NameLocalized string `json:"name_localized,omitempty"`
    Required bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Autocomplete | `bool` |  |
| Choices | `[]any` |  |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| DescriptionLocalized | `string` |  |
| MaxValue | `float64` |  |
| MinValue | `float64` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| NameLocalized | `string` |  |
| Required | `bool` |  |
| Type | `int32` |  |

### ApplicationCommandOptionIntegerChoice
ApplicationCommandOptionIntegerChoice

#### Example Usage

```go
// Create a new ApplicationCommandOptionIntegerChoice
applicationcommandoptionintegerchoice := ApplicationCommandOptionIntegerChoice{
    Name: "example",
    NameLocalizations: map[],
    NameLocalized: "example",
    Value: 42,
}
```

#### Type Definition

```go
type ApplicationCommandOptionIntegerChoice struct {
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    NameLocalized string `json:"name_localized,omitempty"`
    Value int64 `json:"value"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| NameLocalized | `string` |  |
| Value | `int64` |  |

### ApplicationCommandOptionNumberChoice
ApplicationCommandOptionNumberChoice

#### Example Usage

```go
// Create a new ApplicationCommandOptionNumberChoice
applicationcommandoptionnumberchoice := ApplicationCommandOptionNumberChoice{
    Name: "example",
    NameLocalizations: map[],
    NameLocalized: "example",
    Value: 3.14,
}
```

#### Type Definition

```go
type ApplicationCommandOptionNumberChoice struct {
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    NameLocalized string `json:"name_localized,omitempty"`
    Value float64 `json:"value"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| NameLocalized | `string` |  |
| Value | `float64` |  |

### ApplicationCommandOptionStringChoice
ApplicationCommandOptionStringChoice

#### Example Usage

```go
// Create a new ApplicationCommandOptionStringChoice
applicationcommandoptionstringchoice := ApplicationCommandOptionStringChoice{
    Name: "example",
    NameLocalizations: map[],
    Value: "example",
}
```

#### Type Definition

```go
type ApplicationCommandOptionStringChoice struct {
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    Value string `json:"value"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| Value | `string` |  |

### ApplicationCommandOptionType
ApplicationCommandOptionType

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

### ApplicationCommandPatchOptionsPartial
ApplicationCommandPatchOptionsPartial

#### Example Usage

```go
// Create a new ApplicationCommandPatchOptionsPartial
applicationcommandpatchoptionspartial := ApplicationCommandPatchOptionsPartial{
    Contexts: [],
    DefaultMemberPermissions: &42{},
    Description: &"example"{},
    DescriptionLocalizations: map[],
    DmPermission: &true{},
    Handler: any{},
    IntegrationTypes: [],
    Name: "example",
    NameLocalizations: map[],
    Options: [],
}
```

#### Type Definition

```go
type ApplicationCommandPatchOptionsPartial struct {
    Contexts []int32 `json:"contexts,omitempty"`
    DefaultMemberPermissions *int64 `json:"default_member_permissions,omitempty"`
    Description *string `json:"description,omitempty"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    DmPermission *bool `json:"dm_permission,omitempty"`
    Handler any `json:"handler,omitempty"`
    IntegrationTypes []int32 `json:"integration_types,omitempty"`
    Name string `json:"name,omitempty"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    Options []any `json:"options,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Contexts | `[]int32` |  |
| DefaultMemberPermissions | `*int64` |  |
| Description | `*string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| DmPermission | `*bool` |  |
| Handler | `any` |  |
| IntegrationTypes | `[]int32` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| Options | `[]any` |  |

### ApplicationCommandPermission
ApplicationCommandPermission

#### Example Usage

```go
// Create a new ApplicationCommandPermission
applicationcommandpermission := ApplicationCommandPermission{
    ID: "example",
    Permission: true,
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationCommandPermission struct {
    ID string `json:"id"`
    Permission bool `json:"permission"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Permission | `bool` |  |
| Type | `int32` |  |

### ApplicationCommandPermissionType
ApplicationCommandPermissionType

#### Example Usage

```go
// Example usage of ApplicationCommandPermissionType
var value ApplicationCommandPermissionType
// Initialize with appropriate value
```

#### Type Definition

```go
type ApplicationCommandPermissionType int
```

### ApplicationCommandRoleOption
ApplicationCommandRoleOption

#### Example Usage

```go
// Create a new ApplicationCommandRoleOption
applicationcommandroleoption := ApplicationCommandRoleOption{
    Description: "example",
    DescriptionLocalizations: map[],
    Name: "example",
    NameLocalizations: map[],
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationCommandRoleOption struct {
    Description string `json:"description"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### ApplicationCommandStringOption
ApplicationCommandStringOption

#### Example Usage

```go
// Create a new ApplicationCommandStringOption
applicationcommandstringoption := ApplicationCommandStringOption{
    Autocomplete: true,
    Choices: [],
    Description: "example",
    DescriptionLocalizations: map[],
    DescriptionLocalized: "example",
    MaxLength: 42,
    MinLength: 42,
    Name: "example",
    NameLocalizations: map[],
    NameLocalized: "example",
    Required: true,
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationCommandStringOption struct {
    Autocomplete bool `json:"autocomplete,omitempty"`
    Choices []any `json:"choices,omitempty"`
    Description string `json:"description"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    DescriptionLocalized string `json:"description_localized,omitempty"`
    MaxLength int32 `json:"max_length,omitempty"`
    MinLength int32 `json:"min_length,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    NameLocalized string `json:"name_localized,omitempty"`
    Required bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Autocomplete | `bool` |  |
| Choices | `[]any` |  |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| DescriptionLocalized | `string` |  |
| MaxLength | `int32` |  |
| MinLength | `int32` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| NameLocalized | `string` |  |
| Required | `bool` |  |
| Type | `int32` |  |

### ApplicationCommandSubcommandGroupOption
ApplicationCommandSubcommandGroupOption

#### Example Usage

```go
// Create a new ApplicationCommandSubcommandGroupOption
applicationcommandsubcommandgroupoption := ApplicationCommandSubcommandGroupOption{
    Description: "example",
    DescriptionLocalizations: map[],
    DescriptionLocalized: "example",
    Name: "example",
    NameLocalizations: map[],
    NameLocalized: "example",
    Options: [],
    Required: true,
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationCommandSubcommandGroupOption struct {
    Description string `json:"description"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    DescriptionLocalized string `json:"description_localized,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    NameLocalized string `json:"name_localized,omitempty"`
    Options []any `json:"options,omitempty"`
    Required bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| DescriptionLocalized | `string` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| NameLocalized | `string` |  |
| Options | `[]any` |  |
| Required | `bool` |  |
| Type | `int32` |  |

### ApplicationCommandSubcommandOption
ApplicationCommandSubcommandOption

#### Example Usage

```go
// Create a new ApplicationCommandSubcommandOption
applicationcommandsubcommandoption := ApplicationCommandSubcommandOption{
    Description: "example",
    DescriptionLocalizations: map[],
    Name: "example",
    NameLocalizations: map[],
    Options: [],
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationCommandSubcommandOption struct {
    Description string `json:"description"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    Options []any `json:"options,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| Options | `[]any` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### ApplicationCommandType
ApplicationCommandType

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

### ApplicationCommandUpdateOptions
ApplicationCommandUpdateOptions

#### Example Usage

```go
// Create a new ApplicationCommandUpdateOptions
applicationcommandupdateoptions := ApplicationCommandUpdateOptions{
    ID: any{},
    Contexts: [],
    DefaultMemberPermissions: &42{},
    Description: &"example"{},
    DescriptionLocalizations: map[],
    DmPermission: &true{},
    Handler: any{},
    IntegrationTypes: [],
    Name: "example",
    NameLocalizations: map[],
    Options: [],
    Type: any{},
}
```

#### Type Definition

```go
type ApplicationCommandUpdateOptions struct {
    ID any `json:"id,omitempty"`
    Contexts []int32 `json:"contexts,omitempty"`
    DefaultMemberPermissions *int64 `json:"default_member_permissions,omitempty"`
    Description *string `json:"description,omitempty"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    DmPermission *bool `json:"dm_permission,omitempty"`
    Handler any `json:"handler,omitempty"`
    IntegrationTypes []int32 `json:"integration_types,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    Options []any `json:"options,omitempty"`
    Type any `json:"type,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `any` |  |
| Contexts | `[]int32` |  |
| DefaultMemberPermissions | `*int64` |  |
| Description | `*string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| DmPermission | `*bool` |  |
| Handler | `any` |  |
| IntegrationTypes | `[]int32` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| Options | `[]any` |  |
| Type | `any` |  |

### ApplicationCommandUserOption
ApplicationCommandUserOption

#### Example Usage

```go
// Create a new ApplicationCommandUserOption
applicationcommanduseroption := ApplicationCommandUserOption{
    Description: "example",
    DescriptionLocalizations: map[],
    DescriptionLocalized: "example",
    Name: "example",
    NameLocalizations: map[],
    NameLocalized: "example",
    Required: true,
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationCommandUserOption struct {
    Description string `json:"description"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    DescriptionLocalized string `json:"description_localized,omitempty"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    NameLocalized string `json:"name_localized,omitempty"`
    Required bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| DescriptionLocalized | `string` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| NameLocalized | `string` |  |
| Required | `bool` |  |
| Type | `int32` |  |

### ApplicationExplicitContentFilterTypes
ApplicationExplicitContentFilterTypes

#### Example Usage

```go
// Example usage of ApplicationExplicitContentFilterTypes
var value ApplicationExplicitContentFilterTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type ApplicationExplicitContentFilterTypes int
```

### ApplicationFormPartial
ApplicationFormPartial

#### Example Usage

```go
// Create a new ApplicationFormPartial
applicationformpartial := ApplicationFormPartial{
    CoverImage: &"example"{},
    CustomInstallURL: &"example"{},
    Description: any{},
    ExplicitContentFilter: any{},
    Flags: &42{},
    Icon: &"example"{},
    InstallParams: any{},
    IntegrationTypesConfig: map[],
    InteractionsEndpointURL: &"example"{},
    MaxParticipants: &42{},
    RoleConnectionsVerificationURL: &"example"{},
    Tags: [],
    TeamID: any{},
    Type: any{},
}
```

#### Type Definition

```go
type ApplicationFormPartial struct {
    CoverImage *string `json:"cover_image,omitempty"`
    CustomInstallURL *string `json:"custom_install_url,omitempty"`
    Description any `json:"description,omitempty"`
    ExplicitContentFilter any `json:"explicit_content_filter,omitempty"`
    Flags *int64 `json:"flags,omitempty"`
    Icon *string `json:"icon,omitempty"`
    InstallParams any `json:"install_params,omitempty"`
    IntegrationTypesConfig map[string]any `json:"integration_types_config,omitempty"`
    InteractionsEndpointURL *string `json:"interactions_endpoint_url,omitempty"`
    MaxParticipants *int32 `json:"max_participants,omitempty"`
    RoleConnectionsVerificationURL *string `json:"role_connections_verification_url,omitempty"`
    Tags []string `json:"tags,omitempty"`
    TeamID any `json:"team_id,omitempty"`
    Type any `json:"type,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| CoverImage | `*string` |  |
| CustomInstallURL | `*string` |  |
| Description | `any` |  |
| ExplicitContentFilter | `any` |  |
| Flags | `*int64` |  |
| Icon | `*string` |  |
| InstallParams | `any` |  |
| IntegrationTypesConfig | `map[string]any` |  |
| InteractionsEndpointURL | `*string` |  |
| MaxParticipants | `*int32` |  |
| RoleConnectionsVerificationURL | `*string` |  |
| Tags | `[]string` |  |
| TeamID | `any` |  |
| Type | `any` |  |

### ApplicationIdentityProviderAuthType
ApplicationIdentityProviderAuthType

#### Example Usage

```go
// Example usage of ApplicationIdentityProviderAuthType
var value ApplicationIdentityProviderAuthType
// Initialize with appropriate value
```

#### Type Definition

```go
type ApplicationIdentityProviderAuthType string
```

### ApplicationIncomingWebhook
ApplicationIncomingWebhook

#### Example Usage

```go
// Create a new ApplicationIncomingWebhook
applicationincomingwebhook := ApplicationIncomingWebhook{
    ID: "example",
    ApplicationID: any{},
    Avatar: &"example"{},
    ChannelID: any{},
    GuildID: any{},
    Name: "example",
    Type: 42,
    User: any{},
}
```

#### Type Definition

```go
type ApplicationIncomingWebhook struct {
    ID string `json:"id"`
    ApplicationID any `json:"application_id,omitempty"`
    Avatar *string `json:"avatar,omitempty"`
    ChannelID any `json:"channel_id,omitempty"`
    GuildID any `json:"guild_id,omitempty"`
    Name string `json:"name"`
    Type int32 `json:"type"`
    User any `json:"user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ApplicationID | `any` |  |
| Avatar | `*string` |  |
| ChannelID | `any` |  |
| GuildID | `any` |  |
| Name | `string` |  |
| Type | `int32` |  |
| User | `any` |  |

### ApplicationIntegrationType
ApplicationIntegrationType

#### Example Usage

```go
// Example usage of ApplicationIntegrationType
var value ApplicationIntegrationType
// Initialize with appropriate value
```

#### Type Definition

```go
type ApplicationIntegrationType int
```

### ApplicationIntegrationTypeConfiguration
ApplicationIntegrationTypeConfiguration

#### Example Usage

```go
// Create a new ApplicationIntegrationTypeConfiguration
applicationintegrationtypeconfiguration := ApplicationIntegrationTypeConfiguration{
    Oauth2InstallParams: any{},
}
```

#### Type Definition

```go
type ApplicationIntegrationTypeConfiguration struct {
    Oauth2InstallParams any `json:"oauth2_install_params,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Oauth2InstallParams | `any` |  |

### ApplicationOAuth2InstallParams
ApplicationOAuth2InstallParams

#### Example Usage

```go
// Create a new ApplicationOAuth2InstallParams
applicationoauth2installparams := ApplicationOAuth2InstallParams{
    Permissions: &42{},
    Scopes: [],
}
```

#### Type Definition

```go
type ApplicationOAuth2InstallParams struct {
    Permissions *int64 `json:"permissions,omitempty"`
    Scopes []string `json:"scopes,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Permissions | `*int64` |  |
| Scopes | `[]string` |  |

### ApplicationRoleConnectionsMetadataItem
ApplicationRoleConnectionsMetadataItem

#### Example Usage

```go
// Create a new ApplicationRoleConnectionsMetadataItem
applicationroleconnectionsmetadataitem := ApplicationRoleConnectionsMetadataItem{
    Description: "example",
    DescriptionLocalizations: map[],
    Key: "example",
    Name: "example",
    NameLocalizations: map[],
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationRoleConnectionsMetadataItem struct {
    Description string `json:"description"`
    DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
    Key string `json:"key"`
    Name string `json:"name"`
    NameLocalizations map[string]string `json:"name_localizations,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]string` |  |
| Key | `string` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]string` |  |
| Type | `int32` |  |

### ApplicationRoleConnectionsMetadataItemOptions
ApplicationRoleConnectionsMetadataItemOptions

#### Example Usage

```go
// Create a new ApplicationRoleConnectionsMetadataItemOptions
applicationroleconnectionsmetadataitemoptions := ApplicationRoleConnectionsMetadataItemOptions{
    Description: "example",
    DescriptionLocalizations: map[],
    Key: "example",
    Name: "example",
    NameLocalizations: map[],
    Type: 42,
}
```

#### Type Definition

```go
type ApplicationRoleConnectionsMetadataItemOptions struct {
    Description string `json:"description"`
    DescriptionLocalizations map[string]*string `json:"description_localizations,omitempty"`
    Key string `json:"key"`
    Name string `json:"name"`
    NameLocalizations map[string]*string `json:"name_localizations,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `string` |  |
| DescriptionLocalizations | `map[string]*string` |  |
| Key | `string` |  |
| Name | `string` |  |
| NameLocalizations | `map[string]*string` |  |
| Type | `int32` |  |

### ApplicationTypes
ApplicationTypes

#### Example Usage

```go
// Example usage of ApplicationTypes
var value ApplicationTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type ApplicationTypes int
```

### ApplicationUserRoleConnection
ApplicationUserRoleConnection

#### Example Usage

```go
// Create a new ApplicationUserRoleConnection
applicationuserroleconnection := ApplicationUserRoleConnection{
    Metadata: map[],
    PlatformName: &"example"{},
    PlatformUsername: &"example"{},
}
```

#### Type Definition

```go
type ApplicationUserRoleConnection struct {
    Metadata map[string]string `json:"metadata,omitempty"`
    PlatformName *string `json:"platform_name,omitempty"`
    PlatformUsername *string `json:"platform_username,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Metadata | `map[string]string` |  |
| PlatformName | `*string` |  |
| PlatformUsername | `*string` |  |

### Attachment
Attachment

#### Example Usage

```go
// Create a new Attachment
attachment := Attachment{
    ID: "example",
    Application: any{},
    ClipCreatedAt: /* value */,
    ClipParticipants: [],
    ContentType: "example",
    Description: "example",
    DurationSecs: 3.14,
    Ephemeral: true,
    Filename: "example",
    Height: 42,
    ProxyURL: "example",
    Size: 42,
    Title: &"example"{},
    URL: "example",
    Waveform: "example",
    Width: 42,
}
```

#### Type Definition

```go
type Attachment struct {
    ID string `json:"id"`
    Application any `json:"application,omitempty"`
    ClipCreatedAt time.Time `json:"clip_created_at,omitempty"`
    ClipParticipants []any `json:"clip_participants,omitempty"`
    ContentType string `json:"content_type,omitempty"`
    Description string `json:"description,omitempty"`
    DurationSecs float64 `json:"duration_secs,omitempty"`
    Ephemeral bool `json:"ephemeral,omitempty"`
    Filename string `json:"filename"`
    Height int32 `json:"height,omitempty"`
    ProxyURL string `json:"proxy_url"`
    Size int32 `json:"size"`
    Title *string `json:"title,omitempty"`
    URL string `json:"url"`
    Waveform string `json:"waveform,omitempty"`
    Width int32 `json:"width,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Application | `any` |  |
| ClipCreatedAt | `time.Time` |  |
| ClipParticipants | `[]any` |  |
| ContentType | `string` |  |
| Description | `string` |  |
| DurationSecs | `float64` |  |
| Ephemeral | `bool` |  |
| Filename | `string` |  |
| Height | `int32` |  |
| ProxyURL | `string` |  |
| Size | `int32` |  |
| Title | `*string` |  |
| URL | `string` |  |
| Waveform | `string` |  |
| Width | `int32` |  |

### AuditLogActionTypes
AuditLogActionTypes

#### Example Usage

```go
// Example usage of AuditLogActionTypes
var value AuditLogActionTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type AuditLogActionTypes int
```

### AuditLogEntry
AuditLogEntry

#### Example Usage

```go
// Create a new AuditLogEntry
auditlogentry := AuditLogEntry{
    ID: "example",
    ActionType: 42,
    Changes: [],
    Options: map[],
    Reason: "example",
    TargetID: any{},
    UserID: any{},
}
```

#### Type Definition

```go
type AuditLogEntry struct {
    ID string `json:"id"`
    ActionType int32 `json:"action_type"`
    Changes []any `json:"changes,omitempty"`
    Options map[string]string `json:"options,omitempty"`
    Reason string `json:"reason,omitempty"`
    TargetID any `json:"target_id,omitempty"`
    UserID any `json:"user_id,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ActionType | `int32` |  |
| Changes | `[]any` |  |
| Options | `map[string]string` |  |
| Reason | `string` |  |
| TargetID | `any` |  |
| UserID | `any` |  |

### AuditLogObjectChange
AuditLogObjectChange

#### Example Usage

```go
// Create a new AuditLogObjectChange
auditlogobjectchange := AuditLogObjectChange{
    Key: &"example"{},
    NewValue: any{},
    OldValue: any{},
}
```

#### Type Definition

```go
type AuditLogObjectChange struct {
    Key *string `json:"key,omitempty"`
    NewValue any `json:"new_value,omitempty"`
    OldValue any `json:"old_value,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Key | `*string` |  |
| NewValue | `any` |  |
| OldValue | `any` |  |

### AutomodActionType
AutomodActionType

#### Example Usage

```go
// Example usage of AutomodActionType
var value AutomodActionType
// Initialize with appropriate value
```

#### Type Definition

```go
type AutomodActionType int
```

### AutomodEventType
AutomodEventType

#### Example Usage

```go
// Example usage of AutomodEventType
var value AutomodEventType
// Initialize with appropriate value
```

#### Type Definition

```go
type AutomodEventType int
```

### AutomodKeywordPresetType
AutomodKeywordPresetType

#### Example Usage

```go
// Example usage of AutomodKeywordPresetType
var value AutomodKeywordPresetType
// Initialize with appropriate value
```

#### Type Definition

```go
type AutomodKeywordPresetType int
```

### AutomodTriggerType
AutomodTriggerType

#### Example Usage

```go
// Example usage of AutomodTriggerType
var value AutomodTriggerType
// Initialize with appropriate value
```

#### Type Definition

```go
type AutomodTriggerType int
```

### AvailableLocalesEnum
AvailableLocalesEnum

#### Example Usage

```go
// Example usage of AvailableLocalesEnum
var value AvailableLocalesEnum
// Initialize with appropriate value
```

#### Type Definition

```go
type AvailableLocalesEnum string
```

### BanUserFromGuildOptions
BanUserFromGuildOptions

#### Example Usage

```go
// Create a new BanUserFromGuildOptions
banuserfromguildoptions := BanUserFromGuildOptions{
    DeleteMessageDays: &42{},
    DeleteMessageSeconds: &42{},
}
```

#### Type Definition

```go
type BanUserFromGuildOptions struct {
    DeleteMessageDays *int64 `json:"delete_message_days,omitempty"`
    DeleteMessageSeconds *int64 `json:"delete_message_seconds,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| DeleteMessageDays | `*int64` |  |
| DeleteMessageSeconds | `*int64` |  |

### BaseCreateMessageCreateOptions
BaseCreateMessageCreateOptions

#### Example Usage

```go
// Create a new BaseCreateMessageCreateOptions
basecreatemessagecreateoptions := BaseCreateMessageCreateOptions{
    AllowedMentions: any{},
    Attachments: [],
    Components: [],
    ConfettiPotion: any{},
    Content: &"example"{},
    Embeds: [],
    Flags: &42{},
    Poll: any{},
    SharedClientTheme: any{},
    StickerIds: [],
}
```

#### Type Definition

```go
type BaseCreateMessageCreateOptions struct {
    AllowedMentions any `json:"allowed_mentions,omitempty"`
    Attachments []any `json:"attachments,omitempty"`
    Components []any `json:"components,omitempty"`
    ConfettiPotion any `json:"confetti_potion,omitempty"`
    Content *string `json:"content,omitempty"`
    Embeds []any `json:"embeds,omitempty"`
    Flags *int64 `json:"flags,omitempty"`
    Poll any `json:"poll,omitempty"`
    SharedClientTheme any `json:"shared_client_theme,omitempty"`
    StickerIds []string `json:"sticker_ids,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AllowedMentions | `any` |  |
| Attachments | `[]any` |  |
| Components | `[]any` |  |
| ConfettiPotion | `any` |  |
| Content | `*string` |  |
| Embeds | `[]any` |  |
| Flags | `*int64` |  |
| Poll | `any` |  |
| SharedClientTheme | `any` |  |
| StickerIds | `[]string` |  |

### BasicApplication
BasicApplication

#### Example Usage

```go
// Create a new BasicApplication
basicapplication := BasicApplication{
    ID: "example",
    Bot: any{},
    CoverImage: "example",
    Description: "example",
    Icon: &"example"{},
    Name: "example",
    PrimarySkuID: "example",
    Type: any{},
}
```

#### Type Definition

```go
type BasicApplication struct {
    ID string `json:"id"`
    Bot any `json:"bot,omitempty"`
    CoverImage string `json:"cover_image,omitempty"`
    Description string `json:"description"`
    Icon *string `json:"icon,omitempty"`
    Name string `json:"name"`
    PrimarySkuID string `json:"primary_sku_id,omitempty"`
    Type any `json:"type,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Bot | `any` |  |
| CoverImage | `string` |  |
| Description | `string` |  |
| Icon | `*string` |  |
| Name | `string` |  |
| PrimarySkuID | `string` |  |
| Type | `any` |  |

### BasicGuildMember
BasicGuildMember

#### Example Usage

```go
// Create a new BasicGuildMember
basicguildmember := BasicGuildMember{
    Avatar: &"example"{},
    AvatarDecorationData: any{},
    Banner: &"example"{},
    Collectibles: any{},
    CommunicationDisabledUntil: &/* value */{},
    Flags: 42,
    JoinedAt: /* value */,
    Nick: &"example"{},
    Pending: true,
    PremiumSince: &/* value */{},
    Roles: [],
}
```

#### Type Definition

```go
type BasicGuildMember struct {
    Avatar *string `json:"avatar,omitempty"`
    AvatarDecorationData any `json:"avatar_decoration_data,omitempty"`
    Banner *string `json:"banner,omitempty"`
    Collectibles any `json:"collectibles,omitempty"`
    CommunicationDisabledUntil *time.Time `json:"communication_disabled_until,omitempty"`
    Flags int32 `json:"flags"`
    JoinedAt time.Time `json:"joined_at"`
    Nick *string `json:"nick,omitempty"`
    Pending bool `json:"pending"`
    PremiumSince *time.Time `json:"premium_since,omitempty"`
    Roles []string `json:"roles"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Avatar | `*string` |  |
| AvatarDecorationData | `any` |  |
| Banner | `*string` |  |
| Collectibles | `any` |  |
| CommunicationDisabledUntil | `*time.Time` |  |
| Flags | `int32` |  |
| JoinedAt | `time.Time` |  |
| Nick | `*string` |  |
| Pending | `bool` |  |
| PremiumSince | `*time.Time` |  |
| Roles | `[]string` |  |

### BasicMessage
BasicMessage

#### Example Usage

```go
// Create a new BasicMessage
basicmessage := BasicMessage{
    ID: "example",
    Activity: any{},
    Application: any{},
    ApplicationID: "example",
    Attachments: [],
    Author: any{},
    Call: any{},
    ChannelID: "example",
    Components: [],
    Content: "example",
    EditedTimestamp: &/* value */{},
    Embeds: [],
    Flags: 42,
    Interaction: any{},
    InteractionMetadata: any{},
    MentionChannels: [],
    MentionEveryone: true,
    MentionRoles: [],
    Mentions: [],
    MessageReference: any{},
    MessageSnapshots: [],
    Nonce: any{},
    Pinned: true,
    Poll: any{},
    Position: 42,
    PurchaseNotification: any{},
    Resolved: any{},
    RoleSubscriptionData: any{},
    SharedClientTheme: any{},
    StickerItems: [],
    Stickers: [],
    Thread: any{},
    Timestamp: /* value */,
    Tts: true,
    Type: 42,
    WebhookID: "example",
}
```

#### Type Definition

```go
type BasicMessage struct {
    ID string `json:"id"`
    Activity any `json:"activity,omitempty"`
    Application any `json:"application,omitempty"`
    ApplicationID string `json:"application_id,omitempty"`
    Attachments []any `json:"attachments"`
    Author any `json:"author"`
    Call any `json:"call,omitempty"`
    ChannelID string `json:"channel_id"`
    Components []any `json:"components"`
    Content string `json:"content"`
    EditedTimestamp *time.Time `json:"edited_timestamp,omitempty"`
    Embeds []any `json:"embeds"`
    Flags int32 `json:"flags"`
    Interaction any `json:"interaction,omitempty"`
    InteractionMetadata any `json:"interaction_metadata,omitempty"`
    MentionChannels []any `json:"mention_channels,omitempty"`
    MentionEveryone bool `json:"mention_everyone"`
    MentionRoles []string `json:"mention_roles"`
    Mentions []any `json:"mentions"`
    MessageReference any `json:"message_reference,omitempty"`
    MessageSnapshots []any `json:"message_snapshots,omitempty"`
    Nonce any `json:"nonce,omitempty"`
    Pinned bool `json:"pinned"`
    Poll any `json:"poll,omitempty"`
    Position int32 `json:"position,omitempty"`
    PurchaseNotification any `json:"purchase_notification,omitempty"`
    Resolved any `json:"resolved,omitempty"`
    RoleSubscriptionData any `json:"role_subscription_data,omitempty"`
    SharedClientTheme any `json:"shared_client_theme,omitempty"`
    StickerItems []any `json:"sticker_items,omitempty"`
    Stickers []any `json:"stickers,omitempty"`
    Thread any `json:"thread,omitempty"`
    Timestamp time.Time `json:"timestamp"`
    Tts bool `json:"tts"`
    Type int32 `json:"type"`
    WebhookID string `json:"webhook_id,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Activity | `any` |  |
| Application | `any` |  |
| ApplicationID | `string` |  |
| Attachments | `[]any` |  |
| Author | `any` |  |
| Call | `any` |  |
| ChannelID | `string` |  |
| Components | `[]any` |  |
| Content | `string` |  |
| EditedTimestamp | `*time.Time` |  |
| Embeds | `[]any` |  |
| Flags | `int32` |  |
| Interaction | `any` |  |
| InteractionMetadata | `any` |  |
| MentionChannels | `[]any` |  |
| MentionEveryone | `bool` |  |
| MentionRoles | `[]string` |  |
| Mentions | `[]any` |  |
| MessageReference | `any` |  |
| MessageSnapshots | `[]any` |  |
| Nonce | `any` |  |
| Pinned | `bool` |  |
| Poll | `any` |  |
| Position | `int32` |  |
| PurchaseNotification | `any` |  |
| Resolved | `any` |  |
| RoleSubscriptionData | `any` |  |
| SharedClientTheme | `any` |  |
| StickerItems | `[]any` |  |
| Stickers | `[]any` |  |
| Thread | `any` |  |
| Timestamp | `time.Time` |  |
| Tts | `bool` |  |
| Type | `int32` |  |
| WebhookID | `string` |  |

### BlockMessageAction
BlockMessageAction

#### Example Usage

```go
// Create a new BlockMessageAction
blockmessageaction := BlockMessageAction{
    Metadata: any{},
    Type: 42,
}
```

#### Type Definition

```go
type BlockMessageAction struct {
    Metadata any `json:"metadata"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Metadata | `any` |  |
| Type | `int32` |  |

### BlockMessageActionMetadata
BlockMessageActionMetadata

#### Example Usage

```go
// Create a new BlockMessageActionMetadata
blockmessageactionmetadata := BlockMessageActionMetadata{
    CustomMessage: "example",
}
```

#### Type Definition

```go
type BlockMessageActionMetadata struct {
    CustomMessage string `json:"custom_message,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| CustomMessage | `string` |  |

### BotAccountPatchOptions
BotAccountPatchOptions

#### Example Usage

```go
// Create a new BotAccountPatchOptions
botaccountpatchoptions := BotAccountPatchOptions{
    Avatar: &"example"{},
    Banner: &"example"{},
    Username: "example",
}
```

#### Type Definition

```go
type BotAccountPatchOptions struct {
    Avatar *string `json:"avatar,omitempty"`
    Banner *string `json:"banner,omitempty"`
    Username string `json:"username"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Avatar | `*string` |  |
| Banner | `*string` |  |
| Username | `string` |  |

### BotAddGuildMemberOptions
BotAddGuildMemberOptions

#### Example Usage

```go
// Create a new BotAddGuildMemberOptions
botaddguildmemberoptions := BotAddGuildMemberOptions{
    AccessToken: "example",
    Deaf: &true{},
    Flags: &42{},
    Mute: &true{},
    Nick: &"example"{},
    Roles: [],
}
```

#### Type Definition

```go
type BotAddGuildMemberOptions struct {
    AccessToken string `json:"access_token"`
    Deaf *bool `json:"deaf,omitempty"`
    Flags *int64 `json:"flags,omitempty"`
    Mute *bool `json:"mute,omitempty"`
    Nick *string `json:"nick,omitempty"`
    Roles []string `json:"roles,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AccessToken | `string` |  |
| Deaf | `*bool` |  |
| Flags | `*int64` |  |
| Mute | `*bool` |  |
| Nick | `*string` |  |
| Roles | `[]string` |  |

### BulkBanUsers
BulkBanUsers

#### Example Usage

```go
// Create a new BulkBanUsers
bulkbanusers := BulkBanUsers{
    BannedUsers: [],
    FailedUsers: [],
}
```

#### Type Definition

```go
type BulkBanUsers struct {
    BannedUsers []string `json:"banned_users"`
    FailedUsers []string `json:"failed_users"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| BannedUsers | `[]string` |  |
| FailedUsers | `[]string` |  |

### BulkBanUsersOptions
BulkBanUsersOptions

#### Example Usage

```go
// Create a new BulkBanUsersOptions
bulkbanusersoptions := BulkBanUsersOptions{
    DeleteMessageSeconds: &42{},
    UserIds: [],
}
```

#### Type Definition

```go
type BulkBanUsersOptions struct {
    DeleteMessageSeconds *int64 `json:"delete_message_seconds,omitempty"`
    UserIds []string `json:"user_ids"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| DeleteMessageSeconds | `*int64` |  |
| UserIds | `[]string` |  |

### BulkLobbyMemberOptions
BulkLobbyMemberOptions

#### Example Usage

```go
// Create a new BulkLobbyMemberOptions
bulklobbymemberoptions := BulkLobbyMemberOptions{
    ID: "example",
    Flags: any{},
    Metadata: map[],
    RemoveMember: &true{},
}
```

#### Type Definition

```go
type BulkLobbyMemberOptions struct {
    ID string `json:"id"`
    Flags any `json:"flags,omitempty"`
    Metadata map[string]string `json:"metadata,omitempty"`
    RemoveMember *bool `json:"remove_member,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Flags | `any` |  |
| Metadata | `map[string]string` |  |
| RemoveMember | `*bool` |  |

### ButtonComponent
ButtonComponent

#### Example Usage

```go
// Create a new ButtonComponent
buttoncomponent := ButtonComponent{
    ID: 42,
    CustomID: "example",
    Disabled: true,
    Emoji: any{},
    Label: "example",
    SkuID: "example",
    Style: 42,
    Type: 42,
    URL: &"example"{},
}
```

#### Type Definition

```go
type ButtonComponent struct {
    ID int32 `json:"id"`
    CustomID string `json:"custom_id,omitempty"`
    Disabled bool `json:"disabled,omitempty"`
    Emoji any `json:"emoji,omitempty"`
    Label string `json:"label,omitempty"`
    SkuID string `json:"sku_id,omitempty"`
    Style int32 `json:"style"`
    Type int32 `json:"type"`
    URL *string `json:"url,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| CustomID | `string` |  |
| Disabled | `bool` |  |
| Emoji | `any` |  |
| Label | `string` |  |
| SkuID | `string` |  |
| Style | `int32` |  |
| Type | `int32` |  |
| URL | `*string` |  |

### ButtonComponentForMessageOptions
ButtonComponentForMessageOptions

#### Example Usage

```go
// Create a new ButtonComponentForMessageOptions
buttoncomponentformessageoptions := ButtonComponentForMessageOptions{
    ID: &42{},
    CustomID: &"example"{},
    Disabled: &true{},
    Emoji: any{},
    Label: &"example"{},
    SkuID: any{},
    Style: 42,
    Type: 42,
    URL: &"example"{},
}
```

#### Type Definition

```go
type ButtonComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    CustomID *string `json:"custom_id,omitempty"`
    Disabled *bool `json:"disabled,omitempty"`
    Emoji any `json:"emoji,omitempty"`
    Label *string `json:"label,omitempty"`
    SkuID any `json:"sku_id,omitempty"`
    Style int32 `json:"style"`
    Type int32 `json:"type"`
    URL *string `json:"url,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| CustomID | `*string` |  |
| Disabled | `*bool` |  |
| Emoji | `any` |  |
| Label | `*string` |  |
| SkuID | `any` |  |
| Style | `int32` |  |
| Type | `int32` |  |
| URL | `*string` |  |

### ButtonStyleTypes
ButtonStyleTypes

#### Example Usage

```go
// Example usage of ButtonStyleTypes
var value ButtonStyleTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type ButtonStyleTypes int
```

### ChannelFollower
ChannelFollower

#### Example Usage

```go
// Create a new ChannelFollower
channelfollower := ChannelFollower{
    ChannelID: "example",
    WebhookID: "example",
}
```

#### Type Definition

```go
type ChannelFollower struct {
    ChannelID string `json:"channel_id"`
    WebhookID string `json:"webhook_id"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `string` |  |
| WebhookID | `string` |  |

### ChannelFollowerWebhook
ChannelFollowerWebhook

#### Example Usage

```go
// Create a new ChannelFollowerWebhook
channelfollowerwebhook := ChannelFollowerWebhook{
    ID: "example",
    ApplicationID: any{},
    Avatar: &"example"{},
    ChannelID: any{},
    GuildID: any{},
    Name: "example",
    SourceChannel: any{},
    SourceGuild: any{},
    Type: 42,
    User: any{},
}
```

#### Type Definition

```go
type ChannelFollowerWebhook struct {
    ID string `json:"id"`
    ApplicationID any `json:"application_id,omitempty"`
    Avatar *string `json:"avatar,omitempty"`
    ChannelID any `json:"channel_id,omitempty"`
    GuildID any `json:"guild_id,omitempty"`
    Name string `json:"name"`
    SourceChannel any `json:"source_channel,omitempty"`
    SourceGuild any `json:"source_guild,omitempty"`
    Type int32 `json:"type"`
    User any `json:"user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ApplicationID | `any` |  |
| Avatar | `*string` |  |
| ChannelID | `any` |  |
| GuildID | `any` |  |
| Name | `string` |  |
| SourceChannel | `any` |  |
| SourceGuild | `any` |  |
| Type | `int32` |  |
| User | `any` |  |

### ChannelPermissionOverwrite
ChannelPermissionOverwrite

#### Example Usage

```go
// Create a new ChannelPermissionOverwrite
channelpermissionoverwrite := ChannelPermissionOverwrite{
    ID: "example",
    Allow: "example",
    Deny: "example",
    Type: 42,
}
```

#### Type Definition

```go
type ChannelPermissionOverwrite struct {
    ID string `json:"id"`
    Allow string `json:"allow"`
    Deny string `json:"deny"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Allow | `string` |  |
| Deny | `string` |  |
| Type | `int32` |  |

### ChannelPermissionOverwriteOptions
ChannelPermissionOverwriteOptions

#### Example Usage

```go
// Create a new ChannelPermissionOverwriteOptions
channelpermissionoverwriteoptions := ChannelPermissionOverwriteOptions{
    ID: "example",
    Allow: &42{},
    Deny: &42{},
    Type: any{},
}
```

#### Type Definition

```go
type ChannelPermissionOverwriteOptions struct {
    ID string `json:"id"`
    Allow *int64 `json:"allow,omitempty"`
    Deny *int64 `json:"deny,omitempty"`
    Type any `json:"type,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Allow | `*int64` |  |
| Deny | `*int64` |  |
| Type | `any` |  |

### ChannelPermissionOverwrites
ChannelPermissionOverwrites

#### Example Usage

```go
// Example usage of ChannelPermissionOverwrites
var value ChannelPermissionOverwrites
// Initialize with appropriate value
```

#### Type Definition

```go
type ChannelPermissionOverwrites int
```

### ChannelSelectComponent
ChannelSelectComponent

#### Example Usage

```go
// Create a new ChannelSelectComponent
channelselectcomponent := ChannelSelectComponent{
    ID: 42,
    ChannelTypes: [],
    CustomID: "example",
    DefaultValues: [],
    Disabled: true,
    MaxValues: &42{},
    MinValues: &42{},
    Placeholder: "example",
    Type: 42,
}
```

#### Type Definition

```go
type ChannelSelectComponent struct {
    ID int32 `json:"id"`
    ChannelTypes []int32 `json:"channel_types,omitempty"`
    CustomID string `json:"custom_id"`
    DefaultValues []any `json:"default_values,omitempty"`
    Disabled bool `json:"disabled,omitempty"`
    MaxValues *int32 `json:"max_values,omitempty"`
    MinValues *int32 `json:"min_values,omitempty"`
    Placeholder string `json:"placeholder,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| ChannelTypes | `[]int32` |  |
| CustomID | `string` |  |
| DefaultValues | `[]any` |  |
| Disabled | `bool` |  |
| MaxValues | `*int32` |  |
| MinValues | `*int32` |  |
| Placeholder | `string` |  |
| Type | `int32` |  |

### ChannelSelectComponentForMessageOptions
ChannelSelectComponentForMessageOptions

#### Example Usage

```go
// Create a new ChannelSelectComponentForMessageOptions
channelselectcomponentformessageoptions := ChannelSelectComponentForMessageOptions{
    ID: &42{},
    ChannelTypes: [],
    CustomID: "example",
    DefaultValues: [],
    Disabled: &true{},
    MaxValues: &42{},
    MinValues: &42{},
    Placeholder: &"example"{},
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type ChannelSelectComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    ChannelTypes []int32 `json:"channel_types,omitempty"`
    CustomID string `json:"custom_id"`
    DefaultValues []any `json:"default_values,omitempty"`
    Disabled *bool `json:"disabled,omitempty"`
    MaxValues *int64 `json:"max_values,omitempty"`
    MinValues *int64 `json:"min_values,omitempty"`
    Placeholder *string `json:"placeholder,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| ChannelTypes | `[]int32` |  |
| CustomID | `string` |  |
| DefaultValues | `[]any` |  |
| Disabled | `*bool` |  |
| MaxValues | `*int64` |  |
| MinValues | `*int64` |  |
| Placeholder | `*string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### ChannelSelectComponentForModalOptions
ChannelSelectComponentForModalOptions

#### Example Usage

```go
// Create a new ChannelSelectComponentForModalOptions
channelselectcomponentformodaloptions := ChannelSelectComponentForModalOptions{
    ID: &42{},
    ChannelTypes: [],
    CustomID: "example",
    DefaultValues: [],
    Disabled: &true{},
    MaxValues: &42{},
    MinValues: &42{},
    Placeholder: &"example"{},
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type ChannelSelectComponentForModalOptions struct {
    ID *int32 `json:"id,omitempty"`
    ChannelTypes []int32 `json:"channel_types,omitempty"`
    CustomID string `json:"custom_id"`
    DefaultValues []any `json:"default_values,omitempty"`
    Disabled *bool `json:"disabled,omitempty"`
    MaxValues *int64 `json:"max_values,omitempty"`
    MinValues *int64 `json:"min_values,omitempty"`
    Placeholder *string `json:"placeholder,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| ChannelTypes | `[]int32` |  |
| CustomID | `string` |  |
| DefaultValues | `[]any` |  |
| Disabled | `*bool` |  |
| MaxValues | `*int64` |  |
| MinValues | `*int64` |  |
| Placeholder | `*string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### ChannelSelectDefaultValue
ChannelSelectDefaultValue

#### Example Usage

```go
// Create a new ChannelSelectDefaultValue
channelselectdefaultvalue := ChannelSelectDefaultValue{
    ID: "example",
    Type: "example",
}
```

#### Type Definition

```go
type ChannelSelectDefaultValue struct {
    ID string `json:"id"`
    Type string `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Type | `string` |  |

### ChannelTypes
ChannelTypes

#### Example Usage

```go
// Example usage of ChannelTypes
var value ChannelTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type ChannelTypes int
```

### CommandPermission
CommandPermission

#### Example Usage

```go
// Create a new CommandPermission
commandpermission := CommandPermission{
    ID: "example",
    Permission: true,
    Type: 42,
}
```

#### Type Definition

```go
type CommandPermission struct {
    ID string `json:"id"`
    Permission bool `json:"permission"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Permission | `bool` |  |
| Type | `int32` |  |

### CommandPermissions
CommandPermissions

#### Example Usage

```go
// Create a new CommandPermissions
commandpermissions := CommandPermissions{
    ID: "example",
    ApplicationID: "example",
    GuildID: "example",
    Permissions: [],
}
```

#### Type Definition

```go
type CommandPermissions struct {
    ID string `json:"id"`
    ApplicationID string `json:"application_id"`
    GuildID string `json:"guild_id"`
    Permissions []any `json:"permissions"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ApplicationID | `string` |  |
| GuildID | `string` |  |
| Permissions | `[]any` |  |

### ComponentEmoji
ComponentEmoji

#### Example Usage

```go
// Create a new ComponentEmoji
componentemoji := ComponentEmoji{
    ID: "example",
    Animated: true,
    Name: "example",
}
```

#### Type Definition

```go
type ComponentEmoji struct {
    ID string `json:"id,omitempty"`
    Animated bool `json:"animated,omitempty"`
    Name string `json:"name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Animated | `bool` |  |
| Name | `string` |  |

### ComponentEmojiForOptions
ComponentEmojiForOptions

#### Example Usage

```go
// Create a new ComponentEmojiForOptions
componentemojiforoptions := ComponentEmojiForOptions{
    ID: any{},
    Name: "example",
}
```

#### Type Definition

```go
type ComponentEmojiForOptions struct {
    ID any `json:"id,omitempty"`
    Name string `json:"name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `any` |  |
| Name | `string` |  |

### ConfettiPotionCreateOptions
ConfettiPotionCreateOptions

#### Example Usage

```go
// Create a new ConfettiPotionCreateOptions
confettipotioncreateoptions := ConfettiPotionCreateOptions{

}
```

#### Type Definition

```go
type ConfettiPotionCreateOptions struct {
}
```

### ConnectedAccount
ConnectedAccount

#### Example Usage

```go
// Create a new ConnectedAccount
connectedaccount := ConnectedAccount{
    ID: "example",
    FriendSync: true,
    Integrations: [],
    Name: &"example"{},
    Revoked: true,
    ShowActivity: true,
    TwoWayLink: true,
    Type: "example",
    Verified: true,
    Visibility: 42,
}
```

#### Type Definition

```go
type ConnectedAccount struct {
    ID string `json:"id"`
    FriendSync bool `json:"friend_sync"`
    Integrations []any `json:"integrations,omitempty"`
    Name *string `json:"name,omitempty"`
    Revoked bool `json:"revoked,omitempty"`
    ShowActivity bool `json:"show_activity"`
    TwoWayLink bool `json:"two_way_link"`
    Type string `json:"type"`
    Verified bool `json:"verified"`
    Visibility int32 `json:"visibility"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| FriendSync | `bool` |  |
| Integrations | `[]any` |  |
| Name | `*string` |  |
| Revoked | `bool` |  |
| ShowActivity | `bool` |  |
| TwoWayLink | `bool` |  |
| Type | `string` |  |
| Verified | `bool` |  |
| Visibility | `int32` |  |

### ConnectedAccountGuild
ConnectedAccountGuild

#### Example Usage

```go
// Create a new ConnectedAccountGuild
connectedaccountguild := ConnectedAccountGuild{
    ID: "example",
    Icon: &"example"{},
    Name: "example",
}
```

#### Type Definition

```go
type ConnectedAccountGuild struct {
    ID string `json:"id"`
    Icon *string `json:"icon,omitempty"`
    Name string `json:"name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Icon | `*string` |  |
| Name | `string` |  |

### ConnectedAccountIntegration
ConnectedAccountIntegration

#### Example Usage

```go
// Create a new ConnectedAccountIntegration
connectedaccountintegration := ConnectedAccountIntegration{
    ID: "example",
    Account: any{},
    Guild: any{},
    Type: "example",
}
```

#### Type Definition

```go
type ConnectedAccountIntegration struct {
    ID string `json:"id"`
    Account any `json:"account"`
    Guild any `json:"guild"`
    Type string `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Account | `any` |  |
| Guild | `any` |  |
| Type | `string` |  |

### ConnectedAccountProviders
ConnectedAccountProviders

#### Example Usage

```go
// Example usage of ConnectedAccountProviders
var value ConnectedAccountProviders
// Initialize with appropriate value
```

#### Type Definition

```go
type ConnectedAccountProviders string
```

### ConnectedAccountVisibility
ConnectedAccountVisibility

#### Example Usage

```go
// Example usage of ConnectedAccountVisibility
var value ConnectedAccountVisibility
// Initialize with appropriate value
```

#### Type Definition

```go
type ConnectedAccountVisibility int
```

### ContainerComponent
ContainerComponent

#### Example Usage

```go
// Create a new ContainerComponent
containercomponent := ContainerComponent{
    ID: 42,
    AccentColor: &42{},
    Components: [],
    Spoiler: true,
    Type: 42,
}
```

#### Type Definition

```go
type ContainerComponent struct {
    ID int32 `json:"id"`
    AccentColor *int32 `json:"accent_color,omitempty"`
    Components []any `json:"components"`
    Spoiler bool `json:"spoiler"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| AccentColor | `*int32` |  |
| Components | `[]any` |  |
| Spoiler | `bool` |  |
| Type | `int32` |  |

### ContainerComponentForMessageOptions
ContainerComponentForMessageOptions

#### Example Usage

```go
// Create a new ContainerComponentForMessageOptions
containercomponentformessageoptions := ContainerComponentForMessageOptions{
    ID: &42{},
    AccentColor: &42{},
    Components: [],
    Spoiler: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type ContainerComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    AccentColor *int64 `json:"accent_color,omitempty"`
    Components []any `json:"components"`
    Spoiler *bool `json:"spoiler,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| AccentColor | `*int64` |  |
| Components | `[]any` |  |
| Spoiler | `*bool` |  |
| Type | `int32` |  |

### CreateEntitlementOptionsData
CreateEntitlementOptionsData

#### Example Usage

```go
// Create a new CreateEntitlementOptionsData
createentitlementoptionsdata := CreateEntitlementOptionsData{
    OwnerID: "example",
    OwnerType: 42,
    SkuID: "example",
}
```

#### Type Definition

```go
type CreateEntitlementOptionsData struct {
    OwnerID string `json:"owner_id"`
    OwnerType int32 `json:"owner_type"`
    SkuID string `json:"sku_id"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| OwnerID | `string` |  |
| OwnerType | `int32` |  |
| SkuID | `string` |  |

### CreateForumThreadOptions
CreateForumThreadOptions

#### Example Usage

```go
// Create a new CreateForumThreadOptions
createforumthreadoptions := CreateForumThreadOptions{
    AppliedTags: [],
    AutoArchiveDuration: any{},
    Message: any{},
    Name: "example",
    RateLimitPerUser: &42{},
}
```

#### Type Definition

```go
type CreateForumThreadOptions struct {
    AppliedTags []string `json:"applied_tags,omitempty"`
    AutoArchiveDuration any `json:"auto_archive_duration,omitempty"`
    Message any `json:"message"`
    Name string `json:"name"`
    RateLimitPerUser *int64 `json:"rate_limit_per_user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AppliedTags | `[]string` |  |
| AutoArchiveDuration | `any` |  |
| Message | `any` |  |
| Name | `string` |  |
| RateLimitPerUser | `*int64` |  |

### CreateGroupDMInviteOptions
CreateGroupDMInviteOptions

#### Example Usage

```go
// Create a new CreateGroupDMInviteOptions
creategroupdminviteoptions := CreateGroupDMInviteOptions{
    MaxAge: &42{},
}
```

#### Type Definition

```go
type CreateGroupDMInviteOptions struct {
    MaxAge *int64 `json:"max_age,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| MaxAge | `*int64` |  |

### CreateGuildChannelOptions
CreateGuildChannelOptions

#### Example Usage

```go
// Create a new CreateGuildChannelOptions
createguildchanneloptions := CreateGuildChannelOptions{
    AvailableTags: [],
    Bitrate: &42{},
    DefaultAutoArchiveDuration: any{},
    DefaultForumLayout: any{},
    DefaultReactionEmoji: any{},
    DefaultSortOrder: any{},
    DefaultTagSetting: any{},
    DefaultThreadRateLimitPerUser: &42{},
    Name: "example",
    Nsfw: &true{},
    ParentID: any{},
    PermissionOverwrites: [],
    Position: &42{},
    RateLimitPerUser: &42{},
    RtcRegion: &"example"{},
    Topic: &"example"{},
    Type: any{},
    UserLimit: &42{},
    VideoQualityMode: any{},
}
```

#### Type Definition

```go
type CreateGuildChannelOptions struct {
    AvailableTags []any `json:"available_tags,omitempty"`
    Bitrate *int32 `json:"bitrate,omitempty"`
    DefaultAutoArchiveDuration any `json:"default_auto_archive_duration,omitempty"`
    DefaultForumLayout any `json:"default_forum_layout,omitempty"`
    DefaultReactionEmoji any `json:"default_reaction_emoji,omitempty"`
    DefaultSortOrder any `json:"default_sort_order,omitempty"`
    DefaultTagSetting any `json:"default_tag_setting,omitempty"`
    DefaultThreadRateLimitPerUser *int64 `json:"default_thread_rate_limit_per_user,omitempty"`
    Name string `json:"name"`
    Nsfw *bool `json:"nsfw,omitempty"`
    ParentID any `json:"parent_id,omitempty"`
    PermissionOverwrites []any `json:"permission_overwrites,omitempty"`
    Position *int32 `json:"position,omitempty"`
    RateLimitPerUser *int64 `json:"rate_limit_per_user,omitempty"`
    RtcRegion *string `json:"rtc_region,omitempty"`
    Topic *string `json:"topic,omitempty"`
    Type any `json:"type,omitempty"`
    UserLimit *int32 `json:"user_limit,omitempty"`
    VideoQualityMode any `json:"video_quality_mode,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AvailableTags | `[]any` |  |
| Bitrate | `*int32` |  |
| DefaultAutoArchiveDuration | `any` |  |
| DefaultForumLayout | `any` |  |
| DefaultReactionEmoji | `any` |  |
| DefaultSortOrder | `any` |  |
| DefaultTagSetting | `any` |  |
| DefaultThreadRateLimitPerUser | `*int64` |  |
| Name | `string` |  |
| Nsfw | `*bool` |  |
| ParentID | `any` |  |
| PermissionOverwrites | `[]any` |  |
| Position | `*int32` |  |
| RateLimitPerUser | `*int64` |  |
| RtcRegion | `*string` |  |
| Topic | `*string` |  |
| Type | `any` |  |
| UserLimit | `*int32` |  |
| VideoQualityMode | `any` |  |

### CreateGuildInviteOptions
CreateGuildInviteOptions

#### Example Usage

```go
// Create a new CreateGuildInviteOptions
createguildinviteoptions := CreateGuildInviteOptions{
    MaxAge: &42{},
    MaxUses: &42{},
    TargetApplicationID: any{},
    TargetType: any{},
    TargetUserID: any{},
    Temporary: &true{},
    Unique: &true{},
}
```

#### Type Definition

```go
type CreateGuildInviteOptions struct {
    MaxAge *int64 `json:"max_age,omitempty"`
    MaxUses *int64 `json:"max_uses,omitempty"`
    TargetApplicationID any `json:"target_application_id,omitempty"`
    TargetType any `json:"target_type,omitempty"`
    TargetUserID any `json:"target_user_id,omitempty"`
    Temporary *bool `json:"temporary,omitempty"`
    Unique *bool `json:"unique,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| MaxAge | `*int64` |  |
| MaxUses | `*int64` |  |
| TargetApplicationID | `any` |  |
| TargetType | `any` |  |
| TargetUserID | `any` |  |
| Temporary | `*bool` |  |
| Unique | `*bool` |  |

### CreateMessageInteractionCallback
CreateMessageInteractionCallback

#### Example Usage

```go
// Create a new CreateMessageInteractionCallback
createmessageinteractioncallback := CreateMessageInteractionCallback{
    Message: any{},
    Type: 42,
}
```

#### Type Definition

```go
type CreateMessageInteractionCallback struct {
    Message any `json:"message"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Message | `any` |  |
| Type | `int32` |  |

### CreateMessageInteractionCallbackOptions
CreateMessageInteractionCallbackOptions

#### Example Usage

```go
// Create a new CreateMessageInteractionCallbackOptions
createmessageinteractioncallbackoptions := CreateMessageInteractionCallbackOptions{
    Data: any{},
    Type: 42,
}
```

#### Type Definition

```go
type CreateMessageInteractionCallbackOptions struct {
    Data any `json:"data,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Data | `any` |  |
| Type | `int32` |  |

### CreateOrUpdateThreadTagOptions
CreateOrUpdateThreadTagOptions

#### Example Usage

```go
// Create a new CreateOrUpdateThreadTagOptions
createorupdatethreadtagoptions := CreateOrUpdateThreadTagOptions{
    EmojiID: any{},
    EmojiName: &"example"{},
    Moderated: &true{},
    Name: "example",
}
```

#### Type Definition

```go
type CreateOrUpdateThreadTagOptions struct {
    EmojiID any `json:"emoji_id,omitempty"`
    EmojiName *string `json:"emoji_name,omitempty"`
    Moderated *bool `json:"moderated,omitempty"`
    Name string `json:"name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| EmojiID | `any` |  |
| EmojiName | `*string` |  |
| Moderated | `*bool` |  |
| Name | `string` |  |

### CreatePrivateChannelOptions
CreatePrivateChannelOptions

#### Example Usage

```go
// Create a new CreatePrivateChannelOptions
createprivatechanneloptions := CreatePrivateChannelOptions{
    AccessTokens: [],
    Nicks: map[],
    RecipientID: any{},
}
```

#### Type Definition

```go
type CreatePrivateChannelOptions struct {
    AccessTokens []string `json:"access_tokens,omitempty"`
    Nicks map[string]*string `json:"nicks,omitempty"`
    RecipientID any `json:"recipient_id,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AccessTokens | `[]string` |  |
| Nicks | `map[string]*string` |  |
| RecipientID | `any` |  |

### CreateRoleOptions
CreateRoleOptions

#### Example Usage

```go
// Create a new CreateRoleOptions
createroleoptions := CreateRoleOptions{
    Color: &42{},
    Hoist: &true{},
    Icon: &"example"{},
    Mentionable: &true{},
    Name: &"example"{},
    Permissions: &42{},
    UnicodeEmoji: &"example"{},
}
```

#### Type Definition

```go
type CreateRoleOptions struct {
    Color *int64 `json:"color,omitempty"`
    Hoist *bool `json:"hoist,omitempty"`
    Icon *string `json:"icon,omitempty"`
    Mentionable *bool `json:"mentionable,omitempty"`
    Name *string `json:"name,omitempty"`
    Permissions *int64 `json:"permissions,omitempty"`
    UnicodeEmoji *string `json:"unicode_emoji,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Color | `*int64` |  |
| Hoist | `*bool` |  |
| Icon | `*string` |  |
| Mentionable | `*bool` |  |
| Name | `*string` |  |
| Permissions | `*int64` |  |
| UnicodeEmoji | `*string` |  |

### CreateTextThreadWithMessageOptions
CreateTextThreadWithMessageOptions

#### Example Usage

```go
// Create a new CreateTextThreadWithMessageOptions
createtextthreadwithmessageoptions := CreateTextThreadWithMessageOptions{
    AutoArchiveDuration: any{},
    Name: "example",
    RateLimitPerUser: &42{},
}
```

#### Type Definition

```go
type CreateTextThreadWithMessageOptions struct {
    AutoArchiveDuration any `json:"auto_archive_duration,omitempty"`
    Name string `json:"name"`
    RateLimitPerUser *int64 `json:"rate_limit_per_user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AutoArchiveDuration | `any` |  |
| Name | `string` |  |
| RateLimitPerUser | `*int64` |  |

### CreateTextThreadWithoutMessageOptions
CreateTextThreadWithoutMessageOptions

#### Example Usage

```go
// Create a new CreateTextThreadWithoutMessageOptions
createtextthreadwithoutmessageoptions := CreateTextThreadWithoutMessageOptions{
    AutoArchiveDuration: any{},
    Invitable: &true{},
    Name: "example",
    RateLimitPerUser: &42{},
    Type: any{},
}
```

#### Type Definition

```go
type CreateTextThreadWithoutMessageOptions struct {
    AutoArchiveDuration any `json:"auto_archive_duration,omitempty"`
    Invitable *bool `json:"invitable,omitempty"`
    Name string `json:"name"`
    RateLimitPerUser *int64 `json:"rate_limit_per_user,omitempty"`
    Type any `json:"type,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AutoArchiveDuration | `any` |  |
| Invitable | `*bool` |  |
| Name | `string` |  |
| RateLimitPerUser | `*int64` |  |
| Type | `any` |  |

### CreatedThread
CreatedThread

#### Example Usage

```go
// Create a new CreatedThread
createdthread := CreatedThread{
    ID: "example",
    AppliedTags: [],
    Bitrate: 42,
    Flags: 42,
    GuildID: "example",
    LastMessageID: any{},
    LastPinTimestamp: &/* value */{},
    Member: any{},
    MemberCount: 42,
    MessageCount: 42,
    Name: "example",
    OwnerID: "example",
    ParentID: any{},
    Permissions: &"example"{},
    RateLimitPerUser: 42,
    RtcRegion: &"example"{},
    ThreadMetadata: any{},
    TotalMessageSent: 42,
    Type: 42,
    UserLimit: 42,
    VideoQualityMode: 42,
}
```

#### Type Definition

```go
type CreatedThread struct {
    ID string `json:"id"`
    AppliedTags []string `json:"applied_tags,omitempty"`
    Bitrate int32 `json:"bitrate,omitempty"`
    Flags int32 `json:"flags"`
    GuildID string `json:"guild_id"`
    LastMessageID any `json:"last_message_id,omitempty"`
    LastPinTimestamp *time.Time `json:"last_pin_timestamp,omitempty"`
    Member any `json:"member,omitempty"`
    MemberCount int32 `json:"member_count"`
    MessageCount int32 `json:"message_count"`
    Name string `json:"name"`
    OwnerID string `json:"owner_id"`
    ParentID any `json:"parent_id,omitempty"`
    Permissions *string `json:"permissions,omitempty"`
    RateLimitPerUser int32 `json:"rate_limit_per_user,omitempty"`
    RtcRegion *string `json:"rtc_region,omitempty"`
    ThreadMetadata any `json:"thread_metadata"`
    TotalMessageSent int32 `json:"total_message_sent"`
    Type int32 `json:"type"`
    UserLimit int32 `json:"user_limit,omitempty"`
    VideoQualityMode int32 `json:"video_quality_mode,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| AppliedTags | `[]string` |  |
| Bitrate | `int32` |  |
| Flags | `int32` |  |
| GuildID | `string` |  |
| LastMessageID | `any` |  |
| LastPinTimestamp | `*time.Time` |  |
| Member | `any` |  |
| MemberCount | `int32` |  |
| MessageCount | `int32` |  |
| Name | `string` |  |
| OwnerID | `string` |  |
| ParentID | `any` |  |
| Permissions | `*string` |  |
| RateLimitPerUser | `int32` |  |
| RtcRegion | `*string` |  |
| ThreadMetadata | `any` |  |
| TotalMessageSent | `int32` |  |
| Type | `int32` |  |
| UserLimit | `int32` |  |
| VideoQualityMode | `int32` |  |

### CustomClientTheme
CustomClientTheme

#### Example Usage

```go
// Create a new CustomClientTheme
customclienttheme := CustomClientTheme{
    BaseMix: 42,
    BaseTheme: 42,
    Colors: [],
    GradientAngle: 42,
}
```

#### Type Definition

```go
type CustomClientTheme struct {
    BaseMix int32 `json:"base_mix"`
    BaseTheme int32 `json:"base_theme"`
    Colors []string `json:"colors"`
    GradientAngle int32 `json:"gradient_angle"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| BaseMix | `int32` |  |
| BaseTheme | `int32` |  |
| Colors | `[]string` |  |
| GradientAngle | `int32` |  |

### CustomClientThemeShareOptions
CustomClientThemeShareOptions

#### Example Usage

```go
// Create a new CustomClientThemeShareOptions
customclientthemeshareoptions := CustomClientThemeShareOptions{
    BaseMix: 42,
    BaseTheme: any{},
    Colors: [],
    GradientAngle: 42,
}
```

#### Type Definition

```go
type CustomClientThemeShareOptions struct {
    BaseMix int32 `json:"base_mix"`
    BaseTheme any `json:"base_theme,omitempty"`
    Colors []string `json:"colors"`
    GradientAngle int32 `json:"gradient_angle"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| BaseMix | `int32` |  |
| BaseTheme | `any` |  |
| Colors | `[]string` |  |
| GradientAngle | `int32` |  |

### DefaultKeywordListTriggerMetadata
DefaultKeywordListTriggerMetadata

#### Example Usage

```go
// Create a new DefaultKeywordListTriggerMetadata
defaultkeywordlisttriggermetadata := DefaultKeywordListTriggerMetadata{
    AllowList: [],
    Presets: [],
}
```

#### Type Definition

```go
type DefaultKeywordListTriggerMetadata struct {
    AllowList []string `json:"allow_list,omitempty"`
    Presets []int32 `json:"presets,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AllowList | `[]string` |  |
| Presets | `[]int32` |  |

### DefaultKeywordListUpsertOptions
DefaultKeywordListUpsertOptions

#### Example Usage

```go
// Create a new DefaultKeywordListUpsertOptions
defaultkeywordlistupsertoptions := DefaultKeywordListUpsertOptions{
    Actions: [],
    Enabled: &true{},
    EventType: 42,
    ExemptChannels: [],
    ExemptRoles: [],
    Name: "example",
    TriggerMetadata: any{},
    TriggerType: 42,
}
```

#### Type Definition

```go
type DefaultKeywordListUpsertOptions struct {
    Actions []any `json:"actions,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    EventType int32 `json:"event_type"`
    ExemptChannels []string `json:"exempt_channels,omitempty"`
    ExemptRoles []string `json:"exempt_roles,omitempty"`
    Name string `json:"name"`
    TriggerMetadata any `json:"trigger_metadata"`
    TriggerType int32 `json:"trigger_type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Actions | `[]any` |  |
| Enabled | `*bool` |  |
| EventType | `int32` |  |
| ExemptChannels | `[]string` |  |
| ExemptRoles | `[]string` |  |
| Name | `string` |  |
| TriggerMetadata | `any` |  |
| TriggerType | `int32` |  |

### DefaultKeywordListUpsertOptionsPartial
DefaultKeywordListUpsertOptionsPartial

#### Example Usage

```go
// Create a new DefaultKeywordListUpsertOptionsPartial
defaultkeywordlistupsertoptionspartial := DefaultKeywordListUpsertOptionsPartial{
    Actions: [],
    Enabled: &true{},
    EventType: 42,
    ExemptChannels: [],
    ExemptRoles: [],
    Name: "example",
    TriggerMetadata: any{},
    TriggerType: 42,
}
```

#### Type Definition

```go
type DefaultKeywordListUpsertOptionsPartial struct {
    Actions []any `json:"actions,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    EventType int32 `json:"event_type,omitempty"`
    ExemptChannels []string `json:"exempt_channels,omitempty"`
    ExemptRoles []string `json:"exempt_roles,omitempty"`
    Name string `json:"name,omitempty"`
    TriggerMetadata any `json:"trigger_metadata,omitempty"`
    TriggerType int32 `json:"trigger_type,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Actions | `[]any` |  |
| Enabled | `*bool` |  |
| EventType | `int32` |  |
| ExemptChannels | `[]string` |  |
| ExemptRoles | `[]string` |  |
| Name | `string` |  |
| TriggerMetadata | `any` |  |
| TriggerType | `int32` |  |

### DefaultKeywordRule
DefaultKeywordRule

#### Example Usage

```go
// Create a new DefaultKeywordRule
defaultkeywordrule := DefaultKeywordRule{
    ID: "example",
    Actions: [],
    CreatorID: "example",
    Enabled: true,
    EventType: 42,
    ExemptChannels: [],
    ExemptRoles: [],
    GuildID: "example",
    Name: "example",
    TriggerMetadata: any{},
    TriggerType: 42,
}
```

#### Type Definition

```go
type DefaultKeywordRule struct {
    ID string `json:"id"`
    Actions []any `json:"actions"`
    CreatorID string `json:"creator_id"`
    Enabled bool `json:"enabled"`
    EventType int32 `json:"event_type"`
    ExemptChannels []string `json:"exempt_channels"`
    ExemptRoles []string `json:"exempt_roles"`
    GuildID string `json:"guild_id"`
    Name string `json:"name"`
    TriggerMetadata any `json:"trigger_metadata"`
    TriggerType int32 `json:"trigger_type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Actions | `[]any` |  |
| CreatorID | `string` |  |
| Enabled | `bool` |  |
| EventType | `int32` |  |
| ExemptChannels | `[]string` |  |
| ExemptRoles | `[]string` |  |
| GuildID | `string` |  |
| Name | `string` |  |
| TriggerMetadata | `any` |  |
| TriggerType | `int32` |  |

### DefaultReactionEmoji
DefaultReactionEmoji

#### Example Usage

```go
// Create a new DefaultReactionEmoji
defaultreactionemoji := DefaultReactionEmoji{
    EmojiID: any{},
    EmojiName: &"example"{},
}
```

#### Type Definition

```go
type DefaultReactionEmoji struct {
    EmojiID any `json:"emoji_id,omitempty"`
    EmojiName *string `json:"emoji_name,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| EmojiID | `any` |  |
| EmojiName | `*string` |  |

### DiscordIntegration
DiscordIntegration

#### Example Usage

```go
// Create a new DiscordIntegration
discordintegration := DiscordIntegration{
    ID: "example",
    Account: any{},
    Application: any{},
    Enabled: true,
    Name: &"example"{},
    Scopes: [],
    Type: "example",
    User: any{},
}
```

#### Type Definition

```go
type DiscordIntegration struct {
    ID string `json:"id"`
    Account any `json:"account"`
    Application any `json:"application"`
    Enabled bool `json:"enabled"`
    Name *string `json:"name,omitempty"`
    Scopes []string `json:"scopes"`
    Type string `json:"type"`
    User any `json:"user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Account | `any` |  |
| Application | `any` |  |
| Enabled | `bool` |  |
| Name | `*string` |  |
| Scopes | `[]string` |  |
| Type | `string` |  |
| User | `any` |  |

### EmbeddedActivityInstance
EmbeddedActivityInstance

#### Example Usage

```go
// Create a new EmbeddedActivityInstance
embeddedactivityinstance := EmbeddedActivityInstance{
    ApplicationID: "example",
    InstanceID: "example",
    LaunchID: "example",
    Location: any{},
    Users: [],
}
```

#### Type Definition

```go
type EmbeddedActivityInstance struct {
    ApplicationID string `json:"application_id"`
    InstanceID string `json:"instance_id"`
    LaunchID string `json:"launch_id"`
    Location any `json:"location"`
    Users []string `json:"users"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ApplicationID | `string` |  |
| InstanceID | `string` |  |
| LaunchID | `string` |  |
| Location | `any` |  |
| Users | `[]string` |  |

### EmbeddedActivityLocationKind
EmbeddedActivityLocationKind

#### Example Usage

```go
// Example usage of EmbeddedActivityLocationKind
var value EmbeddedActivityLocationKind
// Initialize with appropriate value
```

#### Type Definition

```go
type EmbeddedActivityLocationKind string
```

### Emoji
Emoji

#### Example Usage

```go
// Create a new Emoji
emoji := Emoji{
    ID: "example",
    Animated: true,
    Available: true,
    Managed: true,
    Name: "example",
    RequireColons: true,
    Roles: [],
    User: any{},
}
```

#### Type Definition

```go
type Emoji struct {
    ID string `json:"id"`
    Animated bool `json:"animated"`
    Available bool `json:"available"`
    Managed bool `json:"managed"`
    Name string `json:"name"`
    RequireColons bool `json:"require_colons"`
    Roles []string `json:"roles"`
    User any `json:"user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Animated | `bool` |  |
| Available | `bool` |  |
| Managed | `bool` |  |
| Name | `string` |  |
| RequireColons | `bool` |  |
| Roles | `[]string` |  |
| User | `any` |  |

### Entitlement
Entitlement

#### Example Usage

```go
// Create a new Entitlement
entitlement := Entitlement{
    ID: "example",
    ApplicationID: "example",
    Consumed: &true{},
    Deleted: true,
    EndsAt: &/* value */{},
    FulfilledAt: &/* value */{},
    FulfillmentStatus: any{},
    GifterUserID: any{},
    GuildID: any{},
    ParentID: any{},
    SkuID: "example",
    StartsAt: &/* value */{},
    Type: 42,
    UserID: "example",
}
```

#### Type Definition

```go
type Entitlement struct {
    ID string `json:"id"`
    ApplicationID string `json:"application_id"`
    Consumed *bool `json:"consumed,omitempty"`
    Deleted bool `json:"deleted"`
    EndsAt *time.Time `json:"ends_at,omitempty"`
    FulfilledAt *time.Time `json:"fulfilled_at,omitempty"`
    FulfillmentStatus any `json:"fulfillment_status,omitempty"`
    GifterUserID any `json:"gifter_user_id,omitempty"`
    GuildID any `json:"guild_id,omitempty"`
    ParentID any `json:"parent_id,omitempty"`
    SkuID string `json:"sku_id"`
    StartsAt *time.Time `json:"starts_at,omitempty"`
    Type int32 `json:"type"`
    UserID string `json:"user_id"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ApplicationID | `string` |  |
| Consumed | `*bool` |  |
| Deleted | `bool` |  |
| EndsAt | `*time.Time` |  |
| FulfilledAt | `*time.Time` |  |
| FulfillmentStatus | `any` |  |
| GifterUserID | `any` |  |
| GuildID | `any` |  |
| ParentID | `any` |  |
| SkuID | `string` |  |
| StartsAt | `*time.Time` |  |
| Type | `int32` |  |
| UserID | `string` |  |

### EntitlementOwnerTypes
EntitlementOwnerTypes

#### Example Usage

```go
// Create a new EntitlementOwnerTypes
entitlementownertypes := EntitlementOwnerTypes{

}
```

#### Type Definition

```go
type EntitlementOwnerTypes struct {
}
```

### EntitlementTenantFulfillmentStatus
EntitlementTenantFulfillmentStatus

#### Example Usage

```go
// Example usage of EntitlementTenantFulfillmentStatus
var value EntitlementTenantFulfillmentStatus
// Initialize with appropriate value
```

#### Type Definition

```go
type EntitlementTenantFulfillmentStatus int
```

### EntitlementTypes
EntitlementTypes

#### Example Usage

```go
// Example usage of EntitlementTypes
var value EntitlementTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type EntitlementTypes int
```

### EntityMetadataExternal
EntityMetadataExternal

#### Example Usage

```go
// Create a new EntityMetadataExternal
entitymetadataexternal := EntityMetadataExternal{
    Location: "example",
}
```

#### Type Definition

```go
type EntityMetadataExternal struct {
    Location string `json:"location"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Location | `string` |  |

### EntityMetadataStageInstance
EntityMetadataStageInstance

#### Example Usage

```go
// Create a new EntityMetadataStageInstance
entitymetadatastageinstance := EntityMetadataStageInstance{

}
```

#### Type Definition

```go
type EntityMetadataStageInstance struct {
}
```

### EntityMetadataVoice
EntityMetadataVoice

#### Example Usage

```go
// Create a new EntityMetadataVoice
entitymetadatavoice := EntityMetadataVoice{

}
```

#### Type Definition

```go
type EntityMetadataVoice struct {
}
```

### ErrorDetails
ErrorDetails

#### Example Usage

```go
// Create a new ErrorDetails
errordetails := ErrorDetails{

}
```

#### Type Definition

```go
type ErrorDetails struct {
}
```

### ErrorResponse
Error A single error, either for an API response or a specific field.

#### Example Usage

```go
// Create a new ErrorResponse
errorresponse := ErrorResponse{
    Code: 42,
    Message: "example",
}
```

#### Type Definition

```go
type ErrorResponse struct {
    Code int64 `json:"code"`
    Message string `json:"message"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Code | `int64` | Discord internal error code. See error code reference |
| Message | `string` | Human-readable error message |

### ExternalConnectionIntegration
ExternalConnectionIntegration

#### Example Usage

```go
// Create a new ExternalConnectionIntegration
externalconnectionintegration := ExternalConnectionIntegration{
    ID: "example",
    Account: any{},
    EnableEmoticons: true,
    Enabled: true,
    ExpireBehavior: 42,
    ExpireGracePeriod: 42,
    Name: &"example"{},
    Revoked: true,
    RoleID: any{},
    SubscriberCount: 42,
    SyncedAt: /* value */,
    Syncing: true,
    Type: "example",
    User: any{},
}
```

#### Type Definition

```go
type ExternalConnectionIntegration struct {
    ID string `json:"id"`
    Account any `json:"account"`
    EnableEmoticons bool `json:"enable_emoticons,omitempty"`
    Enabled bool `json:"enabled"`
    ExpireBehavior int32 `json:"expire_behavior,omitempty"`
    ExpireGracePeriod int32 `json:"expire_grace_period,omitempty"`
    Name *string `json:"name,omitempty"`
    Revoked bool `json:"revoked,omitempty"`
    RoleID any `json:"role_id,omitempty"`
    SubscriberCount int32 `json:"subscriber_count,omitempty"`
    SyncedAt time.Time `json:"synced_at,omitempty"`
    Syncing bool `json:"syncing,omitempty"`
    Type string `json:"type"`
    User any `json:"user"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Account | `any` |  |
| EnableEmoticons | `bool` |  |
| Enabled | `bool` |  |
| ExpireBehavior | `int32` |  |
| ExpireGracePeriod | `int32` |  |
| Name | `*string` |  |
| Revoked | `bool` |  |
| RoleID | `any` |  |
| SubscriberCount | `int32` |  |
| SyncedAt | `time.Time` |  |
| Syncing | `bool` |  |
| Type | `string` |  |
| User | `any` |  |

### ExternalScheduledEvent
ExternalScheduledEvent

#### Example Usage

```go
// Create a new ExternalScheduledEvent
externalscheduledevent := ExternalScheduledEvent{
    ID: "example",
    ChannelID: any{},
    Creator: any{},
    CreatorID: any{},
    Description: &"example"{},
    EntityID: any{},
    EntityMetadata: any{},
    EntityType: 42,
    GuildID: "example",
    Image: &"example"{},
    Name: "example",
    PrivacyLevel: 42,
    ScheduledEndTime: &/* value */{},
    ScheduledStartTime: /* value */,
    Status: 42,
    UserCount: 42,
    UserRsvp: any{},
}
```

#### Type Definition

```go
type ExternalScheduledEvent struct {
    ID string `json:"id"`
    ChannelID any `json:"channel_id,omitempty"`
    Creator any `json:"creator,omitempty"`
    CreatorID any `json:"creator_id,omitempty"`
    Description *string `json:"description,omitempty"`
    EntityID any `json:"entity_id,omitempty"`
    EntityMetadata any `json:"entity_metadata"`
    EntityType int32 `json:"entity_type"`
    GuildID string `json:"guild_id"`
    Image *string `json:"image,omitempty"`
    Name string `json:"name"`
    PrivacyLevel int32 `json:"privacy_level"`
    ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
    ScheduledStartTime time.Time `json:"scheduled_start_time"`
    Status int32 `json:"status"`
    UserCount int32 `json:"user_count,omitempty"`
    UserRsvp any `json:"user_rsvp,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ChannelID | `any` |  |
| Creator | `any` |  |
| CreatorID | `any` |  |
| Description | `*string` |  |
| EntityID | `any` |  |
| EntityMetadata | `any` |  |
| EntityType | `int32` |  |
| GuildID | `string` |  |
| Image | `*string` |  |
| Name | `string` |  |
| PrivacyLevel | `int32` |  |
| ScheduledEndTime | `*time.Time` |  |
| ScheduledStartTime | `time.Time` |  |
| Status | `int32` |  |
| UserCount | `int32` |  |
| UserRsvp | `any` |  |

### ExternalScheduledEventCreateOptions
ExternalScheduledEventCreateOptions

#### Example Usage

```go
// Create a new ExternalScheduledEventCreateOptions
externalscheduledeventcreateoptions := ExternalScheduledEventCreateOptions{
    ChannelID: any{},
    Description: &"example"{},
    EntityMetadata: any{},
    EntityType: 42,
    Image: &"example"{},
    Name: "example",
    PrivacyLevel: 42,
    ScheduledEndTime: &/* value */{},
    ScheduledStartTime: /* value */,
}
```

#### Type Definition

```go
type ExternalScheduledEventCreateOptions struct {
    ChannelID any `json:"channel_id,omitempty"`
    Description *string `json:"description,omitempty"`
    EntityMetadata any `json:"entity_metadata"`
    EntityType int32 `json:"entity_type"`
    Image *string `json:"image,omitempty"`
    Name string `json:"name"`
    PrivacyLevel int32 `json:"privacy_level"`
    ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
    ScheduledStartTime time.Time `json:"scheduled_start_time"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `any` |  |
| Description | `*string` |  |
| EntityMetadata | `any` |  |
| EntityType | `int32` |  |
| Image | `*string` |  |
| Name | `string` |  |
| PrivacyLevel | `int32` |  |
| ScheduledEndTime | `*time.Time` |  |
| ScheduledStartTime | `time.Time` |  |

### ExternalScheduledEventPatchOptionsPartial
ExternalScheduledEventPatchOptionsPartial

#### Example Usage

```go
// Create a new ExternalScheduledEventPatchOptionsPartial
externalscheduledeventpatchoptionspartial := ExternalScheduledEventPatchOptionsPartial{
    ChannelID: any{},
    Description: &"example"{},
    EntityMetadata: any{},
    EntityType: any{},
    Image: &"example"{},
    Name: "example",
    PrivacyLevel: 42,
    ScheduledEndTime: &/* value */{},
    ScheduledStartTime: /* value */,
    Status: any{},
}
```

#### Type Definition

```go
type ExternalScheduledEventPatchOptionsPartial struct {
    ChannelID any `json:"channel_id,omitempty"`
    Description *string `json:"description,omitempty"`
    EntityMetadata any `json:"entity_metadata,omitempty"`
    EntityType any `json:"entity_type,omitempty"`
    Image *string `json:"image,omitempty"`
    Name string `json:"name,omitempty"`
    PrivacyLevel int32 `json:"privacy_level,omitempty"`
    ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
    ScheduledStartTime time.Time `json:"scheduled_start_time,omitempty"`
    Status any `json:"status,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `any` |  |
| Description | `*string` |  |
| EntityMetadata | `any` |  |
| EntityType | `any` |  |
| Image | `*string` |  |
| Name | `string` |  |
| PrivacyLevel | `int32` |  |
| ScheduledEndTime | `*time.Time` |  |
| ScheduledStartTime | `time.Time` |  |
| Status | `any` |  |

### FileComponent
FileComponent

#### Example Usage

```go
// Create a new FileComponent
filecomponent := FileComponent{
    ID: 42,
    File: any{},
    Name: &"example"{},
    Size: &42{},
    Spoiler: true,
    Type: 42,
}
```

#### Type Definition

```go
type FileComponent struct {
    ID int32 `json:"id"`
    File any `json:"file"`
    Name *string `json:"name,omitempty"`
    Size *int32 `json:"size,omitempty"`
    Spoiler bool `json:"spoiler"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| File | `any` |  |
| Name | `*string` |  |
| Size | `*int32` |  |
| Spoiler | `bool` |  |
| Type | `int32` |  |

### FileComponentForMessageOptions
FileComponentForMessageOptions

#### Example Usage

```go
// Create a new FileComponentForMessageOptions
filecomponentformessageoptions := FileComponentForMessageOptions{
    ID: &42{},
    File: any{},
    Spoiler: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type FileComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    File any `json:"file"`
    Spoiler *bool `json:"spoiler,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| File | `any` |  |
| Spoiler | `*bool` |  |
| Type | `int32` |  |

### FileUploadComponentForModalOptions
FileUploadComponentForModalOptions

#### Example Usage

```go
// Create a new FileUploadComponentForModalOptions
fileuploadcomponentformodaloptions := FileUploadComponentForModalOptions{
    ID: &42{},
    CustomID: "example",
    MaxValues: &42{},
    MinValues: &42{},
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type FileUploadComponentForModalOptions struct {
    ID *int32 `json:"id,omitempty"`
    CustomID string `json:"custom_id"`
    MaxValues *int64 `json:"max_values,omitempty"`
    MinValues *int64 `json:"min_values,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| CustomID | `string` |  |
| MaxValues | `*int64` |  |
| MinValues | `*int64` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### FlagToChannelAction
FlagToChannelAction

#### Example Usage

```go
// Create a new FlagToChannelAction
flagtochannelaction := FlagToChannelAction{
    Metadata: any{},
    Type: 42,
}
```

#### Type Definition

```go
type FlagToChannelAction struct {
    Metadata any `json:"metadata"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Metadata | `any` |  |
| Type | `int32` |  |

### FlagToChannelActionMetadata
FlagToChannelActionMetadata

#### Example Usage

```go
// Create a new FlagToChannelActionMetadata
flagtochannelactionmetadata := FlagToChannelActionMetadata{
    ChannelID: "example",
}
```

#### Type Definition

```go
type FlagToChannelActionMetadata struct {
    ChannelID string `json:"channel_id"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `string` |  |

### ForumLayout
ForumLayout

#### Example Usage

```go
// Example usage of ForumLayout
var value ForumLayout
// Initialize with appropriate value
```

#### Type Definition

```go
type ForumLayout int
```

### ForumTag
ForumTag

#### Example Usage

```go
// Create a new ForumTag
forumtag := ForumTag{
    ID: "example",
    EmojiID: any{},
    EmojiName: &"example"{},
    Moderated: true,
    Name: "example",
}
```

#### Type Definition

```go
type ForumTag struct {
    ID string `json:"id"`
    EmojiID any `json:"emoji_id,omitempty"`
    EmojiName *string `json:"emoji_name,omitempty"`
    Moderated bool `json:"moderated"`
    Name string `json:"name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| EmojiID | `any` |  |
| EmojiName | `*string` |  |
| Moderated | `bool` |  |
| Name | `string` |  |

### FriendInvite
FriendInvite

#### Example Usage

```go
// Create a new FriendInvite
friendinvite := FriendInvite{
    Channel: any{},
    Code: "example",
    CreatedAt: /* value */,
    ExpiresAt: &/* value */{},
    Flags: 42,
    FriendsCount: 42,
    Inviter: any{},
    IsContact: true,
    MaxAge: 42,
    MaxUses: 42,
    Type: 42,
    Uses: 42,
}
```

#### Type Definition

```go
type FriendInvite struct {
    Channel any `json:"channel,omitempty"`
    Code string `json:"code"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    ExpiresAt *time.Time `json:"expires_at,omitempty"`
    Flags int32 `json:"flags,omitempty"`
    FriendsCount int32 `json:"friends_count,omitempty"`
    Inviter any `json:"inviter,omitempty"`
    IsContact bool `json:"is_contact,omitempty"`
    MaxAge int32 `json:"max_age,omitempty"`
    MaxUses int32 `json:"max_uses,omitempty"`
    Type int32 `json:"type"`
    Uses int32 `json:"uses,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Channel | `any` |  |
| Code | `string` |  |
| CreatedAt | `time.Time` |  |
| ExpiresAt | `*time.Time` |  |
| Flags | `int32` |  |
| FriendsCount | `int32` |  |
| Inviter | `any` |  |
| IsContact | `bool` |  |
| MaxAge | `int32` |  |
| MaxUses | `int32` |  |
| Type | `int32` |  |
| Uses | `int32` |  |

### Gateway
Gateway

#### Example Usage

```go
// Create a new Gateway
gateway := Gateway{
    URL: "example",
}
```

#### Type Definition

```go
type Gateway struct {
    URL string `json:"url"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| URL | `string` |  |

### GatewayBot
GatewayBot

#### Example Usage

```go
// Create a new GatewayBot
gatewaybot := GatewayBot{
    SessionStartLimit: any{},
    Shards: 42,
    URL: "example",
}
```

#### Type Definition

```go
type GatewayBot struct {
    SessionStartLimit any `json:"session_start_limit"`
    Shards int32 `json:"shards"`
    URL string `json:"url"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| SessionStartLimit | `any` |  |
| Shards | `int32` |  |
| URL | `string` |  |

### GatewayBotSessionStartLimit
GatewayBotSessionStartLimit

#### Example Usage

```go
// Create a new GatewayBotSessionStartLimit
gatewaybotsessionstartlimit := GatewayBotSessionStartLimit{
    MaxConcurrency: 42,
    Remaining: 42,
    ResetAfter: 42,
    Total: 42,
}
```

#### Type Definition

```go
type GatewayBotSessionStartLimit struct {
    MaxConcurrency int32 `json:"max_concurrency"`
    Remaining int32 `json:"remaining"`
    ResetAfter int32 `json:"reset_after"`
    Total int32 `json:"total"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| MaxConcurrency | `int32` |  |
| Remaining | `int32` |  |
| ResetAfter | `int32` |  |
| Total | `int32` |  |

### GithubAuthor
GithubAuthor

#### Example Usage

```go
// Create a new GithubAuthor
githubauthor := GithubAuthor{
    Name: "example",
    Username: &"example"{},
}
```

#### Type Definition

```go
type GithubAuthor struct {
    Name string `json:"name"`
    Username *string `json:"username,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| Username | `*string` |  |

### GithubCheckApp
GithubCheckApp

#### Example Usage

```go
// Create a new GithubCheckApp
githubcheckapp := GithubCheckApp{
    Name: "example",
}
```

#### Type Definition

```go
type GithubCheckApp struct {
    Name string `json:"name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |

### GithubCheckPullOptions
GithubCheckPullOptions

#### Example Usage

```go
// Create a new GithubCheckPullOptions
githubcheckpulloptions := GithubCheckPullOptions{
    Number: 42,
}
```

#### Type Definition

```go
type GithubCheckPullOptions struct {
    Number int64 `json:"number"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Number | `int64` |  |

### GithubCheckRun
GithubCheckRun

#### Example Usage

```go
// Create a new GithubCheckRun
githubcheckrun := GithubCheckRun{
    CheckSuite: any{},
    Conclusion: &"example"{},
    DetailsURL: &"example"{},
    HTMLURL: "example",
    Name: "example",
    Output: any{},
    PullRequests: [],
}
```

#### Type Definition

```go
type GithubCheckRun struct {
    CheckSuite any `json:"check_suite"`
    Conclusion *string `json:"conclusion,omitempty"`
    DetailsURL *string `json:"details_url,omitempty"`
    HTMLURL string `json:"html_url"`
    Name string `json:"name"`
    Output any `json:"output,omitempty"`
    PullRequests []any `json:"pull_requests,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| CheckSuite | `any` |  |
| Conclusion | `*string` |  |
| DetailsURL | `*string` |  |
| HTMLURL | `string` |  |
| Name | `string` |  |
| Output | `any` |  |
| PullRequests | `[]any` |  |

### GithubCheckRunOutput
GithubCheckRunOutput

#### Example Usage

```go
// Create a new GithubCheckRunOutput
githubcheckrunoutput := GithubCheckRunOutput{
    Summary: &"example"{},
    Title: &"example"{},
}
```

#### Type Definition

```go
type GithubCheckRunOutput struct {
    Summary *string `json:"summary,omitempty"`
    Title *string `json:"title,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Summary | `*string` |  |
| Title | `*string` |  |

### GithubCheckSuite
GithubCheckSuite

#### Example Usage

```go
// Create a new GithubCheckSuite
githubchecksuite := GithubCheckSuite{
    App: any{},
    Conclusion: &"example"{},
    HeadBranch: &"example"{},
    HeadSha: "example",
    PullRequests: [],
}
```

#### Type Definition

```go
type GithubCheckSuite struct {
    App any `json:"app"`
    Conclusion *string `json:"conclusion,omitempty"`
    HeadBranch *string `json:"head_branch,omitempty"`
    HeadSha string `json:"head_sha"`
    PullRequests []any `json:"pull_requests,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| App | `any` |  |
| Conclusion | `*string` |  |
| HeadBranch | `*string` |  |
| HeadSha | `string` |  |
| PullRequests | `[]any` |  |

### GithubComment
GithubComment

#### Example Usage

```go
// Create a new GithubComment
githubcomment := GithubComment{
    ID: 42,
    Body: "example",
    CommitID: &"example"{},
    HTMLURL: "example",
    User: any{},
}
```

#### Type Definition

```go
type GithubComment struct {
    ID int64 `json:"id"`
    Body string `json:"body"`
    CommitID *string `json:"commit_id,omitempty"`
    HTMLURL string `json:"html_url"`
    User any `json:"user"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int64` |  |
| Body | `string` |  |
| CommitID | `*string` |  |
| HTMLURL | `string` |  |
| User | `any` |  |

### GithubCommit
GithubCommit

#### Example Usage

```go
// Create a new GithubCommit
githubcommit := GithubCommit{
    ID: "example",
    Author: any{},
    Message: "example",
    URL: "example",
}
```

#### Type Definition

```go
type GithubCommit struct {
    ID string `json:"id"`
    Author any `json:"author"`
    Message string `json:"message"`
    URL string `json:"url"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Author | `any` |  |
| Message | `string` |  |
| URL | `string` |  |

### GithubDiscussion
GithubDiscussion

#### Example Usage

```go
// Create a new GithubDiscussion
githubdiscussion := GithubDiscussion{
    AnswerHTMLURL: &"example"{},
    Body: &"example"{},
    HTMLURL: "example",
    Number: 42,
    Title: "example",
    User: any{},
}
```

#### Type Definition

```go
type GithubDiscussion struct {
    AnswerHTMLURL *string `json:"answer_html_url,omitempty"`
    Body *string `json:"body,omitempty"`
    HTMLURL string `json:"html_url"`
    Number int64 `json:"number"`
    Title string `json:"title"`
    User any `json:"user"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AnswerHTMLURL | `*string` |  |
| Body | `*string` |  |
| HTMLURL | `string` |  |
| Number | `int64` |  |
| Title | `string` |  |
| User | `any` |  |

### GithubIssue
GithubIssue

#### Example Usage

```go
// Create a new GithubIssue
githubissue := GithubIssue{
    ID: 42,
    Body: &"example"{},
    HTMLURL: "example",
    Number: 42,
    PullRequest: any{},
    Title: "example",
    User: any{},
}
```

#### Type Definition

```go
type GithubIssue struct {
    ID int64 `json:"id"`
    Body *string `json:"body,omitempty"`
    HTMLURL string `json:"html_url"`
    Number int64 `json:"number"`
    PullRequest any `json:"pull_request,omitempty"`
    Title string `json:"title"`
    User any `json:"user"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int64` |  |
| Body | `*string` |  |
| HTMLURL | `string` |  |
| Number | `int64` |  |
| PullRequest | `any` |  |
| Title | `string` |  |
| User | `any` |  |

### GithubRelease
GithubRelease

#### Example Usage

```go
// Create a new GithubRelease
githubrelease := GithubRelease{
    ID: 42,
    Author: any{},
    HTMLURL: "example",
    TagName: "example",
}
```

#### Type Definition

```go
type GithubRelease struct {
    ID int64 `json:"id"`
    Author any `json:"author"`
    HTMLURL string `json:"html_url"`
    TagName string `json:"tag_name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int64` |  |
| Author | `any` |  |
| HTMLURL | `string` |  |
| TagName | `string` |  |

### GithubRepository
GithubRepository

#### Example Usage

```go
// Create a new GithubRepository
githubrepository := GithubRepository{
    ID: 42,
    FullName: "example",
    HTMLURL: "example",
    Name: "example",
}
```

#### Type Definition

```go
type GithubRepository struct {
    ID int64 `json:"id"`
    FullName string `json:"full_name"`
    HTMLURL string `json:"html_url"`
    Name string `json:"name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int64` |  |
| FullName | `string` |  |
| HTMLURL | `string` |  |
| Name | `string` |  |

### GithubReview
GithubReview

#### Example Usage

```go
// Create a new GithubReview
githubreview := GithubReview{
    Body: &"example"{},
    HTMLURL: "example",
    State: "example",
    User: any{},
}
```

#### Type Definition

```go
type GithubReview struct {
    Body *string `json:"body,omitempty"`
    HTMLURL string `json:"html_url"`
    State string `json:"state"`
    User any `json:"user"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Body | `*string` |  |
| HTMLURL | `string` |  |
| State | `string` |  |
| User | `any` |  |

### GithubUser
GithubUser

#### Example Usage

```go
// Create a new GithubUser
githubuser := GithubUser{
    ID: 42,
    AvatarURL: "example",
    HTMLURL: "example",
    Login: "example",
}
```

#### Type Definition

```go
type GithubUser struct {
    ID int64 `json:"id"`
    AvatarURL string `json:"avatar_url"`
    HTMLURL string `json:"html_url"`
    Login string `json:"login"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int64` |  |
| AvatarURL | `string` |  |
| HTMLURL | `string` |  |
| Login | `string` |  |

### GithubWebhook
GithubWebhook

#### Example Usage

```go
// Create a new GithubWebhook
githubwebhook := GithubWebhook{
    Action: &"example"{},
    Answer: any{},
    CheckRun: any{},
    CheckSuite: any{},
    Comment: any{},
    Commits: [],
    Compare: &"example"{},
    Discussion: any{},
    Forced: &true{},
    Forkee: any{},
    HeadCommit: any{},
    Issue: any{},
    Member: any{},
    PullRequest: any{},
    Ref: &"example"{},
    RefType: &"example"{},
    Release: any{},
    Repository: any{},
    Review: any{},
    Sender: any{},
}
```

#### Type Definition

```go
type GithubWebhook struct {
    Action *string `json:"action,omitempty"`
    Answer any `json:"answer,omitempty"`
    CheckRun any `json:"check_run,omitempty"`
    CheckSuite any `json:"check_suite,omitempty"`
    Comment any `json:"comment,omitempty"`
    Commits []any `json:"commits,omitempty"`
    Compare *string `json:"compare,omitempty"`
    Discussion any `json:"discussion,omitempty"`
    Forced *bool `json:"forced,omitempty"`
    Forkee any `json:"forkee,omitempty"`
    HeadCommit any `json:"head_commit,omitempty"`
    Issue any `json:"issue,omitempty"`
    Member any `json:"member,omitempty"`
    PullRequest any `json:"pull_request,omitempty"`
    Ref *string `json:"ref,omitempty"`
    RefType *string `json:"ref_type,omitempty"`
    Release any `json:"release,omitempty"`
    Repository any `json:"repository,omitempty"`
    Review any `json:"review,omitempty"`
    Sender any `json:"sender"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Action | `*string` |  |
| Answer | `any` |  |
| CheckRun | `any` |  |
| CheckSuite | `any` |  |
| Comment | `any` |  |
| Commits | `[]any` |  |
| Compare | `*string` |  |
| Discussion | `any` |  |
| Forced | `*bool` |  |
| Forkee | `any` |  |
| HeadCommit | `any` |  |
| Issue | `any` |  |
| Member | `any` |  |
| PullRequest | `any` |  |
| Ref | `*string` |  |
| RefType | `*string` |  |
| Release | `any` |  |
| Repository | `any` |  |
| Review | `any` |  |
| Sender | `any` |  |

### GroupDMInvite
GroupDMInvite

#### Example Usage

```go
// Create a new GroupDMInvite
groupdminvite := GroupDMInvite{
    ApproximateMemberCount: &42{},
    Channel: any{},
    Code: "example",
    CreatedAt: /* value */,
    ExpiresAt: &/* value */{},
    Inviter: any{},
    MaxAge: 42,
    Type: 42,
}
```

#### Type Definition

```go
type GroupDMInvite struct {
    ApproximateMemberCount *int32 `json:"approximate_member_count,omitempty"`
    Channel any `json:"channel"`
    Code string `json:"code"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    ExpiresAt *time.Time `json:"expires_at,omitempty"`
    Inviter any `json:"inviter,omitempty"`
    MaxAge int32 `json:"max_age,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ApproximateMemberCount | `*int32` |  |
| Channel | `any` |  |
| Code | `string` |  |
| CreatedAt | `time.Time` |  |
| ExpiresAt | `*time.Time` |  |
| Inviter | `any` |  |
| MaxAge | `int32` |  |
| Type | `int32` |  |

### Guild
Guild

#### Example Usage

```go
// Create a new Guild
guild := Guild{
    ID: "example",
    AfkChannelID: any{},
    AfkTimeout: 42,
    ApplicationID: any{},
    Banner: &"example"{},
    DefaultMessageNotifications: 42,
    Description: &"example"{},
    DiscoverySplash: &"example"{},
    Emojis: [],
    ExplicitContentFilter: 42,
    Features: [],
    HomeHeader: &"example"{},
    Icon: &"example"{},
    MaxMembers: 42,
    MaxPresences: &42{},
    MaxStageVideoChannelUsers: 42,
    MaxVideoChannelUsers: 42,
    MfaLevel: 42,
    Name: "example",
    Nsfw: true,
    NsfwLevel: 42,
    OwnerID: "example",
    PreferredLocale: "example",
    PremiumProgressBarEnabled: true,
    PremiumSubscriptionCount: 42,
    PremiumTier: 42,
    PublicUpdatesChannelID: any{},
    Region: "example",
    Roles: [],
    RulesChannelID: any{},
    SafetyAlertsChannelID: any{},
    Splash: &"example"{},
    Stickers: [],
    SystemChannelFlags: 42,
    SystemChannelID: any{},
    VanityURLCode: &"example"{},
    VerificationLevel: 42,
    WidgetChannelID: any{},
    WidgetEnabled: true,
}
```

#### Type Definition

```go
type Guild struct {
    ID string `json:"id"`
    AfkChannelID any `json:"afk_channel_id,omitempty"`
    AfkTimeout int32 `json:"afk_timeout"`
    ApplicationID any `json:"application_id,omitempty"`
    Banner *string `json:"banner,omitempty"`
    DefaultMessageNotifications int32 `json:"default_message_notifications"`
    Description *string `json:"description,omitempty"`
    DiscoverySplash *string `json:"discovery_splash,omitempty"`
    Emojis []any `json:"emojis"`
    ExplicitContentFilter int32 `json:"explicit_content_filter"`
    Features []string `json:"features"`
    HomeHeader *string `json:"home_header,omitempty"`
    Icon *string `json:"icon,omitempty"`
    MaxMembers int32 `json:"max_members"`
    MaxPresences *int32 `json:"max_presences,omitempty"`
    MaxStageVideoChannelUsers int32 `json:"max_stage_video_channel_users"`
    MaxVideoChannelUsers int32 `json:"max_video_channel_users"`
    MfaLevel int32 `json:"mfa_level"`
    Name string `json:"name"`
    Nsfw bool `json:"nsfw"`
    NsfwLevel int32 `json:"nsfw_level"`
    OwnerID string `json:"owner_id"`
    PreferredLocale string `json:"preferred_locale"`
    PremiumProgressBarEnabled bool `json:"premium_progress_bar_enabled"`
    PremiumSubscriptionCount int32 `json:"premium_subscription_count"`
    PremiumTier int32 `json:"premium_tier"`
    PublicUpdatesChannelID any `json:"public_updates_channel_id,omitempty"`
    Region string `json:"region"`
    Roles []any `json:"roles"`
    RulesChannelID any `json:"rules_channel_id,omitempty"`
    SafetyAlertsChannelID any `json:"safety_alerts_channel_id,omitempty"`
    Splash *string `json:"splash,omitempty"`
    Stickers []any `json:"stickers"`
    SystemChannelFlags int32 `json:"system_channel_flags"`
    SystemChannelID any `json:"system_channel_id,omitempty"`
    VanityURLCode *string `json:"vanity_url_code,omitempty"`
    VerificationLevel int32 `json:"verification_level"`
    WidgetChannelID any `json:"widget_channel_id,omitempty"`
    WidgetEnabled bool `json:"widget_enabled"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| AfkChannelID | `any` |  |
| AfkTimeout | `int32` |  |
| ApplicationID | `any` |  |
| Banner | `*string` |  |
| DefaultMessageNotifications | `int32` |  |
| Description | `*string` |  |
| DiscoverySplash | `*string` |  |
| Emojis | `[]any` |  |
| ExplicitContentFilter | `int32` |  |
| Features | `[]string` |  |
| HomeHeader | `*string` |  |
| Icon | `*string` |  |
| MaxMembers | `int32` |  |
| MaxPresences | `*int32` |  |
| MaxStageVideoChannelUsers | `int32` |  |
| MaxVideoChannelUsers | `int32` |  |
| MfaLevel | `int32` |  |
| Name | `string` |  |
| Nsfw | `bool` |  |
| NsfwLevel | `int32` |  |
| OwnerID | `string` |  |
| PreferredLocale | `string` |  |
| PremiumProgressBarEnabled | `bool` |  |
| PremiumSubscriptionCount | `int32` |  |
| PremiumTier | `int32` |  |
| PublicUpdatesChannelID | `any` |  |
| Region | `string` |  |
| Roles | `[]any` |  |
| RulesChannelID | `any` |  |
| SafetyAlertsChannelID | `any` |  |
| Splash | `*string` |  |
| Stickers | `[]any` |  |
| SystemChannelFlags | `int32` |  |
| SystemChannelID | `any` |  |
| VanityURLCode | `*string` |  |
| VerificationLevel | `int32` |  |
| WidgetChannelID | `any` |  |
| WidgetEnabled | `bool` |  |

### GuildAuditLog
GuildAuditLog

#### Example Usage

```go
// Create a new GuildAuditLog
guildauditlog := GuildAuditLog{
    ApplicationCommands: [],
    AuditLogEntries: [],
    AutoModerationRules: [],
    GuildScheduledEvents: [],
    Integrations: [],
    Threads: [],
    Users: [],
    Webhooks: [],
}
```

#### Type Definition

```go
type GuildAuditLog struct {
    ApplicationCommands []any `json:"application_commands"`
    AuditLogEntries []any `json:"audit_log_entries"`
    AutoModerationRules []any `json:"auto_moderation_rules"`
    GuildScheduledEvents []any `json:"guild_scheduled_events"`
    Integrations []any `json:"integrations"`
    Threads []any `json:"threads"`
    Users []any `json:"users"`
    Webhooks []any `json:"webhooks"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ApplicationCommands | `[]any` |  |
| AuditLogEntries | `[]any` |  |
| AutoModerationRules | `[]any` |  |
| GuildScheduledEvents | `[]any` |  |
| Integrations | `[]any` |  |
| Threads | `[]any` |  |
| Users | `[]any` |  |
| Webhooks | `[]any` |  |

### GuildBan
GuildBan

#### Example Usage

```go
// Create a new GuildBan
guildban := GuildBan{
    Reason: &"example"{},
    User: any{},
}
```

#### Type Definition

```go
type GuildBan struct {
    Reason *string `json:"reason,omitempty"`
    User any `json:"user"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Reason | `*string` |  |
| User | `any` |  |

### GuildChannel
GuildChannel

#### Example Usage

```go
// Create a new GuildChannel
guildchannel := GuildChannel{
    ID: "example",
    AvailableTags: [],
    Bitrate: 42,
    DefaultAutoArchiveDuration: 42,
    DefaultForumLayout: any{},
    DefaultReactionEmoji: any{},
    DefaultSortOrder: any{},
    DefaultTagSetting: any{},
    DefaultThreadRateLimitPerUser: 42,
    Flags: 42,
    GuildID: "example",
    HdStreamingBuyerID: "example",
    HdStreamingUntil: /* value */,
    LastMessageID: any{},
    LastPinTimestamp: &/* value */{},
    Name: "example",
    Nsfw: true,
    ParentID: any{},
    PermissionOverwrites: [],
    Permissions: &"example"{},
    Position: 42,
    RateLimitPerUser: 42,
    RtcRegion: &"example"{},
    Topic: &"example"{},
    Type: 42,
    UserLimit: 42,
    VideoQualityMode: 42,
}
```

#### Type Definition

```go
type GuildChannel struct {
    ID string `json:"id"`
    AvailableTags []any `json:"available_tags,omitempty"`
    Bitrate int32 `json:"bitrate,omitempty"`
    DefaultAutoArchiveDuration int32 `json:"default_auto_archive_duration,omitempty"`
    DefaultForumLayout any `json:"default_forum_layout,omitempty"`
    DefaultReactionEmoji any `json:"default_reaction_emoji,omitempty"`
    DefaultSortOrder any `json:"default_sort_order,omitempty"`
    DefaultTagSetting any `json:"default_tag_setting,omitempty"`
    DefaultThreadRateLimitPerUser int32 `json:"default_thread_rate_limit_per_user,omitempty"`
    Flags int32 `json:"flags"`
    GuildID string `json:"guild_id"`
    HdStreamingBuyerID string `json:"hd_streaming_buyer_id,omitempty"`
    HdStreamingUntil time.Time `json:"hd_streaming_until,omitempty"`
    LastMessageID any `json:"last_message_id,omitempty"`
    LastPinTimestamp *time.Time `json:"last_pin_timestamp,omitempty"`
    Name string `json:"name"`
    Nsfw bool `json:"nsfw,omitempty"`
    ParentID any `json:"parent_id,omitempty"`
    PermissionOverwrites []any `json:"permission_overwrites,omitempty"`
    Permissions *string `json:"permissions,omitempty"`
    Position int32 `json:"position"`
    RateLimitPerUser int32 `json:"rate_limit_per_user,omitempty"`
    RtcRegion *string `json:"rtc_region,omitempty"`
    Topic *string `json:"topic,omitempty"`
    Type int32 `json:"type"`
    UserLimit int32 `json:"user_limit,omitempty"`
    VideoQualityMode int32 `json:"video_quality_mode,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| AvailableTags | `[]any` |  |
| Bitrate | `int32` |  |
| DefaultAutoArchiveDuration | `int32` |  |
| DefaultForumLayout | `any` |  |
| DefaultReactionEmoji | `any` |  |
| DefaultSortOrder | `any` |  |
| DefaultTagSetting | `any` |  |
| DefaultThreadRateLimitPerUser | `int32` |  |
| Flags | `int32` |  |
| GuildID | `string` |  |
| HdStreamingBuyerID | `string` |  |
| HdStreamingUntil | `time.Time` |  |
| LastMessageID | `any` |  |
| LastPinTimestamp | `*time.Time` |  |
| Name | `string` |  |
| Nsfw | `bool` |  |
| ParentID | `any` |  |
| PermissionOverwrites | `[]any` |  |
| Permissions | `*string` |  |
| Position | `int32` |  |
| RateLimitPerUser | `int32` |  |
| RtcRegion | `*string` |  |
| Topic | `*string` |  |
| Type | `int32` |  |
| UserLimit | `int32` |  |
| VideoQualityMode | `int32` |  |

### GuildChannelLocation
GuildChannelLocation

#### Example Usage

```go
// Create a new GuildChannelLocation
guildchannellocation := GuildChannelLocation{
    ID: "example",
    ChannelID: "example",
    GuildID: "example",
    Kind: "example",
}
```

#### Type Definition

```go
type GuildChannelLocation struct {
    ID string `json:"id"`
    ChannelID string `json:"channel_id"`
    GuildID string `json:"guild_id"`
    Kind string `json:"kind"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ChannelID | `string` |  |
| GuildID | `string` |  |
| Kind | `string` |  |

### GuildExplicitContentFilterTypes
GuildExplicitContentFilterTypes

#### Example Usage

```go
// Example usage of GuildExplicitContentFilterTypes
var value GuildExplicitContentFilterTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type GuildExplicitContentFilterTypes int
```

### GuildFeatures
GuildFeatures

#### Example Usage

```go
// Example usage of GuildFeatures
var value GuildFeatures
// Initialize with appropriate value
```

#### Type Definition

```go
type GuildFeatures string
```

### GuildHomeSettings
GuildHomeSettings

#### Example Usage

```go
// Create a new GuildHomeSettings
guildhomesettings := GuildHomeSettings{
    Enabled: true,
    GuildID: "example",
    NewMemberActions: [],
    ResourceChannels: [],
    WelcomeMessage: any{},
}
```

#### Type Definition

```go
type GuildHomeSettings struct {
    Enabled bool `json:"enabled"`
    GuildID string `json:"guild_id"`
    NewMemberActions []any `json:"new_member_actions"`
    ResourceChannels []any `json:"resource_channels"`
    WelcomeMessage any `json:"welcome_message,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Enabled | `bool` |  |
| GuildID | `string` |  |
| NewMemberActions | `[]any` |  |
| ResourceChannels | `[]any` |  |
| WelcomeMessage | `any` |  |

### GuildIncomingWebhook
GuildIncomingWebhook

#### Example Usage

```go
// Create a new GuildIncomingWebhook
guildincomingwebhook := GuildIncomingWebhook{
    ID: "example",
    ApplicationID: any{},
    Avatar: &"example"{},
    ChannelID: any{},
    GuildID: any{},
    Name: "example",
    Token: "example",
    Type: 42,
    URL: "example",
    User: any{},
}
```

#### Type Definition

```go
type GuildIncomingWebhook struct {
    ID string `json:"id"`
    ApplicationID any `json:"application_id,omitempty"`
    Avatar *string `json:"avatar,omitempty"`
    ChannelID any `json:"channel_id,omitempty"`
    GuildID any `json:"guild_id,omitempty"`
    Name string `json:"name"`
    Token string `json:"token,omitempty"`
    Type int32 `json:"type"`
    URL string `json:"url,omitempty"`
    User any `json:"user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ApplicationID | `any` |  |
| Avatar | `*string` |  |
| ChannelID | `any` |  |
| GuildID | `any` |  |
| Name | `string` |  |
| Token | `string` |  |
| Type | `int32` |  |
| URL | `string` |  |
| User | `any` |  |

### GuildInvite
GuildInvite

#### Example Usage

```go
// Create a new GuildInvite
guildinvite := GuildInvite{
    ApproximateMemberCount: &42{},
    ApproximatePresenceCount: &42{},
    Channel: any{},
    Code: "example",
    CreatedAt: /* value */,
    ExpiresAt: &/* value */{},
    Flags: 42,
    Guild: any{},
    GuildID: "example",
    GuildScheduledEvent: any{},
    Inviter: any{},
    IsContact: true,
    IsNicknameChangeable: true,
    MaxAge: 42,
    MaxUses: 42,
    TargetApplication: any{},
    TargetType: 42,
    TargetUser: any{},
    Temporary: true,
    Type: 42,
    Uses: 42,
}
```

#### Type Definition

```go
type GuildInvite struct {
    ApproximateMemberCount *int32 `json:"approximate_member_count,omitempty"`
    ApproximatePresenceCount *int32 `json:"approximate_presence_count,omitempty"`
    Channel any `json:"channel"`
    Code string `json:"code"`
    CreatedAt time.Time `json:"created_at,omitempty"`
    ExpiresAt *time.Time `json:"expires_at,omitempty"`
    Flags int32 `json:"flags,omitempty"`
    Guild any `json:"guild"`
    GuildID string `json:"guild_id"`
    GuildScheduledEvent any `json:"guild_scheduled_event,omitempty"`
    Inviter any `json:"inviter,omitempty"`
    IsContact bool `json:"is_contact,omitempty"`
    IsNicknameChangeable bool `json:"is_nickname_changeable,omitempty"`
    MaxAge int32 `json:"max_age,omitempty"`
    MaxUses int32 `json:"max_uses,omitempty"`
    TargetApplication any `json:"target_application,omitempty"`
    TargetType int32 `json:"target_type,omitempty"`
    TargetUser any `json:"target_user,omitempty"`
    Temporary bool `json:"temporary,omitempty"`
    Type int32 `json:"type"`
    Uses int32 `json:"uses,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ApproximateMemberCount | `*int32` |  |
| ApproximatePresenceCount | `*int32` |  |
| Channel | `any` |  |
| Code | `string` |  |
| CreatedAt | `time.Time` |  |
| ExpiresAt | `*time.Time` |  |
| Flags | `int32` |  |
| Guild | `any` |  |
| GuildID | `string` |  |
| GuildScheduledEvent | `any` |  |
| Inviter | `any` |  |
| IsContact | `bool` |  |
| IsNicknameChangeable | `bool` |  |
| MaxAge | `int32` |  |
| MaxUses | `int32` |  |
| TargetApplication | `any` |  |
| TargetType | `int32` |  |
| TargetUser | `any` |  |
| Temporary | `bool` |  |
| Type | `int32` |  |
| Uses | `int32` |  |

### GuildMFALevel
GuildMFALevel

#### Example Usage

```go
// Example usage of GuildMFALevel
var value GuildMFALevel
// Initialize with appropriate value
```

#### Type Definition

```go
type GuildMFALevel int
```

### GuildMember
GuildMember

#### Example Usage

```go
// Create a new GuildMember
guildmember := GuildMember{
    Avatar: &"example"{},
    AvatarDecorationData: any{},
    Banner: &"example"{},
    Collectibles: any{},
    CommunicationDisabledUntil: &/* value */{},
    Deaf: true,
    Flags: 42,
    JoinedAt: /* value */,
    Mute: true,
    Nick: &"example"{},
    Pending: true,
    PremiumSince: &/* value */{},
    Roles: [],
    User: any{},
}
```

#### Type Definition

```go
type GuildMember struct {
    Avatar *string `json:"avatar,omitempty"`
    AvatarDecorationData any `json:"avatar_decoration_data,omitempty"`
    Banner *string `json:"banner,omitempty"`
    Collectibles any `json:"collectibles,omitempty"`
    CommunicationDisabledUntil *time.Time `json:"communication_disabled_until,omitempty"`
    Deaf bool `json:"deaf"`
    Flags int32 `json:"flags"`
    JoinedAt time.Time `json:"joined_at"`
    Mute bool `json:"mute"`
    Nick *string `json:"nick,omitempty"`
    Pending bool `json:"pending"`
    PremiumSince *time.Time `json:"premium_since,omitempty"`
    Roles []string `json:"roles"`
    User any `json:"user"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Avatar | `*string` |  |
| AvatarDecorationData | `any` |  |
| Banner | `*string` |  |
| Collectibles | `any` |  |
| CommunicationDisabledUntil | `*time.Time` |  |
| Deaf | `bool` |  |
| Flags | `int32` |  |
| JoinedAt | `time.Time` |  |
| Mute | `bool` |  |
| Nick | `*string` |  |
| Pending | `bool` |  |
| PremiumSince | `*time.Time` |  |
| Roles | `[]string` |  |
| User | `any` |  |

### GuildNSFWContentLevel
GuildNSFWContentLevel

#### Example Usage

```go
// Example usage of GuildNSFWContentLevel
var value GuildNSFWContentLevel
// Initialize with appropriate value
```

#### Type Definition

```go
type GuildNSFWContentLevel int
```

### GuildOnboarding
GuildOnboarding

#### Example Usage

```go
// Create a new GuildOnboarding
guildonboarding := GuildOnboarding{
    DefaultChannelIds: [],
    Enabled: true,
    GuildID: "example",
    Prompts: [],
}
```

#### Type Definition

```go
type GuildOnboarding struct {
    DefaultChannelIds []string `json:"default_channel_ids"`
    Enabled bool `json:"enabled"`
    GuildID string `json:"guild_id"`
    Prompts []any `json:"prompts"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| DefaultChannelIds | `[]string` |  |
| Enabled | `bool` |  |
| GuildID | `string` |  |
| Prompts | `[]any` |  |

### GuildOnboardingMode
GuildOnboardingMode

#### Example Usage

```go
// Example usage of GuildOnboardingMode
var value GuildOnboardingMode
// Initialize with appropriate value
```

#### Type Definition

```go
type GuildOnboardingMode int
```

### GuildPatchOptionsPartial
GuildPatchOptionsPartial

#### Example Usage

```go
// Create a new GuildPatchOptionsPartial
guildpatchoptionspartial := GuildPatchOptionsPartial{
    AfkChannelID: any{},
    AfkTimeout: any{},
    Banner: &"example"{},
    DefaultMessageNotifications: any{},
    Description: &"example"{},
    DiscoverySplash: &"example"{},
    ExplicitContentFilter: any{},
    Features: [],
    HomeHeader: &"example"{},
    Icon: &"example"{},
    Name: "example",
    PreferredLocale: any{},
    PremiumProgressBarEnabled: &true{},
    PublicUpdatesChannelID: any{},
    Region: &"example"{},
    RulesChannelID: any{},
    SafetyAlertsChannelID: any{},
    Splash: &"example"{},
    SystemChannelFlags: &42{},
    SystemChannelID: any{},
    VerificationLevel: any{},
}
```

#### Type Definition

```go
type GuildPatchOptionsPartial struct {
    AfkChannelID any `json:"afk_channel_id,omitempty"`
    AfkTimeout any `json:"afk_timeout,omitempty"`
    Banner *string `json:"banner,omitempty"`
    DefaultMessageNotifications any `json:"default_message_notifications,omitempty"`
    Description *string `json:"description,omitempty"`
    DiscoverySplash *string `json:"discovery_splash,omitempty"`
    ExplicitContentFilter any `json:"explicit_content_filter,omitempty"`
    Features []*string `json:"features,omitempty"`
    HomeHeader *string `json:"home_header,omitempty"`
    Icon *string `json:"icon,omitempty"`
    Name string `json:"name,omitempty"`
    PreferredLocale any `json:"preferred_locale,omitempty"`
    PremiumProgressBarEnabled *bool `json:"premium_progress_bar_enabled,omitempty"`
    PublicUpdatesChannelID any `json:"public_updates_channel_id,omitempty"`
    Region *string `json:"region,omitempty"`
    RulesChannelID any `json:"rules_channel_id,omitempty"`
    SafetyAlertsChannelID any `json:"safety_alerts_channel_id,omitempty"`
    Splash *string `json:"splash,omitempty"`
    SystemChannelFlags *int64 `json:"system_channel_flags,omitempty"`
    SystemChannelID any `json:"system_channel_id,omitempty"`
    VerificationLevel any `json:"verification_level,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AfkChannelID | `any` |  |
| AfkTimeout | `any` |  |
| Banner | `*string` |  |
| DefaultMessageNotifications | `any` |  |
| Description | `*string` |  |
| DiscoverySplash | `*string` |  |
| ExplicitContentFilter | `any` |  |
| Features | `[]*string` |  |
| HomeHeader | `*string` |  |
| Icon | `*string` |  |
| Name | `string` |  |
| PreferredLocale | `any` |  |
| PremiumProgressBarEnabled | `*bool` |  |
| PublicUpdatesChannelID | `any` |  |
| Region | `*string` |  |
| RulesChannelID | `any` |  |
| SafetyAlertsChannelID | `any` |  |
| Splash | `*string` |  |
| SystemChannelFlags | `*int64` |  |
| SystemChannelID | `any` |  |
| VerificationLevel | `any` |  |

### GuildPreview
GuildPreview

#### Example Usage

```go
// Create a new GuildPreview
guildpreview := GuildPreview{
    ID: "example",
    ApproximateMemberCount: 42,
    ApproximatePresenceCount: 42,
    Description: &"example"{},
    DiscoverySplash: &"example"{},
    Emojis: [],
    Features: [],
    HomeHeader: &"example"{},
    Icon: &"example"{},
    Name: "example",
    Splash: &"example"{},
    Stickers: [],
}
```

#### Type Definition

```go
type GuildPreview struct {
    ID string `json:"id"`
    ApproximateMemberCount int32 `json:"approximate_member_count"`
    ApproximatePresenceCount int32 `json:"approximate_presence_count"`
    Description *string `json:"description,omitempty"`
    DiscoverySplash *string `json:"discovery_splash,omitempty"`
    Emojis []any `json:"emojis"`
    Features []string `json:"features"`
    HomeHeader *string `json:"home_header,omitempty"`
    Icon *string `json:"icon,omitempty"`
    Name string `json:"name"`
    Splash *string `json:"splash,omitempty"`
    Stickers []any `json:"stickers"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ApproximateMemberCount | `int32` |  |
| ApproximatePresenceCount | `int32` |  |
| Description | `*string` |  |
| DiscoverySplash | `*string` |  |
| Emojis | `[]any` |  |
| Features | `[]string` |  |
| HomeHeader | `*string` |  |
| Icon | `*string` |  |
| Name | `string` |  |
| Splash | `*string` |  |
| Stickers | `[]any` |  |

### GuildProductPurchase
GuildProductPurchase

#### Example Usage

```go
// Create a new GuildProductPurchase
guildproductpurchase := GuildProductPurchase{
    ListingID: "example",
    ProductName: "example",
}
```

#### Type Definition

```go
type GuildProductPurchase struct {
    ListingID string `json:"listing_id"`
    ProductName string `json:"product_name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ListingID | `string` |  |
| ProductName | `string` |  |

### GuildPrune
GuildPrune

#### Example Usage

```go
// Create a new GuildPrune
guildprune := GuildPrune{
    Pruned: &42{},
}
```

#### Type Definition

```go
type GuildPrune struct {
    Pruned *int32 `json:"pruned,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Pruned | `*int32` |  |

### GuildRole
GuildRole

#### Example Usage

```go
// Create a new GuildRole
guildrole := GuildRole{
    ID: "example",
    Color: 42,
    Colors: any{},
    Description: &"example"{},
    Flags: 42,
    Hoist: true,
    Icon: &"example"{},
    Managed: true,
    Mentionable: true,
    Name: "example",
    Permissions: "example",
    Position: 42,
    Tags: any{},
    UnicodeEmoji: &"example"{},
}
```

#### Type Definition

```go
type GuildRole struct {
    ID string `json:"id"`
    Color int32 `json:"color"`
    Colors any `json:"colors"`
    Description *string `json:"description,omitempty"`
    Flags int32 `json:"flags"`
    Hoist bool `json:"hoist"`
    Icon *string `json:"icon,omitempty"`
    Managed bool `json:"managed"`
    Mentionable bool `json:"mentionable"`
    Name string `json:"name"`
    Permissions string `json:"permissions"`
    Position int32 `json:"position"`
    Tags any `json:"tags,omitempty"`
    UnicodeEmoji *string `json:"unicode_emoji,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Color | `int32` |  |
| Colors | `any` |  |
| Description | `*string` |  |
| Flags | `int32` |  |
| Hoist | `bool` |  |
| Icon | `*string` |  |
| Managed | `bool` |  |
| Mentionable | `bool` |  |
| Name | `string` |  |
| Permissions | `string` |  |
| Position | `int32` |  |
| Tags | `any` |  |
| UnicodeEmoji | `*string` |  |

### GuildRoleColors
GuildRoleColors

#### Example Usage

```go
// Create a new GuildRoleColors
guildrolecolors := GuildRoleColors{
    PrimaryColor: 42,
    SecondaryColor: &42{},
    TertiaryColor: &42{},
}
```

#### Type Definition

```go
type GuildRoleColors struct {
    PrimaryColor int32 `json:"primary_color"`
    SecondaryColor *int32 `json:"secondary_color,omitempty"`
    TertiaryColor *int32 `json:"tertiary_color,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| PrimaryColor | `int32` |  |
| SecondaryColor | `*int32` |  |
| TertiaryColor | `*int32` |  |

### GuildRoleTags
GuildRoleTags

#### Example Usage

```go
// Create a new GuildRoleTags
guildroletags := GuildRoleTags{
    AvailableForPurchase: any{},
    BotID: "example",
    GuildConnections: any{},
    IntegrationID: "example",
    PremiumSubscriber: any{},
    SubscriptionListingID: "example",
}
```

#### Type Definition

```go
type GuildRoleTags struct {
    AvailableForPurchase any `json:"available_for_purchase,omitempty"`
    BotID string `json:"bot_id,omitempty"`
    GuildConnections any `json:"guild_connections,omitempty"`
    IntegrationID string `json:"integration_id,omitempty"`
    PremiumSubscriber any `json:"premium_subscriber,omitempty"`
    SubscriptionListingID string `json:"subscription_listing_id,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AvailableForPurchase | `any` |  |
| BotID | `string` |  |
| GuildConnections | `any` |  |
| IntegrationID | `string` |  |
| PremiumSubscriber | `any` |  |
| SubscriptionListingID | `string` |  |

### GuildScheduledEventEntityTypes
GuildScheduledEventEntityTypes

#### Example Usage

```go
// Example usage of GuildScheduledEventEntityTypes
var value GuildScheduledEventEntityTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type GuildScheduledEventEntityTypes int
```

### GuildScheduledEventPrivacyLevels
GuildScheduledEventPrivacyLevels

#### Example Usage

```go
// Example usage of GuildScheduledEventPrivacyLevels
var value GuildScheduledEventPrivacyLevels
// Initialize with appropriate value
```

#### Type Definition

```go
type GuildScheduledEventPrivacyLevels int
```

### GuildScheduledEventStatuses
GuildScheduledEventStatuses

#### Example Usage

```go
// Example usage of GuildScheduledEventStatuses
var value GuildScheduledEventStatuses
// Initialize with appropriate value
```

#### Type Definition

```go
type GuildScheduledEventStatuses int
```

### GuildSticker
GuildSticker

#### Example Usage

```go
// Create a new GuildSticker
guildsticker := GuildSticker{
    ID: "example",
    Available: true,
    Description: &"example"{},
    FormatType: any{},
    GuildID: "example",
    Name: "example",
    Tags: "example",
    Type: 42,
    User: any{},
}
```

#### Type Definition

```go
type GuildSticker struct {
    ID string `json:"id"`
    Available bool `json:"available"`
    Description *string `json:"description,omitempty"`
    FormatType any `json:"format_type,omitempty"`
    GuildID string `json:"guild_id"`
    Name string `json:"name"`
    Tags string `json:"tags"`
    Type int32 `json:"type"`
    User any `json:"user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Available | `bool` |  |
| Description | `*string` |  |
| FormatType | `any` |  |
| GuildID | `string` |  |
| Name | `string` |  |
| Tags | `string` |  |
| Type | `int32` |  |
| User | `any` |  |

### GuildSubscriptionIntegration
GuildSubscriptionIntegration

#### Example Usage

```go
// Create a new GuildSubscriptionIntegration
guildsubscriptionintegration := GuildSubscriptionIntegration{
    ID: "example",
    Account: any{},
    Enabled: true,
    Name: &"example"{},
    Type: "example",
}
```

#### Type Definition

```go
type GuildSubscriptionIntegration struct {
    ID string `json:"id"`
    Account any `json:"account"`
    Enabled bool `json:"enabled"`
    Name *string `json:"name,omitempty"`
    Type string `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Account | `any` |  |
| Enabled | `bool` |  |
| Name | `*string` |  |
| Type | `string` |  |

### GuildTemplate
GuildTemplate

#### Example Usage

```go
// Create a new GuildTemplate
guildtemplate := GuildTemplate{
    Code: "example",
    CreatedAt: /* value */,
    Creator: any{},
    CreatorID: "example",
    Description: &"example"{},
    IsDirty: &true{},
    Name: "example",
    SerializedSourceGuild: any{},
    SourceGuildID: "example",
    UpdatedAt: /* value */,
    UsageCount: 42,
}
```

#### Type Definition

```go
type GuildTemplate struct {
    Code string `json:"code"`
    CreatedAt time.Time `json:"created_at"`
    Creator any `json:"creator,omitempty"`
    CreatorID string `json:"creator_id"`
    Description *string `json:"description,omitempty"`
    IsDirty *bool `json:"is_dirty,omitempty"`
    Name string `json:"name"`
    SerializedSourceGuild any `json:"serialized_source_guild"`
    SourceGuildID string `json:"source_guild_id"`
    UpdatedAt time.Time `json:"updated_at"`
    UsageCount int32 `json:"usage_count"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Code | `string` |  |
| CreatedAt | `time.Time` |  |
| Creator | `any` |  |
| CreatorID | `string` |  |
| Description | `*string` |  |
| IsDirty | `*bool` |  |
| Name | `string` |  |
| SerializedSourceGuild | `any` |  |
| SourceGuildID | `string` |  |
| UpdatedAt | `time.Time` |  |
| UsageCount | `int32` |  |

### GuildTemplateChannel
GuildTemplateChannel

#### Example Usage

```go
// Create a new GuildTemplateChannel
guildtemplatechannel := GuildTemplateChannel{
    ID: &42{},
    AvailableTags: [],
    Bitrate: 42,
    DefaultAutoArchiveDuration: any{},
    DefaultForumLayout: any{},
    DefaultReactionEmoji: any{},
    DefaultSortOrder: any{},
    DefaultTagSetting: any{},
    DefaultThreadRateLimitPerUser: &42{},
    IconEmoji: any{},
    Name: &"example"{},
    Nsfw: true,
    ParentID: any{},
    PermissionOverwrites: [],
    Position: &42{},
    RateLimitPerUser: 42,
    Template: "example",
    ThemeColor: &42{},
    Topic: &"example"{},
    Type: 42,
    UserLimit: 42,
}
```

#### Type Definition

```go
type GuildTemplateChannel struct {
    ID *int32 `json:"id,omitempty"`
    AvailableTags []any `json:"available_tags,omitempty"`
    Bitrate int32 `json:"bitrate"`
    DefaultAutoArchiveDuration any `json:"default_auto_archive_duration,omitempty"`
    DefaultForumLayout any `json:"default_forum_layout,omitempty"`
    DefaultReactionEmoji any `json:"default_reaction_emoji,omitempty"`
    DefaultSortOrder any `json:"default_sort_order,omitempty"`
    DefaultTagSetting any `json:"default_tag_setting,omitempty"`
    DefaultThreadRateLimitPerUser *int32 `json:"default_thread_rate_limit_per_user,omitempty"`
    IconEmoji any `json:"icon_emoji,omitempty"`
    Name *string `json:"name,omitempty"`
    Nsfw bool `json:"nsfw"`
    ParentID any `json:"parent_id,omitempty"`
    PermissionOverwrites []any `json:"permission_overwrites"`
    Position *int32 `json:"position,omitempty"`
    RateLimitPerUser int32 `json:"rate_limit_per_user"`
    Template string `json:"template"`
    ThemeColor *int32 `json:"theme_color,omitempty"`
    Topic *string `json:"topic,omitempty"`
    Type int32 `json:"type"`
    UserLimit int32 `json:"user_limit"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| AvailableTags | `[]any` |  |
| Bitrate | `int32` |  |
| DefaultAutoArchiveDuration | `any` |  |
| DefaultForumLayout | `any` |  |
| DefaultReactionEmoji | `any` |  |
| DefaultSortOrder | `any` |  |
| DefaultTagSetting | `any` |  |
| DefaultThreadRateLimitPerUser | `*int32` |  |
| IconEmoji | `any` |  |
| Name | `*string` |  |
| Nsfw | `bool` |  |
| ParentID | `any` |  |
| PermissionOverwrites | `[]any` |  |
| Position | `*int32` |  |
| RateLimitPerUser | `int32` |  |
| Template | `string` |  |
| ThemeColor | `*int32` |  |
| Topic | `*string` |  |
| Type | `int32` |  |
| UserLimit | `int32` |  |

### GuildTemplateChannelTags
GuildTemplateChannelTags

#### Example Usage

```go
// Create a new GuildTemplateChannelTags
guildtemplatechanneltags := GuildTemplateChannelTags{
    ID: &42{},
    EmojiID: any{},
    EmojiName: &"example"{},
    Moderated: &true{},
    Name: "example",
}
```

#### Type Definition

```go
type GuildTemplateChannelTags struct {
    ID *int32 `json:"id,omitempty"`
    EmojiID any `json:"emoji_id,omitempty"`
    EmojiName *string `json:"emoji_name,omitempty"`
    Moderated *bool `json:"moderated,omitempty"`
    Name string `json:"name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| EmojiID | `any` |  |
| EmojiName | `*string` |  |
| Moderated | `*bool` |  |
| Name | `string` |  |

### GuildTemplateRole
GuildTemplateRole

#### Example Usage

```go
// Create a new GuildTemplateRole
guildtemplaterole := GuildTemplateRole{
    ID: 42,
    Color: 42,
    Colors: any{},
    Hoist: true,
    Icon: &"example"{},
    Mentionable: true,
    Name: "example",
    Permissions: "example",
    UnicodeEmoji: &"example"{},
}
```

#### Type Definition

```go
type GuildTemplateRole struct {
    ID int32 `json:"id"`
    Color int32 `json:"color"`
    Colors any `json:"colors,omitempty"`
    Hoist bool `json:"hoist"`
    Icon *string `json:"icon,omitempty"`
    Mentionable bool `json:"mentionable"`
    Name string `json:"name"`
    Permissions string `json:"permissions"`
    UnicodeEmoji *string `json:"unicode_emoji,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| Color | `int32` |  |
| Colors | `any` |  |
| Hoist | `bool` |  |
| Icon | `*string` |  |
| Mentionable | `bool` |  |
| Name | `string` |  |
| Permissions | `string` |  |
| UnicodeEmoji | `*string` |  |

### GuildTemplateRoleColors
GuildTemplateRoleColors

#### Example Usage

```go
// Create a new GuildTemplateRoleColors
guildtemplaterolecolors := GuildTemplateRoleColors{
    PrimaryColor: 42,
    SecondaryColor: &42{},
    TertiaryColor: &42{},
}
```

#### Type Definition

```go
type GuildTemplateRoleColors struct {
    PrimaryColor int32 `json:"primary_color"`
    SecondaryColor *int32 `json:"secondary_color,omitempty"`
    TertiaryColor *int32 `json:"tertiary_color,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| PrimaryColor | `int32` |  |
| SecondaryColor | `*int32` |  |
| TertiaryColor | `*int32` |  |

### GuildTemplateSnapshot
GuildTemplateSnapshot

#### Example Usage

```go
// Create a new GuildTemplateSnapshot
guildtemplatesnapshot := GuildTemplateSnapshot{
    AfkChannelID: any{},
    AfkTimeout: 42,
    Channels: [],
    DefaultMessageNotifications: 42,
    Description: &"example"{},
    ExplicitContentFilter: 42,
    Name: "example",
    PreferredLocale: "example",
    Region: &"example"{},
    Roles: [],
    SystemChannelFlags: 42,
    SystemChannelID: any{},
    VerificationLevel: 42,
}
```

#### Type Definition

```go
type GuildTemplateSnapshot struct {
    AfkChannelID any `json:"afk_channel_id,omitempty"`
    AfkTimeout int32 `json:"afk_timeout"`
    Channels []any `json:"channels"`
    DefaultMessageNotifications int32 `json:"default_message_notifications"`
    Description *string `json:"description,omitempty"`
    ExplicitContentFilter int32 `json:"explicit_content_filter"`
    Name string `json:"name"`
    PreferredLocale string `json:"preferred_locale"`
    Region *string `json:"region,omitempty"`
    Roles []any `json:"roles"`
    SystemChannelFlags int32 `json:"system_channel_flags"`
    SystemChannelID any `json:"system_channel_id,omitempty"`
    VerificationLevel int32 `json:"verification_level"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AfkChannelID | `any` |  |
| AfkTimeout | `int32` |  |
| Channels | `[]any` |  |
| DefaultMessageNotifications | `int32` |  |
| Description | `*string` |  |
| ExplicitContentFilter | `int32` |  |
| Name | `string` |  |
| PreferredLocale | `string` |  |
| Region | `*string` |  |
| Roles | `[]any` |  |
| SystemChannelFlags | `int32` |  |
| SystemChannelID | `any` |  |
| VerificationLevel | `int32` |  |

### GuildWelcomeChannel
GuildWelcomeChannel

#### Example Usage

```go
// Create a new GuildWelcomeChannel
guildwelcomechannel := GuildWelcomeChannel{
    ChannelID: "example",
    Description: "example",
    EmojiID: any{},
    EmojiName: &"example"{},
}
```

#### Type Definition

```go
type GuildWelcomeChannel struct {
    ChannelID string `json:"channel_id"`
    Description string `json:"description"`
    EmojiID any `json:"emoji_id,omitempty"`
    EmojiName *string `json:"emoji_name,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `string` |  |
| Description | `string` |  |
| EmojiID | `any` |  |
| EmojiName | `*string` |  |

### GuildWelcomeScreen
GuildWelcomeScreen

#### Example Usage

```go
// Create a new GuildWelcomeScreen
guildwelcomescreen := GuildWelcomeScreen{
    Description: &"example"{},
    WelcomeChannels: [],
}
```

#### Type Definition

```go
type GuildWelcomeScreen struct {
    Description *string `json:"description,omitempty"`
    WelcomeChannels []any `json:"welcome_channels"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `*string` |  |
| WelcomeChannels | `[]any` |  |

### GuildWelcomeScreenChannel
GuildWelcomeScreenChannel

#### Example Usage

```go
// Create a new GuildWelcomeScreenChannel
guildwelcomescreenchannel := GuildWelcomeScreenChannel{
    ChannelID: "example",
    Description: "example",
    EmojiID: any{},
    EmojiName: &"example"{},
}
```

#### Type Definition

```go
type GuildWelcomeScreenChannel struct {
    ChannelID string `json:"channel_id"`
    Description string `json:"description"`
    EmojiID any `json:"emoji_id,omitempty"`
    EmojiName *string `json:"emoji_name,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `string` |  |
| Description | `string` |  |
| EmojiID | `any` |  |
| EmojiName | `*string` |  |

### GuildWithCounts
GuildWithCounts

#### Example Usage

```go
// Create a new GuildWithCounts
guildwithcounts := GuildWithCounts{
    ID: "example",
    AfkChannelID: any{},
    AfkTimeout: 42,
    ApplicationID: any{},
    ApproximateMemberCount: &42{},
    ApproximatePresenceCount: &42{},
    Banner: &"example"{},
    DefaultMessageNotifications: 42,
    Description: &"example"{},
    DiscoverySplash: &"example"{},
    Emojis: [],
    ExplicitContentFilter: 42,
    Features: [],
    HomeHeader: &"example"{},
    Icon: &"example"{},
    MaxMembers: 42,
    MaxPresences: &42{},
    MaxStageVideoChannelUsers: 42,
    MaxVideoChannelUsers: 42,
    MfaLevel: 42,
    Name: "example",
    Nsfw: true,
    NsfwLevel: 42,
    OwnerID: "example",
    PreferredLocale: "example",
    PremiumProgressBarEnabled: true,
    PremiumSubscriptionCount: 42,
    PremiumTier: 42,
    PublicUpdatesChannelID: any{},
    Region: "example",
    Roles: [],
    RulesChannelID: any{},
    SafetyAlertsChannelID: any{},
    Splash: &"example"{},
    Stickers: [],
    SystemChannelFlags: 42,
    SystemChannelID: any{},
    VanityURLCode: &"example"{},
    VerificationLevel: 42,
    WidgetChannelID: any{},
    WidgetEnabled: true,
}
```

#### Type Definition

```go
type GuildWithCounts struct {
    ID string `json:"id"`
    AfkChannelID any `json:"afk_channel_id,omitempty"`
    AfkTimeout int32 `json:"afk_timeout"`
    ApplicationID any `json:"application_id,omitempty"`
    ApproximateMemberCount *int32 `json:"approximate_member_count,omitempty"`
    ApproximatePresenceCount *int32 `json:"approximate_presence_count,omitempty"`
    Banner *string `json:"banner,omitempty"`
    DefaultMessageNotifications int32 `json:"default_message_notifications"`
    Description *string `json:"description,omitempty"`
    DiscoverySplash *string `json:"discovery_splash,omitempty"`
    Emojis []any `json:"emojis"`
    ExplicitContentFilter int32 `json:"explicit_content_filter"`
    Features []string `json:"features"`
    HomeHeader *string `json:"home_header,omitempty"`
    Icon *string `json:"icon,omitempty"`
    MaxMembers int32 `json:"max_members"`
    MaxPresences *int32 `json:"max_presences,omitempty"`
    MaxStageVideoChannelUsers int32 `json:"max_stage_video_channel_users"`
    MaxVideoChannelUsers int32 `json:"max_video_channel_users"`
    MfaLevel int32 `json:"mfa_level"`
    Name string `json:"name"`
    Nsfw bool `json:"nsfw"`
    NsfwLevel int32 `json:"nsfw_level"`
    OwnerID string `json:"owner_id"`
    PreferredLocale string `json:"preferred_locale"`
    PremiumProgressBarEnabled bool `json:"premium_progress_bar_enabled"`
    PremiumSubscriptionCount int32 `json:"premium_subscription_count"`
    PremiumTier int32 `json:"premium_tier"`
    PublicUpdatesChannelID any `json:"public_updates_channel_id,omitempty"`
    Region string `json:"region"`
    Roles []any `json:"roles"`
    RulesChannelID any `json:"rules_channel_id,omitempty"`
    SafetyAlertsChannelID any `json:"safety_alerts_channel_id,omitempty"`
    Splash *string `json:"splash,omitempty"`
    Stickers []any `json:"stickers"`
    SystemChannelFlags int32 `json:"system_channel_flags"`
    SystemChannelID any `json:"system_channel_id,omitempty"`
    VanityURLCode *string `json:"vanity_url_code,omitempty"`
    VerificationLevel int32 `json:"verification_level"`
    WidgetChannelID any `json:"widget_channel_id,omitempty"`
    WidgetEnabled bool `json:"widget_enabled"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| AfkChannelID | `any` |  |
| AfkTimeout | `int32` |  |
| ApplicationID | `any` |  |
| ApproximateMemberCount | `*int32` |  |
| ApproximatePresenceCount | `*int32` |  |
| Banner | `*string` |  |
| DefaultMessageNotifications | `int32` |  |
| Description | `*string` |  |
| DiscoverySplash | `*string` |  |
| Emojis | `[]any` |  |
| ExplicitContentFilter | `int32` |  |
| Features | `[]string` |  |
| HomeHeader | `*string` |  |
| Icon | `*string` |  |
| MaxMembers | `int32` |  |
| MaxPresences | `*int32` |  |
| MaxStageVideoChannelUsers | `int32` |  |
| MaxVideoChannelUsers | `int32` |  |
| MfaLevel | `int32` |  |
| Name | `string` |  |
| Nsfw | `bool` |  |
| NsfwLevel | `int32` |  |
| OwnerID | `string` |  |
| PreferredLocale | `string` |  |
| PremiumProgressBarEnabled | `bool` |  |
| PremiumSubscriptionCount | `int32` |  |
| PremiumTier | `int32` |  |
| PublicUpdatesChannelID | `any` |  |
| Region | `string` |  |
| Roles | `[]any` |  |
| RulesChannelID | `any` |  |
| SafetyAlertsChannelID | `any` |  |
| Splash | `*string` |  |
| Stickers | `[]any` |  |
| SystemChannelFlags | `int32` |  |
| SystemChannelID | `any` |  |
| VanityURLCode | `*string` |  |
| VerificationLevel | `int32` |  |
| WidgetChannelID | `any` |  |
| WidgetEnabled | `bool` |  |

### IconEmoji
IconEmoji

#### Example Usage

```go
// Create a new IconEmoji
iconemoji := IconEmoji{

}
```

#### Type Definition

```go
type IconEmoji struct {
}
```

### IncomingWebhookInteractionOptions
IncomingWebhookInteractionOptions

#### Example Usage

```go
// Create a new IncomingWebhookInteractionOptions
incomingwebhookinteractionoptions := IncomingWebhookInteractionOptions{
    AllowedMentions: any{},
    Attachments: [],
    Components: [],
    Content: &"example"{},
    Embeds: [],
    Flags: &42{},
    Poll: any{},
    Tts: &true{},
}
```

#### Type Definition

```go
type IncomingWebhookInteractionOptions struct {
    AllowedMentions any `json:"allowed_mentions,omitempty"`
    Attachments []any `json:"attachments,omitempty"`
    Components []any `json:"components,omitempty"`
    Content *string `json:"content,omitempty"`
    Embeds []any `json:"embeds,omitempty"`
    Flags *int64 `json:"flags,omitempty"`
    Poll any `json:"poll,omitempty"`
    Tts *bool `json:"tts,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AllowedMentions | `any` |  |
| Attachments | `[]any` |  |
| Components | `[]any` |  |
| Content | `*string` |  |
| Embeds | `[]any` |  |
| Flags | `*int64` |  |
| Poll | `any` |  |
| Tts | `*bool` |  |

### IncomingWebhookOptionsPartial
IncomingWebhookOptionsPartial

#### Example Usage

```go
// Create a new IncomingWebhookOptionsPartial
incomingwebhookoptionspartial := IncomingWebhookOptionsPartial{
    AllowedMentions: any{},
    AppliedTags: [],
    Attachments: [],
    AvatarURL: &"example"{},
    Components: [],
    Content: &"example"{},
    Embeds: [],
    Flags: &42{},
    Poll: any{},
    ThreadName: &"example"{},
    Tts: &true{},
    Username: &"example"{},
}
```

#### Type Definition

```go
type IncomingWebhookOptionsPartial struct {
    AllowedMentions any `json:"allowed_mentions,omitempty"`
    AppliedTags []string `json:"applied_tags,omitempty"`
    Attachments []any `json:"attachments,omitempty"`
    AvatarURL *string `json:"avatar_url,omitempty"`
    Components []any `json:"components,omitempty"`
    Content *string `json:"content,omitempty"`
    Embeds []any `json:"embeds,omitempty"`
    Flags *int64 `json:"flags,omitempty"`
    Poll any `json:"poll,omitempty"`
    ThreadName *string `json:"thread_name,omitempty"`
    Tts *bool `json:"tts,omitempty"`
    Username *string `json:"username,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AllowedMentions | `any` |  |
| AppliedTags | `[]string` |  |
| Attachments | `[]any` |  |
| AvatarURL | `*string` |  |
| Components | `[]any` |  |
| Content | `*string` |  |
| Embeds | `[]any` |  |
| Flags | `*int64` |  |
| Poll | `any` |  |
| ThreadName | `*string` |  |
| Tts | `*bool` |  |
| Username | `*string` |  |

### IncomingWebhookUpdateForInteractionCallbackOptionsPartial
IncomingWebhookUpdateForInteractionCallbackOptionsPartial

#### Example Usage

```go
// Create a new IncomingWebhookUpdateForInteractionCallbackOptionsPartial
incomingwebhookupdateforinteractioncallbackoptionspartial := IncomingWebhookUpdateForInteractionCallbackOptionsPartial{
    AllowedMentions: any{},
    Attachments: [],
    Components: [],
    Content: &"example"{},
    Embeds: [],
    Flags: &42{},
}
```

#### Type Definition

```go
type IncomingWebhookUpdateForInteractionCallbackOptionsPartial struct {
    AllowedMentions any `json:"allowed_mentions,omitempty"`
    Attachments []any `json:"attachments,omitempty"`
    Components []any `json:"components,omitempty"`
    Content *string `json:"content,omitempty"`
    Embeds []any `json:"embeds,omitempty"`
    Flags *int64 `json:"flags,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AllowedMentions | `any` |  |
| Attachments | `[]any` |  |
| Components | `[]any` |  |
| Content | `*string` |  |
| Embeds | `[]any` |  |
| Flags | `*int64` |  |

### IncomingWebhookUpdateOptionsPartial
IncomingWebhookUpdateOptionsPartial

#### Example Usage

```go
// Create a new IncomingWebhookUpdateOptionsPartial
incomingwebhookupdateoptionspartial := IncomingWebhookUpdateOptionsPartial{
    AllowedMentions: any{},
    Attachments: [],
    Components: [],
    Content: &"example"{},
    Embeds: [],
    Flags: &42{},
    Poll: any{},
}
```

#### Type Definition

```go
type IncomingWebhookUpdateOptionsPartial struct {
    AllowedMentions any `json:"allowed_mentions,omitempty"`
    Attachments []any `json:"attachments,omitempty"`
    Components []any `json:"components,omitempty"`
    Content *string `json:"content,omitempty"`
    Embeds []any `json:"embeds,omitempty"`
    Flags *int64 `json:"flags,omitempty"`
    Poll any `json:"poll,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AllowedMentions | `any` |  |
| Attachments | `[]any` |  |
| Components | `[]any` |  |
| Content | `*string` |  |
| Embeds | `[]any` |  |
| Flags | `*int64` |  |
| Poll | `any` |  |

### InnerErrors
InnerErrors

#### Example Usage

```go
// Create a new InnerErrors
innererrors := InnerErrors{
    Errors: [],
}
```

#### Type Definition

```go
type InnerErrors struct {
    Errors []any `json:"_errors"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Errors | `[]any` | The list of errors for this field |

### Int53Type
Int53Type

#### Example Usage

```go
// Create a new Int53Type
int53type := Int53Type{

}
```

#### Type Definition

```go
type Int53Type struct {
}
```

### IntegrationApplication
IntegrationApplication

#### Example Usage

```go
// Create a new IntegrationApplication
integrationapplication := IntegrationApplication{
    ID: "example",
    Bot: any{},
    CoverImage: "example",
    Description: "example",
    Icon: &"example"{},
    Name: "example",
    PrimarySkuID: "example",
    Type: any{},
}
```

#### Type Definition

```go
type IntegrationApplication struct {
    ID string `json:"id"`
    Bot any `json:"bot,omitempty"`
    CoverImage string `json:"cover_image,omitempty"`
    Description string `json:"description"`
    Icon *string `json:"icon,omitempty"`
    Name string `json:"name"`
    PrimarySkuID string `json:"primary_sku_id,omitempty"`
    Type any `json:"type,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Bot | `any` |  |
| CoverImage | `string` |  |
| Description | `string` |  |
| Icon | `*string` |  |
| Name | `string` |  |
| PrimarySkuID | `string` |  |
| Type | `any` |  |

### IntegrationExpireBehaviorTypes
IntegrationExpireBehaviorTypes

#### Example Usage

```go
// Example usage of IntegrationExpireBehaviorTypes
var value IntegrationExpireBehaviorTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type IntegrationExpireBehaviorTypes int
```

### IntegrationExpireGracePeriodTypes
IntegrationExpireGracePeriodTypes

#### Example Usage

```go
// Example usage of IntegrationExpireGracePeriodTypes
var value IntegrationExpireGracePeriodTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type IntegrationExpireGracePeriodTypes int
```

### IntegrationTypes
IntegrationTypes

#### Example Usage

```go
// Example usage of IntegrationTypes
var value IntegrationTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type IntegrationTypes string
```

### Interaction
Interaction

#### Example Usage

```go
// Create a new Interaction
interaction := Interaction{
    ID: "example",
    ChannelID: "example",
    GuildID: "example",
    ResponseMessageEphemeral: true,
    ResponseMessageID: "example",
    ResponseMessageLoading: true,
    Type: 42,
}
```

#### Type Definition

```go
type Interaction struct {
    ID string `json:"id"`
    ChannelID string `json:"channel_id,omitempty"`
    GuildID string `json:"guild_id,omitempty"`
    ResponseMessageEphemeral bool `json:"response_message_ephemeral,omitempty"`
    ResponseMessageID string `json:"response_message_id,omitempty"`
    ResponseMessageLoading bool `json:"response_message_loading,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ChannelID | `string` |  |
| GuildID | `string` |  |
| ResponseMessageEphemeral | `bool` |  |
| ResponseMessageID | `string` |  |
| ResponseMessageLoading | `bool` |  |
| Type | `int32` |  |

### InteractionApplicationCommandAutocompleteCallbackIntegerData
InteractionApplicationCommandAutocompleteCallbackIntegerData

#### Example Usage

```go
// Create a new InteractionApplicationCommandAutocompleteCallbackIntegerData
interactionapplicationcommandautocompletecallbackintegerdata := InteractionApplicationCommandAutocompleteCallbackIntegerData{
    Choices: [],
}
```

#### Type Definition

```go
type InteractionApplicationCommandAutocompleteCallbackIntegerData struct {
    Choices []any `json:"choices,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Choices | `[]any` |  |

### InteractionApplicationCommandAutocompleteCallbackNumberData
InteractionApplicationCommandAutocompleteCallbackNumberData

#### Example Usage

```go
// Create a new InteractionApplicationCommandAutocompleteCallbackNumberData
interactionapplicationcommandautocompletecallbacknumberdata := InteractionApplicationCommandAutocompleteCallbackNumberData{
    Choices: [],
}
```

#### Type Definition

```go
type InteractionApplicationCommandAutocompleteCallbackNumberData struct {
    Choices []any `json:"choices,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Choices | `[]any` |  |

### InteractionApplicationCommandAutocompleteCallbackStringData
InteractionApplicationCommandAutocompleteCallbackStringData

#### Example Usage

```go
// Create a new InteractionApplicationCommandAutocompleteCallbackStringData
interactionapplicationcommandautocompletecallbackstringdata := InteractionApplicationCommandAutocompleteCallbackStringData{
    Choices: [],
}
```

#### Type Definition

```go
type InteractionApplicationCommandAutocompleteCallbackStringData struct {
    Choices []any `json:"choices,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Choices | `[]any` |  |

### InteractionCallback
InteractionCallback

#### Example Usage

```go
// Create a new InteractionCallback
interactioncallback := InteractionCallback{
    Interaction: any{},
    Resource: any{},
}
```

#### Type Definition

```go
type InteractionCallback struct {
    Interaction any `json:"interaction"`
    Resource any `json:"resource,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Interaction | `any` |  |
| Resource | `any` |  |

### InteractionCallbackTypes
InteractionCallbackTypes

#### Example Usage

```go
// Example usage of InteractionCallbackTypes
var value InteractionCallbackTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type InteractionCallbackTypes int
```

### InteractionContextType
InteractionContextType

#### Example Usage

```go
// Example usage of InteractionContextType
var value InteractionContextType
// Initialize with appropriate value
```

#### Type Definition

```go
type InteractionContextType int
```

### InteractionTypes
InteractionTypes

#### Example Usage

```go
// Example usage of InteractionTypes
var value InteractionTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type InteractionTypes int
```

### InviteApplication
InviteApplication

#### Example Usage

```go
// Create a new InviteApplication
inviteapplication := InviteApplication{
    ID: "example",
    Bot: any{},
    BotPublic: true,
    BotRequireCodeGrant: true,
    CoverImage: "example",
    CustomInstallURL: "example",
    Description: "example",
    Flags: 42,
    GuildID: "example",
    Icon: &"example"{},
    InstallParams: any{},
    IntegrationTypesConfig: map[],
    MaxParticipants: &42{},
    Name: "example",
    PrimarySkuID: "example",
    PrivacyPolicyURL: "example",
    RpcOrigins: [],
    Slug: "example",
    Tags: [],
    TermsOfServiceURL: "example",
    Type: any{},
    VerifyKey: "example",
}
```

#### Type Definition

```go
type InviteApplication struct {
    ID string `json:"id"`
    Bot any `json:"bot,omitempty"`
    BotPublic bool `json:"bot_public,omitempty"`
    BotRequireCodeGrant bool `json:"bot_require_code_grant,omitempty"`
    CoverImage string `json:"cover_image,omitempty"`
    CustomInstallURL string `json:"custom_install_url,omitempty"`
    Description string `json:"description"`
    Flags int32 `json:"flags"`
    GuildID string `json:"guild_id,omitempty"`
    Icon *string `json:"icon,omitempty"`
    InstallParams any `json:"install_params,omitempty"`
    IntegrationTypesConfig map[string]any `json:"integration_types_config,omitempty"`
    MaxParticipants *int32 `json:"max_participants,omitempty"`
    Name string `json:"name"`
    PrimarySkuID string `json:"primary_sku_id,omitempty"`
    PrivacyPolicyURL string `json:"privacy_policy_url,omitempty"`
    RpcOrigins []*string `json:"rpc_origins,omitempty"`
    Slug string `json:"slug,omitempty"`
    Tags []string `json:"tags,omitempty"`
    TermsOfServiceURL string `json:"terms_of_service_url,omitempty"`
    Type any `json:"type,omitempty"`
    VerifyKey string `json:"verify_key"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Bot | `any` |  |
| BotPublic | `bool` |  |
| BotRequireCodeGrant | `bool` |  |
| CoverImage | `string` |  |
| CustomInstallURL | `string` |  |
| Description | `string` |  |
| Flags | `int32` |  |
| GuildID | `string` |  |
| Icon | `*string` |  |
| InstallParams | `any` |  |
| IntegrationTypesConfig | `map[string]any` |  |
| MaxParticipants | `*int32` |  |
| Name | `string` |  |
| PrimarySkuID | `string` |  |
| PrivacyPolicyURL | `string` |  |
| RpcOrigins | `[]*string` |  |
| Slug | `string` |  |
| Tags | `[]string` |  |
| TermsOfServiceURL | `string` |  |
| Type | `any` |  |
| VerifyKey | `string` |  |

### InviteChannel
InviteChannel

#### Example Usage

```go
// Create a new InviteChannel
invitechannel := InviteChannel{
    ID: "example",
    Icon: "example",
    Name: &"example"{},
    Recipients: [],
    Type: 42,
}
```

#### Type Definition

```go
type InviteChannel struct {
    ID string `json:"id"`
    Icon string `json:"icon,omitempty"`
    Name *string `json:"name,omitempty"`
    Recipients []any `json:"recipients,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Icon | `string` |  |
| Name | `*string` |  |
| Recipients | `[]any` |  |
| Type | `int32` |  |

### InviteChannelRecipient
InviteChannelRecipient

#### Example Usage

```go
// Create a new InviteChannelRecipient
invitechannelrecipient := InviteChannelRecipient{
    Username: "example",
}
```

#### Type Definition

```go
type InviteChannelRecipient struct {
    Username string `json:"username"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Username | `string` |  |

### InviteGuild
InviteGuild

#### Example Usage

```go
// Create a new InviteGuild
inviteguild := InviteGuild{
    ID: "example",
    Banner: &"example"{},
    Description: &"example"{},
    Features: [],
    Icon: &"example"{},
    Name: "example",
    Nsfw: &true{},
    NsfwLevel: any{},
    PremiumSubscriptionCount: 42,
    Splash: &"example"{},
    VanityURLCode: &"example"{},
    VerificationLevel: any{},
}
```

#### Type Definition

```go
type InviteGuild struct {
    ID string `json:"id"`
    Banner *string `json:"banner,omitempty"`
    Description *string `json:"description,omitempty"`
    Features []string `json:"features"`
    Icon *string `json:"icon,omitempty"`
    Name string `json:"name"`
    Nsfw *bool `json:"nsfw,omitempty"`
    NsfwLevel any `json:"nsfw_level,omitempty"`
    PremiumSubscriptionCount int32 `json:"premium_subscription_count"`
    Splash *string `json:"splash,omitempty"`
    VanityURLCode *string `json:"vanity_url_code,omitempty"`
    VerificationLevel any `json:"verification_level,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Banner | `*string` |  |
| Description | `*string` |  |
| Features | `[]string` |  |
| Icon | `*string` |  |
| Name | `string` |  |
| Nsfw | `*bool` |  |
| NsfwLevel | `any` |  |
| PremiumSubscriptionCount | `int32` |  |
| Splash | `*string` |  |
| VanityURLCode | `*string` |  |
| VerificationLevel | `any` |  |

### InviteTargetTypes
InviteTargetTypes

#### Example Usage

```go
// Example usage of InviteTargetTypes
var value InviteTargetTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type InviteTargetTypes int
```

### InviteTypes
InviteTypes

#### Example Usage

```go
// Example usage of InviteTypes
var value InviteTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type InviteTypes int
```

### KeywordRule
KeywordRule

#### Example Usage

```go
// Create a new KeywordRule
keywordrule := KeywordRule{
    ID: "example",
    Actions: [],
    CreatorID: "example",
    Enabled: true,
    EventType: 42,
    ExemptChannels: [],
    ExemptRoles: [],
    GuildID: "example",
    Name: "example",
    TriggerMetadata: any{},
    TriggerType: 42,
}
```

#### Type Definition

```go
type KeywordRule struct {
    ID string `json:"id"`
    Actions []any `json:"actions"`
    CreatorID string `json:"creator_id"`
    Enabled bool `json:"enabled"`
    EventType int32 `json:"event_type"`
    ExemptChannels []string `json:"exempt_channels"`
    ExemptRoles []string `json:"exempt_roles"`
    GuildID string `json:"guild_id"`
    Name string `json:"name"`
    TriggerMetadata any `json:"trigger_metadata"`
    TriggerType int32 `json:"trigger_type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Actions | `[]any` |  |
| CreatorID | `string` |  |
| Enabled | `bool` |  |
| EventType | `int32` |  |
| ExemptChannels | `[]string` |  |
| ExemptRoles | `[]string` |  |
| GuildID | `string` |  |
| Name | `string` |  |
| TriggerMetadata | `any` |  |
| TriggerType | `int32` |  |

### KeywordTriggerMetadata
KeywordTriggerMetadata

#### Example Usage

```go
// Create a new KeywordTriggerMetadata
keywordtriggermetadata := KeywordTriggerMetadata{
    AllowList: [],
    KeywordFilter: [],
    RegexPatterns: [],
}
```

#### Type Definition

```go
type KeywordTriggerMetadata struct {
    AllowList []string `json:"allow_list"`
    KeywordFilter []string `json:"keyword_filter"`
    RegexPatterns []string `json:"regex_patterns"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AllowList | `[]string` |  |
| KeywordFilter | `[]string` |  |
| RegexPatterns | `[]string` |  |

### KeywordUpsertOptions
KeywordUpsertOptions

#### Example Usage

```go
// Create a new KeywordUpsertOptions
keywordupsertoptions := KeywordUpsertOptions{
    Actions: [],
    Enabled: &true{},
    EventType: 42,
    ExemptChannels: [],
    ExemptRoles: [],
    Name: "example",
    TriggerMetadata: any{},
    TriggerType: 42,
}
```

#### Type Definition

```go
type KeywordUpsertOptions struct {
    Actions []any `json:"actions,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    EventType int32 `json:"event_type"`
    ExemptChannels []string `json:"exempt_channels,omitempty"`
    ExemptRoles []string `json:"exempt_roles,omitempty"`
    Name string `json:"name"`
    TriggerMetadata any `json:"trigger_metadata,omitempty"`
    TriggerType int32 `json:"trigger_type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Actions | `[]any` |  |
| Enabled | `*bool` |  |
| EventType | `int32` |  |
| ExemptChannels | `[]string` |  |
| ExemptRoles | `[]string` |  |
| Name | `string` |  |
| TriggerMetadata | `any` |  |
| TriggerType | `int32` |  |

### KeywordUpsertOptionsPartial
KeywordUpsertOptionsPartial

#### Example Usage

```go
// Create a new KeywordUpsertOptionsPartial
keywordupsertoptionspartial := KeywordUpsertOptionsPartial{
    Actions: [],
    Enabled: &true{},
    EventType: 42,
    ExemptChannels: [],
    ExemptRoles: [],
    Name: "example",
    TriggerMetadata: any{},
    TriggerType: 42,
}
```

#### Type Definition

```go
type KeywordUpsertOptionsPartial struct {
    Actions []any `json:"actions,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    EventType int32 `json:"event_type,omitempty"`
    ExemptChannels []string `json:"exempt_channels,omitempty"`
    ExemptRoles []string `json:"exempt_roles,omitempty"`
    Name string `json:"name,omitempty"`
    TriggerMetadata any `json:"trigger_metadata,omitempty"`
    TriggerType int32 `json:"trigger_type,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Actions | `[]any` |  |
| Enabled | `*bool` |  |
| EventType | `int32` |  |
| ExemptChannels | `[]string` |  |
| ExemptRoles | `[]string` |  |
| Name | `string` |  |
| TriggerMetadata | `any` |  |
| TriggerType | `int32` |  |

### LabelComponentForModalOptions
LabelComponentForModalOptions

#### Example Usage

```go
// Create a new LabelComponentForModalOptions
labelcomponentformodaloptions := LabelComponentForModalOptions{
    ID: &42{},
    Component: any{},
    Description: &"example"{},
    Label: "example",
    Type: 42,
}
```

#### Type Definition

```go
type LabelComponentForModalOptions struct {
    ID *int32 `json:"id,omitempty"`
    Component any `json:"component"`
    Description *string `json:"description,omitempty"`
    Label string `json:"label"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| Component | `any` |  |
| Description | `*string` |  |
| Label | `string` |  |
| Type | `int32` |  |

### LaunchActivityInteractionCallback
LaunchActivityInteractionCallback

#### Example Usage

```go
// Create a new LaunchActivityInteractionCallback
launchactivityinteractioncallback := LaunchActivityInteractionCallback{
    Type: 42,
}
```

#### Type Definition

```go
type LaunchActivityInteractionCallback struct {
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Type | `int32` |  |

### LaunchActivityInteractionCallbackOptions
LaunchActivityInteractionCallbackOptions

#### Example Usage

```go
// Create a new LaunchActivityInteractionCallbackOptions
launchactivityinteractioncallbackoptions := LaunchActivityInteractionCallbackOptions{
    Type: 42,
}
```

#### Type Definition

```go
type LaunchActivityInteractionCallbackOptions struct {
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Type | `int32` |  |

### ListApplicationEmojis
ListApplicationEmojis

#### Example Usage

```go
// Create a new ListApplicationEmojis
listapplicationemojis := ListApplicationEmojis{
    Items: [],
}
```

#### Type Definition

```go
type ListApplicationEmojis struct {
    Items []any `json:"items"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Items | `[]any` |  |

### ListGuildSoundboardSounds
ListGuildSoundboardSounds

#### Example Usage

```go
// Create a new ListGuildSoundboardSounds
listguildsoundboardsounds := ListGuildSoundboardSounds{
    Items: [],
}
```

#### Type Definition

```go
type ListGuildSoundboardSounds struct {
    Items []any `json:"items"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Items | `[]any` |  |

### Lobby
Lobby

#### Example Usage

```go
// Create a new Lobby
lobby := Lobby{
    ID: "example",
    ApplicationID: "example",
    Flags: 42,
    LinkedChannel: any{},
    Members: [],
    Metadata: map[],
    OverrideEventWebhooksURL: &"example"{},
}
```

#### Type Definition

```go
type Lobby struct {
    ID string `json:"id"`
    ApplicationID string `json:"application_id"`
    Flags int64 `json:"flags"`
    LinkedChannel any `json:"linked_channel,omitempty"`
    Members []any `json:"members"`
    Metadata map[string]string `json:"metadata,omitempty"`
    OverrideEventWebhooksURL *string `json:"override_event_webhooks_url,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ApplicationID | `string` |  |
| Flags | `int64` |  |
| LinkedChannel | `any` |  |
| Members | `[]any` |  |
| Metadata | `map[string]string` |  |
| OverrideEventWebhooksURL | `*string` |  |

### LobbyGuildInvite
LobbyGuildInvite

#### Example Usage

```go
// Create a new LobbyGuildInvite
lobbyguildinvite := LobbyGuildInvite{
    Code: "example",
}
```

#### Type Definition

```go
type LobbyGuildInvite struct {
    Code string `json:"code"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Code | `string` |  |

### LobbyMember
LobbyMember

#### Example Usage

```go
// Create a new LobbyMember
lobbymember := LobbyMember{
    ID: "example",
    Flags: 42,
    Metadata: map[],
}
```

#### Type Definition

```go
type LobbyMember struct {
    ID string `json:"id"`
    Flags int32 `json:"flags"`
    Metadata map[string]string `json:"metadata,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Flags | `int32` |  |
| Metadata | `map[string]string` |  |

### LobbyMemberOptions
LobbyMemberOptions

#### Example Usage

```go
// Create a new LobbyMemberOptions
lobbymemberoptions := LobbyMemberOptions{
    ID: "example",
    Flags: any{},
    Metadata: map[],
}
```

#### Type Definition

```go
type LobbyMemberOptions struct {
    ID string `json:"id"`
    Flags any `json:"flags,omitempty"`
    Metadata map[string]string `json:"metadata,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Flags | `any` |  |
| Metadata | `map[string]string` |  |

### LobbyMessage
LobbyMessage

#### Example Usage

```go
// Create a new LobbyMessage
lobbymessage := LobbyMessage{
    ID: "example",
    ApplicationID: "example",
    Author: any{},
    ChannelID: "example",
    Content: "example",
    Flags: 42,
    LobbyID: "example",
    Metadata: map[],
    Type: 42,
}
```

#### Type Definition

```go
type LobbyMessage struct {
    ID string `json:"id"`
    ApplicationID string `json:"application_id,omitempty"`
    Author any `json:"author"`
    ChannelID string `json:"channel_id"`
    Content string `json:"content"`
    Flags int32 `json:"flags"`
    LobbyID string `json:"lobby_id"`
    Metadata map[string]string `json:"metadata,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ApplicationID | `string` |  |
| Author | `any` |  |
| ChannelID | `string` |  |
| Content | `string` |  |
| Flags | `int32` |  |
| LobbyID | `string` |  |
| Metadata | `map[string]string` |  |
| Type | `int32` |  |

### MLSpamRule
MLSpamRule

#### Example Usage

```go
// Create a new MLSpamRule
mlspamrule := MLSpamRule{
    ID: "example",
    Actions: [],
    CreatorID: "example",
    Enabled: true,
    EventType: 42,
    ExemptChannels: [],
    ExemptRoles: [],
    GuildID: "example",
    Name: "example",
    TriggerMetadata: any{},
    TriggerType: 42,
}
```

#### Type Definition

```go
type MLSpamRule struct {
    ID string `json:"id"`
    Actions []any `json:"actions"`
    CreatorID string `json:"creator_id"`
    Enabled bool `json:"enabled"`
    EventType int32 `json:"event_type"`
    ExemptChannels []string `json:"exempt_channels"`
    ExemptRoles []string `json:"exempt_roles"`
    GuildID string `json:"guild_id"`
    Name string `json:"name"`
    TriggerMetadata any `json:"trigger_metadata"`
    TriggerType int32 `json:"trigger_type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Actions | `[]any` |  |
| CreatorID | `string` |  |
| Enabled | `bool` |  |
| EventType | `int32` |  |
| ExemptChannels | `[]string` |  |
| ExemptRoles | `[]string` |  |
| GuildID | `string` |  |
| Name | `string` |  |
| TriggerMetadata | `any` |  |
| TriggerType | `int32` |  |

### MLSpamTriggerMetadata
MLSpamTriggerMetadata

#### Example Usage

```go
// Create a new MLSpamTriggerMetadata
mlspamtriggermetadata := MLSpamTriggerMetadata{

}
```

#### Type Definition

```go
type MLSpamTriggerMetadata struct {
}
```

### MLSpamUpsertOptions
MLSpamUpsertOptions

#### Example Usage

```go
// Create a new MLSpamUpsertOptions
mlspamupsertoptions := MLSpamUpsertOptions{
    Actions: [],
    Enabled: &true{},
    EventType: 42,
    ExemptChannels: [],
    ExemptRoles: [],
    Name: "example",
    TriggerMetadata: any{},
    TriggerType: 42,
}
```

#### Type Definition

```go
type MLSpamUpsertOptions struct {
    Actions []any `json:"actions,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    EventType int32 `json:"event_type"`
    ExemptChannels []string `json:"exempt_channels,omitempty"`
    ExemptRoles []string `json:"exempt_roles,omitempty"`
    Name string `json:"name"`
    TriggerMetadata any `json:"trigger_metadata,omitempty"`
    TriggerType int32 `json:"trigger_type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Actions | `[]any` |  |
| Enabled | `*bool` |  |
| EventType | `int32` |  |
| ExemptChannels | `[]string` |  |
| ExemptRoles | `[]string` |  |
| Name | `string` |  |
| TriggerMetadata | `any` |  |
| TriggerType | `int32` |  |

### MLSpamUpsertOptionsPartial
MLSpamUpsertOptionsPartial

#### Example Usage

```go
// Create a new MLSpamUpsertOptionsPartial
mlspamupsertoptionspartial := MLSpamUpsertOptionsPartial{
    Actions: [],
    Enabled: &true{},
    EventType: 42,
    ExemptChannels: [],
    ExemptRoles: [],
    Name: "example",
    TriggerMetadata: any{},
    TriggerType: 42,
}
```

#### Type Definition

```go
type MLSpamUpsertOptionsPartial struct {
    Actions []any `json:"actions,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    EventType int32 `json:"event_type,omitempty"`
    ExemptChannels []string `json:"exempt_channels,omitempty"`
    ExemptRoles []string `json:"exempt_roles,omitempty"`
    Name string `json:"name,omitempty"`
    TriggerMetadata any `json:"trigger_metadata,omitempty"`
    TriggerType int32 `json:"trigger_type,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Actions | `[]any` |  |
| Enabled | `*bool` |  |
| EventType | `int32` |  |
| ExemptChannels | `[]string` |  |
| ExemptRoles | `[]string` |  |
| Name | `string` |  |
| TriggerMetadata | `any` |  |
| TriggerType | `int32` |  |

### MediaGalleryComponent
MediaGalleryComponent

#### Example Usage

```go
// Create a new MediaGalleryComponent
mediagallerycomponent := MediaGalleryComponent{
    ID: 42,
    Items: [],
    Type: 42,
}
```

#### Type Definition

```go
type MediaGalleryComponent struct {
    ID int32 `json:"id"`
    Items []any `json:"items"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| Items | `[]any` |  |
| Type | `int32` |  |

### MediaGalleryComponentForMessageOptions
MediaGalleryComponentForMessageOptions

#### Example Usage

```go
// Create a new MediaGalleryComponentForMessageOptions
mediagallerycomponentformessageoptions := MediaGalleryComponentForMessageOptions{
    ID: &42{},
    Items: [],
    Type: 42,
}
```

#### Type Definition

```go
type MediaGalleryComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    Items []any `json:"items"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| Items | `[]any` |  |
| Type | `int32` |  |

### MediaGalleryItem
MediaGalleryItem

#### Example Usage

```go
// Create a new MediaGalleryItem
mediagalleryitem := MediaGalleryItem{
    Description: &"example"{},
    Media: any{},
    Spoiler: true,
}
```

#### Type Definition

```go
type MediaGalleryItem struct {
    Description *string `json:"description,omitempty"`
    Media any `json:"media"`
    Spoiler bool `json:"spoiler"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `*string` |  |
| Media | `any` |  |
| Spoiler | `bool` |  |

### MediaGalleryItemOptions
MediaGalleryItemOptions

#### Example Usage

```go
// Create a new MediaGalleryItemOptions
mediagalleryitemoptions := MediaGalleryItemOptions{
    Description: &"example"{},
    Media: any{},
    Spoiler: &true{},
}
```

#### Type Definition

```go
type MediaGalleryItemOptions struct {
    Description *string `json:"description,omitempty"`
    Media any `json:"media"`
    Spoiler *bool `json:"spoiler,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `*string` |  |
| Media | `any` |  |
| Spoiler | `*bool` |  |

### MentionSpamRule
MentionSpamRule

#### Example Usage

```go
// Create a new MentionSpamRule
mentionspamrule := MentionSpamRule{
    ID: "example",
    Actions: [],
    CreatorID: "example",
    Enabled: true,
    EventType: 42,
    ExemptChannels: [],
    ExemptRoles: [],
    GuildID: "example",
    Name: "example",
    TriggerMetadata: any{},
    TriggerType: 42,
}
```

#### Type Definition

```go
type MentionSpamRule struct {
    ID string `json:"id"`
    Actions []any `json:"actions"`
    CreatorID string `json:"creator_id"`
    Enabled bool `json:"enabled"`
    EventType int32 `json:"event_type"`
    ExemptChannels []string `json:"exempt_channels"`
    ExemptRoles []string `json:"exempt_roles"`
    GuildID string `json:"guild_id"`
    Name string `json:"name"`
    TriggerMetadata any `json:"trigger_metadata"`
    TriggerType int32 `json:"trigger_type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Actions | `[]any` |  |
| CreatorID | `string` |  |
| Enabled | `bool` |  |
| EventType | `int32` |  |
| ExemptChannels | `[]string` |  |
| ExemptRoles | `[]string` |  |
| GuildID | `string` |  |
| Name | `string` |  |
| TriggerMetadata | `any` |  |
| TriggerType | `int32` |  |

### MentionSpamTriggerMetadata
MentionSpamTriggerMetadata

#### Example Usage

```go
// Create a new MentionSpamTriggerMetadata
mentionspamtriggermetadata := MentionSpamTriggerMetadata{
    MentionRaidProtectionEnabled: &true{},
    MentionTotalLimit: &42{},
}
```

#### Type Definition

```go
type MentionSpamTriggerMetadata struct {
    MentionRaidProtectionEnabled *bool `json:"mention_raid_protection_enabled,omitempty"`
    MentionTotalLimit *int64 `json:"mention_total_limit,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| MentionRaidProtectionEnabled | `*bool` |  |
| MentionTotalLimit | `*int64` |  |

### MentionSpamUpsertOptions
MentionSpamUpsertOptions

#### Example Usage

```go
// Create a new MentionSpamUpsertOptions
mentionspamupsertoptions := MentionSpamUpsertOptions{
    Actions: [],
    Enabled: &true{},
    EventType: 42,
    ExemptChannels: [],
    ExemptRoles: [],
    Name: "example",
    TriggerMetadata: any{},
    TriggerType: 42,
}
```

#### Type Definition

```go
type MentionSpamUpsertOptions struct {
    Actions []any `json:"actions,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    EventType int32 `json:"event_type"`
    ExemptChannels []string `json:"exempt_channels,omitempty"`
    ExemptRoles []string `json:"exempt_roles,omitempty"`
    Name string `json:"name"`
    TriggerMetadata any `json:"trigger_metadata,omitempty"`
    TriggerType int32 `json:"trigger_type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Actions | `[]any` |  |
| Enabled | `*bool` |  |
| EventType | `int32` |  |
| ExemptChannels | `[]string` |  |
| ExemptRoles | `[]string` |  |
| Name | `string` |  |
| TriggerMetadata | `any` |  |
| TriggerType | `int32` |  |

### MentionSpamUpsertOptionsPartial
MentionSpamUpsertOptionsPartial

#### Example Usage

```go
// Create a new MentionSpamUpsertOptionsPartial
mentionspamupsertoptionspartial := MentionSpamUpsertOptionsPartial{
    Actions: [],
    Enabled: &true{},
    EventType: 42,
    ExemptChannels: [],
    ExemptRoles: [],
    Name: "example",
    TriggerMetadata: any{},
    TriggerType: 42,
}
```

#### Type Definition

```go
type MentionSpamUpsertOptionsPartial struct {
    Actions []any `json:"actions,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    EventType int32 `json:"event_type,omitempty"`
    ExemptChannels []string `json:"exempt_channels,omitempty"`
    ExemptRoles []string `json:"exempt_roles,omitempty"`
    Name string `json:"name,omitempty"`
    TriggerMetadata any `json:"trigger_metadata,omitempty"`
    TriggerType int32 `json:"trigger_type,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Actions | `[]any` |  |
| Enabled | `*bool` |  |
| EventType | `int32` |  |
| ExemptChannels | `[]string` |  |
| ExemptRoles | `[]string` |  |
| Name | `string` |  |
| TriggerMetadata | `any` |  |
| TriggerType | `int32` |  |

### MentionableSelectComponent
MentionableSelectComponent

#### Example Usage

```go
// Create a new MentionableSelectComponent
mentionableselectcomponent := MentionableSelectComponent{
    ID: 42,
    CustomID: "example",
    DefaultValues: [],
    Disabled: true,
    MaxValues: &42{},
    MinValues: &42{},
    Placeholder: "example",
    Type: 42,
}
```

#### Type Definition

```go
type MentionableSelectComponent struct {
    ID int32 `json:"id"`
    CustomID string `json:"custom_id"`
    DefaultValues []any `json:"default_values,omitempty"`
    Disabled bool `json:"disabled,omitempty"`
    MaxValues *int32 `json:"max_values,omitempty"`
    MinValues *int32 `json:"min_values,omitempty"`
    Placeholder string `json:"placeholder,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| CustomID | `string` |  |
| DefaultValues | `[]any` |  |
| Disabled | `bool` |  |
| MaxValues | `*int32` |  |
| MinValues | `*int32` |  |
| Placeholder | `string` |  |
| Type | `int32` |  |

### MentionableSelectComponentForMessageOptions
MentionableSelectComponentForMessageOptions

#### Example Usage

```go
// Create a new MentionableSelectComponentForMessageOptions
mentionableselectcomponentformessageoptions := MentionableSelectComponentForMessageOptions{
    ID: &42{},
    CustomID: "example",
    DefaultValues: [],
    Disabled: &true{},
    MaxValues: &42{},
    MinValues: &42{},
    Placeholder: &"example"{},
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type MentionableSelectComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    CustomID string `json:"custom_id"`
    DefaultValues []any `json:"default_values,omitempty"`
    Disabled *bool `json:"disabled,omitempty"`
    MaxValues *int64 `json:"max_values,omitempty"`
    MinValues *int64 `json:"min_values,omitempty"`
    Placeholder *string `json:"placeholder,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| CustomID | `string` |  |
| DefaultValues | `[]any` |  |
| Disabled | `*bool` |  |
| MaxValues | `*int64` |  |
| MinValues | `*int64` |  |
| Placeholder | `*string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### MentionableSelectComponentForModalOptions
MentionableSelectComponentForModalOptions

#### Example Usage

```go
// Create a new MentionableSelectComponentForModalOptions
mentionableselectcomponentformodaloptions := MentionableSelectComponentForModalOptions{
    ID: &42{},
    CustomID: "example",
    DefaultValues: [],
    Disabled: &true{},
    MaxValues: &42{},
    MinValues: &42{},
    Placeholder: &"example"{},
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type MentionableSelectComponentForModalOptions struct {
    ID *int32 `json:"id,omitempty"`
    CustomID string `json:"custom_id"`
    DefaultValues []any `json:"default_values,omitempty"`
    Disabled *bool `json:"disabled,omitempty"`
    MaxValues *int64 `json:"max_values,omitempty"`
    MinValues *int64 `json:"min_values,omitempty"`
    Placeholder *string `json:"placeholder,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| CustomID | `string` |  |
| DefaultValues | `[]any` |  |
| Disabled | `*bool` |  |
| MaxValues | `*int64` |  |
| MinValues | `*int64` |  |
| Placeholder | `*string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### Message
Message

#### Example Usage

```go
// Create a new Message
message := Message{
    ID: "example",
    Activity: any{},
    Application: any{},
    ApplicationID: "example",
    Attachments: [],
    Author: any{},
    Call: any{},
    ChannelID: "example",
    Components: [],
    Content: "example",
    EditedTimestamp: &/* value */{},
    Embeds: [],
    Flags: 42,
    Interaction: any{},
    InteractionMetadata: any{},
    MentionChannels: [],
    MentionEveryone: true,
    MentionRoles: [],
    Mentions: [],
    MessageReference: any{},
    MessageSnapshots: [],
    Nonce: any{},
    Pinned: true,
    Poll: any{},
    Position: 42,
    PurchaseNotification: any{},
    Reactions: [],
    ReferencedMessage: any{},
    Resolved: any{},
    RoleSubscriptionData: any{},
    SharedClientTheme: any{},
    StickerItems: [],
    Stickers: [],
    Thread: any{},
    Timestamp: /* value */,
    Tts: true,
    Type: 42,
    WebhookID: "example",
}
```

#### Type Definition

```go
type Message struct {
    ID string `json:"id"`
    Activity any `json:"activity,omitempty"`
    Application any `json:"application,omitempty"`
    ApplicationID string `json:"application_id,omitempty"`
    Attachments []any `json:"attachments"`
    Author any `json:"author"`
    Call any `json:"call,omitempty"`
    ChannelID string `json:"channel_id"`
    Components []any `json:"components"`
    Content string `json:"content"`
    EditedTimestamp *time.Time `json:"edited_timestamp,omitempty"`
    Embeds []any `json:"embeds"`
    Flags int32 `json:"flags"`
    Interaction any `json:"interaction,omitempty"`
    InteractionMetadata any `json:"interaction_metadata,omitempty"`
    MentionChannels []any `json:"mention_channels,omitempty"`
    MentionEveryone bool `json:"mention_everyone"`
    MentionRoles []string `json:"mention_roles"`
    Mentions []any `json:"mentions"`
    MessageReference any `json:"message_reference,omitempty"`
    MessageSnapshots []any `json:"message_snapshots,omitempty"`
    Nonce any `json:"nonce,omitempty"`
    Pinned bool `json:"pinned"`
    Poll any `json:"poll,omitempty"`
    Position int32 `json:"position,omitempty"`
    PurchaseNotification any `json:"purchase_notification,omitempty"`
    Reactions []any `json:"reactions,omitempty"`
    ReferencedMessage any `json:"referenced_message,omitempty"`
    Resolved any `json:"resolved,omitempty"`
    RoleSubscriptionData any `json:"role_subscription_data,omitempty"`
    SharedClientTheme any `json:"shared_client_theme,omitempty"`
    StickerItems []any `json:"sticker_items,omitempty"`
    Stickers []any `json:"stickers,omitempty"`
    Thread any `json:"thread,omitempty"`
    Timestamp time.Time `json:"timestamp"`
    Tts bool `json:"tts"`
    Type int32 `json:"type"`
    WebhookID string `json:"webhook_id,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Activity | `any` |  |
| Application | `any` |  |
| ApplicationID | `string` |  |
| Attachments | `[]any` |  |
| Author | `any` |  |
| Call | `any` |  |
| ChannelID | `string` |  |
| Components | `[]any` |  |
| Content | `string` |  |
| EditedTimestamp | `*time.Time` |  |
| Embeds | `[]any` |  |
| Flags | `int32` |  |
| Interaction | `any` |  |
| InteractionMetadata | `any` |  |
| MentionChannels | `[]any` |  |
| MentionEveryone | `bool` |  |
| MentionRoles | `[]string` |  |
| Mentions | `[]any` |  |
| MessageReference | `any` |  |
| MessageSnapshots | `[]any` |  |
| Nonce | `any` |  |
| Pinned | `bool` |  |
| Poll | `any` |  |
| Position | `int32` |  |
| PurchaseNotification | `any` |  |
| Reactions | `[]any` |  |
| ReferencedMessage | `any` |  |
| Resolved | `any` |  |
| RoleSubscriptionData | `any` |  |
| SharedClientTheme | `any` |  |
| StickerItems | `[]any` |  |
| Stickers | `[]any` |  |
| Thread | `any` |  |
| Timestamp | `time.Time` |  |
| Tts | `bool` |  |
| Type | `int32` |  |
| WebhookID | `string` |  |

### MessageActivity
MessageActivity

#### Example Usage

```go
// Create a new MessageActivity
messageactivity := MessageActivity{

}
```

#### Type Definition

```go
type MessageActivity struct {
}
```

### MessageAllowedMentionsOptions
MessageAllowedMentionsOptions

#### Example Usage

```go
// Create a new MessageAllowedMentionsOptions
messageallowedmentionsoptions := MessageAllowedMentionsOptions{
    Parse: [],
    RepliedUser: &true{},
    Roles: [],
    Users: [],
}
```

#### Type Definition

```go
type MessageAllowedMentionsOptions struct {
    Parse []any `json:"parse,omitempty"`
    RepliedUser *bool `json:"replied_user,omitempty"`
    Roles []any `json:"roles,omitempty"`
    Users []any `json:"users,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Parse | `[]any` |  |
| RepliedUser | `*bool` |  |
| Roles | `[]any` |  |
| Users | `[]any` |  |

### MessageAttachment
MessageAttachment

#### Example Usage

```go
// Create a new MessageAttachment
messageattachment := MessageAttachment{
    ID: "example",
    Application: any{},
    ClipCreatedAt: /* value */,
    ClipParticipants: [],
    ContentType: "example",
    Description: "example",
    DurationSecs: 3.14,
    Ephemeral: true,
    Filename: "example",
    Height: 42,
    ProxyURL: "example",
    Size: 42,
    Title: &"example"{},
    URL: "example",
    Waveform: "example",
    Width: 42,
}
```

#### Type Definition

```go
type MessageAttachment struct {
    ID string `json:"id"`
    Application any `json:"application,omitempty"`
    ClipCreatedAt time.Time `json:"clip_created_at,omitempty"`
    ClipParticipants []any `json:"clip_participants,omitempty"`
    ContentType string `json:"content_type,omitempty"`
    Description string `json:"description,omitempty"`
    DurationSecs float64 `json:"duration_secs,omitempty"`
    Ephemeral bool `json:"ephemeral,omitempty"`
    Filename string `json:"filename"`
    Height int32 `json:"height,omitempty"`
    ProxyURL string `json:"proxy_url"`
    Size int32 `json:"size"`
    Title *string `json:"title,omitempty"`
    URL string `json:"url"`
    Waveform string `json:"waveform,omitempty"`
    Width int32 `json:"width,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Application | `any` |  |
| ClipCreatedAt | `time.Time` |  |
| ClipParticipants | `[]any` |  |
| ContentType | `string` |  |
| Description | `string` |  |
| DurationSecs | `float64` |  |
| Ephemeral | `bool` |  |
| Filename | `string` |  |
| Height | `int32` |  |
| ProxyURL | `string` |  |
| Size | `int32` |  |
| Title | `*string` |  |
| URL | `string` |  |
| Waveform | `string` |  |
| Width | `int32` |  |

### MessageAttachmentOptions
MessageAttachmentOptions

#### Example Usage

```go
// Create a new MessageAttachmentOptions
messageattachmentoptions := MessageAttachmentOptions{
    ID: "example",
    Description: &"example"{},
    DurationSecs: &3.14{},
    Filename: &"example"{},
    IsRemix: &true{},
    Title: &"example"{},
    Waveform: &"example"{},
}
```

#### Type Definition

```go
type MessageAttachmentOptions struct {
    ID string `json:"id"`
    Description *string `json:"description,omitempty"`
    DurationSecs *float64 `json:"duration_secs,omitempty"`
    Filename *string `json:"filename,omitempty"`
    IsRemix *bool `json:"is_remix,omitempty"`
    Title *string `json:"title,omitempty"`
    Waveform *string `json:"waveform,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Description | `*string` |  |
| DurationSecs | `*float64` |  |
| Filename | `*string` |  |
| IsRemix | `*bool` |  |
| Title | `*string` |  |
| Waveform | `*string` |  |

### MessageCall
MessageCall

#### Example Usage

```go
// Create a new MessageCall
messagecall := MessageCall{
    EndedTimestamp: &/* value */{},
    Participants: [],
}
```

#### Type Definition

```go
type MessageCall struct {
    EndedTimestamp *time.Time `json:"ended_timestamp,omitempty"`
    Participants []string `json:"participants"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| EndedTimestamp | `*time.Time` |  |
| Participants | `[]string` |  |

### MessageComponentInteractionMetadata
MessageComponentInteractionMetadata

#### Example Usage

```go
// Create a new MessageComponentInteractionMetadata
messagecomponentinteractionmetadata := MessageComponentInteractionMetadata{
    ID: "example",
    AuthorizingIntegrationOwners: map[],
    InteractedMessageID: "example",
    OriginalResponseMessageID: "example",
    Type: 42,
    User: any{},
}
```

#### Type Definition

```go
type MessageComponentInteractionMetadata struct {
    ID string `json:"id"`
    AuthorizingIntegrationOwners map[string]string `json:"authorizing_integration_owners"`
    InteractedMessageID string `json:"interacted_message_id"`
    OriginalResponseMessageID string `json:"original_response_message_id,omitempty"`
    Type int32 `json:"type"`
    User any `json:"user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| AuthorizingIntegrationOwners | `map[string]string` |  |
| InteractedMessageID | `string` |  |
| OriginalResponseMessageID | `string` |  |
| Type | `int32` |  |
| User | `any` |  |

### MessageComponentSeparatorSpacingSize
MessageComponentSeparatorSpacingSize

#### Example Usage

```go
// Example usage of MessageComponentSeparatorSpacingSize
var value MessageComponentSeparatorSpacingSize
// Initialize with appropriate value
```

#### Type Definition

```go
type MessageComponentSeparatorSpacingSize int
```

### MessageComponentTypes
MessageComponentTypes

#### Example Usage

```go
// Example usage of MessageComponentTypes
var value MessageComponentTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type MessageComponentTypes int
```

### MessageCreateOptions
MessageCreateOptions

#### Example Usage

```go
// Create a new MessageCreateOptions
messagecreateoptions := MessageCreateOptions{
    AllowedMentions: any{},
    Attachments: [],
    Components: [],
    ConfettiPotion: any{},
    Content: &"example"{},
    Embeds: [],
    EnforceNonce: &true{},
    Flags: &42{},
    MessageReference: any{},
    Nonce: any{},
    Poll: any{},
    SharedClientTheme: any{},
    StickerIds: [],
    Tts: &true{},
}
```

#### Type Definition

```go
type MessageCreateOptions struct {
    AllowedMentions any `json:"allowed_mentions,omitempty"`
    Attachments []any `json:"attachments,omitempty"`
    Components []any `json:"components,omitempty"`
    ConfettiPotion any `json:"confetti_potion,omitempty"`
    Content *string `json:"content,omitempty"`
    Embeds []any `json:"embeds,omitempty"`
    EnforceNonce *bool `json:"enforce_nonce,omitempty"`
    Flags *int64 `json:"flags,omitempty"`
    MessageReference any `json:"message_reference,omitempty"`
    Nonce any `json:"nonce,omitempty"`
    Poll any `json:"poll,omitempty"`
    SharedClientTheme any `json:"shared_client_theme,omitempty"`
    StickerIds []string `json:"sticker_ids,omitempty"`
    Tts *bool `json:"tts,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AllowedMentions | `any` |  |
| Attachments | `[]any` |  |
| Components | `[]any` |  |
| ConfettiPotion | `any` |  |
| Content | `*string` |  |
| Embeds | `[]any` |  |
| EnforceNonce | `*bool` |  |
| Flags | `*int64` |  |
| MessageReference | `any` |  |
| Nonce | `any` |  |
| Poll | `any` |  |
| SharedClientTheme | `any` |  |
| StickerIds | `[]string` |  |
| Tts | `*bool` |  |

### MessageEditOptionsPartial
MessageEditOptionsPartial

#### Example Usage

```go
// Create a new MessageEditOptionsPartial
messageeditoptionspartial := MessageEditOptionsPartial{
    AllowedMentions: any{},
    Attachments: [],
    Components: [],
    Content: &"example"{},
    Embeds: [],
    Flags: &42{},
    StickerIds: [],
}
```

#### Type Definition

```go
type MessageEditOptionsPartial struct {
    AllowedMentions any `json:"allowed_mentions,omitempty"`
    Attachments []any `json:"attachments,omitempty"`
    Components []any `json:"components,omitempty"`
    Content *string `json:"content,omitempty"`
    Embeds []any `json:"embeds,omitempty"`
    Flags *int64 `json:"flags,omitempty"`
    StickerIds []string `json:"sticker_ids,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AllowedMentions | `any` |  |
| Attachments | `[]any` |  |
| Components | `[]any` |  |
| Content | `*string` |  |
| Embeds | `[]any` |  |
| Flags | `*int64` |  |
| StickerIds | `[]string` |  |

### MessageEmbed
MessageEmbed

#### Example Usage

```go
// Create a new MessageEmbed
messageembed := MessageEmbed{
    Author: any{},
    Color: 42,
    Description: "example",
    Fields: [],
    Footer: any{},
    Image: any{},
    Provider: any{},
    Thumbnail: any{},
    Timestamp: /* value */,
    Title: "example",
    Type: "example",
    URL: "example",
    Video: any{},
}
```

#### Type Definition

```go
type MessageEmbed struct {
    Author any `json:"author,omitempty"`
    Color int32 `json:"color,omitempty"`
    Description string `json:"description,omitempty"`
    Fields []any `json:"fields,omitempty"`
    Footer any `json:"footer,omitempty"`
    Image any `json:"image,omitempty"`
    Provider any `json:"provider,omitempty"`
    Thumbnail any `json:"thumbnail,omitempty"`
    Timestamp time.Time `json:"timestamp,omitempty"`
    Title string `json:"title,omitempty"`
    Type string `json:"type"`
    URL string `json:"url,omitempty"`
    Video any `json:"video,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Author | `any` |  |
| Color | `int32` |  |
| Description | `string` |  |
| Fields | `[]any` |  |
| Footer | `any` |  |
| Image | `any` |  |
| Provider | `any` |  |
| Thumbnail | `any` |  |
| Timestamp | `time.Time` |  |
| Title | `string` |  |
| Type | `string` |  |
| URL | `string` |  |
| Video | `any` |  |

### MessageEmbedAuthor
MessageEmbedAuthor

#### Example Usage

```go
// Create a new MessageEmbedAuthor
messageembedauthor := MessageEmbedAuthor{
    IconURL: "example",
    Name: "example",
    ProxyIconURL: "example",
    URL: "example",
}
```

#### Type Definition

```go
type MessageEmbedAuthor struct {
    IconURL string `json:"icon_url,omitempty"`
    Name string `json:"name"`
    ProxyIconURL string `json:"proxy_icon_url,omitempty"`
    URL string `json:"url,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| IconURL | `string` |  |
| Name | `string` |  |
| ProxyIconURL | `string` |  |
| URL | `string` |  |

### MessageEmbedField
MessageEmbedField

#### Example Usage

```go
// Create a new MessageEmbedField
messageembedfield := MessageEmbedField{
    Inline: true,
    Name: "example",
    Value: "example",
}
```

#### Type Definition

```go
type MessageEmbedField struct {
    Inline bool `json:"inline"`
    Name string `json:"name"`
    Value string `json:"value"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Inline | `bool` |  |
| Name | `string` |  |
| Value | `string` |  |

### MessageEmbedFooter
MessageEmbedFooter

#### Example Usage

```go
// Create a new MessageEmbedFooter
messageembedfooter := MessageEmbedFooter{
    IconURL: "example",
    ProxyIconURL: "example",
    Text: "example",
}
```

#### Type Definition

```go
type MessageEmbedFooter struct {
    IconURL string `json:"icon_url,omitempty"`
    ProxyIconURL string `json:"proxy_icon_url,omitempty"`
    Text string `json:"text"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| IconURL | `string` |  |
| ProxyIconURL | `string` |  |
| Text | `string` |  |

### MessageEmbedImage
MessageEmbedImage

#### Example Usage

```go
// Create a new MessageEmbedImage
messageembedimage := MessageEmbedImage{
    ContentType: "example",
    Description: "example",
    Flags: 42,
    Height: 42,
    Placeholder: "example",
    PlaceholderVersion: 42,
    ProxyURL: "example",
    URL: "example",
    Width: 42,
}
```

#### Type Definition

```go
type MessageEmbedImage struct {
    ContentType string `json:"content_type,omitempty"`
    Description string `json:"description,omitempty"`
    Flags int64 `json:"flags,omitempty"`
    Height int64 `json:"height,omitempty"`
    Placeholder string `json:"placeholder,omitempty"`
    PlaceholderVersion int64 `json:"placeholder_version,omitempty"`
    ProxyURL string `json:"proxy_url,omitempty"`
    URL string `json:"url,omitempty"`
    Width int64 `json:"width,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ContentType | `string` |  |
| Description | `string` |  |
| Flags | `int64` |  |
| Height | `int64` |  |
| Placeholder | `string` |  |
| PlaceholderVersion | `int64` |  |
| ProxyURL | `string` |  |
| URL | `string` |  |
| Width | `int64` |  |

### MessageEmbedProvider
MessageEmbedProvider

#### Example Usage

```go
// Create a new MessageEmbedProvider
messageembedprovider := MessageEmbedProvider{
    Name: "example",
    URL: "example",
}
```

#### Type Definition

```go
type MessageEmbedProvider struct {
    Name string `json:"name"`
    URL string `json:"url,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |
| URL | `string` |  |

### MessageEmbedVideo
MessageEmbedVideo

#### Example Usage

```go
// Create a new MessageEmbedVideo
messageembedvideo := MessageEmbedVideo{
    ContentType: "example",
    Description: "example",
    Flags: 42,
    Height: 42,
    Placeholder: "example",
    PlaceholderVersion: 42,
    ProxyURL: "example",
    URL: "example",
    Width: 42,
}
```

#### Type Definition

```go
type MessageEmbedVideo struct {
    ContentType string `json:"content_type,omitempty"`
    Description string `json:"description,omitempty"`
    Flags int64 `json:"flags,omitempty"`
    Height int64 `json:"height,omitempty"`
    Placeholder string `json:"placeholder,omitempty"`
    PlaceholderVersion int64 `json:"placeholder_version,omitempty"`
    ProxyURL string `json:"proxy_url,omitempty"`
    URL string `json:"url,omitempty"`
    Width int64 `json:"width,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ContentType | `string` |  |
| Description | `string` |  |
| Flags | `int64` |  |
| Height | `int64` |  |
| Placeholder | `string` |  |
| PlaceholderVersion | `int64` |  |
| ProxyURL | `string` |  |
| URL | `string` |  |
| Width | `int64` |  |

### MessageInteraction
MessageInteraction

#### Example Usage

```go
// Create a new MessageInteraction
messageinteraction := MessageInteraction{
    ID: "example",
    Name: "example",
    NameLocalized: "example",
    Type: 42,
    User: any{},
}
```

#### Type Definition

```go
type MessageInteraction struct {
    ID string `json:"id"`
    Name string `json:"name"`
    NameLocalized string `json:"name_localized,omitempty"`
    Type int32 `json:"type"`
    User any `json:"user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Name | `string` |  |
| NameLocalized | `string` |  |
| Type | `int32` |  |
| User | `any` |  |

### MessageMentionChannel
MessageMentionChannel

#### Example Usage

```go
// Create a new MessageMentionChannel
messagementionchannel := MessageMentionChannel{
    ID: "example",
    GuildID: "example",
    Name: "example",
    Type: 42,
}
```

#### Type Definition

```go
type MessageMentionChannel struct {
    ID string `json:"id"`
    GuildID string `json:"guild_id"`
    Name string `json:"name"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| GuildID | `string` |  |
| Name | `string` |  |
| Type | `int32` |  |

### MessageReaction
MessageReaction

#### Example Usage

```go
// Create a new MessageReaction
messagereaction := MessageReaction{
    BurstColors: [],
    Count: 42,
    CountDetails: any{},
    Emoji: any{},
    Me: true,
    MeBurst: true,
}
```

#### Type Definition

```go
type MessageReaction struct {
    BurstColors []string `json:"burst_colors"`
    Count int32 `json:"count"`
    CountDetails any `json:"count_details"`
    Emoji any `json:"emoji"`
    Me bool `json:"me"`
    MeBurst bool `json:"me_burst"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| BurstColors | `[]string` |  |
| Count | `int32` |  |
| CountDetails | `any` |  |
| Emoji | `any` |  |
| Me | `bool` |  |
| MeBurst | `bool` |  |

### MessageReactionCountDetails
MessageReactionCountDetails

#### Example Usage

```go
// Create a new MessageReactionCountDetails
messagereactioncountdetails := MessageReactionCountDetails{
    Burst: 42,
    Normal: 42,
}
```

#### Type Definition

```go
type MessageReactionCountDetails struct {
    Burst int32 `json:"burst"`
    Normal int32 `json:"normal"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Burst | `int32` |  |
| Normal | `int32` |  |

### MessageReactionEmoji
MessageReactionEmoji

#### Example Usage

```go
// Create a new MessageReactionEmoji
messagereactionemoji := MessageReactionEmoji{
    ID: any{},
    Animated: true,
    Name: &"example"{},
}
```

#### Type Definition

```go
type MessageReactionEmoji struct {
    ID any `json:"id,omitempty"`
    Animated bool `json:"animated,omitempty"`
    Name *string `json:"name,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `any` |  |
| Animated | `bool` |  |
| Name | `*string` |  |

### MessageReference
MessageReference

#### Example Usage

```go
// Create a new MessageReference
messagereference := MessageReference{
    ChannelID: "example",
    GuildID: "example",
    MessageID: "example",
    Type: 42,
}
```

#### Type Definition

```go
type MessageReference struct {
    ChannelID string `json:"channel_id"`
    GuildID string `json:"guild_id,omitempty"`
    MessageID string `json:"message_id,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `string` |  |
| GuildID | `string` |  |
| MessageID | `string` |  |
| Type | `int32` |  |

### MessageReferenceOptions
MessageReferenceOptions

#### Example Usage

```go
// Create a new MessageReferenceOptions
messagereferenceoptions := MessageReferenceOptions{
    ChannelID: any{},
    FailIfNotExists: &true{},
    GuildID: any{},
    MessageID: "example",
    Type: any{},
}
```

#### Type Definition

```go
type MessageReferenceOptions struct {
    ChannelID any `json:"channel_id,omitempty"`
    FailIfNotExists *bool `json:"fail_if_not_exists,omitempty"`
    GuildID any `json:"guild_id,omitempty"`
    MessageID string `json:"message_id"`
    Type any `json:"type,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `any` |  |
| FailIfNotExists | `*bool` |  |
| GuildID | `any` |  |
| MessageID | `string` |  |
| Type | `any` |  |

### MessageReferenceType
MessageReferenceType

#### Example Usage

```go
// Example usage of MessageReferenceType
var value MessageReferenceType
// Initialize with appropriate value
```

#### Type Definition

```go
type MessageReferenceType int
```

### MessageRoleSubscriptionData
MessageRoleSubscriptionData

#### Example Usage

```go
// Create a new MessageRoleSubscriptionData
messagerolesubscriptiondata := MessageRoleSubscriptionData{
    IsRenewal: true,
    RoleSubscriptionListingID: "example",
    TierName: "example",
    TotalMonthsSubscribed: 42,
}
```

#### Type Definition

```go
type MessageRoleSubscriptionData struct {
    IsRenewal bool `json:"is_renewal"`
    RoleSubscriptionListingID string `json:"role_subscription_listing_id"`
    TierName string `json:"tier_name"`
    TotalMonthsSubscribed int32 `json:"total_months_subscribed"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| IsRenewal | `bool` |  |
| RoleSubscriptionListingID | `string` |  |
| TierName | `string` |  |
| TotalMonthsSubscribed | `int32` |  |

### MessageShareCustomUserThemeBaseTheme
MessageShareCustomUserThemeBaseTheme

#### Example Usage

```go
// Example usage of MessageShareCustomUserThemeBaseTheme
var value MessageShareCustomUserThemeBaseTheme
// Initialize with appropriate value
```

#### Type Definition

```go
type MessageShareCustomUserThemeBaseTheme int
```

### MessageSnapshot
MessageSnapshot

#### Example Usage

```go
// Create a new MessageSnapshot
messagesnapshot := MessageSnapshot{
    Message: any{},
}
```

#### Type Definition

```go
type MessageSnapshot struct {
    Message any `json:"message"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Message | `any` |  |

### MessageStickerItem
MessageStickerItem

#### Example Usage

```go
// Create a new MessageStickerItem
messagestickeritem := MessageStickerItem{
    ID: "example",
    FormatType: 42,
    Name: "example",
}
```

#### Type Definition

```go
type MessageStickerItem struct {
    ID string `json:"id"`
    FormatType int32 `json:"format_type"`
    Name string `json:"name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| FormatType | `int32` |  |
| Name | `string` |  |

### MessageType
MessageType

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

### MetadataItemTypes
MetadataItemTypes

#### Example Usage

```go
// Example usage of MetadataItemTypes
var value MetadataItemTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type MetadataItemTypes int
```

### MinimalContentMessage
MinimalContentMessage

#### Example Usage

```go
// Create a new MinimalContentMessage
minimalcontentmessage := MinimalContentMessage{
    Attachments: [],
    Components: [],
    Content: "example",
    EditedTimestamp: &/* value */{},
    Embeds: [],
    Flags: 42,
    MentionRoles: [],
    Mentions: [],
    StickerItems: [],
    Stickers: [],
    Timestamp: /* value */,
    Type: 42,
}
```

#### Type Definition

```go
type MinimalContentMessage struct {
    Attachments []any `json:"attachments"`
    Components []any `json:"components"`
    Content string `json:"content"`
    EditedTimestamp *time.Time `json:"edited_timestamp,omitempty"`
    Embeds []any `json:"embeds"`
    Flags int32 `json:"flags"`
    MentionRoles []string `json:"mention_roles"`
    Mentions []any `json:"mentions"`
    StickerItems []any `json:"sticker_items,omitempty"`
    Stickers []any `json:"stickers,omitempty"`
    Timestamp time.Time `json:"timestamp"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Attachments | `[]any` |  |
| Components | `[]any` |  |
| Content | `string` |  |
| EditedTimestamp | `*time.Time` |  |
| Embeds | `[]any` |  |
| Flags | `int32` |  |
| MentionRoles | `[]string` |  |
| Mentions | `[]any` |  |
| StickerItems | `[]any` |  |
| Stickers | `[]any` |  |
| Timestamp | `time.Time` |  |
| Type | `int32` |  |

### ModalInteractionCallbackOptions
ModalInteractionCallbackOptions

#### Example Usage

```go
// Create a new ModalInteractionCallbackOptions
modalinteractioncallbackoptions := ModalInteractionCallbackOptions{
    Data: any{},
    Type: 42,
}
```

#### Type Definition

```go
type ModalInteractionCallbackOptions struct {
    Data any `json:"data"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Data | `any` |  |
| Type | `int32` |  |

### ModalInteractionCallbackOptionsData
ModalInteractionCallbackOptionsData

#### Example Usage

```go
// Create a new ModalInteractionCallbackOptionsData
modalinteractioncallbackoptionsdata := ModalInteractionCallbackOptionsData{
    Components: [],
    CustomID: "example",
    Title: "example",
}
```

#### Type Definition

```go
type ModalInteractionCallbackOptionsData struct {
    Components []any `json:"components"`
    CustomID string `json:"custom_id"`
    Title string `json:"title"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Components | `[]any` |  |
| CustomID | `string` |  |
| Title | `string` |  |

### ModalSubmitInteractionMetadata
ModalSubmitInteractionMetadata

#### Example Usage

```go
// Create a new ModalSubmitInteractionMetadata
modalsubmitinteractionmetadata := ModalSubmitInteractionMetadata{
    ID: "example",
    AuthorizingIntegrationOwners: map[],
    OriginalResponseMessageID: "example",
    TriggeringInteractionMetadata: any{},
    Type: 42,
    User: any{},
}
```

#### Type Definition

```go
type ModalSubmitInteractionMetadata struct {
    ID string `json:"id"`
    AuthorizingIntegrationOwners map[string]string `json:"authorizing_integration_owners"`
    OriginalResponseMessageID string `json:"original_response_message_id,omitempty"`
    TriggeringInteractionMetadata any `json:"triggering_interaction_metadata"`
    Type int32 `json:"type"`
    User any `json:"user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| AuthorizingIntegrationOwners | `map[string]string` |  |
| OriginalResponseMessageID | `string` |  |
| TriggeringInteractionMetadata | `any` |  |
| Type | `int32` |  |
| User | `any` |  |

### MyGuild
MyGuild

#### Example Usage

```go
// Create a new MyGuild
myguild := MyGuild{
    ID: "example",
    ApproximateMemberCount: &42{},
    ApproximatePresenceCount: &42{},
    Banner: &"example"{},
    Features: [],
    Icon: &"example"{},
    Name: "example",
    Owner: true,
    Permissions: "example",
}
```

#### Type Definition

```go
type MyGuild struct {
    ID string `json:"id"`
    ApproximateMemberCount *int32 `json:"approximate_member_count,omitempty"`
    ApproximatePresenceCount *int32 `json:"approximate_presence_count,omitempty"`
    Banner *string `json:"banner,omitempty"`
    Features []string `json:"features"`
    Icon *string `json:"icon,omitempty"`
    Name string `json:"name"`
    Owner bool `json:"owner"`
    Permissions string `json:"permissions"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ApproximateMemberCount | `*int32` |  |
| ApproximatePresenceCount | `*int32` |  |
| Banner | `*string` |  |
| Features | `[]string` |  |
| Icon | `*string` |  |
| Name | `string` |  |
| Owner | `bool` |  |
| Permissions | `string` |  |

### NameplatePalette
NameplatePalette

#### Example Usage

```go
// Create a new NameplatePalette
nameplatepalette := NameplatePalette{

}
```

#### Type Definition

```go
type NameplatePalette struct {
}
```

### NewMemberAction
NewMemberAction

#### Example Usage

```go
// Create a new NewMemberAction
newmemberaction := NewMemberAction{
    ActionType: 42,
    ChannelID: "example",
    Description: "example",
    Emoji: any{},
    Icon: "example",
    Title: "example",
}
```

#### Type Definition

```go
type NewMemberAction struct {
    ActionType int32 `json:"action_type"`
    ChannelID string `json:"channel_id"`
    Description string `json:"description"`
    Emoji any `json:"emoji,omitempty"`
    Icon string `json:"icon,omitempty"`
    Title string `json:"title"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ActionType | `int32` |  |
| ChannelID | `string` |  |
| Description | `string` |  |
| Emoji | `any` |  |
| Icon | `string` |  |
| Title | `string` |  |

### NewMemberActionType
NewMemberActionType

#### Example Usage

```go
// Example usage of NewMemberActionType
var value NewMemberActionType
// Initialize with appropriate value
```

#### Type Definition

```go
type NewMemberActionType int
```

### OAuth2GetAuthorization
OAuth2GetAuthorization

#### Example Usage

```go
// Create a new OAuth2GetAuthorization
oauth2getauthorization := OAuth2GetAuthorization{
    Application: any{},
    Expires: /* value */,
    Scopes: [],
    User: any{},
}
```

#### Type Definition

```go
type OAuth2GetAuthorization struct {
    Application any `json:"application"`
    Expires time.Time `json:"expires"`
    Scopes []string `json:"scopes"`
    User any `json:"user,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Application | `any` |  |
| Expires | `time.Time` |  |
| Scopes | `[]string` |  |
| User | `any` |  |

### OAuth2GetKeys
OAuth2GetKeys

#### Example Usage

```go
// Create a new OAuth2GetKeys
oauth2getkeys := OAuth2GetKeys{
    Keys: [],
}
```

#### Type Definition

```go
type OAuth2GetKeys struct {
    Keys []any `json:"keys"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Keys | `[]any` |  |

### OAuth2GetOpenIDConnectUserInfo
OAuth2GetOpenIDConnectUserInfo

#### Example Usage

```go
// Create a new OAuth2GetOpenIDConnectUserInfo
oauth2getopenidconnectuserinfo := OAuth2GetOpenIDConnectUserInfo{
    Email: &"example"{},
    EmailVerified: true,
    Locale: "example",
    Nickname: &"example"{},
    Picture: "example",
    PreferredUsername: "example",
    Sub: "example",
}
```

#### Type Definition

```go
type OAuth2GetOpenIDConnectUserInfo struct {
    Email *string `json:"email,omitempty"`
    EmailVerified bool `json:"email_verified,omitempty"`
    Locale string `json:"locale,omitempty"`
    Nickname *string `json:"nickname,omitempty"`
    Picture string `json:"picture,omitempty"`
    PreferredUsername string `json:"preferred_username,omitempty"`
    Sub string `json:"sub"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Email | `*string` |  |
| EmailVerified | `bool` |  |
| Locale | `string` |  |
| Nickname | `*string` |  |
| Picture | `string` |  |
| PreferredUsername | `string` |  |
| Sub | `string` |  |

### OAuth2Key
OAuth2Key

#### Example Usage

```go
// Create a new OAuth2Key
oauth2key := OAuth2Key{
    Alg: "example",
    E: "example",
    Kid: "example",
    Kty: "example",
    N: "example",
    Use: "example",
}
```

#### Type Definition

```go
type OAuth2Key struct {
    Alg string `json:"alg"`
    E string `json:"e"`
    Kid string `json:"kid"`
    Kty string `json:"kty"`
    N string `json:"n"`
    Use string `json:"use"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Alg | `string` |  |
| E | `string` |  |
| Kid | `string` |  |
| Kty | `string` |  |
| N | `string` |  |
| Use | `string` |  |

### OAuth2Scopes
OAuth2Scopes

#### Example Usage

```go
// Example usage of OAuth2Scopes
var value OAuth2Scopes
// Initialize with appropriate value
```

#### Type Definition

```go
type OAuth2Scopes string
```

### OnboardingPrompt
OnboardingPrompt

#### Example Usage

```go
// Create a new OnboardingPrompt
onboardingprompt := OnboardingPrompt{
    ID: "example",
    InOnboarding: true,
    Options: [],
    Required: true,
    SingleSelect: true,
    Title: "example",
    Type: 42,
}
```

#### Type Definition

```go
type OnboardingPrompt struct {
    ID string `json:"id"`
    InOnboarding bool `json:"in_onboarding"`
    Options []any `json:"options"`
    Required bool `json:"required"`
    SingleSelect bool `json:"single_select"`
    Title string `json:"title"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| InOnboarding | `bool` |  |
| Options | `[]any` |  |
| Required | `bool` |  |
| SingleSelect | `bool` |  |
| Title | `string` |  |
| Type | `int32` |  |

### OnboardingPromptOption
OnboardingPromptOption

#### Example Usage

```go
// Create a new OnboardingPromptOption
onboardingpromptoption := OnboardingPromptOption{
    ID: "example",
    ChannelIds: [],
    Description: "example",
    Emoji: any{},
    RoleIds: [],
    Title: "example",
}
```

#### Type Definition

```go
type OnboardingPromptOption struct {
    ID string `json:"id"`
    ChannelIds []string `json:"channel_ids"`
    Description string `json:"description"`
    Emoji any `json:"emoji"`
    RoleIds []string `json:"role_ids"`
    Title string `json:"title"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ChannelIds | `[]string` |  |
| Description | `string` |  |
| Emoji | `any` |  |
| RoleIds | `[]string` |  |
| Title | `string` |  |

### OnboardingPromptOptionOptions
OnboardingPromptOptionOptions

#### Example Usage

```go
// Create a new OnboardingPromptOptionOptions
onboardingpromptoptionoptions := OnboardingPromptOptionOptions{
    ID: any{},
    ChannelIds: [],
    Description: &"example"{},
    EmojiAnimated: &true{},
    EmojiID: any{},
    EmojiName: &"example"{},
    RoleIds: [],
    Title: "example",
}
```

#### Type Definition

```go
type OnboardingPromptOptionOptions struct {
    ID any `json:"id,omitempty"`
    ChannelIds []string `json:"channel_ids,omitempty"`
    Description *string `json:"description,omitempty"`
    EmojiAnimated *bool `json:"emoji_animated,omitempty"`
    EmojiID any `json:"emoji_id,omitempty"`
    EmojiName *string `json:"emoji_name,omitempty"`
    RoleIds []string `json:"role_ids,omitempty"`
    Title string `json:"title"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `any` |  |
| ChannelIds | `[]string` |  |
| Description | `*string` |  |
| EmojiAnimated | `*bool` |  |
| EmojiID | `any` |  |
| EmojiName | `*string` |  |
| RoleIds | `[]string` |  |
| Title | `string` |  |

### OnboardingPromptType
OnboardingPromptType

#### Example Usage

```go
// Example usage of OnboardingPromptType
var value OnboardingPromptType
// Initialize with appropriate value
```

#### Type Definition

```go
type OnboardingPromptType int
```

### PartialDiscordIntegration
PartialDiscordIntegration

#### Example Usage

```go
// Create a new PartialDiscordIntegration
partialdiscordintegration := PartialDiscordIntegration{
    ID: "example",
    Account: any{},
    ApplicationID: "example",
    Name: &"example"{},
    Type: "example",
}
```

#### Type Definition

```go
type PartialDiscordIntegration struct {
    ID string `json:"id"`
    Account any `json:"account"`
    ApplicationID string `json:"application_id"`
    Name *string `json:"name,omitempty"`
    Type string `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Account | `any` |  |
| ApplicationID | `string` |  |
| Name | `*string` |  |
| Type | `string` |  |

### PartialExternalConnectionIntegration
PartialExternalConnectionIntegration

#### Example Usage

```go
// Create a new PartialExternalConnectionIntegration
partialexternalconnectionintegration := PartialExternalConnectionIntegration{
    ID: "example",
    Account: any{},
    Name: &"example"{},
    Type: "example",
}
```

#### Type Definition

```go
type PartialExternalConnectionIntegration struct {
    ID string `json:"id"`
    Account any `json:"account"`
    Name *string `json:"name,omitempty"`
    Type string `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Account | `any` |  |
| Name | `*string` |  |
| Type | `string` |  |

### PartialGuildSubscriptionIntegration
PartialGuildSubscriptionIntegration

#### Example Usage

```go
// Create a new PartialGuildSubscriptionIntegration
partialguildsubscriptionintegration := PartialGuildSubscriptionIntegration{
    ID: "example",
    Account: any{},
    Name: &"example"{},
    Type: "example",
}
```

#### Type Definition

```go
type PartialGuildSubscriptionIntegration struct {
    ID string `json:"id"`
    Account any `json:"account"`
    Name *string `json:"name,omitempty"`
    Type string `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Account | `any` |  |
| Name | `*string` |  |
| Type | `string` |  |

### PinnedMessage
PinnedMessage

#### Example Usage

```go
// Create a new PinnedMessage
pinnedmessage := PinnedMessage{
    Message: any{},
    PinnedAt: /* value */,
}
```

#### Type Definition

```go
type PinnedMessage struct {
    Message any `json:"message"`
    PinnedAt time.Time `json:"pinned_at"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Message | `any` |  |
| PinnedAt | `time.Time` |  |

### PinnedMessages
PinnedMessages

#### Example Usage

```go
// Create a new PinnedMessages
pinnedmessages := PinnedMessages{
    HasMore: true,
    Items: [],
}
```

#### Type Definition

```go
type PinnedMessages struct {
    HasMore bool `json:"has_more"`
    Items []any `json:"items"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| HasMore | `bool` |  |
| Items | `[]any` |  |

### Poll
Poll

#### Example Usage

```go
// Create a new Poll
poll := Poll{
    AllowMultiselect: true,
    Answers: [],
    Expiry: /* value */,
    LayoutType: 42,
    Question: any{},
    Results: any{},
}
```

#### Type Definition

```go
type Poll struct {
    AllowMultiselect bool `json:"allow_multiselect"`
    Answers []any `json:"answers"`
    Expiry time.Time `json:"expiry"`
    LayoutType int32 `json:"layout_type"`
    Question any `json:"question"`
    Results any `json:"results"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AllowMultiselect | `bool` |  |
| Answers | `[]any` |  |
| Expiry | `time.Time` |  |
| LayoutType | `int32` |  |
| Question | `any` |  |
| Results | `any` |  |

### PollAnswer
PollAnswer

#### Example Usage

```go
// Create a new PollAnswer
pollanswer := PollAnswer{
    AnswerID: 42,
    PollMedia: any{},
}
```

#### Type Definition

```go
type PollAnswer struct {
    AnswerID int32 `json:"answer_id"`
    PollMedia any `json:"poll_media"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AnswerID | `int32` |  |
| PollMedia | `any` |  |

### PollAnswerCreateOptions
PollAnswerCreateOptions

#### Example Usage

```go
// Create a new PollAnswerCreateOptions
pollanswercreateoptions := PollAnswerCreateOptions{
    PollMedia: any{},
}
```

#### Type Definition

```go
type PollAnswerCreateOptions struct {
    PollMedia any `json:"poll_media"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| PollMedia | `any` |  |

### PollAnswerDetails
PollAnswerDetails

#### Example Usage

```go
// Create a new PollAnswerDetails
pollanswerdetails := PollAnswerDetails{
    Users: [],
}
```

#### Type Definition

```go
type PollAnswerDetails struct {
    Users []any `json:"users"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Users | `[]any` |  |

### PollCreateOptions
PollCreateOptions

#### Example Usage

```go
// Create a new PollCreateOptions
pollcreateoptions := PollCreateOptions{
    AllowMultiselect: &true{},
    Answers: [],
    Duration: &42{},
    LayoutType: any{},
    Question: any{},
}
```

#### Type Definition

```go
type PollCreateOptions struct {
    AllowMultiselect *bool `json:"allow_multiselect,omitempty"`
    Answers []any `json:"answers"`
    Duration *int32 `json:"duration,omitempty"`
    LayoutType any `json:"layout_type,omitempty"`
    Question any `json:"question"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AllowMultiselect | `*bool` |  |
| Answers | `[]any` |  |
| Duration | `*int32` |  |
| LayoutType | `any` |  |
| Question | `any` |  |

### PollEmoji
PollEmoji

#### Example Usage

```go
// Create a new PollEmoji
pollemoji := PollEmoji{
    ID: any{},
    Animated: &true{},
    Name: &"example"{},
}
```

#### Type Definition

```go
type PollEmoji struct {
    ID any `json:"id,omitempty"`
    Animated *bool `json:"animated,omitempty"`
    Name *string `json:"name,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `any` |  |
| Animated | `*bool` |  |
| Name | `*string` |  |

### PollEmojiCreateOptions
PollEmojiCreateOptions

#### Example Usage

```go
// Create a new PollEmojiCreateOptions
pollemojicreateoptions := PollEmojiCreateOptions{
    ID: any{},
    Animated: &true{},
    Name: &"example"{},
}
```

#### Type Definition

```go
type PollEmojiCreateOptions struct {
    ID any `json:"id,omitempty"`
    Animated *bool `json:"animated,omitempty"`
    Name *string `json:"name,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `any` |  |
| Animated | `*bool` |  |
| Name | `*string` |  |

### PollLayoutTypes
PollLayoutTypes

#### Example Usage

```go
// Create a new PollLayoutTypes
polllayouttypes := PollLayoutTypes{

}
```

#### Type Definition

```go
type PollLayoutTypes struct {
}
```

### PollMedia
PollMedia

#### Example Usage

```go
// Create a new PollMedia
pollmedia := PollMedia{
    Emoji: any{},
    Text: "example",
}
```

#### Type Definition

```go
type PollMedia struct {
    Emoji any `json:"emoji,omitempty"`
    Text string `json:"text,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Emoji | `any` |  |
| Text | `string` |  |

### PollMediaCreateOptions
PollMediaCreateOptions

#### Example Usage

```go
// Create a new PollMediaCreateOptions
pollmediacreateoptions := PollMediaCreateOptions{
    Emoji: any{},
    Text: &"example"{},
}
```

#### Type Definition

```go
type PollMediaCreateOptions struct {
    Emoji any `json:"emoji,omitempty"`
    Text *string `json:"text,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Emoji | `any` |  |
| Text | `*string` |  |

### PollResults
PollResults

#### Example Usage

```go
// Create a new PollResults
pollresults := PollResults{
    AnswerCounts: [],
    IsFinalized: true,
}
```

#### Type Definition

```go
type PollResults struct {
    AnswerCounts []any `json:"answer_counts"`
    IsFinalized bool `json:"is_finalized"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AnswerCounts | `[]any` |  |
| IsFinalized | `bool` |  |

### PollResultsEntry
PollResultsEntry

#### Example Usage

```go
// Create a new PollResultsEntry
pollresultsentry := PollResultsEntry{
    ID: 42,
    Count: 42,
    MeVoted: true,
}
```

#### Type Definition

```go
type PollResultsEntry struct {
    ID int32 `json:"id"`
    Count int32 `json:"count"`
    MeVoted bool `json:"me_voted"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| Count | `int32` |  |
| MeVoted | `bool` |  |

### PongInteractionCallbackOptions
PongInteractionCallbackOptions

#### Example Usage

```go
// Create a new PongInteractionCallbackOptions
ponginteractioncallbackoptions := PongInteractionCallbackOptions{
    Type: 42,
}
```

#### Type Definition

```go
type PongInteractionCallbackOptions struct {
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Type | `int32` |  |

### PremiumGuildTiers
PremiumGuildTiers

#### Example Usage

```go
// Example usage of PremiumGuildTiers
var value PremiumGuildTiers
// Initialize with appropriate value
```

#### Type Definition

```go
type PremiumGuildTiers int
```

### PremiumTypes
PremiumTypes

#### Example Usage

```go
// Example usage of PremiumTypes
var value PremiumTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type PremiumTypes int
```

### PrivateApplication
PrivateApplication

#### Example Usage

```go
// Create a new PrivateApplication
privateapplication := PrivateApplication{
    ID: "example",
    ApproximateGuildCount: &42{},
    ApproximateUserAuthorizationCount: 42,
    ApproximateUserInstallCount: 42,
    Bot: any{},
    BotPublic: true,
    BotRequireCodeGrant: true,
    CoverImage: "example",
    CustomInstallURL: "example",
    Description: "example",
    ExplicitContentFilter: 42,
    Flags: 42,
    GuildID: "example",
    Icon: &"example"{},
    InstallParams: any{},
    IntegrationTypesConfig: map[],
    InteractionsEndpointURL: &"example"{},
    MaxParticipants: &42{},
    Name: "example",
    Owner: any{},
    PrimarySkuID: "example",
    PrivacyPolicyURL: "example",
    RedirectUris: [],
    RoleConnectionsVerificationURL: &"example"{},
    RpcOrigins: [],
    Slug: "example",
    Tags: [],
    Team: any{},
    TermsOfServiceURL: "example",
    Type: any{},
    VerifyKey: "example",
}
```

#### Type Definition

```go
type PrivateApplication struct {
    ID string `json:"id"`
    ApproximateGuildCount *int32 `json:"approximate_guild_count,omitempty"`
    ApproximateUserAuthorizationCount int32 `json:"approximate_user_authorization_count"`
    ApproximateUserInstallCount int32 `json:"approximate_user_install_count"`
    Bot any `json:"bot,omitempty"`
    BotPublic bool `json:"bot_public,omitempty"`
    BotRequireCodeGrant bool `json:"bot_require_code_grant,omitempty"`
    CoverImage string `json:"cover_image,omitempty"`
    CustomInstallURL string `json:"custom_install_url,omitempty"`
    Description string `json:"description"`
    ExplicitContentFilter int32 `json:"explicit_content_filter"`
    Flags int32 `json:"flags"`
    GuildID string `json:"guild_id,omitempty"`
    Icon *string `json:"icon,omitempty"`
    InstallParams any `json:"install_params,omitempty"`
    IntegrationTypesConfig map[string]any `json:"integration_types_config,omitempty"`
    InteractionsEndpointURL *string `json:"interactions_endpoint_url,omitempty"`
    MaxParticipants *int32 `json:"max_participants,omitempty"`
    Name string `json:"name"`
    Owner any `json:"owner"`
    PrimarySkuID string `json:"primary_sku_id,omitempty"`
    PrivacyPolicyURL string `json:"privacy_policy_url,omitempty"`
    RedirectUris []*string `json:"redirect_uris"`
    RoleConnectionsVerificationURL *string `json:"role_connections_verification_url,omitempty"`
    RpcOrigins []*string `json:"rpc_origins,omitempty"`
    Slug string `json:"slug,omitempty"`
    Tags []string `json:"tags,omitempty"`
    Team any `json:"team,omitempty"`
    TermsOfServiceURL string `json:"terms_of_service_url,omitempty"`
    Type any `json:"type,omitempty"`
    VerifyKey string `json:"verify_key"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ApproximateGuildCount | `*int32` |  |
| ApproximateUserAuthorizationCount | `int32` |  |
| ApproximateUserInstallCount | `int32` |  |
| Bot | `any` |  |
| BotPublic | `bool` |  |
| BotRequireCodeGrant | `bool` |  |
| CoverImage | `string` |  |
| CustomInstallURL | `string` |  |
| Description | `string` |  |
| ExplicitContentFilter | `int32` |  |
| Flags | `int32` |  |
| GuildID | `string` |  |
| Icon | `*string` |  |
| InstallParams | `any` |  |
| IntegrationTypesConfig | `map[string]any` |  |
| InteractionsEndpointURL | `*string` |  |
| MaxParticipants | `*int32` |  |
| Name | `string` |  |
| Owner | `any` |  |
| PrimarySkuID | `string` |  |
| PrivacyPolicyURL | `string` |  |
| RedirectUris | `[]*string` |  |
| RoleConnectionsVerificationURL | `*string` |  |
| RpcOrigins | `[]*string` |  |
| Slug | `string` |  |
| Tags | `[]string` |  |
| Team | `any` |  |
| TermsOfServiceURL | `string` |  |
| Type | `any` |  |
| VerifyKey | `string` |  |

### PrivateChannel
PrivateChannel

#### Example Usage

```go
// Create a new PrivateChannel
privatechannel := PrivateChannel{
    ID: "example",
    Flags: 42,
    LastMessageID: any{},
    LastPinTimestamp: &/* value */{},
    Recipients: [],
    Type: 42,
}
```

#### Type Definition

```go
type PrivateChannel struct {
    ID string `json:"id"`
    Flags int32 `json:"flags"`
    LastMessageID any `json:"last_message_id,omitempty"`
    LastPinTimestamp *time.Time `json:"last_pin_timestamp,omitempty"`
    Recipients []any `json:"recipients"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Flags | `int32` |  |
| LastMessageID | `any` |  |
| LastPinTimestamp | `*time.Time` |  |
| Recipients | `[]any` |  |
| Type | `int32` |  |

### PrivateChannelLocation
PrivateChannelLocation

#### Example Usage

```go
// Create a new PrivateChannelLocation
privatechannellocation := PrivateChannelLocation{
    ID: "example",
    ChannelID: "example",
    Kind: "example",
}
```

#### Type Definition

```go
type PrivateChannelLocation struct {
    ID string `json:"id"`
    ChannelID string `json:"channel_id"`
    Kind string `json:"kind"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ChannelID | `string` |  |
| Kind | `string` |  |

### PrivateGroupChannel
PrivateGroupChannel

#### Example Usage

```go
// Create a new PrivateGroupChannel
privategroupchannel := PrivateGroupChannel{
    ID: "example",
    ApplicationID: "example",
    Flags: 42,
    Icon: &"example"{},
    LastMessageID: any{},
    LastPinTimestamp: &/* value */{},
    Managed: true,
    Name: &"example"{},
    OwnerID: "example",
    Recipients: [],
    Type: 42,
}
```

#### Type Definition

```go
type PrivateGroupChannel struct {
    ID string `json:"id"`
    ApplicationID string `json:"application_id,omitempty"`
    Flags int32 `json:"flags"`
    Icon *string `json:"icon,omitempty"`
    LastMessageID any `json:"last_message_id,omitempty"`
    LastPinTimestamp *time.Time `json:"last_pin_timestamp,omitempty"`
    Managed bool `json:"managed,omitempty"`
    Name *string `json:"name,omitempty"`
    OwnerID string `json:"owner_id"`
    Recipients []any `json:"recipients"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ApplicationID | `string` |  |
| Flags | `int32` |  |
| Icon | `*string` |  |
| LastMessageID | `any` |  |
| LastPinTimestamp | `*time.Time` |  |
| Managed | `bool` |  |
| Name | `*string` |  |
| OwnerID | `string` |  |
| Recipients | `[]any` |  |
| Type | `int32` |  |

### PrivateGuildMember
PrivateGuildMember

#### Example Usage

```go
// Create a new PrivateGuildMember
privateguildmember := PrivateGuildMember{
    Avatar: &"example"{},
    AvatarDecorationData: any{},
    Banner: &"example"{},
    Collectibles: any{},
    CommunicationDisabledUntil: &/* value */{},
    Deaf: true,
    Flags: 42,
    JoinedAt: /* value */,
    Mute: true,
    Nick: &"example"{},
    Pending: true,
    Permissions: "example",
    PremiumSince: &/* value */{},
    Roles: [],
    User: any{},
}
```

#### Type Definition

```go
type PrivateGuildMember struct {
    Avatar *string `json:"avatar,omitempty"`
    AvatarDecorationData any `json:"avatar_decoration_data,omitempty"`
    Banner *string `json:"banner,omitempty"`
    Collectibles any `json:"collectibles,omitempty"`
    CommunicationDisabledUntil *time.Time `json:"communication_disabled_until,omitempty"`
    Deaf bool `json:"deaf"`
    Flags int32 `json:"flags"`
    JoinedAt time.Time `json:"joined_at"`
    Mute bool `json:"mute"`
    Nick *string `json:"nick,omitempty"`
    Pending bool `json:"pending"`
    Permissions string `json:"permissions,omitempty"`
    PremiumSince *time.Time `json:"premium_since,omitempty"`
    Roles []string `json:"roles"`
    User any `json:"user"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Avatar | `*string` |  |
| AvatarDecorationData | `any` |  |
| Banner | `*string` |  |
| Collectibles | `any` |  |
| CommunicationDisabledUntil | `*time.Time` |  |
| Deaf | `bool` |  |
| Flags | `int32` |  |
| JoinedAt | `time.Time` |  |
| Mute | `bool` |  |
| Nick | `*string` |  |
| Pending | `bool` |  |
| Permissions | `string` |  |
| PremiumSince | `*time.Time` |  |
| Roles | `[]string` |  |
| User | `any` |  |

### ProvisionalToken
ProvisionalToken

#### Example Usage

```go
// Create a new ProvisionalToken
provisionaltoken := ProvisionalToken{
    AccessToken: "example",
    ExpiresAtS: &42{},
    ExpiresIn: 42,
    IDToken: "example",
    RefreshToken: &"example"{},
    Scope: "example",
    Scopes: [],
    TokenType: "example",
}
```

#### Type Definition

```go
type ProvisionalToken struct {
    AccessToken string `json:"access_token"`
    ExpiresAtS *int32 `json:"expires_at_s,omitempty"`
    ExpiresIn int32 `json:"expires_in"`
    IDToken string `json:"id_token"`
    RefreshToken *string `json:"refresh_token,omitempty"`
    Scope string `json:"scope"`
    Scopes []string `json:"scopes,omitempty"`
    TokenType string `json:"token_type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AccessToken | `string` |  |
| ExpiresAtS | `*int32` |  |
| ExpiresIn | `int32` |  |
| IDToken | `string` |  |
| RefreshToken | `*string` |  |
| Scope | `string` |  |
| Scopes | `[]string` |  |
| TokenType | `string` |  |

### PruneGuildOptions
PruneGuildOptions

#### Example Usage

```go
// Create a new PruneGuildOptions
pruneguildoptions := PruneGuildOptions{
    ComputePruneCount: &true{},
    Days: &42{},
    IncludeRoles: any{},
}
```

#### Type Definition

```go
type PruneGuildOptions struct {
    ComputePruneCount *bool `json:"compute_prune_count,omitempty"`
    Days *int64 `json:"days,omitempty"`
    IncludeRoles any `json:"include_roles,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ComputePruneCount | `*bool` |  |
| Days | `*int64` |  |
| IncludeRoles | `any` |  |

### PurchaseNotification
PurchaseNotification

#### Example Usage

```go
// Create a new PurchaseNotification
purchasenotification := PurchaseNotification{
    GuildProductPurchase: any{},
    Type: 42,
}
```

#### Type Definition

```go
type PurchaseNotification struct {
    GuildProductPurchase any `json:"guild_product_purchase,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| GuildProductPurchase | `any` |  |
| Type | `int32` |  |

### PurchaseType
PurchaseType

#### Example Usage

```go
// Example usage of PurchaseType
var value PurchaseType
// Initialize with appropriate value
```

#### Type Definition

```go
type PurchaseType int
```

### QuarantineUserAction
QuarantineUserAction

#### Example Usage

```go
// Create a new QuarantineUserAction
quarantineuseraction := QuarantineUserAction{
    Metadata: any{},
    Type: 42,
}
```

#### Type Definition

```go
type QuarantineUserAction struct {
    Metadata any `json:"metadata"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Metadata | `any` |  |
| Type | `int32` |  |

### QuarantineUserActionMetadata
QuarantineUserActionMetadata

#### Example Usage

```go
// Create a new QuarantineUserActionMetadata
quarantineuseractionmetadata := QuarantineUserActionMetadata{

}
```

#### Type Definition

```go
type QuarantineUserActionMetadata struct {
}
```

### RatelimitedResponse
Ratelimited Ratelimit error object returned by the Discord API

#### Example Usage

```go
// Create a new RatelimitedResponse
ratelimitedresponse := RatelimitedResponse{
    Code: 42,
    Message: "example",
    Global: true,
    RetryAfter: 3.14,
}
```

#### Type Definition

```go
type RatelimitedResponse struct {
    Code int64 `json:"code"`
    Message string `json:"message"`
    Global bool `json:"global"`
    RetryAfter float64 `json:"retry_after"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Code | `int64` | Discord internal error code. See error code reference |
| Message | `string` | Human-readable error message |
| Global | `bool` | Whether you are being ratelimited by the global ratelimit or a per-endpoint ratelimit |
| RetryAfter | `float64` | The number of seconds to wait before retrying your request |

### ReactionTypes
ReactionTypes

#### Example Usage

```go
// Example usage of ReactionTypes
var value ReactionTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type ReactionTypes int
```

### ResolvedObjects
ResolvedObjects

#### Example Usage

```go
// Create a new ResolvedObjects
resolvedobjects := ResolvedObjects{
    Channels: map[],
    Members: map[],
    Roles: map[],
    Users: map[],
}
```

#### Type Definition

```go
type ResolvedObjects struct {
    Channels map[string]any `json:"channels,omitempty"`
    Members map[string]any `json:"members,omitempty"`
    Roles map[string]any `json:"roles,omitempty"`
    Users map[string]any `json:"users,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Channels | `map[string]any` |  |
| Members | `map[string]any` |  |
| Roles | `map[string]any` |  |
| Users | `map[string]any` |  |

### ResourceChannel
ResourceChannel

#### Example Usage

```go
// Create a new ResourceChannel
resourcechannel := ResourceChannel{
    ChannelID: "example",
    Description: "example",
    Emoji: any{},
    Icon: "example",
    Title: "example",
}
```

#### Type Definition

```go
type ResourceChannel struct {
    ChannelID string `json:"channel_id"`
    Description string `json:"description"`
    Emoji any `json:"emoji,omitempty"`
    Icon string `json:"icon,omitempty"`
    Title string `json:"title"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `string` |  |
| Description | `string` |  |
| Emoji | `any` |  |
| Icon | `string` |  |
| Title | `string` |  |

### RichEmbed
RichEmbed

#### Example Usage

```go
// Create a new RichEmbed
richembed := RichEmbed{
    Author: any{},
    Color: &42{},
    Description: &"example"{},
    Fields: [],
    Footer: any{},
    Image: any{},
    Provider: any{},
    Thumbnail: any{},
    Timestamp: &/* value */{},
    Title: &"example"{},
    Type: &"example"{},
    URL: &"example"{},
    Video: any{},
}
```

#### Type Definition

```go
type RichEmbed struct {
    Author any `json:"author,omitempty"`
    Color *int64 `json:"color,omitempty"`
    Description *string `json:"description,omitempty"`
    Fields []any `json:"fields,omitempty"`
    Footer any `json:"footer,omitempty"`
    Image any `json:"image,omitempty"`
    Provider any `json:"provider,omitempty"`
    Thumbnail any `json:"thumbnail,omitempty"`
    Timestamp *time.Time `json:"timestamp,omitempty"`
    Title *string `json:"title,omitempty"`
    Type *string `json:"type,omitempty"`
    URL *string `json:"url,omitempty"`
    Video any `json:"video,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Author | `any` |  |
| Color | `*int64` |  |
| Description | `*string` |  |
| Fields | `[]any` |  |
| Footer | `any` |  |
| Image | `any` |  |
| Provider | `any` |  |
| Thumbnail | `any` |  |
| Timestamp | `*time.Time` |  |
| Title | `*string` |  |
| Type | `*string` |  |
| URL | `*string` |  |
| Video | `any` |  |

### RichEmbedAuthor
RichEmbedAuthor

#### Example Usage

```go
// Create a new RichEmbedAuthor
richembedauthor := RichEmbedAuthor{
    IconURL: &"example"{},
    Name: &"example"{},
    URL: &"example"{},
}
```

#### Type Definition

```go
type RichEmbedAuthor struct {
    IconURL *string `json:"icon_url,omitempty"`
    Name *string `json:"name,omitempty"`
    URL *string `json:"url,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| IconURL | `*string` |  |
| Name | `*string` |  |
| URL | `*string` |  |

### RichEmbedField
RichEmbedField

#### Example Usage

```go
// Create a new RichEmbedField
richembedfield := RichEmbedField{
    Inline: &true{},
    Name: "example",
    Value: "example",
}
```

#### Type Definition

```go
type RichEmbedField struct {
    Inline *bool `json:"inline,omitempty"`
    Name string `json:"name"`
    Value string `json:"value"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Inline | `*bool` |  |
| Name | `string` |  |
| Value | `string` |  |

### RichEmbedFooter
RichEmbedFooter

#### Example Usage

```go
// Create a new RichEmbedFooter
richembedfooter := RichEmbedFooter{
    IconURL: &"example"{},
    Text: &"example"{},
}
```

#### Type Definition

```go
type RichEmbedFooter struct {
    IconURL *string `json:"icon_url,omitempty"`
    Text *string `json:"text,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| IconURL | `*string` |  |
| Text | `*string` |  |

### RichEmbedImage
RichEmbedImage

#### Example Usage

```go
// Create a new RichEmbedImage
richembedimage := RichEmbedImage{
    Description: &"example"{},
    Height: &42{},
    IsAnimated: &true{},
    Placeholder: &"example"{},
    PlaceholderVersion: &42{},
    URL: &"example"{},
    Width: &42{},
}
```

#### Type Definition

```go
type RichEmbedImage struct {
    Description *string `json:"description,omitempty"`
    Height *int64 `json:"height,omitempty"`
    IsAnimated *bool `json:"is_animated,omitempty"`
    Placeholder *string `json:"placeholder,omitempty"`
    PlaceholderVersion *int64 `json:"placeholder_version,omitempty"`
    URL *string `json:"url,omitempty"`
    Width *int64 `json:"width,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `*string` |  |
| Height | `*int64` |  |
| IsAnimated | `*bool` |  |
| Placeholder | `*string` |  |
| PlaceholderVersion | `*int64` |  |
| URL | `*string` |  |
| Width | `*int64` |  |

### RichEmbedProvider
RichEmbedProvider

#### Example Usage

```go
// Create a new RichEmbedProvider
richembedprovider := RichEmbedProvider{
    Name: &"example"{},
    URL: &"example"{},
}
```

#### Type Definition

```go
type RichEmbedProvider struct {
    Name *string `json:"name,omitempty"`
    URL *string `json:"url,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `*string` |  |
| URL | `*string` |  |

### RichEmbedThumbnail
RichEmbedThumbnail

#### Example Usage

```go
// Create a new RichEmbedThumbnail
richembedthumbnail := RichEmbedThumbnail{
    Description: &"example"{},
    Height: &42{},
    IsAnimated: &true{},
    Placeholder: &"example"{},
    PlaceholderVersion: &42{},
    URL: &"example"{},
    Width: &42{},
}
```

#### Type Definition

```go
type RichEmbedThumbnail struct {
    Description *string `json:"description,omitempty"`
    Height *int64 `json:"height,omitempty"`
    IsAnimated *bool `json:"is_animated,omitempty"`
    Placeholder *string `json:"placeholder,omitempty"`
    PlaceholderVersion *int64 `json:"placeholder_version,omitempty"`
    URL *string `json:"url,omitempty"`
    Width *int64 `json:"width,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `*string` |  |
| Height | `*int64` |  |
| IsAnimated | `*bool` |  |
| Placeholder | `*string` |  |
| PlaceholderVersion | `*int64` |  |
| URL | `*string` |  |
| Width | `*int64` |  |

### RichEmbedVideo
RichEmbedVideo

#### Example Usage

```go
// Create a new RichEmbedVideo
richembedvideo := RichEmbedVideo{
    Description: &"example"{},
    Height: &42{},
    IsAnimated: &true{},
    Placeholder: &"example"{},
    PlaceholderVersion: &42{},
    URL: &"example"{},
    Width: &42{},
}
```

#### Type Definition

```go
type RichEmbedVideo struct {
    Description *string `json:"description,omitempty"`
    Height *int64 `json:"height,omitempty"`
    IsAnimated *bool `json:"is_animated,omitempty"`
    Placeholder *string `json:"placeholder,omitempty"`
    PlaceholderVersion *int64 `json:"placeholder_version,omitempty"`
    URL *string `json:"url,omitempty"`
    Width *int64 `json:"width,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `*string` |  |
| Height | `*int64` |  |
| IsAnimated | `*bool` |  |
| Placeholder | `*string` |  |
| PlaceholderVersion | `*int64` |  |
| URL | `*string` |  |
| Width | `*int64` |  |

### RoleSelectComponent
RoleSelectComponent

#### Example Usage

```go
// Create a new RoleSelectComponent
roleselectcomponent := RoleSelectComponent{
    ID: 42,
    CustomID: "example",
    DefaultValues: [],
    Disabled: true,
    MaxValues: &42{},
    MinValues: &42{},
    Placeholder: "example",
    Type: 42,
}
```

#### Type Definition

```go
type RoleSelectComponent struct {
    ID int32 `json:"id"`
    CustomID string `json:"custom_id"`
    DefaultValues []any `json:"default_values,omitempty"`
    Disabled bool `json:"disabled,omitempty"`
    MaxValues *int32 `json:"max_values,omitempty"`
    MinValues *int32 `json:"min_values,omitempty"`
    Placeholder string `json:"placeholder,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| CustomID | `string` |  |
| DefaultValues | `[]any` |  |
| Disabled | `bool` |  |
| MaxValues | `*int32` |  |
| MinValues | `*int32` |  |
| Placeholder | `string` |  |
| Type | `int32` |  |

### RoleSelectComponentForMessageOptions
RoleSelectComponentForMessageOptions

#### Example Usage

```go
// Create a new RoleSelectComponentForMessageOptions
roleselectcomponentformessageoptions := RoleSelectComponentForMessageOptions{
    ID: &42{},
    CustomID: "example",
    DefaultValues: [],
    Disabled: &true{},
    MaxValues: &42{},
    MinValues: &42{},
    Placeholder: &"example"{},
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type RoleSelectComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    CustomID string `json:"custom_id"`
    DefaultValues []any `json:"default_values,omitempty"`
    Disabled *bool `json:"disabled,omitempty"`
    MaxValues *int64 `json:"max_values,omitempty"`
    MinValues *int64 `json:"min_values,omitempty"`
    Placeholder *string `json:"placeholder,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| CustomID | `string` |  |
| DefaultValues | `[]any` |  |
| Disabled | `*bool` |  |
| MaxValues | `*int64` |  |
| MinValues | `*int64` |  |
| Placeholder | `*string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### RoleSelectComponentForModalOptions
RoleSelectComponentForModalOptions

#### Example Usage

```go
// Create a new RoleSelectComponentForModalOptions
roleselectcomponentformodaloptions := RoleSelectComponentForModalOptions{
    ID: &42{},
    CustomID: "example",
    DefaultValues: [],
    Disabled: &true{},
    MaxValues: &42{},
    MinValues: &42{},
    Placeholder: &"example"{},
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type RoleSelectComponentForModalOptions struct {
    ID *int32 `json:"id,omitempty"`
    CustomID string `json:"custom_id"`
    DefaultValues []any `json:"default_values,omitempty"`
    Disabled *bool `json:"disabled,omitempty"`
    MaxValues *int64 `json:"max_values,omitempty"`
    MinValues *int64 `json:"min_values,omitempty"`
    Placeholder *string `json:"placeholder,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| CustomID | `string` |  |
| DefaultValues | `[]any` |  |
| Disabled | `*bool` |  |
| MaxValues | `*int64` |  |
| MinValues | `*int64` |  |
| Placeholder | `*string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### RoleSelectDefaultValue
RoleSelectDefaultValue

#### Example Usage

```go
// Create a new RoleSelectDefaultValue
roleselectdefaultvalue := RoleSelectDefaultValue{
    ID: "example",
    Type: "example",
}
```

#### Type Definition

```go
type RoleSelectDefaultValue struct {
    ID string `json:"id"`
    Type string `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Type | `string` |  |

### SDKMessageOptions
SDKMessageOptions

#### Example Usage

```go
// Create a new SDKMessageOptions
sdkmessageoptions := SDKMessageOptions{
    AllowedMentions: any{},
    Attachments: [],
    Components: [],
    ConfettiPotion: any{},
    Content: &"example"{},
    Embeds: [],
    EnforceNonce: &true{},
    Flags: &42{},
    MessageReference: any{},
    Nonce: any{},
    Poll: any{},
    SharedClientTheme: any{},
    StickerIds: [],
    Tts: &true{},
}
```

#### Type Definition

```go
type SDKMessageOptions struct {
    AllowedMentions any `json:"allowed_mentions,omitempty"`
    Attachments []any `json:"attachments,omitempty"`
    Components []any `json:"components,omitempty"`
    ConfettiPotion any `json:"confetti_potion,omitempty"`
    Content *string `json:"content,omitempty"`
    Embeds []any `json:"embeds,omitempty"`
    EnforceNonce *bool `json:"enforce_nonce,omitempty"`
    Flags *int64 `json:"flags,omitempty"`
    MessageReference any `json:"message_reference,omitempty"`
    Nonce any `json:"nonce,omitempty"`
    Poll any `json:"poll,omitempty"`
    SharedClientTheme any `json:"shared_client_theme,omitempty"`
    StickerIds []string `json:"sticker_ids,omitempty"`
    Tts *bool `json:"tts,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AllowedMentions | `any` |  |
| Attachments | `[]any` |  |
| Components | `[]any` |  |
| ConfettiPotion | `any` |  |
| Content | `*string` |  |
| Embeds | `[]any` |  |
| EnforceNonce | `*bool` |  |
| Flags | `*int64` |  |
| MessageReference | `any` |  |
| Nonce | `any` |  |
| Poll | `any` |  |
| SharedClientTheme | `any` |  |
| StickerIds | `[]string` |  |
| Tts | `*bool` |  |

### ScheduledEvent
ScheduledEvent

#### Example Usage

```go
// Create a new ScheduledEvent
scheduledevent := ScheduledEvent{
    ID: "example",
    ChannelID: any{},
    Creator: any{},
    CreatorID: any{},
    Description: &"example"{},
    EntityID: any{},
    EntityType: 42,
    GuildID: "example",
    Image: &"example"{},
    Name: "example",
    PrivacyLevel: 42,
    ScheduledEndTime: &/* value */{},
    ScheduledStartTime: /* value */,
    Status: 42,
    UserCount: 42,
    UserRsvp: any{},
}
```

#### Type Definition

```go
type ScheduledEvent struct {
    ID string `json:"id"`
    ChannelID any `json:"channel_id,omitempty"`
    Creator any `json:"creator,omitempty"`
    CreatorID any `json:"creator_id,omitempty"`
    Description *string `json:"description,omitempty"`
    EntityID any `json:"entity_id,omitempty"`
    EntityType int32 `json:"entity_type"`
    GuildID string `json:"guild_id"`
    Image *string `json:"image,omitempty"`
    Name string `json:"name"`
    PrivacyLevel int32 `json:"privacy_level"`
    ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
    ScheduledStartTime time.Time `json:"scheduled_start_time"`
    Status int32 `json:"status"`
    UserCount int32 `json:"user_count,omitempty"`
    UserRsvp any `json:"user_rsvp,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ChannelID | `any` |  |
| Creator | `any` |  |
| CreatorID | `any` |  |
| Description | `*string` |  |
| EntityID | `any` |  |
| EntityType | `int32` |  |
| GuildID | `string` |  |
| Image | `*string` |  |
| Name | `string` |  |
| PrivacyLevel | `int32` |  |
| ScheduledEndTime | `*time.Time` |  |
| ScheduledStartTime | `time.Time` |  |
| Status | `int32` |  |
| UserCount | `int32` |  |
| UserRsvp | `any` |  |

### ScheduledEventUser
ScheduledEventUser

#### Example Usage

```go
// Create a new ScheduledEventUser
scheduledeventuser := ScheduledEventUser{
    GuildScheduledEventID: "example",
    Member: any{},
    User: any{},
    UserID: "example",
}
```

#### Type Definition

```go
type ScheduledEventUser struct {
    GuildScheduledEventID string `json:"guild_scheduled_event_id"`
    Member any `json:"member,omitempty"`
    User any `json:"user,omitempty"`
    UserID string `json:"user_id"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| GuildScheduledEventID | `string` |  |
| Member | `any` |  |
| User | `any` |  |
| UserID | `string` |  |

### SectionComponent
SectionComponent

#### Example Usage

```go
// Create a new SectionComponent
sectioncomponent := SectionComponent{
    ID: 42,
    Accessory: any{},
    Components: [],
    Type: 42,
}
```

#### Type Definition

```go
type SectionComponent struct {
    ID int32 `json:"id"`
    Accessory any `json:"accessory"`
    Components []any `json:"components"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| Accessory | `any` |  |
| Components | `[]any` |  |
| Type | `int32` |  |

### SectionComponentForMessageOptions
SectionComponentForMessageOptions

#### Example Usage

```go
// Create a new SectionComponentForMessageOptions
sectioncomponentformessageoptions := SectionComponentForMessageOptions{
    ID: &42{},
    Accessory: any{},
    Components: [],
    Type: 42,
}
```

#### Type Definition

```go
type SectionComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    Accessory any `json:"accessory"`
    Components []any `json:"components"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| Accessory | `any` |  |
| Components | `[]any` |  |
| Type | `int32` |  |

### SeparatorComponent
SeparatorComponent

#### Example Usage

```go
// Create a new SeparatorComponent
separatorcomponent := SeparatorComponent{
    ID: 42,
    Divider: true,
    Spacing: 42,
    Type: 42,
}
```

#### Type Definition

```go
type SeparatorComponent struct {
    ID int32 `json:"id"`
    Divider bool `json:"divider"`
    Spacing int32 `json:"spacing"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| Divider | `bool` |  |
| Spacing | `int32` |  |
| Type | `int32` |  |

### SeparatorComponentForMessageOptions
SeparatorComponentForMessageOptions

#### Example Usage

```go
// Create a new SeparatorComponentForMessageOptions
separatorcomponentformessageoptions := SeparatorComponentForMessageOptions{
    ID: &42{},
    Divider: &true{},
    Spacing: any{},
    Type: 42,
}
```

#### Type Definition

```go
type SeparatorComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    Divider *bool `json:"divider,omitempty"`
    Spacing any `json:"spacing,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| Divider | `*bool` |  |
| Spacing | `any` |  |
| Type | `int32` |  |

### SettingsEmoji
SettingsEmoji

#### Example Usage

```go
// Create a new SettingsEmoji
settingsemoji := SettingsEmoji{
    ID: any{},
    Animated: true,
    Name: &"example"{},
}
```

#### Type Definition

```go
type SettingsEmoji struct {
    ID any `json:"id,omitempty"`
    Animated bool `json:"animated"`
    Name *string `json:"name,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `any` |  |
| Animated | `bool` |  |
| Name | `*string` |  |

### SlackWebhook
SlackWebhook

#### Example Usage

```go
// Create a new SlackWebhook
slackwebhook := SlackWebhook{
    Attachments: [],
    IconURL: &"example"{},
    Text: &"example"{},
    Username: &"example"{},
}
```

#### Type Definition

```go
type SlackWebhook struct {
    Attachments []any `json:"attachments,omitempty"`
    IconURL *string `json:"icon_url,omitempty"`
    Text *string `json:"text,omitempty"`
    Username *string `json:"username,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Attachments | `[]any` |  |
| IconURL | `*string` |  |
| Text | `*string` |  |
| Username | `*string` |  |

### SnowflakeSelectDefaultValueTypes
SnowflakeSelectDefaultValueTypes

#### Example Usage

```go
// Example usage of SnowflakeSelectDefaultValueTypes
var value SnowflakeSelectDefaultValueTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type SnowflakeSelectDefaultValueTypes string
```

### SnowflakeType
SnowflakeType is a Discord snowflake ID (just a string).

#### Example Usage

```go
// Example usage of SnowflakeType
var value SnowflakeType
// Initialize with appropriate value
```

#### Type Definition

```go
type SnowflakeType string
```

### SortingOrder
SortingOrder

#### Example Usage

```go
// Example usage of SortingOrder
var value SortingOrder
// Initialize with appropriate value
```

#### Type Definition

```go
type SortingOrder string
```

### SoundboardCreateOptions
SoundboardCreateOptions

#### Example Usage

```go
// Create a new SoundboardCreateOptions
soundboardcreateoptions := SoundboardCreateOptions{
    EmojiID: any{},
    EmojiName: &"example"{},
    Name: "example",
    Sound: "example",
    Volume: &3.14{},
}
```

#### Type Definition

```go
type SoundboardCreateOptions struct {
    EmojiID any `json:"emoji_id,omitempty"`
    EmojiName *string `json:"emoji_name,omitempty"`
    Name string `json:"name"`
    Sound string `json:"sound"`
    Volume *float64 `json:"volume,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| EmojiID | `any` |  |
| EmojiName | `*string` |  |
| Name | `string` |  |
| Sound | `string` |  |
| Volume | `*float64` |  |

### SoundboardPatchOptionsPartial
SoundboardPatchOptionsPartial

#### Example Usage

```go
// Create a new SoundboardPatchOptionsPartial
soundboardpatchoptionspartial := SoundboardPatchOptionsPartial{
    EmojiID: any{},
    EmojiName: &"example"{},
    Name: "example",
    Volume: &3.14{},
}
```

#### Type Definition

```go
type SoundboardPatchOptionsPartial struct {
    EmojiID any `json:"emoji_id,omitempty"`
    EmojiName *string `json:"emoji_name,omitempty"`
    Name string `json:"name,omitempty"`
    Volume *float64 `json:"volume,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| EmojiID | `any` |  |
| EmojiName | `*string` |  |
| Name | `string` |  |
| Volume | `*float64` |  |

### SoundboardSound
SoundboardSound

#### Example Usage

```go
// Create a new SoundboardSound
soundboardsound := SoundboardSound{
    Available: true,
    EmojiID: any{},
    EmojiName: &"example"{},
    GuildID: "example",
    Name: "example",
    SoundID: "example",
    User: any{},
    Volume: 3.14,
}
```

#### Type Definition

```go
type SoundboardSound struct {
    Available bool `json:"available"`
    EmojiID any `json:"emoji_id,omitempty"`
    EmojiName *string `json:"emoji_name,omitempty"`
    GuildID string `json:"guild_id,omitempty"`
    Name string `json:"name"`
    SoundID string `json:"sound_id"`
    User any `json:"user,omitempty"`
    Volume float64 `json:"volume"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Available | `bool` |  |
| EmojiID | `any` |  |
| EmojiName | `*string` |  |
| GuildID | `string` |  |
| Name | `string` |  |
| SoundID | `string` |  |
| User | `any` |  |
| Volume | `float64` |  |

### SoundboardSoundSendOptions
SoundboardSoundSendOptions

#### Example Usage

```go
// Create a new SoundboardSoundSendOptions
soundboardsoundsendoptions := SoundboardSoundSendOptions{
    SoundID: "example",
    SourceGuildID: any{},
}
```

#### Type Definition

```go
type SoundboardSoundSendOptions struct {
    SoundID string `json:"sound_id"`
    SourceGuildID any `json:"source_guild_id,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| SoundID | `string` |  |
| SourceGuildID | `any` |  |

### SpamLinkRule
SpamLinkRule

#### Example Usage

```go
// Create a new SpamLinkRule
spamlinkrule := SpamLinkRule{
    ID: "example",
    Actions: [],
    CreatorID: "example",
    Enabled: true,
    EventType: 42,
    ExemptChannels: [],
    ExemptRoles: [],
    GuildID: "example",
    Name: "example",
    TriggerMetadata: any{},
    TriggerType: 42,
}
```

#### Type Definition

```go
type SpamLinkRule struct {
    ID string `json:"id"`
    Actions []any `json:"actions"`
    CreatorID string `json:"creator_id"`
    Enabled bool `json:"enabled"`
    EventType int32 `json:"event_type"`
    ExemptChannels []string `json:"exempt_channels"`
    ExemptRoles []string `json:"exempt_roles"`
    GuildID string `json:"guild_id"`
    Name string `json:"name"`
    TriggerMetadata any `json:"trigger_metadata"`
    TriggerType int32 `json:"trigger_type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Actions | `[]any` |  |
| CreatorID | `string` |  |
| Enabled | `bool` |  |
| EventType | `int32` |  |
| ExemptChannels | `[]string` |  |
| ExemptRoles | `[]string` |  |
| GuildID | `string` |  |
| Name | `string` |  |
| TriggerMetadata | `any` |  |
| TriggerType | `int32` |  |

### SpamLinkTriggerMetadata
SpamLinkTriggerMetadata

#### Example Usage

```go
// Create a new SpamLinkTriggerMetadata
spamlinktriggermetadata := SpamLinkTriggerMetadata{

}
```

#### Type Definition

```go
type SpamLinkTriggerMetadata struct {
}
```

### StageInstance
StageInstance

#### Example Usage

```go
// Create a new StageInstance
stageinstance := StageInstance{
    ID: "example",
    ChannelID: "example",
    DiscoverableDisabled: true,
    GuildID: "example",
    GuildScheduledEventID: any{},
    PrivacyLevel: 42,
    Topic: "example",
}
```

#### Type Definition

```go
type StageInstance struct {
    ID string `json:"id"`
    ChannelID string `json:"channel_id"`
    DiscoverableDisabled bool `json:"discoverable_disabled"`
    GuildID string `json:"guild_id"`
    GuildScheduledEventID any `json:"guild_scheduled_event_id,omitempty"`
    PrivacyLevel int32 `json:"privacy_level"`
    Topic string `json:"topic"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ChannelID | `string` |  |
| DiscoverableDisabled | `bool` |  |
| GuildID | `string` |  |
| GuildScheduledEventID | `any` |  |
| PrivacyLevel | `int32` |  |
| Topic | `string` |  |

### StageInstancesPrivacyLevels
StageInstancesPrivacyLevels

#### Example Usage

```go
// Example usage of StageInstancesPrivacyLevels
var value StageInstancesPrivacyLevels
// Initialize with appropriate value
```

#### Type Definition

```go
type StageInstancesPrivacyLevels int
```

### StageScheduledEvent
StageScheduledEvent

#### Example Usage

```go
// Create a new StageScheduledEvent
stagescheduledevent := StageScheduledEvent{
    ID: "example",
    ChannelID: any{},
    Creator: any{},
    CreatorID: any{},
    Description: &"example"{},
    EntityID: any{},
    EntityMetadata: any{},
    EntityType: 42,
    GuildID: "example",
    Image: &"example"{},
    Name: "example",
    PrivacyLevel: 42,
    ScheduledEndTime: &/* value */{},
    ScheduledStartTime: /* value */,
    Status: 42,
    UserCount: 42,
    UserRsvp: any{},
}
```

#### Type Definition

```go
type StageScheduledEvent struct {
    ID string `json:"id"`
    ChannelID any `json:"channel_id,omitempty"`
    Creator any `json:"creator,omitempty"`
    CreatorID any `json:"creator_id,omitempty"`
    Description *string `json:"description,omitempty"`
    EntityID any `json:"entity_id,omitempty"`
    EntityMetadata any `json:"entity_metadata,omitempty"`
    EntityType int32 `json:"entity_type"`
    GuildID string `json:"guild_id"`
    Image *string `json:"image,omitempty"`
    Name string `json:"name"`
    PrivacyLevel int32 `json:"privacy_level"`
    ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
    ScheduledStartTime time.Time `json:"scheduled_start_time"`
    Status int32 `json:"status"`
    UserCount int32 `json:"user_count,omitempty"`
    UserRsvp any `json:"user_rsvp,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ChannelID | `any` |  |
| Creator | `any` |  |
| CreatorID | `any` |  |
| Description | `*string` |  |
| EntityID | `any` |  |
| EntityMetadata | `any` |  |
| EntityType | `int32` |  |
| GuildID | `string` |  |
| Image | `*string` |  |
| Name | `string` |  |
| PrivacyLevel | `int32` |  |
| ScheduledEndTime | `*time.Time` |  |
| ScheduledStartTime | `time.Time` |  |
| Status | `int32` |  |
| UserCount | `int32` |  |
| UserRsvp | `any` |  |

### StageScheduledEventCreateOptions
StageScheduledEventCreateOptions

#### Example Usage

```go
// Create a new StageScheduledEventCreateOptions
stagescheduledeventcreateoptions := StageScheduledEventCreateOptions{
    ChannelID: any{},
    Description: &"example"{},
    EntityMetadata: any{},
    EntityType: 42,
    Image: &"example"{},
    Name: "example",
    PrivacyLevel: 42,
    ScheduledEndTime: &/* value */{},
    ScheduledStartTime: /* value */,
}
```

#### Type Definition

```go
type StageScheduledEventCreateOptions struct {
    ChannelID any `json:"channel_id,omitempty"`
    Description *string `json:"description,omitempty"`
    EntityMetadata any `json:"entity_metadata,omitempty"`
    EntityType int32 `json:"entity_type"`
    Image *string `json:"image,omitempty"`
    Name string `json:"name"`
    PrivacyLevel int32 `json:"privacy_level"`
    ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
    ScheduledStartTime time.Time `json:"scheduled_start_time"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `any` |  |
| Description | `*string` |  |
| EntityMetadata | `any` |  |
| EntityType | `int32` |  |
| Image | `*string` |  |
| Name | `string` |  |
| PrivacyLevel | `int32` |  |
| ScheduledEndTime | `*time.Time` |  |
| ScheduledStartTime | `time.Time` |  |

### StageScheduledEventPatchOptionsPartial
StageScheduledEventPatchOptionsPartial

#### Example Usage

```go
// Create a new StageScheduledEventPatchOptionsPartial
stagescheduledeventpatchoptionspartial := StageScheduledEventPatchOptionsPartial{
    ChannelID: any{},
    Description: &"example"{},
    EntityMetadata: any{},
    EntityType: any{},
    Image: &"example"{},
    Name: "example",
    PrivacyLevel: 42,
    ScheduledEndTime: &/* value */{},
    ScheduledStartTime: /* value */,
    Status: any{},
}
```

#### Type Definition

```go
type StageScheduledEventPatchOptionsPartial struct {
    ChannelID any `json:"channel_id,omitempty"`
    Description *string `json:"description,omitempty"`
    EntityMetadata any `json:"entity_metadata,omitempty"`
    EntityType any `json:"entity_type,omitempty"`
    Image *string `json:"image,omitempty"`
    Name string `json:"name,omitempty"`
    PrivacyLevel int32 `json:"privacy_level,omitempty"`
    ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
    ScheduledStartTime time.Time `json:"scheduled_start_time,omitempty"`
    Status any `json:"status,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `any` |  |
| Description | `*string` |  |
| EntityMetadata | `any` |  |
| EntityType | `any` |  |
| Image | `*string` |  |
| Name | `string` |  |
| PrivacyLevel | `int32` |  |
| ScheduledEndTime | `*time.Time` |  |
| ScheduledStartTime | `time.Time` |  |
| Status | `any` |  |

### StandardSticker
StandardSticker

#### Example Usage

```go
// Create a new StandardSticker
standardsticker := StandardSticker{
    ID: "example",
    Description: &"example"{},
    FormatType: any{},
    Name: "example",
    PackID: "example",
    SortValue: 42,
    Tags: "example",
    Type: 42,
}
```

#### Type Definition

```go
type StandardSticker struct {
    ID string `json:"id"`
    Description *string `json:"description,omitempty"`
    FormatType any `json:"format_type,omitempty"`
    Name string `json:"name"`
    PackID string `json:"pack_id"`
    SortValue int32 `json:"sort_value"`
    Tags string `json:"tags"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Description | `*string` |  |
| FormatType | `any` |  |
| Name | `string` |  |
| PackID | `string` |  |
| SortValue | `int32` |  |
| Tags | `string` |  |
| Type | `int32` |  |

### StickerFormatTypes
StickerFormatTypes

#### Example Usage

```go
// Example usage of StickerFormatTypes
var value StickerFormatTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type StickerFormatTypes int
```

### StickerPack
StickerPack

#### Example Usage

```go
// Create a new StickerPack
stickerpack := StickerPack{
    ID: "example",
    BannerAssetID: "example",
    CoverStickerID: "example",
    Description: &"example"{},
    Name: "example",
    SkuID: "example",
    Stickers: [],
}
```

#### Type Definition

```go
type StickerPack struct {
    ID string `json:"id"`
    BannerAssetID string `json:"banner_asset_id,omitempty"`
    CoverStickerID string `json:"cover_sticker_id,omitempty"`
    Description *string `json:"description,omitempty"`
    Name string `json:"name"`
    SkuID string `json:"sku_id"`
    Stickers []any `json:"stickers"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| BannerAssetID | `string` |  |
| CoverStickerID | `string` |  |
| Description | `*string` |  |
| Name | `string` |  |
| SkuID | `string` |  |
| Stickers | `[]any` |  |

### StickerPackCollection
StickerPackCollection

#### Example Usage

```go
// Create a new StickerPackCollection
stickerpackcollection := StickerPackCollection{
    StickerPacks: [],
}
```

#### Type Definition

```go
type StickerPackCollection struct {
    StickerPacks []any `json:"sticker_packs"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| StickerPacks | `[]any` |  |

### StickerTypes
StickerTypes

#### Example Usage

```go
// Example usage of StickerTypes
var value StickerTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type StickerTypes int
```

### StringSelectComponent
StringSelectComponent

#### Example Usage

```go
// Create a new StringSelectComponent
stringselectcomponent := StringSelectComponent{
    ID: 42,
    CustomID: "example",
    Disabled: true,
    MaxValues: &42{},
    MinValues: &42{},
    Options: [],
    Placeholder: "example",
    Type: 42,
}
```

#### Type Definition

```go
type StringSelectComponent struct {
    ID int32 `json:"id"`
    CustomID string `json:"custom_id"`
    Disabled bool `json:"disabled,omitempty"`
    MaxValues *int32 `json:"max_values,omitempty"`
    MinValues *int32 `json:"min_values,omitempty"`
    Options []any `json:"options"`
    Placeholder string `json:"placeholder,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| CustomID | `string` |  |
| Disabled | `bool` |  |
| MaxValues | `*int32` |  |
| MinValues | `*int32` |  |
| Options | `[]any` |  |
| Placeholder | `string` |  |
| Type | `int32` |  |

### StringSelectComponentForMessageOptions
StringSelectComponentForMessageOptions

#### Example Usage

```go
// Create a new StringSelectComponentForMessageOptions
stringselectcomponentformessageoptions := StringSelectComponentForMessageOptions{
    ID: &42{},
    CustomID: "example",
    Disabled: &true{},
    MaxValues: &42{},
    MinValues: &42{},
    Options: [],
    Placeholder: &"example"{},
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type StringSelectComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    CustomID string `json:"custom_id"`
    Disabled *bool `json:"disabled,omitempty"`
    MaxValues *int64 `json:"max_values,omitempty"`
    MinValues *int64 `json:"min_values,omitempty"`
    Options []any `json:"options"`
    Placeholder *string `json:"placeholder,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| CustomID | `string` |  |
| Disabled | `*bool` |  |
| MaxValues | `*int64` |  |
| MinValues | `*int64` |  |
| Options | `[]any` |  |
| Placeholder | `*string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### StringSelectComponentForModalOptions
StringSelectComponentForModalOptions

#### Example Usage

```go
// Create a new StringSelectComponentForModalOptions
stringselectcomponentformodaloptions := StringSelectComponentForModalOptions{
    ID: &42{},
    CustomID: "example",
    Disabled: &true{},
    MaxValues: &42{},
    MinValues: &42{},
    Options: [],
    Placeholder: &"example"{},
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type StringSelectComponentForModalOptions struct {
    ID *int32 `json:"id,omitempty"`
    CustomID string `json:"custom_id"`
    Disabled *bool `json:"disabled,omitempty"`
    MaxValues *int64 `json:"max_values,omitempty"`
    MinValues *int64 `json:"min_values,omitempty"`
    Options []any `json:"options"`
    Placeholder *string `json:"placeholder,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| CustomID | `string` |  |
| Disabled | `*bool` |  |
| MaxValues | `*int64` |  |
| MinValues | `*int64` |  |
| Options | `[]any` |  |
| Placeholder | `*string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### StringSelectOption
StringSelectOption

#### Example Usage

```go
// Create a new StringSelectOption
stringselectoption := StringSelectOption{
    Default: true,
    Description: "example",
    Emoji: any{},
    Label: "example",
    Value: "example",
}
```

#### Type Definition

```go
type StringSelectOption struct {
    Default bool `json:"default,omitempty"`
    Description string `json:"description,omitempty"`
    Emoji any `json:"emoji,omitempty"`
    Label string `json:"label"`
    Value string `json:"value"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Default | `bool` |  |
| Description | `string` |  |
| Emoji | `any` |  |
| Label | `string` |  |
| Value | `string` |  |

### StringSelectOptionForOptions
StringSelectOptionForOptions

#### Example Usage

```go
// Create a new StringSelectOptionForOptions
stringselectoptionforoptions := StringSelectOptionForOptions{
    Default: &true{},
    Description: &"example"{},
    Emoji: any{},
    Label: "example",
    Value: "example",
}
```

#### Type Definition

```go
type StringSelectOptionForOptions struct {
    Default *bool `json:"default,omitempty"`
    Description *string `json:"description,omitempty"`
    Emoji any `json:"emoji,omitempty"`
    Label string `json:"label"`
    Value string `json:"value"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Default | `*bool` |  |
| Description | `*string` |  |
| Emoji | `any` |  |
| Label | `string` |  |
| Value | `string` |  |

### Team
Team

#### Example Usage

```go
// Create a new Team
team := Team{
    ID: "example",
    Icon: &"example"{},
    Members: [],
    Name: "example",
    OwnerUserID: "example",
}
```

#### Type Definition

```go
type Team struct {
    ID string `json:"id"`
    Icon *string `json:"icon,omitempty"`
    Members []any `json:"members"`
    Name string `json:"name"`
    OwnerUserID string `json:"owner_user_id"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Icon | `*string` |  |
| Members | `[]any` |  |
| Name | `string` |  |
| OwnerUserID | `string` |  |

### TeamMember
TeamMember

#### Example Usage

```go
// Create a new TeamMember
teammember := TeamMember{
    MembershipState: 42,
    TeamID: "example",
    User: any{},
}
```

#### Type Definition

```go
type TeamMember struct {
    MembershipState int32 `json:"membership_state"`
    TeamID string `json:"team_id"`
    User any `json:"user"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| MembershipState | `int32` |  |
| TeamID | `string` |  |
| User | `any` |  |

### TeamMembershipStates
TeamMembershipStates

#### Example Usage

```go
// Example usage of TeamMembershipStates
var value TeamMembershipStates
// Initialize with appropriate value
```

#### Type Definition

```go
type TeamMembershipStates int
```

### TextDisplayComponent
TextDisplayComponent

#### Example Usage

```go
// Create a new TextDisplayComponent
textdisplaycomponent := TextDisplayComponent{
    ID: 42,
    Content: "example",
    Type: 42,
}
```

#### Type Definition

```go
type TextDisplayComponent struct {
    ID int32 `json:"id"`
    Content string `json:"content"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| Content | `string` |  |
| Type | `int32` |  |

### TextDisplayComponentForMessageOptions
TextDisplayComponentForMessageOptions

#### Example Usage

```go
// Create a new TextDisplayComponentForMessageOptions
textdisplaycomponentformessageoptions := TextDisplayComponentForMessageOptions{
    ID: &42{},
    Content: "example",
    Type: 42,
}
```

#### Type Definition

```go
type TextDisplayComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    Content string `json:"content"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| Content | `string` |  |
| Type | `int32` |  |

### TextDisplayComponentForModalOptions
TextDisplayComponentForModalOptions

#### Example Usage

```go
// Create a new TextDisplayComponentForModalOptions
textdisplaycomponentformodaloptions := TextDisplayComponentForModalOptions{
    ID: &42{},
    Content: "example",
    Type: 42,
}
```

#### Type Definition

```go
type TextDisplayComponentForModalOptions struct {
    ID *int32 `json:"id,omitempty"`
    Content string `json:"content"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| Content | `string` |  |
| Type | `int32` |  |

### TextInputComponent
TextInputComponent

#### Example Usage

```go
// Create a new TextInputComponent
textinputcomponent := TextInputComponent{
    ID: 42,
    CustomID: "example",
    Label: &"example"{},
    MaxLength: &42{},
    MinLength: &42{},
    Placeholder: "example",
    Required: true,
    Style: 42,
    Type: 42,
    Value: "example",
}
```

#### Type Definition

```go
type TextInputComponent struct {
    ID int32 `json:"id"`
    CustomID string `json:"custom_id"`
    Label *string `json:"label,omitempty"`
    MaxLength *int32 `json:"max_length,omitempty"`
    MinLength *int32 `json:"min_length,omitempty"`
    Placeholder string `json:"placeholder,omitempty"`
    Required bool `json:"required,omitempty"`
    Style int32 `json:"style"`
    Type int32 `json:"type"`
    Value string `json:"value,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| CustomID | `string` |  |
| Label | `*string` |  |
| MaxLength | `*int32` |  |
| MinLength | `*int32` |  |
| Placeholder | `string` |  |
| Required | `bool` |  |
| Style | `int32` |  |
| Type | `int32` |  |
| Value | `string` |  |

### TextInputComponentForModalOptions
TextInputComponentForModalOptions

#### Example Usage

```go
// Create a new TextInputComponentForModalOptions
textinputcomponentformodaloptions := TextInputComponentForModalOptions{
    ID: &42{},
    CustomID: "example",
    Label: &"example"{},
    MaxLength: &42{},
    MinLength: &42{},
    Placeholder: &"example"{},
    Required: &true{},
    Style: 42,
    Type: 42,
    Value: &"example"{},
}
```

#### Type Definition

```go
type TextInputComponentForModalOptions struct {
    ID *int32 `json:"id,omitempty"`
    CustomID string `json:"custom_id"`
    Label *string `json:"label,omitempty"`
    MaxLength *int64 `json:"max_length,omitempty"`
    MinLength *int64 `json:"min_length,omitempty"`
    Placeholder *string `json:"placeholder,omitempty"`
    Required *bool `json:"required,omitempty"`
    Style int32 `json:"style"`
    Type int32 `json:"type"`
    Value *string `json:"value,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| CustomID | `string` |  |
| Label | `*string` |  |
| MaxLength | `*int64` |  |
| MinLength | `*int64` |  |
| Placeholder | `*string` |  |
| Required | `*bool` |  |
| Style | `int32` |  |
| Type | `int32` |  |
| Value | `*string` |  |

### TextInputStyleTypes
TextInputStyleTypes

#### Example Usage

```go
// Example usage of TextInputStyleTypes
var value TextInputStyleTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type TextInputStyleTypes int
```

### Thread
Thread

#### Example Usage

```go
// Create a new Thread
thread := Thread{
    ID: "example",
    AppliedTags: [],
    Bitrate: 42,
    Flags: 42,
    GuildID: "example",
    LastMessageID: any{},
    LastPinTimestamp: &/* value */{},
    Member: any{},
    MemberCount: 42,
    MessageCount: 42,
    Name: "example",
    OwnerID: "example",
    ParentID: any{},
    Permissions: &"example"{},
    RateLimitPerUser: 42,
    RtcRegion: &"example"{},
    ThreadMetadata: any{},
    TotalMessageSent: 42,
    Type: 42,
    UserLimit: 42,
    VideoQualityMode: 42,
}
```

#### Type Definition

```go
type Thread struct {
    ID string `json:"id"`
    AppliedTags []string `json:"applied_tags,omitempty"`
    Bitrate int32 `json:"bitrate,omitempty"`
    Flags int32 `json:"flags"`
    GuildID string `json:"guild_id"`
    LastMessageID any `json:"last_message_id,omitempty"`
    LastPinTimestamp *time.Time `json:"last_pin_timestamp,omitempty"`
    Member any `json:"member,omitempty"`
    MemberCount int32 `json:"member_count"`
    MessageCount int32 `json:"message_count"`
    Name string `json:"name"`
    OwnerID string `json:"owner_id"`
    ParentID any `json:"parent_id,omitempty"`
    Permissions *string `json:"permissions,omitempty"`
    RateLimitPerUser int32 `json:"rate_limit_per_user,omitempty"`
    RtcRegion *string `json:"rtc_region,omitempty"`
    ThreadMetadata any `json:"thread_metadata"`
    TotalMessageSent int32 `json:"total_message_sent"`
    Type int32 `json:"type"`
    UserLimit int32 `json:"user_limit,omitempty"`
    VideoQualityMode int32 `json:"video_quality_mode,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| AppliedTags | `[]string` |  |
| Bitrate | `int32` |  |
| Flags | `int32` |  |
| GuildID | `string` |  |
| LastMessageID | `any` |  |
| LastPinTimestamp | `*time.Time` |  |
| Member | `any` |  |
| MemberCount | `int32` |  |
| MessageCount | `int32` |  |
| Name | `string` |  |
| OwnerID | `string` |  |
| ParentID | `any` |  |
| Permissions | `*string` |  |
| RateLimitPerUser | `int32` |  |
| RtcRegion | `*string` |  |
| ThreadMetadata | `any` |  |
| TotalMessageSent | `int32` |  |
| Type | `int32` |  |
| UserLimit | `int32` |  |
| VideoQualityMode | `int32` |  |

### ThreadAutoArchiveDuration
ThreadAutoArchiveDuration

#### Example Usage

```go
// Example usage of ThreadAutoArchiveDuration
var value ThreadAutoArchiveDuration
// Initialize with appropriate value
```

#### Type Definition

```go
type ThreadAutoArchiveDuration int
```

### ThreadMember
ThreadMember

#### Example Usage

```go
// Create a new ThreadMember
threadmember := ThreadMember{
    ID: "example",
    Flags: 42,
    JoinTimestamp: /* value */,
    Member: any{},
    UserID: "example",
}
```

#### Type Definition

```go
type ThreadMember struct {
    ID string `json:"id"`
    Flags int32 `json:"flags"`
    JoinTimestamp time.Time `json:"join_timestamp"`
    Member any `json:"member,omitempty"`
    UserID string `json:"user_id"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Flags | `int32` |  |
| JoinTimestamp | `time.Time` |  |
| Member | `any` |  |
| UserID | `string` |  |

### ThreadMetadata
ThreadMetadata

#### Example Usage

```go
// Create a new ThreadMetadata
threadmetadata := ThreadMetadata{
    ArchiveTimestamp: &/* value */{},
    Archived: true,
    AutoArchiveDuration: 42,
    CreateTimestamp: /* value */,
    Invitable: true,
    Locked: true,
}
```

#### Type Definition

```go
type ThreadMetadata struct {
    ArchiveTimestamp *time.Time `json:"archive_timestamp,omitempty"`
    Archived bool `json:"archived"`
    AutoArchiveDuration int32 `json:"auto_archive_duration"`
    CreateTimestamp time.Time `json:"create_timestamp,omitempty"`
    Invitable bool `json:"invitable,omitempty"`
    Locked bool `json:"locked"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ArchiveTimestamp | `*time.Time` |  |
| Archived | `bool` |  |
| AutoArchiveDuration | `int32` |  |
| CreateTimestamp | `time.Time` |  |
| Invitable | `bool` |  |
| Locked | `bool` |  |

### ThreadSearch
ThreadSearch

#### Example Usage

```go
// Create a new ThreadSearch
threadsearch := ThreadSearch{
    FirstMessages: [],
    HasMore: true,
    Members: [],
    Threads: [],
    TotalResults: 42,
}
```

#### Type Definition

```go
type ThreadSearch struct {
    FirstMessages []any `json:"first_messages,omitempty"`
    HasMore bool `json:"has_more"`
    Members []any `json:"members"`
    Threads []any `json:"threads"`
    TotalResults int32 `json:"total_results"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| FirstMessages | `[]any` |  |
| HasMore | `bool` |  |
| Members | `[]any` |  |
| Threads | `[]any` |  |
| TotalResults | `int32` |  |

### ThreadSearchTagSetting
ThreadSearchTagSetting

#### Example Usage

```go
// Example usage of ThreadSearchTagSetting
var value ThreadSearchTagSetting
// Initialize with appropriate value
```

#### Type Definition

```go
type ThreadSearchTagSetting string
```

### ThreadSortOrder
ThreadSortOrder

#### Example Usage

```go
// Example usage of ThreadSortOrder
var value ThreadSortOrder
// Initialize with appropriate value
```

#### Type Definition

```go
type ThreadSortOrder int
```

### ThreadSortingMode
ThreadSortingMode

#### Example Usage

```go
// Example usage of ThreadSortingMode
var value ThreadSortingMode
// Initialize with appropriate value
```

#### Type Definition

```go
type ThreadSortingMode string
```

### Threads
Threads

#### Example Usage

```go
// Create a new Threads
threads := Threads{
    FirstMessages: [],
    HasMore: true,
    Members: [],
    Threads: [],
}
```

#### Type Definition

```go
type Threads struct {
    FirstMessages []any `json:"first_messages,omitempty"`
    HasMore bool `json:"has_more"`
    Members []any `json:"members"`
    Threads []any `json:"threads"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| FirstMessages | `[]any` |  |
| HasMore | `bool` |  |
| Members | `[]any` |  |
| Threads | `[]any` |  |

### ThumbnailComponent
ThumbnailComponent

#### Example Usage

```go
// Create a new ThumbnailComponent
thumbnailcomponent := ThumbnailComponent{
    ID: 42,
    Description: &"example"{},
    Media: any{},
    Spoiler: true,
    Type: 42,
}
```

#### Type Definition

```go
type ThumbnailComponent struct {
    ID int32 `json:"id"`
    Description *string `json:"description,omitempty"`
    Media any `json:"media"`
    Spoiler bool `json:"spoiler"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| Description | `*string` |  |
| Media | `any` |  |
| Spoiler | `bool` |  |
| Type | `int32` |  |

### ThumbnailComponentForMessageOptions
ThumbnailComponentForMessageOptions

#### Example Usage

```go
// Create a new ThumbnailComponentForMessageOptions
thumbnailcomponentformessageoptions := ThumbnailComponentForMessageOptions{
    ID: &42{},
    Description: &"example"{},
    Media: any{},
    Spoiler: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type ThumbnailComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    Description *string `json:"description,omitempty"`
    Media any `json:"media"`
    Spoiler *bool `json:"spoiler,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| Description | `*string` |  |
| Media | `any` |  |
| Spoiler | `*bool` |  |
| Type | `int32` |  |

### TypingIndicator
TypingIndicator

#### Example Usage

```go
// Create a new TypingIndicator
typingindicator := TypingIndicator{

}
```

#### Type Definition

```go
type TypingIndicator struct {
}
```

### UInt32Type
UInt32Type

#### Example Usage

```go
// Create a new UInt32Type
uint32type := UInt32Type{

}
```

#### Type Definition

```go
type UInt32Type struct {
}
```

### UnbanUserFromGuildOptions
UnbanUserFromGuildOptions

#### Example Usage

```go
// Create a new UnbanUserFromGuildOptions
unbanuserfromguildoptions := UnbanUserFromGuildOptions{

}
```

#### Type Definition

```go
type UnbanUserFromGuildOptions struct {
}
```

### UnfurledMedia
UnfurledMedia

#### Example Usage

```go
// Create a new UnfurledMedia
unfurledmedia := UnfurledMedia{
    ID: "example",
    AttachmentID: "example",
    ContentType: &"example"{},
    Height: &42{},
    ProxyURL: "example",
    URL: "example",
    Width: &42{},
}
```

#### Type Definition

```go
type UnfurledMedia struct {
    ID string `json:"id"`
    AttachmentID string `json:"attachment_id,omitempty"`
    ContentType *string `json:"content_type,omitempty"`
    Height *int32 `json:"height,omitempty"`
    ProxyURL string `json:"proxy_url"`
    URL string `json:"url"`
    Width *int32 `json:"width,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| AttachmentID | `string` |  |
| ContentType | `*string` |  |
| Height | `*int32` |  |
| ProxyURL | `string` |  |
| URL | `string` |  |
| Width | `*int32` |  |

### UnfurledMediaOptions
UnfurledMediaOptions

#### Example Usage

```go
// Create a new UnfurledMediaOptions
unfurledmediaoptions := UnfurledMediaOptions{
    URL: "example",
}
```

#### Type Definition

```go
type UnfurledMediaOptions struct {
    URL string `json:"url"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| URL | `string` |  |

### UnfurledMediaOptionsWithAttachmentReferenceRequired
UnfurledMediaOptionsWithAttachmentReferenceRequired

#### Example Usage

```go
// Create a new UnfurledMediaOptionsWithAttachmentReferenceRequired
unfurledmediaoptionswithattachmentreferencerequired := UnfurledMediaOptionsWithAttachmentReferenceRequired{
    URL: "example",
}
```

#### Type Definition

```go
type UnfurledMediaOptionsWithAttachmentReferenceRequired struct {
    URL string `json:"url"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| URL | `string` |  |

### UpdateApplicationUserRoleConnectionOptions
UpdateApplicationUserRoleConnectionOptions

#### Example Usage

```go
// Create a new UpdateApplicationUserRoleConnectionOptions
updateapplicationuserroleconnectionoptions := UpdateApplicationUserRoleConnectionOptions{
    Metadata: map[],
    PlatformName: &"example"{},
    PlatformUsername: &"example"{},
}
```

#### Type Definition

```go
type UpdateApplicationUserRoleConnectionOptions struct {
    Metadata map[string]string `json:"metadata,omitempty"`
    PlatformName *string `json:"platform_name,omitempty"`
    PlatformUsername *string `json:"platform_username,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Metadata | `map[string]string` |  |
| PlatformName | `*string` |  |
| PlatformUsername | `*string` |  |

### UpdateDMOptionsPartial
UpdateDMOptionsPartial

#### Example Usage

```go
// Create a new UpdateDMOptionsPartial
updatedmoptionspartial := UpdateDMOptionsPartial{
    Name: &"example"{},
}
```

#### Type Definition

```go
type UpdateDMOptionsPartial struct {
    Name *string `json:"name,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `*string` |  |

### UpdateDefaultReactionEmojiOptions
UpdateDefaultReactionEmojiOptions

#### Example Usage

```go
// Create a new UpdateDefaultReactionEmojiOptions
updatedefaultreactionemojioptions := UpdateDefaultReactionEmojiOptions{
    EmojiID: any{},
    EmojiName: &"example"{},
}
```

#### Type Definition

```go
type UpdateDefaultReactionEmojiOptions struct {
    EmojiID any `json:"emoji_id,omitempty"`
    EmojiName *string `json:"emoji_name,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| EmojiID | `any` |  |
| EmojiName | `*string` |  |

### UpdateGroupDMOptionsPartial
UpdateGroupDMOptionsPartial

#### Example Usage

```go
// Create a new UpdateGroupDMOptionsPartial
updategroupdmoptionspartial := UpdateGroupDMOptionsPartial{
    Icon: &"example"{},
    Name: &"example"{},
}
```

#### Type Definition

```go
type UpdateGroupDMOptionsPartial struct {
    Icon *string `json:"icon,omitempty"`
    Name *string `json:"name,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Icon | `*string` |  |
| Name | `*string` |  |

### UpdateGuildChannelOptionsPartial
UpdateGuildChannelOptionsPartial

#### Example Usage

```go
// Create a new UpdateGuildChannelOptionsPartial
updateguildchanneloptionspartial := UpdateGuildChannelOptionsPartial{
    AvailableTags: [],
    Bitrate: &42{},
    DefaultAutoArchiveDuration: any{},
    DefaultForumLayout: any{},
    DefaultReactionEmoji: any{},
    DefaultSortOrder: any{},
    DefaultTagSetting: any{},
    DefaultThreadRateLimitPerUser: &42{},
    Flags: &42{},
    Name: "example",
    Nsfw: &true{},
    ParentID: any{},
    PermissionOverwrites: [],
    Position: &42{},
    RateLimitPerUser: &42{},
    RtcRegion: &"example"{},
    Topic: &"example"{},
    Type: any{},
    UserLimit: &42{},
    VideoQualityMode: any{},
}
```

#### Type Definition

```go
type UpdateGuildChannelOptionsPartial struct {
    AvailableTags []any `json:"available_tags,omitempty"`
    Bitrate *int32 `json:"bitrate,omitempty"`
    DefaultAutoArchiveDuration any `json:"default_auto_archive_duration,omitempty"`
    DefaultForumLayout any `json:"default_forum_layout,omitempty"`
    DefaultReactionEmoji any `json:"default_reaction_emoji,omitempty"`
    DefaultSortOrder any `json:"default_sort_order,omitempty"`
    DefaultTagSetting any `json:"default_tag_setting,omitempty"`
    DefaultThreadRateLimitPerUser *int64 `json:"default_thread_rate_limit_per_user,omitempty"`
    Flags *int64 `json:"flags,omitempty"`
    Name string `json:"name,omitempty"`
    Nsfw *bool `json:"nsfw,omitempty"`
    ParentID any `json:"parent_id,omitempty"`
    PermissionOverwrites []any `json:"permission_overwrites,omitempty"`
    Position *int32 `json:"position,omitempty"`
    RateLimitPerUser *int64 `json:"rate_limit_per_user,omitempty"`
    RtcRegion *string `json:"rtc_region,omitempty"`
    Topic *string `json:"topic,omitempty"`
    Type any `json:"type,omitempty"`
    UserLimit *int32 `json:"user_limit,omitempty"`
    VideoQualityMode any `json:"video_quality_mode,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AvailableTags | `[]any` |  |
| Bitrate | `*int32` |  |
| DefaultAutoArchiveDuration | `any` |  |
| DefaultForumLayout | `any` |  |
| DefaultReactionEmoji | `any` |  |
| DefaultSortOrder | `any` |  |
| DefaultTagSetting | `any` |  |
| DefaultThreadRateLimitPerUser | `*int64` |  |
| Flags | `*int64` |  |
| Name | `string` |  |
| Nsfw | `*bool` |  |
| ParentID | `any` |  |
| PermissionOverwrites | `[]any` |  |
| Position | `*int32` |  |
| RateLimitPerUser | `*int64` |  |
| RtcRegion | `*string` |  |
| Topic | `*string` |  |
| Type | `any` |  |
| UserLimit | `*int32` |  |
| VideoQualityMode | `any` |  |

### UpdateGuildOnboardingOptions
UpdateGuildOnboardingOptions

#### Example Usage

```go
// Create a new UpdateGuildOnboardingOptions
updateguildonboardingoptions := UpdateGuildOnboardingOptions{
    DefaultChannelIds: [],
    Enabled: &true{},
    Mode: any{},
    Prompts: [],
}
```

#### Type Definition

```go
type UpdateGuildOnboardingOptions struct {
    DefaultChannelIds []string `json:"default_channel_ids,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    Mode any `json:"mode,omitempty"`
    Prompts []any `json:"prompts,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| DefaultChannelIds | `[]string` |  |
| Enabled | `*bool` |  |
| Mode | `any` |  |
| Prompts | `[]any` |  |

### UpdateMessageInteractionCallback
UpdateMessageInteractionCallback

#### Example Usage

```go
// Create a new UpdateMessageInteractionCallback
updatemessageinteractioncallback := UpdateMessageInteractionCallback{
    Message: any{},
    Type: 42,
}
```

#### Type Definition

```go
type UpdateMessageInteractionCallback struct {
    Message any `json:"message"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Message | `any` |  |
| Type | `int32` |  |

### UpdateMessageInteractionCallbackOptions
UpdateMessageInteractionCallbackOptions

#### Example Usage

```go
// Create a new UpdateMessageInteractionCallbackOptions
updatemessageinteractioncallbackoptions := UpdateMessageInteractionCallbackOptions{
    Data: any{},
    Type: 42,
}
```

#### Type Definition

```go
type UpdateMessageInteractionCallbackOptions struct {
    Data any `json:"data,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Data | `any` |  |
| Type | `int32` |  |

### UpdateOnboardingPromptOptions
UpdateOnboardingPromptOptions

#### Example Usage

```go
// Create a new UpdateOnboardingPromptOptions
updateonboardingpromptoptions := UpdateOnboardingPromptOptions{
    ID: "example",
    InOnboarding: &true{},
    Options: [],
    Required: &true{},
    SingleSelect: &true{},
    Title: "example",
    Type: any{},
}
```

#### Type Definition

```go
type UpdateOnboardingPromptOptions struct {
    ID string `json:"id"`
    InOnboarding *bool `json:"in_onboarding,omitempty"`
    Options []any `json:"options"`
    Required *bool `json:"required,omitempty"`
    SingleSelect *bool `json:"single_select,omitempty"`
    Title string `json:"title"`
    Type any `json:"type,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| InOnboarding | `*bool` |  |
| Options | `[]any` |  |
| Required | `*bool` |  |
| SingleSelect | `*bool` |  |
| Title | `string` |  |
| Type | `any` |  |

### UpdateRoleOptionsPartial
UpdateRoleOptionsPartial

#### Example Usage

```go
// Create a new UpdateRoleOptionsPartial
updateroleoptionspartial := UpdateRoleOptionsPartial{
    Color: &42{},
    Hoist: &true{},
    Icon: &"example"{},
    Mentionable: &true{},
    Name: &"example"{},
    Permissions: &42{},
    UnicodeEmoji: &"example"{},
}
```

#### Type Definition

```go
type UpdateRoleOptionsPartial struct {
    Color *int64 `json:"color,omitempty"`
    Hoist *bool `json:"hoist,omitempty"`
    Icon *string `json:"icon,omitempty"`
    Mentionable *bool `json:"mentionable,omitempty"`
    Name *string `json:"name,omitempty"`
    Permissions *int64 `json:"permissions,omitempty"`
    UnicodeEmoji *string `json:"unicode_emoji,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Color | `*int64` |  |
| Hoist | `*bool` |  |
| Icon | `*string` |  |
| Mentionable | `*bool` |  |
| Name | `*string` |  |
| Permissions | `*int64` |  |
| UnicodeEmoji | `*string` |  |

### UpdateRolePositionsOptions
UpdateRolePositionsOptions

#### Example Usage

```go
// Create a new UpdateRolePositionsOptions
updaterolepositionsoptions := UpdateRolePositionsOptions{
    ID: any{},
    Position: &42{},
}
```

#### Type Definition

```go
type UpdateRolePositionsOptions struct {
    ID any `json:"id,omitempty"`
    Position *int32 `json:"position,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `any` |  |
| Position | `*int32` |  |

### UpdateSelfVoiceStateOptionsPartial
UpdateSelfVoiceStateOptionsPartial

#### Example Usage

```go
// Create a new UpdateSelfVoiceStateOptionsPartial
updateselfvoicestateoptionspartial := UpdateSelfVoiceStateOptionsPartial{
    ChannelID: any{},
    RequestToSpeakTimestamp: &/* value */{},
    Suppress: &true{},
}
```

#### Type Definition

```go
type UpdateSelfVoiceStateOptionsPartial struct {
    ChannelID any `json:"channel_id,omitempty"`
    RequestToSpeakTimestamp *time.Time `json:"request_to_speak_timestamp,omitempty"`
    Suppress *bool `json:"suppress,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `any` |  |
| RequestToSpeakTimestamp | `*time.Time` |  |
| Suppress | `*bool` |  |

### UpdateThreadOptionsPartial
UpdateThreadOptionsPartial

#### Example Usage

```go
// Create a new UpdateThreadOptionsPartial
updatethreadoptionspartial := UpdateThreadOptionsPartial{
    AppliedTags: [],
    Archived: &true{},
    AutoArchiveDuration: any{},
    Bitrate: &42{},
    Flags: &42{},
    Invitable: &true{},
    Locked: &true{},
    Name: &"example"{},
    RateLimitPerUser: &42{},
    RtcRegion: &"example"{},
    UserLimit: &42{},
    VideoQualityMode: any{},
}
```

#### Type Definition

```go
type UpdateThreadOptionsPartial struct {
    AppliedTags []string `json:"applied_tags,omitempty"`
    Archived *bool `json:"archived,omitempty"`
    AutoArchiveDuration any `json:"auto_archive_duration,omitempty"`
    Bitrate *int32 `json:"bitrate,omitempty"`
    Flags *int64 `json:"flags,omitempty"`
    Invitable *bool `json:"invitable,omitempty"`
    Locked *bool `json:"locked,omitempty"`
    Name *string `json:"name,omitempty"`
    RateLimitPerUser *int64 `json:"rate_limit_per_user,omitempty"`
    RtcRegion *string `json:"rtc_region,omitempty"`
    UserLimit *int64 `json:"user_limit,omitempty"`
    VideoQualityMode any `json:"video_quality_mode,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AppliedTags | `[]string` |  |
| Archived | `*bool` |  |
| AutoArchiveDuration | `any` |  |
| Bitrate | `*int32` |  |
| Flags | `*int64` |  |
| Invitable | `*bool` |  |
| Locked | `*bool` |  |
| Name | `*string` |  |
| RateLimitPerUser | `*int64` |  |
| RtcRegion | `*string` |  |
| UserLimit | `*int64` |  |
| VideoQualityMode | `any` |  |

### UpdateThreadTagOptions
UpdateThreadTagOptions

#### Example Usage

```go
// Create a new UpdateThreadTagOptions
updatethreadtagoptions := UpdateThreadTagOptions{
    ID: any{},
    EmojiID: any{},
    EmojiName: &"example"{},
    Moderated: &true{},
    Name: "example",
}
```

#### Type Definition

```go
type UpdateThreadTagOptions struct {
    ID any `json:"id,omitempty"`
    EmojiID any `json:"emoji_id,omitempty"`
    EmojiName *string `json:"emoji_name,omitempty"`
    Moderated *bool `json:"moderated,omitempty"`
    Name string `json:"name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `any` |  |
| EmojiID | `any` |  |
| EmojiName | `*string` |  |
| Moderated | `*bool` |  |
| Name | `string` |  |

### UpdateVoiceStateOptionsPartial
UpdateVoiceStateOptionsPartial

#### Example Usage

```go
// Create a new UpdateVoiceStateOptionsPartial
updatevoicestateoptionspartial := UpdateVoiceStateOptionsPartial{
    ChannelID: any{},
    Suppress: &true{},
}
```

#### Type Definition

```go
type UpdateVoiceStateOptionsPartial struct {
    ChannelID any `json:"channel_id,omitempty"`
    Suppress *bool `json:"suppress,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `any` |  |
| Suppress | `*bool` |  |

### User
User

#### Example Usage

```go
// Create a new User
user := User{
    ID: "example",
    AccentColor: &42{},
    Avatar: &"example"{},
    AvatarDecorationData: any{},
    Banner: &"example"{},
    Bot: true,
    Collectibles: any{},
    Discriminator: "example",
    Flags: 42,
    GlobalName: &"example"{},
    PrimaryGuild: any{},
    PublicFlags: 42,
    System: true,
    Username: "example",
}
```

#### Type Definition

```go
type User struct {
    ID string `json:"id"`
    AccentColor *int32 `json:"accent_color,omitempty"`
    Avatar *string `json:"avatar,omitempty"`
    AvatarDecorationData any `json:"avatar_decoration_data,omitempty"`
    Banner *string `json:"banner,omitempty"`
    Bot bool `json:"bot,omitempty"`
    Collectibles any `json:"collectibles,omitempty"`
    Discriminator string `json:"discriminator"`
    Flags int64 `json:"flags"`
    GlobalName *string `json:"global_name,omitempty"`
    PrimaryGuild any `json:"primary_guild,omitempty"`
    PublicFlags int32 `json:"public_flags"`
    System bool `json:"system,omitempty"`
    Username string `json:"username"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| AccentColor | `*int32` |  |
| Avatar | `*string` |  |
| AvatarDecorationData | `any` |  |
| Banner | `*string` |  |
| Bot | `bool` |  |
| Collectibles | `any` |  |
| Discriminator | `string` |  |
| Flags | `int64` |  |
| GlobalName | `*string` |  |
| PrimaryGuild | `any` |  |
| PublicFlags | `int32` |  |
| System | `bool` |  |
| Username | `string` |  |

### UserAvatarDecoration
UserAvatarDecoration

#### Example Usage

```go
// Create a new UserAvatarDecoration
useravatardecoration := UserAvatarDecoration{
    Asset: "example",
    SkuID: any{},
}
```

#### Type Definition

```go
type UserAvatarDecoration struct {
    Asset string `json:"asset"`
    SkuID any `json:"sku_id,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Asset | `string` |  |
| SkuID | `any` |  |

### UserCollectibles
UserCollectibles

#### Example Usage

```go
// Create a new UserCollectibles
usercollectibles := UserCollectibles{
    Nameplate: any{},
}
```

#### Type Definition

```go
type UserCollectibles struct {
    Nameplate any `json:"nameplate,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Nameplate | `any` |  |

### UserCommunicationDisabledAction
UserCommunicationDisabledAction

#### Example Usage

```go
// Create a new UserCommunicationDisabledAction
usercommunicationdisabledaction := UserCommunicationDisabledAction{
    Metadata: any{},
    Type: 42,
}
```

#### Type Definition

```go
type UserCommunicationDisabledAction struct {
    Metadata any `json:"metadata"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Metadata | `any` |  |
| Type | `int32` |  |

### UserCommunicationDisabledActionMetadata
UserCommunicationDisabledActionMetadata

#### Example Usage

```go
// Create a new UserCommunicationDisabledActionMetadata
usercommunicationdisabledactionmetadata := UserCommunicationDisabledActionMetadata{
    DurationSeconds: 42,
}
```

#### Type Definition

```go
type UserCommunicationDisabledActionMetadata struct {
    DurationSeconds int32 `json:"duration_seconds"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| DurationSeconds | `int32` |  |

### UserGuildOnboarding
UserGuildOnboarding

#### Example Usage

```go
// Create a new UserGuildOnboarding
userguildonboarding := UserGuildOnboarding{
    DefaultChannelIds: [],
    Enabled: true,
    GuildID: "example",
    Prompts: [],
}
```

#### Type Definition

```go
type UserGuildOnboarding struct {
    DefaultChannelIds []string `json:"default_channel_ids"`
    Enabled bool `json:"enabled"`
    GuildID string `json:"guild_id"`
    Prompts []any `json:"prompts"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| DefaultChannelIds | `[]string` |  |
| Enabled | `bool` |  |
| GuildID | `string` |  |
| Prompts | `[]any` |  |

### UserNameplate
UserNameplate

#### Example Usage

```go
// Create a new UserNameplate
usernameplate := UserNameplate{
    Asset: "example",
    Label: "example",
    Palette: "example",
    SkuID: any{},
}
```

#### Type Definition

```go
type UserNameplate struct {
    Asset string `json:"asset"`
    Label string `json:"label"`
    Palette string `json:"palette"`
    SkuID any `json:"sku_id,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Asset | `string` |  |
| Label | `string` |  |
| Palette | `string` |  |
| SkuID | `any` |  |

### UserNotificationSettings
UserNotificationSettings

#### Example Usage

```go
// Example usage of UserNotificationSettings
var value UserNotificationSettings
// Initialize with appropriate value
```

#### Type Definition

```go
type UserNotificationSettings int
```

### UserPII
UserPII

#### Example Usage

```go
// Create a new UserPII
userpii := UserPII{
    ID: "example",
    AccentColor: &42{},
    Avatar: &"example"{},
    AvatarDecorationData: any{},
    Banner: &"example"{},
    Bot: true,
    Collectibles: any{},
    Discriminator: "example",
    Email: &"example"{},
    Flags: 42,
    GlobalName: &"example"{},
    Locale: "example",
    MfaEnabled: true,
    PremiumType: 42,
    PrimaryGuild: any{},
    PublicFlags: 42,
    System: true,
    Username: "example",
    Verified: true,
}
```

#### Type Definition

```go
type UserPII struct {
    ID string `json:"id"`
    AccentColor *int32 `json:"accent_color,omitempty"`
    Avatar *string `json:"avatar,omitempty"`
    AvatarDecorationData any `json:"avatar_decoration_data,omitempty"`
    Banner *string `json:"banner,omitempty"`
    Bot bool `json:"bot,omitempty"`
    Collectibles any `json:"collectibles,omitempty"`
    Discriminator string `json:"discriminator"`
    Email *string `json:"email,omitempty"`
    Flags int64 `json:"flags"`
    GlobalName *string `json:"global_name,omitempty"`
    Locale string `json:"locale"`
    MfaEnabled bool `json:"mfa_enabled"`
    PremiumType int32 `json:"premium_type,omitempty"`
    PrimaryGuild any `json:"primary_guild,omitempty"`
    PublicFlags int32 `json:"public_flags"`
    System bool `json:"system,omitempty"`
    Username string `json:"username"`
    Verified bool `json:"verified,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| AccentColor | `*int32` |  |
| Avatar | `*string` |  |
| AvatarDecorationData | `any` |  |
| Banner | `*string` |  |
| Bot | `bool` |  |
| Collectibles | `any` |  |
| Discriminator | `string` |  |
| Email | `*string` |  |
| Flags | `int64` |  |
| GlobalName | `*string` |  |
| Locale | `string` |  |
| MfaEnabled | `bool` |  |
| PremiumType | `int32` |  |
| PrimaryGuild | `any` |  |
| PublicFlags | `int32` |  |
| System | `bool` |  |
| Username | `string` |  |
| Verified | `bool` |  |

### UserPrimaryGuild
UserPrimaryGuild

#### Example Usage

```go
// Create a new UserPrimaryGuild
userprimaryguild := UserPrimaryGuild{
    Badge: &"example"{},
    IdentityEnabled: &true{},
    IdentityGuildID: any{},
    Tag: &"example"{},
}
```

#### Type Definition

```go
type UserPrimaryGuild struct {
    Badge *string `json:"badge,omitempty"`
    IdentityEnabled *bool `json:"identity_enabled,omitempty"`
    IdentityGuildID any `json:"identity_guild_id,omitempty"`
    Tag *string `json:"tag,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Badge | `*string` |  |
| IdentityEnabled | `*bool` |  |
| IdentityGuildID | `any` |  |
| Tag | `*string` |  |

### UserSelectComponent
UserSelectComponent

#### Example Usage

```go
// Create a new UserSelectComponent
userselectcomponent := UserSelectComponent{
    ID: 42,
    CustomID: "example",
    DefaultValues: [],
    Disabled: true,
    MaxValues: &42{},
    MinValues: &42{},
    Placeholder: "example",
    Type: 42,
}
```

#### Type Definition

```go
type UserSelectComponent struct {
    ID int32 `json:"id"`
    CustomID string `json:"custom_id"`
    DefaultValues []any `json:"default_values,omitempty"`
    Disabled bool `json:"disabled,omitempty"`
    MaxValues *int32 `json:"max_values,omitempty"`
    MinValues *int32 `json:"min_values,omitempty"`
    Placeholder string `json:"placeholder,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `int32` |  |
| CustomID | `string` |  |
| DefaultValues | `[]any` |  |
| Disabled | `bool` |  |
| MaxValues | `*int32` |  |
| MinValues | `*int32` |  |
| Placeholder | `string` |  |
| Type | `int32` |  |

### UserSelectComponentForMessageOptions
UserSelectComponentForMessageOptions

#### Example Usage

```go
// Create a new UserSelectComponentForMessageOptions
userselectcomponentformessageoptions := UserSelectComponentForMessageOptions{
    ID: &42{},
    CustomID: "example",
    DefaultValues: [],
    Disabled: &true{},
    MaxValues: &42{},
    MinValues: &42{},
    Placeholder: &"example"{},
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type UserSelectComponentForMessageOptions struct {
    ID *int32 `json:"id,omitempty"`
    CustomID string `json:"custom_id"`
    DefaultValues []any `json:"default_values,omitempty"`
    Disabled *bool `json:"disabled,omitempty"`
    MaxValues *int64 `json:"max_values,omitempty"`
    MinValues *int64 `json:"min_values,omitempty"`
    Placeholder *string `json:"placeholder,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| CustomID | `string` |  |
| DefaultValues | `[]any` |  |
| Disabled | `*bool` |  |
| MaxValues | `*int64` |  |
| MinValues | `*int64` |  |
| Placeholder | `*string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### UserSelectComponentForModalOptions
UserSelectComponentForModalOptions

#### Example Usage

```go
// Create a new UserSelectComponentForModalOptions
userselectcomponentformodaloptions := UserSelectComponentForModalOptions{
    ID: &42{},
    CustomID: "example",
    DefaultValues: [],
    Disabled: &true{},
    MaxValues: &42{},
    MinValues: &42{},
    Placeholder: &"example"{},
    Required: &true{},
    Type: 42,
}
```

#### Type Definition

```go
type UserSelectComponentForModalOptions struct {
    ID *int32 `json:"id,omitempty"`
    CustomID string `json:"custom_id"`
    DefaultValues []any `json:"default_values,omitempty"`
    Disabled *bool `json:"disabled,omitempty"`
    MaxValues *int64 `json:"max_values,omitempty"`
    MinValues *int64 `json:"min_values,omitempty"`
    Placeholder *string `json:"placeholder,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type int32 `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `*int32` |  |
| CustomID | `string` |  |
| DefaultValues | `[]any` |  |
| Disabled | `*bool` |  |
| MaxValues | `*int64` |  |
| MinValues | `*int64` |  |
| Placeholder | `*string` |  |
| Required | `*bool` |  |
| Type | `int32` |  |

### UserSelectDefaultValue
UserSelectDefaultValue

#### Example Usage

```go
// Create a new UserSelectDefaultValue
userselectdefaultvalue := UserSelectDefaultValue{
    ID: "example",
    Type: "example",
}
```

#### Type Definition

```go
type UserSelectDefaultValue struct {
    ID string `json:"id"`
    Type string `json:"type"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Type | `string` |  |

### VanityURL
VanityURL

#### Example Usage

```go
// Create a new VanityURL
vanityurl := VanityURL{
    Code: &"example"{},
    Error: any{},
    Uses: 42,
}
```

#### Type Definition

```go
type VanityURL struct {
    Code *string `json:"code,omitempty"`
    Error any `json:"error,omitempty"`
    Uses int32 `json:"uses"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Code | `*string` |  |
| Error | `any` |  |
| Uses | `int32` |  |

### VanityURLError
VanityURLError

#### Example Usage

```go
// Create a new VanityURLError
vanityurlerror := VanityURLError{
    Code: 42,
    Message: "example",
}
```

#### Type Definition

```go
type VanityURLError struct {
    Code int32 `json:"code"`
    Message string `json:"message"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Code | `int32` |  |
| Message | `string` |  |

### VerificationLevels
VerificationLevels

#### Example Usage

```go
// Example usage of VerificationLevels
var value VerificationLevels
// Initialize with appropriate value
```

#### Type Definition

```go
type VerificationLevels int
```

### VideoQualityModes
VideoQualityModes

#### Example Usage

```go
// Example usage of VideoQualityModes
var value VideoQualityModes
// Initialize with appropriate value
```

#### Type Definition

```go
type VideoQualityModes int
```

### VoiceRegion
VoiceRegion

#### Example Usage

```go
// Create a new VoiceRegion
voiceregion := VoiceRegion{
    ID: "example",
    Custom: true,
    Deprecated: true,
    Name: "example",
    Optimal: true,
}
```

#### Type Definition

```go
type VoiceRegion struct {
    ID string `json:"id"`
    Custom bool `json:"custom"`
    Deprecated bool `json:"deprecated"`
    Name string `json:"name"`
    Optimal bool `json:"optimal"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Custom | `bool` |  |
| Deprecated | `bool` |  |
| Name | `string` |  |
| Optimal | `bool` |  |

### VoiceScheduledEvent
VoiceScheduledEvent

#### Example Usage

```go
// Create a new VoiceScheduledEvent
voicescheduledevent := VoiceScheduledEvent{
    ID: "example",
    ChannelID: any{},
    Creator: any{},
    CreatorID: any{},
    Description: &"example"{},
    EntityID: any{},
    EntityMetadata: any{},
    EntityType: 42,
    GuildID: "example",
    Image: &"example"{},
    Name: "example",
    PrivacyLevel: 42,
    ScheduledEndTime: &/* value */{},
    ScheduledStartTime: /* value */,
    Status: 42,
    UserCount: 42,
    UserRsvp: any{},
}
```

#### Type Definition

```go
type VoiceScheduledEvent struct {
    ID string `json:"id"`
    ChannelID any `json:"channel_id,omitempty"`
    Creator any `json:"creator,omitempty"`
    CreatorID any `json:"creator_id,omitempty"`
    Description *string `json:"description,omitempty"`
    EntityID any `json:"entity_id,omitempty"`
    EntityMetadata any `json:"entity_metadata,omitempty"`
    EntityType int32 `json:"entity_type"`
    GuildID string `json:"guild_id"`
    Image *string `json:"image,omitempty"`
    Name string `json:"name"`
    PrivacyLevel int32 `json:"privacy_level"`
    ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
    ScheduledStartTime time.Time `json:"scheduled_start_time"`
    Status int32 `json:"status"`
    UserCount int32 `json:"user_count,omitempty"`
    UserRsvp any `json:"user_rsvp,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| ChannelID | `any` |  |
| Creator | `any` |  |
| CreatorID | `any` |  |
| Description | `*string` |  |
| EntityID | `any` |  |
| EntityMetadata | `any` |  |
| EntityType | `int32` |  |
| GuildID | `string` |  |
| Image | `*string` |  |
| Name | `string` |  |
| PrivacyLevel | `int32` |  |
| ScheduledEndTime | `*time.Time` |  |
| ScheduledStartTime | `time.Time` |  |
| Status | `int32` |  |
| UserCount | `int32` |  |
| UserRsvp | `any` |  |

### VoiceScheduledEventCreateOptions
VoiceScheduledEventCreateOptions

#### Example Usage

```go
// Create a new VoiceScheduledEventCreateOptions
voicescheduledeventcreateoptions := VoiceScheduledEventCreateOptions{
    ChannelID: any{},
    Description: &"example"{},
    EntityMetadata: any{},
    EntityType: 42,
    Image: &"example"{},
    Name: "example",
    PrivacyLevel: 42,
    ScheduledEndTime: &/* value */{},
    ScheduledStartTime: /* value */,
}
```

#### Type Definition

```go
type VoiceScheduledEventCreateOptions struct {
    ChannelID any `json:"channel_id,omitempty"`
    Description *string `json:"description,omitempty"`
    EntityMetadata any `json:"entity_metadata,omitempty"`
    EntityType int32 `json:"entity_type"`
    Image *string `json:"image,omitempty"`
    Name string `json:"name"`
    PrivacyLevel int32 `json:"privacy_level"`
    ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
    ScheduledStartTime time.Time `json:"scheduled_start_time"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `any` |  |
| Description | `*string` |  |
| EntityMetadata | `any` |  |
| EntityType | `int32` |  |
| Image | `*string` |  |
| Name | `string` |  |
| PrivacyLevel | `int32` |  |
| ScheduledEndTime | `*time.Time` |  |
| ScheduledStartTime | `time.Time` |  |

### VoiceScheduledEventPatchOptionsPartial
VoiceScheduledEventPatchOptionsPartial

#### Example Usage

```go
// Create a new VoiceScheduledEventPatchOptionsPartial
voicescheduledeventpatchoptionspartial := VoiceScheduledEventPatchOptionsPartial{
    ChannelID: any{},
    Description: &"example"{},
    EntityMetadata: any{},
    EntityType: any{},
    Image: &"example"{},
    Name: "example",
    PrivacyLevel: 42,
    ScheduledEndTime: &/* value */{},
    ScheduledStartTime: /* value */,
    Status: any{},
}
```

#### Type Definition

```go
type VoiceScheduledEventPatchOptionsPartial struct {
    ChannelID any `json:"channel_id,omitempty"`
    Description *string `json:"description,omitempty"`
    EntityMetadata any `json:"entity_metadata,omitempty"`
    EntityType any `json:"entity_type,omitempty"`
    Image *string `json:"image,omitempty"`
    Name string `json:"name,omitempty"`
    PrivacyLevel int32 `json:"privacy_level,omitempty"`
    ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
    ScheduledStartTime time.Time `json:"scheduled_start_time,omitempty"`
    Status any `json:"status,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `any` |  |
| Description | `*string` |  |
| EntityMetadata | `any` |  |
| EntityType | `any` |  |
| Image | `*string` |  |
| Name | `string` |  |
| PrivacyLevel | `int32` |  |
| ScheduledEndTime | `*time.Time` |  |
| ScheduledStartTime | `time.Time` |  |
| Status | `any` |  |

### VoiceState
VoiceState

#### Example Usage

```go
// Create a new VoiceState
voicestate := VoiceState{
    ChannelID: any{},
    Deaf: true,
    GuildID: any{},
    Member: any{},
    Mute: true,
    RequestToSpeakTimestamp: &/* value */{},
    SelfDeaf: true,
    SelfMute: true,
    SelfStream: &true{},
    SelfVideo: true,
    SessionID: "example",
    Suppress: true,
    UserID: "example",
}
```

#### Type Definition

```go
type VoiceState struct {
    ChannelID any `json:"channel_id,omitempty"`
    Deaf bool `json:"deaf"`
    GuildID any `json:"guild_id,omitempty"`
    Member any `json:"member,omitempty"`
    Mute bool `json:"mute"`
    RequestToSpeakTimestamp *time.Time `json:"request_to_speak_timestamp,omitempty"`
    SelfDeaf bool `json:"self_deaf"`
    SelfMute bool `json:"self_mute"`
    SelfStream *bool `json:"self_stream,omitempty"`
    SelfVideo bool `json:"self_video"`
    SessionID string `json:"session_id"`
    Suppress bool `json:"suppress"`
    UserID string `json:"user_id"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `any` |  |
| Deaf | `bool` |  |
| GuildID | `any` |  |
| Member | `any` |  |
| Mute | `bool` |  |
| RequestToSpeakTimestamp | `*time.Time` |  |
| SelfDeaf | `bool` |  |
| SelfMute | `bool` |  |
| SelfStream | `*bool` |  |
| SelfVideo | `bool` |  |
| SessionID | `string` |  |
| Suppress | `bool` |  |
| UserID | `string` |  |

### WebhookSlackEmbed
WebhookSlackEmbed

#### Example Usage

```go
// Create a new WebhookSlackEmbed
webhookslackembed := WebhookSlackEmbed{
    AuthorIcon: &"example"{},
    AuthorLink: &"example"{},
    AuthorName: &"example"{},
    Color: &"example"{},
    Fields: [],
    Footer: &"example"{},
    FooterIcon: &"example"{},
    ImageURL: &"example"{},
    Pretext: &"example"{},
    Text: &"example"{},
    ThumbURL: &"example"{},
    Title: &"example"{},
    TitleLink: &"example"{},
    Ts: &42{},
}
```

#### Type Definition

```go
type WebhookSlackEmbed struct {
    AuthorIcon *string `json:"author_icon,omitempty"`
    AuthorLink *string `json:"author_link,omitempty"`
    AuthorName *string `json:"author_name,omitempty"`
    Color *string `json:"color,omitempty"`
    Fields []any `json:"fields,omitempty"`
    Footer *string `json:"footer,omitempty"`
    FooterIcon *string `json:"footer_icon,omitempty"`
    ImageURL *string `json:"image_url,omitempty"`
    Pretext *string `json:"pretext,omitempty"`
    Text *string `json:"text,omitempty"`
    ThumbURL *string `json:"thumb_url,omitempty"`
    Title *string `json:"title,omitempty"`
    TitleLink *string `json:"title_link,omitempty"`
    Ts *int64 `json:"ts,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AuthorIcon | `*string` |  |
| AuthorLink | `*string` |  |
| AuthorName | `*string` |  |
| Color | `*string` |  |
| Fields | `[]any` |  |
| Footer | `*string` |  |
| FooterIcon | `*string` |  |
| ImageURL | `*string` |  |
| Pretext | `*string` |  |
| Text | `*string` |  |
| ThumbURL | `*string` |  |
| Title | `*string` |  |
| TitleLink | `*string` |  |
| Ts | `*int64` |  |

### WebhookSlackEmbedField
WebhookSlackEmbedField

#### Example Usage

```go
// Create a new WebhookSlackEmbedField
webhookslackembedfield := WebhookSlackEmbedField{
    Inline: &true{},
    Name: &"example"{},
    Value: &"example"{},
}
```

#### Type Definition

```go
type WebhookSlackEmbedField struct {
    Inline *bool `json:"inline,omitempty"`
    Name *string `json:"name,omitempty"`
    Value *string `json:"value,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Inline | `*bool` |  |
| Name | `*string` |  |
| Value | `*string` |  |

### WebhookSourceChannel
WebhookSourceChannel

#### Example Usage

```go
// Create a new WebhookSourceChannel
webhooksourcechannel := WebhookSourceChannel{
    ID: "example",
    Name: "example",
}
```

#### Type Definition

```go
type WebhookSourceChannel struct {
    ID string `json:"id"`
    Name string `json:"name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Name | `string` |  |

### WebhookSourceGuild
WebhookSourceGuild

#### Example Usage

```go
// Create a new WebhookSourceGuild
webhooksourceguild := WebhookSourceGuild{
    ID: "example",
    Icon: &"example"{},
    Name: "example",
}
```

#### Type Definition

```go
type WebhookSourceGuild struct {
    ID string `json:"id"`
    Icon *string `json:"icon,omitempty"`
    Name string `json:"name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Icon | `*string` |  |
| Name | `string` |  |

### WebhookTypes
WebhookTypes

#### Example Usage

```go
// Example usage of WebhookTypes
var value WebhookTypes
// Initialize with appropriate value
```

#### Type Definition

```go
type WebhookTypes int
```

### WelcomeMessage
WelcomeMessage

#### Example Usage

```go
// Create a new WelcomeMessage
welcomemessage := WelcomeMessage{
    AuthorIds: [],
    Message: "example",
}
```

#### Type Definition

```go
type WelcomeMessage struct {
    AuthorIds []string `json:"author_ids"`
    Message string `json:"message"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| AuthorIds | `[]string` |  |
| Message | `string` |  |

### WelcomeScreenPatchOptionsPartial
WelcomeScreenPatchOptionsPartial

#### Example Usage

```go
// Create a new WelcomeScreenPatchOptionsPartial
welcomescreenpatchoptionspartial := WelcomeScreenPatchOptionsPartial{
    Description: &"example"{},
    Enabled: &true{},
    WelcomeChannels: [],
}
```

#### Type Definition

```go
type WelcomeScreenPatchOptionsPartial struct {
    Description *string `json:"description,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    WelcomeChannels []any `json:"welcome_channels,omitempty"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Description | `*string` |  |
| Enabled | `*bool` |  |
| WelcomeChannels | `[]any` |  |

### Widget
Widget

#### Example Usage

```go
// Create a new Widget
widget := Widget{
    ID: "example",
    Channels: [],
    InstantInvite: &"example"{},
    Members: [],
    Name: "example",
    PresenceCount: 42,
}
```

#### Type Definition

```go
type Widget struct {
    ID string `json:"id"`
    Channels []any `json:"channels"`
    InstantInvite *string `json:"instant_invite,omitempty"`
    Members []any `json:"members"`
    Name string `json:"name"`
    PresenceCount int32 `json:"presence_count"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Channels | `[]any` |  |
| InstantInvite | `*string` |  |
| Members | `[]any` |  |
| Name | `string` |  |
| PresenceCount | `int32` |  |

### WidgetActivity
WidgetActivity

#### Example Usage

```go
// Create a new WidgetActivity
widgetactivity := WidgetActivity{
    Name: "example",
}
```

#### Type Definition

```go
type WidgetActivity struct {
    Name string `json:"name"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Name | `string` |  |

### WidgetChannel
WidgetChannel

#### Example Usage

```go
// Create a new WidgetChannel
widgetchannel := WidgetChannel{
    ID: "example",
    Name: "example",
    Position: 42,
}
```

#### Type Definition

```go
type WidgetChannel struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Position int32 `json:"position"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Name | `string` |  |
| Position | `int32` |  |

### WidgetImageStyles
WidgetImageStyles

#### Example Usage

```go
// Example usage of WidgetImageStyles
var value WidgetImageStyles
// Initialize with appropriate value
```

#### Type Definition

```go
type WidgetImageStyles string
```

### WidgetMember
WidgetMember

#### Example Usage

```go
// Create a new WidgetMember
widgetmember := WidgetMember{
    ID: "example",
    Activity: any{},
    Avatar: any{},
    AvatarURL: "example",
    ChannelID: "example",
    Deaf: true,
    Discriminator: "example",
    Mute: true,
    SelfDeaf: true,
    SelfMute: true,
    Status: "example",
    Suppress: true,
    Username: "example",
}
```

#### Type Definition

```go
type WidgetMember struct {
    ID string `json:"id"`
    Activity any `json:"activity,omitempty"`
    Avatar any `json:"avatar,omitempty"`
    AvatarURL string `json:"avatar_url"`
    ChannelID string `json:"channel_id,omitempty"`
    Deaf bool `json:"deaf,omitempty"`
    Discriminator string `json:"discriminator"`
    Mute bool `json:"mute,omitempty"`
    SelfDeaf bool `json:"self_deaf,omitempty"`
    SelfMute bool `json:"self_mute,omitempty"`
    Status string `json:"status"`
    Suppress bool `json:"suppress,omitempty"`
    Username string `json:"username"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | `string` |  |
| Activity | `any` |  |
| Avatar | `any` |  |
| AvatarURL | `string` |  |
| ChannelID | `string` |  |
| Deaf | `bool` |  |
| Discriminator | `string` |  |
| Mute | `bool` |  |
| SelfDeaf | `bool` |  |
| SelfMute | `bool` |  |
| Status | `string` |  |
| Suppress | `bool` |  |
| Username | `string` |  |

### WidgetSettings
WidgetSettings

#### Example Usage

```go
// Create a new WidgetSettings
widgetsettings := WidgetSettings{
    ChannelID: any{},
    Enabled: true,
}
```

#### Type Definition

```go
type WidgetSettings struct {
    ChannelID any `json:"channel_id,omitempty"`
    Enabled bool `json:"enabled"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| ChannelID | `any` |  |
| Enabled | `bool` |  |

### WidgetUserDiscriminator
WidgetUserDiscriminator

#### Example Usage

```go
// Example usage of WidgetUserDiscriminator
var value WidgetUserDiscriminator
// Initialize with appropriate value
```

#### Type Definition

```go
type WidgetUserDiscriminator string
```

## External Links

- [Package Overview](../packages/models.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/discord/models)
- [Source Code](https://github.com/kolosys/discord/tree/main/models)

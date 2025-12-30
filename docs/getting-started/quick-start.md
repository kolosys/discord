# Quick Start

This guide will help you get started with discord quickly with a basic example.

## Basic Usage

Here's a simple example to get you started:

```go
package main

import (
    "fmt"
    "log"
    "github.com/kolosys/discord/commands"
    "github.com/kolosys/discord"
    "github.com/kolosys/discord/components"
    "github.com/kolosys/discord/embed"
    "github.com/kolosys/discord/events"
    "github.com/kolosys/discord/examples/basic/bot"
    "github.com/kolosys/discord/examples/basic/commands"
    "github.com/kolosys/discord/examples/basic/events"
    "github.com/kolosys/discord/examples/basic/routes"
    "github.com/kolosys/discord/examples/internal"
    "github.com/kolosys/discord/gateway"
    "github.com/kolosys/discord/models"
    "github.com/kolosys/discord/rest/internal"
    "github.com/kolosys/discord/rest"
    "github.com/kolosys/discord/server"
    "github.com/kolosys/discord/state"
    "github.com/kolosys/discord/types"
    "github.com/kolosys/discord/voice"
)

func main() {
    // Basic usage example
    fmt.Println("Welcome to discord!")
    
    // TODO: Add your code here
}
```

## Common Use Cases

### Using commands

**Import Path:** `github.com/kolosys/discord/commands`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/commands"
)

func main() {
    // Example usage of commands
    fmt.Println("Using commands package")
}
```

#### Available Types
- **AllowedMentions** - AllowedMentions configures which mentions are parsed.
- **ApplicationCommandChoice** - ApplicationCommandChoice represents a choice for an option.
- **ApplicationCommandCreate** - ApplicationCommandCreate represents a command to create with Discord.
- **ApplicationCommandOption** - ApplicationCommandOption represents an option for an application command.
- **AutocompleteHandler** - AutocompleteHandler is an optional interface for commands that support autocomplete.
- **Choice** - Choice represents an autocomplete choice.
- **Command** - Command is the interface all commands implement.
- **CommandGroup** - CommandGroup represents a group of related commands.
- **CommandSyncer** - CommandSyncer is an interface for syncing commands with Discord.
- **CommandType** - CommandType represents the type of application command.
- **ComponentContext** - - Defer() - show loading indicator for longer operations # Token Lifetime The interaction token expires after 15 minutes. After expiration, Followup(), Edit(), and Delete() calls will fail. # Modal Submissions For modal submissions, use ModalValue(customID) to retrieve text input values.
- **ComponentHandlerFunc** - ComponentHandlerFunc is the signature for component handler functions.
- **ComponentMiddleware** - ComponentMiddleware wraps a component handler function.
- **Context** - if err := ctx.Defer(); err != nil { return err } // Perform long operation (up to 15 minutes) result := performLongOperation() // Edit the deferred response return ctx.Reply(result) }
- **Embed** - Embed represents a Discord embed.
- **EmbedAuthor** - 
- **EmbedField** - 
- **EmbedFooter** - 
- **EmbedImage** - 
- **EmbedProvider** - 
- **EmbedThumbnail** - 
- **EmbedVideo** - 
- **HandlerFunc** - HandlerFunc is the signature for command handler functions.
- **Interaction** - Interaction represents a Discord interaction.
- **InteractionCallbackType** - InteractionCallbackType represents the type of interaction response.
- **InteractionData** - InteractionData contains the data for an interaction.
- **InteractionMember** - InteractionMember represents a guild member in an interaction context. This differs from models.GuildMember as it includes computed permissions.
- **InteractionOption** - InteractionOption represents an option in an interaction.
- **InteractionResponse** - InteractionResponse represents a response to an interaction.
- **InteractionResponseData** - InteractionResponseData contains the data for an interaction response.
- **InteractionType** - InteractionType represents the type of interaction.
- **MessageCommand** - MessageCommand represents a message context menu command.
- **MessageCreate** - MessageCreate represents a new message to send.
- **MessageEdit** - MessageEdit represents edits to a message.
- **MessageFlags** - MessageFlags represents message flags.
- **Middleware** - Middleware wraps a handler function to provide cross-cutting concerns.
- **MiddlewareProvider** - MiddlewareProvider is an optional interface for commands that have their own middleware.
- **Option** - Option represents a command option.
- **OptionBuilder** - OptionBuilder builds command options with a fluent API.
- **OptionMap** - OptionMap provides type-safe access to command options.
- **OptionProvider** - OptionProvider is an optional interface for commands that have options.
- **OptionType** - OptionType represents the type of a command option.
- **ResolvedData** - ResolvedData contains resolved entities from interaction options.
- **ResolvedInteractionData** - ResolvedInteractionData contains resolved entities.
- **Responder** - Responder is an interface for responding to interactions. This is implemented by the Bot to avoid circular imports.
- **Router** - Router handles routing interactions to commands and components.
- **ServiceRegistry** - ServiceRegistry holds registered services for dependency injection.
- **SlashCommand** - SlashCommand represents a slash command with typed options.
- **SubcommandGroup** - SubcommandGroup represents a group of subcommands.
- **SubcommandProvider** - SubcommandProvider is an optional interface for commands that have subcommands.
- **SyncOptions** - SyncOptions configures command synchronization.
- **SyncResult** - SyncResult contains the result of a sync operation.
- **UserCommand** - UserCommand represents a user context menu command.

#### Available Functions
- **ComponentService** - ComponentService retrieves a typed service from a ComponentContext. Returns the zero value if the service is not registered.
- **Service** - Service retrieves a typed service from the context. Returns the zero value if the service is not registered. Example: users := commands.Service[*UserService](ctx) if users == nil { return ctx.ReplyEphemeral("Service unavailable") }

For detailed API documentation, see the [commands API Reference](../api-reference/commands.md).

### Using discord

**Import Path:** `github.com/kolosys/discord`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord"
)

func main() {
    // Example usage of discord
    fmt.Println("Using discord package")
}
```

#### Available Types
- **Bot** - Bot is a unified Discord bot that combines: - Discord Gateway (real-time events via WebSocket) - Discord REST client (API calls) - Helix HTTP server (webhooks, interactions, admin API) [optional] - State cache (automatic entity caching) [enabled by default]
- **GuildCreateEvent** - 
- **InteractionCreateEvent** - 
- **MessageCreateEvent** - 
- **Module** - // Register HTTP routes api := bot.Group("/api/moderation") api.GET("/warnings/{userID}", m.handleGetWarnings) api.POST("/warnings", m.handleCreateWarning) return nil } See examples/modular for a complete example.
- **Options** - Options configures the Discord bot.
- **ReadyEvent** - Event type aliases for convenience
- **Snowflake** - Exports for developer convenience

#### Available Functions
- **Listen** - Listen is a convenience helper that registers a typed handler for an event. Example: discord.Listen(bot, events.MessageCreate, func(ctx context.Context, e *events.MessageCreateEvent) { fmt.Printf("Message: %s\n", e.Content) })
- **ListenRaw** - ListenRaw registers a handler that receives raw JSON data for an event.
- **MustRegisterModules** - MustRegisterModules is like RegisterModules but panics on error. Useful for application startup where registration errors are fatal. Example: discord.MustRegisterModules(bot, &utils.Module{}, &moderation.Module{}, &leveling.Module{}, )
- **OnEvent** - OnEvent registers a typed event handler using generics.
- **RegisterModules** - RegisterModules registers multiple modules with the bot. Modules are registered in the order they are provided. If any module returns an error, registration stops and that error is returned. Example: err := discord.RegisterModules(bot, &utils.Module{}, &moderation.Module{}, &leveling.Module{}, &admin.Module{}, ) if err != nil { return fmt.Errorf("failed to register modules: %w", err) }

For detailed API documentation, see the [discord API Reference](../api-reference/discord.md).

### Using components

**Import Path:** `github.com/kolosys/discord/components`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/components"
)

func main() {
    // Example usage of components
    fmt.Println("Using components package")
}
```

#### Available Types
- **ActionRowComponent** - ActionRowComponent is a container for other components.
- **ButtonBuilder** - ButtonBuilder builds buttons with a fluent API.
- **ButtonComponent** - ButtonComponent represents a clickable button.
- **ButtonEmoji** - ButtonEmoji represents an emoji on a button.
- **ChannelSelectBuilder** - ChannelSelectBuilder builds channel select menus.
- **ChannelSelectComponent** - ChannelSelectComponent represents a select menu for channels.
- **Component** - Component is the interface for all message components.
- **MentionableSelectBuilder** - MentionableSelectBuilder builds mentionable select menus.
- **MentionableSelectComponent** - MentionableSelectComponent represents a select menu for mentionable entities.
- **ModalBuilder** - ModalBuilder builds modals with a fluent API.
- **ModalData** - ModalData represents the data for a modal response.
- **ModalTextInputBuilder** - ModalTextInputBuilder is a builder for text inputs within a modal.
- **RoleSelectBuilder** - RoleSelectBuilder builds role select menus.
- **RoleSelectComponent** - RoleSelectComponent represents a select menu for roles.
- **SelectOption** - SelectOption represents an option in a string select menu.
- **StringSelectBuilder** - StringSelectBuilder builds string select menus.
- **StringSelectComponent** - StringSelectComponent represents a dropdown with string options.
- **TextInputBuilder** - TextInputBuilder builds text inputs with a fluent API.
- **TextInputComponent** - TextInputComponent represents a text input in a modal.
- **UserSelectBuilder** - UserSelectBuilder builds user select menus.
- **UserSelectComponent** - UserSelectComponent represents a select menu for users.

#### Available Functions
- **BuildComponents** - BuildComponents takes individual components and wraps them in action rows. Buttons are grouped together (max 5 per row), select menus get their own rows.
- **BuildComponentsSafe** - BuildComponentsSafe takes individual components and wraps them in action rows, returning an error if Discord limits are exceeded.
- **ValidateActionRowCount** - ValidateActionRowCount checks if the number of action rows is within limits.
- **ValidateCustomID** - ValidateCustomID checks if a custom_id is within the allowed length.
- **ValidateModalInputCount** - ValidateModalInputCount checks if the number of modal inputs is within limits.
- **ValidateSelectOptionCount** - ValidateSelectOptionCount checks if the number of select options is within limits.

For detailed API documentation, see the [components API Reference](../api-reference/components.md).

### Using embed

**Import Path:** `github.com/kolosys/discord/embed`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/embed"
)

func main() {
    // Example usage of embed
    fmt.Println("Using embed package")
}
```

#### Available Types
- **Builder** - Builder provides a fluent API for constructing Discord embeds.

For detailed API documentation, see the [embed API Reference](../api-reference/embed.md).

### Using events

**Import Path:** `github.com/kolosys/discord/events`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/events"
)

func main() {
    // Example usage of events
    fmt.Println("Using events package")
}
```

#### Available Types
- **Activity** - Activity represents a user activity.
- **Base** - Base contains common fields for all events.
- **ChannelCreateEvent** - ChannelCreateEvent is dispatched when a channel is created.
- **ChannelDeleteEvent** - ChannelDeleteEvent is dispatched when a channel is deleted.
- **ChannelUpdateEvent** - ChannelUpdateEvent is dispatched when a channel is updated.
- **ClientStatus** - ClientStatus represents a user's client status across platforms.
- **Dispatcher** - Dispatcher manages event subscriptions and dispatching.
- **GuildBanAddEvent** - GuildBanAddEvent is dispatched when a user is banned from a guild.
- **GuildBanRemoveEvent** - GuildBanRemoveEvent is dispatched when a user is unbanned from a guild.
- **GuildCreateEvent** - GuildCreateEvent is dispatched when a guild becomes available or the bot joins a new guild.
- **GuildDeleteEvent** - GuildDeleteEvent is dispatched when a guild becomes unavailable or the bot leaves/is removed.
- **GuildMemberAddEvent** - GuildMemberAddEvent is dispatched when a user joins a guild.
- **GuildMemberRemoveEvent** - GuildMemberRemoveEvent is dispatched when a user leaves or is removed from a guild.
- **GuildMemberUpdateEvent** - GuildMemberUpdateEvent is dispatched when a guild member is updated.
- **GuildMembersChunkEvent** - GuildMembersChunkEvent is dispatched in response to RequestGuildMembers.
- **GuildRoleCreateEvent** - GuildRoleCreateEvent is dispatched when a role is created.
- **GuildRoleDeleteEvent** - GuildRoleDeleteEvent is dispatched when a role is deleted.
- **GuildRoleUpdateEvent** - GuildRoleUpdateEvent is dispatched when a role is updated.
- **GuildUpdateEvent** - GuildUpdateEvent is dispatched when a guild is updated.
- **Handler** - Handler is a type-safe event handler function.
- **InteractionCreateEvent** - InteractionCreateEvent is dispatched when a user interacts with the application.
- **InviteCreateEvent** - InviteCreateEvent is dispatched when an invite is created.
- **InviteDeleteEvent** - InviteDeleteEvent is dispatched when an invite is deleted.
- **MessageCreateEvent** - MessageCreateEvent is dispatched when a message is created.
- **MessageDeleteBulkEvent** - MessageDeleteBulkEvent is dispatched when multiple messages are deleted at once.
- **MessageDeleteEvent** - MessageDeleteEvent is dispatched when a message is deleted.
- **MessageReactionAddEvent** - MessageReactionAddEvent is dispatched when a reaction is added to a message.
- **MessageReactionRemoveEvent** - MessageReactionRemoveEvent is dispatched when a reaction is removed from a message.
- **MessageUpdateEvent** - MessageUpdateEvent is dispatched when a message is updated.
- **PartialApplication** - PartialApplication contains partial application info.
- **PresenceUpdateEvent** - PresenceUpdateEvent is dispatched when a user's presence is updated.
- **RawHandler** - RawHandler handles raw event data without type conversion.
- **ReactionEmoji** - ReactionEmoji represents an emoji in a reaction.
- **ReadyEvent** - ReadyEvent is dispatched when the client has completed the initial handshake.
- **Type** - Type represents a Discord gateway event type.
- **TypingStartEvent** - TypingStartEvent is dispatched when a user starts typing.
- **UnavailableGuild** - UnavailableGuild represents a guild that is initially unavailable.
- **VoiceServerUpdateEvent** - VoiceServerUpdateEvent is dispatched when a guild's voice server is updated.
- **VoiceStateUpdateEvent** - VoiceStateUpdateEvent is dispatched when a user's voice state changes.

#### Available Functions
- **On** - On registers a type-safe handler for a specific event type. The handler will receive the event data already deserialized into the correct type. Context is propagated from the event bus for cancellation and deadline support.

For detailed API documentation, see the [events API Reference](../api-reference/events.md).

### Using bot

**Import Path:** `github.com/kolosys/discord/examples/basic/bot`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/examples/basic/bot"
)

func main() {
    // Example usage of bot
    fmt.Println("Using bot package")
}
```

#### Available Functions
- **New** - 

For detailed API documentation, see the [bot API Reference](../api-reference/bot.md).

### Using commands

**Import Path:** `github.com/kolosys/discord/examples/basic/commands`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/examples/basic/commands"
)

func main() {
    // Example usage of commands
    fmt.Println("Using commands package")
}
```

#### Available Functions
- **Register** - 

For detailed API documentation, see the [commands API Reference](../api-reference/commands.md).

### Using events

**Import Path:** `github.com/kolosys/discord/examples/basic/events`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/examples/basic/events"
)

func main() {
    // Example usage of events
    fmt.Println("Using events package")
}
```

#### Available Functions
- **Register** - 

For detailed API documentation, see the [events API Reference](../api-reference/events.md).

### Using routes

**Import Path:** `github.com/kolosys/discord/examples/basic/routes`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/examples/basic/routes"
)

func main() {
    // Example usage of routes
    fmt.Println("Using routes package")
}
```

#### Available Types
- **GatewayStatus** - 
- **HealthResponse** - 
- **StatusResponse** - 

#### Available Functions
- **Register** - 

For detailed API documentation, see the [routes API Reference](../api-reference/routes.md).

### Using internal

**Import Path:** `github.com/kolosys/discord/examples/internal`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/examples/internal"
)

func main() {
    // Example usage of internal
    fmt.Println("Using internal package")
}
```

#### Available Functions
- **LoadEnv** - 

For detailed API documentation, see the [internal API Reference](../api-reference/internal.md).

### Using gateway

**Import Path:** `github.com/kolosys/discord/gateway`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/gateway"
)

func main() {
    // Example usage of gateway
    fmt.Println("Using gateway package")
}
```

#### Available Types
- **Activity** - Activity represents a user activity.
- **ActivityType** - ActivityType represents the type of activity.
- **CloseCode** - CloseCode represents a Discord gateway close code. For standard WebSocket codes (1000-1015), use axon.CloseCode directly. This type focuses on Discord-specific codes (4000+) with detailed descriptions.
- **Gateway** - Gateway manages Discord gateway connections across multiple shards.
- **GatewayClient** - GatewayClient defines the interface for WebSocket client operations required by Shard. It enables dependency injection for testing with axon/mock while maintaining type safety.
- **GatewayPayload** - GatewayPayload represents a Discord gateway message.
- **HeartbeatManager** - HeartbeatManager manages the heartbeat loop for a Discord gateway connection.
- **HelloData** - HelloData is the data sent in the Hello event (Opcode 10).
- **IdentifyData** - IdentifyData is sent to identify a new session.
- **IdentifyProperties** - IdentifyProperties contains connection properties for identification.
- **Intent** - Intent represents a Discord gateway intent bitfield.
- **InvalidSessionData** - InvalidSessionData indicates whether the session can be resumed.
- **Opcode** - Opcode represents a Discord gateway opcode.
- **Options** - Options configures the Gateway with optional dependency overrides.
- **PartialApplication** - PartialApplication contains partial application info from Ready.
- **PresenceUpdate** - PresenceUpdate represents a presence update payload.
- **ReadyData** - ReadyData is the data sent in the Ready event after successful identification.
- **RequestGuildMembersData** - RequestGuildMembersData is sent to request guild members.
- **ResumeData** - ResumeData is sent to resume a disconnected session.
- **Session** - Session manages the state of a Discord gateway session.
- **Shard** - Shard represents a single Discord gateway shard connection.
- **UnavailableGuild** - UnavailableGuild represents a guild that is initially unavailable in the Ready event.
- **VoiceStateUpdateData** - VoiceStateUpdateData is sent to join/leave/move voice channels.

#### Available Functions
- **Listen** - Listen is a convenience helper that registers a typed handler for an event. Example: gateway.Listen(gw, events.MessageCreate, func(ctx context.Context, e *events.MessageCreateEvent) { fmt.Printf("Message: %s\n", e.Content) })
- **ListenRaw** - ListenRaw registers a handler that receives raw JSON data for an event.
- **OnEvent** - OnEvent registers a typed event handler using generics.

For detailed API documentation, see the [gateway API Reference](../api-reference/gateway.md).

### Using models

**Import Path:** `github.com/kolosys/discord/models`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/models"
)

func main() {
    // Example usage of models
    fmt.Println("Using models package")
}
```

#### Available Types
- **Account** - Account
- **ActionRowComponent** - ActionRowComponent
- **ActionRowComponentForMessageOptions** - ActionRowComponentForMessageOptions
- **ActionRowComponentForModalOptions** - ActionRowComponentForModalOptions
- **ActivitiesAttachment** - ActivitiesAttachment
- **AfkTimeouts** - AfkTimeouts
- **AllowedMentionTypes** - AllowedMentionTypes
- **Application** - Application
- **ApplicationCommand** - ApplicationCommand
- **ApplicationCommandAttachmentOption** - ApplicationCommandAttachmentOption
- **ApplicationCommandAutocompleteCallbackOptions** - ApplicationCommandAutocompleteCallbackOptions
- **ApplicationCommandBooleanOption** - ApplicationCommandBooleanOption
- **ApplicationCommandChannelOption** - ApplicationCommandChannelOption
- **ApplicationCommandCreateOptions** - ApplicationCommandCreateOptions
- **ApplicationCommandHandler** - ApplicationCommandHandler
- **ApplicationCommandIntegerOption** - ApplicationCommandIntegerOption
- **ApplicationCommandInteractionMetadata** - ApplicationCommandInteractionMetadata
- **ApplicationCommandMentionableOption** - ApplicationCommandMentionableOption
- **ApplicationCommandNumberOption** - ApplicationCommandNumberOption
- **ApplicationCommandOptionIntegerChoice** - ApplicationCommandOptionIntegerChoice
- **ApplicationCommandOptionNumberChoice** - ApplicationCommandOptionNumberChoice
- **ApplicationCommandOptionStringChoice** - ApplicationCommandOptionStringChoice
- **ApplicationCommandOptionType** - ApplicationCommandOptionType
- **ApplicationCommandPatchOptionsPartial** - ApplicationCommandPatchOptionsPartial
- **ApplicationCommandPermission** - ApplicationCommandPermission
- **ApplicationCommandPermissionType** - ApplicationCommandPermissionType
- **ApplicationCommandRoleOption** - ApplicationCommandRoleOption
- **ApplicationCommandStringOption** - ApplicationCommandStringOption
- **ApplicationCommandSubcommandGroupOption** - ApplicationCommandSubcommandGroupOption
- **ApplicationCommandSubcommandOption** - ApplicationCommandSubcommandOption
- **ApplicationCommandType** - ApplicationCommandType
- **ApplicationCommandUpdateOptions** - ApplicationCommandUpdateOptions
- **ApplicationCommandUserOption** - ApplicationCommandUserOption
- **ApplicationExplicitContentFilterTypes** - ApplicationExplicitContentFilterTypes
- **ApplicationFormPartial** - ApplicationFormPartial
- **ApplicationIdentityProviderAuthType** - ApplicationIdentityProviderAuthType
- **ApplicationIncomingWebhook** - ApplicationIncomingWebhook
- **ApplicationIntegrationType** - ApplicationIntegrationType
- **ApplicationIntegrationTypeConfiguration** - ApplicationIntegrationTypeConfiguration
- **ApplicationOAuth2InstallParams** - ApplicationOAuth2InstallParams
- **ApplicationRoleConnectionsMetadataItem** - ApplicationRoleConnectionsMetadataItem
- **ApplicationRoleConnectionsMetadataItemOptions** - ApplicationRoleConnectionsMetadataItemOptions
- **ApplicationTypes** - ApplicationTypes
- **ApplicationUserRoleConnection** - ApplicationUserRoleConnection
- **Attachment** - Attachment
- **AuditLogActionTypes** - AuditLogActionTypes
- **AuditLogEntry** - AuditLogEntry
- **AuditLogObjectChange** - AuditLogObjectChange
- **AutomodActionType** - AutomodActionType
- **AutomodEventType** - AutomodEventType
- **AutomodKeywordPresetType** - AutomodKeywordPresetType
- **AutomodTriggerType** - AutomodTriggerType
- **AvailableLocalesEnum** - AvailableLocalesEnum
- **BanUserFromGuildOptions** - BanUserFromGuildOptions
- **BaseCreateMessageCreateOptions** - BaseCreateMessageCreateOptions
- **BasicApplication** - BasicApplication
- **BasicGuildMember** - BasicGuildMember
- **BasicMessage** - BasicMessage
- **BlockMessageAction** - BlockMessageAction
- **BlockMessageActionMetadata** - BlockMessageActionMetadata
- **BotAccountPatchOptions** - BotAccountPatchOptions
- **BotAddGuildMemberOptions** - BotAddGuildMemberOptions
- **BulkBanUsers** - BulkBanUsers
- **BulkBanUsersOptions** - BulkBanUsersOptions
- **BulkLobbyMemberOptions** - BulkLobbyMemberOptions
- **ButtonComponent** - ButtonComponent
- **ButtonComponentForMessageOptions** - ButtonComponentForMessageOptions
- **ButtonStyleTypes** - ButtonStyleTypes
- **ChannelFollower** - ChannelFollower
- **ChannelFollowerWebhook** - ChannelFollowerWebhook
- **ChannelPermissionOverwrite** - ChannelPermissionOverwrite
- **ChannelPermissionOverwriteOptions** - ChannelPermissionOverwriteOptions
- **ChannelPermissionOverwrites** - ChannelPermissionOverwrites
- **ChannelSelectComponent** - ChannelSelectComponent
- **ChannelSelectComponentForMessageOptions** - ChannelSelectComponentForMessageOptions
- **ChannelSelectComponentForModalOptions** - ChannelSelectComponentForModalOptions
- **ChannelSelectDefaultValue** - ChannelSelectDefaultValue
- **ChannelTypes** - ChannelTypes
- **CommandPermission** - CommandPermission
- **CommandPermissions** - CommandPermissions
- **ComponentEmoji** - ComponentEmoji
- **ComponentEmojiForOptions** - ComponentEmojiForOptions
- **ConfettiPotionCreateOptions** - ConfettiPotionCreateOptions
- **ConnectedAccount** - ConnectedAccount
- **ConnectedAccountGuild** - ConnectedAccountGuild
- **ConnectedAccountIntegration** - ConnectedAccountIntegration
- **ConnectedAccountProviders** - ConnectedAccountProviders
- **ConnectedAccountVisibility** - ConnectedAccountVisibility
- **ContainerComponent** - ContainerComponent
- **ContainerComponentForMessageOptions** - ContainerComponentForMessageOptions
- **CreateEntitlementOptionsData** - CreateEntitlementOptionsData
- **CreateForumThreadOptions** - CreateForumThreadOptions
- **CreateGroupDMInviteOptions** - CreateGroupDMInviteOptions
- **CreateGuildChannelOptions** - CreateGuildChannelOptions
- **CreateGuildInviteOptions** - CreateGuildInviteOptions
- **CreateMessageInteractionCallback** - CreateMessageInteractionCallback
- **CreateMessageInteractionCallbackOptions** - CreateMessageInteractionCallbackOptions
- **CreateOrUpdateThreadTagOptions** - CreateOrUpdateThreadTagOptions
- **CreatePrivateChannelOptions** - CreatePrivateChannelOptions
- **CreateRoleOptions** - CreateRoleOptions
- **CreateTextThreadWithMessageOptions** - CreateTextThreadWithMessageOptions
- **CreateTextThreadWithoutMessageOptions** - CreateTextThreadWithoutMessageOptions
- **CreatedThread** - CreatedThread
- **CustomClientTheme** - CustomClientTheme
- **CustomClientThemeShareOptions** - CustomClientThemeShareOptions
- **DefaultKeywordListTriggerMetadata** - DefaultKeywordListTriggerMetadata
- **DefaultKeywordListUpsertOptions** - DefaultKeywordListUpsertOptions
- **DefaultKeywordListUpsertOptionsPartial** - DefaultKeywordListUpsertOptionsPartial
- **DefaultKeywordRule** - DefaultKeywordRule
- **DefaultReactionEmoji** - DefaultReactionEmoji
- **DiscordIntegration** - DiscordIntegration
- **EmbeddedActivityInstance** - EmbeddedActivityInstance
- **EmbeddedActivityLocationKind** - EmbeddedActivityLocationKind
- **Emoji** - Emoji
- **Entitlement** - Entitlement
- **EntitlementOwnerTypes** - EntitlementOwnerTypes
- **EntitlementTenantFulfillmentStatus** - EntitlementTenantFulfillmentStatus
- **EntitlementTypes** - EntitlementTypes
- **EntityMetadataExternal** - EntityMetadataExternal
- **EntityMetadataStageInstance** - EntityMetadataStageInstance
- **EntityMetadataVoice** - EntityMetadataVoice
- **ErrorDetails** - ErrorDetails
- **ErrorResponse** - Error A single error, either for an API response or a specific field.
- **ExternalConnectionIntegration** - ExternalConnectionIntegration
- **ExternalScheduledEvent** - ExternalScheduledEvent
- **ExternalScheduledEventCreateOptions** - ExternalScheduledEventCreateOptions
- **ExternalScheduledEventPatchOptionsPartial** - ExternalScheduledEventPatchOptionsPartial
- **FileComponent** - FileComponent
- **FileComponentForMessageOptions** - FileComponentForMessageOptions
- **FileUploadComponentForModalOptions** - FileUploadComponentForModalOptions
- **FlagToChannelAction** - FlagToChannelAction
- **FlagToChannelActionMetadata** - FlagToChannelActionMetadata
- **ForumLayout** - ForumLayout
- **ForumTag** - ForumTag
- **FriendInvite** - FriendInvite
- **Gateway** - Gateway
- **GatewayBot** - GatewayBot
- **GatewayBotSessionStartLimit** - GatewayBotSessionStartLimit
- **GithubAuthor** - GithubAuthor
- **GithubCheckApp** - GithubCheckApp
- **GithubCheckPullOptions** - GithubCheckPullOptions
- **GithubCheckRun** - GithubCheckRun
- **GithubCheckRunOutput** - GithubCheckRunOutput
- **GithubCheckSuite** - GithubCheckSuite
- **GithubComment** - GithubComment
- **GithubCommit** - GithubCommit
- **GithubDiscussion** - GithubDiscussion
- **GithubIssue** - GithubIssue
- **GithubRelease** - GithubRelease
- **GithubRepository** - GithubRepository
- **GithubReview** - GithubReview
- **GithubUser** - GithubUser
- **GithubWebhook** - GithubWebhook
- **GroupDMInvite** - GroupDMInvite
- **Guild** - Guild
- **GuildAuditLog** - GuildAuditLog
- **GuildBan** - GuildBan
- **GuildChannel** - GuildChannel
- **GuildChannelLocation** - GuildChannelLocation
- **GuildExplicitContentFilterTypes** - GuildExplicitContentFilterTypes
- **GuildFeatures** - GuildFeatures
- **GuildHomeSettings** - GuildHomeSettings
- **GuildIncomingWebhook** - GuildIncomingWebhook
- **GuildInvite** - GuildInvite
- **GuildMFALevel** - GuildMFALevel
- **GuildMember** - GuildMember
- **GuildNSFWContentLevel** - GuildNSFWContentLevel
- **GuildOnboarding** - GuildOnboarding
- **GuildOnboardingMode** - GuildOnboardingMode
- **GuildPatchOptionsPartial** - GuildPatchOptionsPartial
- **GuildPreview** - GuildPreview
- **GuildProductPurchase** - GuildProductPurchase
- **GuildPrune** - GuildPrune
- **GuildRole** - GuildRole
- **GuildRoleColors** - GuildRoleColors
- **GuildRoleTags** - GuildRoleTags
- **GuildScheduledEventEntityTypes** - GuildScheduledEventEntityTypes
- **GuildScheduledEventPrivacyLevels** - GuildScheduledEventPrivacyLevels
- **GuildScheduledEventStatuses** - GuildScheduledEventStatuses
- **GuildSticker** - GuildSticker
- **GuildSubscriptionIntegration** - GuildSubscriptionIntegration
- **GuildTemplate** - GuildTemplate
- **GuildTemplateChannel** - GuildTemplateChannel
- **GuildTemplateChannelTags** - GuildTemplateChannelTags
- **GuildTemplateRole** - GuildTemplateRole
- **GuildTemplateRoleColors** - GuildTemplateRoleColors
- **GuildTemplateSnapshot** - GuildTemplateSnapshot
- **GuildWelcomeChannel** - GuildWelcomeChannel
- **GuildWelcomeScreen** - GuildWelcomeScreen
- **GuildWelcomeScreenChannel** - GuildWelcomeScreenChannel
- **GuildWithCounts** - GuildWithCounts
- **IconEmoji** - IconEmoji
- **IncomingWebhookInteractionOptions** - IncomingWebhookInteractionOptions
- **IncomingWebhookOptionsPartial** - IncomingWebhookOptionsPartial
- **IncomingWebhookUpdateForInteractionCallbackOptionsPartial** - IncomingWebhookUpdateForInteractionCallbackOptionsPartial
- **IncomingWebhookUpdateOptionsPartial** - IncomingWebhookUpdateOptionsPartial
- **InnerErrors** - InnerErrors
- **Int53Type** - Int53Type
- **IntegrationApplication** - IntegrationApplication
- **IntegrationExpireBehaviorTypes** - IntegrationExpireBehaviorTypes
- **IntegrationExpireGracePeriodTypes** - IntegrationExpireGracePeriodTypes
- **IntegrationTypes** - IntegrationTypes
- **Interaction** - Interaction
- **InteractionApplicationCommandAutocompleteCallbackIntegerData** - InteractionApplicationCommandAutocompleteCallbackIntegerData
- **InteractionApplicationCommandAutocompleteCallbackNumberData** - InteractionApplicationCommandAutocompleteCallbackNumberData
- **InteractionApplicationCommandAutocompleteCallbackStringData** - InteractionApplicationCommandAutocompleteCallbackStringData
- **InteractionCallback** - InteractionCallback
- **InteractionCallbackTypes** - InteractionCallbackTypes
- **InteractionContextType** - InteractionContextType
- **InteractionTypes** - InteractionTypes
- **InviteApplication** - InviteApplication
- **InviteChannel** - InviteChannel
- **InviteChannelRecipient** - InviteChannelRecipient
- **InviteGuild** - InviteGuild
- **InviteTargetTypes** - InviteTargetTypes
- **InviteTypes** - InviteTypes
- **KeywordRule** - KeywordRule
- **KeywordTriggerMetadata** - KeywordTriggerMetadata
- **KeywordUpsertOptions** - KeywordUpsertOptions
- **KeywordUpsertOptionsPartial** - KeywordUpsertOptionsPartial
- **LabelComponentForModalOptions** - LabelComponentForModalOptions
- **LaunchActivityInteractionCallback** - LaunchActivityInteractionCallback
- **LaunchActivityInteractionCallbackOptions** - LaunchActivityInteractionCallbackOptions
- **ListApplicationEmojis** - ListApplicationEmojis
- **ListGuildSoundboardSounds** - ListGuildSoundboardSounds
- **Lobby** - Lobby
- **LobbyGuildInvite** - LobbyGuildInvite
- **LobbyMember** - LobbyMember
- **LobbyMemberOptions** - LobbyMemberOptions
- **LobbyMessage** - LobbyMessage
- **MLSpamRule** - MLSpamRule
- **MLSpamTriggerMetadata** - MLSpamTriggerMetadata
- **MLSpamUpsertOptions** - MLSpamUpsertOptions
- **MLSpamUpsertOptionsPartial** - MLSpamUpsertOptionsPartial
- **MediaGalleryComponent** - MediaGalleryComponent
- **MediaGalleryComponentForMessageOptions** - MediaGalleryComponentForMessageOptions
- **MediaGalleryItem** - MediaGalleryItem
- **MediaGalleryItemOptions** - MediaGalleryItemOptions
- **MentionSpamRule** - MentionSpamRule
- **MentionSpamTriggerMetadata** - MentionSpamTriggerMetadata
- **MentionSpamUpsertOptions** - MentionSpamUpsertOptions
- **MentionSpamUpsertOptionsPartial** - MentionSpamUpsertOptionsPartial
- **MentionableSelectComponent** - MentionableSelectComponent
- **MentionableSelectComponentForMessageOptions** - MentionableSelectComponentForMessageOptions
- **MentionableSelectComponentForModalOptions** - MentionableSelectComponentForModalOptions
- **Message** - Message
- **MessageActivity** - MessageActivity
- **MessageAllowedMentionsOptions** - MessageAllowedMentionsOptions
- **MessageAttachment** - MessageAttachment
- **MessageAttachmentOptions** - MessageAttachmentOptions
- **MessageCall** - MessageCall
- **MessageComponentInteractionMetadata** - MessageComponentInteractionMetadata
- **MessageComponentSeparatorSpacingSize** - MessageComponentSeparatorSpacingSize
- **MessageComponentTypes** - MessageComponentTypes
- **MessageCreateOptions** - MessageCreateOptions
- **MessageEditOptionsPartial** - MessageEditOptionsPartial
- **MessageEmbed** - MessageEmbed
- **MessageEmbedAuthor** - MessageEmbedAuthor
- **MessageEmbedField** - MessageEmbedField
- **MessageEmbedFooter** - MessageEmbedFooter
- **MessageEmbedImage** - MessageEmbedImage
- **MessageEmbedProvider** - MessageEmbedProvider
- **MessageEmbedVideo** - MessageEmbedVideo
- **MessageInteraction** - MessageInteraction
- **MessageMentionChannel** - MessageMentionChannel
- **MessageReaction** - MessageReaction
- **MessageReactionCountDetails** - MessageReactionCountDetails
- **MessageReactionEmoji** - MessageReactionEmoji
- **MessageReference** - MessageReference
- **MessageReferenceOptions** - MessageReferenceOptions
- **MessageReferenceType** - MessageReferenceType
- **MessageRoleSubscriptionData** - MessageRoleSubscriptionData
- **MessageShareCustomUserThemeBaseTheme** - MessageShareCustomUserThemeBaseTheme
- **MessageSnapshot** - MessageSnapshot
- **MessageStickerItem** - MessageStickerItem
- **MessageType** - MessageType
- **MetadataItemTypes** - MetadataItemTypes
- **MinimalContentMessage** - MinimalContentMessage
- **ModalInteractionCallbackOptions** - ModalInteractionCallbackOptions
- **ModalInteractionCallbackOptionsData** - ModalInteractionCallbackOptionsData
- **ModalSubmitInteractionMetadata** - ModalSubmitInteractionMetadata
- **MyGuild** - MyGuild
- **NameplatePalette** - NameplatePalette
- **NewMemberAction** - NewMemberAction
- **NewMemberActionType** - NewMemberActionType
- **OAuth2GetAuthorization** - OAuth2GetAuthorization
- **OAuth2GetKeys** - OAuth2GetKeys
- **OAuth2GetOpenIDConnectUserInfo** - OAuth2GetOpenIDConnectUserInfo
- **OAuth2Key** - OAuth2Key
- **OAuth2Scopes** - OAuth2Scopes
- **OnboardingPrompt** - OnboardingPrompt
- **OnboardingPromptOption** - OnboardingPromptOption
- **OnboardingPromptOptionOptions** - OnboardingPromptOptionOptions
- **OnboardingPromptType** - OnboardingPromptType
- **PartialDiscordIntegration** - PartialDiscordIntegration
- **PartialExternalConnectionIntegration** - PartialExternalConnectionIntegration
- **PartialGuildSubscriptionIntegration** - PartialGuildSubscriptionIntegration
- **PinnedMessage** - PinnedMessage
- **PinnedMessages** - PinnedMessages
- **Poll** - Poll
- **PollAnswer** - PollAnswer
- **PollAnswerCreateOptions** - PollAnswerCreateOptions
- **PollAnswerDetails** - PollAnswerDetails
- **PollCreateOptions** - PollCreateOptions
- **PollEmoji** - PollEmoji
- **PollEmojiCreateOptions** - PollEmojiCreateOptions
- **PollLayoutTypes** - PollLayoutTypes
- **PollMedia** - PollMedia
- **PollMediaCreateOptions** - PollMediaCreateOptions
- **PollResults** - PollResults
- **PollResultsEntry** - PollResultsEntry
- **PongInteractionCallbackOptions** - PongInteractionCallbackOptions
- **PremiumGuildTiers** - PremiumGuildTiers
- **PremiumTypes** - PremiumTypes
- **PrivateApplication** - PrivateApplication
- **PrivateChannel** - PrivateChannel
- **PrivateChannelLocation** - PrivateChannelLocation
- **PrivateGroupChannel** - PrivateGroupChannel
- **PrivateGuildMember** - PrivateGuildMember
- **ProvisionalToken** - ProvisionalToken
- **PruneGuildOptions** - PruneGuildOptions
- **PurchaseNotification** - PurchaseNotification
- **PurchaseType** - PurchaseType
- **QuarantineUserAction** - QuarantineUserAction
- **QuarantineUserActionMetadata** - QuarantineUserActionMetadata
- **RatelimitedResponse** - Ratelimited Ratelimit error object returned by the Discord API
- **ReactionTypes** - ReactionTypes
- **ResolvedObjects** - ResolvedObjects
- **ResourceChannel** - ResourceChannel
- **RichEmbed** - RichEmbed
- **RichEmbedAuthor** - RichEmbedAuthor
- **RichEmbedField** - RichEmbedField
- **RichEmbedFooter** - RichEmbedFooter
- **RichEmbedImage** - RichEmbedImage
- **RichEmbedProvider** - RichEmbedProvider
- **RichEmbedThumbnail** - RichEmbedThumbnail
- **RichEmbedVideo** - RichEmbedVideo
- **RoleSelectComponent** - RoleSelectComponent
- **RoleSelectComponentForMessageOptions** - RoleSelectComponentForMessageOptions
- **RoleSelectComponentForModalOptions** - RoleSelectComponentForModalOptions
- **RoleSelectDefaultValue** - RoleSelectDefaultValue
- **SDKMessageOptions** - SDKMessageOptions
- **ScheduledEvent** - ScheduledEvent
- **ScheduledEventUser** - ScheduledEventUser
- **SectionComponent** - SectionComponent
- **SectionComponentForMessageOptions** - SectionComponentForMessageOptions
- **SeparatorComponent** - SeparatorComponent
- **SeparatorComponentForMessageOptions** - SeparatorComponentForMessageOptions
- **SettingsEmoji** - SettingsEmoji
- **SlackWebhook** - SlackWebhook
- **SnowflakeSelectDefaultValueTypes** - SnowflakeSelectDefaultValueTypes
- **SnowflakeType** - SnowflakeType is a Discord snowflake ID (just a string).
- **SortingOrder** - SortingOrder
- **SoundboardCreateOptions** - SoundboardCreateOptions
- **SoundboardPatchOptionsPartial** - SoundboardPatchOptionsPartial
- **SoundboardSound** - SoundboardSound
- **SoundboardSoundSendOptions** - SoundboardSoundSendOptions
- **SpamLinkRule** - SpamLinkRule
- **SpamLinkTriggerMetadata** - SpamLinkTriggerMetadata
- **StageInstance** - StageInstance
- **StageInstancesPrivacyLevels** - StageInstancesPrivacyLevels
- **StageScheduledEvent** - StageScheduledEvent
- **StageScheduledEventCreateOptions** - StageScheduledEventCreateOptions
- **StageScheduledEventPatchOptionsPartial** - StageScheduledEventPatchOptionsPartial
- **StandardSticker** - StandardSticker
- **StickerFormatTypes** - StickerFormatTypes
- **StickerPack** - StickerPack
- **StickerPackCollection** - StickerPackCollection
- **StickerTypes** - StickerTypes
- **StringSelectComponent** - StringSelectComponent
- **StringSelectComponentForMessageOptions** - StringSelectComponentForMessageOptions
- **StringSelectComponentForModalOptions** - StringSelectComponentForModalOptions
- **StringSelectOption** - StringSelectOption
- **StringSelectOptionForOptions** - StringSelectOptionForOptions
- **Team** - Team
- **TeamMember** - TeamMember
- **TeamMembershipStates** - TeamMembershipStates
- **TextDisplayComponent** - TextDisplayComponent
- **TextDisplayComponentForMessageOptions** - TextDisplayComponentForMessageOptions
- **TextDisplayComponentForModalOptions** - TextDisplayComponentForModalOptions
- **TextInputComponent** - TextInputComponent
- **TextInputComponentForModalOptions** - TextInputComponentForModalOptions
- **TextInputStyleTypes** - TextInputStyleTypes
- **Thread** - Thread
- **ThreadAutoArchiveDuration** - ThreadAutoArchiveDuration
- **ThreadMember** - ThreadMember
- **ThreadMetadata** - ThreadMetadata
- **ThreadSearch** - ThreadSearch
- **ThreadSearchTagSetting** - ThreadSearchTagSetting
- **ThreadSortOrder** - ThreadSortOrder
- **ThreadSortingMode** - ThreadSortingMode
- **Threads** - Threads
- **ThumbnailComponent** - ThumbnailComponent
- **ThumbnailComponentForMessageOptions** - ThumbnailComponentForMessageOptions
- **TypingIndicator** - TypingIndicator
- **UInt32Type** - UInt32Type
- **UnbanUserFromGuildOptions** - UnbanUserFromGuildOptions
- **UnfurledMedia** - UnfurledMedia
- **UnfurledMediaOptions** - UnfurledMediaOptions
- **UnfurledMediaOptionsWithAttachmentReferenceRequired** - UnfurledMediaOptionsWithAttachmentReferenceRequired
- **UpdateApplicationUserRoleConnectionOptions** - UpdateApplicationUserRoleConnectionOptions
- **UpdateDMOptionsPartial** - UpdateDMOptionsPartial
- **UpdateDefaultReactionEmojiOptions** - UpdateDefaultReactionEmojiOptions
- **UpdateGroupDMOptionsPartial** - UpdateGroupDMOptionsPartial
- **UpdateGuildChannelOptionsPartial** - UpdateGuildChannelOptionsPartial
- **UpdateGuildOnboardingOptions** - UpdateGuildOnboardingOptions
- **UpdateMessageInteractionCallback** - UpdateMessageInteractionCallback
- **UpdateMessageInteractionCallbackOptions** - UpdateMessageInteractionCallbackOptions
- **UpdateOnboardingPromptOptions** - UpdateOnboardingPromptOptions
- **UpdateRoleOptionsPartial** - UpdateRoleOptionsPartial
- **UpdateRolePositionsOptions** - UpdateRolePositionsOptions
- **UpdateSelfVoiceStateOptionsPartial** - UpdateSelfVoiceStateOptionsPartial
- **UpdateThreadOptionsPartial** - UpdateThreadOptionsPartial
- **UpdateThreadTagOptions** - UpdateThreadTagOptions
- **UpdateVoiceStateOptionsPartial** - UpdateVoiceStateOptionsPartial
- **User** - User
- **UserAvatarDecoration** - UserAvatarDecoration
- **UserCollectibles** - UserCollectibles
- **UserCommunicationDisabledAction** - UserCommunicationDisabledAction
- **UserCommunicationDisabledActionMetadata** - UserCommunicationDisabledActionMetadata
- **UserGuildOnboarding** - UserGuildOnboarding
- **UserNameplate** - UserNameplate
- **UserNotificationSettings** - UserNotificationSettings
- **UserPII** - UserPII
- **UserPrimaryGuild** - UserPrimaryGuild
- **UserSelectComponent** - UserSelectComponent
- **UserSelectComponentForMessageOptions** - UserSelectComponentForMessageOptions
- **UserSelectComponentForModalOptions** - UserSelectComponentForModalOptions
- **UserSelectDefaultValue** - UserSelectDefaultValue
- **VanityURL** - VanityURL
- **VanityURLError** - VanityURLError
- **VerificationLevels** - VerificationLevels
- **VideoQualityModes** - VideoQualityModes
- **VoiceRegion** - VoiceRegion
- **VoiceScheduledEvent** - VoiceScheduledEvent
- **VoiceScheduledEventCreateOptions** - VoiceScheduledEventCreateOptions
- **VoiceScheduledEventPatchOptionsPartial** - VoiceScheduledEventPatchOptionsPartial
- **VoiceState** - VoiceState
- **WebhookSlackEmbed** - WebhookSlackEmbed
- **WebhookSlackEmbedField** - WebhookSlackEmbedField
- **WebhookSourceChannel** - WebhookSourceChannel
- **WebhookSourceGuild** - WebhookSourceGuild
- **WebhookTypes** - WebhookTypes
- **WelcomeMessage** - WelcomeMessage
- **WelcomeScreenPatchOptionsPartial** - WelcomeScreenPatchOptionsPartial
- **Widget** - Widget
- **WidgetActivity** - WidgetActivity
- **WidgetChannel** - WidgetChannel
- **WidgetImageStyles** - WidgetImageStyles
- **WidgetMember** - WidgetMember
- **WidgetSettings** - WidgetSettings
- **WidgetUserDiscriminator** - WidgetUserDiscriminator

For detailed API documentation, see the [models API Reference](../api-reference/models.md).

### Using client

**Import Path:** `github.com/kolosys/discord/rest/internal`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/rest/internal"
)

func main() {
    // Example usage of client
    fmt.Println("Using client package")
}
```

#### Available Types
- **ClientErrorResponseResponse** - ClientErrorResponseResponse wraps the ClientErrorResponse response
- **ClientOptions** - ClientOptions holds client configuration
- **ClientRatelimitedResponseResponse** - ClientRatelimitedResponseResponse wraps the ClientRatelimitedResponse response
- **EmptyResponse** - Response type aliases for common responses
- **HeaderValues** - HeaderValues holds typed header values
- **RESTClient** - RESTClient wraps a neuron.Client with generated route methods
- **RequestBuilder** - RequestBuilder helps build requests with context
- **SecurityScheme** - SecurityScheme defines a security scheme

#### Available Functions
- **ApplyAPIKey** - ApplyAPIKey adds an API key to the request
- **ApplyBasicAuth** - ApplyBasicAuth adds basic authentication to the request
- **ApplyBearerToken** - ApplyBearerToken adds a bearer token to the request headers

For detailed API documentation, see the [client API Reference](../api-reference/client.md).

### Using rest

**Import Path:** `github.com/kolosys/discord/rest`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/rest"
)

func main() {
    // Example usage of rest
    fmt.Println("Using rest package")
}
```

#### Available Types
- **DiscordRateLimiter** - DiscordRateLimiter adapts Ion's MultiTierLimiter to Neuron's RateLimitHandler interface for Discord's bucket-based rate limiting system.
- **ListMessagesParams** - ListMessagesParams configures ListMessages requests.
- **Options** - Options holds REST client configuration.
- **REST** - REST provides access to the Discord REST API with rate limiting and authentication.
- **RateLimiterConfig** - RateLimiterConfig configures the Discord rate limiter.

For detailed API documentation, see the [rest API Reference](../api-reference/rest.md).

### Using server

**Import Path:** `github.com/kolosys/discord/server`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/server"
)

func main() {
    // Example usage of server
    fmt.Println("Using server package")
}
```

#### Available Types
- **Manager** - 
- **Options** - Options configures the Discord bot's HTTP server and lifecycle management. All fields are optional and will use sensible defaults if not specified. For HTTP-disabled bots, only ConnectFn and StopFn are required. For HTTP-enabled bots, the server defaults to listening on :8080.

For detailed API documentation, see the [server API Reference](../api-reference/server.md).

### Using state

**Import Path:** `github.com/kolosys/discord/state`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/state"
)

func main() {
    // Example usage of state
    fmt.Println("Using state package")
}
```

#### Available Types
- **CachedGuild** - CachedGuild extends Guild with additional cached data from GUILD_CREATE.
- **CachedMember** - CachedMember wraps GuildMember with guild and user references.
- **ChannelStore** - ChannelStore manages channel cache.
- **GuildStore** - GuildStore manages guild cache with relationship methods.
- **MemberStore** - MemberStore manages member cache with composite keys.
- **Options** - Options configures state caching behavior.
- **Presence** - Presence represents a user's presence state in a guild.
- **PresenceStore** - PresenceStore manages presence cache.
- **RoleStore** - RoleStore manages role cache with composite keys.
- **State** - State manages Discord entity caches with automatic population from gateway events.
- **StateStats** - StateStats contains aggregated statistics from all caches.
- **Store** - Store is a generic entity cache with standard CRUD operations.
- **UserStore** - UserStore manages user cache.
- **VoiceState** - VoiceState represents a user's voice connection state in a guild.
- **VoiceStateStore** - VoiceStateStore manages voice state cache.

For detailed API documentation, see the [state API Reference](../api-reference/state.md).

### Using types

**Import Path:** `github.com/kolosys/discord/types`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/types"
)

func main() {
    // Example usage of types
    fmt.Println("Using types package")
}
```

#### Available Types
- **ActivityType** - ActivityType represents the type of a user activity.
- **AllowedMentionType** - AllowedMentionType represents types allowed in message mentions.
- **ApplicationCommandOptionType** - ApplicationCommandOptionType represents the type of a command option.
- **ApplicationCommandType** - ApplicationCommandType represents the type of an application command.
- **AuditLogEvent** - AuditLogEvent represents the type of audit log event.
- **ButtonStyle** - ButtonStyle represents the visual style of a button.
- **ChannelType** - ChannelType represents the type of a Discord channel.
- **Color** - Color represents a 24-bit RGB color value.
- **ComponentType** - ComponentType represents the type of a message component.
- **DefaultMessageNotificationLevel** - DefaultMessageNotificationLevel represents guild default notification settings.
- **ExplicitContentFilterLevel** - ExplicitContentFilterLevel represents a guild's media content filter.
- **GuildScheduledEventEntityType** - GuildScheduledEventEntityType represents the entity type of a scheduled event.
- **GuildScheduledEventPrivacyLevel** - GuildScheduledEventPrivacyLevel represents the privacy level of a scheduled event.
- **GuildScheduledEventStatus** - GuildScheduledEventStatus represents the status of a scheduled event.
- **InteractionCallbackType** - InteractionCallbackType represents the type of an interaction response.
- **InteractionType** - InteractionType represents the type of an interaction.
- **InviteTargetType** - InviteTargetType represents the target type for an invite.
- **MFALevel** - MFALevel represents a guild's MFA requirement for moderators.
- **MessageType** - MessageType represents the type of a Discord message.
- **NSFWLevel** - NSFWLevel represents a guild's NSFW classification.
- **NullableTimestamp** - NullableTimestamp is a timestamp that can be null in JSON.
- **OverwriteType** - OverwriteType represents the type of a permission overwrite.
- **Permissions** - Permissions represents a Discord permission bitfield.
- **PremiumTier** - PremiumTier represents a guild's boost level.
- **PremiumType** - PremiumType represents a user's Nitro subscription type.
- **Snowflake** - Snowflake is a Discord snowflake ID.
- **StatusType** - StatusType represents a user's online status.
- **StickerFormatType** - StickerFormatType represents the format of a sticker.
- **StickerType** - StickerType represents the type of a sticker.
- **TextInputStyle** - TextInputStyle represents the visual style of a text input.
- **Timestamp** - Timestamp is an ISO8601/RFC3339 timestamp.
- **VerificationLevel** - VerificationLevel represents a guild's verification requirements.
- **WebhookType** - WebhookType represents the type of a webhook.

For detailed API documentation, see the [types API Reference](../api-reference/types.md).

### Using voice

**Import Path:** `github.com/kolosys/discord/voice`



```go
package main

import (
    "fmt"
    "github.com/kolosys/discord/voice"
)

func main() {
    // Example usage of voice
    fmt.Println("Using voice package")
}
```

#### Available Types
- **OpusEncoder** - 
- **VoiceConnection** - 
- **VoicePayload** - 

For detailed API documentation, see the [voice API Reference](../api-reference/voice.md).

## Step-by-Step Tutorial

### Step 1: Import the Package

First, import the necessary packages in your Go file:

```go
import (
    "fmt"
    "github.com/kolosys/discord/commands"
    "github.com/kolosys/discord"
    "github.com/kolosys/discord/components"
    "github.com/kolosys/discord/embed"
    "github.com/kolosys/discord/events"
    "github.com/kolosys/discord/examples/basic/bot"
    "github.com/kolosys/discord/examples/basic/commands"
    "github.com/kolosys/discord/examples/basic/events"
    "github.com/kolosys/discord/examples/basic/routes"
    "github.com/kolosys/discord/examples/internal"
    "github.com/kolosys/discord/gateway"
    "github.com/kolosys/discord/models"
    "github.com/kolosys/discord/rest/internal"
    "github.com/kolosys/discord/rest"
    "github.com/kolosys/discord/server"
    "github.com/kolosys/discord/state"
    "github.com/kolosys/discord/types"
    "github.com/kolosys/discord/voice"
)
```

### Step 2: Initialize

Set up the basic configuration:

```go
func main() {
    // Initialize your application
    fmt.Println("Initializing discord...")
}
```

### Step 3: Use the Library

Implement your specific use case:

```go
func main() {
    // Your implementation here
}
```

## Running Your Code

To run your Go program:

```bash
go run main.go
```

To build an executable:

```bash
go build -o myapp
./myapp
```

## Configuration Options

discord can be configured to suit your needs. Check the [Core Concepts](../core-concepts/) section for detailed information about configuration options.

## Error Handling

Always handle errors appropriately:

```go
result, err := someFunction()
if err != nil {
    log.Fatalf("Error: %v", err)
}
```

## Best Practices

- Always handle errors returned by library functions
- Check the API documentation for detailed parameter information
- Use meaningful variable and function names
- Add comments to document your code

## Complete Example

Here's a complete working example:

```go
package main

import (
    "fmt"
    "log"
    "github.com/kolosys/discord/commands"
    "github.com/kolosys/discord"
    "github.com/kolosys/discord/components"
    "github.com/kolosys/discord/embed"
    "github.com/kolosys/discord/events"
    "github.com/kolosys/discord/examples/basic/bot"
    "github.com/kolosys/discord/examples/basic/commands"
    "github.com/kolosys/discord/examples/basic/events"
    "github.com/kolosys/discord/examples/basic/routes"
    "github.com/kolosys/discord/examples/internal"
    "github.com/kolosys/discord/gateway"
    "github.com/kolosys/discord/models"
    "github.com/kolosys/discord/rest/internal"
    "github.com/kolosys/discord/rest"
    "github.com/kolosys/discord/server"
    "github.com/kolosys/discord/state"
    "github.com/kolosys/discord/types"
    "github.com/kolosys/discord/voice"
)

func main() {
    fmt.Println("Starting discord application...")
    
    // Add your implementation here
    
    fmt.Println("Application completed successfully!")
}
```

## Next Steps

Now that you've seen the basics, explore:

- **[Core Concepts](../core-concepts/)** - Understanding the library architecture
- **[API Reference](../api-reference/)** - Complete API documentation
- **[Examples](../examples/README.md)** - More detailed examples
- **[Advanced Topics](../advanced/)** - Performance tuning and advanced patterns

## Getting Help

If you run into issues:

1. Check the [API Reference](../api-reference/)
2. Browse the [Examples](../examples/README.md)
3. Visit the [GitHub Issues](https://github.com/kolosys/discord/issues) page


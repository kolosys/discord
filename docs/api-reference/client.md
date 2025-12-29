# client API

Complete API documentation for the client package.

**Import Path:** `github.com/kolosys/discord/rest/internal`

## Package Documentation



## Constants

**HeaderContentType, HeaderAccept, HeaderAuthorization, HeaderUserAgent, HeaderXRequestID, HeaderXCorrelationID**

Standard header names


```go
const HeaderContentType = "Content-Type"
const HeaderAccept = "Accept"
const HeaderAuthorization = "Authorization"
const HeaderUserAgent = "User-Agent"
const HeaderXRequestID = "X-Request-ID"
const HeaderXCorrelationID = "X-Correlation-ID"
```

**ContentTypeJSON, ContentTypeXML, ContentTypeForm**

Content types


```go
const ContentTypeJSON = "application/json"
const ContentTypeXML = "application/xml"
const ContentTypeForm = "application/x-www-form-urlencoded"
```

**HeaderXRateLimitBucket, HeaderXRateLimitLimit, HeaderXRateLimitRemaining, HeaderXRateLimitReset, HeaderXRateLimitResetAfter**

API-specific headers


```go
const HeaderXRateLimitBucket = "X-RateLimit-Bucket"
const HeaderXRateLimitLimit = "X-RateLimit-Limit"
const HeaderXRateLimitRemaining = "X-RateLimit-Remaining"
const HeaderXRateLimitReset = "X-RateLimit-Reset"
const HeaderXRateLimitResetAfter = "X-RateLimit-Reset-After"
```

**SecurityTypeHTTP, SecurityTypeAPIKey, SecurityTypeOAuth2, SecurityTypeOpenID**

Security scheme types


```go
const SecurityTypeHTTP = "http"
const SecurityTypeAPIKey = "apiKey"
const SecurityTypeOAuth2 = "oauth2"
const SecurityTypeOpenID = "openIdConnect"
```

## Variables

**GetMyApplicationRoute, UpdateMyApplicationRoute, GetApplicationRoute, UpdateApplicationRoute, ApplicationsGetActivityInstanceRoute, UploadApplicationAttachmentRoute, ListApplicationCommandsRoute, CreateApplicationCommandRoute, BulkSetApplicationCommandsRoute, GetApplicationCommandRoute, UpdateApplicationCommandRoute, DeleteApplicationCommandRoute, ListApplicationEmojisRoute, CreateApplicationEmojiRoute, GetApplicationEmojiRoute, UpdateApplicationEmojiRoute, DeleteApplicationEmojiRoute, GetEntitlementsRoute, CreateEntitlementRoute, GetEntitlementRoute, DeleteEntitlementRoute, ConsumeEntitlementRoute, ListGuildApplicationCommandsRoute, CreateGuildApplicationCommandRoute, BulkSetGuildApplicationCommandsRoute, ListGuildApplicationCommandPermissionsRoute, GetGuildApplicationCommandRoute, UpdateGuildApplicationCommandRoute, DeleteGuildApplicationCommandRoute, GetGuildApplicationCommandPermissionsRoute, SetGuildApplicationCommandPermissionsRoute, GetApplicationRoleConnectionsMetadataRoute, UpdateApplicationRoleConnectionsMetadataRoute, GetChannelRoute, UpdateChannelRoute, DeleteChannelRoute, FollowChannelRoute, ListChannelInvitesRoute, CreateChannelInviteRoute, ListMessagesRoute, CreateMessageRoute, BulkDeleteMessagesRoute, ListPinsRoute, CreatePinRoute, DeletePinRoute, GetMessageRoute, UpdateMessageRoute, DeleteMessageRoute, CrosspostMessageRoute, DeleteAllMessageReactionsRoute, ListMessageReactionsByEmojiRoute, DeleteAllMessageReactionsByEmojiRoute, AddMyMessageReactionRoute, DeleteMyMessageReactionRoute, DeleteUserMessageReactionRoute, CreateThreadFromMessageRoute, SetChannelPermissionOverwriteRoute, DeleteChannelPermissionOverwriteRoute, DeprecatedListPinsRoute, DeprecatedCreatePinRoute, DeprecatedDeletePinRoute, GetAnswerVotersRoute, PollExpireRoute, AddGroupDmUserRoute, DeleteGroupDmUserRoute, SendSoundboardSoundRoute, ListThreadMembersRoute, JoinThreadRoute, LeaveThreadRoute, GetThreadMemberRoute, AddThreadMemberRoute, DeleteThreadMemberRoute, CreateThreadRoute, ListPrivateArchivedThreadsRoute, ListPublicArchivedThreadsRoute, ThreadSearchRoute, TriggerTypingIndicatorRoute, ListMyPrivateArchivedThreadsRoute, ListChannelWebhooksRoute, CreateWebhookRoute, GetGatewayRoute, GetBotGatewayRoute, GetGuildTemplateRoute, GetGuildRoute, UpdateGuildRoute, ListGuildAuditLogEntriesRoute, ListAutoModerationRulesRoute, CreateAutoModerationRuleRoute, GetAutoModerationRuleRoute, UpdateAutoModerationRuleRoute, DeleteAutoModerationRuleRoute, ListGuildBansRoute, GetGuildBanRoute, BanUserFromGuildRoute, UnbanUserFromGuildRoute, BulkBanUsersFromGuildRoute, ListGuildChannelsRoute, CreateGuildChannelRoute, BulkUpdateGuildChannelsRoute, ListGuildEmojisRoute, CreateGuildEmojiRoute, GetGuildEmojiRoute, UpdateGuildEmojiRoute, DeleteGuildEmojiRoute, ListGuildIntegrationsRoute, DeleteGuildIntegrationRoute, ListGuildInvitesRoute, ListGuildMembersRoute, UpdateMyGuildMemberRoute, SearchGuildMembersRoute, GetGuildMemberRoute, AddGuildMemberRoute, UpdateGuildMemberRoute, DeleteGuildMemberRoute, AddGuildMemberRoleRoute, DeleteGuildMemberRoleRoute, GetGuildNewMemberWelcomeRoute, GetGuildsOnboardingRoute, PutGuildsOnboardingRoute, GetGuildPreviewRoute, PreviewPruneGuildRoute, PruneGuildRoute, ListGuildVoiceRegionsRoute, ListGuildRolesRoute, CreateGuildRoleRoute, BulkUpdateGuildRolesRoute, GuildRoleMemberCountsRoute, GetGuildRoleRoute, UpdateGuildRoleRoute, DeleteGuildRoleRoute, ListGuildScheduledEventsRoute, CreateGuildScheduledEventRoute, GetGuildScheduledEventRoute, UpdateGuildScheduledEventRoute, DeleteGuildScheduledEventRoute, ListGuildScheduledEventUsersRoute, ListGuildSoundboardSoundsRoute, CreateGuildSoundboardSoundRoute, GetGuildSoundboardSoundRoute, UpdateGuildSoundboardSoundRoute, DeleteGuildSoundboardSoundRoute, ListGuildStickersRoute, CreateGuildStickerRoute, GetGuildStickerRoute, UpdateGuildStickerRoute, DeleteGuildStickerRoute, ListGuildTemplatesRoute, CreateGuildTemplateRoute, SyncGuildTemplateRoute, UpdateGuildTemplateRoute, DeleteGuildTemplateRoute, GetActiveGuildThreadsRoute, GetGuildVanityURLRoute, GetSelfVoiceStateRoute, UpdateSelfVoiceStateRoute, GetVoiceStateRoute, UpdateVoiceStateRoute, GetGuildWebhooksRoute, GetGuildWelcomeScreenRoute, UpdateGuildWelcomeScreenRoute, GetGuildWidgetSettingsRoute, UpdateGuildWidgetSettingsRoute, GetGuildWidgetRoute, GetGuildWidgetPngRoute, CreateInteractionResponseRoute, InviteResolveRoute, InviteRevokeRoute, CreateLobbyRoute, CreateOrJoinLobbyRoute, GetLobbyRoute, EditLobbyRoute, EditLobbyChannelLinkRoute, LeaveLobbyRoute, CreateLinkedLobbyGuildInviteForSelfRoute, BulkUpdateLobbyMembersRoute, AddLobbyMemberRoute, DeleteLobbyMemberRoute, CreateLinkedLobbyGuildInviteForUserRoute, GetLobbyMessagesRoute, CreateLobbyMessageRoute, GetMyOauth2AuthorizationRoute, GetMyOauth2ApplicationRoute, GetPublicKeysRoute, GetOpenidConnectUserinfoRoute, PartnerSdkUnmergeProvisionalAccountRoute, BotPartnerSdkUnmergeProvisionalAccountRoute, PartnerSdkTokenRoute, BotPartnerSdkTokenRoute, GetSoundboardDefaultSoundsRoute, CreateStageInstanceRoute, GetStageInstanceRoute, UpdateStageInstanceRoute, DeleteStageInstanceRoute, ListStickerPacksRoute, GetStickerPackRoute, GetStickerRoute, GetMyUserRoute, UpdateMyUserRoute, GetCurrentUserApplicationEntitlementsRoute, GetApplicationUserRoleConnectionRoute, UpdateApplicationUserRoleConnectionRoute, DeleteApplicationUserRoleConnectionRoute, CreateDmRoute, ListMyConnectionsRoute, ListMyGuildsRoute, LeaveGuildRoute, GetMyGuildMemberRoute, GetUserRoute, ListVoiceRegionsRoute, GetWebhookRoute, UpdateWebhookRoute, DeleteWebhookRoute, GetWebhookByTokenRoute, ExecuteWebhookRoute, UpdateWebhookByTokenRoute, DeleteWebhookByTokenRoute, ExecuteGithubCompatibleWebhookRoute, GetOriginalWebhookMessageRoute, UpdateOriginalWebhookMessageRoute, DeleteOriginalWebhookMessageRoute, GetWebhookMessageRoute, UpdateWebhookMessageRoute, DeleteWebhookMessageRoute, ExecuteSlackCompatibleWebhookRoute**

Route definitions


```go
var GetMyApplicationRoute = neuron.NewRoute[neuron.EmptyRequest, models.PrivateApplication](neuron.MethodGET, "/applications/@me")
var UpdateMyApplicationRoute = neuron.NewRoute[models.ApplicationFormPartial, models.PrivateApplication](neuron.MethodPATCH, "/applications/@me")
var GetApplicationRoute = neuron.NewRoute[neuron.EmptyRequest, models.PrivateApplication](neuron.MethodGET, "/applications/{application_id}")
var UpdateApplicationRoute = neuron.NewRoute[models.ApplicationFormPartial, models.PrivateApplication](neuron.MethodPATCH, "/applications/{application_id}")
var ApplicationsGetActivityInstanceRoute = neuron.NewRoute[neuron.EmptyRequest, models.EmbeddedActivityInstance](neuron.MethodGET, "/applications/{application_id}/activity-instances/{instance_id}")
var UploadApplicationAttachmentRoute = neuron.NewRoute[any, models.ActivitiesAttachment](neuron.MethodPOST, "/applications/{application_id}/attachment")
var ListApplicationCommandsRoute = neuron.NewRoute[neuron.EmptyRequest, []models.ApplicationCommand](neuron.MethodGET, "/applications/{application_id}/commands")
var CreateApplicationCommandRoute = neuron.NewRoute[models.ApplicationCommandCreateOptions, models.ApplicationCommand](neuron.MethodPOST, "/applications/{application_id}/commands")
var BulkSetApplicationCommandsRoute = neuron.NewRoute[any, []models.ApplicationCommand](neuron.MethodPUT, "/applications/{application_id}/commands")
var GetApplicationCommandRoute = neuron.NewRoute[neuron.EmptyRequest, models.ApplicationCommand](neuron.MethodGET, "/applications/{application_id}/commands/{command_id}")
var UpdateApplicationCommandRoute = neuron.NewRoute[models.ApplicationCommandPatchOptionsPartial, models.ApplicationCommand](neuron.MethodPATCH, "/applications/{application_id}/commands/{command_id}")
var DeleteApplicationCommandRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/applications/{application_id}/commands/{command_id}")
var ListApplicationEmojisRoute = neuron.NewRoute[neuron.EmptyRequest, models.ListApplicationEmojis](neuron.MethodGET, "/applications/{application_id}/emojis")
var CreateApplicationEmojiRoute = neuron.NewRoute[any, models.Emoji](neuron.MethodPOST, "/applications/{application_id}/emojis")
var GetApplicationEmojiRoute = neuron.NewRoute[neuron.EmptyRequest, models.Emoji](neuron.MethodGET, "/applications/{application_id}/emojis/{emoji_id}")
var UpdateApplicationEmojiRoute = neuron.NewRoute[any, models.Emoji](neuron.MethodPATCH, "/applications/{application_id}/emojis/{emoji_id}")
var DeleteApplicationEmojiRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/applications/{application_id}/emojis/{emoji_id}")
var GetEntitlementsRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/applications/{application_id}/entitlements")
var CreateEntitlementRoute = neuron.NewRoute[models.CreateEntitlementOptionsData, models.Entitlement](neuron.MethodPOST, "/applications/{application_id}/entitlements")
var GetEntitlementRoute = neuron.NewRoute[neuron.EmptyRequest, models.Entitlement](neuron.MethodGET, "/applications/{application_id}/entitlements/{entitlement_id}")
var DeleteEntitlementRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/applications/{application_id}/entitlements/{entitlement_id}")
var ConsumeEntitlementRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodPOST, "/applications/{application_id}/entitlements/{entitlement_id}/consume")
var ListGuildApplicationCommandsRoute = neuron.NewRoute[neuron.EmptyRequest, []models.ApplicationCommand](neuron.MethodGET, "/applications/{application_id}/guilds/{guild_id}/commands")
var CreateGuildApplicationCommandRoute = neuron.NewRoute[models.ApplicationCommandCreateOptions, models.ApplicationCommand](neuron.MethodPOST, "/applications/{application_id}/guilds/{guild_id}/commands")
var BulkSetGuildApplicationCommandsRoute = neuron.NewRoute[any, []models.ApplicationCommand](neuron.MethodPUT, "/applications/{application_id}/guilds/{guild_id}/commands")
var ListGuildApplicationCommandPermissionsRoute = neuron.NewRoute[neuron.EmptyRequest, []models.CommandPermissions](neuron.MethodGET, "/applications/{application_id}/guilds/{guild_id}/commands/permissions")
var GetGuildApplicationCommandRoute = neuron.NewRoute[neuron.EmptyRequest, models.ApplicationCommand](neuron.MethodGET, "/applications/{application_id}/guilds/{guild_id}/commands/{command_id}")
var UpdateGuildApplicationCommandRoute = neuron.NewRoute[models.ApplicationCommandPatchOptionsPartial, models.ApplicationCommand](neuron.MethodPATCH, "/applications/{application_id}/guilds/{guild_id}/commands/{command_id}")
var DeleteGuildApplicationCommandRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/applications/{application_id}/guilds/{guild_id}/commands/{command_id}")
var GetGuildApplicationCommandPermissionsRoute = neuron.NewRoute[neuron.EmptyRequest, models.CommandPermissions](neuron.MethodGET, "/applications/{application_id}/guilds/{guild_id}/commands/{command_id}/permissions")
var SetGuildApplicationCommandPermissionsRoute = neuron.NewRoute[any, models.CommandPermissions](neuron.MethodPUT, "/applications/{application_id}/guilds/{guild_id}/commands/{command_id}/permissions")
var GetApplicationRoleConnectionsMetadataRoute = neuron.NewRoute[neuron.EmptyRequest, []models.ApplicationRoleConnectionsMetadataItem](neuron.MethodGET, "/applications/{application_id}/role-connections/metadata")
var UpdateApplicationRoleConnectionsMetadataRoute = neuron.NewRoute[any, []models.ApplicationRoleConnectionsMetadataItem](neuron.MethodPUT, "/applications/{application_id}/role-connections/metadata")
var GetChannelRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/channels/{channel_id}")
var UpdateChannelRoute = neuron.NewRoute[any, any](neuron.MethodPATCH, "/channels/{channel_id}")
var DeleteChannelRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodDELETE, "/channels/{channel_id}")
var FollowChannelRoute = neuron.NewRoute[any, models.ChannelFollower](neuron.MethodPOST, "/channels/{channel_id}/followers")
var ListChannelInvitesRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/channels/{channel_id}/invites")
var CreateChannelInviteRoute = neuron.NewRoute[any, any](neuron.MethodPOST, "/channels/{channel_id}/invites")
var ListMessagesRoute = neuron.NewRoute[neuron.EmptyRequest, []models.Message](neuron.MethodGET, "/channels/{channel_id}/messages")
var CreateMessageRoute = neuron.NewRoute[models.MessageCreateOptions, models.Message](neuron.MethodPOST, "/channels/{channel_id}/messages")
var BulkDeleteMessagesRoute = neuron.NewRoute[any, neuron.EmptyResponse](neuron.MethodPOST, "/channels/{channel_id}/messages/bulk-delete")
var ListPinsRoute = neuron.NewRoute[neuron.EmptyRequest, models.PinnedMessages](neuron.MethodGET, "/channels/{channel_id}/messages/pins")
var CreatePinRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodPUT, "/channels/{channel_id}/messages/pins/{message_id}")
var DeletePinRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/channels/{channel_id}/messages/pins/{message_id}")
var GetMessageRoute = neuron.NewRoute[neuron.EmptyRequest, models.Message](neuron.MethodGET, "/channels/{channel_id}/messages/{message_id}")
var UpdateMessageRoute = neuron.NewRoute[models.MessageEditOptionsPartial, models.Message](neuron.MethodPATCH, "/channels/{channel_id}/messages/{message_id}")
var DeleteMessageRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/channels/{channel_id}/messages/{message_id}")
var CrosspostMessageRoute = neuron.NewRoute[neuron.EmptyRequest, models.Message](neuron.MethodPOST, "/channels/{channel_id}/messages/{message_id}/crosspost")
var DeleteAllMessageReactionsRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/channels/{channel_id}/messages/{message_id}/reactions")
var ListMessageReactionsByEmojiRoute = neuron.NewRoute[neuron.EmptyRequest, []models.User](neuron.MethodGET, "/channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}")
var DeleteAllMessageReactionsByEmojiRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}")
var AddMyMessageReactionRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodPUT, "/channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}/@me")
var DeleteMyMessageReactionRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}/@me")
var DeleteUserMessageReactionRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}/{user_id}")
var CreateThreadFromMessageRoute = neuron.NewRoute[models.CreateTextThreadWithMessageOptions, models.Thread](neuron.MethodPOST, "/channels/{channel_id}/messages/{message_id}/threads")
var SetChannelPermissionOverwriteRoute = neuron.NewRoute[any, neuron.EmptyResponse](neuron.MethodPUT, "/channels/{channel_id}/permissions/{overwrite_id}")
var DeleteChannelPermissionOverwriteRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/channels/{channel_id}/permissions/{overwrite_id}")
var DeprecatedListPinsRoute = neuron.NewRoute[neuron.EmptyRequest, []models.Message](neuron.MethodGET, "/channels/{channel_id}/pins")
var DeprecatedCreatePinRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodPUT, "/channels/{channel_id}/pins/{message_id}")
var DeprecatedDeletePinRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/channels/{channel_id}/pins/{message_id}")
var GetAnswerVotersRoute = neuron.NewRoute[neuron.EmptyRequest, models.PollAnswerDetails](neuron.MethodGET, "/channels/{channel_id}/polls/{message_id}/answers/{answer_id}")
var PollExpireRoute = neuron.NewRoute[neuron.EmptyRequest, models.Message](neuron.MethodPOST, "/channels/{channel_id}/polls/{message_id}/expire")
var AddGroupDmUserRoute = neuron.NewRoute[any, any](neuron.MethodPUT, "/channels/{channel_id}/recipients/{user_id}")
var DeleteGroupDmUserRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/channels/{channel_id}/recipients/{user_id}")
var SendSoundboardSoundRoute = neuron.NewRoute[models.SoundboardSoundSendOptions, neuron.EmptyResponse](neuron.MethodPOST, "/channels/{channel_id}/send-soundboard-sound")
var ListThreadMembersRoute = neuron.NewRoute[neuron.EmptyRequest, []models.ThreadMember](neuron.MethodGET, "/channels/{channel_id}/thread-members")
var JoinThreadRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodPUT, "/channels/{channel_id}/thread-members/@me")
var LeaveThreadRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/channels/{channel_id}/thread-members/@me")
var GetThreadMemberRoute = neuron.NewRoute[neuron.EmptyRequest, models.ThreadMember](neuron.MethodGET, "/channels/{channel_id}/thread-members/{user_id}")
var AddThreadMemberRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodPUT, "/channels/{channel_id}/thread-members/{user_id}")
var DeleteThreadMemberRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/channels/{channel_id}/thread-members/{user_id}")
var CreateThreadRoute = neuron.NewRoute[any, models.CreatedThread](neuron.MethodPOST, "/channels/{channel_id}/threads")
var ListPrivateArchivedThreadsRoute = neuron.NewRoute[neuron.EmptyRequest, models.Threads](neuron.MethodGET, "/channels/{channel_id}/threads/archived/private")
var ListPublicArchivedThreadsRoute = neuron.NewRoute[neuron.EmptyRequest, models.Threads](neuron.MethodGET, "/channels/{channel_id}/threads/archived/public")
var ThreadSearchRoute = neuron.NewRoute[neuron.EmptyRequest, models.ThreadSearch](neuron.MethodGET, "/channels/{channel_id}/threads/search")
var TriggerTypingIndicatorRoute = neuron.NewRoute[neuron.EmptyRequest, models.TypingIndicator](neuron.MethodPOST, "/channels/{channel_id}/typing")
var ListMyPrivateArchivedThreadsRoute = neuron.NewRoute[neuron.EmptyRequest, models.Threads](neuron.MethodGET, "/channels/{channel_id}/users/@me/threads/archived/private")
var ListChannelWebhooksRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/channels/{channel_id}/webhooks")
var CreateWebhookRoute = neuron.NewRoute[any, models.GuildIncomingWebhook](neuron.MethodPOST, "/channels/{channel_id}/webhooks")
var GetGatewayRoute = neuron.NewRoute[neuron.EmptyRequest, models.Gateway](neuron.MethodGET, "/gateway")
var GetBotGatewayRoute = neuron.NewRoute[neuron.EmptyRequest, models.GatewayBot](neuron.MethodGET, "/gateway/bot")
var GetGuildTemplateRoute = neuron.NewRoute[neuron.EmptyRequest, models.GuildTemplate](neuron.MethodGET, "/guilds/templates/{code}")
var GetGuildRoute = neuron.NewRoute[neuron.EmptyRequest, models.GuildWithCounts](neuron.MethodGET, "/guilds/{guild_id}")
var UpdateGuildRoute = neuron.NewRoute[models.GuildPatchOptionsPartial, models.Guild](neuron.MethodPATCH, "/guilds/{guild_id}")
var ListGuildAuditLogEntriesRoute = neuron.NewRoute[neuron.EmptyRequest, models.GuildAuditLog](neuron.MethodGET, "/guilds/{guild_id}/audit-logs")
var ListAutoModerationRulesRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/guilds/{guild_id}/auto-moderation/rules")
var CreateAutoModerationRuleRoute = neuron.NewRoute[any, any](neuron.MethodPOST, "/guilds/{guild_id}/auto-moderation/rules")
var GetAutoModerationRuleRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/guilds/{guild_id}/auto-moderation/rules/{rule_id}")
var UpdateAutoModerationRuleRoute = neuron.NewRoute[any, any](neuron.MethodPATCH, "/guilds/{guild_id}/auto-moderation/rules/{rule_id}")
var DeleteAutoModerationRuleRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/guilds/{guild_id}/auto-moderation/rules/{rule_id}")
var ListGuildBansRoute = neuron.NewRoute[neuron.EmptyRequest, []models.GuildBan](neuron.MethodGET, "/guilds/{guild_id}/bans")
var GetGuildBanRoute = neuron.NewRoute[neuron.EmptyRequest, models.GuildBan](neuron.MethodGET, "/guilds/{guild_id}/bans/{user_id}")
var BanUserFromGuildRoute = neuron.NewRoute[models.BanUserFromGuildOptions, neuron.EmptyResponse](neuron.MethodPUT, "/guilds/{guild_id}/bans/{user_id}")
var UnbanUserFromGuildRoute = neuron.NewRoute[models.UnbanUserFromGuildOptions, neuron.EmptyResponse](neuron.MethodDELETE, "/guilds/{guild_id}/bans/{user_id}")
var BulkBanUsersFromGuildRoute = neuron.NewRoute[models.BulkBanUsersOptions, models.BulkBanUsers](neuron.MethodPOST, "/guilds/{guild_id}/bulk-ban")
var ListGuildChannelsRoute = neuron.NewRoute[neuron.EmptyRequest, []models.GuildChannel](neuron.MethodGET, "/guilds/{guild_id}/channels")
var CreateGuildChannelRoute = neuron.NewRoute[models.CreateGuildChannelOptions, models.GuildChannel](neuron.MethodPOST, "/guilds/{guild_id}/channels")
var BulkUpdateGuildChannelsRoute = neuron.NewRoute[any, neuron.EmptyResponse](neuron.MethodPATCH, "/guilds/{guild_id}/channels")
var ListGuildEmojisRoute = neuron.NewRoute[neuron.EmptyRequest, []models.Emoji](neuron.MethodGET, "/guilds/{guild_id}/emojis")
var CreateGuildEmojiRoute = neuron.NewRoute[any, models.Emoji](neuron.MethodPOST, "/guilds/{guild_id}/emojis")
var GetGuildEmojiRoute = neuron.NewRoute[neuron.EmptyRequest, models.Emoji](neuron.MethodGET, "/guilds/{guild_id}/emojis/{emoji_id}")
var UpdateGuildEmojiRoute = neuron.NewRoute[any, models.Emoji](neuron.MethodPATCH, "/guilds/{guild_id}/emojis/{emoji_id}")
var DeleteGuildEmojiRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/guilds/{guild_id}/emojis/{emoji_id}")
var ListGuildIntegrationsRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/guilds/{guild_id}/integrations")
var DeleteGuildIntegrationRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/guilds/{guild_id}/integrations/{integration_id}")
var ListGuildInvitesRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/guilds/{guild_id}/invites")
var ListGuildMembersRoute = neuron.NewRoute[neuron.EmptyRequest, []models.GuildMember](neuron.MethodGET, "/guilds/{guild_id}/members")
var UpdateMyGuildMemberRoute = neuron.NewRoute[any, models.PrivateGuildMember](neuron.MethodPATCH, "/guilds/{guild_id}/members/@me")
var SearchGuildMembersRoute = neuron.NewRoute[neuron.EmptyRequest, []models.GuildMember](neuron.MethodGET, "/guilds/{guild_id}/members/search")
var GetGuildMemberRoute = neuron.NewRoute[neuron.EmptyRequest, models.GuildMember](neuron.MethodGET, "/guilds/{guild_id}/members/{user_id}")
var AddGuildMemberRoute = neuron.NewRoute[models.BotAddGuildMemberOptions, models.GuildMember](neuron.MethodPUT, "/guilds/{guild_id}/members/{user_id}")
var UpdateGuildMemberRoute = neuron.NewRoute[any, models.GuildMember](neuron.MethodPATCH, "/guilds/{guild_id}/members/{user_id}")
var DeleteGuildMemberRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/guilds/{guild_id}/members/{user_id}")
var AddGuildMemberRoleRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodPUT, "/guilds/{guild_id}/members/{user_id}/roles/{role_id}")
var DeleteGuildMemberRoleRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/guilds/{guild_id}/members/{user_id}/roles/{role_id}")
var GetGuildNewMemberWelcomeRoute = neuron.NewRoute[neuron.EmptyRequest, models.GuildHomeSettings](neuron.MethodGET, "/guilds/{guild_id}/new-member-welcome")
var GetGuildsOnboardingRoute = neuron.NewRoute[neuron.EmptyRequest, models.UserGuildOnboarding](neuron.MethodGET, "/guilds/{guild_id}/onboarding")
var PutGuildsOnboardingRoute = neuron.NewRoute[models.UpdateGuildOnboardingOptions, models.GuildOnboarding](neuron.MethodPUT, "/guilds/{guild_id}/onboarding")
var GetGuildPreviewRoute = neuron.NewRoute[neuron.EmptyRequest, models.GuildPreview](neuron.MethodGET, "/guilds/{guild_id}/preview")
var PreviewPruneGuildRoute = neuron.NewRoute[neuron.EmptyRequest, models.GuildPrune](neuron.MethodGET, "/guilds/{guild_id}/prune")
var PruneGuildRoute = neuron.NewRoute[models.PruneGuildOptions, models.GuildPrune](neuron.MethodPOST, "/guilds/{guild_id}/prune")
var ListGuildVoiceRegionsRoute = neuron.NewRoute[neuron.EmptyRequest, []models.VoiceRegion](neuron.MethodGET, "/guilds/{guild_id}/regions")
var ListGuildRolesRoute = neuron.NewRoute[neuron.EmptyRequest, []models.GuildRole](neuron.MethodGET, "/guilds/{guild_id}/roles")
var CreateGuildRoleRoute = neuron.NewRoute[models.CreateRoleOptions, models.GuildRole](neuron.MethodPOST, "/guilds/{guild_id}/roles")
var BulkUpdateGuildRolesRoute = neuron.NewRoute[any, []models.GuildRole](neuron.MethodPATCH, "/guilds/{guild_id}/roles")
var GuildRoleMemberCountsRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/guilds/{guild_id}/roles/member-counts")
var GetGuildRoleRoute = neuron.NewRoute[neuron.EmptyRequest, models.GuildRole](neuron.MethodGET, "/guilds/{guild_id}/roles/{role_id}")
var UpdateGuildRoleRoute = neuron.NewRoute[models.UpdateRoleOptionsPartial, models.GuildRole](neuron.MethodPATCH, "/guilds/{guild_id}/roles/{role_id}")
var DeleteGuildRoleRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/guilds/{guild_id}/roles/{role_id}")
var ListGuildScheduledEventsRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/guilds/{guild_id}/scheduled-events")
var CreateGuildScheduledEventRoute = neuron.NewRoute[any, any](neuron.MethodPOST, "/guilds/{guild_id}/scheduled-events")
var GetGuildScheduledEventRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}")
var UpdateGuildScheduledEventRoute = neuron.NewRoute[any, any](neuron.MethodPATCH, "/guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}")
var DeleteGuildScheduledEventRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}")
var ListGuildScheduledEventUsersRoute = neuron.NewRoute[neuron.EmptyRequest, []models.ScheduledEventUser](neuron.MethodGET, "/guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}/users")
var ListGuildSoundboardSoundsRoute = neuron.NewRoute[neuron.EmptyRequest, models.ListGuildSoundboardSounds](neuron.MethodGET, "/guilds/{guild_id}/soundboard-sounds")
var CreateGuildSoundboardSoundRoute = neuron.NewRoute[models.SoundboardCreateOptions, models.SoundboardSound](neuron.MethodPOST, "/guilds/{guild_id}/soundboard-sounds")
var GetGuildSoundboardSoundRoute = neuron.NewRoute[neuron.EmptyRequest, models.SoundboardSound](neuron.MethodGET, "/guilds/{guild_id}/soundboard-sounds/{sound_id}")
var UpdateGuildSoundboardSoundRoute = neuron.NewRoute[models.SoundboardPatchOptionsPartial, models.SoundboardSound](neuron.MethodPATCH, "/guilds/{guild_id}/soundboard-sounds/{sound_id}")
var DeleteGuildSoundboardSoundRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/guilds/{guild_id}/soundboard-sounds/{sound_id}")
var ListGuildStickersRoute = neuron.NewRoute[neuron.EmptyRequest, []models.GuildSticker](neuron.MethodGET, "/guilds/{guild_id}/stickers")
var CreateGuildStickerRoute = neuron.NewRoute[any, models.GuildSticker](neuron.MethodPOST, "/guilds/{guild_id}/stickers")
var GetGuildStickerRoute = neuron.NewRoute[neuron.EmptyRequest, models.GuildSticker](neuron.MethodGET, "/guilds/{guild_id}/stickers/{sticker_id}")
var UpdateGuildStickerRoute = neuron.NewRoute[any, models.GuildSticker](neuron.MethodPATCH, "/guilds/{guild_id}/stickers/{sticker_id}")
var DeleteGuildStickerRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/guilds/{guild_id}/stickers/{sticker_id}")
var ListGuildTemplatesRoute = neuron.NewRoute[neuron.EmptyRequest, []models.GuildTemplate](neuron.MethodGET, "/guilds/{guild_id}/templates")
var CreateGuildTemplateRoute = neuron.NewRoute[any, models.GuildTemplate](neuron.MethodPOST, "/guilds/{guild_id}/templates")
var SyncGuildTemplateRoute = neuron.NewRoute[neuron.EmptyRequest, models.GuildTemplate](neuron.MethodPUT, "/guilds/{guild_id}/templates/{code}")
var UpdateGuildTemplateRoute = neuron.NewRoute[any, models.GuildTemplate](neuron.MethodPATCH, "/guilds/{guild_id}/templates/{code}")
var DeleteGuildTemplateRoute = neuron.NewRoute[neuron.EmptyRequest, models.GuildTemplate](neuron.MethodDELETE, "/guilds/{guild_id}/templates/{code}")
var GetActiveGuildThreadsRoute = neuron.NewRoute[neuron.EmptyRequest, models.Threads](neuron.MethodGET, "/guilds/{guild_id}/threads/active")
var GetGuildVanityURLRoute = neuron.NewRoute[neuron.EmptyRequest, models.VanityURL](neuron.MethodGET, "/guilds/{guild_id}/vanity-url")
var GetSelfVoiceStateRoute = neuron.NewRoute[neuron.EmptyRequest, models.VoiceState](neuron.MethodGET, "/guilds/{guild_id}/voice-states/@me")
var UpdateSelfVoiceStateRoute = neuron.NewRoute[models.UpdateSelfVoiceStateOptionsPartial, neuron.EmptyResponse](neuron.MethodPATCH, "/guilds/{guild_id}/voice-states/@me")
var GetVoiceStateRoute = neuron.NewRoute[neuron.EmptyRequest, models.VoiceState](neuron.MethodGET, "/guilds/{guild_id}/voice-states/{user_id}")
var UpdateVoiceStateRoute = neuron.NewRoute[models.UpdateVoiceStateOptionsPartial, neuron.EmptyResponse](neuron.MethodPATCH, "/guilds/{guild_id}/voice-states/{user_id}")
var GetGuildWebhooksRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/guilds/{guild_id}/webhooks")
var GetGuildWelcomeScreenRoute = neuron.NewRoute[neuron.EmptyRequest, models.GuildWelcomeScreen](neuron.MethodGET, "/guilds/{guild_id}/welcome-screen")
var UpdateGuildWelcomeScreenRoute = neuron.NewRoute[models.WelcomeScreenPatchOptionsPartial, models.GuildWelcomeScreen](neuron.MethodPATCH, "/guilds/{guild_id}/welcome-screen")
var GetGuildWidgetSettingsRoute = neuron.NewRoute[neuron.EmptyRequest, models.WidgetSettings](neuron.MethodGET, "/guilds/{guild_id}/widget")
var UpdateGuildWidgetSettingsRoute = neuron.NewRoute[any, models.WidgetSettings](neuron.MethodPATCH, "/guilds/{guild_id}/widget")
var GetGuildWidgetRoute = neuron.NewRoute[neuron.EmptyRequest, models.Widget](neuron.MethodGET, "/guilds/{guild_id}/widget.json")
var GetGuildWidgetPngRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodGET, "/guilds/{guild_id}/widget.png")
var CreateInteractionResponseRoute = neuron.NewRoute[any, models.InteractionCallback](neuron.MethodPOST, "/interactions/{interaction_id}/{interaction_token}/callback")
var InviteResolveRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/invites/{code}")
var InviteRevokeRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodDELETE, "/invites/{code}")
var CreateLobbyRoute = neuron.NewRoute[any, models.Lobby](neuron.MethodPOST, "/lobbies")
var CreateOrJoinLobbyRoute = neuron.NewRoute[any, models.Lobby](neuron.MethodPUT, "/lobbies")
var GetLobbyRoute = neuron.NewRoute[neuron.EmptyRequest, models.Lobby](neuron.MethodGET, "/lobbies/{lobby_id}")
var EditLobbyRoute = neuron.NewRoute[any, models.Lobby](neuron.MethodPATCH, "/lobbies/{lobby_id}")
var EditLobbyChannelLinkRoute = neuron.NewRoute[any, models.Lobby](neuron.MethodPATCH, "/lobbies/{lobby_id}/channel-linking")
var LeaveLobbyRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/lobbies/{lobby_id}/members/@me")
var CreateLinkedLobbyGuildInviteForSelfRoute = neuron.NewRoute[neuron.EmptyRequest, models.LobbyGuildInvite](neuron.MethodPOST, "/lobbies/{lobby_id}/members/@me/invites")
var BulkUpdateLobbyMembersRoute = neuron.NewRoute[any, []models.LobbyMember](neuron.MethodPOST, "/lobbies/{lobby_id}/members/bulk")
var AddLobbyMemberRoute = neuron.NewRoute[any, models.LobbyMember](neuron.MethodPUT, "/lobbies/{lobby_id}/members/{user_id}")
var DeleteLobbyMemberRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/lobbies/{lobby_id}/members/{user_id}")
var CreateLinkedLobbyGuildInviteForUserRoute = neuron.NewRoute[neuron.EmptyRequest, models.LobbyGuildInvite](neuron.MethodPOST, "/lobbies/{lobby_id}/members/{user_id}/invites")
var GetLobbyMessagesRoute = neuron.NewRoute[neuron.EmptyRequest, []models.LobbyMessage](neuron.MethodGET, "/lobbies/{lobby_id}/messages")
var CreateLobbyMessageRoute = neuron.NewRoute[models.SDKMessageOptions, models.LobbyMessage](neuron.MethodPOST, "/lobbies/{lobby_id}/messages")
var GetMyOauth2AuthorizationRoute = neuron.NewRoute[neuron.EmptyRequest, models.OAuth2GetAuthorization](neuron.MethodGET, "/oauth2/@me")
var GetMyOauth2ApplicationRoute = neuron.NewRoute[neuron.EmptyRequest, models.PrivateApplication](neuron.MethodGET, "/oauth2/applications/@me")
var GetPublicKeysRoute = neuron.NewRoute[neuron.EmptyRequest, models.OAuth2GetKeys](neuron.MethodGET, "/oauth2/keys")
var GetOpenidConnectUserinfoRoute = neuron.NewRoute[neuron.EmptyRequest, models.OAuth2GetOpenIDConnectUserInfo](neuron.MethodGET, "/oauth2/userinfo")
var PartnerSdkUnmergeProvisionalAccountRoute = neuron.NewRoute[any, neuron.EmptyResponse](neuron.MethodPOST, "/partner-sdk/provisional-accounts/unmerge")
var BotPartnerSdkUnmergeProvisionalAccountRoute = neuron.NewRoute[any, neuron.EmptyResponse](neuron.MethodPOST, "/partner-sdk/provisional-accounts/unmerge/bot")
var PartnerSdkTokenRoute = neuron.NewRoute[any, models.ProvisionalToken](neuron.MethodPOST, "/partner-sdk/token")
var BotPartnerSdkTokenRoute = neuron.NewRoute[any, models.ProvisionalToken](neuron.MethodPOST, "/partner-sdk/token/bot")
var GetSoundboardDefaultSoundsRoute = neuron.NewRoute[neuron.EmptyRequest, []models.SoundboardSound](neuron.MethodGET, "/soundboard-default-sounds")
var CreateStageInstanceRoute = neuron.NewRoute[any, models.StageInstance](neuron.MethodPOST, "/stage-instances")
var GetStageInstanceRoute = neuron.NewRoute[neuron.EmptyRequest, models.StageInstance](neuron.MethodGET, "/stage-instances/{channel_id}")
var UpdateStageInstanceRoute = neuron.NewRoute[any, models.StageInstance](neuron.MethodPATCH, "/stage-instances/{channel_id}")
var DeleteStageInstanceRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/stage-instances/{channel_id}")
var ListStickerPacksRoute = neuron.NewRoute[neuron.EmptyRequest, models.StickerPackCollection](neuron.MethodGET, "/sticker-packs")
var GetStickerPackRoute = neuron.NewRoute[neuron.EmptyRequest, models.StickerPack](neuron.MethodGET, "/sticker-packs/{pack_id}")
var GetStickerRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/stickers/{sticker_id}")
var GetMyUserRoute = neuron.NewRoute[neuron.EmptyRequest, models.UserPII](neuron.MethodGET, "/users/@me")
var UpdateMyUserRoute = neuron.NewRoute[models.BotAccountPatchOptions, models.UserPII](neuron.MethodPATCH, "/users/@me")
var GetCurrentUserApplicationEntitlementsRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/users/@me/applications/{application_id}/entitlements")
var GetApplicationUserRoleConnectionRoute = neuron.NewRoute[neuron.EmptyRequest, models.ApplicationUserRoleConnection](neuron.MethodGET, "/users/@me/applications/{application_id}/role-connection")
var UpdateApplicationUserRoleConnectionRoute = neuron.NewRoute[models.UpdateApplicationUserRoleConnectionOptions, models.ApplicationUserRoleConnection](neuron.MethodPUT, "/users/@me/applications/{application_id}/role-connection")
var DeleteApplicationUserRoleConnectionRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/users/@me/applications/{application_id}/role-connection")
var CreateDmRoute = neuron.NewRoute[models.CreatePrivateChannelOptions, any](neuron.MethodPOST, "/users/@me/channels")
var ListMyConnectionsRoute = neuron.NewRoute[neuron.EmptyRequest, []models.ConnectedAccount](neuron.MethodGET, "/users/@me/connections")
var ListMyGuildsRoute = neuron.NewRoute[neuron.EmptyRequest, []models.MyGuild](neuron.MethodGET, "/users/@me/guilds")
var LeaveGuildRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/users/@me/guilds/{guild_id}")
var GetMyGuildMemberRoute = neuron.NewRoute[neuron.EmptyRequest, models.PrivateGuildMember](neuron.MethodGET, "/users/@me/guilds/{guild_id}/member")
var GetUserRoute = neuron.NewRoute[neuron.EmptyRequest, models.User](neuron.MethodGET, "/users/{user_id}")
var ListVoiceRegionsRoute = neuron.NewRoute[neuron.EmptyRequest, []models.VoiceRegion](neuron.MethodGET, "/voice/regions")
var GetWebhookRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/webhooks/{webhook_id}")
var UpdateWebhookRoute = neuron.NewRoute[any, any](neuron.MethodPATCH, "/webhooks/{webhook_id}")
var DeleteWebhookRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/webhooks/{webhook_id}")
var GetWebhookByTokenRoute = neuron.NewRoute[neuron.EmptyRequest, any](neuron.MethodGET, "/webhooks/{webhook_id}/{webhook_token}")
var ExecuteWebhookRoute = neuron.NewRoute[any, models.Message](neuron.MethodPOST, "/webhooks/{webhook_id}/{webhook_token}")
var UpdateWebhookByTokenRoute = neuron.NewRoute[any, any](neuron.MethodPATCH, "/webhooks/{webhook_id}/{webhook_token}")
var DeleteWebhookByTokenRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/webhooks/{webhook_id}/{webhook_token}")
var ExecuteGithubCompatibleWebhookRoute = neuron.NewRoute[models.GithubWebhook, neuron.EmptyResponse](neuron.MethodPOST, "/webhooks/{webhook_id}/{webhook_token}/github")
var GetOriginalWebhookMessageRoute = neuron.NewRoute[neuron.EmptyRequest, models.Message](neuron.MethodGET, "/webhooks/{webhook_id}/{webhook_token}/messages/@original")
var UpdateOriginalWebhookMessageRoute = neuron.NewRoute[models.IncomingWebhookUpdateOptionsPartial, models.Message](neuron.MethodPATCH, "/webhooks/{webhook_id}/{webhook_token}/messages/@original")
var DeleteOriginalWebhookMessageRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/webhooks/{webhook_id}/{webhook_token}/messages/@original")
var GetWebhookMessageRoute = neuron.NewRoute[neuron.EmptyRequest, models.Message](neuron.MethodGET, "/webhooks/{webhook_id}/{webhook_token}/messages/{message_id}")
var UpdateWebhookMessageRoute = neuron.NewRoute[models.IncomingWebhookUpdateOptionsPartial, models.Message](neuron.MethodPATCH, "/webhooks/{webhook_id}/{webhook_token}/messages/{message_id}")
var DeleteWebhookMessageRoute = neuron.NewRoute[neuron.EmptyRequest, neuron.EmptyResponse](neuron.MethodDELETE, "/webhooks/{webhook_id}/{webhook_token}/messages/{message_id}")
var ExecuteSlackCompatibleWebhookRoute = neuron.NewRoute[models.SlackWebhook, any](neuron.MethodPOST, "/webhooks/{webhook_id}/{webhook_token}/slack")
```

**BotTokenSecurity**

BotTokenSecurity represents the BotToken security scheme


```go
var BotTokenSecurity = SecurityScheme{Type: "apiKey", In: "header", Name: "Authorization"}
```

**OAuth2Security**

OAuth2Security represents the OAuth2 security scheme


```go
var OAuth2Security = SecurityScheme{Type: "oauth2"}
```

## Types

### ClientErrorResponseResponse
ClientErrorResponseResponse wraps the ClientErrorResponse response

#### Example Usage

```go
// Create a new ClientErrorResponseResponse
clienterrorresponseresponse := ClientErrorResponseResponse{
    Data: /* value */,
    XRateLimitLimit: "example",
    XRateLimitRemaining: "example",
    XRateLimitReset: "example",
    XRateLimitResetAfter: "example",
    XRateLimitBucket: "example",
}
```

#### Type Definition

```go
type ClientErrorResponseResponse struct {
    Data models.ErrorResponse
    XRateLimitLimit string `header:"X-RateLimit-Limit"`
    XRateLimitRemaining string `header:"X-RateLimit-Remaining"`
    XRateLimitReset string `header:"X-RateLimit-Reset"`
    XRateLimitResetAfter string `header:"X-RateLimit-Reset-After"`
    XRateLimitBucket string `header:"X-RateLimit-Bucket"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Data | `models.ErrorResponse` |  |
| XRateLimitLimit | `string` |  |
| XRateLimitRemaining | `string` |  |
| XRateLimitReset | `string` |  |
| XRateLimitResetAfter | `string` |  |
| XRateLimitBucket | `string` |  |

### ClientOptions
ClientOptions holds client configuration

#### Example Usage

```go
// Create a new ClientOptions
clientoptions := ClientOptions{
    BaseURL: "example",
    UserAgent: "example",
    Headers: /* value */,
    Timeout: /* value */,
    MaxRetries: 42,
}
```

#### Type Definition

```go
type ClientOptions struct {
    BaseURL string
    UserAgent string
    Headers http.Header
    Timeout time.Duration
    MaxRetries int
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| BaseURL | `string` |  |
| UserAgent | `string` |  |
| Headers | `http.Header` |  |
| Timeout | `time.Duration` |  |
| MaxRetries | `int` |  |

### ClientRatelimitedResponseResponse
ClientRatelimitedResponseResponse wraps the ClientRatelimitedResponse response

#### Example Usage

```go
// Create a new ClientRatelimitedResponseResponse
clientratelimitedresponseresponse := ClientRatelimitedResponseResponse{
    Data: /* value */,
    XRateLimitReset: "example",
    XRateLimitResetAfter: "example",
    XRateLimitBucket: "example",
    XRateLimitLimit: "example",
    XRateLimitRemaining: "example",
}
```

#### Type Definition

```go
type ClientRatelimitedResponseResponse struct {
    Data models.RatelimitedResponse
    XRateLimitReset string `header:"X-RateLimit-Reset"`
    XRateLimitResetAfter string `header:"X-RateLimit-Reset-After"`
    XRateLimitBucket string `header:"X-RateLimit-Bucket"`
    XRateLimitLimit string `header:"X-RateLimit-Limit"`
    XRateLimitRemaining string `header:"X-RateLimit-Remaining"`
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Data | `models.RatelimitedResponse` |  |
| XRateLimitReset | `string` |  |
| XRateLimitResetAfter | `string` |  |
| XRateLimitBucket | `string` |  |
| XRateLimitLimit | `string` |  |
| XRateLimitRemaining | `string` |  |

### EmptyResponse
Response type aliases for common responses

#### Example Usage

```go
// Create a new EmptyResponse
emptyresponse := EmptyResponse{

}
```

#### Type Definition

```go
type EmptyResponse struct {
}
```

### HeaderValues
HeaderValues holds typed header values

#### Example Usage

```go
// Create a new HeaderValues
headervalues := HeaderValues{
    XRateLimitBucket: "example",
    XRateLimitLimit: 42,
    XRateLimitRemaining: 42,
    XRateLimitReset: 3.14,
    XRateLimitResetAfter: 3.14,
}
```

#### Type Definition

```go
type HeaderValues struct {
    XRateLimitBucket string
    XRateLimitLimit int64
    XRateLimitRemaining int64
    XRateLimitReset float64
    XRateLimitResetAfter float64
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| XRateLimitBucket | `string` | The bucket that the request belongs to |
| XRateLimitLimit | `int64` | The maximum number of requests that can be made in the current ratelimit window |
| XRateLimitRemaining | `int64` | The number of requests remaining in the current ratelimit window |
| XRateLimitReset | `float64` | A unix timestamp in seconds at which the current ratelimit window resets |
| XRateLimitResetAfter | `float64` | The duration in seconds until the current ratelimit window resets |

### RESTClient
RESTClient wraps a neuron.Client with generated route methods

#### Example Usage

```go
// Create a new RESTClient
restclient := RESTClient{

}
```

#### Type Definition

```go
type RESTClient struct {
    *neuron.Client
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| **neuron.Client | `*neuron.Client` |  |

### Constructor Functions

### NewRESTClient

NewRESTClient creates a new REST client

```go
func NewRESTClient(opts ClientOptions) *RESTClient
```

**Parameters:**
- `opts` (ClientOptions)

**Returns:**
- *RESTClient

## Methods

### AddGroupDmUser

AddGroupDmUser executes PUT /channels/{channel_id}/recipients/{user_id}

```go
func (*RESTClient) AddGroupDmUser(ctx context.Context, channelid models.SnowflakeType, userid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### AddGuildMember

AddGuildMember executes PUT /guilds/{guild_id}/members/{user_id}

```go
func (*RESTClient) AddGuildMember(ctx context.Context, guildid models.SnowflakeType, userid models.SnowflakeType, body models.BotAddGuildMemberOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `body` (models.BotAddGuildMemberOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### AddGuildMemberRole

AddGuildMemberRole executes PUT /guilds/{guild_id}/members/{user_id}/roles/{role_id}

```go
func (*RESTClient) AddGuildMemberRole(ctx context.Context, guildid models.SnowflakeType, userid models.SnowflakeType, roleid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `roleid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### AddLobbyMember

AddLobbyMember executes PUT /lobbies/{lobby_id}/members/{user_id}

```go
func (*RESTClient) AddLobbyMember(ctx context.Context, lobbyid models.SnowflakeType, userid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `lobbyid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### AddMyMessageReaction

AddMyMessageReaction executes PUT /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}/@me

```go
func (*RESTClient) AddMyMessageReaction(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, emojiname string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `emojiname` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### AddThreadMember

AddThreadMember executes PUT /channels/{channel_id}/thread-members/{user_id}

```go
func (*RESTClient) AddThreadMember(ctx context.Context, channelid models.SnowflakeType, userid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ApplicationsGetActivityInstance

ApplicationsGetActivityInstance executes GET /applications/{application_id}/activity-instances/{instance_id}

```go
func (*RESTClient) ApplicationsGetActivityInstance(ctx context.Context, applicationid models.SnowflakeType, instanceid string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `instanceid` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### BanUserFromGuild

BanUserFromGuild executes PUT /guilds/{guild_id}/bans/{user_id}

```go
func (*RESTClient) BanUserFromGuild(ctx context.Context, guildid models.SnowflakeType, userid models.SnowflakeType, body models.BanUserFromGuildOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `body` (models.BanUserFromGuildOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### BaseURL

BaseURL returns the client's base URL

```go
func (*RESTClient) BaseURL() string
```

**Parameters:**
  None

**Returns:**
- string

### BotPartnerSdkToken

BotPartnerSdkToken executes POST /partner-sdk/token/bot

```go
func (*RESTClient) BotPartnerSdkToken(ctx context.Context, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### BotPartnerSdkUnmergeProvisionalAccount

BotPartnerSdkUnmergeProvisionalAccount executes POST /partner-sdk/provisional-accounts/unmerge/bot

```go
func (*RESTClient) BotPartnerSdkUnmergeProvisionalAccount(ctx context.Context, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### BulkBanUsersFromGuild

BulkBanUsersFromGuild executes POST /guilds/{guild_id}/bulk-ban

```go
func (*RESTClient) BulkBanUsersFromGuild(ctx context.Context, guildid models.SnowflakeType, body models.BulkBanUsersOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (models.BulkBanUsersOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### BulkDeleteMessages

BulkDeleteMessages executes POST /channels/{channel_id}/messages/bulk-delete

```go
func (*RESTClient) BulkDeleteMessages(ctx context.Context, channelid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### BulkSetApplicationCommands

BulkSetApplicationCommands executes PUT /applications/{application_id}/commands

```go
func (*RESTClient) BulkSetApplicationCommands(ctx context.Context, applicationid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### BulkSetGuildApplicationCommands

BulkSetGuildApplicationCommands executes PUT /applications/{application_id}/guilds/{guild_id}/commands

```go
func (*RESTClient) BulkSetGuildApplicationCommands(ctx context.Context, applicationid models.SnowflakeType, guildid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `guildid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### BulkUpdateGuildChannels

BulkUpdateGuildChannels executes PATCH /guilds/{guild_id}/channels

```go
func (*RESTClient) BulkUpdateGuildChannels(ctx context.Context, guildid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### BulkUpdateGuildRoles

BulkUpdateGuildRoles executes PATCH /guilds/{guild_id}/roles

```go
func (*RESTClient) BulkUpdateGuildRoles(ctx context.Context, guildid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### BulkUpdateLobbyMembers

BulkUpdateLobbyMembers executes POST /lobbies/{lobby_id}/members/bulk

```go
func (*RESTClient) BulkUpdateLobbyMembers(ctx context.Context, lobbyid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `lobbyid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ConsumeEntitlement

ConsumeEntitlement executes POST /applications/{application_id}/entitlements/{entitlement_id}/consume

```go
func (*RESTClient) ConsumeEntitlement(ctx context.Context, applicationid models.SnowflakeType, entitlementid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `entitlementid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateApplicationCommand

CreateApplicationCommand executes POST /applications/{application_id}/commands

```go
func (*RESTClient) CreateApplicationCommand(ctx context.Context, applicationid models.SnowflakeType, body models.ApplicationCommandCreateOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `body` (models.ApplicationCommandCreateOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateApplicationEmoji

CreateApplicationEmoji executes POST /applications/{application_id}/emojis

```go
func (*RESTClient) CreateApplicationEmoji(ctx context.Context, applicationid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateAutoModerationRule

CreateAutoModerationRule executes POST /guilds/{guild_id}/auto-moderation/rules

```go
func (*RESTClient) CreateAutoModerationRule(ctx context.Context, guildid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateChannelInvite

CreateChannelInvite executes POST /channels/{channel_id}/invites

```go
func (*RESTClient) CreateChannelInvite(ctx context.Context, channelid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateDm

CreateDm executes POST /users/@me/channels

```go
func (*RESTClient) CreateDm(ctx context.Context, body models.CreatePrivateChannelOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `body` (models.CreatePrivateChannelOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateEntitlement

CreateEntitlement executes POST /applications/{application_id}/entitlements

```go
func (*RESTClient) CreateEntitlement(ctx context.Context, applicationid models.SnowflakeType, body models.CreateEntitlementOptionsData, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `body` (models.CreateEntitlementOptionsData)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateGuildApplicationCommand

CreateGuildApplicationCommand executes POST /applications/{application_id}/guilds/{guild_id}/commands

```go
func (*RESTClient) CreateGuildApplicationCommand(ctx context.Context, applicationid models.SnowflakeType, guildid models.SnowflakeType, body models.ApplicationCommandCreateOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `guildid` (models.SnowflakeType)
- `body` (models.ApplicationCommandCreateOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateGuildChannel

CreateGuildChannel executes POST /guilds/{guild_id}/channels

```go
func (*RESTClient) CreateGuildChannel(ctx context.Context, guildid models.SnowflakeType, body models.CreateGuildChannelOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (models.CreateGuildChannelOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateGuildEmoji

CreateGuildEmoji executes POST /guilds/{guild_id}/emojis

```go
func (*RESTClient) CreateGuildEmoji(ctx context.Context, guildid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateGuildRole

CreateGuildRole executes POST /guilds/{guild_id}/roles

```go
func (*RESTClient) CreateGuildRole(ctx context.Context, guildid models.SnowflakeType, body models.CreateRoleOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (models.CreateRoleOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateGuildScheduledEvent

CreateGuildScheduledEvent executes POST /guilds/{guild_id}/scheduled-events

```go
func (*RESTClient) CreateGuildScheduledEvent(ctx context.Context, guildid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateGuildSoundboardSound

CreateGuildSoundboardSound executes POST /guilds/{guild_id}/soundboard-sounds

```go
func (*RESTClient) CreateGuildSoundboardSound(ctx context.Context, guildid models.SnowflakeType, body models.SoundboardCreateOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (models.SoundboardCreateOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateGuildSticker

CreateGuildSticker executes POST /guilds/{guild_id}/stickers

```go
func (*RESTClient) CreateGuildSticker(ctx context.Context, guildid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateGuildTemplate

CreateGuildTemplate executes POST /guilds/{guild_id}/templates

```go
func (*RESTClient) CreateGuildTemplate(ctx context.Context, guildid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateInteractionResponse

CreateInteractionResponse executes POST /interactions/{interaction_id}/{interaction_token}/callback

```go
func (*RESTClient) CreateInteractionResponse(ctx context.Context, interactionid models.SnowflakeType, interactiontoken string, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `interactionid` (models.SnowflakeType)
- `interactiontoken` (string)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateLinkedLobbyGuildInviteForSelf

CreateLinkedLobbyGuildInviteForSelf executes POST /lobbies/{lobby_id}/members/@me/invites

```go
func (*RESTClient) CreateLinkedLobbyGuildInviteForSelf(ctx context.Context, lobbyid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `lobbyid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateLinkedLobbyGuildInviteForUser

CreateLinkedLobbyGuildInviteForUser executes POST /lobbies/{lobby_id}/members/{user_id}/invites

```go
func (*RESTClient) CreateLinkedLobbyGuildInviteForUser(ctx context.Context, lobbyid models.SnowflakeType, userid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `lobbyid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateLobby

CreateLobby executes POST /lobbies

```go
func (*RESTClient) CreateLobby(ctx context.Context, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateLobbyMessage

CreateLobbyMessage executes POST /lobbies/{lobby_id}/messages

```go
func (*RESTClient) CreateLobbyMessage(ctx context.Context, lobbyid models.SnowflakeType, body models.SDKMessageOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `lobbyid` (models.SnowflakeType)
- `body` (models.SDKMessageOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateMessage

CreateMessage executes POST /channels/{channel_id}/messages

```go
func (*RESTClient) CreateMessage(ctx context.Context, channelid models.SnowflakeType, body models.MessageCreateOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `body` (models.MessageCreateOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateOrJoinLobby

CreateOrJoinLobby executes PUT /lobbies

```go
func (*RESTClient) CreateOrJoinLobby(ctx context.Context, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreatePin

CreatePin executes PUT /channels/{channel_id}/messages/pins/{message_id}

```go
func (*RESTClient) CreatePin(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateStageInstance

CreateStageInstance executes POST /stage-instances

```go
func (*RESTClient) CreateStageInstance(ctx context.Context, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateThread

CreateThread executes POST /channels/{channel_id}/threads

```go
func (*RESTClient) CreateThread(ctx context.Context, channelid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateThreadFromMessage

CreateThreadFromMessage executes POST /channels/{channel_id}/messages/{message_id}/threads

```go
func (*RESTClient) CreateThreadFromMessage(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, body models.CreateTextThreadWithMessageOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `body` (models.CreateTextThreadWithMessageOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CreateWebhook

CreateWebhook executes POST /channels/{channel_id}/webhooks

```go
func (*RESTClient) CreateWebhook(ctx context.Context, channelid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### CrosspostMessage

CrosspostMessage executes POST /channels/{channel_id}/messages/{message_id}/crosspost

```go
func (*RESTClient) CrosspostMessage(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteAllMessageReactions

DeleteAllMessageReactions executes DELETE /channels/{channel_id}/messages/{message_id}/reactions

```go
func (*RESTClient) DeleteAllMessageReactions(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteAllMessageReactionsByEmoji

DeleteAllMessageReactionsByEmoji executes DELETE /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}

```go
func (*RESTClient) DeleteAllMessageReactionsByEmoji(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, emojiname string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `emojiname` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteApplicationCommand

DeleteApplicationCommand executes DELETE /applications/{application_id}/commands/{command_id}

```go
func (*RESTClient) DeleteApplicationCommand(ctx context.Context, applicationid models.SnowflakeType, commandid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `commandid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteApplicationEmoji

DeleteApplicationEmoji executes DELETE /applications/{application_id}/emojis/{emoji_id}

```go
func (*RESTClient) DeleteApplicationEmoji(ctx context.Context, applicationid models.SnowflakeType, emojiid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `emojiid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteApplicationUserRoleConnection

DeleteApplicationUserRoleConnection executes DELETE /users/@me/applications/{application_id}/role-connection

```go
func (*RESTClient) DeleteApplicationUserRoleConnection(ctx context.Context, applicationid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteAutoModerationRule

DeleteAutoModerationRule executes DELETE /guilds/{guild_id}/auto-moderation/rules/{rule_id}

```go
func (*RESTClient) DeleteAutoModerationRule(ctx context.Context, guildid models.SnowflakeType, ruleid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `ruleid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteChannel

DeleteChannel executes DELETE /channels/{channel_id}

```go
func (*RESTClient) DeleteChannel(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteChannelPermissionOverwrite

DeleteChannelPermissionOverwrite executes DELETE /channels/{channel_id}/permissions/{overwrite_id}

```go
func (*RESTClient) DeleteChannelPermissionOverwrite(ctx context.Context, channelid models.SnowflakeType, overwriteid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `overwriteid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteEntitlement

DeleteEntitlement executes DELETE /applications/{application_id}/entitlements/{entitlement_id}

```go
func (*RESTClient) DeleteEntitlement(ctx context.Context, applicationid models.SnowflakeType, entitlementid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `entitlementid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteGroupDmUser

DeleteGroupDmUser executes DELETE /channels/{channel_id}/recipients/{user_id}

```go
func (*RESTClient) DeleteGroupDmUser(ctx context.Context, channelid models.SnowflakeType, userid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteGuildApplicationCommand

DeleteGuildApplicationCommand executes DELETE /applications/{application_id}/guilds/{guild_id}/commands/{command_id}

```go
func (*RESTClient) DeleteGuildApplicationCommand(ctx context.Context, applicationid models.SnowflakeType, guildid models.SnowflakeType, commandid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `guildid` (models.SnowflakeType)
- `commandid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteGuildEmoji

DeleteGuildEmoji executes DELETE /guilds/{guild_id}/emojis/{emoji_id}

```go
func (*RESTClient) DeleteGuildEmoji(ctx context.Context, guildid models.SnowflakeType, emojiid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `emojiid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteGuildIntegration

DeleteGuildIntegration executes DELETE /guilds/{guild_id}/integrations/{integration_id}

```go
func (*RESTClient) DeleteGuildIntegration(ctx context.Context, guildid models.SnowflakeType, integrationid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `integrationid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteGuildMember

DeleteGuildMember executes DELETE /guilds/{guild_id}/members/{user_id}

```go
func (*RESTClient) DeleteGuildMember(ctx context.Context, guildid models.SnowflakeType, userid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteGuildMemberRole

DeleteGuildMemberRole executes DELETE /guilds/{guild_id}/members/{user_id}/roles/{role_id}

```go
func (*RESTClient) DeleteGuildMemberRole(ctx context.Context, guildid models.SnowflakeType, userid models.SnowflakeType, roleid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `roleid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteGuildRole

DeleteGuildRole executes DELETE /guilds/{guild_id}/roles/{role_id}

```go
func (*RESTClient) DeleteGuildRole(ctx context.Context, guildid models.SnowflakeType, roleid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `roleid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteGuildScheduledEvent

DeleteGuildScheduledEvent executes DELETE /guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}

```go
func (*RESTClient) DeleteGuildScheduledEvent(ctx context.Context, guildid models.SnowflakeType, guildscheduledeventid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `guildscheduledeventid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteGuildSoundboardSound

DeleteGuildSoundboardSound executes DELETE /guilds/{guild_id}/soundboard-sounds/{sound_id}

```go
func (*RESTClient) DeleteGuildSoundboardSound(ctx context.Context, guildid models.SnowflakeType, soundid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `soundid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteGuildSticker

DeleteGuildSticker executes DELETE /guilds/{guild_id}/stickers/{sticker_id}

```go
func (*RESTClient) DeleteGuildSticker(ctx context.Context, guildid models.SnowflakeType, stickerid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `stickerid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteGuildTemplate

DeleteGuildTemplate executes DELETE /guilds/{guild_id}/templates/{code}

```go
func (*RESTClient) DeleteGuildTemplate(ctx context.Context, guildid models.SnowflakeType, code string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `code` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteLobbyMember

DeleteLobbyMember executes DELETE /lobbies/{lobby_id}/members/{user_id}

```go
func (*RESTClient) DeleteLobbyMember(ctx context.Context, lobbyid models.SnowflakeType, userid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `lobbyid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteMessage

DeleteMessage executes DELETE /channels/{channel_id}/messages/{message_id}

```go
func (*RESTClient) DeleteMessage(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteMyMessageReaction

DeleteMyMessageReaction executes DELETE /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}/@me

```go
func (*RESTClient) DeleteMyMessageReaction(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, emojiname string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `emojiname` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteOriginalWebhookMessage

DeleteOriginalWebhookMessage executes DELETE /webhooks/{webhook_id}/{webhook_token}/messages/@original

```go
func (*RESTClient) DeleteOriginalWebhookMessage(ctx context.Context, webhookid models.SnowflakeType, webhooktoken string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `webhooktoken` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeletePin

DeletePin executes DELETE /channels/{channel_id}/messages/pins/{message_id}

```go
func (*RESTClient) DeletePin(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteStageInstance

DeleteStageInstance executes DELETE /stage-instances/{channel_id}

```go
func (*RESTClient) DeleteStageInstance(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteThreadMember

DeleteThreadMember executes DELETE /channels/{channel_id}/thread-members/{user_id}

```go
func (*RESTClient) DeleteThreadMember(ctx context.Context, channelid models.SnowflakeType, userid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteUserMessageReaction

DeleteUserMessageReaction executes DELETE /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}/{user_id}

```go
func (*RESTClient) DeleteUserMessageReaction(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, emojiname string, userid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `emojiname` (string)
- `userid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteWebhook

DeleteWebhook executes DELETE /webhooks/{webhook_id}

```go
func (*RESTClient) DeleteWebhook(ctx context.Context, webhookid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteWebhookByToken

DeleteWebhookByToken executes DELETE /webhooks/{webhook_id}/{webhook_token}

```go
func (*RESTClient) DeleteWebhookByToken(ctx context.Context, webhookid models.SnowflakeType, webhooktoken string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `webhooktoken` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeleteWebhookMessage

DeleteWebhookMessage executes DELETE /webhooks/{webhook_id}/{webhook_token}/messages/{message_id}

```go
func (*RESTClient) DeleteWebhookMessage(ctx context.Context, webhookid models.SnowflakeType, webhooktoken string, messageid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `webhooktoken` (string)
- `messageid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeprecatedCreatePin

DeprecatedCreatePin executes PUT /channels/{channel_id}/pins/{message_id}

```go
func (*RESTClient) DeprecatedCreatePin(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeprecatedDeletePin

DeprecatedDeletePin executes DELETE /channels/{channel_id}/pins/{message_id}

```go
func (*RESTClient) DeprecatedDeletePin(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### DeprecatedListPins

DeprecatedListPins executes GET /channels/{channel_id}/pins

```go
func (*RESTClient) DeprecatedListPins(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### EditLobby

EditLobby executes PATCH /lobbies/{lobby_id}

```go
func (*RESTClient) EditLobby(ctx context.Context, lobbyid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `lobbyid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### EditLobbyChannelLink

EditLobbyChannelLink executes PATCH /lobbies/{lobby_id}/channel-linking

```go
func (*RESTClient) EditLobbyChannelLink(ctx context.Context, lobbyid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `lobbyid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ExecuteGithubCompatibleWebhook

ExecuteGithubCompatibleWebhook executes POST /webhooks/{webhook_id}/{webhook_token}/github

```go
func (*RESTClient) ExecuteGithubCompatibleWebhook(ctx context.Context, webhookid models.SnowflakeType, webhooktoken string, body models.GithubWebhook, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `webhooktoken` (string)
- `body` (models.GithubWebhook)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ExecuteSlackCompatibleWebhook

ExecuteSlackCompatibleWebhook executes POST /webhooks/{webhook_id}/{webhook_token}/slack

```go
func (*RESTClient) ExecuteSlackCompatibleWebhook(ctx context.Context, webhookid models.SnowflakeType, webhooktoken string, body models.SlackWebhook, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `webhooktoken` (string)
- `body` (models.SlackWebhook)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ExecuteWebhook

ExecuteWebhook executes POST /webhooks/{webhook_id}/{webhook_token}

```go
func (*RESTClient) ExecuteWebhook(ctx context.Context, webhookid models.SnowflakeType, webhooktoken string, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `webhooktoken` (string)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### FollowChannel

FollowChannel executes POST /channels/{channel_id}/followers

```go
func (*RESTClient) FollowChannel(ctx context.Context, channelid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetActiveGuildThreads

GetActiveGuildThreads executes GET /guilds/{guild_id}/threads/active

```go
func (*RESTClient) GetActiveGuildThreads(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetAnswerVoters

GetAnswerVoters executes GET /channels/{channel_id}/polls/{message_id}/answers/{answer_id}

```go
func (*RESTClient) GetAnswerVoters(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, answerid int32, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `answerid` (int32)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetApplication

GetApplication executes GET /applications/{application_id}

```go
func (*RESTClient) GetApplication(ctx context.Context, applicationid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetApplicationCommand

GetApplicationCommand executes GET /applications/{application_id}/commands/{command_id}

```go
func (*RESTClient) GetApplicationCommand(ctx context.Context, applicationid models.SnowflakeType, commandid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `commandid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetApplicationEmoji

GetApplicationEmoji executes GET /applications/{application_id}/emojis/{emoji_id}

```go
func (*RESTClient) GetApplicationEmoji(ctx context.Context, applicationid models.SnowflakeType, emojiid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `emojiid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetApplicationRoleConnectionsMetadata

GetApplicationRoleConnectionsMetadata executes GET /applications/{application_id}/role-connections/metadata

```go
func (*RESTClient) GetApplicationRoleConnectionsMetadata(ctx context.Context, applicationid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetApplicationUserRoleConnection

GetApplicationUserRoleConnection executes GET /users/@me/applications/{application_id}/role-connection

```go
func (*RESTClient) GetApplicationUserRoleConnection(ctx context.Context, applicationid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetAutoModerationRule

GetAutoModerationRule executes GET /guilds/{guild_id}/auto-moderation/rules/{rule_id}

```go
func (*RESTClient) GetAutoModerationRule(ctx context.Context, guildid models.SnowflakeType, ruleid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `ruleid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetBotGateway

GetBotGateway executes GET /gateway/bot

```go
func (*RESTClient) GetBotGateway(ctx context.Context, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetChannel

GetChannel executes GET /channels/{channel_id}

```go
func (*RESTClient) GetChannel(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetCurrentUserApplicationEntitlements

GetCurrentUserApplicationEntitlements executes GET /users/@me/applications/{application_id}/entitlements

```go
func (*RESTClient) GetCurrentUserApplicationEntitlements(ctx context.Context, applicationid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetEntitlement

GetEntitlement executes GET /applications/{application_id}/entitlements/{entitlement_id}

```go
func (*RESTClient) GetEntitlement(ctx context.Context, applicationid models.SnowflakeType, entitlementid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `entitlementid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetEntitlements

GetEntitlements executes GET /applications/{application_id}/entitlements

```go
func (*RESTClient) GetEntitlements(ctx context.Context, applicationid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGateway

GetGateway executes GET /gateway

```go
func (*RESTClient) GetGateway(ctx context.Context, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuild

GetGuild executes GET /guilds/{guild_id}

```go
func (*RESTClient) GetGuild(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildApplicationCommand

GetGuildApplicationCommand executes GET /applications/{application_id}/guilds/{guild_id}/commands/{command_id}

```go
func (*RESTClient) GetGuildApplicationCommand(ctx context.Context, applicationid models.SnowflakeType, guildid models.SnowflakeType, commandid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `guildid` (models.SnowflakeType)
- `commandid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildApplicationCommandPermissions

GetGuildApplicationCommandPermissions executes GET /applications/{application_id}/guilds/{guild_id}/commands/{command_id}/permissions

```go
func (*RESTClient) GetGuildApplicationCommandPermissions(ctx context.Context, applicationid models.SnowflakeType, guildid models.SnowflakeType, commandid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `guildid` (models.SnowflakeType)
- `commandid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildBan

GetGuildBan executes GET /guilds/{guild_id}/bans/{user_id}

```go
func (*RESTClient) GetGuildBan(ctx context.Context, guildid models.SnowflakeType, userid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildEmoji

GetGuildEmoji executes GET /guilds/{guild_id}/emojis/{emoji_id}

```go
func (*RESTClient) GetGuildEmoji(ctx context.Context, guildid models.SnowflakeType, emojiid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `emojiid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildMember

GetGuildMember executes GET /guilds/{guild_id}/members/{user_id}

```go
func (*RESTClient) GetGuildMember(ctx context.Context, guildid models.SnowflakeType, userid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildNewMemberWelcome

GetGuildNewMemberWelcome executes GET /guilds/{guild_id}/new-member-welcome

```go
func (*RESTClient) GetGuildNewMemberWelcome(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildPreview

GetGuildPreview executes GET /guilds/{guild_id}/preview

```go
func (*RESTClient) GetGuildPreview(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildRole

GetGuildRole executes GET /guilds/{guild_id}/roles/{role_id}

```go
func (*RESTClient) GetGuildRole(ctx context.Context, guildid models.SnowflakeType, roleid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `roleid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildScheduledEvent

GetGuildScheduledEvent executes GET /guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}

```go
func (*RESTClient) GetGuildScheduledEvent(ctx context.Context, guildid models.SnowflakeType, guildscheduledeventid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `guildscheduledeventid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildSoundboardSound

GetGuildSoundboardSound executes GET /guilds/{guild_id}/soundboard-sounds/{sound_id}

```go
func (*RESTClient) GetGuildSoundboardSound(ctx context.Context, guildid models.SnowflakeType, soundid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `soundid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildSticker

GetGuildSticker executes GET /guilds/{guild_id}/stickers/{sticker_id}

```go
func (*RESTClient) GetGuildSticker(ctx context.Context, guildid models.SnowflakeType, stickerid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `stickerid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildTemplate

GetGuildTemplate executes GET /guilds/templates/{code}

```go
func (*RESTClient) GetGuildTemplate(ctx context.Context, code string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `code` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildVanityURL

GetGuildVanityURL executes GET /guilds/{guild_id}/vanity-url

```go
func (*RESTClient) GetGuildVanityURL(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildWebhooks

GetGuildWebhooks executes GET /guilds/{guild_id}/webhooks

```go
func (*RESTClient) GetGuildWebhooks(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildWelcomeScreen

GetGuildWelcomeScreen executes GET /guilds/{guild_id}/welcome-screen

```go
func (*RESTClient) GetGuildWelcomeScreen(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildWidget

GetGuildWidget executes GET /guilds/{guild_id}/widget.json

```go
func (*RESTClient) GetGuildWidget(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildWidgetPng

GetGuildWidgetPng executes GET /guilds/{guild_id}/widget.png

```go
func (*RESTClient) GetGuildWidgetPng(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildWidgetSettings

GetGuildWidgetSettings executes GET /guilds/{guild_id}/widget

```go
func (*RESTClient) GetGuildWidgetSettings(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetGuildsOnboarding

GetGuildsOnboarding executes GET /guilds/{guild_id}/onboarding

```go
func (*RESTClient) GetGuildsOnboarding(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetLobby

GetLobby executes GET /lobbies/{lobby_id}

```go
func (*RESTClient) GetLobby(ctx context.Context, lobbyid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `lobbyid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetLobbyMessages

GetLobbyMessages executes GET /lobbies/{lobby_id}/messages

```go
func (*RESTClient) GetLobbyMessages(ctx context.Context, lobbyid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `lobbyid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetMessage

GetMessage executes GET /channels/{channel_id}/messages/{message_id}

```go
func (*RESTClient) GetMessage(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetMyApplication

GetMyApplication executes GET /applications/@me

```go
func (*RESTClient) GetMyApplication(ctx context.Context, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetMyGuildMember

GetMyGuildMember executes GET /users/@me/guilds/{guild_id}/member

```go
func (*RESTClient) GetMyGuildMember(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetMyOauth2Application

GetMyOauth2Application executes GET /oauth2/applications/@me

```go
func (*RESTClient) GetMyOauth2Application(ctx context.Context, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetMyOauth2Authorization

GetMyOauth2Authorization executes GET /oauth2/@me

```go
func (*RESTClient) GetMyOauth2Authorization(ctx context.Context, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetMyUser

GetMyUser executes GET /users/@me

```go
func (*RESTClient) GetMyUser(ctx context.Context, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetOpenidConnectUserinfo

GetOpenidConnectUserinfo executes GET /oauth2/userinfo

```go
func (*RESTClient) GetOpenidConnectUserinfo(ctx context.Context, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetOriginalWebhookMessage

GetOriginalWebhookMessage executes GET /webhooks/{webhook_id}/{webhook_token}/messages/@original

```go
func (*RESTClient) GetOriginalWebhookMessage(ctx context.Context, webhookid models.SnowflakeType, webhooktoken string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `webhooktoken` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetPublicKeys

GetPublicKeys executes GET /oauth2/keys

```go
func (*RESTClient) GetPublicKeys(ctx context.Context, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetSelfVoiceState

GetSelfVoiceState executes GET /guilds/{guild_id}/voice-states/@me

```go
func (*RESTClient) GetSelfVoiceState(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetSoundboardDefaultSounds

GetSoundboardDefaultSounds executes GET /soundboard-default-sounds

```go
func (*RESTClient) GetSoundboardDefaultSounds(ctx context.Context, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetStageInstance

GetStageInstance executes GET /stage-instances/{channel_id}

```go
func (*RESTClient) GetStageInstance(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetSticker

GetSticker executes GET /stickers/{sticker_id}

```go
func (*RESTClient) GetSticker(ctx context.Context, stickerid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `stickerid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetStickerPack

GetStickerPack executes GET /sticker-packs/{pack_id}

```go
func (*RESTClient) GetStickerPack(ctx context.Context, packid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `packid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetThreadMember

GetThreadMember executes GET /channels/{channel_id}/thread-members/{user_id}

```go
func (*RESTClient) GetThreadMember(ctx context.Context, channelid models.SnowflakeType, userid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetUser

GetUser executes GET /users/{user_id}

```go
func (*RESTClient) GetUser(ctx context.Context, userid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `userid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetVoiceState

GetVoiceState executes GET /guilds/{guild_id}/voice-states/{user_id}

```go
func (*RESTClient) GetVoiceState(ctx context.Context, guildid models.SnowflakeType, userid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetWebhook

GetWebhook executes GET /webhooks/{webhook_id}

```go
func (*RESTClient) GetWebhook(ctx context.Context, webhookid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetWebhookByToken

GetWebhookByToken executes GET /webhooks/{webhook_id}/{webhook_token}

```go
func (*RESTClient) GetWebhookByToken(ctx context.Context, webhookid models.SnowflakeType, webhooktoken string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `webhooktoken` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GetWebhookMessage

GetWebhookMessage executes GET /webhooks/{webhook_id}/{webhook_token}/messages/{message_id}

```go
func (*RESTClient) GetWebhookMessage(ctx context.Context, webhookid models.SnowflakeType, webhooktoken string, messageid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `webhooktoken` (string)
- `messageid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### GuildRoleMemberCounts

GuildRoleMemberCounts executes GET /guilds/{guild_id}/roles/member-counts

```go
func (*RESTClient) GuildRoleMemberCounts(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### InviteResolve

InviteResolve executes GET /invites/{code}

```go
func (*RESTClient) InviteResolve(ctx context.Context, code string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `code` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### InviteRevoke

InviteRevoke executes DELETE /invites/{code}

```go
func (*RESTClient) InviteRevoke(ctx context.Context, code string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `code` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### JoinThread

JoinThread executes PUT /channels/{channel_id}/thread-members/@me

```go
func (*RESTClient) JoinThread(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### LeaveGuild

LeaveGuild executes DELETE /users/@me/guilds/{guild_id}

```go
func (*RESTClient) LeaveGuild(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### LeaveLobby

LeaveLobby executes DELETE /lobbies/{lobby_id}/members/@me

```go
func (*RESTClient) LeaveLobby(ctx context.Context, lobbyid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `lobbyid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### LeaveThread

LeaveThread executes DELETE /channels/{channel_id}/thread-members/@me

```go
func (*RESTClient) LeaveThread(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListApplicationCommands

ListApplicationCommands executes GET /applications/{application_id}/commands

```go
func (*RESTClient) ListApplicationCommands(ctx context.Context, applicationid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListApplicationEmojis

ListApplicationEmojis executes GET /applications/{application_id}/emojis

```go
func (*RESTClient) ListApplicationEmojis(ctx context.Context, applicationid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListAutoModerationRules

ListAutoModerationRules executes GET /guilds/{guild_id}/auto-moderation/rules

```go
func (*RESTClient) ListAutoModerationRules(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListChannelInvites

ListChannelInvites executes GET /channels/{channel_id}/invites

```go
func (*RESTClient) ListChannelInvites(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListChannelWebhooks

ListChannelWebhooks executes GET /channels/{channel_id}/webhooks

```go
func (*RESTClient) ListChannelWebhooks(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildApplicationCommandPermissions

ListGuildApplicationCommandPermissions executes GET /applications/{application_id}/guilds/{guild_id}/commands/permissions

```go
func (*RESTClient) ListGuildApplicationCommandPermissions(ctx context.Context, applicationid models.SnowflakeType, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildApplicationCommands

ListGuildApplicationCommands executes GET /applications/{application_id}/guilds/{guild_id}/commands

```go
func (*RESTClient) ListGuildApplicationCommands(ctx context.Context, applicationid models.SnowflakeType, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildAuditLogEntries

ListGuildAuditLogEntries executes GET /guilds/{guild_id}/audit-logs

```go
func (*RESTClient) ListGuildAuditLogEntries(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildBans

ListGuildBans executes GET /guilds/{guild_id}/bans

```go
func (*RESTClient) ListGuildBans(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildChannels

ListGuildChannels executes GET /guilds/{guild_id}/channels

```go
func (*RESTClient) ListGuildChannels(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildEmojis

ListGuildEmojis executes GET /guilds/{guild_id}/emojis

```go
func (*RESTClient) ListGuildEmojis(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildIntegrations

ListGuildIntegrations executes GET /guilds/{guild_id}/integrations

```go
func (*RESTClient) ListGuildIntegrations(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildInvites

ListGuildInvites executes GET /guilds/{guild_id}/invites

```go
func (*RESTClient) ListGuildInvites(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildMembers

ListGuildMembers executes GET /guilds/{guild_id}/members

```go
func (*RESTClient) ListGuildMembers(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildRoles

ListGuildRoles executes GET /guilds/{guild_id}/roles

```go
func (*RESTClient) ListGuildRoles(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildScheduledEventUsers

ListGuildScheduledEventUsers executes GET /guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}/users

```go
func (*RESTClient) ListGuildScheduledEventUsers(ctx context.Context, guildid models.SnowflakeType, guildscheduledeventid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `guildscheduledeventid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildScheduledEvents

ListGuildScheduledEvents executes GET /guilds/{guild_id}/scheduled-events

```go
func (*RESTClient) ListGuildScheduledEvents(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildSoundboardSounds

ListGuildSoundboardSounds executes GET /guilds/{guild_id}/soundboard-sounds

```go
func (*RESTClient) ListGuildSoundboardSounds(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildStickers

ListGuildStickers executes GET /guilds/{guild_id}/stickers

```go
func (*RESTClient) ListGuildStickers(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildTemplates

ListGuildTemplates executes GET /guilds/{guild_id}/templates

```go
func (*RESTClient) ListGuildTemplates(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListGuildVoiceRegions

ListGuildVoiceRegions executes GET /guilds/{guild_id}/regions

```go
func (*RESTClient) ListGuildVoiceRegions(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListMessageReactionsByEmoji

ListMessageReactionsByEmoji executes GET /channels/{channel_id}/messages/{message_id}/reactions/{emoji_name}

```go
func (*RESTClient) ListMessageReactionsByEmoji(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, emojiname string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `emojiname` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListMessages

ListMessages executes GET /channels/{channel_id}/messages

```go
func (*RESTClient) ListMessages(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListMyConnections

ListMyConnections executes GET /users/@me/connections

```go
func (*RESTClient) ListMyConnections(ctx context.Context, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListMyGuilds

ListMyGuilds executes GET /users/@me/guilds

```go
func (*RESTClient) ListMyGuilds(ctx context.Context, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListMyPrivateArchivedThreads

ListMyPrivateArchivedThreads executes GET /channels/{channel_id}/users/@me/threads/archived/private

```go
func (*RESTClient) ListMyPrivateArchivedThreads(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListPins

ListPins executes GET /channels/{channel_id}/messages/pins

```go
func (*RESTClient) ListPins(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListPrivateArchivedThreads

ListPrivateArchivedThreads executes GET /channels/{channel_id}/threads/archived/private

```go
func (*RESTClient) ListPrivateArchivedThreads(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListPublicArchivedThreads

ListPublicArchivedThreads executes GET /channels/{channel_id}/threads/archived/public

```go
func (*RESTClient) ListPublicArchivedThreads(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListStickerPacks

ListStickerPacks executes GET /sticker-packs

```go
func (*RESTClient) ListStickerPacks(ctx context.Context, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListThreadMembers

ListThreadMembers executes GET /channels/{channel_id}/thread-members

```go
func (*RESTClient) ListThreadMembers(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ListVoiceRegions

ListVoiceRegions executes GET /voice/regions

```go
func (*RESTClient) ListVoiceRegions(ctx context.Context, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### PartnerSdkToken

PartnerSdkToken executes POST /partner-sdk/token

```go
func (*RESTClient) PartnerSdkToken(ctx context.Context, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### PartnerSdkUnmergeProvisionalAccount

PartnerSdkUnmergeProvisionalAccount executes POST /partner-sdk/provisional-accounts/unmerge

```go
func (*RESTClient) PartnerSdkUnmergeProvisionalAccount(ctx context.Context, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### PollExpire

PollExpire executes POST /channels/{channel_id}/polls/{message_id}/expire

```go
func (*RESTClient) PollExpire(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### PreviewPruneGuild

PreviewPruneGuild executes GET /guilds/{guild_id}/prune

```go
func (*RESTClient) PreviewPruneGuild(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### PruneGuild

PruneGuild executes POST /guilds/{guild_id}/prune

```go
func (*RESTClient) PruneGuild(ctx context.Context, guildid models.SnowflakeType, body models.PruneGuildOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (models.PruneGuildOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### PutGuildsOnboarding

PutGuildsOnboarding executes PUT /guilds/{guild_id}/onboarding

```go
func (*RESTClient) PutGuildsOnboarding(ctx context.Context, guildid models.SnowflakeType, body models.UpdateGuildOnboardingOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (models.UpdateGuildOnboardingOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### SearchGuildMembers

SearchGuildMembers executes GET /guilds/{guild_id}/members/search

```go
func (*RESTClient) SearchGuildMembers(ctx context.Context, guildid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### SendSoundboardSound

SendSoundboardSound executes POST /channels/{channel_id}/send-soundboard-sound

```go
func (*RESTClient) SendSoundboardSound(ctx context.Context, channelid models.SnowflakeType, body models.SoundboardSoundSendOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `body` (models.SoundboardSoundSendOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### SetChannelPermissionOverwrite

SetChannelPermissionOverwrite executes PUT /channels/{channel_id}/permissions/{overwrite_id}

```go
func (*RESTClient) SetChannelPermissionOverwrite(ctx context.Context, channelid models.SnowflakeType, overwriteid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `overwriteid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### SetGuildApplicationCommandPermissions

SetGuildApplicationCommandPermissions executes PUT /applications/{application_id}/guilds/{guild_id}/commands/{command_id}/permissions

```go
func (*RESTClient) SetGuildApplicationCommandPermissions(ctx context.Context, applicationid models.SnowflakeType, guildid models.SnowflakeType, commandid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `guildid` (models.SnowflakeType)
- `commandid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### SyncGuildTemplate

SyncGuildTemplate executes PUT /guilds/{guild_id}/templates/{code}

```go
func (*RESTClient) SyncGuildTemplate(ctx context.Context, guildid models.SnowflakeType, code string, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `code` (string)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### ThreadSearch

ThreadSearch executes GET /channels/{channel_id}/threads/search

```go
func (*RESTClient) ThreadSearch(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### TriggerTypingIndicator

TriggerTypingIndicator executes POST /channels/{channel_id}/typing

```go
func (*RESTClient) TriggerTypingIndicator(ctx context.Context, channelid models.SnowflakeType, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UnbanUserFromGuild

UnbanUserFromGuild executes DELETE /guilds/{guild_id}/bans/{user_id}

```go
func (*RESTClient) UnbanUserFromGuild(ctx context.Context, guildid models.SnowflakeType, userid models.SnowflakeType, body models.UnbanUserFromGuildOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `body` (models.UnbanUserFromGuildOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateApplication

UpdateApplication executes PATCH /applications/{application_id}

```go
func (*RESTClient) UpdateApplication(ctx context.Context, applicationid models.SnowflakeType, body models.ApplicationFormPartial, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `body` (models.ApplicationFormPartial)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateApplicationCommand

UpdateApplicationCommand executes PATCH /applications/{application_id}/commands/{command_id}

```go
func (*RESTClient) UpdateApplicationCommand(ctx context.Context, applicationid models.SnowflakeType, commandid models.SnowflakeType, body models.ApplicationCommandPatchOptionsPartial, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `commandid` (models.SnowflakeType)
- `body` (models.ApplicationCommandPatchOptionsPartial)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateApplicationEmoji

UpdateApplicationEmoji executes PATCH /applications/{application_id}/emojis/{emoji_id}

```go
func (*RESTClient) UpdateApplicationEmoji(ctx context.Context, applicationid models.SnowflakeType, emojiid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `emojiid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateApplicationRoleConnectionsMetadata

UpdateApplicationRoleConnectionsMetadata executes PUT /applications/{application_id}/role-connections/metadata

```go
func (*RESTClient) UpdateApplicationRoleConnectionsMetadata(ctx context.Context, applicationid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateApplicationUserRoleConnection

UpdateApplicationUserRoleConnection executes PUT /users/@me/applications/{application_id}/role-connection

```go
func (*RESTClient) UpdateApplicationUserRoleConnection(ctx context.Context, applicationid models.SnowflakeType, body models.UpdateApplicationUserRoleConnectionOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `body` (models.UpdateApplicationUserRoleConnectionOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateAutoModerationRule

UpdateAutoModerationRule executes PATCH /guilds/{guild_id}/auto-moderation/rules/{rule_id}

```go
func (*RESTClient) UpdateAutoModerationRule(ctx context.Context, guildid models.SnowflakeType, ruleid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `ruleid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateChannel

UpdateChannel executes PATCH /channels/{channel_id}

```go
func (*RESTClient) UpdateChannel(ctx context.Context, channelid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateGuild

UpdateGuild executes PATCH /guilds/{guild_id}

```go
func (*RESTClient) UpdateGuild(ctx context.Context, guildid models.SnowflakeType, body models.GuildPatchOptionsPartial, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (models.GuildPatchOptionsPartial)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateGuildApplicationCommand

UpdateGuildApplicationCommand executes PATCH /applications/{application_id}/guilds/{guild_id}/commands/{command_id}

```go
func (*RESTClient) UpdateGuildApplicationCommand(ctx context.Context, applicationid models.SnowflakeType, guildid models.SnowflakeType, commandid models.SnowflakeType, body models.ApplicationCommandPatchOptionsPartial, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `guildid` (models.SnowflakeType)
- `commandid` (models.SnowflakeType)
- `body` (models.ApplicationCommandPatchOptionsPartial)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateGuildEmoji

UpdateGuildEmoji executes PATCH /guilds/{guild_id}/emojis/{emoji_id}

```go
func (*RESTClient) UpdateGuildEmoji(ctx context.Context, guildid models.SnowflakeType, emojiid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `emojiid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateGuildMember

UpdateGuildMember executes PATCH /guilds/{guild_id}/members/{user_id}

```go
func (*RESTClient) UpdateGuildMember(ctx context.Context, guildid models.SnowflakeType, userid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateGuildRole

UpdateGuildRole executes PATCH /guilds/{guild_id}/roles/{role_id}

```go
func (*RESTClient) UpdateGuildRole(ctx context.Context, guildid models.SnowflakeType, roleid models.SnowflakeType, body models.UpdateRoleOptionsPartial, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `roleid` (models.SnowflakeType)
- `body` (models.UpdateRoleOptionsPartial)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateGuildScheduledEvent

UpdateGuildScheduledEvent executes PATCH /guilds/{guild_id}/scheduled-events/{guild_scheduled_event_id}

```go
func (*RESTClient) UpdateGuildScheduledEvent(ctx context.Context, guildid models.SnowflakeType, guildscheduledeventid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `guildscheduledeventid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateGuildSoundboardSound

UpdateGuildSoundboardSound executes PATCH /guilds/{guild_id}/soundboard-sounds/{sound_id}

```go
func (*RESTClient) UpdateGuildSoundboardSound(ctx context.Context, guildid models.SnowflakeType, soundid models.SnowflakeType, body models.SoundboardPatchOptionsPartial, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `soundid` (models.SnowflakeType)
- `body` (models.SoundboardPatchOptionsPartial)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateGuildSticker

UpdateGuildSticker executes PATCH /guilds/{guild_id}/stickers/{sticker_id}

```go
func (*RESTClient) UpdateGuildSticker(ctx context.Context, guildid models.SnowflakeType, stickerid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `stickerid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateGuildTemplate

UpdateGuildTemplate executes PATCH /guilds/{guild_id}/templates/{code}

```go
func (*RESTClient) UpdateGuildTemplate(ctx context.Context, guildid models.SnowflakeType, code string, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `code` (string)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateGuildWelcomeScreen

UpdateGuildWelcomeScreen executes PATCH /guilds/{guild_id}/welcome-screen

```go
func (*RESTClient) UpdateGuildWelcomeScreen(ctx context.Context, guildid models.SnowflakeType, body models.WelcomeScreenPatchOptionsPartial, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (models.WelcomeScreenPatchOptionsPartial)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateGuildWidgetSettings

UpdateGuildWidgetSettings executes PATCH /guilds/{guild_id}/widget

```go
func (*RESTClient) UpdateGuildWidgetSettings(ctx context.Context, guildid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateMessage

UpdateMessage executes PATCH /channels/{channel_id}/messages/{message_id}

```go
func (*RESTClient) UpdateMessage(ctx context.Context, channelid models.SnowflakeType, messageid models.SnowflakeType, body models.MessageEditOptionsPartial, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `messageid` (models.SnowflakeType)
- `body` (models.MessageEditOptionsPartial)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateMyApplication

UpdateMyApplication executes PATCH /applications/@me

```go
func (*RESTClient) UpdateMyApplication(ctx context.Context, body models.ApplicationFormPartial, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `body` (models.ApplicationFormPartial)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateMyGuildMember

UpdateMyGuildMember executes PATCH /guilds/{guild_id}/members/@me

```go
func (*RESTClient) UpdateMyGuildMember(ctx context.Context, guildid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateMyUser

UpdateMyUser executes PATCH /users/@me

```go
func (*RESTClient) UpdateMyUser(ctx context.Context, body models.BotAccountPatchOptions, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `body` (models.BotAccountPatchOptions)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateOriginalWebhookMessage

UpdateOriginalWebhookMessage executes PATCH /webhooks/{webhook_id}/{webhook_token}/messages/@original

```go
func (*RESTClient) UpdateOriginalWebhookMessage(ctx context.Context, webhookid models.SnowflakeType, webhooktoken string, body models.IncomingWebhookUpdateOptionsPartial, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `webhooktoken` (string)
- `body` (models.IncomingWebhookUpdateOptionsPartial)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateSelfVoiceState

UpdateSelfVoiceState executes PATCH /guilds/{guild_id}/voice-states/@me

```go
func (*RESTClient) UpdateSelfVoiceState(ctx context.Context, guildid models.SnowflakeType, body models.UpdateSelfVoiceStateOptionsPartial, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `body` (models.UpdateSelfVoiceStateOptionsPartial)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateStageInstance

UpdateStageInstance executes PATCH /stage-instances/{channel_id}

```go
func (*RESTClient) UpdateStageInstance(ctx context.Context, channelid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `channelid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateVoiceState

UpdateVoiceState executes PATCH /guilds/{guild_id}/voice-states/{user_id}

```go
func (*RESTClient) UpdateVoiceState(ctx context.Context, guildid models.SnowflakeType, userid models.SnowflakeType, body models.UpdateVoiceStateOptionsPartial, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `guildid` (models.SnowflakeType)
- `userid` (models.SnowflakeType)
- `body` (models.UpdateVoiceStateOptionsPartial)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateWebhook

UpdateWebhook executes PATCH /webhooks/{webhook_id}

```go
func (*RESTClient) UpdateWebhook(ctx context.Context, webhookid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateWebhookByToken

UpdateWebhookByToken executes PATCH /webhooks/{webhook_id}/{webhook_token}

```go
func (*RESTClient) UpdateWebhookByToken(ctx context.Context, webhookid models.SnowflakeType, webhooktoken string, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `webhooktoken` (string)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UpdateWebhookMessage

UpdateWebhookMessage executes PATCH /webhooks/{webhook_id}/{webhook_token}/messages/{message_id}

```go
func (*RESTClient) UpdateWebhookMessage(ctx context.Context, webhookid models.SnowflakeType, webhooktoken string, messageid models.SnowflakeType, body models.IncomingWebhookUpdateOptionsPartial, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `webhookid` (models.SnowflakeType)
- `webhooktoken` (string)
- `messageid` (models.SnowflakeType)
- `body` (models.IncomingWebhookUpdateOptionsPartial)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### UploadApplicationAttachment

UploadApplicationAttachment executes POST /applications/{application_id}/attachment

```go
func (*RESTClient) UploadApplicationAttachment(ctx context.Context, applicationid models.SnowflakeType, body any, opts ...*neuron.RequestOptions) (**ast.IndexExpr, error)
```

**Parameters:**
- `ctx` (context.Context)
- `applicationid` (models.SnowflakeType)
- `body` (any)
- `opts` (...*neuron.RequestOptions)

**Returns:**
- **ast.IndexExpr
- error

### WithContext

WithContext returns a context-aware request builder

```go
func (*RESTClient) WithContext(ctx context.Context) *RequestBuilder
```

**Parameters:**
- `ctx` (context.Context)

**Returns:**
- *RequestBuilder

### RequestBuilder
RequestBuilder helps build requests with context

#### Example Usage

```go
// Create a new RequestBuilder
requestbuilder := RequestBuilder{

}
```

#### Type Definition

```go
type RequestBuilder struct {
}
```

## Methods

### Options

Options returns RequestOptions for the builder

```go
func (*RequestBuilder) Options() *neuron.RequestOptions
```

**Parameters:**
  None

**Returns:**
- *neuron.RequestOptions

### WithHeader

WithHeader adds a header to the request

```go
func (*RequestBuilder) WithHeader(key, value string) *RequestBuilder
```

**Parameters:**
- `key` (string)
- `value` (string)

**Returns:**
- *RequestBuilder

### WithQuery

WithQuery adds a query parameter

```go
func (*RequestBuilder) WithQuery(key string, value any) *RequestBuilder
```

**Parameters:**
- `key` (string)
- `value` (any)

**Returns:**
- *RequestBuilder

### SecurityScheme
SecurityScheme defines a security scheme

#### Example Usage

```go
// Create a new SecurityScheme
securityscheme := SecurityScheme{
    Type: "example",
    Scheme: "example",
    BearerFormat: "example",
    In: "example",
    Name: "example",
}
```

#### Type Definition

```go
type SecurityScheme struct {
    Type string
    Scheme string
    BearerFormat string
    In string
    Name string
}
```

### Fields

| Field | Type | Description |
| ----- | ---- | ----------- |
| Type | `string` |  |
| Scheme | `string` |  |
| BearerFormat | `string` |  |
| In | `string` |  |
| Name | `string` |  |

## Functions

### ApplyAPIKey
ApplyAPIKey adds an API key to the request

```go
func ApplyAPIKey(name, value, location string) func(*http.Request) error
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `name` | `string` | |
| `value` | `string` | |
| `location` | `string` | |

**Returns:**
| Type | Description |
|------|-------------|
| `func(*http.Request) error` | |

**Example:**

```go
// Example usage of ApplyAPIKey
result := ApplyAPIKey(/* parameters */)
```

### ApplyBasicAuth
ApplyBasicAuth adds basic authentication to the request

```go
func ApplyBasicAuth(username, password string) func(*http.Request) error
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `username` | `string` | |
| `password` | `string` | |

**Returns:**
| Type | Description |
|------|-------------|
| `func(*http.Request) error` | |

**Example:**

```go
// Example usage of ApplyBasicAuth
result := ApplyBasicAuth(/* parameters */)
```

### ApplyBearerToken
ApplyBearerToken adds a bearer token to the request headers

```go
func ApplyBearerToken(token string) func(*http.Request) error
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `token` | `string` | |

**Returns:**
| Type | Description |
|------|-------------|
| `func(*http.Request) error` | |

**Example:**

```go
// Example usage of ApplyBearerToken
result := ApplyBearerToken(/* parameters */)
```

## External Links

- [Package Overview](../packages/client.md)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/kolosys/discord/rest/internal)
- [Source Code](https://github.com/kolosys/discord/tree/main/rest/internal)

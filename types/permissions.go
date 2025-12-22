package types

import (
	"encoding/json"
	"strconv"
)

// Permissions represents a Discord permission bitfield.
type Permissions uint64

const (
	PermissionCreateInstantInvite Permissions = 1 << 0
	PermissionKickMembers         Permissions = 1 << 1
	PermissionBanMembers          Permissions = 1 << 2
	PermissionAdministrator       Permissions = 1 << 3
	PermissionManageChannels      Permissions = 1 << 4
	PermissionManageGuild         Permissions = 1 << 5
	PermissionAddReactions        Permissions = 1 << 6
	PermissionViewAuditLog        Permissions = 1 << 7
	PermissionPrioritySpeaker     Permissions = 1 << 8
	PermissionStream              Permissions = 1 << 9
	PermissionViewChannel         Permissions = 1 << 10
	PermissionSendMessages        Permissions = 1 << 11
	PermissionSendTTSMessages     Permissions = 1 << 12
	PermissionManageMessages      Permissions = 1 << 13
	PermissionEmbedLinks          Permissions = 1 << 14
	PermissionAttachFiles         Permissions = 1 << 15
	PermissionReadMessageHistory  Permissions = 1 << 16
	PermissionMentionEveryone     Permissions = 1 << 17
	PermissionUseExternalEmojis   Permissions = 1 << 18
	PermissionViewGuildInsights   Permissions = 1 << 19
	PermissionConnect             Permissions = 1 << 20
	PermissionSpeak               Permissions = 1 << 21
	PermissionMuteMembers         Permissions = 1 << 22
	PermissionDeafenMembers       Permissions = 1 << 23
	PermissionMoveMembers         Permissions = 1 << 24
	PermissionUseVAD              Permissions = 1 << 25
	PermissionChangeNickname      Permissions = 1 << 26
	PermissionManageNicknames     Permissions = 1 << 27
	PermissionManageRoles         Permissions = 1 << 28
	PermissionManageWebhooks      Permissions = 1 << 29
	PermissionManageGuildExpressions Permissions = 1 << 30
	PermissionUseApplicationCommands Permissions = 1 << 31
	PermissionRequestToSpeak         Permissions = 1 << 32
	PermissionManageEvents           Permissions = 1 << 33
	PermissionManageThreads          Permissions = 1 << 34
	PermissionCreatePublicThreads    Permissions = 1 << 35
	PermissionCreatePrivateThreads   Permissions = 1 << 36
	PermissionUseExternalStickers    Permissions = 1 << 37
	PermissionSendMessagesInThreads  Permissions = 1 << 38
	PermissionUseEmbeddedActivities  Permissions = 1 << 39
	PermissionModerateMembers        Permissions = 1 << 40
	PermissionViewCreatorMonetizationAnalytics Permissions = 1 << 41
	PermissionUseSoundboard          Permissions = 1 << 42
	PermissionCreateGuildExpressions Permissions = 1 << 43
	PermissionCreateEvents           Permissions = 1 << 44
	PermissionUseExternalSounds      Permissions = 1 << 45
	PermissionSendVoiceMessages      Permissions = 1 << 46
	PermissionSendPolls              Permissions = 1 << 49
	PermissionUseExternalApps        Permissions = 1 << 50

	PermissionAllText = PermissionViewChannel |
		PermissionSendMessages |
		PermissionSendTTSMessages |
		PermissionManageMessages |
		PermissionEmbedLinks |
		PermissionAttachFiles |
		PermissionReadMessageHistory |
		PermissionMentionEveryone |
		PermissionUseExternalEmojis |
		PermissionAddReactions

	PermissionAllVoice = PermissionConnect |
		PermissionSpeak |
		PermissionMuteMembers |
		PermissionDeafenMembers |
		PermissionMoveMembers |
		PermissionUseVAD |
		PermissionPrioritySpeaker |
		PermissionStream

	PermissionAllChannel = PermissionAllText | PermissionAllVoice | PermissionManageChannels

	PermissionAll = PermissionAllChannel |
		PermissionKickMembers |
		PermissionBanMembers |
		PermissionManageGuild |
		PermissionAdministrator |
		PermissionManageRoles |
		PermissionManageWebhooks |
		PermissionManageGuildExpressions
)

func (p Permissions) Has(perm Permissions) bool {
	return (p & perm) == perm
}

func (p Permissions) Add(perm Permissions) Permissions {
	return p | perm
}

func (p Permissions) Remove(perm Permissions) Permissions {
	return p &^ perm
}

func (p Permissions) Toggle(perm Permissions) Permissions {
	return p ^ perm
}

func (p *Permissions) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		var num uint64
		if err := json.Unmarshal(data, &num); err != nil {
			return err
		}
		*p = Permissions(num)
		return nil
	}
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*p = Permissions(val)
	return nil
}

func (p Permissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(p), 10))
}

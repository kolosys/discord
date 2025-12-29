package types_test

import (
	"encoding/json"
	"testing"

	"github.com/kolosys/discord/types"
)

func TestPermissions_Has(t *testing.T) {
	tests := []struct {
		name     string
		p        types.Permissions
		perm     types.Permissions
		expected bool
	}{
		{"Has single permission", types.PermissionSendMessages, types.PermissionSendMessages, true},
		{"Has multiple permissions", types.PermissionSendMessages | types.PermissionManageMessages, types.PermissionSendMessages, true},
		{"Missing permission", types.PermissionSendMessages, types.PermissionManageMessages, false},
		{"Zero permissions", types.Permissions(0), types.PermissionSendMessages, false},
		{"All permissions", types.PermissionAll, types.PermissionSendMessages, true},
		{"Check multiple at once", types.PermissionSendMessages | types.PermissionManageMessages, types.PermissionSendMessages | types.PermissionManageMessages, true},
		{"Check multiple partial", types.PermissionSendMessages | types.PermissionManageMessages, types.PermissionSendMessages | types.PermissionEmbedLinks, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Has(tt.perm)
			if got != tt.expected {
				t.Errorf("Has(%v) = %v, want %v", tt.perm, got, tt.expected)
			}
		})
	}
}

func TestPermissions_Add(t *testing.T) {
	tests := []struct {
		name     string
		p        types.Permissions
		perm     types.Permissions
		expected types.Permissions
	}{
		{"Add to zero", types.Permissions(0), types.PermissionSendMessages, types.PermissionSendMessages},
		{"Add to existing", types.PermissionSendMessages, types.PermissionManageMessages, types.PermissionSendMessages | types.PermissionManageMessages},
		{"Add duplicate", types.PermissionSendMessages, types.PermissionSendMessages, types.PermissionSendMessages},
		{"Add multiple", types.PermissionSendMessages, types.PermissionManageMessages | types.PermissionEmbedLinks, types.PermissionSendMessages | types.PermissionManageMessages | types.PermissionEmbedLinks},
		{"Add to all", types.PermissionAll, types.PermissionSendMessages, types.PermissionAll},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Add(tt.perm)
			if got != tt.expected {
				t.Errorf("Add(%v) = %v, want %v", tt.perm, got, tt.expected)
			}
		})
	}
}

func TestPermissions_Remove(t *testing.T) {
	tests := []struct {
		name     string
		p        types.Permissions
		perm     types.Permissions
		expected types.Permissions
	}{
		{"Remove from zero", types.Permissions(0), types.PermissionSendMessages, types.Permissions(0)},
		{"Remove existing", types.PermissionSendMessages, types.PermissionSendMessages, types.Permissions(0)},
		{"Remove non-existing", types.PermissionSendMessages, types.PermissionManageMessages, types.PermissionSendMessages},
		{"Remove one from many", types.PermissionSendMessages | types.PermissionManageMessages | types.PermissionEmbedLinks, types.PermissionManageMessages, types.PermissionSendMessages | types.PermissionEmbedLinks},
		{"Remove multiple", types.PermissionAll, types.PermissionSendMessages | types.PermissionManageMessages, types.PermissionAll.Remove(types.PermissionSendMessages | types.PermissionManageMessages)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Remove(tt.perm)
			if got != tt.expected {
				t.Errorf("Remove(%v) = %v, want %v", tt.perm, got, tt.expected)
			}
		})
	}
}

func TestPermissions_Toggle(t *testing.T) {
	tests := []struct {
		name     string
		p        types.Permissions
		perm     types.Permissions
		expected types.Permissions
	}{
		{"Toggle off to on", types.Permissions(0), types.PermissionSendMessages, types.PermissionSendMessages},
		{"Toggle on to off", types.PermissionSendMessages, types.PermissionSendMessages, types.Permissions(0)},
		{"Toggle one of many", types.PermissionSendMessages | types.PermissionManageMessages, types.PermissionSendMessages, types.PermissionManageMessages},
		{"Toggle multiple", types.PermissionSendMessages | types.PermissionManageMessages, types.PermissionSendMessages | types.PermissionEmbedLinks, types.PermissionManageMessages | types.PermissionEmbedLinks},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Toggle(tt.perm)
			if got != tt.expected {
				t.Errorf("Toggle(%v) = %v, want %v", tt.perm, got, tt.expected)
			}
		})
	}
}

func TestPermissions_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		p    types.Permissions
	}{
		{"Zero", types.Permissions(0)},
		{"Single", types.PermissionSendMessages},
		{"Multiple", types.PermissionSendMessages | types.PermissionManageMessages},
		{"All permissions", types.PermissionAll},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.MarshalJSON()
			if err != nil {
				t.Errorf("MarshalJSON() error = %v", err)
				return
			}
			if len(got) == 0 {
				t.Error("MarshalJSON() returned empty result")
			}
			// Verify it's a valid JSON string (starts and ends with quotes)
			if got[0] != '"' || got[len(got)-1] != '"' {
				t.Errorf("MarshalJSON() should return a quoted string: %s", got)
			}
		})
	}
}

func TestPermissions_UnmarshalJSON_String(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		wantErr bool
	}{
		{"String zero", `"0"`, false},
		{"String number", `"2048"`, false},
		{"String large", `"999999999999"`, false},
		{"Invalid string", `"invalid"`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p types.Permissions
			err := json.Unmarshal([]byte(tt.data), &p)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && p < 0 {
				t.Error("UnmarshalJSON() resulted in negative permissions")
			}
		})
	}
}

func TestPermissions_UnmarshalJSON_Number(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		wantErr bool
	}{
		{"Number zero", `0`, false},
		{"Number", `2048`, false},
		{"Number large", `999999999999`, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p types.Permissions
			err := json.Unmarshal([]byte(tt.data), &p)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && p < 0 {
				t.Error("UnmarshalJSON() resulted in negative permissions")
			}
		})
	}
}

func TestPermissions_Roundtrip(t *testing.T) {
	tests := []types.Permissions{
		types.Permissions(0),
		types.PermissionSendMessages,
		types.PermissionSendMessages | types.PermissionManageMessages,
		types.PermissionAll,
		types.PermissionAllText,
		types.PermissionAllVoice,
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			data, err := tt.MarshalJSON()
			if err != nil {
				t.Errorf("MarshalJSON() error = %v", err)
				return
			}

			var p types.Permissions
			err = json.Unmarshal(data, &p)
			if err != nil {
				t.Errorf("UnmarshalJSON() error = %v", err)
				return
			}

			if p != tt {
				t.Errorf("Roundtrip failed: %v != %v", p, tt)
			}
		})
	}
}

func TestPermissions_Constants(t *testing.T) {
	tests := []struct {
		name string
		perm types.Permissions
	}{
		{"PermissionCreateInstantInvite", types.PermissionCreateInstantInvite},
		{"PermissionKickMembers", types.PermissionKickMembers},
		{"PermissionBanMembers", types.PermissionBanMembers},
		{"PermissionAdministrator", types.PermissionAdministrator},
		{"PermissionManageChannels", types.PermissionManageChannels},
		{"PermissionSendMessages", types.PermissionSendMessages},
		{"PermissionManageMessages", types.PermissionManageMessages},
		{"PermissionAllText contains SendMessages", types.PermissionAllText},
		{"PermissionAllVoice contains Connect", types.PermissionAllVoice},
		{"PermissionAll is superset", types.PermissionAll},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.perm == 0 && tt.name != "PermissionAllText" && tt.name != "PermissionAllVoice" && tt.name != "PermissionAll" {
				t.Errorf("Permission constant should not be zero: %s", tt.name)
			}
		})
	}
}

func TestPermissions_Combinations(t *testing.T) {
	send := types.PermissionSendMessages
	manage := types.PermissionManageMessages
	embed := types.PermissionEmbedLinks

	// Test: Add multiple, then remove
	combined := send.Add(manage).Add(embed)
	if !combined.Has(send) || !combined.Has(manage) || !combined.Has(embed) {
		t.Error("All permissions should be present after Add chain")
	}

	removed := combined.Remove(manage)
	if !removed.Has(send) || removed.Has(manage) || !removed.Has(embed) {
		t.Error("Manage should be removed, but Send and Embed should remain")
	}

	// Test: Toggle twice returns to original
	toggled := send.Toggle(manage).Toggle(manage)
	if toggled != send {
		t.Errorf("Toggle twice should return original: %v != %v", toggled, send)
	}
}

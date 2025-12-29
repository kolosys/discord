package components_test

import (
	"testing"

	"github.com/kolosys/discord/components"
	"github.com/kolosys/discord/types"
)

func TestStringSelect_Valid(t *testing.T) {
	b := components.StringSelect("select_id")
	b.Option("Option 1", "value_1")
	b.Option("Option 2", "value_2")

	err := b.Validate()
	if err != nil {
		t.Errorf("Validate() error = %v", err)
	}
}

func TestStringSelect_NoOptions(t *testing.T) {
	b := components.StringSelect("select_id")
	err := b.Validate()
	if err == nil {
		t.Error("Validate() should error for no options")
	}
}

func TestStringSelect_TooManyOptions(t *testing.T) {
	b := components.StringSelect("select_id")
	for i := 0; i <= components.MaxSelectOptionsPerMenu; i++ {
		b.Option("Option", "value")
	}
	err := b.Validate()
	if err == nil {
		t.Error("Validate() should error for too many options")
	}
}

func TestStringSelect_OptionLabelTooLong(t *testing.T) {
	b := components.StringSelect("select_id")
	b.Option(string(make([]byte, components.MaxSelectOptionLabelLength+1)), "value")
	err := b.Validate()
	if err != components.ErrSelectOptionLabelTooLong {
		t.Errorf("Validate() = %v, want ErrSelectOptionLabelTooLong", err)
	}
}

func TestStringSelect_OptionValueTooLong(t *testing.T) {
	b := components.StringSelect("select_id")
	b.Option("Label", string(make([]byte, components.MaxSelectOptionValueLength+1)))
	err := b.Validate()
	if err != components.ErrSelectOptionValueTooLong {
		t.Errorf("Validate() = %v, want ErrSelectOptionValueTooLong", err)
	}
}

func TestStringSelect_OptionDescriptionTooLong(t *testing.T) {
	b := components.StringSelect("select_id")
	opt := components.SelectOption{
		Label:       "Label",
		Value:       "value",
		Description: string(make([]byte, components.MaxSelectOptionDescriptionLength+1)),
	}
	b.Options(opt)
	err := b.Validate()
	if err != components.ErrSelectOptionDescTooLong {
		t.Errorf("Validate() = %v, want ErrSelectOptionDescTooLong", err)
	}
}

func TestStringSelect_PlaceholderTooLong(t *testing.T) {
	b := components.StringSelect("select_id")
	b.Option("Option", "value")
	b.Placeholder(string(make([]byte, components.MaxSelectPlaceholderLength+1)))
	err := b.Validate()
	if err != components.ErrSelectPlaceholderTooLong {
		t.Errorf("Validate() = %v, want ErrSelectPlaceholderTooLong", err)
	}
}

func TestStringSelect_MinMaxValues(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(*components.StringSelectBuilder)
		wantErr bool
	}{
		{"Valid min/max", func(b *components.StringSelectBuilder) {
			b.Option("O1", "v1").Option("O2", "v2").MinValues(1).MaxValues(2)
		}, false},
		{"Min zero", func(b *components.StringSelectBuilder) {
			b.Option("O1", "v1").MinValues(0)
		}, false},
		{"Max 25", func(b *components.StringSelectBuilder) {
			for i := 0; i < 25; i++ {
				b.Option("O", "v")
			}
			b.MaxValues(25)
		}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := components.StringSelect("id")
			tt.setup(b)
			err := b.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringSelect_Placeholder(t *testing.T) {
	b := components.StringSelect("id")
	b.Option("Option", "value")
	b.Placeholder("Select something...")
	sel := b.Build()
	if sel.Placeholder != "Select something..." {
		t.Errorf("Placeholder = %q, want %q", sel.Placeholder, "Select something...")
	}
}

func TestStringSelect_Disabled(t *testing.T) {
	tests := []bool{true, false}
	for _, disabled := range tests {
		b := components.StringSelect("id")
		b.Option("Option", "value").Disabled(disabled)
		sel := b.Build()
		if sel.Disabled != disabled {
			t.Errorf("Disabled = %v, want %v", sel.Disabled, disabled)
		}
	}
}

func TestStringSelect_Options(t *testing.T) {
	b := components.StringSelect("id")
	b.Option("Label 1", "value_1")
	b.Option("Label 2", "value_2")
	b.OptionWithDescription("Label 3", "value_3", "Description")

	sel := b.Build()
	if len(sel.Options) != 3 {
		t.Errorf("Options count = %d, want 3", len(sel.Options))
	}

	if sel.Options[0].Label != "Label 1" || sel.Options[0].Value != "value_1" {
		t.Error("Option 1 mismatch")
	}

	if sel.Options[2].Description != "Description" {
		t.Errorf("Option 3 description = %q, want %q", sel.Options[2].Description, "Description")
	}
}

func TestStringSelect_Chaining(t *testing.T) {
	b := components.StringSelect("id")
	result := b.Option("O1", "v1").Option("O2", "v2").Placeholder("Select")
	if result != b {
		t.Error("Builder methods should return self for chaining")
	}
}

func TestStringSelect_Build(t *testing.T) {
	b := components.StringSelect("test_id")
	b.Option("Option", "value")
	sel := b.Build()

	if sel == nil {
		t.Fatal("Build() returned nil")
	}
	if sel.CustomID != "test_id" {
		t.Errorf("CustomID = %q, want %q", sel.CustomID, "test_id")
	}
	if sel.Type != types.ComponentTypeStringSelect {
		t.Errorf("Type = %v, want %v", sel.Type, types.ComponentTypeStringSelect)
	}
}

func TestStringSelect_BuildSafe(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(*components.StringSelectBuilder)
		wantErr bool
	}{
		{"Valid", func(b *components.StringSelectBuilder) { b.Option("O", "v") }, false},
		{"No options", func(b *components.StringSelectBuilder) {}, true},
		{"Label too long", func(b *components.StringSelectBuilder) {
			b.Option(string(make([]byte, 101)), "v")
		}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := components.StringSelect("id")
			tt.setup(b)
			sel, err := b.BuildSafe()
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildSafe() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && sel == nil {
				t.Error("BuildSafe() returned nil for valid select")
			}
		})
	}
}

func TestUserSelect_Valid(t *testing.T) {
	b := components.UserSelect("user_select")
	err := b.Validate()
	if err != nil {
		t.Errorf("Validate() error = %v", err)
	}
}

func TestUserSelect_Placeholder(t *testing.T) {
	b := components.UserSelect("id")
	b.Placeholder("Select users...")
	sel := b.Build()
	if sel.Placeholder != "Select users..." {
		t.Errorf("Placeholder = %q, want %q", sel.Placeholder, "Select users...")
	}
}

func TestUserSelect_MinMaxValues(t *testing.T) {
	b := components.UserSelect("id")
	b.MinValues(1).MaxValues(5)
	sel := b.Build()

	if sel.MinValues == nil || *sel.MinValues != 1 {
		t.Error("MinValues mismatch")
	}
	if sel.MaxValues == nil || *sel.MaxValues != 5 {
		t.Error("MaxValues mismatch")
	}
}

func TestUserSelect_Disabled(t *testing.T) {
	b := components.UserSelect("id").Disabled(true)
	sel := b.Build()
	if !sel.Disabled {
		t.Error("Disabled should be true")
	}
}

func TestUserSelect_Build(t *testing.T) {
	b := components.UserSelect("user_id")
	sel := b.Build()

	if sel == nil {
		t.Fatal("Build() returned nil")
	}
	if sel.Type != types.ComponentTypeUserSelect {
		t.Errorf("Type = %v, want %v", sel.Type, types.ComponentTypeUserSelect)
	}
}

func TestRoleSelect_Valid(t *testing.T) {
	b := components.RoleSelect("role_select")
	err := b.Validate()
	if err != nil {
		t.Errorf("Validate() error = %v", err)
	}
}

func TestRoleSelect_Placeholder(t *testing.T) {
	b := components.RoleSelect("id")
	b.Placeholder("Select a role...")
	sel := b.Build()
	if sel.Placeholder != "Select a role..." {
		t.Errorf("Placeholder = %q, want %q", sel.Placeholder, "Select a role...")
	}
}

func TestRoleSelect_Build(t *testing.T) {
	b := components.RoleSelect("role_id")
	sel := b.Build()

	if sel == nil {
		t.Fatal("Build() returned nil")
	}
	if sel.Type != types.ComponentTypeRoleSelect {
		t.Errorf("Type = %v, want %v", sel.Type, types.ComponentTypeRoleSelect)
	}
}

func TestChannelSelect_Valid(t *testing.T) {
	b := components.ChannelSelect("channel_select")
	err := b.Validate()
	if err != nil {
		t.Errorf("Validate() error = %v", err)
	}
}

func TestChannelSelect_ChannelTypes(t *testing.T) {
	b := components.ChannelSelect("id")
	b.ChannelTypes(0, 1) // text, DM
	sel := b.Build()

	if len(sel.ChannelTypes) != 2 {
		t.Errorf("ChannelTypes count = %d, want 2", len(sel.ChannelTypes))
	}
}

func TestChannelSelect_Placeholder(t *testing.T) {
	b := components.ChannelSelect("id")
	b.Placeholder("Select a channel...")
	sel := b.Build()
	if sel.Placeholder != "Select a channel..." {
		t.Errorf("Placeholder = %q, want %q", sel.Placeholder, "Select a channel...")
	}
}

func TestChannelSelect_Build(t *testing.T) {
	b := components.ChannelSelect("channel_id")
	sel := b.Build()

	if sel == nil {
		t.Fatal("Build() returned nil")
	}
	if sel.Type != types.ComponentTypeChannelSelect {
		t.Errorf("Type = %v, want %v", sel.Type, types.ComponentTypeChannelSelect)
	}
}

func TestMentionableSelect_Valid(t *testing.T) {
	b := components.MentionableSelect("mentionable_select")
	err := b.Validate()
	if err != nil {
		t.Errorf("Validate() error = %v", err)
	}
}

func TestMentionableSelect_Placeholder(t *testing.T) {
	b := components.MentionableSelect("id")
	b.Placeholder("Select someone...")
	sel := b.Build()
	if sel.Placeholder != "Select someone..." {
		t.Errorf("Placeholder = %q, want %q", sel.Placeholder, "Select someone...")
	}
}

func TestMentionableSelect_Build(t *testing.T) {
	b := components.MentionableSelect("mentionable_id")
	sel := b.Build()

	if sel == nil {
		t.Fatal("Build() returned nil")
	}
	if sel.Type != types.ComponentTypeMentionableSelect {
		t.Errorf("Type = %v, want %v", sel.Type, types.ComponentTypeMentionableSelect)
	}
}

func TestSelectOption_Construction(t *testing.T) {
	tests := []struct {
		name string
		opt  components.SelectOption
	}{
		{"Basic option", components.SelectOption{
			Label: "Label",
			Value: "value",
		}},
		{"With description", components.SelectOption{
			Label:       "Label",
			Value:       "value",
			Description: "Description",
		}},
		{"With emoji", components.SelectOption{
			Label: "Label",
			Value: "value",
			Emoji: &components.ButtonEmoji{Name: "👍"},
		}},
		{"Default option", components.SelectOption{
			Label:   "Label",
			Value:   "value",
			Default: true,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.opt.Label == "" {
				t.Error("Label should not be empty")
			}
			if tt.opt.Value == "" {
				t.Error("Value should not be empty")
			}
		})
	}
}

func TestStringSelect_MaxOptionsBoundary(t *testing.T) {
	tests := []struct {
		name    string
		count   int
		wantErr bool
	}{
		{"Max options", components.MaxSelectOptionsPerMenu, false},
		{"Over max", components.MaxSelectOptionsPerMenu + 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := components.StringSelect("id")
			for i := 0; i < tt.count; i++ {
				b.Option("Option", "value")
			}
			err := b.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringSelect_PlaceholderMaxBoundary(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{"Max length", components.MaxSelectPlaceholderLength, false},
		{"Over max", components.MaxSelectPlaceholderLength + 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := components.StringSelect("id")
			b.Option("O", "v")
			b.Placeholder(string(make([]byte, tt.length)))
			err := b.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringSelect_OptionLabelMaxBoundary(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{"Max length", components.MaxSelectOptionLabelLength, false},
		{"Over max", components.MaxSelectOptionLabelLength + 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := components.StringSelect("id")
			b.Option(string(make([]byte, tt.length)), "value")
			err := b.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringSelect_OptionValueMaxBoundary(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{"Max length", components.MaxSelectOptionValueLength, false},
		{"Over max", components.MaxSelectOptionValueLength + 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := components.StringSelect("id")
			b.Option("Label", string(make([]byte, tt.length)))
			err := b.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringSelect_OptionDescriptionMaxBoundary(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{"Max length", components.MaxSelectOptionDescriptionLength, false},
		{"Over max", components.MaxSelectOptionDescriptionLength + 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := components.StringSelect("id")
			opt := components.SelectOption{
				Label:       "Label",
				Value:       "value",
				Description: string(make([]byte, tt.length)),
			}
			b.Options(opt)
			err := b.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

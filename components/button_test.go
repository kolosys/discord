package components_test

import (
	"testing"

	"github.com/kolosys/discord/components"
	"github.com/kolosys/discord/types"
)

func TestButton_Valid(t *testing.T) {
	tests := []struct {
		name  string
		setup func(*components.ButtonBuilder)
	}{
		{"Primary button", func(b *components.ButtonBuilder) { b.Label("Click me").Primary() }},
		{"Secondary button", func(b *components.ButtonBuilder) { b.Label("Click me").Secondary() }},
		{"Success button", func(b *components.ButtonBuilder) { b.Label("Click me").Success() }},
		{"Danger button", func(b *components.ButtonBuilder) { b.Label("Click me").Danger() }},
		{"Button with emoji", func(b *components.ButtonBuilder) { b.Label("React").Emoji("👍") }},
		{"Button with custom emoji", func(b *components.ButtonBuilder) { b.Label("React").CustomEmoji("123456", "name", false) }},
		{"Disabled button", func(b *components.ButtonBuilder) { b.Label("Disabled").Disabled(true) }},
		{"Button with label only", func(b *components.ButtonBuilder) { b.Label("Label") }},
		{"Button with emoji only", func(b *components.ButtonBuilder) { b.Emoji("✅") }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := components.Button("test_id")
			tt.setup(b)
			err := b.Validate()
			if err != nil {
				t.Errorf("Validate() error = %v", err)
			}
		})
	}
}

func TestButton_LinkButton(t *testing.T) {
	tests := []struct {
		name  string
		url   string
		setup func(*components.ButtonBuilder)
		valid bool
	}{
		{"Link button valid", "https://example.com", func(b *components.ButtonBuilder) { b.Label("Link") }, true},
		{"Link button with emoji", "https://example.com", func(b *components.ButtonBuilder) { b.Label("Link").Emoji("🔗") }, true},
		{"Link button without label", "https://example.com", func(b *components.ButtonBuilder) {}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := components.LinkButton(tt.url)
			tt.setup(b)
			err := b.Validate()
			if tt.valid && err != nil {
				t.Errorf("Validate() error = %v, want nil", err)
			}
			if !tt.valid && err == nil {
				t.Error("Validate() should error for invalid link button")
			}
		})
	}
}

func TestButton_LinkWithCustomID(t *testing.T) {
	// Test that a link button built normally doesn't have this error
	b := components.LinkButton("https://example.com").Label("Link")
	err := b.Validate()
	if err != nil {
		t.Errorf("Valid link button should not error: %v", err)
	}

	// Note: There is no public API to set CustomID on a link button.
	// The ErrLinkButtonWithCustomID error would only occur if:
	// 1. Internal code logic creates a link button with a custom_id
	// 2. Or external code directly constructs a ButtonComponent with both URL and CustomID
	// Both of these are outside normal usage patterns and would be caught by validation
}

func TestButton_MissingCustomID(t *testing.T) {
	b := components.Button("").Label("Missing ID")
	err := b.Validate()
	if err != components.ErrNonLinkButtonMissingCustomID {
		t.Errorf("Validate() = %v, want ErrNonLinkButtonMissingCustomID", err)
	}
}

func TestButton_LabelTooLong(t *testing.T) {
	b := components.Button("id")
	b.Label(string(make([]byte, components.MaxButtonLabelLength+1)))
	err := b.Validate()
	if err != components.ErrButtonLabelTooLong {
		t.Errorf("Validate() = %v, want ErrButtonLabelTooLong", err)
	}
}

func TestButton_CustomIDTooLong(t *testing.T) {
	b := components.Button(string(make([]byte, 101)))
	b.Label("Label")
	err := b.Validate()
	if err != components.ErrCustomIDTooLong {
		t.Errorf("Validate() = %v, want ErrCustomIDTooLong", err)
	}
}

func TestButton_MissingLabelAndEmoji(t *testing.T) {
	b := components.Button("id")
	err := b.Validate()
	if err != components.ErrButtonMissingLabelOrEmoji {
		t.Errorf("Validate() = %v, want ErrButtonMissingLabelOrEmoji", err)
	}
}

func TestButton_Disabled(t *testing.T) {
	tests := []bool{true, false}
	for _, disabled := range tests {
		b := components.Button("id").Label("Test").Disabled(disabled)
		btn := b.Build()
		if btn.Disabled != disabled {
			t.Errorf("Disabled = %v, want %v", btn.Disabled, disabled)
		}
	}
}

func TestButton_Emoji(t *testing.T) {
	b := components.Button("id").Label("React").Emoji("👍")
	btn := b.Build()
	if btn.Emoji == nil {
		t.Fatal("Emoji is nil")
	}
	if btn.Emoji.Name != "👍" {
		t.Errorf("Emoji.Name = %q, want %q", btn.Emoji.Name, "👍")
	}
	if btn.Emoji.ID != "" {
		t.Errorf("Emoji.ID = %q, want empty", btn.Emoji.ID)
	}
}

func TestButton_CustomEmoji(t *testing.T) {
	b := components.Button("id").Label("React").CustomEmoji("123456", "custom", true)
	btn := b.Build()
	if btn.Emoji == nil {
		t.Fatal("Emoji is nil")
	}
	if btn.Emoji.ID != "123456" {
		t.Errorf("Emoji.ID = %q, want %q", btn.Emoji.ID, "123456")
	}
	if btn.Emoji.Name != "custom" {
		t.Errorf("Emoji.Name = %q, want %q", btn.Emoji.Name, "custom")
	}
	if btn.Emoji.Animated != true {
		t.Errorf("Emoji.Animated = %v, want true", btn.Emoji.Animated)
	}
}

func TestButton_Chaining(t *testing.T) {
	b := components.Button("id")
	result := b.Label("Test").Primary().Disabled(true)
	if result != b {
		t.Error("Builder methods should return self for chaining")
	}
}

func TestButton_Style(t *testing.T) {
	tests := []struct {
		name  string
		style types.ButtonStyle
		fn    func(*components.ButtonBuilder) *components.ButtonBuilder
	}{
		{"Primary", types.ButtonStylePrimary, func(b *components.ButtonBuilder) *components.ButtonBuilder { return b.Primary() }},
		{"Secondary", types.ButtonStyleSecondary, func(b *components.ButtonBuilder) *components.ButtonBuilder { return b.Secondary() }},
		{"Success", types.ButtonStyleSuccess, func(b *components.ButtonBuilder) *components.ButtonBuilder { return b.Success() }},
		{"Danger", types.ButtonStyleDanger, func(b *components.ButtonBuilder) *components.ButtonBuilder { return b.Danger() }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := components.Button("id")
			tt.fn(b)
			btn := b.Build()
			if btn.Style != tt.style {
				t.Errorf("Style = %v, want %v", btn.Style, tt.style)
			}
		})
	}
}

func TestButton_Build(t *testing.T) {
	b := components.Button("test_id").Label("Click").Primary()
	btn := b.Build()

	if btn == nil {
		t.Fatal("Build() returned nil")
	}
	if btn.CustomID != "test_id" {
		t.Errorf("CustomID = %q, want %q", btn.CustomID, "test_id")
	}
	if btn.Label != "Click" {
		t.Errorf("Label = %q, want %q", btn.Label, "Click")
	}
	if btn.Style != types.ButtonStylePrimary {
		t.Errorf("Style = %v, want %v", btn.Style, types.ButtonStylePrimary)
	}
	if btn.Type != types.ComponentTypeButton {
		t.Errorf("Type = %v, want %v", btn.Type, types.ComponentTypeButton)
	}
}

func TestButton_BuildSafe(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(*components.ButtonBuilder)
		wantErr bool
	}{
		{"Valid button", func(b *components.ButtonBuilder) { b.Label("Valid") }, false},
		{"Missing label", func(b *components.ButtonBuilder) {}, true},
		{"Label too long", func(b *components.ButtonBuilder) { b.Label(string(make([]byte, 81))) }, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := components.Button("id")
			tt.setup(b)
			btn, err := b.BuildSafe()
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildSafe() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && btn == nil {
				t.Error("BuildSafe() returned nil for valid button")
			}
		})
	}
}

func TestButton_ConvenienceFunctions(t *testing.T) {
	tests := []struct {
		name     string
		fn       func() *components.ButtonComponent
		expected types.ButtonStyle
	}{
		{"PrimaryButton", func() *components.ButtonComponent { return components.PrimaryButton("id", "Label") }, types.ButtonStylePrimary},
		{"SecondaryButton", func() *components.ButtonComponent { return components.SecondaryButton("id", "Label") }, types.ButtonStyleSecondary},
		{"SuccessButton", func() *components.ButtonComponent { return components.SuccessButton("id", "Label") }, types.ButtonStyleSuccess},
		{"DangerButton", func() *components.ButtonComponent { return components.DangerButton("id", "Label") }, types.ButtonStyleDanger},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			btn := tt.fn()
			if btn.Style != tt.expected {
				t.Errorf("Style = %v, want %v", btn.Style, tt.expected)
			}
			if btn.Label != "Label" {
				t.Errorf("Label = %q, want %q", btn.Label, "Label")
			}
		})
	}
}

func TestButton_Link(t *testing.T) {
	btn := components.Link("https://example.com", "Visit")
	if btn.Style != types.ButtonStyleLink {
		t.Errorf("Style = %v, want %v", btn.Style, types.ButtonStyleLink)
	}
	if btn.URL != "https://example.com" {
		t.Errorf("URL = %q, want %q", btn.URL, "https://example.com")
	}
	if btn.Label != "Visit" {
		t.Errorf("Label = %q, want %q", btn.Label, "Visit")
	}
}

func TestButton_PremiumButton(t *testing.T) {
	b := components.PremiumButton("sku_123")
	btn := b.Build()
	if btn.Style != types.ButtonStylePremium {
		t.Errorf("Style = %v, want %v", btn.Style, types.ButtonStylePremium)
	}
	if btn.SkuID != "sku_123" {
		t.Errorf("SkuID = %q, want %q", btn.SkuID, "sku_123")
	}
}

func TestButton_MaxLabelBoundary(t *testing.T) {
	tests := []struct {
		name  string
		label string
		valid bool
	}{
		{"Max length", string(make([]byte, components.MaxButtonLabelLength)), true},
		{"Over max", string(make([]byte, components.MaxButtonLabelLength+1)), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := components.Button("id").Label(tt.label)
			err := b.Validate()
			if tt.valid && err != nil {
				t.Errorf("Validate() = %v, want nil", err)
			}
			if !tt.valid && err == nil {
				t.Error("Validate() should error for label too long")
			}
		})
	}
}

func TestButton_CustomIDMaxBoundary(t *testing.T) {
	tests := []struct {
		name  string
		id    string
		valid bool
	}{
		{"Max length", string(make([]byte, 100)), true},
		{"Over max", string(make([]byte, 101)), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := components.Button(tt.id).Label("Label")
			err := b.Validate()
			if tt.valid && err != nil {
				t.Errorf("Validate() = %v, want nil", err)
			}
			if !tt.valid && err == nil {
				t.Error("Validate() should error for custom_id too long")
			}
		})
	}
}

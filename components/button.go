package components

import "github.com/kolosys/discord/types"

// ButtonComponent represents a clickable button.
type ButtonComponent struct {
	Type     types.ComponentType `json:"type"`
	Style    types.ButtonStyle   `json:"style"`
	Label    string              `json:"label,omitempty"`
	Emoji    *ButtonEmoji        `json:"emoji,omitempty"`
	CustomID string              `json:"custom_id,omitempty"`
	URL      string              `json:"url,omitempty"`
	Disabled bool                `json:"disabled,omitempty"`
	SkuID    string              `json:"sku_id,omitempty"`
}

// ButtonEmoji represents an emoji on a button.
type ButtonEmoji struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Animated bool   `json:"animated,omitempty"`
}

// ButtonBuilder builds buttons with a fluent API.
type ButtonBuilder struct {
	button ButtonComponent
}

// Button creates a new button builder.
func Button(customID string) *ButtonBuilder {
	return &ButtonBuilder{
		button: ButtonComponent{
			Type:     types.ComponentTypeButton,
			CustomID: customID,
			Style:    types.ButtonStylePrimary,
		},
	}
}

// LinkButton creates a link button.
func LinkButton(url string) *ButtonBuilder {
	return &ButtonBuilder{
		button: ButtonComponent{
			Type:  types.ComponentTypeButton,
			URL:   url,
			Style: types.ButtonStyleLink,
		},
	}
}

// PremiumButton creates a premium/SKU button.
func PremiumButton(skuID string) *ButtonBuilder {
	return &ButtonBuilder{
		button: ButtonComponent{
			Type:  types.ComponentTypeButton,
			SkuID: skuID,
			Style: types.ButtonStylePremium,
		},
	}
}

// Label sets the button label.
func (b *ButtonBuilder) Label(label string) *ButtonBuilder {
	b.button.Label = label
	return b
}

// Style sets the button style.
func (b *ButtonBuilder) Style(style types.ButtonStyle) *ButtonBuilder {
	b.button.Style = style
	return b
}

// Primary sets the button to primary style.
func (b *ButtonBuilder) Primary() *ButtonBuilder {
	return b.Style(types.ButtonStylePrimary)
}

// Secondary sets the button to secondary style.
func (b *ButtonBuilder) Secondary() *ButtonBuilder {
	return b.Style(types.ButtonStyleSecondary)
}

// Success sets the button to success style.
func (b *ButtonBuilder) Success() *ButtonBuilder {
	return b.Style(types.ButtonStyleSuccess)
}

// Danger sets the button to danger style.
func (b *ButtonBuilder) Danger() *ButtonBuilder {
	return b.Style(types.ButtonStyleDanger)
}

// Emoji sets the button emoji by name (for standard emoji).
func (b *ButtonBuilder) Emoji(name string) *ButtonBuilder {
	b.button.Emoji = &ButtonEmoji{Name: name}
	return b
}

// CustomEmoji sets the button emoji by ID (for custom emoji).
func (b *ButtonBuilder) CustomEmoji(id, name string, animated bool) *ButtonBuilder {
	b.button.Emoji = &ButtonEmoji{
		ID:       id,
		Name:     name,
		Animated: animated,
	}
	return b
}

// Disabled sets whether the button is disabled.
func (b *ButtonBuilder) Disabled(disabled bool) *ButtonBuilder {
	b.button.Disabled = disabled
	return b
}

// Validate checks if the button configuration is valid according to Discord limits.
func (b *ButtonBuilder) Validate() error {
	btn := &b.button

	// Link buttons shouldn't have custom_id, non-link buttons require it
	if btn.Style == types.ButtonStyleLink {
		if btn.CustomID != "" {
			return ErrLinkButtonWithCustomID
		}
	} else if btn.Style != types.ButtonStylePremium {
		if btn.CustomID == "" {
			return ErrNonLinkButtonMissingCustomID
		}
		if err := ValidateCustomID(btn.CustomID); err != nil {
			return err
		}
	}

	// Must have label or emoji
	if btn.Label == "" && btn.Emoji == nil && btn.Style != types.ButtonStylePremium {
		return ErrButtonMissingLabelOrEmoji
	}

	// Label length
	if len(btn.Label) > MaxButtonLabelLength {
		return ErrButtonLabelTooLong
	}

	return nil
}

// Build returns the built button component.
func (b *ButtonBuilder) Build() *ButtonComponent {
	return &b.button
}

// BuildSafe validates and returns the built button component.
func (b *ButtonBuilder) BuildSafe() (*ButtonComponent, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}
	return &b.button, nil
}

// Convenience functions for quick button creation

// PrimaryButton creates a primary style button.
func PrimaryButton(customID, label string) *ButtonComponent {
	return Button(customID).Label(label).Primary().Build()
}

// SecondaryButton creates a secondary style button.
func SecondaryButton(customID, label string) *ButtonComponent {
	return Button(customID).Label(label).Secondary().Build()
}

// SuccessButton creates a success style button.
func SuccessButton(customID, label string) *ButtonComponent {
	return Button(customID).Label(label).Success().Build()
}

// DangerButton creates a danger style button.
func DangerButton(customID, label string) *ButtonComponent {
	return Button(customID).Label(label).Danger().Build()
}

// Link creates a link button.
func Link(url, label string) *ButtonComponent {
	return LinkButton(url).Label(label).Build()
}

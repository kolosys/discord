package components

import "github.com/kolosys/discord/types"

// SelectOption represents an option in a string select menu.
type SelectOption struct {
	Label       string       `json:"label"`
	Value       string       `json:"value"`
	Description string       `json:"description,omitempty"`
	Emoji       *ButtonEmoji `json:"emoji,omitempty"`
	Default     bool         `json:"default,omitempty"`
}

// StringSelectComponent represents a dropdown with string options.
type StringSelectComponent struct {
	Type        types.ComponentType `json:"type"`
	CustomID    string              `json:"custom_id"`
	Options     []SelectOption      `json:"options"`
	Placeholder string              `json:"placeholder,omitempty"`
	MinValues   *int                `json:"min_values,omitempty"`
	MaxValues   *int                `json:"max_values,omitempty"`
	Disabled    bool                `json:"disabled,omitempty"`
}

// StringSelectBuilder builds string select menus.
type StringSelectBuilder struct {
	sel StringSelectComponent
}

// StringSelect creates a new string select builder.
func StringSelect(customID string) *StringSelectBuilder {
	return &StringSelectBuilder{
		sel: StringSelectComponent{
			Type:     types.ComponentTypeStringSelect,
			CustomID: customID,
		},
	}
}

// Placeholder sets the placeholder text.
func (b *StringSelectBuilder) Placeholder(text string) *StringSelectBuilder {
	b.sel.Placeholder = text
	return b
}

// MinValues sets the minimum number of selections.
func (b *StringSelectBuilder) MinValues(min int) *StringSelectBuilder {
	b.sel.MinValues = &min
	return b
}

// MaxValues sets the maximum number of selections.
func (b *StringSelectBuilder) MaxValues(max int) *StringSelectBuilder {
	b.sel.MaxValues = &max
	return b
}

// Disabled sets whether the select is disabled.
func (b *StringSelectBuilder) Disabled(disabled bool) *StringSelectBuilder {
	b.sel.Disabled = disabled
	return b
}

// Options adds options to the select menu.
func (b *StringSelectBuilder) Options(opts ...SelectOption) *StringSelectBuilder {
	b.sel.Options = append(b.sel.Options, opts...)
	return b
}

// Option adds a single option.
func (b *StringSelectBuilder) Option(label, value string) *StringSelectBuilder {
	b.sel.Options = append(b.sel.Options, SelectOption{Label: label, Value: value})
	return b
}

// OptionWithDescription adds an option with description.
func (b *StringSelectBuilder) OptionWithDescription(label, value, description string) *StringSelectBuilder {
	b.sel.Options = append(b.sel.Options, SelectOption{
		Label:       label,
		Value:       value,
		Description: description,
	})
	return b
}

// Validate checks if the select menu configuration is valid according to Discord limits.
func (b *StringSelectBuilder) Validate() error {
	sel := &b.sel

	if err := ValidateCustomID(sel.CustomID); err != nil {
		return err
	}

	if err := ValidateSelectOptionCount(len(sel.Options)); err != nil {
		return err
	}

	if len(sel.Placeholder) > MaxSelectPlaceholderLength {
		return ErrSelectPlaceholderTooLong
	}

	for _, opt := range sel.Options {
		if len(opt.Label) > MaxSelectOptionLabelLength {
			return ErrSelectOptionLabelTooLong
		}
		if len(opt.Value) > MaxSelectOptionValueLength {
			return ErrSelectOptionValueTooLong
		}
		if len(opt.Description) > MaxSelectOptionDescriptionLength {
			return ErrSelectOptionDescTooLong
		}
	}

	return nil
}

// Build returns the built select component.
func (b *StringSelectBuilder) Build() *StringSelectComponent {
	return &b.sel
}

// BuildSafe validates and returns the built select component.
func (b *StringSelectBuilder) BuildSafe() (*StringSelectComponent, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}
	return &b.sel, nil
}

// UserSelectComponent represents a select menu for users.
type UserSelectComponent struct {
	Type         types.ComponentType `json:"type"`
	CustomID     string              `json:"custom_id"`
	Placeholder  string              `json:"placeholder,omitempty"`
	MinValues    *int                `json:"min_values,omitempty"`
	MaxValues    *int                `json:"max_values,omitempty"`
	Disabled     bool                `json:"disabled,omitempty"`
	DefaultUsers []string            `json:"default_values,omitempty"`
}

// UserSelectBuilder builds user select menus.
type UserSelectBuilder struct {
	sel UserSelectComponent
}

// UserSelect creates a new user select builder.
func UserSelect(customID string) *UserSelectBuilder {
	return &UserSelectBuilder{
		sel: UserSelectComponent{
			Type:     types.ComponentTypeUserSelect,
			CustomID: customID,
		},
	}
}

// Placeholder sets the placeholder text.
func (b *UserSelectBuilder) Placeholder(text string) *UserSelectBuilder {
	b.sel.Placeholder = text
	return b
}

// MinValues sets the minimum number of selections.
func (b *UserSelectBuilder) MinValues(min int) *UserSelectBuilder {
	b.sel.MinValues = &min
	return b
}

// MaxValues sets the maximum number of selections.
func (b *UserSelectBuilder) MaxValues(max int) *UserSelectBuilder {
	b.sel.MaxValues = &max
	return b
}

// Disabled sets whether the select is disabled.
func (b *UserSelectBuilder) Disabled(disabled bool) *UserSelectBuilder {
	b.sel.Disabled = disabled
	return b
}

// Validate checks if the select menu configuration is valid according to Discord limits.
func (b *UserSelectBuilder) Validate() error {
	if err := ValidateCustomID(b.sel.CustomID); err != nil {
		return err
	}
	if len(b.sel.Placeholder) > MaxSelectPlaceholderLength {
		return ErrSelectPlaceholderTooLong
	}
	return nil
}

// Build returns the built select component.
func (b *UserSelectBuilder) Build() *UserSelectComponent {
	return &b.sel
}

// BuildSafe validates and returns the built select component.
func (b *UserSelectBuilder) BuildSafe() (*UserSelectComponent, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}
	return &b.sel, nil
}

// RoleSelectComponent represents a select menu for roles.
type RoleSelectComponent struct {
	Type        types.ComponentType `json:"type"`
	CustomID    string              `json:"custom_id"`
	Placeholder string              `json:"placeholder,omitempty"`
	MinValues   *int                `json:"min_values,omitempty"`
	MaxValues   *int                `json:"max_values,omitempty"`
	Disabled    bool                `json:"disabled,omitempty"`
}

// RoleSelectBuilder builds role select menus.
type RoleSelectBuilder struct {
	sel RoleSelectComponent
}

// RoleSelect creates a new role select builder.
func RoleSelect(customID string) *RoleSelectBuilder {
	return &RoleSelectBuilder{
		sel: RoleSelectComponent{
			Type:     types.ComponentTypeRoleSelect,
			CustomID: customID,
		},
	}
}

// Placeholder sets the placeholder text.
func (b *RoleSelectBuilder) Placeholder(text string) *RoleSelectBuilder {
	b.sel.Placeholder = text
	return b
}

// MinValues sets the minimum number of selections.
func (b *RoleSelectBuilder) MinValues(min int) *RoleSelectBuilder {
	b.sel.MinValues = &min
	return b
}

// MaxValues sets the maximum number of selections.
func (b *RoleSelectBuilder) MaxValues(max int) *RoleSelectBuilder {
	b.sel.MaxValues = &max
	return b
}

// Disabled sets whether the select is disabled.
func (b *RoleSelectBuilder) Disabled(disabled bool) *RoleSelectBuilder {
	b.sel.Disabled = disabled
	return b
}

// Validate checks if the select menu configuration is valid according to Discord limits.
func (b *RoleSelectBuilder) Validate() error {
	if err := ValidateCustomID(b.sel.CustomID); err != nil {
		return err
	}
	if len(b.sel.Placeholder) > MaxSelectPlaceholderLength {
		return ErrSelectPlaceholderTooLong
	}
	return nil
}

// Build returns the built select component.
func (b *RoleSelectBuilder) Build() *RoleSelectComponent {
	return &b.sel
}

// BuildSafe validates and returns the built select component.
func (b *RoleSelectBuilder) BuildSafe() (*RoleSelectComponent, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}
	return &b.sel, nil
}

// ChannelSelectComponent represents a select menu for channels.
type ChannelSelectComponent struct {
	Type         types.ComponentType `json:"type"`
	CustomID     string              `json:"custom_id"`
	Placeholder  string              `json:"placeholder,omitempty"`
	MinValues    *int                `json:"min_values,omitempty"`
	MaxValues    *int                `json:"max_values,omitempty"`
	Disabled     bool                `json:"disabled,omitempty"`
	ChannelTypes []int32             `json:"channel_types,omitempty"`
}

// ChannelSelectBuilder builds channel select menus.
type ChannelSelectBuilder struct {
	sel ChannelSelectComponent
}

// ChannelSelect creates a new channel select builder.
func ChannelSelect(customID string) *ChannelSelectBuilder {
	return &ChannelSelectBuilder{
		sel: ChannelSelectComponent{
			Type:     types.ComponentTypeChannelSelect,
			CustomID: customID,
		},
	}
}

// Placeholder sets the placeholder text.
func (b *ChannelSelectBuilder) Placeholder(text string) *ChannelSelectBuilder {
	b.sel.Placeholder = text
	return b
}

// MinValues sets the minimum number of selections.
func (b *ChannelSelectBuilder) MinValues(min int) *ChannelSelectBuilder {
	b.sel.MinValues = &min
	return b
}

// MaxValues sets the maximum number of selections.
func (b *ChannelSelectBuilder) MaxValues(max int) *ChannelSelectBuilder {
	b.sel.MaxValues = &max
	return b
}

// Disabled sets whether the select is disabled.
func (b *ChannelSelectBuilder) Disabled(disabled bool) *ChannelSelectBuilder {
	b.sel.Disabled = disabled
	return b
}

// ChannelTypes filters the channels by type.
func (b *ChannelSelectBuilder) ChannelTypes(types ...int32) *ChannelSelectBuilder {
	b.sel.ChannelTypes = types
	return b
}

// Validate checks if the select menu configuration is valid according to Discord limits.
func (b *ChannelSelectBuilder) Validate() error {
	if err := ValidateCustomID(b.sel.CustomID); err != nil {
		return err
	}
	if len(b.sel.Placeholder) > MaxSelectPlaceholderLength {
		return ErrSelectPlaceholderTooLong
	}
	return nil
}

// Build returns the built select component.
func (b *ChannelSelectBuilder) Build() *ChannelSelectComponent {
	return &b.sel
}

// BuildSafe validates and returns the built select component.
func (b *ChannelSelectBuilder) BuildSafe() (*ChannelSelectComponent, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}
	return &b.sel, nil
}

// MentionableSelectComponent represents a select menu for mentionable entities.
type MentionableSelectComponent struct {
	Type        types.ComponentType `json:"type"`
	CustomID    string              `json:"custom_id"`
	Placeholder string              `json:"placeholder,omitempty"`
	MinValues   *int                `json:"min_values,omitempty"`
	MaxValues   *int                `json:"max_values,omitempty"`
	Disabled    bool                `json:"disabled,omitempty"`
}

// MentionableSelectBuilder builds mentionable select menus.
type MentionableSelectBuilder struct {
	sel MentionableSelectComponent
}

// MentionableSelect creates a new mentionable select builder.
func MentionableSelect(customID string) *MentionableSelectBuilder {
	return &MentionableSelectBuilder{
		sel: MentionableSelectComponent{
			Type:     types.ComponentTypeMentionableSelect,
			CustomID: customID,
		},
	}
}

// Placeholder sets the placeholder text.
func (b *MentionableSelectBuilder) Placeholder(text string) *MentionableSelectBuilder {
	b.sel.Placeholder = text
	return b
}

// MinValues sets the minimum number of selections.
func (b *MentionableSelectBuilder) MinValues(min int) *MentionableSelectBuilder {
	b.sel.MinValues = &min
	return b
}

// MaxValues sets the maximum number of selections.
func (b *MentionableSelectBuilder) MaxValues(max int) *MentionableSelectBuilder {
	b.sel.MaxValues = &max
	return b
}

// Disabled sets whether the select is disabled.
func (b *MentionableSelectBuilder) Disabled(disabled bool) *MentionableSelectBuilder {
	b.sel.Disabled = disabled
	return b
}

// Validate checks if the select menu configuration is valid according to Discord limits.
func (b *MentionableSelectBuilder) Validate() error {
	if err := ValidateCustomID(b.sel.CustomID); err != nil {
		return err
	}
	if len(b.sel.Placeholder) > MaxSelectPlaceholderLength {
		return ErrSelectPlaceholderTooLong
	}
	return nil
}

// Build returns the built select component.
func (b *MentionableSelectBuilder) Build() *MentionableSelectComponent {
	return &b.sel
}

// BuildSafe validates and returns the built select component.
func (b *MentionableSelectBuilder) BuildSafe() (*MentionableSelectComponent, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}
	return &b.sel, nil
}

// Quick creation functions

// NewSelectOption creates a new select option.
func NewSelectOption(label, value string) SelectOption {
	return SelectOption{Label: label, Value: value}
}

// NewSelectOptionWithEmoji creates a select option with an emoji.
func NewSelectOptionWithEmoji(label, value, emoji string) SelectOption {
	return SelectOption{
		Label: label,
		Value: value,
		Emoji: &ButtonEmoji{Name: emoji},
	}
}

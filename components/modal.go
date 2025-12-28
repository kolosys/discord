package components

import "github.com/kolosys/discord/types"

// TextInputComponent represents a text input in a modal.
type TextInputComponent struct {
	Type        types.ComponentType  `json:"type"`
	CustomID    string               `json:"custom_id"`
	Style       types.TextInputStyle `json:"style"`
	Label       string               `json:"label"`
	MinLength   *int                 `json:"min_length,omitempty"`
	MaxLength   *int                 `json:"max_length,omitempty"`
	Required    *bool                `json:"required,omitempty"`
	Value       string               `json:"value,omitempty"`
	Placeholder string               `json:"placeholder,omitempty"`
}

// TextInputBuilder builds text inputs with a fluent API.
type TextInputBuilder struct {
	input TextInputComponent
}

// TextInput creates a new text input builder.
func TextInput(customID, label string) *TextInputBuilder {
	return &TextInputBuilder{
		input: TextInputComponent{
			Type:     types.ComponentTypeTextInput,
			CustomID: customID,
			Label:    label,
			Style:    types.TextInputStyleShort,
		},
	}
}

// Short sets the input to single-line style.
func (b *TextInputBuilder) Short() *TextInputBuilder {
	b.input.Style = types.TextInputStyleShort
	return b
}

// Paragraph sets the input to multi-line style.
func (b *TextInputBuilder) Paragraph() *TextInputBuilder {
	b.input.Style = types.TextInputStyleParagraph
	return b
}

// Style sets the input style.
func (b *TextInputBuilder) Style(style types.TextInputStyle) *TextInputBuilder {
	b.input.Style = style
	return b
}

// MinLength sets the minimum length.
func (b *TextInputBuilder) MinLength(min int) *TextInputBuilder {
	b.input.MinLength = &min
	return b
}

// MaxLength sets the maximum length.
func (b *TextInputBuilder) MaxLength(max int) *TextInputBuilder {
	b.input.MaxLength = &max
	return b
}

// Required sets whether the input is required.
func (b *TextInputBuilder) Required(required bool) *TextInputBuilder {
	b.input.Required = &required
	return b
}

// Value sets the default value.
func (b *TextInputBuilder) Value(value string) *TextInputBuilder {
	b.input.Value = value
	return b
}

// Placeholder sets the placeholder text.
func (b *TextInputBuilder) Placeholder(text string) *TextInputBuilder {
	b.input.Placeholder = text
	return b
}

// Validate checks if the text input configuration is valid according to Discord limits.
func (b *TextInputBuilder) Validate() error {
	if err := ValidateCustomID(b.input.CustomID); err != nil {
		return err
	}

	if len(b.input.Label) > MaxTextInputLabelLength {
		return ErrTextInputLabelTooLong
	}

	if len(b.input.Placeholder) > MaxTextInputPlaceholderLength {
		return ErrTextInputPlaceholderTooLong
	}

	return nil
}

// Build returns the built text input wrapped in an action row.
func (b *TextInputBuilder) Build() *ActionRowComponent {
	return ActionRow(&b.input)
}

// BuildSafe validates and returns the built text input wrapped in an action row.
func (b *TextInputBuilder) BuildSafe() (*ActionRowComponent, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}
	return ActionRow(&b.input), nil
}

// BuildRaw returns just the text input component.
func (b *TextInputBuilder) BuildRaw() *TextInputComponent {
	return &b.input
}

// ModalBuilder builds modals with a fluent API.
type ModalBuilder struct {
	customID   string
	title      string
	components []any
}

// Modal creates a new modal builder.
func Modal(customID, title string) *ModalBuilder {
	return &ModalBuilder{
		customID: customID,
		title:    title,
	}
}

// TextInput adds a text input to the modal.
func (m *ModalBuilder) TextInput(customID, label string) *ModalTextInputBuilder {
	return &ModalTextInputBuilder{
		modal: m,
		input: TextInputComponent{
			Type:     types.ComponentTypeTextInput,
			CustomID: customID,
			Label:    label,
			Style:    types.TextInputStyleShort,
		},
	}
}

// ShortInput adds a short text input.
func (m *ModalBuilder) ShortInput(customID, label string) *ModalBuilder {
	m.components = append(m.components, ActionRow(&TextInputComponent{
		Type:     types.ComponentTypeTextInput,
		CustomID: customID,
		Label:    label,
		Style:    types.TextInputStyleShort,
	}))
	return m
}

// ParagraphInput adds a paragraph text input.
func (m *ModalBuilder) ParagraphInput(customID, label string) *ModalBuilder {
	m.components = append(m.components, ActionRow(&TextInputComponent{
		Type:     types.ComponentTypeTextInput,
		CustomID: customID,
		Label:    label,
		Style:    types.TextInputStyleParagraph,
	}))
	return m
}

// Component adds a raw component.
func (m *ModalBuilder) Component(comp any) *ModalBuilder {
	m.components = append(m.components, comp)
	return m
}

// Validate checks if the modal configuration is valid according to Discord limits.
func (m *ModalBuilder) Validate() error {
	if err := ValidateCustomID(m.customID); err != nil {
		return err
	}

	if len(m.title) > MaxModalTitleLength {
		return ErrModalTitleTooLong
	}

	if err := ValidateModalInputCount(len(m.components)); err != nil {
		return err
	}

	return nil
}

// Build returns the modal data for responding.
func (m *ModalBuilder) Build() *ModalData {
	return &ModalData{
		CustomID:   m.customID,
		Title:      m.title,
		Components: m.components,
	}
}

// BuildSafe validates and returns the modal data.
func (m *ModalBuilder) BuildSafe() (*ModalData, error) {
	if err := m.Validate(); err != nil {
		return nil, err
	}
	return m.Build(), nil
}

// ModalData represents the data for a modal response.
type ModalData struct {
	CustomID   string `json:"custom_id"`
	Title      string `json:"title"`
	Components []any  `json:"components"`
}

// ModalTextInputBuilder is a builder for text inputs within a modal.
type ModalTextInputBuilder struct {
	modal *ModalBuilder
	input TextInputComponent
}

// Short sets the input to single-line style.
func (b *ModalTextInputBuilder) Short() *ModalTextInputBuilder {
	b.input.Style = types.TextInputStyleShort
	return b
}

// Paragraph sets the input to multi-line style.
func (b *ModalTextInputBuilder) Paragraph() *ModalTextInputBuilder {
	b.input.Style = types.TextInputStyleParagraph
	return b
}

// MinLength sets the minimum length.
func (b *ModalTextInputBuilder) MinLength(min int) *ModalTextInputBuilder {
	b.input.MinLength = &min
	return b
}

// MaxLength sets the maximum length.
func (b *ModalTextInputBuilder) MaxLength(max int) *ModalTextInputBuilder {
	b.input.MaxLength = &max
	return b
}

// Required sets whether the input is required.
func (b *ModalTextInputBuilder) Required(required bool) *ModalTextInputBuilder {
	b.input.Required = &required
	return b
}

// Value sets the default value.
func (b *ModalTextInputBuilder) Value(value string) *ModalTextInputBuilder {
	b.input.Value = value
	return b
}

// Placeholder sets the placeholder text.
func (b *ModalTextInputBuilder) Placeholder(text string) *ModalTextInputBuilder {
	b.input.Placeholder = text
	return b
}

// Done finishes the text input and returns to the modal builder.
func (b *ModalTextInputBuilder) Done() *ModalBuilder {
	b.modal.components = append(b.modal.components, ActionRow(&b.input))
	return b.modal
}

// Convenience functions

// ShortTextInput creates a short text input.
func ShortTextInput(customID, label string) *TextInputBuilder {
	return TextInput(customID, label).Short()
}

// ParagraphTextInput creates a paragraph text input.
func ParagraphTextInput(customID, label string) *TextInputBuilder {
	return TextInput(customID, label).Paragraph()
}

package components

import "github.com/kolosys/discord/types"

// ActionRow creates a new action row with the given components.
func ActionRow(components ...Component) *ActionRowComponent {
	return &ActionRowComponent{
		Type:       types.ComponentTypeActionRow,
		Components: components,
	}
}

// ActionRowComponent is a container for other components.
type ActionRowComponent struct {
	Type       types.ComponentType `json:"type"`
	Components []Component         `json:"components"`
}

// Component is the interface for all message components.
type Component interface {
	component() // marker method
}

func (b *ButtonComponent) component()            {}
func (s *StringSelectComponent) component()      {}
func (s *UserSelectComponent) component()        {}
func (s *RoleSelectComponent) component()        {}
func (s *ChannelSelectComponent) component()     {}
func (s *MentionableSelectComponent) component() {}
func (t *TextInputComponent) component()         {}
func (a *ActionRowComponent) component()         {}

// BuildComponents takes individual components and wraps them in action rows.
// Buttons are grouped together (max 5 per row), select menus get their own rows.
func BuildComponents(components ...Component) []any {
	var result []any
	var buttonRow []Component

	flushButtons := func() {
		if len(buttonRow) > 0 {
			result = append(result, ActionRow(buttonRow...))
			buttonRow = nil
		}
	}

	for _, comp := range components {
		switch c := comp.(type) {
		case *ButtonComponent:
			buttonRow = append(buttonRow, c)
			if len(buttonRow) >= MaxButtonsPerRow {
				flushButtons()
			}
		case *ActionRowComponent:
			flushButtons()
			result = append(result, c)
		default:
			flushButtons()
			result = append(result, ActionRow(comp))
		}
	}

	flushButtons()
	return result
}

// BuildComponentsSafe takes individual components and wraps them in action rows,
// returning an error if Discord limits are exceeded.
func BuildComponentsSafe(components ...Component) ([]any, error) {
	result := BuildComponents(components...)

	if err := ValidateActionRowCount(len(result)); err != nil {
		return nil, err
	}

	totalComponents := 0
	for _, row := range result {
		if ar, ok := row.(*ActionRowComponent); ok {
			totalComponents += len(ar.Components)
		}
	}

	if totalComponents > MaxComponentsPerMessage {
		return nil, ErrTooManyComponents
	}

	return result, nil
}

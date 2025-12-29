package components

import "errors"

// Discord API limits for message components.
const (
	// MaxCustomIDLength is the maximum length of a custom_id field.
	MaxCustomIDLength = 100

	// MaxActionRows is the maximum number of action rows per message.
	MaxActionRows = 5

	// MaxButtonsPerRow is the maximum number of buttons per action row.
	MaxButtonsPerRow = 5

	// MaxSelectOptionsPerMenu is the maximum number of options in a string select menu.
	MaxSelectOptionsPerMenu = 25

	// MaxComponentsPerMessage is the maximum total components per message.
	MaxComponentsPerMessage = 25

	// MaxTextInputsPerModal is the maximum number of text inputs per modal.
	MaxTextInputsPerModal = 5

	// MaxButtonLabelLength is the maximum length of a button label.
	MaxButtonLabelLength = 80

	// MaxSelectPlaceholderLength is the maximum length of a select menu placeholder.
	MaxSelectPlaceholderLength = 150

	// MaxSelectOptionLabelLength is the maximum length of a select option label.
	MaxSelectOptionLabelLength = 100

	// MaxSelectOptionValueLength is the maximum length of a select option value.
	MaxSelectOptionValueLength = 100

	// MaxSelectOptionDescriptionLength is the maximum length of a select option description.
	MaxSelectOptionDescriptionLength = 100

	// MaxModalTitleLength is the maximum length of a modal title.
	MaxModalTitleLength = 45

	// MaxTextInputLabelLength is the maximum length of a text input label.
	MaxTextInputLabelLength = 45

	// MaxTextInputPlaceholderLength is the maximum length of a text input placeholder.
	MaxTextInputPlaceholderLength = 100

	// MaxTextInputValueLength is the maximum length of a text input value.
	MaxTextInputValueLength = 4000
)

// Validation errors.
var (
	ErrCustomIDTooLong              = errors.New("components: custom_id exceeds 100 characters")
	ErrTooManyActionRows            = errors.New("components: message cannot have more than 5 action rows")
	ErrTooManyButtonsPerRow         = errors.New("components: action row cannot have more than 5 buttons")
	ErrTooManySelectOptions         = errors.New("components: select menu cannot have more than 25 options")
	ErrTooManyComponents            = errors.New("components: message cannot have more than 25 components")
	ErrTooManyModalInputs           = errors.New("components: modal cannot have more than 5 text inputs")
	ErrButtonLabelTooLong           = errors.New("components: button label exceeds 80 characters")
	ErrSelectPlaceholderTooLong     = errors.New("components: select placeholder exceeds 150 characters")
	ErrSelectOptionLabelTooLong     = errors.New("components: select option label exceeds 100 characters")
	ErrSelectOptionValueTooLong     = errors.New("components: select option value exceeds 100 characters")
	ErrSelectOptionDescTooLong      = errors.New("components: select option description exceeds 100 characters")
	ErrModalTitleTooLong            = errors.New("components: modal title exceeds 45 characters")
	ErrTextInputLabelTooLong        = errors.New("components: text input label exceeds 45 characters")
	ErrTextInputPlaceholderTooLong  = errors.New("components: text input placeholder exceeds 100 characters")
	ErrNoSelectOptions              = errors.New("components: select menu must have at least 1 option")
	ErrButtonMissingLabelOrEmoji    = errors.New("components: button must have a label or emoji")
	ErrLinkButtonWithCustomID       = errors.New("components: link button cannot have a custom_id")
	ErrNonLinkButtonMissingCustomID = errors.New("components: non-link button must have a custom_id")
)

// ValidateCustomID checks if a custom_id is within the allowed length.
func ValidateCustomID(customID string) error {
	if len(customID) > MaxCustomIDLength {
		return ErrCustomIDTooLong
	}
	return nil
}

// ValidateActionRowCount checks if the number of action rows is within limits.
func ValidateActionRowCount(count int) error {
	if count > MaxActionRows {
		return ErrTooManyActionRows
	}
	return nil
}

// ValidateSelectOptionCount checks if the number of select options is within limits.
func ValidateSelectOptionCount(count int) error {
	if count > MaxSelectOptionsPerMenu {
		return ErrTooManySelectOptions
	}
	if count == 0 {
		return ErrNoSelectOptions
	}
	return nil
}

// ValidateModalInputCount checks if the number of modal inputs is within limits.
func ValidateModalInputCount(count int) error {
	if count > MaxTextInputsPerModal {
		return ErrTooManyModalInputs
	}
	return nil
}

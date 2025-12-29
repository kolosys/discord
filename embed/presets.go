package embed

// SuccessEmbed creates a green embed with the given title and description.
func SuccessEmbed(title, description string) *Builder {
	return New().Title(title).Description(description).Success()
}

// ErrorEmbed creates a red embed with the given title and description.
func ErrorEmbed(title, description string) *Builder {
	return New().Title(title).Description(description).Error()
}

// WarningEmbed creates a yellow embed with the given title and description.
func WarningEmbed(title, description string) *Builder {
	return New().Title(title).Description(description).Warning()
}

// InfoEmbed creates a blurple embed with the given title and description.
func InfoEmbed(title, description string) *Builder {
	return New().Title(title).Description(description).Info()
}

// SimpleEmbed creates a basic embed with just a description.
func SimpleEmbed(description string) *Builder {
	return New().Description(description)
}

// TitleEmbed creates an embed with just a title.
func TitleEmbed(title string) *Builder {
	return New().Title(title)
}

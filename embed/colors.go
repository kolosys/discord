package embed

import "github.com/kolosys/discord/types"

// Preset colors for common embed types
const (
	ColorSuccess = int(types.ColorGreen)   // 0x57F287 - Green
	ColorError   = int(types.ColorRed)     // 0xED4245 - Red
	ColorWarning = int(types.ColorYellow)  // 0xFEE75C - Yellow
	ColorInfo    = int(types.ColorBlurple) // 0x5865F2 - Discord Blurple
)

// Success sets the embed color to green (success).
func (b *Builder) Success() *Builder {
	return b.Color(ColorSuccess)
}

// Error sets the embed color to red (error).
func (b *Builder) Error() *Builder {
	return b.Color(ColorError)
}

// Warning sets the embed color to yellow (warning).
func (b *Builder) Warning() *Builder {
	return b.Color(ColorWarning)
}

// Info sets the embed color to blurple (info).
func (b *Builder) Info() *Builder {
	return b.Color(ColorInfo)
}

// Blurple sets the embed color to Discord's brand color.
func (b *Builder) Blurple() *Builder {
	return b.Color(int(types.ColorBlurple))
}

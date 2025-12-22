package types

import "strconv"

// Color represents a 24-bit RGB color value.
type Color int32

const (
	ColorDefault       Color = 0x000000
	ColorBlurple       Color = 0x5865F2
	ColorGreen         Color = 0x57F287
	ColorYellow        Color = 0xFEE75C
	ColorFuchsia       Color = 0xEB459E
	ColorRed           Color = 0xED4245
	ColorWhite         Color = 0xFFFFFF
	ColorBlack         Color = 0x000000
	ColorDiscordOld    Color = 0x7289DA
	ColorDiscordGreen  Color = 0x3BA55C
	ColorDiscordYellow Color = 0xFAA61A
	ColorDiscordRed    Color = 0xED4245
)

func (c Color) R() uint8 { return uint8((c >> 16) & 0xFF) }
func (c Color) G() uint8 { return uint8((c >> 8) & 0xFF) }
func (c Color) B() uint8 { return uint8(c & 0xFF) }

func NewColor(r, g, b uint8) Color {
	return Color((int32(r) << 16) | (int32(g) << 8) | int32(b))
}

func ColorFromHex(hex string) Color {
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}
	val, _ := strconv.ParseInt(hex, 16, 32)
	return Color(val)
}

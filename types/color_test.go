package types_test

import (
	"testing"

	"github.com/kolosys/discord/types"
)

func TestColor_Components(t *testing.T) {
	tests := []struct {
		name  string
		color types.Color
		r     uint8
		g     uint8
		b     uint8
	}{
		{"Black", 0x000000, 0x00, 0x00, 0x00},
		{"White", 0xFFFFFF, 0xFF, 0xFF, 0xFF},
		{"Red", 0xFF0000, 0xFF, 0x00, 0x00},
		{"Green", 0x00FF00, 0x00, 0xFF, 0x00},
		{"Blue", 0x0000FF, 0x00, 0x00, 0xFF},
		{"Discord Blurple", 0x5865F2, 0x58, 0x65, 0xF2},
		{"Discord Green", 0x57F287, 0x57, 0xF2, 0x87},
		{"Discord Yellow", 0xFEE75C, 0xFE, 0xE7, 0x5C},
		{"Discord Red", 0xED4245, 0xED, 0x42, 0x45},
		{"Mixed", 0xA1B2C3, 0xA1, 0xB2, 0xC3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			color := types.Color(tt.color)
			if r := color.R(); r != tt.r {
				t.Errorf("R() = %02x, want %02x", r, tt.r)
			}
			if g := color.G(); g != tt.g {
				t.Errorf("G() = %02x, want %02x", g, tt.g)
			}
			if b := color.B(); b != tt.b {
				t.Errorf("B() = %02x, want %02x", b, tt.b)
			}
		})
	}
}

func TestColor_NewColor(t *testing.T) {
	tests := []struct {
		name     string
		r        uint8
		g        uint8
		b        uint8
		expected types.Color
	}{
		{"Black", 0x00, 0x00, 0x00, 0x000000},
		{"White", 0xFF, 0xFF, 0xFF, 0xFFFFFF},
		{"Red", 0xFF, 0x00, 0x00, 0xFF0000},
		{"Green", 0x00, 0xFF, 0x00, 0x00FF00},
		{"Blue", 0x00, 0x00, 0xFF, 0x0000FF},
		{"Discord Blurple", 0x58, 0x65, 0xF2, 0x5865F2},
		{"Discord Green", 0x57, 0xF2, 0x87, 0x57F287},
		{"Discord Yellow", 0xFE, 0xE7, 0x5C, 0xFEE75C},
		{"Discord Red", 0xED, 0x42, 0x45, 0xED4245},
		{"Mixed", 0xA1, 0xB2, 0xC3, 0xA1B2C3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := types.NewColor(tt.r, tt.g, tt.b)
			if got != tt.expected {
				t.Errorf("NewColor(%02x, %02x, %02x) = %06x, want %06x", tt.r, tt.g, tt.b, got, tt.expected)
			}
		})
	}
}

func TestColor_FromHex(t *testing.T) {
	tests := []struct {
		name     string
		hex      string
		expected types.Color
	}{
		{"Lowercase", "5865f2", 0x5865F2},
		{"Uppercase", "5865F2", 0x5865F2},
		{"Mixed case", "5865f2", 0x5865F2},
		{"With hash", "#5865F2", 0x5865F2},
		{"With hash lowercase", "#5865f2", 0x5865F2},
		{"Black", "000000", 0x000000},
		{"Black with hash", "#000000", 0x000000},
		{"White", "FFFFFF", 0xFFFFFF},
		{"White with hash", "#FFFFFF", 0xFFFFFF},
		{"Red", "FF0000", 0xFF0000},
		{"Green", "00FF00", 0x00FF00},
		{"Blue", "0000FF", 0x0000FF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := types.ColorFromHex(tt.hex)
			if got != tt.expected {
				t.Errorf("ColorFromHex(%q) = %06x, want %06x", tt.hex, got, tt.expected)
			}
		})
	}
}

func TestColor_Constants(t *testing.T) {
	tests := []struct {
		name  string
		color types.Color
		hex   string
	}{
		{"ColorDefault", types.ColorDefault, "000000"},
		{"ColorBlurple", types.ColorBlurple, "5865F2"},
		{"ColorGreen", types.ColorGreen, "57F287"},
		{"ColorYellow", types.ColorYellow, "FEE75C"},
		{"ColorFuchsia", types.ColorFuchsia, "EB459E"},
		{"ColorRed", types.ColorRed, "ED4245"},
		{"ColorWhite", types.ColorWhite, "FFFFFF"},
		{"ColorBlack", types.ColorBlack, "000000"},
		{"ColorDiscordOld", types.ColorDiscordOld, "7289DA"},
		{"ColorDiscordGreen", types.ColorDiscordGreen, "3BA55C"},
		{"ColorDiscordYellow", types.ColorDiscordYellow, "FAA61A"},
		{"ColorDiscordRed", types.ColorDiscordRed, "ED4245"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verify constant equals expected value
			expected := types.ColorFromHex(tt.hex)
			if tt.color != expected {
				t.Errorf("%s = %06x, want %06x", tt.name, tt.color, expected)
			}
		})
	}
}

func TestColor_Roundtrip(t *testing.T) {
	tests := []struct {
		name string
		hex  string
	}{
		{"Blurple", "5865F2"},
		{"Green", "57F287"},
		{"Yellow", "FEE75C"},
		{"Red", "ED4245"},
		{"Black", "000000"},
		{"White", "FFFFFF"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse hex to color
			color := types.ColorFromHex(tt.hex)

			// Extract components
			r := color.R()
			g := color.G()
			b := color.B()

			// Create new color from components
			newColor := types.NewColor(r, g, b)

			// Should match original
			if newColor != color {
				t.Errorf("Roundtrip failed: NewColor(%02x, %02x, %02x) = %06x, want %06x", r, g, b, newColor, color)
			}
		})
	}
}

func TestColor_ComponentExtraction(t *testing.T) {
	// Test that extracting and recombining gives original
	original := types.Color(0x123456)

	r := original.R()
	g := original.G()
	b := original.B()

	reconstructed := types.NewColor(r, g, b)

	if original != reconstructed {
		t.Errorf("Component roundtrip failed: %06x != %06x", original, reconstructed)
	}
}

func TestColor_EdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		color types.Color
	}{
		{"All zeros", 0x000000},
		{"All ones", 0xFFFFFF},
		{"Alternating", 0xAAAAA},
		{"First byte only", 0xFF0000},
		{"Second byte only", 0x00FF00},
		{"Third byte only", 0x0000FF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Extract components
			r := tt.color.R()
			g := tt.color.G()
			b := tt.color.B()

			// Create new color from components
			newColor := types.NewColor(r, g, b)

			// Should match original
			if newColor != tt.color {
				t.Errorf("NewColor(%02x, %02x, %02x) = %06x, want %06x", r, g, b, newColor, tt.color)
			}
		})
	}
}

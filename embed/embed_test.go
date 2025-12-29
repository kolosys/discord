package embed_test

import (
	"strings"
	"testing"
	"time"

	"github.com/kolosys/discord/embed"
)

func TestEmbed_Valid(t *testing.T) {
	tests := []struct {
		name  string
		setup func(*embed.Builder)
	}{
		{"Empty", func(b *embed.Builder) {}},
		{"Title only", func(b *embed.Builder) { b.Title("Test") }},
		{"Description only", func(b *embed.Builder) { b.Description("Test") }},
		{"Title and description", func(b *embed.Builder) { b.Title("Test").Description("Desc") }},
		{"With color", func(b *embed.Builder) { b.Title("T").Color(0xFF0000) }},
		{"With field", func(b *embed.Builder) { b.Field("Name", "Value", false) }},
		{"With multiple fields", func(b *embed.Builder) {
			for i := 0; i < 5; i++ {
				b.Field("F"+string(rune(i)), "V"+string(rune(i)), false)
			}
		}},
		{"With author", func(b *embed.Builder) { b.Author("Author") }},
		{"With footer", func(b *embed.Builder) { b.Footer("Footer") }},
		{"With image", func(b *embed.Builder) { b.Image("https://example.com/image.png") }},
		{"With thumbnail", func(b *embed.Builder) { b.Thumbnail("https://example.com/thumb.png") }},
		{"All fields", func(b *embed.Builder) {
			b.Title("T").Description("D").Author("A").Footer("F").Color(0xFF0000).
				Image("img").Thumbnail("thumb").Field("N", "V", false)
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := embed.New()
			tt.setup(b)
			e := b.Build()
			if e == nil {
				t.Error("Build() returned nil")
			}
		})
	}
}

func TestEmbed_TitleTooLong(t *testing.T) {
	b := embed.New()
	b.Title(strings.Repeat("a", embed.MaxTitleLength+1))
	_, err := b.BuildSafe()
	if err == nil {
		t.Error("BuildSafe() should error for title too long")
	}
	if err != embed.ErrTitleTooLong {
		t.Errorf("Expected ErrTitleTooLong, got %v", err)
	}
}

func TestEmbed_DescriptionTooLong(t *testing.T) {
	b := embed.New()
	b.Description(strings.Repeat("a", embed.MaxDescriptionLength+1))
	_, err := b.BuildSafe()
	if err == nil {
		t.Error("BuildSafe() should error for description too long")
	}
	if err != embed.ErrDescriptionTooLong {
		t.Errorf("Expected ErrDescriptionTooLong, got %v", err)
	}
}

func TestEmbed_TooManyFields(t *testing.T) {
	b := embed.New()
	for i := 0; i <= embed.MaxFields; i++ {
		b.Field("Name", "Value", false)
	}
	_, err := b.BuildSafe()
	if err == nil {
		t.Error("BuildSafe() should error for too many fields")
	}
	if err != embed.ErrTooManyFields {
		t.Errorf("Expected ErrTooManyFields, got %v", err)
	}
}

func TestEmbed_FieldNameTooLong(t *testing.T) {
	b := embed.New()
	b.Field(strings.Repeat("a", embed.MaxFieldNameLength+1), "value", false)
	_, err := b.BuildSafe()
	if err == nil {
		t.Error("BuildSafe() should error for field name too long")
	}
	if err != embed.ErrFieldNameTooLong {
		t.Errorf("Expected ErrFieldNameTooLong, got %v", err)
	}
}

func TestEmbed_FieldValueTooLong(t *testing.T) {
	b := embed.New()
	b.Field("name", strings.Repeat("a", embed.MaxFieldValueLength+1), false)
	_, err := b.BuildSafe()
	if err == nil {
		t.Error("BuildSafe() should error for field value too long")
	}
	if err != embed.ErrFieldValueTooLong {
		t.Errorf("Expected ErrFieldValueTooLong, got %v", err)
	}
}

func TestEmbed_FooterTooLong(t *testing.T) {
	b := embed.New()
	b.Footer(strings.Repeat("a", embed.MaxFooterLength+1))
	_, err := b.BuildSafe()
	if err == nil {
		t.Error("BuildSafe() should error for footer too long")
	}
	if err != embed.ErrFooterTooLong {
		t.Errorf("Expected ErrFooterTooLong, got %v", err)
	}
}

func TestEmbed_AuthorTooLong(t *testing.T) {
	b := embed.New()
	b.Author(strings.Repeat("a", embed.MaxAuthorNameLength+1))
	_, err := b.BuildSafe()
	if err == nil {
		t.Error("BuildSafe() should error for author too long")
	}
	if err != embed.ErrAuthorNameTooLong {
		t.Errorf("Expected ErrAuthorNameTooLong, got %v", err)
	}
}

func TestEmbed_TotalTooLarge(t *testing.T) {
	b := embed.New()
	// Create content that exceeds total limit but doesn't hit individual limits
	// Use multiple shorter fields to avoid hitting field value limit
	b.Description(strings.Repeat("a", 3000))
	for i := 0; i < 10; i++ {
		b.Field("F", strings.Repeat("a", 300), false)
	}
	_, err := b.BuildSafe()
	if err == nil {
		t.Error("BuildSafe() should error for total length exceeded")
	}
	if err != embed.ErrEmbedTooLarge {
		t.Errorf("Expected ErrEmbedTooLarge, got %v", err)
	}
}

func TestEmbed_Chaining(t *testing.T) {
	b := embed.New()
	result := b.Title("T").Description("D").Color(0xFF0000)
	if result != b {
		t.Error("Builder methods should return self for chaining")
	}
}

func TestEmbed_Copy(t *testing.T) {
	original := embed.New().
		Title("Original").
		Description("Desc").
		Author("Auth").
		Field("F1", "V1", false)

	copy := original.Copy()

	// Verify copy has same values
	origEmbed := original.Build()
	copyEmbed := copy.Build()

	if origEmbed.Title != copyEmbed.Title {
		t.Errorf("Title mismatch: %s != %s", origEmbed.Title, copyEmbed.Title)
	}
	if origEmbed.Description != copyEmbed.Description {
		t.Errorf("Description mismatch: %s != %s", origEmbed.Description, copyEmbed.Description)
	}

	// Modify copy
	copy.Title("Modified")

	// Verify original unchanged
	origEmbed = original.Build()
	if origEmbed.Title != "Original" {
		t.Errorf("Original was modified: expected 'Original', got %q", origEmbed.Title)
	}
}

func TestEmbed_TimestampNow(t *testing.T) {
	b := embed.New()
	before := time.Now().UTC()
	b.TimestampNow()
	after := time.Now().UTC()

	e := b.Build()
	if e.Timestamp == "" {
		t.Error("TimestampNow() did not set timestamp")
	}

	// Parse timestamp
	ts, err := time.Parse(time.RFC3339, e.Timestamp)
	if err != nil {
		t.Errorf("Invalid timestamp format: %v", err)
	}

	// Verify timestamp is close to now (truncate to seconds since RFC3339 doesn't include nanoseconds)
	before = before.Truncate(time.Second)
	after = after.Truncate(time.Second).Add(time.Second)
	if ts.Before(before) || ts.After(after) {
		t.Errorf("Timestamp not close to now: %v (before: %v, after: %v)", ts, before, after)
	}
}

func TestEmbed_TimestampTime(t *testing.T) {
	b := embed.New()
	expectedTime := time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)
	b.TimestampTime(expectedTime)

	e := b.Build()
	if e.Timestamp == "" {
		t.Error("TimestampTime() did not set timestamp")
	}

	// Parse timestamp
	ts, err := time.Parse(time.RFC3339, e.Timestamp)
	if err != nil {
		t.Errorf("Invalid timestamp format: %v", err)
	}

	if !ts.Equal(expectedTime) {
		t.Errorf("Timestamp mismatch: %v != %v", ts, expectedTime)
	}
}

func TestEmbed_Colors(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func(*embed.Builder) *embed.Builder
		expected  int
	}{
		{"Success", func(b *embed.Builder) *embed.Builder { return b.Success() }, embed.ColorSuccess},
		{"Error", func(b *embed.Builder) *embed.Builder { return b.Error() }, embed.ColorError},
		{"Warning", func(b *embed.Builder) *embed.Builder { return b.Warning() }, embed.ColorWarning},
		{"Info", func(b *embed.Builder) *embed.Builder { return b.Info() }, embed.ColorInfo},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := embed.New()
			tt.setupFunc(b)
			e := b.Build()
			if e.Color != tt.expected {
				t.Errorf("Color = %d, want %d", e.Color, tt.expected)
			}
		})
	}
}

func TestEmbed_Presets(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func(string, string) *embed.Builder
		checkFunc func(*embed.Builder) bool
	}{
		{"SuccessEmbed", embed.SuccessEmbed, func(b *embed.Builder) bool {
			e := b.Build()
			return e.Color == embed.ColorSuccess
		}},
		{"ErrorEmbed", embed.ErrorEmbed, func(b *embed.Builder) bool {
			e := b.Build()
			return e.Color == embed.ColorError
		}},
		{"WarningEmbed", embed.WarningEmbed, func(b *embed.Builder) bool {
			e := b.Build()
			return e.Color == embed.ColorWarning
		}},
		{"InfoEmbed", embed.InfoEmbed, func(b *embed.Builder) bool {
			e := b.Build()
			return e.Color == embed.ColorInfo
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.setupFunc("Title", "Description")
			e := b.Build()
			if e.Title != "Title" {
				t.Errorf("Title = %q, want %q", e.Title, "Title")
			}
			if e.Description != "Description" {
				t.Errorf("Description = %q, want %q", e.Description, "Description")
			}
			if !tt.checkFunc(b) {
				t.Errorf("Color check failed for %s", tt.name)
			}
		})
	}
}

func TestEmbed_BuildSafe_Valid(t *testing.T) {
	b := embed.New().Title("Valid").Description("Embed")
	e, err := b.BuildSafe()
	if err != nil {
		t.Errorf("BuildSafe() error = %v, want nil", err)
	}
	if e == nil {
		t.Error("BuildSafe() returned nil embed")
	}
}

func TestEmbed_BuildSafe_Invalid(t *testing.T) {
	b := embed.New().Title(strings.Repeat("a", embed.MaxTitleLength+1))
	e, err := b.BuildSafe()
	if err == nil {
		t.Error("BuildSafe() should error for invalid embed")
	}
	if e != nil {
		t.Error("BuildSafe() should not return embed when error occurs")
	}
}

func TestEmbed_InlineField(t *testing.T) {
	b := embed.New().InlineField("Name", "Value")
	e := b.Build()
	if len(e.Fields) != 1 {
		t.Errorf("Expected 1 field, got %d", len(e.Fields))
	}
	if !e.Fields[0].Inline {
		t.Error("InlineField() should set Inline to true")
	}
}

func TestEmbed_BlockField(t *testing.T) {
	b := embed.New().BlockField("Name", "Value")
	e := b.Build()
	if len(e.Fields) != 1 {
		t.Errorf("Expected 1 field, got %d", len(e.Fields))
	}
	if e.Fields[0].Inline {
		t.Error("BlockField() should set Inline to false")
	}
}

func TestEmbed_AuthorFull(t *testing.T) {
	b := embed.New().AuthorFull("Name", "http://example.com", "http://example.com/icon.png")
	e := b.Build()
	if e.Author == nil {
		t.Fatal("Author is nil")
	}
	if e.Author.Name != "Name" {
		t.Errorf("Author.Name = %q, want %q", e.Author.Name, "Name")
	}
	if e.Author.URL != "http://example.com" {
		t.Errorf("Author.URL = %q, want %q", e.Author.URL, "http://example.com")
	}
	if e.Author.IconURL != "http://example.com/icon.png" {
		t.Errorf("Author.IconURL = %q, want %q", e.Author.IconURL, "http://example.com/icon.png")
	}
}

func TestEmbed_FooterFull(t *testing.T) {
	b := embed.New().FooterFull("Text", "http://example.com/icon.png")
	e := b.Build()
	if e.Footer == nil {
		t.Fatal("Footer is nil")
	}
	if e.Footer.Text != "Text" {
		t.Errorf("Footer.Text = %q, want %q", e.Footer.Text, "Text")
	}
	if e.Footer.IconURL != "http://example.com/icon.png" {
		t.Errorf("Footer.IconURL = %q, want %q", e.Footer.IconURL, "http://example.com/icon.png")
	}
}

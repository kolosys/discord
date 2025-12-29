package embed

import (
	"time"

	"github.com/kolosys/discord/commands"
)

// Builder provides a fluent API for constructing Discord embeds.
type Builder struct {
	embed commands.Embed
}

// New creates a new embed builder.
func New() *Builder {
	return &Builder{}
}

// Title sets the embed title.
func (b *Builder) Title(title string) *Builder {
	b.embed.Title = title
	return b
}

// Description sets the embed description.
func (b *Builder) Description(desc string) *Builder {
	b.embed.Description = desc
	return b
}

// URL sets the embed URL (makes title a hyperlink).
func (b *Builder) URL(url string) *Builder {
	b.embed.URL = url
	return b
}

// Color sets the embed color.
func (b *Builder) Color(color int) *Builder {
	b.embed.Color = color
	return b
}

// Timestamp sets the embed timestamp.
func (b *Builder) Timestamp(ts string) *Builder {
	b.embed.Timestamp = ts
	return b
}

// TimestampNow sets the timestamp to the current time.
func (b *Builder) TimestampNow() *Builder {
	b.embed.Timestamp = time.Now().UTC().Format(time.RFC3339)
	return b
}

// TimestampTime sets the timestamp from a time.Time value.
func (b *Builder) TimestampTime(t time.Time) *Builder {
	b.embed.Timestamp = t.UTC().Format(time.RFC3339)
	return b
}

// Author sets the embed author name.
func (b *Builder) Author(name string) *Builder {
	if b.embed.Author == nil {
		b.embed.Author = &commands.EmbedAuthor{}
	}
	b.embed.Author.Name = name
	return b
}

// AuthorIcon sets the author icon URL.
func (b *Builder) AuthorIcon(iconURL string) *Builder {
	if b.embed.Author == nil {
		b.embed.Author = &commands.EmbedAuthor{}
	}
	b.embed.Author.IconURL = iconURL
	return b
}

// AuthorURL sets the author URL (makes author name a hyperlink).
func (b *Builder) AuthorURL(url string) *Builder {
	if b.embed.Author == nil {
		b.embed.Author = &commands.EmbedAuthor{}
	}
	b.embed.Author.URL = url
	return b
}

// AuthorFull sets all author fields at once.
func (b *Builder) AuthorFull(name, url, iconURL string) *Builder {
	b.embed.Author = &commands.EmbedAuthor{
		Name:    name,
		URL:     url,
		IconURL: iconURL,
	}
	return b
}

// Footer sets the embed footer text.
func (b *Builder) Footer(text string) *Builder {
	if b.embed.Footer == nil {
		b.embed.Footer = &commands.EmbedFooter{}
	}
	b.embed.Footer.Text = text
	return b
}

// FooterIcon sets the footer icon URL.
func (b *Builder) FooterIcon(iconURL string) *Builder {
	if b.embed.Footer == nil {
		b.embed.Footer = &commands.EmbedFooter{}
	}
	b.embed.Footer.IconURL = iconURL
	return b
}

// FooterFull sets both footer text and icon.
func (b *Builder) FooterFull(text, iconURL string) *Builder {
	b.embed.Footer = &commands.EmbedFooter{
		Text:    text,
		IconURL: iconURL,
	}
	return b
}

// Image sets the embed image URL.
func (b *Builder) Image(url string) *Builder {
	b.embed.Image = &commands.EmbedImage{URL: url}
	return b
}

// Thumbnail sets the embed thumbnail URL.
func (b *Builder) Thumbnail(url string) *Builder {
	b.embed.Thumbnail = &commands.EmbedThumbnail{URL: url}
	return b
}

// Field adds a field to the embed.
func (b *Builder) Field(name, value string, inline bool) *Builder {
	b.embed.Fields = append(b.embed.Fields, &commands.EmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	})
	return b
}

// InlineField adds an inline field to the embed.
func (b *Builder) InlineField(name, value string) *Builder {
	return b.Field(name, value, true)
}

// BlockField adds a block (non-inline) field to the embed.
func (b *Builder) BlockField(name, value string) *Builder {
	return b.Field(name, value, false)
}

// Fields adds multiple fields at once.
func (b *Builder) Fields(fields ...*commands.EmbedField) *Builder {
	b.embed.Fields = append(b.embed.Fields, fields...)
	return b
}

// Build returns the constructed embed.
func (b *Builder) Build() *commands.Embed {
	return &b.embed
}

// BuildSafe validates and returns the embed, or an error if invalid.
func (b *Builder) BuildSafe() (*commands.Embed, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}
	return &b.embed, nil
}

// Validate checks if the embed conforms to Discord's limits.
func (b *Builder) Validate() error {
	if len(b.embed.Title) > MaxTitleLength {
		return ErrTitleTooLong
	}
	if len(b.embed.Description) > MaxDescriptionLength {
		return ErrDescriptionTooLong
	}
	if len(b.embed.Fields) > MaxFields {
		return ErrTooManyFields
	}
	if b.embed.Footer != nil && len(b.embed.Footer.Text) > MaxFooterLength {
		return ErrFooterTooLong
	}
	if b.embed.Author != nil && len(b.embed.Author.Name) > MaxAuthorNameLength {
		return ErrAuthorNameTooLong
	}
	for _, f := range b.embed.Fields {
		if len(f.Name) > MaxFieldNameLength {
			return ErrFieldNameTooLong
		}
		if len(f.Value) > MaxFieldValueLength {
			return ErrFieldValueTooLong
		}
	}
	if b.totalLength() > MaxTotalLength {
		return ErrEmbedTooLarge
	}
	return nil
}

// totalLength calculates the total character count of the embed.
func (b *Builder) totalLength() int {
	total := len(b.embed.Title) + len(b.embed.Description)
	if b.embed.Footer != nil {
		total += len(b.embed.Footer.Text)
	}
	if b.embed.Author != nil {
		total += len(b.embed.Author.Name)
	}
	for _, f := range b.embed.Fields {
		total += len(f.Name) + len(f.Value)
	}
	return total
}

// Copy returns a new builder with a copy of the current embed.
func (b *Builder) Copy() *Builder {
	newBuilder := &Builder{
		embed: commands.Embed{
			Title:       b.embed.Title,
			Type:        b.embed.Type,
			Description: b.embed.Description,
			URL:         b.embed.URL,
			Timestamp:   b.embed.Timestamp,
			Color:       b.embed.Color,
		},
	}
	if b.embed.Footer != nil {
		newBuilder.embed.Footer = &commands.EmbedFooter{
			Text:    b.embed.Footer.Text,
			IconURL: b.embed.Footer.IconURL,
		}
	}
	if b.embed.Image != nil {
		newBuilder.embed.Image = &commands.EmbedImage{URL: b.embed.Image.URL}
	}
	if b.embed.Thumbnail != nil {
		newBuilder.embed.Thumbnail = &commands.EmbedThumbnail{URL: b.embed.Thumbnail.URL}
	}
	if b.embed.Author != nil {
		newBuilder.embed.Author = &commands.EmbedAuthor{
			Name:    b.embed.Author.Name,
			URL:     b.embed.Author.URL,
			IconURL: b.embed.Author.IconURL,
		}
	}
	if len(b.embed.Fields) > 0 {
		newBuilder.embed.Fields = make([]*commands.EmbedField, len(b.embed.Fields))
		for i, f := range b.embed.Fields {
			newBuilder.embed.Fields[i] = &commands.EmbedField{
				Name:   f.Name,
				Value:  f.Value,
				Inline: f.Inline,
			}
		}
	}
	return newBuilder
}

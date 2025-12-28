package commands

import "github.com/kolosys/discord/models"

// OptionType represents the type of a command option.
type OptionType int32

const (
	OptionTypeSubCommand      OptionType = 1
	OptionTypeSubCommandGroup OptionType = 2
	OptionTypeString          OptionType = 3
	OptionTypeInteger         OptionType = 4
	OptionTypeBoolean         OptionType = 5
	OptionTypeUser            OptionType = 6
	OptionTypeChannel         OptionType = 7
	OptionTypeRole            OptionType = 8
	OptionTypeMentionable     OptionType = 9
	OptionTypeNumber          OptionType = 10
	OptionTypeAttachment      OptionType = 11
)

// Option represents a command option.
type Option struct {
	Name         string
	Description  string
	Type         OptionType
	Required     bool
	Choices      []Choice
	Options      []Option // For subcommands
	ChannelTypes []int32
	MinValue     *float64
	MaxValue     *float64
	MinLength    *int32
	MaxLength    *int32
	Autocomplete bool
}

// String creates a string option.
func String(name, description string) *OptionBuilder {
	return &OptionBuilder{opt: Option{Name: name, Description: description, Type: OptionTypeString}}
}

// Integer creates an integer option.
func Integer(name, description string) *OptionBuilder {
	return &OptionBuilder{opt: Option{Name: name, Description: description, Type: OptionTypeInteger}}
}

// Number creates a number (float) option.
func Number(name, description string) *OptionBuilder {
	return &OptionBuilder{opt: Option{Name: name, Description: description, Type: OptionTypeNumber}}
}

// Boolean creates a boolean option.
func Boolean(name, description string) *OptionBuilder {
	return &OptionBuilder{opt: Option{Name: name, Description: description, Type: OptionTypeBoolean}}
}

// User creates a user option.
func User(name, description string) *OptionBuilder {
	return &OptionBuilder{opt: Option{Name: name, Description: description, Type: OptionTypeUser}}
}

// Channel creates a channel option.
func Channel(name, description string) *OptionBuilder {
	return &OptionBuilder{opt: Option{Name: name, Description: description, Type: OptionTypeChannel}}
}

// Role creates a role option.
func Role(name, description string) *OptionBuilder {
	return &OptionBuilder{opt: Option{Name: name, Description: description, Type: OptionTypeRole}}
}

// Mentionable creates a mentionable option.
func Mentionable(name, description string) *OptionBuilder {
	return &OptionBuilder{opt: Option{Name: name, Description: description, Type: OptionTypeMentionable}}
}

// Attachment creates an attachment option.
func Attachment(name, description string) *OptionBuilder {
	return &OptionBuilder{opt: Option{Name: name, Description: description, Type: OptionTypeAttachment}}
}

// OptionBuilder builds command options with a fluent API.
type OptionBuilder struct {
	opt Option
}

func (b *OptionBuilder) Required() *OptionBuilder {
	b.opt.Required = true
	return b
}

func (b *OptionBuilder) Choices(choices ...Choice) *OptionBuilder {
	b.opt.Choices = choices
	return b
}

func (b *OptionBuilder) MinValue(min float64) *OptionBuilder {
	b.opt.MinValue = &min
	return b
}

func (b *OptionBuilder) MaxValue(max float64) *OptionBuilder {
	b.opt.MaxValue = &max
	return b
}

func (b *OptionBuilder) MinLength(min int32) *OptionBuilder {
	b.opt.MinLength = &min
	return b
}

func (b *OptionBuilder) MaxLength(max int32) *OptionBuilder {
	b.opt.MaxLength = &max
	return b
}

func (b *OptionBuilder) ChannelTypes(types ...int32) *OptionBuilder {
	b.opt.ChannelTypes = types
	return b
}

func (b *OptionBuilder) Autocomplete() *OptionBuilder {
	b.opt.Autocomplete = true
	return b
}

func (b *OptionBuilder) Build() Option {
	return b.opt
}

// OptionMap provides type-safe access to command options.
type OptionMap struct {
	options  map[string]any
	resolved *ResolvedData
}

// ResolvedData contains resolved entities from interaction options.
type ResolvedData struct {
	Users       map[string]*models.User
	Members     map[string]*models.GuildMember
	Roles       map[string]*models.GuildRole
	Channels    map[string]*models.GuildChannel
	Messages    map[string]*models.Message
	Attachments map[string]*models.Attachment
}

// NewOptionMap creates a new OptionMap.
func NewOptionMap() *OptionMap {
	return &OptionMap{
		options: make(map[string]any),
		resolved: &ResolvedData{
			Users:       make(map[string]*models.User),
			Members:     make(map[string]*models.GuildMember),
			Roles:       make(map[string]*models.GuildRole),
			Channels:    make(map[string]*models.GuildChannel),
			Messages:    make(map[string]*models.Message),
			Attachments: make(map[string]*models.Attachment),
		},
	}
}

// Set sets an option value.
func (m *OptionMap) Set(name string, value any) {
	m.options[name] = value
}

// Get returns the raw option value.
func (m *OptionMap) Get(name string) (any, bool) {
	v, ok := m.options[name]
	return v, ok
}

// String returns a string option value.
func (m *OptionMap) String(name string) string {
	v, ok := m.options[name]
	if !ok {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

// StringDefault returns a string option value or a default.
func (m *OptionMap) StringDefault(name, def string) string {
	if s := m.String(name); s != "" {
		return s
	}
	return def
}

// Int returns an integer option value.
func (m *OptionMap) Int(name string) int64 {
	v, ok := m.options[name]
	if !ok {
		return 0
	}
	switch n := v.(type) {
	case int64:
		return n
	case float64:
		return int64(n)
	case int:
		return int64(n)
	}
	return 0
}

// IntDefault returns an integer option value or a default.
func (m *OptionMap) IntDefault(name string, def int64) int64 {
	if _, ok := m.options[name]; !ok {
		return def
	}
	return m.Int(name)
}

// Float returns a number option value.
func (m *OptionMap) Float(name string) float64 {
	v, ok := m.options[name]
	if !ok {
		return 0
	}
	switch n := v.(type) {
	case float64:
		return n
	case int64:
		return float64(n)
	case int:
		return float64(n)
	}
	return 0
}

// FloatDefault returns a number option value or a default.
func (m *OptionMap) FloatDefault(name string, def float64) float64 {
	if _, ok := m.options[name]; !ok {
		return def
	}
	return m.Float(name)
}

// Bool returns a boolean option value.
func (m *OptionMap) Bool(name string) bool {
	v, ok := m.options[name]
	if !ok {
		return false
	}
	if b, ok := v.(bool); ok {
		return b
	}
	return false
}

// User returns a resolved user.
func (m *OptionMap) User(name string) *models.User {
	id := m.String(name)
	if id == "" {
		return nil
	}
	return m.resolved.Users[id]
}

// Member returns a resolved guild member.
func (m *OptionMap) Member(name string) *models.GuildMember {
	id := m.String(name)
	if id == "" {
		return nil
	}
	return m.resolved.Members[id]
}

// Role returns a resolved role.
func (m *OptionMap) Role(name string) *models.GuildRole {
	id := m.String(name)
	if id == "" {
		return nil
	}
	return m.resolved.Roles[id]
}

// Channel returns a resolved channel.
func (m *OptionMap) Channel(name string) *models.GuildChannel {
	id := m.String(name)
	if id == "" {
		return nil
	}
	return m.resolved.Channels[id]
}

// Attachment returns a resolved attachment.
func (m *OptionMap) Attachment(name string) *models.Attachment {
	id := m.String(name)
	if id == "" {
		return nil
	}
	return m.resolved.Attachments[id]
}

// Message returns a resolved message.
func (m *OptionMap) Message(name string) *models.Message {
	id := m.String(name)
	if id == "" {
		return nil
	}
	return m.resolved.Messages[id]
}

// Resolved returns the resolved data for direct access.
func (m *OptionMap) Resolved() *ResolvedData {
	return m.resolved
}

// SetResolved sets the resolved data.
func (m *OptionMap) SetResolved(r *ResolvedData) {
	if r != nil {
		m.resolved = r
	}
}

// Has returns true if an option was provided.
func (m *OptionMap) Has(name string) bool {
	_, ok := m.options[name]
	return ok
}

// All returns all options as a map.
func (m *OptionMap) All() map[string]any {
	result := make(map[string]any, len(m.options))
	for k, v := range m.options {
		result[k] = v
	}
	return result
}

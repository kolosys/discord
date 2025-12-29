package commands_test

import (
	"testing"

	"github.com/kolosys/discord/commands"
)

func TestString_Creation(t *testing.T) {
	opt := commands.String("name", "description")
	if opt == nil {
		t.Fatal("String() returned nil")
	}

	built := opt.Build()
	if built.Name != "name" {
		t.Errorf("Name = %q, want %q", built.Name, "name")
	}
	if built.Description != "description" {
		t.Errorf("Description = %q, want %q", built.Description, "description")
	}
	if built.Type != commands.OptionTypeString {
		t.Errorf("Type = %v, want %v", built.Type, commands.OptionTypeString)
	}
	if built.Required {
		t.Error("String option should not be required by default")
	}
}

func TestInteger_Creation(t *testing.T) {
	opt := commands.Integer("number", "A number")
	built := opt.Build()

	if built.Type != commands.OptionTypeInteger {
		t.Errorf("Type = %v, want %v", built.Type, commands.OptionTypeInteger)
	}
}

func TestNumber_Creation(t *testing.T) {
	opt := commands.Number("decimal", "A decimal")
	built := opt.Build()

	if built.Type != commands.OptionTypeNumber {
		t.Errorf("Type = %v, want %v", built.Type, commands.OptionTypeNumber)
	}
}

func TestBoolean_Creation(t *testing.T) {
	opt := commands.Boolean("flag", "A flag")
	built := opt.Build()

	if built.Type != commands.OptionTypeBoolean {
		t.Errorf("Type = %v, want %v", built.Type, commands.OptionTypeBoolean)
	}
}

func TestUser_Creation(t *testing.T) {
	opt := commands.User("user", "A user")
	built := opt.Build()

	if built.Type != commands.OptionTypeUser {
		t.Errorf("Type = %v, want %v", built.Type, commands.OptionTypeUser)
	}
}

func TestChannel_Creation(t *testing.T) {
	opt := commands.Channel("channel", "A channel")
	built := opt.Build()

	if built.Type != commands.OptionTypeChannel {
		t.Errorf("Type = %v, want %v", built.Type, commands.OptionTypeChannel)
	}
}

func TestRole_Creation(t *testing.T) {
	opt := commands.Role("role", "A role")
	built := opt.Build()

	if built.Type != commands.OptionTypeRole {
		t.Errorf("Type = %v, want %v", built.Type, commands.OptionTypeRole)
	}
}

func TestMentionable_Creation(t *testing.T) {
	opt := commands.Mentionable("mentionable", "A mentionable")
	built := opt.Build()

	if built.Type != commands.OptionTypeMentionable {
		t.Errorf("Type = %v, want %v", built.Type, commands.OptionTypeMentionable)
	}
}

func TestAttachment_Creation(t *testing.T) {
	opt := commands.Attachment("file", "A file")
	built := opt.Build()

	if built.Type != commands.OptionTypeAttachment {
		t.Errorf("Type = %v, want %v", built.Type, commands.OptionTypeAttachment)
	}
}

func TestOptionBuilder_Required(t *testing.T) {
	opt := commands.String("name", "desc").Required()
	built := opt.Build()

	if !built.Required {
		t.Error("Required() should set required to true")
	}
}

func TestOptionBuilder_Chaining(t *testing.T) {
	opt := commands.String("name", "desc")
	result := opt.Required().MinLength(1).MaxLength(100)

	if result != opt {
		t.Error("Builder methods should return self for chaining")
	}
}

func TestOptionBuilder_MinValue(t *testing.T) {
	opt := commands.Integer("num", "desc").MinValue(5)
	built := opt.Build()

	if built.MinValue == nil || *built.MinValue != 5 {
		t.Error("MinValue not set correctly")
	}
}

func TestOptionBuilder_MaxValue(t *testing.T) {
	opt := commands.Integer("num", "desc").MaxValue(100)
	built := opt.Build()

	if built.MaxValue == nil || *built.MaxValue != 100 {
		t.Error("MaxValue not set correctly")
	}
}

func TestOptionBuilder_MinMaxValue(t *testing.T) {
	opt := commands.Integer("num", "desc").MinValue(5).MaxValue(100)
	built := opt.Build()

	if built.MinValue == nil || *built.MinValue != 5 {
		t.Error("MinValue not set")
	}
	if built.MaxValue == nil || *built.MaxValue != 100 {
		t.Error("MaxValue not set")
	}
}

func TestOptionBuilder_MinLength(t *testing.T) {
	opt := commands.String("str", "desc").MinLength(3)
	built := opt.Build()

	if built.MinLength == nil || *built.MinLength != 3 {
		t.Error("MinLength not set correctly")
	}
}

func TestOptionBuilder_MaxLength(t *testing.T) {
	opt := commands.String("str", "desc").MaxLength(255)
	built := opt.Build()

	if built.MaxLength == nil || *built.MaxLength != 255 {
		t.Error("MaxLength not set correctly")
	}
}

func TestOptionBuilder_Choices(t *testing.T) {
	choices := []commands.Choice{
		{Name: "Choice 1", Value: "value1"},
		{Name: "Choice 2", Value: "value2"},
	}

	opt := commands.String("choice", "desc").Choices(choices...)
	built := opt.Build()

	if len(built.Choices) != 2 {
		t.Errorf("Expected 2 choices, got %d", len(built.Choices))
	}

	if built.Choices[0].Name != "Choice 1" {
		t.Errorf("Choice name = %q, want %q", built.Choices[0].Name, "Choice 1")
	}
}

func TestOptionBuilder_ChannelTypes(t *testing.T) {
	opt := commands.Channel("chan", "desc").ChannelTypes(0, 1, 4)
	built := opt.Build()

	if len(built.ChannelTypes) != 3 {
		t.Errorf("Expected 3 channel types, got %d", len(built.ChannelTypes))
	}

	if built.ChannelTypes[0] != 0 || built.ChannelTypes[1] != 1 || built.ChannelTypes[2] != 4 {
		t.Error("ChannelTypes mismatch")
	}
}

func TestOptionBuilder_Autocomplete(t *testing.T) {
	opt := commands.String("search", "desc").Autocomplete()
	built := opt.Build()

	if !built.Autocomplete {
		t.Error("Autocomplete() should set autocomplete to true")
	}
}

func TestChoice_Construction(t *testing.T) {
	choice := commands.Choice{
		Name:  "Option",
		Value: "value",
	}

	if choice.Name != "Option" {
		t.Errorf("Name = %q, want %q", choice.Name, "Option")
	}
	if choice.Value != "value" {
		t.Errorf("Value = %q, want %q", choice.Value, "value")
	}
}

func TestOptionTypes(t *testing.T) {
	tests := []struct {
		name     string
		builder  *commands.OptionBuilder
		expected commands.OptionType
	}{
		{"String", commands.String("n", "d"), commands.OptionTypeString},
		{"Integer", commands.Integer("n", "d"), commands.OptionTypeInteger},
		{"Boolean", commands.Boolean("n", "d"), commands.OptionTypeBoolean},
		{"User", commands.User("n", "d"), commands.OptionTypeUser},
		{"Channel", commands.Channel("n", "d"), commands.OptionTypeChannel},
		{"Role", commands.Role("n", "d"), commands.OptionTypeRole},
		{"Mentionable", commands.Mentionable("n", "d"), commands.OptionTypeMentionable},
		{"Attachment", commands.Attachment("n", "d"), commands.OptionTypeAttachment},
		{"Number", commands.Number("n", "d"), commands.OptionTypeNumber},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			built := tt.builder.Build()
			if built.Type != tt.expected {
				t.Errorf("Type = %v, want %v", built.Type, tt.expected)
			}
		})
	}
}

func TestOption_Build(t *testing.T) {
	builder := commands.String("name", "description").
		Required().
		MinLength(1).
		MaxLength(100).
		Autocomplete()

	opt := builder.Build()

	if opt.Name != "name" {
		t.Errorf("Name = %q, want %q", opt.Name, "name")
	}
	if opt.Description != "description" {
		t.Errorf("Description = %q, want %q", opt.Description, "description")
	}
	if opt.Type != commands.OptionTypeString {
		t.Errorf("Type = %v, want %v", opt.Type, commands.OptionTypeString)
	}
	if !opt.Required {
		t.Error("Required should be true")
	}
	if opt.MinLength == nil || *opt.MinLength != 1 {
		t.Error("MinLength should be 1")
	}
	if opt.MaxLength == nil || *opt.MaxLength != 100 {
		t.Error("MaxLength should be 100")
	}
	if !opt.Autocomplete {
		t.Error("Autocomplete should be true")
	}
}

func TestIntegerOption_MinMaxConstraints(t *testing.T) {
	opt := commands.Integer("level", "User level").
		MinValue(1).
		MaxValue(100)

	built := opt.Build()

	if built.MinValue == nil || *built.MinValue != 1 {
		t.Error("MinValue should be 1")
	}
	if built.MaxValue == nil || *built.MaxValue != 100 {
		t.Error("MaxValue should be 100")
	}
}

func TestStringOption_LengthConstraints(t *testing.T) {
	opt := commands.String("username", "User name").
		MinLength(3).
		MaxLength(32)

	built := opt.Build()

	if built.MinLength == nil || *built.MinLength != 3 {
		t.Error("MinLength should be 3")
	}
	if built.MaxLength == nil || *built.MaxLength != 32 {
		t.Error("MaxLength should be 32")
	}
}

func TestNumberOption_Floats(t *testing.T) {
	opt := commands.Number("price", "Price").
		MinValue(0.01).
		MaxValue(9999.99)

	built := opt.Build()

	if built.MinValue == nil || *built.MinValue != 0.01 {
		t.Error("MinValue should be 0.01")
	}
	if built.MaxValue == nil || *built.MaxValue != 9999.99 {
		t.Error("MaxValue should be 9999.99")
	}
}

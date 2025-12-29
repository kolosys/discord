package commands_test

import (
	"testing"

	"github.com/kolosys/discord/commands"
	"github.com/kolosys/discord/models"
)

func TestOptionMap_Creation(t *testing.T) {
	m := commands.NewOptionMap()
	if m == nil {
		t.Fatal("NewOptionMap() returned nil")
	}
}

func TestOptionMap_String(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Normal string", "hello", "hello"},
		{"With spaces", "hello world", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := commands.NewOptionMap()
			m.Set("test", tt.value)

			got := m.String("test")
			if got != tt.expected {
				t.Errorf("String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestOptionMap_String_Missing(t *testing.T) {
	m := commands.NewOptionMap()
	got := m.String("nonexistent")

	if got != "" {
		t.Errorf("String() for missing key = %q, want empty string", got)
	}
}

func TestOptionMap_StringDefault(t *testing.T) {
	tests := []struct {
		name   string
		set    bool
		value  string
		def    string
		expect string
	}{
		{"Missing uses default", false, "", "default", "default"},
		{"Set uses value", true, "custom", "default", "custom"},
		{"Empty string uses default", true, "", "default", "default"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := commands.NewOptionMap()
			if tt.set {
				m.Set("test", tt.value)
			}

			got := m.StringDefault("test", tt.def)
			if got != tt.expect {
				t.Errorf("StringDefault() = %q, want %q", got, tt.expect)
			}
		})
	}
}

func TestOptionMap_Int(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected int64
	}{
		{"int64", int64(42), 42},
		{"int", int(42), 42},
		{"float64", float64(42.0), 42},
		{"Missing", nil, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := commands.NewOptionMap()
			if tt.value != nil {
				m.Set("test", tt.value)
			}

			got := m.Int("test")
			if got != tt.expected {
				t.Errorf("Int() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestOptionMap_IntDefault(t *testing.T) {
	tests := []struct {
		name   string
		set    bool
		value  int64
		def    int64
		expect int64
	}{
		{"Missing uses default", false, 0, 100, 100},
		{"Set uses value", true, 42, 100, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := commands.NewOptionMap()
			if tt.set {
				m.Set("test", tt.value)
			}

			got := m.IntDefault("test", tt.def)
			if got != tt.expect {
				t.Errorf("IntDefault() = %d, want %d", got, tt.expect)
			}
		})
	}
}

func TestOptionMap_Float(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected float64
	}{
		{"float64", float64(3.14), 3.14},
		{"int64", int64(42), 42.0},
		{"int", int(42), 42.0},
		{"Missing", nil, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := commands.NewOptionMap()
			if tt.value != nil {
				m.Set("test", tt.value)
			}

			got := m.Float("test")
			if got != tt.expected {
				t.Errorf("Float() = %f, want %f", got, tt.expected)
			}
		})
	}
}

func TestOptionMap_FloatDefault(t *testing.T) {
	tests := []struct {
		name   string
		set    bool
		value  float64
		def    float64
		expect float64
	}{
		{"Missing uses default", false, 0, 3.14, 3.14},
		{"Set uses value", true, 2.71, 3.14, 2.71},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := commands.NewOptionMap()
			if tt.set {
				m.Set("test", tt.value)
			}

			got := m.FloatDefault("test", tt.def)
			if got != tt.expect {
				t.Errorf("FloatDefault() = %f, want %f", got, tt.expect)
			}
		})
	}
}

func TestOptionMap_Bool(t *testing.T) {
	tests := []struct {
		name     string
		value    any
		expected bool
	}{
		{"true", true, true},
		{"false", false, false},
		{"Missing", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := commands.NewOptionMap()
			if tt.value != nil {
				m.Set("test", tt.value)
			}

			got := m.Bool("test")
			if got != tt.expected {
				t.Errorf("Bool() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestOptionMap_User(t *testing.T) {
	m := commands.NewOptionMap()
	m.Set("user_id", "123")

	user := &models.User{ID: "123", Username: "testuser"}
	m.Resolved().Users["123"] = user

	got := m.User("user_id")
	if got == nil {
		t.Fatal("User() returned nil")
	}
	if got.ID != "123" {
		t.Errorf("User.ID = %q, want %q", got.ID, "123")
	}
	if got.Username != "testuser" {
		t.Errorf("User.Username = %q, want %q", got.Username, "testuser")
	}
}

func TestOptionMap_User_Missing(t *testing.T) {
	m := commands.NewOptionMap()
	got := m.User("nonexistent")

	if got != nil {
		t.Error("User() for missing key should return nil")
	}
}

func TestOptionMap_Member(t *testing.T) {
	m := commands.NewOptionMap()
	m.Set("member_id", "456")

	user := &models.User{ID: "456", Username: "member"}
	member := &models.GuildMember{User: user}
	m.Resolved().Members["456"] = member

	got := m.Member("member_id")
	if got == nil {
		t.Fatal("Member() returned nil")
	}
	if got.User == nil {
		t.Fatal("Member.User is nil")
	}
	// User is stored as any, so we need to cast it
	if userAny, ok := got.User.(*models.User); ok {
		if userAny.ID != "456" {
			t.Errorf("Member.User.ID = %q, want %q", userAny.ID, "456")
		}
	} else {
		t.Error("Member.User is not a *models.User")
	}
}

func TestOptionMap_Role(t *testing.T) {
	m := commands.NewOptionMap()
	m.Set("role_id", "789")

	role := &models.GuildRole{ID: "789", Name: "admin"}
	m.Resolved().Roles["789"] = role

	got := m.Role("role_id")
	if got == nil {
		t.Fatal("Role() returned nil")
	}
	if got.ID != "789" {
		t.Errorf("Role.ID = %q, want %q", got.ID, "789")
	}
	if got.Name != "admin" {
		t.Errorf("Role.Name = %q, want %q", got.Name, "admin")
	}
}

func TestOptionMap_Channel(t *testing.T) {
	m := commands.NewOptionMap()
	m.Set("channel_id", "999")

	channel := &models.GuildChannel{ID: "999", Name: "general"}
	m.Resolved().Channels["999"] = channel

	got := m.Channel("channel_id")
	if got == nil {
		t.Fatal("Channel() returned nil")
	}
	if got.ID != "999" {
		t.Errorf("Channel.ID = %q, want %q", got.ID, "999")
	}
	if got.Name != "general" {
		t.Errorf("Channel.Name = %q, want %q", got.Name, "general")
	}
}

func TestOptionMap_Attachment(t *testing.T) {
	m := commands.NewOptionMap()
	m.Set("file_id", "111")

	attachment := &models.Attachment{ID: "111", Filename: "test.txt"}
	m.Resolved().Attachments["111"] = attachment

	got := m.Attachment("file_id")
	if got == nil {
		t.Fatal("Attachment() returned nil")
	}
	if got.ID != "111" {
		t.Errorf("Attachment.ID = %q, want %q", got.ID, "111")
	}
	if got.Filename != "test.txt" {
		t.Errorf("Attachment.Filename = %q, want %q", got.Filename, "test.txt")
	}
}

func TestOptionMap_Message(t *testing.T) {
	m := commands.NewOptionMap()
	m.Set("message_id", "222")

	message := &models.Message{ID: "222", Content: "test"}
	m.Resolved().Messages["222"] = message

	got := m.Message("message_id")
	if got == nil {
		t.Fatal("Message() returned nil")
	}
	if got.ID != "222" {
		t.Errorf("Message.ID = %q, want %q", got.ID, "222")
	}
	if got.Content != "test" {
		t.Errorf("Message.Content = %q, want %q", got.Content, "test")
	}
}

func TestOptionMap_Get(t *testing.T) {
	m := commands.NewOptionMap()
	m.Set("key", "value")

	val, ok := m.Get("key")
	if !ok {
		t.Error("Get() should return true for existing key")
	}
	if val != "value" {
		t.Errorf("Get() value = %v, want %q", val, "value")
	}

	_, notOk := m.Get("nonexistent")
	if notOk {
		t.Error("Get() should return false for missing key")
	}
}

func TestOptionMap_Has(t *testing.T) {
	m := commands.NewOptionMap()
	m.Set("test", "value")

	if !m.Has("test") {
		t.Error("Has() should return true for existing key")
	}

	if m.Has("nonexistent") {
		t.Error("Has() should return false for missing key")
	}
}

func TestOptionMap_All(t *testing.T) {
	m := commands.NewOptionMap()
	m.Set("key1", "value1")
	m.Set("key2", 42)
	m.Set("key3", true)

	all := m.All()
	if len(all) != 3 {
		t.Errorf("All() returned %d items, want 3", len(all))
	}

	if all["key1"] != "value1" {
		t.Error("All() missing key1")
	}
	if all["key2"] != 42 {
		t.Error("All() missing key2")
	}
	if all["key3"] != true {
		t.Error("All() missing key3")
	}
}

func TestOptionMap_Set(t *testing.T) {
	m := commands.NewOptionMap()
	m.Set("test", "value")

	got := m.String("test")
	if got != "value" {
		t.Errorf("Set() didn't store value correctly: %q", got)
	}

	// Overwrite
	m.Set("test", "updated")
	got = m.String("test")
	if got != "updated" {
		t.Errorf("Set() didn't overwrite value: %q", got)
	}
}

func TestOptionMap_Resolved(t *testing.T) {
	m := commands.NewOptionMap()
	resolved := m.Resolved()

	if resolved == nil {
		t.Fatal("Resolved() returned nil")
	}

	if resolved.Users == nil {
		t.Error("Resolved().Users is nil")
	}
	if resolved.Members == nil {
		t.Error("Resolved().Members is nil")
	}
	if resolved.Roles == nil {
		t.Error("Resolved().Roles is nil")
	}
	if resolved.Channels == nil {
		t.Error("Resolved().Channels is nil")
	}
}

func TestOptionMap_SetResolved(t *testing.T) {
	m := commands.NewOptionMap()
	m.Set("user_option", "123")

	customResolved := &commands.ResolvedData{
		Users: make(map[string]*models.User),
	}
	customResolved.Users["123"] = &models.User{ID: "123"}

	m.SetResolved(customResolved)
	got := m.Resolved()

	if got != customResolved {
		t.Error("SetResolved() didn't set the resolved data")
	}

	user := m.User("user_option")
	if user == nil {
		t.Error("User() should find the resolved user")
	}
	if user.ID != "123" {
		t.Errorf("User.ID = %q, want %q", user.ID, "123")
	}
}

func TestOptionMap_SetResolvedNil(t *testing.T) {
	m := commands.NewOptionMap()
	original := m.Resolved()

	// SetResolved with nil should not change anything
	m.SetResolved(nil)
	afterNil := m.Resolved()

	if original != afterNil {
		t.Error("SetResolved(nil) should not change resolved data")
	}
}

func TestOptionMap_MixedTypes(t *testing.T) {
	m := commands.NewOptionMap()
	m.Set("string", "hello")
	m.Set("int", int64(42))
	m.Set("float", 3.14)
	m.Set("bool", true)

	if m.String("string") != "hello" {
		t.Error("String access failed")
	}
	if m.Int("int") != 42 {
		t.Error("Int access failed")
	}
	if m.Float("float") != 3.14 {
		t.Error("Float access failed")
	}
	if !m.Bool("bool") {
		t.Error("Bool access failed")
	}
}

func TestOptionMap_TypeConversion(t *testing.T) {
	m := commands.NewOptionMap()

	// Store int64, retrieve as float64
	m.Set("number", int64(42))
	floatVal := m.Float("number")
	if floatVal != 42.0 {
		t.Errorf("Type conversion int64 to float64 = %f, want 42.0", floatVal)
	}

	// Store float64, retrieve as int64
	m.Set("decimal", 99.5)
	intVal := m.Int("decimal")
	if intVal != 99 {
		t.Errorf("Type conversion float64 to int64 = %d, want 99", intVal)
	}
}

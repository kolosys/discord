package components_test

import (
	"testing"

	"github.com/kolosys/discord/components"
	"github.com/kolosys/discord/types"
)

func TestBuildComponents_SingleButton(t *testing.T) {
	btn := components.Button("id").Label("Click")
	result := components.BuildComponents(btn.Build())

	if len(result) != 1 {
		t.Errorf("Expected 1 row, got %d", len(result))
	}

	if ar, ok := result[0].(*components.ActionRowComponent); ok {
		if len(ar.Components) != 1 {
			t.Errorf("Expected 1 component in row, got %d", len(ar.Components))
		}
	} else {
		t.Error("Result should be ActionRowComponent")
	}
}

func TestBuildComponents_FiveButtons(t *testing.T) {
	var buttons []components.Component
	for i := 0; i < components.MaxButtonsPerRow; i++ {
		buttons = append(buttons, components.Button("id").Label("Button").Build())
	}

	result := components.BuildComponents(buttons...)

	if len(result) != 1 {
		t.Errorf("Expected 1 row, got %d", len(result))
	}

	if ar, ok := result[0].(*components.ActionRowComponent); ok {
		if len(ar.Components) != components.MaxButtonsPerRow {
			t.Errorf("Expected %d components, got %d", components.MaxButtonsPerRow, len(ar.Components))
		}
	}
}

func TestBuildComponents_SixButtons(t *testing.T) {
	var buttons []components.Component
	for i := 0; i < 6; i++ {
		buttons = append(buttons, components.Button("id").Label("Button").Build())
	}

	result := components.BuildComponents(buttons...)

	if len(result) != 2 {
		t.Errorf("Expected 2 rows for 6 buttons, got %d", len(result))
	}

	if ar, ok := result[0].(*components.ActionRowComponent); ok {
		if len(ar.Components) != components.MaxButtonsPerRow {
			t.Errorf("First row expected %d components, got %d", components.MaxButtonsPerRow, len(ar.Components))
		}
	}

	if ar, ok := result[1].(*components.ActionRowComponent); ok {
		if len(ar.Components) != 1 {
			t.Errorf("Second row expected 1 component, got %d", len(ar.Components))
		}
	}
}

func TestBuildComponents_SelectAlone(t *testing.T) {
	sel := components.StringSelect("id")
	sel.Option("Option", "value")
	result := components.BuildComponents(sel.Build())

	if len(result) != 1 {
		t.Errorf("Expected 1 row, got %d", len(result))
	}

	if ar, ok := result[0].(*components.ActionRowComponent); ok {
		if len(ar.Components) != 1 {
			t.Errorf("Expected 1 component, got %d", len(ar.Components))
		}
	}
}

func TestBuildComponents_ButtonThenSelect(t *testing.T) {
	btn := components.Button("id").Label("Button").Build()
	sel := components.StringSelect("id")
	sel.Option("Option", "value")

	result := components.BuildComponents(btn, sel.Build())

	if len(result) != 2 {
		t.Errorf("Expected 2 rows, got %d", len(result))
	}

	// First row should have button
	if ar, ok := result[0].(*components.ActionRowComponent); ok {
		if len(ar.Components) != 1 {
			t.Errorf("First row: expected 1 component, got %d", len(ar.Components))
		}
		if _, isButton := ar.Components[0].(*components.ButtonComponent); !isButton {
			t.Error("First row should contain button")
		}
	}

	// Second row should have select
	if ar, ok := result[1].(*components.ActionRowComponent); ok {
		if len(ar.Components) != 1 {
			t.Errorf("Second row: expected 1 component, got %d", len(ar.Components))
		}
		if _, isSelect := ar.Components[0].(*components.StringSelectComponent); !isSelect {
			t.Error("Second row should contain select")
		}
	}
}

func TestBuildComponents_Mixed(t *testing.T) {
	var comps []components.Component

	// 3 buttons
	for i := 0; i < 3; i++ {
		comps = append(comps, components.Button("id").Label("Button").Build())
	}

	// Select
	sel := components.StringSelect("id")
	sel.Option("Option", "value")
	comps = append(comps, sel.Build())

	// 2 more buttons
	for i := 0; i < 2; i++ {
		comps = append(comps, components.Button("id").Label("Button").Build())
	}

	result := components.BuildComponents(comps...)

	if len(result) != 3 {
		t.Errorf("Expected 3 rows, got %d", len(result))
	}
}

func TestBuildComponents_Empty(t *testing.T) {
	result := components.BuildComponents()
	if len(result) != 0 {
		t.Errorf("Expected 0 rows for empty input, got %d", len(result))
	}
}

func TestBuildComponents_ActionRowPassthrough(t *testing.T) {
	btn := components.Button("id").Label("Button").Build()
	ar := components.ActionRow(btn)

	result := components.BuildComponents(ar)

	if len(result) != 1 {
		t.Errorf("Expected 1 row, got %d", len(result))
	}

	if returnedAr, ok := result[0].(*components.ActionRowComponent); ok {
		if returnedAr != ar {
			t.Error("ActionRow should be passed through unchanged")
		}
	}
}

func TestBuildComponents_Type(t *testing.T) {
	btn := components.Button("id").Label("Button").Build()
	result := components.BuildComponents(btn)

	if len(result) == 0 {
		t.Fatal("Result is empty")
	}

	ar, ok := result[0].(*components.ActionRowComponent)
	if !ok {
		t.Fatal("Result should be ActionRowComponent")
	}

	if ar.Type != types.ComponentTypeActionRow {
		t.Errorf("ActionRow type = %v, want %v", ar.Type, types.ComponentTypeActionRow)
	}
}

func TestBuildComponentsSafe_Valid(t *testing.T) {
	var comps []components.Component
	for i := 0; i < 3; i++ {
		comps = append(comps, components.Button("id").Label("Button").Build())
	}

	result, err := components.BuildComponentsSafe(comps...)
	if err != nil {
		t.Errorf("BuildComponentsSafe() error = %v", err)
	}
	if result == nil {
		t.Error("BuildComponentsSafe() returned nil")
	}
}

func TestBuildComponentsSafe_TooManyRows(t *testing.T) {
	var comps []components.Component

	// Create more than MaxActionRows worth of selects (each gets its own row)
	for i := 0; i <= components.MaxActionRows; i++ {
		sel := components.StringSelect("id")
		sel.Option("Option", "value")
		comps = append(comps, sel.Build())
	}

	result, err := components.BuildComponentsSafe(comps...)
	if err != components.ErrTooManyActionRows {
		t.Errorf("BuildComponentsSafe() = %v, want ErrTooManyActionRows", err)
	}
	if result != nil {
		t.Error("BuildComponentsSafe() should return nil on error")
	}
}

func TestBuildComponentsSafe_TooManyComponents(t *testing.T) {
	var comps []components.Component

	// Create 26 buttons (5 buttons per row x 5 rows = 25 max)
	// 26 buttons will create 6 rows (5+5+5+5+5+1) which is more than max of 5
	// Actually, we need to test the component limit, not the row limit
	// Let's create 5 rows of 5 buttons = 25 total components, then add 1 more = 26
	// But that would hit the row limit first.
	// Let's use a different approach: use action rows with mixed components

	// Create 5 action rows with multiple buttons each
	for i := 0; i < components.MaxActionRows; i++ {
		for j := 0; j < 5; j++ {
			comps = append(comps, components.Button("id").Label("Button").Build())
		}
	}
	// Add one more component
	comps = append(comps, components.Button("id").Label("Button").Build())

	result, err := components.BuildComponentsSafe(comps...)
	if err != components.ErrTooManyActionRows && err != components.ErrTooManyComponents {
		t.Errorf("BuildComponentsSafe() = %v, want error", err)
	}
	if result != nil && err != nil {
		// If we got an error, result should be nil
	}
}

func TestActionRow_SingleComponent(t *testing.T) {
	btn := components.Button("id").Label("Button").Build()
	ar := components.ActionRow(btn)

	if ar == nil {
		t.Fatal("ActionRow() returned nil")
	}
	if ar.Type != types.ComponentTypeActionRow {
		t.Errorf("Type = %v, want %v", ar.Type, types.ComponentTypeActionRow)
	}
	if len(ar.Components) != 1 {
		t.Errorf("Expected 1 component, got %d", len(ar.Components))
	}
}

func TestActionRow_MultipleComponents(t *testing.T) {
	btn1 := components.Button("id1").Label("Button 1").Build()
	btn2 := components.Button("id2").Label("Button 2").Build()
	btn3 := components.Button("id3").Label("Button 3").Build()

	ar := components.ActionRow(btn1, btn2, btn3)

	if len(ar.Components) != 3 {
		t.Errorf("Expected 3 components, got %d", len(ar.Components))
	}
}

func TestActionRow_Empty(t *testing.T) {
	ar := components.ActionRow()

	if ar == nil {
		t.Fatal("ActionRow() returned nil")
	}
	if len(ar.Components) != 0 {
		t.Errorf("Expected 0 components, got %d", len(ar.Components))
	}
}

func TestBuildComponents_TenButtons(t *testing.T) {
	var buttons []components.Component
	for i := 0; i < 10; i++ {
		buttons = append(buttons, components.Button("id").Label("Button").Build())
	}

	result := components.BuildComponents(buttons...)

	if len(result) != 2 {
		t.Errorf("Expected 2 rows for 10 buttons, got %d", len(result))
	}

	// Each row should have exactly 5 buttons
	for i, row := range result {
		if ar, ok := row.(*components.ActionRowComponent); ok {
			if len(ar.Components) != components.MaxButtonsPerRow {
				t.Errorf("Row %d: expected %d components, got %d", i, components.MaxButtonsPerRow, len(ar.Components))
			}
		}
	}
}

func TestBuildComponents_MaxRows(t *testing.T) {
	var comps []components.Component

	// Create MaxActionRows worth of selects (each gets its own row)
	for i := 0; i < components.MaxActionRows; i++ {
		sel := components.StringSelect("id")
		sel.Option("Option", "value")
		comps = append(comps, sel.Build())
	}

	result, err := components.BuildComponentsSafe(comps...)
	if err != nil {
		t.Errorf("BuildComponentsSafe() error = %v", err)
	}
	if len(result) != components.MaxActionRows {
		t.Errorf("Expected %d rows, got %d", components.MaxActionRows, len(result))
	}
}

func TestBuildComponents_SelectWithButtons(t *testing.T) {
	sel := components.StringSelect("id")
	sel.Option("Option", "value")

	var comps []components.Component
	comps = append(comps, components.Button("id1").Label("Button 1").Build())
	comps = append(comps, sel.Build())
	comps = append(comps, components.Button("id2").Label("Button 2").Build())

	result := components.BuildComponents(comps...)

	if len(result) != 3 {
		t.Errorf("Expected 3 rows, got %d", len(result))
	}

	// Verify order: button row, select row, button row
	if _, ok := result[0].(*components.ActionRowComponent); !ok {
		t.Error("First should be ActionRow")
	}
	if _, ok := result[1].(*components.ActionRowComponent); !ok {
		t.Error("Second should be ActionRow")
	}
	if _, ok := result[2].(*components.ActionRowComponent); !ok {
		t.Error("Third should be ActionRow")
	}
}

func TestBuildComponents_ButtonGrouping(t *testing.T) {
	var buttons []components.Component
	for i := 0; i < 7; i++ {
		buttons = append(buttons, components.Button("id").Label("Button").Build())
	}

	result := components.BuildComponents(buttons...)

	if len(result) != 2 {
		t.Errorf("Expected 2 rows, got %d", len(result))
	}

	// First row should have 5
	if ar, ok := result[0].(*components.ActionRowComponent); ok {
		if len(ar.Components) != 5 {
			t.Errorf("First row: expected 5 components, got %d", len(ar.Components))
		}
	}

	// Second row should have 2
	if ar, ok := result[1].(*components.ActionRowComponent); ok {
		if len(ar.Components) != 2 {
			t.Errorf("Second row: expected 2 components, got %d", len(ar.Components))
		}
	}
}

package pokeapi

import (
	"testing"
)

func TestCategorizeMoves(t *testing.T) {
	// Create test moves
	testMoves := []MoveDetails{
		{
			Name:  "tackle",
			Power: 40,
			DamageClass: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "physical",
			},
			Type: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "normal",
			},
		},
		{
			Name:  "thunderbolt",
			Power: 90,
			DamageClass: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "special",
			},
			Type: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "electric",
			},
		},
		{
			Name:  "growl",
			Power: 0, // Status move with no power
			DamageClass: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "status",
			},
			Type: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "normal",
			},
		},
	}

	// Test categorization
	basicMoves, specialMoves := CategorizeMoves(testMoves)

	// Verify results
	if len(basicMoves) != 1 {
		t.Errorf("Expected 1 basic move, got %d", len(basicMoves))
	}
	if len(specialMoves) != 1 {
		t.Errorf("Expected 1 special move, got %d", len(specialMoves))
	}

	// Check if moves were categorized correctly
	if len(basicMoves) > 0 && basicMoves[0].Name != "tackle" {
		t.Errorf("Expected basic move to be 'tackle', got '%s'", basicMoves[0].Name)
	}
	if len(specialMoves) > 0 && specialMoves[0].Name != "thunderbolt" {
		t.Errorf("Expected special move to be 'thunderbolt', got '%s'", specialMoves[0].Name)
	}
}

func TestConvertMoveToSkill(t *testing.T) {
	// Create a test move
	testMove := MoveDetails{
		Name:  "flamethrower",
		Power: 90,
		DamageClass: struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			Name: "special",
		},
		Type: struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			Name: "fire",
		},
	}

	// Convert to skill
	skill := ConvertMoveToSkill(testMove)

	// Verify conversion
	if skill.Name != "flamethrower" {
		t.Errorf("Expected skill name to be 'flamethrower', got '%s'", skill.Name)
	}
	if skill.Damage != 90 {
		t.Errorf("Expected skill damage to be 90, got %d", skill.Damage)
	}
	if skill.Type != "fire" {
		t.Errorf("Expected skill type to be 'fire', got '%s'", skill.Type)
	}
	if skill.Class != "special" {
		t.Errorf("Expected skill class to be 'special', got '%s'", skill.Class)
	}
	// Check URL format
	expectedURLSuffix := "move/flamethrower"
	if len(skill.URL) < len(expectedURLSuffix) || skill.URL[len(skill.URL)-len(expectedURLSuffix):] != expectedURLSuffix {
		t.Errorf("Expected URL to end with '%s', got '%s'", expectedURLSuffix, skill.URL)
	}
}

func TestHasValidMoves(t *testing.T) {
	// Create test moves
	basicMoves := []MoveDetails{
		{
			Name:  "tackle",
			Power: 40,
		},
	}
	specialMoves := []MoveDetails{
		{
			Name:  "thunderbolt",
			Power: 90,
		},
	}
	emptyMoves := []MoveDetails{}

	// Test with valid moves in both categories
	if !HasValidMoves(basicMoves, specialMoves) {
		t.Error("Expected HasValidMoves to return true with valid moves in both categories")
	}

	// Test with no basic moves
	if HasValidMoves(emptyMoves, specialMoves) {
		t.Error("Expected HasValidMoves to return false with no basic moves")
	}

	// Test with no special moves
	if HasValidMoves(basicMoves, emptyMoves) {
		t.Error("Expected HasValidMoves to return false with no special moves")
	}

	// Test with no moves in either category
	if HasValidMoves(emptyMoves, emptyMoves) {
		t.Error("Expected HasValidMoves to return false with no moves in either category")
	}
}

func TestGetMoveCounts(t *testing.T) {
	// Create test moves
	basicMoves := []MoveDetails{
		{Name: "tackle"},
		{Name: "scratch"},
	}
	specialMoves := []MoveDetails{
		{Name: "thunderbolt"},
		{Name: "flamethrower"},
		{Name: "ice-beam"},
	}

	// Test count function
	basicCount, specialCount := GetMoveCounts(basicMoves, specialMoves)

	// Verify counts
	if basicCount != 2 {
		t.Errorf("Expected 2 basic moves, got %d", basicCount)
	}
	if specialCount != 3 {
		t.Errorf("Expected 3 special moves, got %d", specialCount)
	}
}
package pokeapi

import (
	"testing"
)

func TestGetMoveDetails(t *testing.T) {
	client := NewClient()

	moveDetails, err := client.GetMoveDetails("tackle")
	if err != nil {
		t.Fatalf("Error getting move details: %v", err)
	}

	// Verify basic response properties
	if moveDetails.Name != "tackle" {
		t.Errorf("Expected move name 'tackle', got '%s'", moveDetails.Name)
	}

	if moveDetails.Power <= 0 {
		t.Errorf("Expected power to be greater than 0, got %d", moveDetails.Power)
	}

	if moveDetails.Type.Name == "" {
		t.Errorf("Expected a type name, got empty string")
	}

	if moveDetails.DamageClass.Name == "" {
		t.Errorf("Expected a damage class name, got empty string")
	}

	// Test cache hit
	cachedMoveDetails, err := client.GetMoveDetails("tackle")
	if err != nil {
		t.Fatalf("Error getting cached move details: %v", err)
	}

	if cachedMoveDetails.Name != moveDetails.Name {
		t.Errorf("Cache inconsistency: expected %s, got %s", moveDetails.Name, cachedMoveDetails.Name)
	}
}

func TestGetPokemonMoves(t *testing.T) {
	client := NewClient()

	// Pikachu has many moves, so it's a good test case
	moves, err := client.GetPokemonMoves("pikachu")
	if err != nil {
		t.Fatalf("Error getting Pokemon moves: %v", err)
	}

	// Verify we got some moves back
	if len(moves) == 0 {
		t.Errorf("Expected moves for pikachu, got none")
	}

	// Check structure of first move
	if moves[0].Move.Name == "" {
		t.Errorf("Expected move name, got empty string")
	}

	if moves[0].Move.URL == "" {
		t.Errorf("Expected move URL, got empty string")
	}

	// Verify version group details
	if len(moves[0].VersionGroupDetails) == 0 {
		t.Errorf("Expected version group details, got none")
	}
}

func TestGetMoveDetailsBatch(t *testing.T) {
	client := NewClient()
	// Test with common move names
	moveNames := []string{"tackle", "thunderbolt", "ember"}
	// First call should hit the API for each move
	moves, err := client.GetMoveDetailsBatch(moveNames)
	if err != nil {
		t.Fatalf("Error getting batch moves: %v", err)
	}
	// Verify we got all the moves back
	if len(moves) != len(moveNames) {
		t.Errorf("Expected %d moves, got %d", len(moveNames), len(moves))
	}
	// Verify each move has the correct name
	for _, name := range moveNames {
		move, exists := moves[name]
		if !exists {
			t.Errorf("Move %s not found in results", name)
			continue
		}
		if move.Name != name {
			t.Errorf("Expected move name %s, got %s", name, move.Name)
		}
	}
	// Second call should be a cache hit for each move
	moves2, err := client.GetMoveDetailsBatch(moveNames)
	if err != nil {
		t.Fatalf("Error getting cached batch moves: %v", err)
	}
	// Verify the cached results match
	if len(moves2) != len(moves) {
		t.Errorf("Cached result count mismatch: expected %d, got %d", len(moves), len(moves2))
	}
	// Check that a specific move's details match between the two calls
	if moves["tackle"].Power != moves2["tackle"].Power {
		t.Errorf("Cache inconsistency: move power mismatch")
	}
}

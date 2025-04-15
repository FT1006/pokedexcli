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
package pokeapi

import (
	"testing"
	
	"github.com/FT1006/pokedexcli/internal/models"
)

func TestSelectRandomBasicMove(t *testing.T) {
	// Test with empty list
	emptyMoves := []MoveDetails{}
	_, ok := SelectRandomBasicMove(emptyMoves)
	if ok {
		t.Error("Expected SelectRandomBasicMove to return false with empty moves list")
	}

	// Test with populated list
	basicMoves := []MoveDetails{
		{Name: "tackle", Power: 40},
		{Name: "scratch", Power: 40},
		{Name: "pound", Power: 40},
	}
	move, ok := SelectRandomBasicMove(basicMoves)
	if !ok {
		t.Error("Expected SelectRandomBasicMove to return true with valid moves list")
	}

	// Check if the selected move is in the original list
	found := false
	for _, m := range basicMoves {
		if m.Name == move.Name {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Selected move %s not found in the original list", move.Name)
	}
}

func TestSelectRandomSpecialMove(t *testing.T) {
	// Test with empty list
	emptyMoves := []MoveDetails{}
	_, ok := SelectRandomSpecialMove(emptyMoves)
	if ok {
		t.Error("Expected SelectRandomSpecialMove to return false with empty moves list")
	}

	// Test with populated list
	specialMoves := []MoveDetails{
		{Name: "thunderbolt", Power: 90},
		{Name: "flamethrower", Power: 90},
		{Name: "ice-beam", Power: 90},
	}
	move, ok := SelectRandomSpecialMove(specialMoves)
	if !ok {
		t.Error("Expected SelectRandomSpecialMove to return true with valid moves list")
	}

	// Check if the selected move is in the original list
	found := false
	for _, m := range specialMoves {
		if m.Name == move.Name {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Selected move %s not found in the original list", move.Name)
	}
}

func TestSelectMoves(t *testing.T) {
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
	}

	// Test move selection
	basicSkill, specialSkill, err := SelectMoves(testMoves)
	if err != nil {
		t.Errorf("Unexpected error in SelectMoves: %v", err)
	}

	// Verify basic skill
	if basicSkill.Name != "tackle" {
		t.Errorf("Expected basic skill name to be 'tackle', got '%s'", basicSkill.Name)
	}
	if basicSkill.Type != "normal" {
		t.Errorf("Expected basic skill type to be 'normal', got '%s'", basicSkill.Type)
	}
	if basicSkill.Class != "physical" {
		t.Errorf("Expected basic skill class to be 'physical', got '%s'", basicSkill.Class)
	}

	// Verify special skill
	if specialSkill.Name != "thunderbolt" {
		t.Errorf("Expected special skill name to be 'thunderbolt', got '%s'", specialSkill.Name)
	}
	if specialSkill.Type != "electric" {
		t.Errorf("Expected special skill type to be 'electric', got '%s'", specialSkill.Type)
	}
	if specialSkill.Class != "special" {
		t.Errorf("Expected special skill class to be 'special', got '%s'", specialSkill.Class)
	}
}

func TestSelectMovesEmptyCategories(t *testing.T) {
	// Test with no moves available
	emptyMoves := []MoveDetails{}
	basicSkill, specialSkill, err := SelectMoves(emptyMoves)
	if err != nil {
		t.Errorf("Unexpected error in SelectMoves with empty list: %v", err)
	}

	// Verify fallback to struggle move
	if basicSkill.Name != "struggle" {
		t.Errorf("Expected basic skill to fallback to 'struggle', got '%s'", basicSkill.Name)
	}
	if specialSkill.Name != "struggle" {
		t.Errorf("Expected special skill to fallback to 'struggle', got '%s'", specialSkill.Name)
	}

	// Test with moves only in one category
	onlySpeicalMoves := []MoveDetails{
		{
			Name:  "thunderbolt",
			Power: 90,
			DamageClass: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "special",
			},
		},
	}
	basicSkill, specialSkill, err = SelectMoves(onlySpeicalMoves)
	if err != nil {
		t.Errorf("Unexpected error in SelectMoves with one category: %v", err)
	}

	// Both should fallback to struggle since we need both categories
	if basicSkill.Name != "struggle" {
		t.Errorf("Expected basic skill to fallback to 'struggle' when only special moves available, got '%s'", basicSkill.Name)
	}
}

// MockClient implements a mock version of the Client interface for testing
type MockClient struct {
	MockGetPokemonDetails func(name string) (PokemonRes, error)
	MockGetMoveDetailsBatch func(moveNames []string) (map[string]MoveDetails, error)
}

func (m *MockClient) GetPokemonDetails(name string) (PokemonRes, error) {
	if m.MockGetPokemonDetails != nil {
		return m.MockGetPokemonDetails(name)
	}
	return PokemonRes{}, nil
}

func (m *MockClient) GetMoveDetailsBatch(moveNames []string) (map[string]MoveDetails, error) {
	if m.MockGetMoveDetailsBatch != nil {
		return m.MockGetMoveDetailsBatch(moveNames)
	}
	return map[string]MoveDetails{}, nil
}

// Implement the GetMovesForPokemon method for MockClient
func (m *MockClient) GetMovesForPokemon(name string) (models.Skill, models.Skill, error) {
	// Get Pokemon details
	pokemonDetails, err := m.GetPokemonDetails(name)
	if err != nil {
		return models.Skill{}, models.Skill{}, err
	}

	// Extract move references
	var moveRefs []string
	for _, move := range pokemonDetails.Moves {
		moveRefs = append(moveRefs, move.Move.Name)
	}

	// Fetch move details
	moveDetailsMap, err := m.GetMoveDetailsBatch(moveRefs)
	if err != nil {
		return models.Skill{}, models.Skill{}, err
	}

	// Convert map to slice
	var moveDetails []MoveDetails
	for _, details := range moveDetailsMap {
		moveDetails = append(moveDetails, details)
	}

	// Select the moves using the quality-based selection
	return SelectQualityMoves(moveDetails)
}

func TestGetMovesForPokemon(t *testing.T) {
	// Create a mock client
	mockClient := &MockClient{}
	
	// Set up the mock to return predefined responses
	mockClient.MockGetPokemonDetails = func(name string) (PokemonRes, error) {
		// Return a Pokemon with a set of moves
		return PokemonRes{
			Name: "pikachu",
			Moves: []PokemonMove{
				{
					Move: struct {
						Name string `json:"name"`
						URL  string `json:"url"`
					}{
						Name: "tackle",
					},
				},
				{
					Move: struct {
						Name string `json:"name"`
						URL  string `json:"url"`
					}{
						Name: "thunderbolt",
					},
				},
			},
		}, nil
	}
	
	mockClient.MockGetMoveDetailsBatch = func(moveNames []string) (map[string]MoveDetails, error) {
		// Return move details for the provided move names
		result := make(map[string]MoveDetails)
		
		for _, name := range moveNames {
			if name == "tackle" {
				result[name] = MoveDetails{
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
				}
			} else if name == "thunderbolt" {
				result[name] = MoveDetails{
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
				}
			}
		}
		
		return result, nil
	}
	
	// Test the client.GetMovesForPokemon method
	basicSkill, specialSkill, err := mockClient.GetMovesForPokemon("pikachu")
	if err != nil {
		t.Errorf("Unexpected error in GetMovesForPokemon: %v", err)
	}
	
	// Verify the selected skills
	if basicSkill.Name != "tackle" || basicSkill.Class != "physical" {
		t.Errorf("Expected basic skill to be 'tackle' with class 'physical', got '%s' with class '%s'", 
			basicSkill.Name, basicSkill.Class)
	}
	
	if specialSkill.Name != "thunderbolt" || specialSkill.Class != "special" {
		t.Errorf("Expected special skill to be 'thunderbolt' with class 'special', got '%s' with class '%s'", 
			specialSkill.Name, specialSkill.Class)
	}
}
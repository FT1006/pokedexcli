package service

import (
	"testing"

	"github.com/FT1006/pokedexcli/internal/database"
	"github.com/FT1006/pokedexcli/internal/models"
)

// MockDB implements a mock database service for testing
type MockDB struct{}

func (m *MockDB) Queries() interface{} {
	return &MockQueries{}
}

// MockQueries implements mock database queries
type MockQueries struct{}

// Test the ConvertToOwnedPokemon function with skills
func TestConvertToOwnedPokemonWithSkills(t *testing.T) {
	// Create a PokemonService with a mock DB
	service := NewPokemonService(&database.Service{})

	// Create test data
	basicSkill := &models.Skill{
		Name:   "tackle",
		URL:    "https://pokeapi.co/api/v2/move/tackle",
		Damage: 40,
		Type:   "normal",
		Class:  "physical",
	}

	specialSkill := &models.Skill{
		Name:   "flamethrower",
		URL:    "https://pokeapi.co/api/v2/move/flamethrower",
		Damage: 90,
		Type:   "fire",
		Class:  "special",
	}

	pokemon := models.Pokemon{
		Name:           "charmander",
		Height:         6,
		Weight:         85,
		BaseExperience: 62,
		Stats:          []models.Stats{},
		Types:          []models.Types{},
		BasicSkill:     basicSkill,
		SpecialSkill:   specialSkill,
	}

	// Convert to DB model
	result, err := service.ConvertToOwnedPokemon(1, pokemon)
	if err != nil {
		t.Fatalf("Error converting Pokemon to owned Pokemon: %v", err)
	}

	// Verify the basic skill was converted correctly
	if len(result.BasicSkill) == 0 {
		t.Error("Expected BasicSkill to be marshaled, but got empty result")
	}

	// Verify the special skill was converted correctly
	if len(result.SpecialSkill) == 0 {
		t.Error("Expected SpecialSkill to be marshaled, but got empty result")
	}
}

// Test handling of nil skills in ConvertToOwnedPokemon
func TestConvertToOwnedPokemonWithNilSkills(t *testing.T) {
	// Create a PokemonService with a mock DB
	service := NewPokemonService(&database.Service{})

	// Create test data with nil skills
	pokemon := models.Pokemon{
		Name:           "charmander",
		Height:         6,
		Weight:         85,
		BaseExperience: 62,
		Stats:          []models.Stats{},
		Types:          []models.Types{},
		BasicSkill:     nil,
		SpecialSkill:   nil,
	}

	// Convert to DB model
	result, err := service.ConvertToOwnedPokemon(1, pokemon)
	if err != nil {
		t.Fatalf("Error converting Pokemon to owned Pokemon: %v", err)
	}

	// Verify the skills are nil/empty
	if len(result.BasicSkill) != 0 {
		t.Error("Expected BasicSkill to be empty, but got non-empty result")
	}

	if len(result.SpecialSkill) != 0 {
		t.Error("Expected SpecialSkill to be empty, but got non-empty result")
	}
}

// Test UnmarshalSkill function
func TestUnmarshalSkill(t *testing.T) {
	// Create a PokemonService with a mock DB
	service := NewPokemonService(&database.Service{})

	// Create test skill
	skill := models.Skill{
		Name:   "tackle",
		URL:    "https://pokeapi.co/api/v2/move/tackle",
		Damage: 40,
		Type:   "normal",
		Class:  "physical",
	}

	// Marshal to JSON
	pokemon := models.Pokemon{
		BasicSkill: &skill,
	}
	result, err := service.ConvertToOwnedPokemon(1, pokemon)
	if err != nil {
		t.Fatalf("Error converting Pokemon to owned Pokemon: %v", err)
	}

	// Unmarshal using UnmarshalSkill
	unmarshaledSkill, err := service.UnmarshalSkill(result.BasicSkill)
	if err != nil {
		t.Fatalf("Error unmarshaling skill: %v", err)
	}

	// Verify the unmarshaled skill
	if unmarshaledSkill == nil {
		t.Fatal("Expected non-nil skill, but got nil")
	}

	if unmarshaledSkill.Name != skill.Name {
		t.Errorf("Expected skill name %s, but got %s", skill.Name, unmarshaledSkill.Name)
	}

	if unmarshaledSkill.Damage != skill.Damage {
		t.Errorf("Expected skill damage %d, but got %d", skill.Damage, unmarshaledSkill.Damage)
	}

	if unmarshaledSkill.Type != skill.Type {
		t.Errorf("Expected skill type %s, but got %s", skill.Type, unmarshaledSkill.Type)
	}

	if unmarshaledSkill.Class != skill.Class {
		t.Errorf("Expected skill class %s, but got %s", skill.Class, unmarshaledSkill.Class)
	}
}
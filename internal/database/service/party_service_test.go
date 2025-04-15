package service

import (
	"testing"

	"github.com/FT1006/pokedexcli/internal/models"
)

// Test that the PartyPokemon struct can properly hold skill data
func TestPartyPokemonWithSkills(t *testing.T) {
	// Create test skills
	basicSkill := &models.Skill{
		Name:   "tackle",
		URL:    "https://pokeapi.co/api/v2/move/tackle",
		Damage: 40,
		Type:   "normal",
		Class:  "physical",
	}

	specialSkill := &models.Skill{
		Name:   "thunderbolt",
		URL:    "https://pokeapi.co/api/v2/move/thunderbolt",
		Damage: 90,
		Type:   "electric",
		Class:  "special",
	}

	// Create a PartyPokemon with skills
	partyPokemon := PartyPokemon{
		Slot:           1,
		OwnpokeID:      123,
		Name:           "pikachu",
		Height:         4,
		Weight:         60,
		Stats:          []models.Stats{},
		Types:          []models.Types{},
		BaseExperience: 112,
		BasicSkill:     basicSkill,
		SpecialSkill:   specialSkill,
	}

	// Test that skills are properly stored and accessible
	if partyPokemon.BasicSkill == nil {
		t.Error("Expected BasicSkill to be non-nil")
	} else {
		if partyPokemon.BasicSkill.Name != "tackle" {
			t.Errorf("Expected BasicSkill name to be 'tackle', got '%s'", partyPokemon.BasicSkill.Name)
		}
		if partyPokemon.BasicSkill.Type != "normal" {
			t.Errorf("Expected BasicSkill type to be 'normal', got '%s'", partyPokemon.BasicSkill.Type)
		}
		if partyPokemon.BasicSkill.Class != "physical" {
			t.Errorf("Expected BasicSkill class to be 'physical', got '%s'", partyPokemon.BasicSkill.Class)
		}
	}

	if partyPokemon.SpecialSkill == nil {
		t.Error("Expected SpecialSkill to be non-nil")
	} else {
		if partyPokemon.SpecialSkill.Name != "thunderbolt" {
			t.Errorf("Expected SpecialSkill name to be 'thunderbolt', got '%s'", partyPokemon.SpecialSkill.Name)
		}
		if partyPokemon.SpecialSkill.Type != "electric" {
			t.Errorf("Expected SpecialSkill type to be 'electric', got '%s'", partyPokemon.SpecialSkill.Type)
		}
		if partyPokemon.SpecialSkill.Class != "special" {
			t.Errorf("Expected SpecialSkill class to be 'special', got '%s'", partyPokemon.SpecialSkill.Class)
		}
	}
}
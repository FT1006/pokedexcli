package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/FT1006/pokedexcli/internal/database"
	dbsqlc "github.com/FT1006/pokedexcli/internal/database/sqlc/db"
	"github.com/FT1006/pokedexcli/internal/models"
)

// APIStruct aliases to shared model types
// We use these to avoid naming conflicts with the generated DB types
type APIPokemon = models.Pokemon
type APIStat = models.Stat
type APIStats = models.Stats
type APIType = models.Type
type APITypes = models.Types
type APISkill = models.Skill

// OwnedPokemon represents a caught pokemon with timestamp
type OwnedPokemon struct {
	ID           int32     // Database ID
	Name         string
	CaughtAt     time.Time
	BasicSkill   *APISkill
	SpecialSkill *APISkill
}

type PokemonService struct {
	db *database.Service
}

func NewPokemonService(db *database.Service) *PokemonService {
	return &PokemonService{
		db: db,
	}
}

// Convert API Pokemon to DB Pokedex entry
func (s *PokemonService) ConvertToPokedex(trainerID int32, p APIPokemon) (dbsqlc.CreatePokedexEntryParams, error) {
	// Convert stats to JSON
	statsJSON, err := json.Marshal(p.Stats)
	if err != nil {
		return dbsqlc.CreatePokedexEntryParams{}, fmt.Errorf("error marshaling stats: %w", err)
	}

	// Convert types to JSON
	typesJSON, err := json.Marshal(p.Types)
	if err != nil {
		return dbsqlc.CreatePokedexEntryParams{}, fmt.Errorf("error marshaling types: %w", err)
	}

	return dbsqlc.CreatePokedexEntryParams{
		TrainerID:      trainerID,
		Name:           p.Name,
		Height:         int32(p.Height),
		Weight:         int32(p.Weight),
		BaseExperience: int32(p.BaseExperience),
		Stats:          statsJSON,
		Types:          typesJSON,
	}, nil
}

// Convert API Pokemon to DB Owned Pokemon entry
func (s *PokemonService) ConvertToOwnedPokemon(trainerID int32, p APIPokemon) (dbsqlc.AddOwnedPokemonParams, error) {
	// Convert stats to JSON
	statsJSON, err := json.Marshal(p.Stats)
	if err != nil {
		return dbsqlc.AddOwnedPokemonParams{}, fmt.Errorf("error marshaling stats: %w", err)
	}

	// Convert types to JSON
	typesJSON, err := json.Marshal(p.Types)
	if err != nil {
		return dbsqlc.AddOwnedPokemonParams{}, fmt.Errorf("error marshaling types: %w", err)
	}

	// Convert skills to JSON (handle nil gracefully)
	var basicSkillJSON, specialSkillJSON []byte

	if p.BasicSkill != nil {
		basicSkillJSON, err = json.Marshal(p.BasicSkill)
		if err != nil {
			return dbsqlc.AddOwnedPokemonParams{}, fmt.Errorf("error marshaling basic skill: %w", err)
		}
	}

	if p.SpecialSkill != nil {
		specialSkillJSON, err = json.Marshal(p.SpecialSkill)
		if err != nil {
			return dbsqlc.AddOwnedPokemonParams{}, fmt.Errorf("error marshaling special skill: %w", err)
		}
	}

	return dbsqlc.AddOwnedPokemonParams{
		TrainerID:      trainerID,
		Name:           p.Name,
		Height:         int32(p.Height),
		Weight:         int32(p.Weight),
		BaseExperience: int32(p.BaseExperience),
		Stats:          statsJSON,
		Types:          typesJSON,
		BasicSkill:     basicSkillJSON,
		SpecialSkill:   specialSkillJSON,
	}, nil
}

// Helper function to unmarshal stats and types JSON
func (s *PokemonService) UnmarshalStatsAndTypes(statsJSON, typesJSON []byte) ([]APIStats, []APITypes, error) {
	var stats []APIStats
	var types []APITypes

	// Unmarshal stats
	if err := json.Unmarshal(statsJSON, &stats); err != nil {
		return nil, nil, fmt.Errorf("error unmarshaling stats: %w", err)
	}

	// Unmarshal types
	if err := json.Unmarshal(typesJSON, &types); err != nil {
		return nil, nil, fmt.Errorf("error unmarshaling types: %w", err)
	}

	return stats, types, nil
}

// Convert DB Pokedex to API Pokemon
func (s *PokemonService) ConvertFromPokedex(p dbsqlc.Pokedex) (APIPokemon, error) {
	stats, types, err := s.UnmarshalStatsAndTypes(p.Stats, p.Types)
	if err != nil {
		return APIPokemon{}, err
	}

	return APIPokemon{
		Name:           p.Name,
		Height:         int(p.Height),
		Weight:         int(p.Weight),
		Stats:          stats,
		Types:          types,
		BaseExperience: int(p.BaseExperience),
	}, nil
}

// Helper function to unmarshal skill JSON
func (s *PokemonService) UnmarshalSkill(skillJSON []byte) (*APISkill, error) {
	if len(skillJSON) == 0 {
		return nil, nil // No skill data available
	}

	var skill APISkill
	if err := json.Unmarshal(skillJSON, &skill); err != nil {
		return nil, fmt.Errorf("error unmarshaling skill: %w", err)
	}

	return &skill, nil
}

// Convert DB OwnPoke to API Pokemon with caught time
func (s *PokemonService) ConvertFromOwnPoke(p dbsqlc.Ownpoke) (OwnedPokemon, error) {
	var basicSkill, specialSkill *APISkill
	var err error

	// Unmarshal skills if present
	if len(p.BasicSkill) > 0 {
		basicSkill, err = s.UnmarshalSkill(p.BasicSkill)
		if err != nil {
			return OwnedPokemon{}, fmt.Errorf("error with basic skill: %w", err)
		}
	}

	if len(p.SpecialSkill) > 0 {
		specialSkill, err = s.UnmarshalSkill(p.SpecialSkill)
		if err != nil {
			return OwnedPokemon{}, fmt.Errorf("error with special skill: %w", err)
		}
	}

	return OwnedPokemon{
		ID:           p.ID,
		Name:         p.Name,
		CaughtAt:     p.CaughtAt.Time,
		BasicSkill:   basicSkill,
		SpecialSkill: specialSkill,
	}, nil
}

// Save Pokemon to pokedex and ownpoke for a trainer (used when catching)
func (s *PokemonService) SavePokemon(ctx context.Context, trainerID int32, pokemon APIPokemon) error {
	// Add to pokedex (non-duplicated)
	err := s.AddToPokedex(ctx, trainerID, pokemon)
	if err != nil {
		return err
	}

	// Add to ownpoke (allows duplicates)
	err = s.AddToOwnPoke(ctx, trainerID, pokemon)
	if err != nil {
		return err
	}

	return nil
}

// AddToPokedex adds a Pokemon to the pokedex (no duplicates)
func (s *PokemonService) AddToPokedex(ctx context.Context, trainerID int32, pokemon APIPokemon) error {
	pokedexEntry, err := s.ConvertToPokedex(trainerID, pokemon)
	if err != nil {
		return err
	}

	err = s.db.Queries().CreatePokedexEntry(ctx, pokedexEntry)
	if err != nil {
		return fmt.Errorf("error creating pokedex entry: %w", err)
	}

	return nil
}

// AddToOwnPoke adds a Pokemon to the ownpoke table (allows duplicates)
func (s *PokemonService) AddToOwnPoke(ctx context.Context, trainerID int32, pokemon APIPokemon) error {
	ownedPokemon, err := s.ConvertToOwnedPokemon(trainerID, pokemon)
	if err != nil {
		return err
	}

	_, err = s.db.Queries().AddOwnedPokemon(ctx, ownedPokemon)
	if err != nil {
		return fmt.Errorf("error adding owned pokemon: %w", err)
	}

	return nil
}

// Get Pokemon by name from pokedex for a trainer
func (s *PokemonService) GetPokemonByName(ctx context.Context, trainerID int32, name string) (APIPokemon, error) {
	dbPokemon, err := s.db.Queries().GetPokedexEntryByNameAndTrainer(ctx, dbsqlc.GetPokedexEntryByNameAndTrainerParams{
		Name:      name,
		TrainerID: trainerID,
	})
	if err != nil {
		return APIPokemon{}, fmt.Errorf("error getting pokemon: %w", err)
	}

	return s.ConvertFromPokedex(dbPokemon)
}

// Get all Pokemon from pokedex for a trainer
func (s *PokemonService) GetAllPokemon(ctx context.Context, trainerID int32) ([]APIPokemon, error) {
	dbPokemons, err := s.db.Queries().ListPokedexByTrainer(ctx, trainerID)
	if err != nil {
		return nil, fmt.Errorf("error listing pokemon: %w", err)
	}

	pokemons := make([]APIPokemon, 0, len(dbPokemons))
	for _, dbPokemon := range dbPokemons {
		pokemon, err := s.ConvertFromPokedex(dbPokemon)
		if err != nil {
			return nil, err
		}
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

// Get all owned Pokemon for a trainer
func (s *PokemonService) GetAllOwnedPokemon(ctx context.Context, trainerID int32) ([]OwnedPokemon, error) {
	dbPokemons, err := s.db.Queries().ListOwnedPokemonByTrainer(ctx, trainerID)
	if err != nil {
		return nil, fmt.Errorf("error listing owned pokemon: %w", err)
	}

	ownedPokemons := make([]OwnedPokemon, 0, len(dbPokemons))
	for _, dbPokemon := range dbPokemons {
		ownedPokemon, err := s.ConvertFromOwnPoke(dbPokemon)
		if err != nil {
			return nil, err
		}
		ownedPokemons = append(ownedPokemons, ownedPokemon)
	}

	return ownedPokemons, nil
}

// UpdatePokemonSkills updates the skills for an owned Pokemon
func (s *PokemonService) UpdatePokemonSkills(ctx context.Context, pokemonID int32, basicSkill, specialSkill *APISkill) error {
	// Marshal skills to JSON (handle nil gracefully)
	var basicSkillJSON, specialSkillJSON []byte
	var err error

	if basicSkill != nil {
		basicSkillJSON, err = json.Marshal(basicSkill)
		if err != nil {
			return fmt.Errorf("error marshaling basic skill: %w", err)
		}
	}

	if specialSkill != nil {
		specialSkillJSON, err = json.Marshal(specialSkill)
		if err != nil {
			return fmt.Errorf("error marshaling special skill: %w", err)
		}
	}

	// Update the Pokemon's skills in the database
	err = s.db.Queries().UpdateOwnedPokemonSkills(ctx, dbsqlc.UpdateOwnedPokemonSkillsParams{
		ID:           pokemonID,
		BasicSkill:   basicSkillJSON,
		SpecialSkill: specialSkillJSON,
	})

	if err != nil {
		return fmt.Errorf("error updating pokemon skills: %w", err)
	}

	return nil
}

// AddPokemonWithSkills adds a Pokemon to both pokedex and ownpoke with skills
func (s *PokemonService) AddPokemonWithSkills(ctx context.Context, trainerID int32, pokemon APIPokemon, basicSkill, specialSkill *APISkill) (int32, error) {
	// Set the skills on the Pokemon
	pokemonWithSkills := pokemon
	pokemonWithSkills.BasicSkill = basicSkill
	pokemonWithSkills.SpecialSkill = specialSkill

	// First, add to pokedex (or update if exists)
	// This ensures the Pokemon is in the trainer's Pokedex for the inspect command
	pokedexEntry, err := s.ConvertToPokedex(trainerID, pokemonWithSkills)
	if err != nil {
		return 0, fmt.Errorf("error creating pokedex entry: %w", err)
	}

	// Try to add to pokedex - this may fail if already exists, which is OK
	err = s.db.Queries().CreatePokedexEntry(ctx, pokedexEntry)
	if err != nil {
		// If error is not "already exists", return it
		if !isAlreadyExistsError(err) {
			return 0, fmt.Errorf("error adding to pokedex: %w", err)
		}
		// Otherwise continue - already having it in pokedex is fine
	}

	// Convert to DB model for ownpoke
	ownedPokemon, err := s.ConvertToOwnedPokemon(trainerID, pokemonWithSkills)
	if err != nil {
		return 0, err
	}

	// Add to the ownpoke table
	result, err := s.db.Queries().AddOwnedPokemon(ctx, ownedPokemon)
	if err != nil {
		return 0, fmt.Errorf("error adding owned pokemon with skills: %w", err)
	}

	return result.ID, nil
}

// Helper function to check if an error is "already exists"
func isAlreadyExistsError(err error) bool {
	return err != nil && (err.Error() == "ERROR: duplicate key value violates unique constraint \"pokedex_trainer_id_name_key\" (SQLSTATE 23505)" ||
		err.Error() == "ERROR: duplicate key value violates unique constraint \"pokedex_pkey\" (SQLSTATE 23505)")
}

// GetPokemonWithSkills gets a Pokemon with its skills by ID
func (s *PokemonService) GetPokemonWithSkills(ctx context.Context, pokemonID int32) (OwnedPokemon, error) {
	dbPokemon, err := s.db.Queries().GetOwnedPokemonByID(ctx, pokemonID)
	if err != nil {
		return OwnedPokemon{}, fmt.Errorf("error getting owned pokemon: %w", err)
	}

	return s.ConvertFromOwnPoke(dbPokemon)
}

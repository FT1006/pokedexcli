package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FT1006/pokedexcli/internal/models"
)

// APIStruct aliases to shared model types
// We use these to avoid naming conflicts with the generated DB types
type APIPokemon = models.Pokemon
type APIStat = models.Stat
type APIStats = models.Stats
type APIType = models.Type
type APITypes = models.Types

type PokemonService struct {
	db *Service
}

func NewPokemonService(db *Service) *PokemonService {
	return &PokemonService{
		db: db,
	}
}

// Convert API Pokemon to DB Pokemon
func (s *PokemonService) ConvertToDB(trainerID int32, p APIPokemon) (CreatePokemonParams, error) {
	// Convert stats to JSON
	statsJSON, err := json.Marshal(p.Stats)
	if err != nil {
		return CreatePokemonParams{}, fmt.Errorf("error marshaling stats: %w", err)
	}

	// Convert types to JSON
	typesJSON, err := json.Marshal(p.Types)
	if err != nil {
		return CreatePokemonParams{}, fmt.Errorf("error marshaling types: %w", err)
	}

	return CreatePokemonParams{
		TrainerID:      trainerID,
		Name:           p.Name,
		Height:         int32(p.Height),
		Weight:         int32(p.Weight),
		BaseExperience: int32(p.BaseExperience),
		Stats:          statsJSON,
		Types:          typesJSON,
	}, nil
}

// Convert DB Pokemon to API Pokemon
func (s *PokemonService) ConvertFromDB(p Pokemon) (APIPokemon, error) {
	var stats []APIStats
	var types []APITypes

	// Unmarshal stats
	if err := json.Unmarshal(p.Stats, &stats); err != nil {
		return APIPokemon{}, fmt.Errorf("error unmarshaling stats: %w", err)
	}

	// Unmarshal types
	if err := json.Unmarshal(p.Types, &types); err != nil {
		return APIPokemon{}, fmt.Errorf("error unmarshaling types: %w", err)
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

// Save Pokemon for a trainer
func (s *PokemonService) SavePokemon(ctx context.Context, trainerID int32, pokemon APIPokemon) error {
	dbPokemon, err := s.ConvertToDB(trainerID, pokemon)
	if err != nil {
		return err
	}

	_, err = s.db.Queries().CreatePokemon(ctx, dbPokemon)
	if err != nil {
		return fmt.Errorf("error creating pokemon: %w", err)
	}

	return nil
}

// Get Pokemon by name for a trainer
func (s *PokemonService) GetPokemonByName(ctx context.Context, trainerID int32, name string) (APIPokemon, error) {
	dbPokemon, err := s.db.Queries().GetPokemonByNameAndTrainer(ctx, GetPokemonByNameAndTrainerParams{
		Name:      name,
		TrainerID: trainerID,
	})
	if err != nil {
		return APIPokemon{}, fmt.Errorf("error getting pokemon: %w", err)
	}

	return s.ConvertFromDB(dbPokemon)
}

// Get all Pokemon for a trainer
func (s *PokemonService) GetAllPokemon(ctx context.Context, trainerID int32) ([]APIPokemon, error) {
	dbPokemons, err := s.db.Queries().ListPokemonByTrainer(ctx, trainerID)
	if err != nil {
		return nil, fmt.Errorf("error listing pokemon: %w", err)
	}

	pokemons := make([]APIPokemon, 0, len(dbPokemons))
	for _, dbPokemon := range dbPokemons {
		pokemon, err := s.ConvertFromDB(dbPokemon)
		if err != nil {
			return nil, err
		}
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}
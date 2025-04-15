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

// OwnedPokemon represents a caught pokemon with timestamp
type OwnedPokemon struct {
	Name     string
	CaughtAt time.Time
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

	return dbsqlc.AddOwnedPokemonParams{
		TrainerID:      trainerID,
		Name:           p.Name,
		Height:         int32(p.Height),
		Weight:         int32(p.Weight),
		BaseExperience: int32(p.BaseExperience),
		Stats:          statsJSON,
		Types:          typesJSON,
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

// Convert DB OwnPoke to API Pokemon with caught time
func (s *PokemonService) ConvertFromOwnPoke(p dbsqlc.Ownpoke) (OwnedPokemon, error) {
	return OwnedPokemon{
		Name:     p.Name,
		CaughtAt: p.CaughtAt.Time,
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
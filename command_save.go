package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/FT1006/pokedexcli/internal/models"
)

func commandSave(name string, cfg *Config) error {
	// Check if database service is available
	if cfg.dbService == nil {
		return errors.New("database is not available, cannot save")
	}

	// If no name is provided and no current trainer, prompt for one
	if name == "" && cfg.currentTrainer == nil {
		return errors.New("please provide a trainer name to save as")
	}

	// Use current trainer name if none provided
	if name == "" && cfg.currentTrainer != nil {
		name = cfg.currentTrainer.Name.String
	}

	// Create or update trainer
	ctx := context.Background()
	trainer, err := cfg.trainerService.CreateOrUpdateTrainer(ctx, name)
	if err != nil {
		return fmt.Errorf("error creating trainer: %w", err)
	}

	// Update current trainer
	cfg.currentTrainer = &trainer

	// Only add newly caught Pokemon to Pokedex (no duplicates)
	// We don't add to ownpoke here because they're already added when caught
	if len(cfg.newlyCaughtPokemon) > 0 {
		for _, pokemon := range cfg.newlyCaughtPokemon {
			err := cfg.pokemonService.AddToPokedex(ctx, trainer.ID, pokemon)
			if err != nil {
				return fmt.Errorf("error adding pokemon to pokedex: %w", err)
			}
		}
		// Clear the newly caught Pokemon map after saving
		cfg.newlyCaughtPokemon = make(map[string]models.Pokemon)
	}

	// Get counts for Pokedex and owned Pokemon
	pokedexPokemon, err := cfg.pokemonService.GetAllPokemon(ctx, trainer.ID)
	if err != nil {
		return fmt.Errorf("error getting pokedex count: %w", err)
	}

	ownedPokemon, err := cfg.pokemonService.GetAllOwnedPokemon(ctx, trainer.ID)
	if err != nil {
		return fmt.Errorf("error getting owned pokemon count: %w", err)
	}

	fmt.Printf("Saved trainer '%s' with %d Pokemon owned, %d unique Pokemon in Pokedex\n",
		name, len(ownedPokemon), len(pokedexPokemon))
	return nil
}

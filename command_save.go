package main

import (
	"context"
	"errors"
	"fmt"
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

	// Save caught Pokemon
	for _, pokemon := range cfg.caughtPokemon {
		err := cfg.pokemonService.SavePokemon(ctx, trainer.ID, pokemon)
		if err != nil {
			return fmt.Errorf("error saving pokemon: %w", err)
		}
	}

	fmt.Printf("Saved trainer '%s' with %d Pokemon\n", name, len(cfg.caughtPokemon))
	return nil
}
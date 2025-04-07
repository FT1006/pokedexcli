package main

import (
	"context"
	"errors"
	"fmt"
	
	"github.com/FT1006/pokedexcli/internal/models"
)

func commandLoad(name string, cfg *Config) error {
	// Check if database service is available
	if cfg.dbService == nil {
		return errors.New("database is not available, cannot load")
	}

	// Name is required
	if name == "" {
		return errors.New("please provide a trainer name to load")
	}

	// Find trainer by name
	ctx := context.Background()
	trainer, err := cfg.trainerService.CreateOrUpdateTrainer(ctx, name)
	if err != nil {
		return fmt.Errorf("error finding trainer: %w", err)
	}

	// Update current trainer
	cfg.currentTrainer = &trainer

	// Load caught Pokemon
	pokemons, err := cfg.pokemonService.GetAllPokemon(ctx, trainer.ID)
	if err != nil {
		return fmt.Errorf("error loading pokemon: %w", err)
	}

	// Replace caught Pokemon map
	cfg.caughtPokemon = make(map[string]models.Pokemon)
	for _, pokemon := range pokemons {
		cfg.caughtPokemon[pokemon.Name] = pokemon
	}

	fmt.Printf("Loaded trainer '%s' with %d Pokemon\n", name, len(cfg.caughtPokemon))
	return nil
}
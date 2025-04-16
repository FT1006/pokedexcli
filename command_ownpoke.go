package main

import (
	"context"
	"fmt"
)

func commandOwnPoke(additionalInput string, c *Config) error {
	if c.dbService == nil {
		return fmt.Errorf("database not initialized")
	}

	if c.currentTrainer == nil {
		return fmt.Errorf("no trainer logged in - use 'load' command first")
	}

	ctx := context.Background()
	ownedPokemon, err := c.pokemonService.GetAllOwnedPokemon(ctx, c.currentTrainer.ID)
	if err != nil {
		return fmt.Errorf("error getting owned pokemon: %w", err)
	}

	if len(ownedPokemon) == 0 {
		fmt.Println("You don't have any caught Pokemon yet!")
		return nil
	}

	fmt.Printf("Your owned Pokemon collection (%d total):\n", len(ownedPokemon))

	// Display all Pokemon with their skills using our shared helper
	displayPokemonInstances(ownedPokemon)

	return nil
}

package main

import (
	"context"
	"fmt"
)

func commandPokedex(additionalInput string, c *Config) error {
	// Use database if available
	if c.dbService != nil && c.currentTrainer != nil {
		ctx := context.Background()
		pokemon, err := c.pokemonService.GetAllPokemon(ctx, c.currentTrainer.ID)
		if err != nil {
			return fmt.Errorf("error getting pokemon from database: %w", err)
		}

		if len(pokemon) == 0 {
			fmt.Println("Your Pokedex is empty! Go catch some Pokemon!")
			return nil
		}

		fmt.Println("Your Pokedex:")
		for _, p := range pokemon {
			fmt.Printf("- %s\n", p.Name)
		}
		return nil
	}

	// Fall back to in-memory storage if no database
	if len(c.caughtPokemon) == 0 {
		fmt.Println("You have not caught any Pokemon yet")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range c.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}
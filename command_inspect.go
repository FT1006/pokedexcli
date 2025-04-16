package main

import (
	"context"
	"fmt"

	"github.com/FT1006/pokedexcli/internal/database/service"
	"github.com/FT1006/pokedexcli/internal/models"
)

func commandInspect(pokemon string, c *Config) error {
	if pokemon == "" {
		fmt.Println("no pokemon provided")
		return nil
	}

	// Check if we have a database connection and a trainer loaded
	if c.dbService == nil {
		fmt.Println("Database not available. Running in offline mode.")
		// Check if we have this pokemon in memory
		if pkm, ok := c.caughtPokemon[pokemon]; !ok {
			fmt.Println("you have not caught that pokemon")
			return nil
		} else {
			displayPokemonDetails(pkm)
			fmt.Println("\nNote: Skills not available in offline mode.")
			return nil
		}
	}

	if c.currentTrainer == nil {
		fmt.Println("No trainer profile loaded. Use 'load <name>' to load or create a trainer profile.")
		fmt.Println("For now, showing basic info only:")

		// Check if we have this pokemon in memory
		if pkm, ok := c.caughtPokemon[pokemon]; !ok {
			fmt.Println("you have not caught that pokemon")
			return nil
		} else {
			displayPokemonDetails(pkm)
			fmt.Println("\nNote: Skill details require a loaded trainer profile.")
			return nil
		}
	}

	ctx := context.Background()

	// Get all owned Pokemon directly from ownpoke table
	ownedPokemon, err := c.pokemonService.GetAllOwnedPokemon(ctx, c.currentTrainer.ID)
	if err != nil {
		return fmt.Errorf("error getting owned pokemon: %w", err)
	}

	// Filter to just the requested Pokemon
	var instances []service.OwnedPokemon
	for _, p := range ownedPokemon {
		if p.Name == pokemon {
			instances = append(instances, p)
		}
	}

	if len(instances) == 0 {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	// Use the first instance to show base Pokemon details
	firstInstance := instances[0]

	// Convert to Pokemon object for common display format
	basePokemon := models.Pokemon{
		Name:   firstInstance.Name,
		Height: 0, // These will be filled in below
		Weight: 0,
	}

	// Get details from pokedex for complete information if needed
	pokedexPokemon, err := c.pokemonService.GetPokemonByName(ctx, c.currentTrainer.ID, pokemon)
	if err == nil {
		// Use pokedex data for base stats and type info
		basePokemon = pokedexPokemon
	} else {
		// If not in pokedex, try to extract info from in-memory map
		if pkm, ok := c.caughtPokemon[pokemon]; ok {
			basePokemon = pkm
		}
	}

	// Show base details
	displayPokemonDetails(basePokemon)

	// Display instances with skills
	fmt.Printf("\nYou own %d %s:\n", len(instances), pokemon)

	// Display each instance with its skills
	displayPokemonInstances(instances)

	return nil
}

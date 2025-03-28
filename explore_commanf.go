package main

import (
	"fmt"
)

func commandExplore(area string, c *Config) error {
	if area == "" {
		return fmt.Errorf("no area provided")
	}

	exploreds, err := c.pokeapiClient.GetExploredPokemonList(area)
	if err != nil {
		return err // Return the error instead of log.Fatal
	}

	fmt.Println("Exploring " + area + "...")
	fmt.Println("Found Pokemon:")

	for _, explored := range exploreds.PokemonEncounters {
		fmt.Println("- " + explored.Pokemon.Name)
	}

	return nil
}

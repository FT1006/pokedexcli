package main

import "fmt"

func commandPokedex(additionalInput string, c *Config) error {
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

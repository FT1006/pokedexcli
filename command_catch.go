package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/FT1006/pokedexcli/internal/models"
)

func commandCatch(pokemon string, c *Config) error {
	if pokemon == "" {
		return fmt.Errorf("no pokemon provided")
	}

	pokemonInfo, err := c.pokeapiClient.GetPokemonDetails(pokemon)
	if err != nil {
		return err // Return the error instead of log.Fatal
	}

	var catchChance float64
	catchChance = 1.0 - float64(pokemonInfo.BaseExperience)/300.0
	if catchChance > 0.9 {
		catchChance = 0.9
	} else if catchChance < 0.1 {
		catchChance = 0.1
	}

	fmt.Println("catch chance: ", catchChance*100, "%")
	fmt.Println("Throwing a Pokeball at " + pokemon + "...")
	time.Sleep(time.Second * 2)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Float64() < catchChance {
		fmt.Println(pokemonInfo.Name + " was caught!")
		fmt.Println("You may now inspect it with the inspect command.")

		caughtStats := make([]models.Stats, len(pokemonInfo.Stats))

		for i, stat := range pokemonInfo.Stats {
			caughtStats[i] = models.Stats{
				BaseStat: stat.BaseStat,
				Effort:   stat.Effort,
				Stat: models.Stat{
					Name: stat.Stat.Name,
					URL:  stat.Stat.URL,
				},
			}
		}

		caughtTypes := make([]models.Types, len(pokemonInfo.Types))
		for i, typ := range pokemonInfo.Types {
			caughtTypes[i] = models.Types{
				Slot: typ.Slot,
				Type: models.Type{
					Name: typ.Type.Name,
					URL:  typ.Type.URL,
				},
			}
		}

		caughtPokemon := models.Pokemon{
			Name:           pokemonInfo.Name,
			Height:         pokemonInfo.Height,
			Weight:         pokemonInfo.Weight,
			Stats:          caughtStats,
			Types:          caughtTypes,
			BaseExperience: pokemonInfo.BaseExperience,
		}

		// Add to in-memory maps
		c.caughtPokemon[pokemon] = caughtPokemon
		c.newlyCaughtPokemon[pokemon] = caughtPokemon

		// If database is initialized, save to database
		if c.dbService != nil && c.currentTrainer != nil {
			ctx := context.Background()
			err := c.pokemonService.SavePokemon(ctx, c.currentTrainer.ID, caughtPokemon)
			if err != nil {
				fmt.Printf("Warning: could not save pokemon to database: %v\n", err)
			}
		}

	} else {
		fmt.Println(pokemonInfo.Name + " escaped!")
	}

	return nil
}
package main

import (
	"fmt"
	"math/rand"
	"time"
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

		caughtStats := make([]Stats, len(pokemonInfo.Stats))

		for i, stat := range pokemonInfo.Stats {
			caughtStats[i] = Stats{
				BaseStat: stat.BaseStat,
				Effort:   stat.Effort,
				Stat: Stat{
					Name: stat.Stat.Name,
					URL:  stat.Stat.URL,
				},
			}
		}

		caughtTypes := make([]Types, len(pokemonInfo.Types))
		for i, typ := range pokemonInfo.Types {
			caughtTypes[i] = Types{
				Slot: typ.Slot,
				Type: Type{
					Name: typ.Type.Name,
					URL:  typ.Type.URL,
				},
			}
		}

		c.caughtPokemon[pokemon] = Pokemon{
			Name:   pokemonInfo.Name,
			Height: pokemonInfo.Height,
			Weight: pokemonInfo.Weight,
			Stats:  caughtStats,
			Types:  caughtTypes,
		}

	} else {
		fmt.Println(pokemonInfo.Name + " escaped!")
	}

	return nil
}

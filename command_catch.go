package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
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

		// Create base Pokemon structure
		caughtPokemon := models.Pokemon{
			Name:           pokemonInfo.Name,
			Height:         pokemonInfo.Height,
			Weight:         pokemonInfo.Weight,
			Stats:          caughtStats,
			Types:          caughtTypes,
			BaseExperience: pokemonInfo.BaseExperience,
		}

		// Get skills for the Pokemon
		fmt.Println("Determining " + pokemon + "'s skills...")
		basicSkill, specialSkill, err := c.pokeapiClient.GetMovesForPokemon(pokemon)
		if err != nil {
			fmt.Printf("Warning: could not get skills for %s: %v\n", pokemon, err)
			fmt.Println("Pokemon will be caught without skills.")
		} else {
			// Assign skills to the Pokemon
			caughtPokemon.BasicSkill = &basicSkill
			caughtPokemon.SpecialSkill = &specialSkill

			// Display the skills
			fmt.Printf("%s learned %s (%s type, %s class) and %s (%s type, %s class)!\n",
				pokemon,
				basicSkill.Name,
				basicSkill.Type,
				basicSkill.Class,
				specialSkill.Name,
				specialSkill.Type,
				specialSkill.Class)
		}
		fmt.Println("You may now inspect it with the inspect command.")

		// Add to in-memory maps
		c.caughtPokemon[pokemon] = caughtPokemon
		c.newlyCaughtPokemon[pokemon] = caughtPokemon

		// If database is initialized, save to database and add to party
		if c.dbService != nil && c.currentTrainer != nil {
			ctx := context.Background()

			// Save the Pokemon to the database with skills
			var latestPokemonID int32

			if caughtPokemon.BasicSkill != nil && caughtPokemon.SpecialSkill != nil {
				// Use the new method to save Pokemon with skills
				pokemonID, err := c.pokemonService.AddPokemonWithSkills(
					ctx,
					c.currentTrainer.ID,
					caughtPokemon,
					caughtPokemon.BasicSkill,
					caughtPokemon.SpecialSkill)

				if err != nil {
					fmt.Printf("Warning: could not save pokemon with skills: %v\n", err)
					fmt.Println("Falling back to saving without skills...")

					// Fallback to saving without skills
					err = c.pokemonService.SavePokemon(ctx, c.currentTrainer.ID, caughtPokemon)
					if err != nil {
						fmt.Printf("Warning: could not save pokemon to database: %v\n", err)
						return nil
					}
				} else {
					latestPokemonID = pokemonID
					// We already have the Pokemon ID, skip to party processing
					goto ProcessParty
				}
			} else {
				// Save without skills (default behavior)
				err = c.pokemonService.SavePokemon(ctx, c.currentTrainer.ID, caughtPokemon)
				if err != nil {
					fmt.Printf("Warning: could not save pokemon to database: %v\n", err)
					return nil
				}
			}

			// If we don't already have the Pokemon ID from AddPokemonWithSkills
			if latestPokemonID == 0 {
				// Get the OwnpokeID of the Pokemon we just caught
				ownedPokemon, err := c.dbService.Queries().ListOwnedPokemonByTrainer(ctx, c.currentTrainer.ID)
				if err != nil || len(ownedPokemon) == 0 {
					fmt.Printf("Warning: could not get caught Pokemon ID: %v\n", err)
					return nil
				}

				// Assuming the most recently caught Pokemon is at index 0 (ordered by caught_at DESC)
				latestPokemonID = ownedPokemon[0].ID
			}

		ProcessParty:
			// Check if party is full
			partyCount, err := c.partyService.GetPartyCount(ctx, c.currentTrainer.ID)
			if err != nil {
				fmt.Printf("Warning: could not check party count: %v\n", err)
				return nil
			}

			if partyCount < 6 {
				// Add to next available slot
				slot, err := c.partyService.AddToNextAvailableSlot(ctx, c.currentTrainer.ID, latestPokemonID)
				if err != nil {
					fmt.Printf("Warning: could not add Pokemon to party: %v\n", err)
				} else {
					fmt.Printf("%s was added to your party in slot %d!\n", caughtPokemon.Name, slot)
				}
			} else {
				// Party is full, prompt user to replace
				fmt.Println("Your party is full! Would you like to replace a Pokemon?")

				// Display current party
				partyPokemon, err := c.partyService.GetParty(ctx, c.currentTrainer.ID)
				if err != nil {
					fmt.Printf("Warning: could not get party: %v\n", err)
					return nil
				}

				for _, p := range partyPokemon {
					fmt.Printf("%d. %s", p.Slot, p.Name)
					// Display skills if available
					if p.BasicSkill != nil {
						fmt.Printf(" (Basic: %s)", p.BasicSkill.Name)
					}
					if p.SpecialSkill != nil {
						fmt.Printf(" (Special: %s)", p.SpecialSkill.Name)
					}
					fmt.Println()
				}

				fmt.Println("Enter a slot number (1-6) to replace, or 'n' to keep your current party:")
				var input string
				if _, err := fmt.Scanln(&input); err != nil {
					fmt.Println("Invalid input. Your party remains unchanged.")
				}

				// Parse input
				if input == "n" || input == "N" {
					fmt.Println("Your party remains unchanged.")
				} else {
					slot, err := strconv.Atoi(input)
					if err != nil || slot < 1 || slot > 6 {
						fmt.Println("Invalid input. Your party remains unchanged.")
					} else {
						// Replace Pokemon in selected slot
						err = c.partyService.AddPokemonToParty(ctx, c.currentTrainer.ID, latestPokemonID, slot)
						if err != nil {
							fmt.Printf("Warning: could not replace Pokemon: %v\n", err)
						} else {
							fmt.Printf("%s was added to your party in slot %d!\n", caughtPokemon.Name, slot)
						}
					}
				}
			}
		}

	} else {
		fmt.Println(pokemonInfo.Name + " escaped!")
	}

	return nil
}

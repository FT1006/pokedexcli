package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/FT1006/pokedexcli/internal/models"
	"github.com/FT1006/pokedexcli/internal/pokeapi"
)

// BattlePokemon conducts a battle between userPokemon and wildPokemon
// Returns true if the user wins, false if the user loses
func BattlePokemon(userPokemon models.Pokemon, wildPokemonInput interface{}) bool {
	// Convert wildPokemonInput to models.Pokemon if needed
	var wildPokemon models.Pokemon

	switch wp := wildPokemonInput.(type) {
	case models.Pokemon:
		// If it's already a models.Pokemon, use it directly
		wildPokemon = wp
	case pokeapi.PokemonRes:
		// If it's a PokemonRes, convert it to models.Pokemon
		// Create Stats array
		wildPokemonStats := make([]models.Stats, len(wp.Stats))
		for i, stat := range wp.Stats {
			wildPokemonStats[i] = models.Stats{
				BaseStat: stat.BaseStat,
				Effort:   stat.Effort,
				Stat: models.Stat{
					Name: stat.Stat.Name,
					URL:  stat.Stat.URL,
				},
			}
		}

		// Create Types array
		wildPokemonTypes := make([]models.Types, len(wp.Types))
		for i, typ := range wp.Types {
			wildPokemonTypes[i] = models.Types{
				Slot: typ.Slot,
				Type: models.Type{
					Name: typ.Type.Name,
					URL:  typ.Type.URL,
				},
			}
		}

		// Build Pokemon model
		wildPokemon = models.Pokemon{
			Name:           wp.Name,
			Height:         wp.Height,
			Weight:         wp.Weight,
			Stats:          wildPokemonStats,
			Types:          wildPokemonTypes,
			BaseExperience: wp.BaseExperience,
		}
	default:
		// For unexpected types, log warning and create a default Pokemon
		fmt.Println("Warning: Unexpected Pokemon type in battle. Using default wild Pokemon.")
		wildPokemon = models.Pokemon{
			Name: "Wild Pokemon",
			Stats: []models.Stats{
				{BaseStat: 50, Stat: models.Stat{Name: "hp"}},
				{BaseStat: 30, Stat: models.Stat{Name: "attack"}},
				{BaseStat: 20, Stat: models.Stat{Name: "defense"}},
				{BaseStat: 30, Stat: models.Stat{Name: "special-attack"}},
			},
			BaseExperience: 50,
		}
	}
	// Battle state tracking
	type battlePokemon struct {
		pokemon      models.Pokemon
		currentHP    int
		maxHP        int
		specialUsed  bool
		name         string
		attackStat   int
		defenseStat  int
		spAttackStat int
	}

	// Find stats for both Pokemon
	var userAttack, userDefense, userSpAttack int
	var wildAttack, wildDefense, wildSpAttack int

	// Get user Pokemon stats
	for _, stat := range userPokemon.Stats {
		switch stat.Stat.Name {
		case "hp":
			// Skip HP, we'll set it separately
		case "attack":
			userAttack = stat.BaseStat
		case "defense":
			userDefense = stat.BaseStat
		case "special-attack":
			userSpAttack = stat.BaseStat
		}
	}

	// Get wild Pokemon stats
	for _, stat := range wildPokemon.Stats {
		switch stat.Stat.Name {
		case "hp":
			// Skip HP, we'll set it separately
		case "attack":
			wildAttack = stat.BaseStat
		case "defense":
			wildDefense = stat.BaseStat
		case "special-attack":
			wildSpAttack = stat.BaseStat
		}
	}

	// Set up battle Pokemon with initial state
	// Get HP stat or use default if not found
	userMaxHP := 100
	wildMaxHP := 100

	for _, stat := range userPokemon.Stats {
		if stat.Stat.Name == "hp" {
			userMaxHP = stat.BaseStat
			break
		}
	}

	for _, stat := range wildPokemon.Stats {
		if stat.Stat.Name == "hp" {
			wildMaxHP = stat.BaseStat
			break
		}
	}

	user := battlePokemon{
		pokemon:      userPokemon,
		currentHP:    userMaxHP,
		maxHP:        userMaxHP,
		specialUsed:  false,
		name:         userPokemon.Name,
		attackStat:   userAttack,
		defenseStat:  userDefense,
		spAttackStat: userSpAttack,
	}

	wild := battlePokemon{
		pokemon:      wildPokemon,
		currentHP:    wildMaxHP,
		maxHP:        wildMaxHP,
		specialUsed:  false,
		name:         wildPokemon.Name,
		attackStat:   wildAttack,
		defenseStat:  wildDefense,
		spAttackStat: wildSpAttack,
	}

	// Begin battle
	fmt.Printf("\n--- BATTLE START ---\n")
	fmt.Printf("Your %s (HP: %d) vs Wild %s (HP: %d)\n\n",
		user.name, user.currentHP, wild.name, wild.currentHP)

	// Check if Pokemon have skills
	if user.pokemon.BasicSkill == nil || user.pokemon.SpecialSkill == nil {
		fmt.Println("Your Pokemon doesn't have skills! Using default attacks.")
		user.pokemon.BasicSkill = &models.Skill{
			Name:   "Tackle",
			Damage: 40,
			Type:   "normal",
			Class:  "physical",
		}
		user.pokemon.SpecialSkill = &models.Skill{
			Name:   "Quick Attack",
			Damage: 60,
			Type:   "normal",
			Class:  "physical",
		}
	}

	if wild.pokemon.BasicSkill == nil || wild.pokemon.SpecialSkill == nil {
		wild.pokemon.BasicSkill = &models.Skill{
			Name:   "Tackle",
			Damage: 40,
			Type:   "normal",
			Class:  "physical",
		}
		wild.pokemon.SpecialSkill = &models.Skill{
			Name:   "Quick Attack",
			Damage: 60,
			Type:   "normal",
			Class:  "physical",
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Main battle loop
	for user.currentHP > 0 && wild.currentHP > 0 {
		// Display HP
		fmt.Printf("Your %s: %d/%d HP | Wild %s: %d/%d HP\n",
			user.name, user.currentHP, user.maxHP,
			wild.name, wild.currentHP, wild.maxHP)

		// User's turn
		fmt.Println("\nYour turn! Choose an attack:")
		fmt.Printf("1. %s (Basic)\n", user.pokemon.BasicSkill.Name)
		specialOption := "2. " + user.pokemon.SpecialSkill.Name + " (Special)"

		if user.specialUsed {
			specialOption += " [USED]"
		}
		fmt.Println(specialOption)

		// Get user choice
		var choice string
		var moveChoice int
		validChoice := false

		for !validChoice {
			fmt.Print("Enter choice (1-2): ")
			fmt.Scanln(&choice)

			moveChoice, _ = strconv.Atoi(choice)

			if moveChoice == 1 {
				validChoice = true
			} else if moveChoice == 2 && !user.specialUsed {
				validChoice = true
			} else if moveChoice == 2 && user.specialUsed {
				fmt.Println("You've already used your special move!")
			} else {
				fmt.Println("Invalid choice. Try again.")
			}
		}

		// Execute user's move
		userHit := true
		if moveChoice == 2 {
			// Special move has 50% chance to hit
			userHit = r.Float64() < 0.5
			user.specialUsed = true
		}

		var damage int

		if userHit {
			if moveChoice == 1 {
				// Basic attack
				damage = int(float64(user.pokemon.BasicSkill.Damage+user.attackStat-wild.defenseStat) * 0.1)
				if damage < 1 {
					damage = 1
				}
				fmt.Printf("Your %s used %s! It dealt %d damage!\n",
					user.name, user.pokemon.BasicSkill.Name, damage)
			} else {
				// Special attack
				specialDamage := float64(user.pokemon.SpecialSkill.Damage+user.spAttackStat) * 1.5
				damage = int((specialDamage - float64(wild.defenseStat)) * 0.1)
				if damage < 1 {
					damage = 1
				}
				fmt.Printf("Your %s used %s! It dealt %d damage!\n",
					user.name, user.pokemon.SpecialSkill.Name, damage)
			}

			wild.currentHP -= damage
			if wild.currentHP < 0 {
				wild.currentHP = 0
			}
		} else {
			fmt.Printf("Your %s used %s! But it missed!\n",
				user.name, user.pokemon.SpecialSkill.Name)
		}

		// Check if wild Pokemon fainted
		if wild.currentHP <= 0 {
			fmt.Printf("Wild %s fainted!\n", wild.name)
			break
		}

		// Wild Pokemon's turn
		fmt.Printf("\nWild %s's turn!\n", wild.name)

		// Wild Pokemon chooses move (75% basic, 25% special if not used)
		wildMoveChoice := 1 // Default to basic
		if !wild.specialUsed && r.Float64() < 0.25 {
			wildMoveChoice = 2
			wild.specialUsed = true
		}

		// Execute wild Pokemon's move
		wildHit := true
		if wildMoveChoice == 2 {
			// Special move has 50% chance to hit
			wildHit = r.Float64() < 0.5
		}

		if wildHit {
			if wildMoveChoice == 1 {
				// Basic attack
				damage = int(float64(wild.pokemon.BasicSkill.Damage+wild.attackStat-user.defenseStat) * 0.1)
				if damage < 1 {
					damage = 1
				}
				fmt.Printf("Wild %s used %s! It dealt %d damage!\n",
					wild.name, wild.pokemon.BasicSkill.Name, damage)
			} else {
				// Special attack
				specialDamage := float64(wild.pokemon.SpecialSkill.Damage+wild.spAttackStat) * 1.5
				damage = int((specialDamage - float64(user.defenseStat)) * 0.1)
				if damage < 1 {
					damage = 1
				}
				fmt.Printf("Wild %s used %s! It dealt %d damage!\n",
					wild.name, wild.pokemon.SpecialSkill.Name, damage)
			}

			user.currentHP -= damage
			if user.currentHP < 0 {
				user.currentHP = 0
			}
		} else {
			fmt.Printf("Wild %s used %s! But it missed!\n",
				wild.name, wild.pokemon.SpecialSkill.Name)
		}

		// Check if user's Pokemon fainted
		if user.currentHP <= 0 {
			fmt.Printf("Your %s fainted!\n", user.name)
			break
		}

		// Pause before next turn
		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln()
	}

	// Determine battle outcome
	fmt.Println("\n--- BATTLE END ---")

	if wild.currentHP <= 0 {
		fmt.Printf("You won the battle against wild %s!\n", wild.name)
		return true
	} else {
		fmt.Printf("You lost the battle against wild %s!\n", wild.name)
		return false
	}
}

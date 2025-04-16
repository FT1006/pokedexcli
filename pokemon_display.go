package main

import (
	"fmt"
	"time"

	"github.com/FT1006/pokedexcli/internal/database/service"
	"github.com/FT1006/pokedexcli/internal/models"
)

// FormatTime returns a nicely formatted time string
func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// DisplayPokemonDetails prints the general info about a Pokemon
func displayPokemonDetails(pkm models.Pokemon) {
	fmt.Printf("Name: %s\n", pkm.Name)
	fmt.Printf("Height: %d\n", pkm.Height)
	fmt.Printf("Weight: %d\n", pkm.Weight)
	fmt.Println("Stats:")
	for _, stat := range pkm.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pkm.Types {
		fmt.Printf("  -%s\n", typ.Type.Name)
	}
}

// DisplayPokemonInstances displays a list of Pokemon instances with their skills
func displayPokemonInstances(instances []service.OwnedPokemon) {
	for _, p := range instances {
		fmt.Printf("\nOwnpoke ID[%d] - %s - Caught: %s\n",
			p.ID,
			p.Name,
			formatTime(p.CaughtAt))

		// Display skills
		fmt.Println("Skills:")
		if p.BasicSkill != nil {
			fmt.Printf("  • Basic: %s (%d damage) - %s type, %s class\n",
				p.BasicSkill.Name,
				p.BasicSkill.Damage,
				p.BasicSkill.Type,
				p.BasicSkill.Class)
		} else {
			fmt.Println("  • Basic: None")
		}

		if p.SpecialSkill != nil {
			fmt.Printf("  • Special: %s (%d damage) - %s type, %s class\n",
				p.SpecialSkill.Name,
				p.SpecialSkill.Damage,
				p.SpecialSkill.Type,
				p.SpecialSkill.Class)
		} else {
			fmt.Println("  • Special: None")
		}
	}
}

// FormatTypesString formats Pokemon types as a comma-separated string
func formatTypesString(types []models.Types) string {
	typesStr := ""
	for i, t := range types {
		if i > 0 {
			typesStr += ", "
		}
		typesStr += t.Type.Name
	}
	return typesStr
}

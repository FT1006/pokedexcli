package main

import (
	"context"
	"fmt"
)

func commandParty(additionalInput string, c *Config) error {
	if c.dbService == nil {
		return fmt.Errorf("database not initialized")
	}

	if c.currentTrainer == nil {
		return fmt.Errorf("no trainer logged in - use 'load' command first")
	}

	ctx := context.Background()
	partyPokemon, err := c.partyService.GetParty(ctx, c.currentTrainer.ID)
	if err != nil {
		return fmt.Errorf("error getting party: %w", err)
	}

	if len(partyPokemon) == 0 {
		fmt.Println("Your party is empty! Catch some Pokemon and add them to your party.")
		return nil
	}

	fmt.Printf("Your Pokemon party (%d/6):\n", len(partyPokemon))

	// Use a custom format for party display to include the slot
	for _, p := range partyPokemon {
		fmt.Printf("\nSlot %d - Ownpoke ID[%d] - %s\n",
			p.Slot,
			p.OwnpokeID,
			p.Name,
		)

		// Format types
		typesStr := formatTypesString(p.Types)
		fmt.Printf("Types: %s\n", typesStr)

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

	return nil
}

package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
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

	fmt.Printf("Your Pokemon party (%d/6):\n\n", len(partyPokemon))

	// Create a tabwriter for nicely aligned output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "Slot\tName\tTypes\tBase Experience")
	fmt.Fprintln(w, "----\t----\t-----\t---------------")

	for _, p := range partyPokemon {
		// Format types as comma-separated list
		typesStr := ""
		for i, t := range p.Types {
			if i > 0 {
				typesStr += ", "
			}
			typesStr += t.Type.Name
		}

		fmt.Fprintf(w, "%d\t%s\t%s\t%d\n", 
			p.Slot, 
			p.Name, 
			typesStr,
			p.BaseExperience)
	}
	w.Flush()

	return nil
}
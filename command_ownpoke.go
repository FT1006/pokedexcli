package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func commandOwnPoke(additionalInput string, c *Config) error {
	if c.dbService == nil {
		return fmt.Errorf("database not initialized")
	}

	if c.currentTrainer == nil {
		return fmt.Errorf("no trainer logged in - use 'load' command first")
	}

	ctx := context.Background()
	ownedPokemon, err := c.pokemonService.GetAllOwnedPokemon(ctx, c.currentTrainer.ID)
	if err != nil {
		return fmt.Errorf("error getting owned pokemon: %w", err)
	}

	if len(ownedPokemon) == 0 {
		fmt.Println("You don't have any caught Pokemon yet!")
		return nil
	}

	fmt.Printf("Your owned Pokemon collection (%d total):\n\n", len(ownedPokemon))

	// Create a tabwriter for nicely aligned output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "Name\tCaught At")
	fmt.Fprintln(w, "----\t---------")

	for _, p := range ownedPokemon {
		fmt.Fprintf(w, "%s\t%s\n", p.Name, formatCaughtTime(p.CaughtAt))
	}
	w.Flush()

	return nil
}

func formatCaughtTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
package main

import (
	"context"
	"fmt"
	
	dbsqlc "github.com/FT1006/pokedexcli/internal/database/sqlc/db"
)

func commandTrainer(_ string, c *Config) error {
	// Check if database service is initialized
	if c.dbService == nil {
		return fmt.Errorf("database not initialized, please run the application with a database connection")
	}

	// Query all trainers from the database
	ctx := context.Background()
	
	// ListTrainers requires limit and offset parameters
	params := dbsqlc.ListTrainersParams{
		Limit:  100, // Reasonable limit for most users
		Offset: 0,   // Start from the first trainer
	}
	
	trainers, err := c.dbService.Queries().ListTrainers(ctx, params)
	if err != nil {
		return fmt.Errorf("error retrieving trainers: %v", err)
	}

	// Display trainer information
	if len(trainers) == 0 {
		fmt.Println("No trainers found. Use 'save <trainer-name>' to create a new trainer.")
		return nil
	}

	fmt.Println("\n===== Available Trainers =====")
	fmt.Println("ID\tName\tPokemon\tCreated")
	fmt.Println("------------------------------------------")

	for _, t := range trainers {
		// Get count of Pokemon for this trainer by querying directly
		var pokemonCount int
		rows, err := c.dbService.Queries().ListOwnedPokemonByTrainer(ctx, t.ID)
		if err != nil {
			pokemonCount = 0 // Default to 0 if there's an error
		} else {
			pokemonCount = len(rows)
		}

		// Format creation time
		createdAt := t.CreatedAt.Time
		timeStr := createdAt.Format("Jan 02, 2006")

		// Display trainer with their Pokemon count
		fmt.Printf("%d\t%s\t%d\t%s\n", 
			t.ID, 
			t.Name.String, // Convert pgtype.Text to string
			pokemonCount,
			timeStr)
	}
	
	fmt.Println("\nUse 'load <trainer-name>' to load a trainer profile.")
	fmt.Println("==============================")

	return nil
}
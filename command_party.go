package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/FT1006/pokedexcli/internal/database/service"
)

func commandParty(additionalInput string, c *Config) error {
	if c.dbService == nil {
		return fmt.Errorf("database not initialized")
	}

	if c.currentTrainer == nil {
		return fmt.Errorf("no trainer logged in - use 'load' command first")
	}

	ctx := context.Background()

	// Check for the "change" argument
	if strings.TrimSpace(additionalInput) == "change" {
		return handlePartyChange(ctx, c)
	}

	// Display current party
	return displayParty(ctx, c)
}

// displayParty shows the current party with all details
func displayParty(ctx context.Context, c *Config) error {
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

// handlePartyChange manages the party change workflow
func handlePartyChange(ctx context.Context, c *Config) error {
	// Step 1: Get current party
	partyPokemon, err := c.partyService.GetParty(ctx, c.currentTrainer.ID)
	if err != nil {
		return fmt.Errorf("error getting party: %w", err)
	}

	if len(partyPokemon) == 0 {
		fmt.Println("Your party is empty! Catch some Pokemon first.")
		return nil
	}

	// Display current party in compact format
	fmt.Println("Current party:")
	for _, p := range partyPokemon {
		fmt.Printf("Slot %d: %s (ID: %d)\n", p.Slot, p.Name, p.OwnpokeID)
	}
	fmt.Println()

	// Step 2: Get all owned Pokemon that are not in the party
	allPokemon, err := c.pokemonService.GetAllOwnedPokemon(ctx, c.currentTrainer.ID)
	if err != nil {
		return fmt.Errorf("error getting owned pokemon: %w", err)
	}

	// Map of party Pokemon IDs for quick lookup
	partyIDs := make(map[int32]bool)
	for _, p := range partyPokemon {
		partyIDs[p.OwnpokeID] = true
	}

	// Filter out Pokemon already in party
	var availablePokemon []service.OwnedPokemon
	for _, p := range allPokemon {
		if !partyIDs[p.ID] {
			availablePokemon = append(availablePokemon, p)
		}
	}

	if len(availablePokemon) == 0 {
		fmt.Println("You don't have any other Pokemon available to swap into your party.")
		return nil
	}

	// Step 3: Display available Pokemon
	fmt.Println("Available Pokemon:")
	for i, p := range availablePokemon {
		fmt.Printf("%d. %s (ID: %d)\n",
			i+1,
			p.Name,
			p.ID)
		// Display skills in compact format
		var skills []string
		if p.BasicSkill != nil {
			skills = append(skills, fmt.Sprintf("%s (%s)", p.BasicSkill.Name, p.BasicSkill.Type))
		}
		if p.SpecialSkill != nil {
			skills = append(skills, fmt.Sprintf("%s (%s)", p.SpecialSkill.Name, p.SpecialSkill.Type))
		}

		if len(skills) > 0 {
			fmt.Printf("   Skills: %s\n", strings.Join(skills, ", "))
		} else {
			fmt.Println("   No skills")
		}
	}

	// Step 4: Ask user to select a Pokemon by ID
	fmt.Print("\nEnter the ID of the Pokemon you want to add to your party: ")
	var pokemonIDInput string
	fmt.Scanln(&pokemonIDInput)

	pokemonID, err := strconv.ParseInt(pokemonIDInput, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid Pokemon ID: %s", pokemonIDInput)
	}

	// Verify the Pokemon exists
	var selectedPokemon *service.OwnedPokemon
	for i := range availablePokemon {
		if availablePokemon[i].ID == int32(pokemonID) {
			selectedPokemon = &availablePokemon[i]
			break
		}
	}

	if selectedPokemon == nil {
		return fmt.Errorf("no Pokemon found with ID %d or it's already in your party", pokemonID)
	}

	// Step 5: Ask user which party slot to replace
	fmt.Print("\nEnter the party slot (1-6) to replace: ")
	var slotInput string
	fmt.Scanln(&slotInput)

	slot, err := strconv.Atoi(slotInput)
	if err != nil || slot < 1 || slot > 6 {
		return fmt.Errorf("invalid slot number: must be between 1 and 6")
	}

	// Check if the slot is occupied
	slotOccupied := false
	for _, p := range partyPokemon {
		if p.Slot == slot {
			slotOccupied = true
			break
		}
	}

	// Step 6: Add the Pokemon to the party
	err = c.partyService.AddPokemonToParty(ctx, c.currentTrainer.ID, int32(pokemonID), slot)
	if err != nil {
		return fmt.Errorf("error updating party: %w", err)
	}

	// Step 7: Display success message
	if slotOccupied {
		fmt.Printf("Successfully replaced the Pokemon in slot %d with %s (ID: %d)!\n",
			slot, selectedPokemon.Name, selectedPokemon.ID)
	} else {
		fmt.Printf("Successfully added %s (ID: %d) to slot %d!\n",
			selectedPokemon.Name, selectedPokemon.ID, slot)
	}

	// Step 8: Display updated party
	fmt.Println("\nUpdated party:")
	return displayParty(ctx, c)
}

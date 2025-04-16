package main

import (
	"fmt"
	"slices"
	"sort"
)

// Helper function to print commands by category
func printCommandsByCategory(allCommands []string, commandMap map[string]cliCommand, categoryCommands []string) {
	for _, cmd := range categoryCommands {
		if slices.Contains(allCommands, cmd) {
			cmdVal := commandMap[cmd]
			fmt.Printf("  %-15s %s\n", cmdVal.name+":", cmdVal.description)
		}
	}
}

func commandHelp(additionalInput string, c *Config) error {
	if additionalInput != "" {
		fmt.Println("Additional input ignored")
	}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	// Get commands and prepare for display
	commands := callCommand()
	var cmdNames []string
	for cmd := range commands {
		cmdNames = append(cmdNames, cmd)
	}

	// Sort commands alphabetically
	sort.Strings(cmdNames)

	// Group commands by category for better organization
	fmt.Println("Navigation & Basic:")
	fmt.Println("------------------")
	printCommandsByCategory(cmdNames, commands, []string{"exit", "help", "map", "mapb", "explore"})

	fmt.Println("\nPokemon Management:")
	fmt.Println("------------------")
	printCommandsByCategory(cmdNames, commands, []string{"catch", "inspect", "pokedex", "ownpoke", "party"})

	fmt.Println("\nSave & Load:")
	fmt.Println("------------------")
	printCommandsByCategory(cmdNames, commands, []string{"trainer", "save", "load"})

	fmt.Println("\nInformation:")
	fmt.Println("------------------")
	printCommandsByCategory(cmdNames, commands, []string{"battle"})

	fmt.Println()
	return nil
}

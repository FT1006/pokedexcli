package main

import (
	"fmt"
	"os"
)

func commandExit(additionalInput string, c *Config) error {
	if additionalInput != "" {
		fmt.Println("Additional input has been ignored")
	}

	// Make sure we have a clean line with proper cursor position
	// Clear any terminal escape sequences that might be in effect
	os.Stdout.WriteString("\r\n")

	// Print farewell message
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	// Use a sentinel error to signal exit
	return nil
}

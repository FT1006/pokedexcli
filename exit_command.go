package main

import (
	"fmt"
	"os"
)

func commandExit(additionalInput string, c *Config) error {
	if additionalInput != "" {
		fmt.Println("Additional input has been ignored")
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

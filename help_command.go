package main

import (
	"fmt"
	"sort"
)

func commandHelp(additionalInput string, c *Config) error {
	if additionalInput != "" {
		fmt.Println("Additional input ignored")
	}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage\n\n")
	var cmdNames []string
	for cmd := range callCommand() {
		cmdNames = append(cmdNames, cmd)
	}
	sort.Slice(cmdNames, func(i, j int) bool {
		return cmdNames[i] < cmdNames[j]
	})
	for _, cmdName := range cmdNames {
		cmdVal := callCommand()[cmdName]
		fmt.Printf("%s: %s\n", cmdVal.name, cmdVal.description)
	}
	return nil
}

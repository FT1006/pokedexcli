package main

import (
	"fmt"
	"sort"
)

func commandHelp(c *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage\n\n")
	var cmdNames []string
	for _, c := range callCommand() {
		cmdNames = append(cmdNames, c.name)
	}
	sort.Slice(cmdNames, func(i, j int) bool {
		return cmdNames[i] < cmdNames[j]
	})
	for _, cmdName := range cmdNames {
		fmt.Printf("%s: %s\n", cmdName, callCommand()[cmdName].description)
	}
	return nil
}

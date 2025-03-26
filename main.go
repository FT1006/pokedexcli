package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FT1006/pokedexcli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*internal.Config) error
}

func callCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 maps",
			callback:    internal.MapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 maps",
			callback:    internal.MapBackCommand,
		},
	}
}

func main() {
	cfg := &internal.Config{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		if c, exist := callCommand()[cleanedInput[0]]; !exist {
			fmt.Println("Unknown command")
		} else {
			if err := c.callback(cfg); err != nil {
				fmt.Println(err)
			}
		}
	}
}

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for {
// 		fmt.Print("Pokedex > ")
// 		scanner.Scan()
// 		input := scanner.Text()
// 		cleanedInput := cleanInput(input)
// 		fmt.Printf("Your command was: %s\n", cleanedInput[0])
// 	}
// }

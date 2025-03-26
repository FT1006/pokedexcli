package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func callCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
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
			fmt.Println("Unkown command")
		} else {
			if err := c.callback(); err != nil {
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

package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func cleanInput(text string) []string {
// 	text = strings.ToLower(text)
// 	words := strings.Fields(text)
// 	return words
// }

// type cliCommand struct {
// 	name        string
// 	description string
// 	callback    func(string, *Config) error
// }

// func repl(cfg *Config) {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for {
// 		fmt.Print("Pokedex > ")
// 		scanner.Scan()
// 		input := scanner.Text()
// 		cleanedInput := cleanInput(input)

// 		if len(cleanedInput) == 0 {
// 			continue
// 		}

// 		command := cleanedInput[0]
// 		additionalInput := ""

// 		if len(cleanedInput) > 1 {
// 			additionalInput = cleanedInput[1]
// 		}
// 		if c, exist := callCommand()[command]; !exist {
// 			fmt.Println("Unknown command")
// 		} else {
// 			if err := c.callback(additionalInput, cfg); err != nil {
// 				fmt.Println(err)
// 			}
// 		}
// 	}
// }

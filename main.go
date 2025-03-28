package main

import "github.com/FT1006/pokedexcli/internal/pokeapi"

func main() {
	pokeapiClient := pokeapi.NewClient()

	cfg := Config{
		pokeapiClient: pokeapiClient,
		next:          "",
		prev:          "",
	}

	repl(&cfg)
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

package main

func main() {
	var cfg Config
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

package main

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
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 maps",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore [area name]",
			description: "Explore an area. Example: explore pastoria-city-area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch [pokemon name]",
			description: "Try to catch a pokemon. Example: catch pikachu",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect [pokemon name]",
			description: "Inspect a pokemon. Example: inspect pikachu",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show pokedex",
			callback:    commandPokedex,
		},
	}
}

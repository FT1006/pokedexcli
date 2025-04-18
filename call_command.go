package main

func callCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"battle": {
			name:        "battle",
			description: "Show information about the battle system",
			callback:    commandBattle,
		},
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
			name:        "explore <area-name>",
			description: "Explore an area. Example: explore pastoria-city-area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Try to catch a pokemon. Example: catch pikachu",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Inspect a pokemon with all its instances and skills. Example: inspect pikachu",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View unique Pokemon you've caught (no duplicates)",
			callback:    commandPokedex,
		},
		"ownpoke": {
			name:        "ownpoke",
			description: "View all Pokemon you've caught including duplicates with skills and catch times",
			callback:    commandOwnPoke,
		},
		"party": {
			name:        "party [change]",
			description: "View your active Pokemon party (up to 6 Pokemon) with skills. Use 'party change' to substitute party members",
			callback:    commandParty,
		},
		"save": {
			name:        "save [trainer-name]",
			description: "Save your trainer data and caught Pokemon. The trainer name is optional.",
			callback:    commandSave,
		},
		"load": {
			name:        "load <trainer-name>",
			description: "Load a trainer's data and caught Pokemon. Example: load ash",
			callback:    commandLoad,
		},
		"trainer": {
			name:        "trainer",
			description: "View all available trainers that can be loaded",
			callback:    commandTrainer,
		},
	}
}
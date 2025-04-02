# PokedexCLI

A CLI tool for explore the Pokémon map, catch Pokémon, inspect their details.

## Description

PokedexCLI interacts with the [PokéAPI](https://pokeapi.co/).

## Installation

Install the CLI directly using Go:

```bash
go install github.com/FT1006/pokedexcli@latest
```

After installation, the `pokedexcli` binary will be available in your PATH.

## Usage

Once the PokedexCLI is running, you’ll see a prompt: `Pokedex >`. Enter commands as follows:

- **`catch [pokemon name]`**: Try to catch a Pokémon.  
  *Example*: `catch pikachu`

- **`exit`**: Exit the PokedexCLI.  
  *Example*: `exit`

- **`explore [area name]`**: Explore a specific area and see what Pokémon might be lurking there.  
  *Example*: `explore pastoria-city-area`

- **`help`**: Display a help message with available commands.  
  *Example*: `help`

- **`inspect [pokemon name]`**: Inspect details of a Pokémon (e.g., stats, types).  
  *Example*: `inspect pikachu`

- **`map`**: Display the next 20 location areas from the PokéAPI.  
  *Example*: `map`

- **`mapb`**: Display the previous 20 location areas (if applicable).  
  *Example*: `mapb`

- **`pokedex`**: Show the Pokémon you’ve caught so far.  
  *Example*: `pokedex`

### Example Session
```
Pokedex > map
Cache miss! Making API request to: https://pokeapi.co/api/v2/location-area/
canalave-city-area
eterna-city-area
pastoria-city-area
...
Pokedex > catch pikachu
Trying to catch pikachu... Success!
Pokedex > pokedex
Your Pokedex: [pikachu]
Pokedex > exit
Closing the Pokedex... Goodbye!
```

## Branches
- **`main`**: Stable version using `bufio.Scanner` for input handling.
- **`raw-mode-experiment`**: Experimental branch with raw terminal mode (`golang.org/x/term`) for arrow key support.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

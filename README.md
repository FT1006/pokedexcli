# PokedexCLI

Welcome to **PokedexCLI**! This is a command-line interface (CLI) tool built in Go that lets you explore the Pokémon universe, catch Pokémon, inspect their details, and navigate through various locations—all from your terminal!

## Description

PokedexCLI interacts with the [PokéAPI](https://pokeapi.co/) to bring you a lightweight, text-based Pokémon experience. Whether you're catching Pikachu, exploring Pastoria City, or checking your Pokedex, this tool is perfect for Pokémon fans who love the command line.

## Installation

To get started, you’ll need [Go](https://golang.org/dl/) installed on your system. Then follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/FT1006/pokedexcli.git
   cd pokedexcli
   ```

2. Build and run the CLI:
   ```bash
   go run .
   ```

3. (Optional) Install it locally for easy access:
   ```bash
   go install
   pokedexcli
   ```

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

## Features
- Simple, text-based interface.
- Caching of API responses to reduce network requests.
- Command history (in some versions—see branches).
- Extensible command system.

## Contributing
Contributions are welcome! Here’s how you can help:
1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m "Add your feature"`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

Feel free to report bugs or suggest enhancements via [issues](https://github.com/yourusername/pokedexcli/issues).

## Branches
- **`main`**: Stable version using `bufio.Scanner` for input handling.
- **`raw-mode-experiment`**: Experimental branch with raw terminal mode (`golang.org/x/term`) for arrow key support.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Happy Pokémon hunting!

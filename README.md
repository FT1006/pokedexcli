# Pokedex CLI

A command-line interface for exploring and catching Pokemon.

## Features

- Browse location areas from the Pokemon world
- Explore areas to find Pokemon
- Catch Pokemon and add them to your Pokedex
- Inspect your caught Pokemon for details
- Command history with arrow key navigation
- Save and load your trainer data and Pokemon collection

## Database Setup

The app uses PostgreSQL to store trainer and Pokemon data. To set up the database:

1. Make sure PostgreSQL is running locally
2. Run the setup script: `./scripts/setup_db.sh`

Alternatively, you can set a custom database URL with the DATABASE_URL environment variable:

```
export DATABASE_URL="postgresql://username:password@hostname:port/pokedexdbv1?sslmode=disable"
```

## Commands

- `help`: Display available commands
- `exit`: Exit the Pokedex
- `map`: Display the next 20 location areas
- `mapb`: Display the previous 20 location areas
- `explore [area name]`: Explore a location area for Pokemon
- `catch [pokemon name]`: Try to catch a Pokemon
- `inspect [pokemon name]`: View details about a caught Pokemon
- `pokedex`: List all your caught Pokemon
- `save [name]`: Save your trainer data and caught Pokemon (name is optional if already loaded)
- `load [name]`: Load a trainer's data and caught Pokemon

## Development

This project uses Go modules for dependency management.
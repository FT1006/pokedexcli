# Pokedex CLI

A command-line interface for exploring, catching, and managing Pokemon.

## Features

- Browse location areas from the Pokemon world
- Explore areas to find Pokemon
- Catch Pokemon and add them to your Pokedex
- Manage your active Pokemon party (up to 6 Pokemon)
- View all your caught Pokemon, including duplicates
- Inspect your caught Pokemon for details
- Command history with arrow key navigation
- Save and load your trainer data and Pokemon collection
- Terminal raw mode for improved user experience
- In-memory caching for API responses

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
- `catch [pokemon name]`: Try to catch a Pokemon and add it to your party
- `inspect [pokemon name]`: View details about a caught Pokemon
- `pokedex`: List all your unique caught Pokemon
- `ownpoke`: View all your caught Pokemon, including duplicates with timestamps
- `party`: Display your active party of up to 6 Pokemon
- `save [name]`: Save your trainer data and caught Pokemon (name is optional if already loaded)
- `load [name]`: Load a trainer's data and caught Pokemon

## Party System

The party system allows you to maintain an active group of up to 6 Pokemon:

- When you catch a Pokemon, it's automatically added to your party if you have an open slot
- If your party is full when catching a new Pokemon, you'll be prompted to either:
  - Replace an existing Pokemon in your party by specifying the slot number (1-6)
  - Keep your current party unchanged
- Use the `party` command to view your active Pokemon party

## Technical Architecture

- CLI with REPL interface and raw terminal mode for better UX
- PostgreSQL database with tables for:
  - Trainers
  - Unique Pokemon collection (pokedex)
  - All caught Pokemon with timestamps (ownpoke)
  - Party management (active Pokemon)
- External PokeAPI integration via HTTP client
- In-memory caching layer for API responses
- Service layer for database operations

## Development

This project uses Go modules for dependency management.
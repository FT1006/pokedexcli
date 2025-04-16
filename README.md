# Pokedex CLI

An interactive command-line interface featuring turn-based battles, skill systems, and party management for a complete Pokemon adventure. This Go-based application combines the nostalgic world of Pokemon with sophisticated software architecture.

## Key Technologies

- **Go**: Core language with concurrency via goroutines and channels
- **PostgreSQL**: Persistent storage with JSON capabilities
- **SQLC**: Type-safe SQL query generation
- **pgx/v5**: High-performance PostgreSQL driver with connection pooling
- **PokeAPI**: RESTful data source with custom client and caching layer

## Design & Architecture

- **Evolutionary Architecture**: Incremental feature development with compatibility preservation
- **Service-Based Structure**: Specialized components with clear separation of concerns 
- **Layered Design**: Distinct API, business logic, and data access layers
- **Repository Pattern**: Abstracted data access through specialized service objects
- **Command Pattern**: Encapsulated operations with metadata and callbacks
- **Concurrent Processing**: Parallel data retrieval with channels for API optimization

## Features

### Core Functionality
- Browse location areas from the Pokemon world
- Explore areas to find Pokemon
- Catch Pokemon and add them to your Pokedex
- Inspect your caught Pokemon for details

### Advanced Features
- **Battle System**: Turn-based combat with wild Pokemon that escape catch attempts
- **Party Management**: Maintain a party of up to 6 active Pokemon for battles
- **Skill System**: Pokemon have unique basic and special skills with different damage characteristics
- **Command History**: Navigate previous commands with arrow keys
- **Persistence**: Save and load your trainer data and Pokemon collection
- **Comprehensive Pokemon Management**: View all caught Pokemon, including duplicates and timestamps
- **Enhanced UX**: Terminal raw mode for improved user experience
- **Performance Optimization**: In-memory caching for API responses

## Quick Start

### Running Without Database (In-Memory Mode)

Clone the repository and run:

```bash
# Clone the repository
git clone https://github.com/FT1006/pokedexcli.git
cd pokedexcli

# Run the application (falls back to in-memory mode without database)
go run .
```

### With Database (Full Features)

For full functionality including saving trainers and Pokemon:

1. Make sure PostgreSQL is installed and running
2. Use the provided setup script:

```bash
# Run the setup script (creates database and applies migrations if possible)
./scripts/setup_db.sh

# Run the application with the suggested DATABASE_URL from the script
go run .
```

The script will:
- Create the database (default: pokedexdbv1)
- Check for and run migrations if possible
- Configure the connection string based on your system

You can also customize the database connection:

```bash
# Set custom database parameters (optional)
export DB_NAME=customdbname
export DB_USER=customuser
export DB_HOST=localhost
export DB_PORT=5432

# Run the setup script with custom parameters
./scripts/setup_db.sh

# Or set the connection string directly
export DATABASE_URL="postgresql://username:password@localhost:5432/pokedexdbv1?sslmode=disable"
```

The application will automatically detect the database connection and enable full functionality.

## Commands

### Navigation
- `help`: Display available commands
- `exit`: Exit the Pokedex
- `map`: Display the next 20 location areas
- `mapb`: Display the previous 20 location areas

### Pokemon Interaction
- `explore <area-name>`: Explore a location area for Pokemon
- `catch <pokemon-name>`: Try to catch a Pokemon; may trigger a battle if the Pokemon escapes
- `inspect <pokemon-name>`: View details about a caught Pokemon, including its skills

### Collection Management
- `pokedex`: List all your unique caught Pokemon
- `ownpoke`: View all your caught Pokemon, including duplicates with timestamps
- `party [change]`: Display your active party of up to 6 Pokemon; use with 'change' to modify party members

### Trainer System
- `save [trainer-name]`: Save your trainer data and caught Pokemon (name is optional if already loaded)
- `load <trainer-name>`: Load a trainer's data and caught Pokemon
- `trainer`: View all available trainers that can be loaded

### Information
- `battle`: Display detailed information about the battle system and mechanics

## Party System

The party system allows you to maintain an active group of up to 6 Pokemon:

- When you catch a Pokemon, it's automatically added to your party if you have an open slot
- If your party is full when catching a new Pokemon, you'll be prompted to either:
  - Replace an existing Pokemon in your party by specifying the slot number (1-6)
  - Keep your current party unchanged
- Use the `party` command to view your active Pokemon party
- Pokemon in your party are used in battles when a wild Pokemon escapes a catch attempt

## Battle System

The battle system enables turn-based combat with wild Pokemon:

### Battle Triggers
- When a Pokemon escapes a catch attempt, you're given the option to battle
- A random Pokemon from your party is selected to battle the wild Pokemon

### Battle Mechanics
- **Turn-Based**: Your Pokemon always attacks first, followed by the wild Pokemon
- **Attack Types**:
  - Basic attacks: Always hit, lower damage
  - Special attacks: 50% hit chance, higher damage, can only be used once per battle
- **AI Behavior**: Wild Pokemon use basic attacks (75% chance) or special attacks (25% chance)
- **Battle Outcome**:
  - Win: You catch the Pokemon that previously escaped
  - Lose: The Pokemon escapes

### Battle Formulas
- Basic attack damage: (skill damage + attacker's attack - defender's defense) × 0.1
- Special attack damage: ((skill damage + attacker's special-attack) × 1.5 - defender's defense) × 0.1

### Skill System
- Each Pokemon has a basic skill and a special skill
- Skills have different damage values, types, and classes
- The skill selection algorithm picks appropriate moves based on the Pokemon's type

## Technical Architecture

### Evolutionary Design Approach
This project follows an evolutionary architecture approach, with incremental feature additions that build upon a strong foundation:

1. **Core CLI Implementation**: Basic REPL pattern with command mapping and processing
2. **Command History Enhancement**: Arrow key navigation with raw terminal mode integration
3. **Persistence Layer Addition**: PostgreSQL database with normalized schema design
4. **Party System Implementation**: Slot-based management with constraint enforcement
5. **Skill System Development**: Advanced move selection and categorization algorithms
6. **Battle System Integration**: Turn-based mechanics with stat-based damage calculations

### Pattern Implementations

#### Data Access Layer
- **Database Abstraction**: Wrapper service for database operations and connection management
- **Specialized Service Components**:
  - Database service with connection pooling and explicit transaction support (WithTx)
  - Pokemon service for JSON serialization/deserialization of Pokemon data
  - Party service with slot-based management (1-6) and constraint enforcement
  - Trainer service for profile tracking and basic session management
- **SQL Generation**: Type-safe database access through SQLC-generated Go code
- **Database Migrations**: Progressive schema evolution with numbered migration files (001-005)

#### API Integration
- **Client Abstraction**: Dedicated PokeAPI client with method-based endpoints
- **Quality-Based Selection**: Sophisticated move selection algorithm with:
  - Quality scoring based on power, accuracy, priority, and effect chance
  - Weighted randomization favoring higher-quality moves
  - Top-N filtering to focus on the best moves
- **Smart Categorization**: Classification of moves into basic and special categories
- **Multi-Level Fallbacks**: Graceful degradation for Pokemon with limited movesets:
  - Category-agnostic selection when move categories are imbalanced
  - Default "struggle" move when no viable moves are available

#### Performance Optimizations
- **Concurrent API Fetching**: Parallel data retrieval using goroutines and channels
- **Batch Processing**: Efficient multi-item fetching with GetMoveDetailsBatch
- **Time-Based Cache**: In-memory caching with automatic expiration via reapLoop
- **Connection Pooling**: pgxpool for optimized database connection management
- **Background Prefetching**: Preloading of common moves to reduce API latency

#### User Experience
- **Raw Terminal Mode**: Direct terminal control for improved interactivity
- **Command History**: Arrow-key navigation through previous commands
- **Structured Command System**: Rich metadata with name, description, and callback
- **Interactive Battle UI**: Turn-based interface with attack selection and state tracking
- **Category-Based Help**: Commands organized into logical groups for easier discovery

## Development

This project uses Go modules for dependency management.
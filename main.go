package main

import (
	"fmt"
	"os"

	"github.com/FT1006/pokedexcli/internal/database"
	"github.com/FT1006/pokedexcli/internal/database/service"
	"github.com/FT1006/pokedexcli/internal/models"
	"github.com/FT1006/pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient()

	// Initialize database connection
	connStr := "postgresql://spaceship@localhost:5432/pokedexdbv1?sslmode=disable"
	// Use DATABASE_URL environment variable if available
	if envConnStr := os.Getenv("DATABASE_URL"); envConnStr != "" {
		connStr = envConnStr
	}

	dbService, err := database.NewService(connStr)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		// Fall back to local mode without database
		cfg := Config{
			pokeapiClient:      pokeapiClient,
			caughtPokemon:      make(map[string]models.Pokemon),
			newlyCaughtPokemon: make(map[string]models.Pokemon),
			next:               "",
			prev:               "",
		}
		repl(&cfg)
		return
	}
	defer dbService.Close()

	// Create services
	pokemonService := service.NewPokemonService(dbService)
	trainerService := service.NewTrainerService(dbService)
	partyService := service.NewPartyService(dbService, pokemonService)

	cfg := Config{
		pokeapiClient:      pokeapiClient,
		caughtPokemon:      make(map[string]models.Pokemon),
		newlyCaughtPokemon: make(map[string]models.Pokemon),
		next:               "",
		prev:               "",
		dbService:          dbService,
		pokemonService:     pokemonService,
		trainerService:     trainerService,
		partyService:       partyService,
		currentTrainer:     nil, // Will be set when user saves or loads a trainer
	}

	repl(&cfg)
}
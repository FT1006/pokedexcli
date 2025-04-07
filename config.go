package main

import (
	"github.com/FT1006/pokedexcli/internal/database"
	"github.com/FT1006/pokedexcli/internal/models"
	"github.com/FT1006/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient  *pokeapi.Client
	caughtPokemon  map[string]models.Pokemon
	next           string
	prev           string
	dbService      *database.Service
	pokemonService *database.PokemonService
	trainerService *database.TrainerService
	currentTrainer *database.Trainer
}

package main

import (
	"github.com/FT1006/pokedexcli/internal/database"
	"github.com/FT1006/pokedexcli/internal/database/service"
	"github.com/FT1006/pokedexcli/internal/models"
	"github.com/FT1006/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient     *pokeapi.Client
	caughtPokemon     map[string]models.Pokemon
	newlyCaughtPokemon map[string]models.Pokemon // Tracks Pokemon caught since last save
	next              string
	prev              string
	dbService         *database.Service
	pokemonService    *service.PokemonService
	trainerService    *service.TrainerService
	partyService      *service.PartyService
	currentTrainer    *service.Trainer
}

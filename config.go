package main

import (
	"github.com/FT1006/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient *pokeapi.Client
	caughtPokemon map[string]Pokemon
	next          string
	prev          string
}

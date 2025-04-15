package service

import (
	dbsqlc "github.com/FT1006/pokedexcli/internal/database/sqlc/db"
)

// Type aliases for sqlc generated types
type Ownpoke = dbsqlc.Ownpoke
type Party = dbsqlc.Party
type Pokedex = dbsqlc.Pokedex
type Trainer = dbsqlc.Trainer

// Parameter type aliases
type AddOwnedPokemonParams = dbsqlc.AddOwnedPokemonParams
type AddPokemonToPartyParams = dbsqlc.AddPokemonToPartyParams
type CreatePokedexEntryParams = dbsqlc.CreatePokedexEntryParams
type GetPartyByTrainerRow = dbsqlc.GetPartyByTrainerRow
type GetPartySlotOccupiedParams = dbsqlc.GetPartySlotOccupiedParams
type GetPokedexEntryByNameAndTrainerParams = dbsqlc.GetPokedexEntryByNameAndTrainerParams
type ListTrainersParams = dbsqlc.ListTrainersParams
type UpdateTrainerParams = dbsqlc.UpdateTrainerParams
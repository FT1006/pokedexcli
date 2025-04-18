// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddOwnedPokemon(ctx context.Context, arg AddOwnedPokemonParams) (Ownpoke, error)
	AddPokemonToParty(ctx context.Context, arg AddPokemonToPartyParams) (Party, error)
	CreatePokedexEntry(ctx context.Context, arg CreatePokedexEntryParams) error
	CreateTrainer(ctx context.Context, name pgtype.Text) (Trainer, error)
	DeletePokedexEntry(ctx context.Context, id int32) error
	GetOwnedPokemonByID(ctx context.Context, id int32) (Ownpoke, error)
	GetPartyByTrainer(ctx context.Context, trainerID int32) ([]GetPartyByTrainerRow, error)
	GetPartyCount(ctx context.Context, trainerID int32) (int64, error)
	GetPartySlotOccupied(ctx context.Context, arg GetPartySlotOccupiedParams) (bool, error)
	GetPokedexEntry(ctx context.Context, id int32) (Pokedex, error)
	GetPokedexEntryByNameAndTrainer(ctx context.Context, arg GetPokedexEntryByNameAndTrainerParams) (Pokedex, error)
	GetTrainer(ctx context.Context, id int32) (Trainer, error)
	GetTrainerByName(ctx context.Context, name pgtype.Text) (Trainer, error)
	ListOwnedPokemonByTrainer(ctx context.Context, trainerID int32) ([]Ownpoke, error)
	ListPokedexByTrainer(ctx context.Context, trainerID int32) ([]Pokedex, error)
	ListTrainers(ctx context.Context, arg ListTrainersParams) ([]Trainer, error)
	UpdateOwnedPokemonSkills(ctx context.Context, arg UpdateOwnedPokemonSkillsParams) error
	UpdateTrainer(ctx context.Context, arg UpdateTrainerParams) (Trainer, error)
}

var _ Querier = (*Queries)(nil)

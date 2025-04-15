package service

import (
	"context"
	"fmt"
	"time"

	"github.com/FT1006/pokedexcli/internal/database"
	dbsqlc "github.com/FT1006/pokedexcli/internal/database/sqlc/db"
	"github.com/FT1006/pokedexcli/internal/models"
)

const MAX_PARTY_SIZE = 6

// PartyPokemon represents a pokemon in a trainer's party
type PartyPokemon struct {
	Slot         int
	OwnpokeID    int32
	Name         string
	AddedAt      time.Time
	Height       int
	Weight       int
	Stats        []models.Stats 
	Types        []models.Types
	BaseExperience int
}

type PartyService struct {
	db *database.Service
	pokemonService *PokemonService
}

func NewPartyService(db *database.Service, pokemonService *PokemonService) *PartyService {
	return &PartyService{
		db: db,
		pokemonService: pokemonService,
	}
}

// Add a Pokemon to a trainer's party in the specified slot (1-6)
func (s *PartyService) AddPokemonToParty(ctx context.Context, trainerID int32, ownpokeID int32, slot int) error {
	if slot < 1 || slot > MAX_PARTY_SIZE {
		return fmt.Errorf("invalid slot: must be between 1 and %d", MAX_PARTY_SIZE)
	}

	_, err := s.db.Queries().AddPokemonToParty(ctx, dbsqlc.AddPokemonToPartyParams{
		TrainerID: trainerID,
		OwnpokeID: ownpokeID,
		Slot:     int32(slot),
	})
	if err != nil {
		return fmt.Errorf("error adding Pokemon to party: %w", err)
	}

	return nil
}

// Get a trainer's entire party
func (s *PartyService) GetParty(ctx context.Context, trainerID int32) ([]PartyPokemon, error) {
	partyRows, err := s.db.Queries().GetPartyByTrainer(ctx, trainerID)
	if err != nil {
		return nil, fmt.Errorf("error getting party: %w", err)
	}

	partyPokemon := make([]PartyPokemon, 0, len(partyRows))
	for _, row := range partyRows {
		// Unmarshal stats and types JSON
		stats, types, err := s.pokemonService.UnmarshalStatsAndTypes(row.Stats, row.Types)
		if err != nil {
			return nil, err
		}

		pokemon := PartyPokemon{
			Slot:         int(row.Slot),
			OwnpokeID:    row.OwnpokeID,
			Name:         row.Name,
			AddedAt:      row.AddedAt.Time,
			Height:       int(row.Height),
			Weight:       int(row.Weight),
			Stats:        stats,
			Types:        types,
			BaseExperience: int(row.BaseExperience),
		}

		partyPokemon = append(partyPokemon, pokemon)
	}

	return partyPokemon, nil
}

// Get the number of Pokemon in a trainer's party
func (s *PartyService) GetPartyCount(ctx context.Context, trainerID int32) (int, error) {
	count, err := s.db.Queries().GetPartyCount(ctx, trainerID)
	if err != nil {
		return 0, fmt.Errorf("error getting party count: %w", err)
	}

	return int(count), nil
}

// Check if a specific party slot is occupied
func (s *PartyService) IsSlotOccupied(ctx context.Context, trainerID int32, slot int) (bool, error) {
	if slot < 1 || slot > MAX_PARTY_SIZE {
		return false, fmt.Errorf("invalid slot: must be between 1 and %d", MAX_PARTY_SIZE)
	}

	occupied, err := s.db.Queries().GetPartySlotOccupied(ctx, dbsqlc.GetPartySlotOccupiedParams{
		TrainerID: trainerID,
		Slot:     int32(slot),
	})
	if err != nil {
		return false, fmt.Errorf("error checking slot: %w", err)
	}

	return occupied, nil
}

// Add Pokemon to the next available party slot, or return an error if party is full
func (s *PartyService) AddToNextAvailableSlot(ctx context.Context, trainerID int32, ownpokeID int32) (int, error) {
	count, err := s.GetPartyCount(ctx, trainerID)
	if err != nil {
		return 0, err
	}

	if count >= MAX_PARTY_SIZE {
		return 0, fmt.Errorf("party is full (max %d Pokemon)", MAX_PARTY_SIZE)
	}

	// Find the next empty slot
	for slot := 1; slot <= MAX_PARTY_SIZE; slot++ {
		occupied, err := s.IsSlotOccupied(ctx, trainerID, slot)
		if err != nil {
			return 0, err
		}

		if !occupied {
			err = s.AddPokemonToParty(ctx, trainerID, ownpokeID, slot)
			if err != nil {
				return 0, err
			}
			return slot, nil
		}
	}

	// This should never happen if GetPartyCount works correctly
	return 0, fmt.Errorf("could not find an empty slot despite party count being less than %d", MAX_PARTY_SIZE)
}
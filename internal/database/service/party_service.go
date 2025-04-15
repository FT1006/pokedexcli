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
	BasicSkill   *models.Skill
	SpecialSkill *models.Skill
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
		return fmt.Errorf("invalid party slot: must be between 1 and %d", MAX_PARTY_SIZE)
	}

	params := dbsqlc.AddPokemonToPartyParams{
		TrainerID: trainerID,
		OwnpokeID: ownpokeID,
		Slot:      int32(slot),
	}

	_, err := s.db.Queries().AddPokemonToParty(ctx, params)
	if err != nil {
		return fmt.Errorf("error adding pokemon to party: %w", err)
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

		// Unmarshal skills if available
		var basicSkill, specialSkill *models.Skill
		if len(row.BasicSkill) > 0 {
			basicSkill, err = s.pokemonService.UnmarshalSkill(row.BasicSkill)
			if err != nil {
				return nil, fmt.Errorf("error unmarshaling basic skill: %w", err)
			}
		}
		
		if len(row.SpecialSkill) > 0 {
			specialSkill, err = s.pokemonService.UnmarshalSkill(row.SpecialSkill)
			if err != nil {
				return nil, fmt.Errorf("error unmarshaling special skill: %w", err)
			}
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
			BasicSkill:   basicSkill,
			SpecialSkill: specialSkill,
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

// Check if a party slot is already occupied
func (s *PartyService) IsSlotOccupied(ctx context.Context, trainerID int32, slot int) (bool, error) {
	if slot < 1 || slot > MAX_PARTY_SIZE {
		return false, fmt.Errorf("invalid party slot: must be between 1 and %d", MAX_PARTY_SIZE)
	}
	
	occupied, err := s.db.Queries().GetPartySlotOccupied(ctx, dbsqlc.GetPartySlotOccupiedParams{
		TrainerID: trainerID,
		Slot:      int32(slot),
	})
	if err != nil {
		return false, fmt.Errorf("error checking if slot is occupied: %w", err)
	}
	
	return occupied, nil
}

// Add a Pokemon to the first available slot in the party
func (s *PartyService) AddToNextAvailableSlot(ctx context.Context, trainerID int32, ownpokeID int32) (int, error) {
	// Check if party is full
	count, err := s.GetPartyCount(ctx, trainerID)
	if err != nil {
		return 0, err
	}
	
	if count >= MAX_PARTY_SIZE {
		return 0, fmt.Errorf("party is full")
	}
	
	// Find the first available slot
	for slot := 1; slot <= MAX_PARTY_SIZE; slot++ {
		occupied, err := s.IsSlotOccupied(ctx, trainerID, slot)
		if err != nil {
			return 0, err
		}
		
		if !occupied {
			// Found an empty slot, add the Pokemon
			err = s.AddPokemonToParty(ctx, trainerID, ownpokeID, slot)
			if err != nil {
				return 0, err
			}
			return slot, nil
		}
	}
	
	// This should never happen if GetPartyCount works correctly
	return 0, fmt.Errorf("no available slots found despite party not being full")
}
package service

import (
	"context"
	
	"github.com/FT1006/pokedexcli/internal/database"
	dbsqlc "github.com/FT1006/pokedexcli/internal/database/sqlc/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type TrainerService struct {
	db *database.Service
}

func NewTrainerService(db *database.Service) *TrainerService {
	return &TrainerService{
		db: db,
	}
}

// Create or update a trainer by name
func (s *TrainerService) CreateOrUpdateTrainer(ctx context.Context, name string) (dbsqlc.Trainer, error) {
	nameParam := pgtype.Text{String: name, Valid: true}
	
	// Check if trainer exists
	trainer, err := s.db.Queries().GetTrainerByName(ctx, nameParam)
	if err == nil {
		// Trainer exists, return it
		return trainer, nil
	}

	// Trainer doesn't exist, create a new one
	return s.db.Queries().CreateTrainer(ctx, nameParam)
}

// Get a trainer by ID
func (s *TrainerService) GetTrainer(ctx context.Context, id int32) (dbsqlc.Trainer, error) {
	return s.db.Queries().GetTrainer(ctx, id)
}
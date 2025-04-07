package database

import (
	"context"
	
	"github.com/jackc/pgx/v5/pgtype"
)

type TrainerService struct {
	db *Service
}

func NewTrainerService(db *Service) *TrainerService {
	return &TrainerService{
		db: db,
	}
}

// Create or update a trainer by name
func (s *TrainerService) CreateOrUpdateTrainer(ctx context.Context, name string) (Trainer, error) {
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
func (s *TrainerService) GetTrainer(ctx context.Context, id int32) (Trainer, error) {
	return s.db.Queries().GetTrainer(ctx, id)
}
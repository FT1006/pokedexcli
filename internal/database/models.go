// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Pokemon struct {
	ID             int32            `json:"id"`
	TrainerID      int32            `json:"trainer_id"`
	Name           string           `json:"name"`
	Height         int32            `json:"height"`
	Weight         int32            `json:"weight"`
	BaseExperience int32            `json:"base_experience"`
	Stats          []byte           `json:"stats"`
	Types          []byte           `json:"types"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
}

type Trainer struct {
	ID        int32            `json:"id"`
	Name      pgtype.Text      `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

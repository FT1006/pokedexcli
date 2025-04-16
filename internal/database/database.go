package database

import (
	"context"
	"fmt"

	dbsqlc "github.com/FT1006/pokedexcli/internal/database/sqlc/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	db  *pgxpool.Pool
	dbq *dbsqlc.Queries
}

func NewService(connStr string) (*Service, error) {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing database connection string: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	// Verify connection
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	queries := dbsqlc.New(pool)

	return &Service{
		db:  pool,
		dbq: queries,
	}, nil
}

func (s *Service) Close() {
	if s.db != nil {
		s.db.Close()
	}
}

func (s *Service) Queries() *dbsqlc.Queries {
	return s.dbq
}

// Runs a function within a transaction
func (s *Service) WithTx(ctx context.Context, fn func(*dbsqlc.Queries) error) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}

	q := s.dbq.WithTx(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("error rolling back transaction: %v, original error: %w", rbErr, err)
		}
		return err
	}

	return tx.Commit(ctx)
}

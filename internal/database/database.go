package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	db  *pgxpool.Pool
	dbq *Queries
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

	queries := New(pool)

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

func (s *Service) Queries() *Queries {
	return s.dbq
}

// Runs a function within a transaction
func (s *Service) WithTx(ctx context.Context, fn func(*Queries) error) error {
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

// Execute migrations on the database - this would use goose in a real application
func (s *Service) Migrate() error {
	// In a real application, this would run migrations with goose
	// For simplicity, we're skipping this step and assuming migrations are run manually
	return nil
}
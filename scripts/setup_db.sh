#!/bin/bash
set -e

# Create the database if it doesn't exist
createdb pokedexdbv1 || echo "Database already exists"

# Run migrations
cd /Users/spaceship/project/pokedexcli
goose -dir internal/database/migrations postgres "postgresql://spaceship@localhost:5432/pokedexdbv1?sslmode=disable" up

echo "Database setup complete"
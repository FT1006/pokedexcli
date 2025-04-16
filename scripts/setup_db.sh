#!/bin/bash
set -e

# Check if psql is installed
if ! command -v psql &> /dev/null; then
    echo "PostgreSQL client (psql) is not installed. Please install PostgreSQL first."
    exit 1
fi

# Set database connection parameters (use environment variables if provided)
DB_NAME=${DB_NAME:-pokedexdbv1}
DB_USER=${DB_USER:-$(whoami)}
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
CONN_STRING="postgresql://${DB_USER}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

# Display the configuration
echo "Database setup configuration:"
echo " - Database: $DB_NAME"
echo " - User: $DB_USER"
echo " - Host: $DB_HOST:$DB_PORT"

# Create the database if it doesn't exist
if psql -lqt | cut -d \| -f 1 | grep -qw $DB_NAME; then
    echo "Database $DB_NAME already exists"
else
    echo "Creating database $DB_NAME..."
    createdb $DB_NAME
fi

# Check if goose is installed
if ! command -v goose &> /dev/null; then
    echo "Goose migration tool is not installed."
    echo "To run migrations, please install goose: go install github.com/pressly/goose/v3/cmd/goose@latest"
    echo "Then run: goose -dir internal/database/migrations postgres \"$CONN_STRING\" up"
    echo ""
    echo "Database created but migrations not applied."
    echo "You can still use the application with the DATABASE_URL environment variable:"
    echo "export DATABASE_URL=\"$CONN_STRING\""
    exit 0
fi

# Run migrations with goose if it's installed
echo "Running database migrations..."
# Use the directory relative to the script location
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$( cd "$SCRIPT_DIR/.." && pwd )"

goose -dir "$PROJECT_ROOT/internal/database/migrations" postgres "$CONN_STRING" up

echo "Database setup complete!"
echo "You can now run the application with:"
echo "export DATABASE_URL=\"$CONN_STRING\""
echo "go run ."
-- +goose Up
-- +goose StatementBegin
CREATE TABLE pokedex (
    id SERIAL PRIMARY KEY,
    trainer_id INTEGER NOT NULL REFERENCES trainers(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    height INTEGER NOT NULL,
    weight INTEGER NOT NULL,
    base_experience INTEGER NOT NULL,
    stats JSONB NOT NULL,
    types JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(trainer_id, name)
);

CREATE INDEX pokedex_trainer_id_idx ON pokedex(trainer_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pokedex;
-- +goose StatementEnd
-- +goose Up
-- +goose StatementBegin
CREATE TABLE trainers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

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

CREATE TABLE ownpoke (
    id SERIAL PRIMARY KEY,
    trainer_id INTEGER NOT NULL REFERENCES trainers(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    height INTEGER NOT NULL,
    weight INTEGER NOT NULL,
    base_experience INTEGER NOT NULL,
    stats JSONB NOT NULL,
    types JSONB NOT NULL,
    caught_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX pokedex_trainer_id_idx ON pokedex(trainer_id);
CREATE INDEX ownpoke_trainer_id_idx ON ownpoke(trainer_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ownpoke;
DROP TABLE pokedex;
DROP TABLE trainers;
-- +goose StatementEnd
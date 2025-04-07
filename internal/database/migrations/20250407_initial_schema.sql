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

CREATE TABLE party (
    id SERIAL PRIMARY KEY,
    trainer_id INTEGER NOT NULL REFERENCES trainers(id) ON DELETE CASCADE,
    ownpoke_id INTEGER NOT NULL REFERENCES ownpoke(id) ON DELETE CASCADE,
    slot INTEGER NOT NULL CHECK (slot >= 1 AND slot <= 6), -- Enforce slot range 1-6
    added_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(trainer_id, slot) -- Each slot can only have one Pokemon
);

CREATE INDEX pokedex_trainer_id_idx ON pokedex(trainer_id);
CREATE INDEX ownpoke_trainer_id_idx ON ownpoke(trainer_id);
CREATE INDEX party_trainer_id_idx ON party(trainer_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE party;
DROP TABLE ownpoke;
DROP TABLE pokedex;
DROP TABLE trainers;
-- +goose StatementEnd
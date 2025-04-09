-- +goose Up
-- +goose StatementBegin
CREATE TABLE party (
    id SERIAL PRIMARY KEY,
    trainer_id INTEGER NOT NULL REFERENCES trainers(id) ON DELETE CASCADE,
    ownpoke_id INTEGER NOT NULL REFERENCES ownpoke(id) ON DELETE CASCADE,
    slot INTEGER NOT NULL CHECK (slot >= 1 AND slot <= 6), -- Enforce slot range 1-6
    added_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(trainer_id, slot) -- Each slot can only have one Pokemon
);

CREATE INDEX party_trainer_id_idx ON party(trainer_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE party;
-- +goose StatementEnd
-- +goose Up
-- +goose StatementBegin
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

CREATE INDEX ownpoke_trainer_id_idx ON ownpoke(trainer_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ownpoke;
-- +goose StatementEnd
-- +goose Up
-- +goose StatementBegin
CREATE TABLE trainers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE trainers;
-- +goose StatementEnd
-- +goose Up
-- +goose StatementBegin
ALTER TABLE ownpoke 
    ADD COLUMN basic_skill JSONB,
    ADD COLUMN special_skill JSONB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE ownpoke 
    DROP COLUMN basic_skill,
    DROP COLUMN special_skill;
-- +goose StatementEnd
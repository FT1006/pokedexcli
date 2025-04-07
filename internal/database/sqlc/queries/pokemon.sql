-- name: CreatePokedexEntry :exec
INSERT INTO pokedex (
    trainer_id,
    name,
    height,
    weight,
    base_experience,
    stats,
    types
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) 
ON CONFLICT (trainer_id, name) DO NOTHING;

-- name: GetPokedexEntry :one
SELECT * FROM pokedex
WHERE id = $1 LIMIT 1;

-- name: GetPokedexEntryByNameAndTrainer :one
SELECT * FROM pokedex
WHERE name = $1 AND trainer_id = $2 LIMIT 1;

-- name: ListPokedexByTrainer :many
SELECT * FROM pokedex
WHERE trainer_id = $1
ORDER BY id;

-- name: DeletePokedexEntry :exec
DELETE FROM pokedex
WHERE id = $1;

-- name: AddOwnedPokemon :one
INSERT INTO ownpoke (
    trainer_id,
    name,
    height,
    weight,
    base_experience,
    stats,
    types
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: ListOwnedPokemonByTrainer :many
SELECT * FROM ownpoke
WHERE trainer_id = $1
ORDER BY caught_at DESC;
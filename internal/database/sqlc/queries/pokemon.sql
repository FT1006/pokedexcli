-- name: CreatePokemon :one
INSERT INTO pokemon (
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

-- name: GetPokemon :one
SELECT * FROM pokemon
WHERE id = $1 LIMIT 1;

-- name: GetPokemonByNameAndTrainer :one
SELECT * FROM pokemon
WHERE name = $1 AND trainer_id = $2 LIMIT 1;

-- name: ListPokemonByTrainer :many
SELECT * FROM pokemon
WHERE trainer_id = $1
ORDER BY id;

-- name: DeletePokemon :exec
DELETE FROM pokemon
WHERE id = $1;
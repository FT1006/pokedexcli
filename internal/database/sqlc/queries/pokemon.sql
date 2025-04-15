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
    types,
    basic_skill,
    special_skill
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: ListOwnedPokemonByTrainer :many
SELECT * FROM ownpoke
WHERE trainer_id = $1
ORDER BY caught_at DESC;

-- name: AddPokemonToParty :one
INSERT INTO party (
    trainer_id,
    ownpoke_id,
    slot
) VALUES (
    $1, $2, $3
)
ON CONFLICT (trainer_id, slot) 
DO UPDATE SET ownpoke_id = EXCLUDED.ownpoke_id, added_at = NOW()
RETURNING *;

-- name: GetPartyByTrainer :many
SELECT p.*, o.name, o.stats, o.types, o.height, o.weight, o.base_experience, o.basic_skill, o.special_skill
FROM party p
JOIN ownpoke o ON p.ownpoke_id = o.id
WHERE p.trainer_id = $1
ORDER BY p.slot;

-- name: GetPartyCount :one
SELECT COUNT(*) FROM party
WHERE trainer_id = $1;

-- name: GetPartySlotOccupied :one
SELECT EXISTS(
    SELECT 1 FROM party
    WHERE trainer_id = $1 AND slot = $2
);
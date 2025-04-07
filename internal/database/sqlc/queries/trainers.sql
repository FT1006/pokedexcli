-- name: CreateTrainer :one
INSERT INTO trainers (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetTrainer :one
SELECT * FROM trainers
WHERE id = $1 LIMIT 1;

-- name: GetTrainerByName :one
SELECT * FROM trainers
WHERE name = $1 LIMIT 1;

-- name: ListTrainers :many
SELECT * FROM trainers
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateTrainer :one
UPDATE trainers 
SET name = $2, updated_at = NOW()
WHERE id = $1
RETURNING *;
-- name: GetAllWishes :many
SELECT id, name, wish_description, created_at FROM wishes WHERE deleted_at IS NULL ORDER BY created_at DESC;

-- name: GetWish :one
SELECT id, name, wish_description, created_at FROM wishes WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: CreateWish :one
INSERT INTO wishes (
    name, wish_description
) VALUES (
             $1, $2
         )
    RETURNING *;

-- name: DeleteWish :exec
UPDATE wishes SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL;
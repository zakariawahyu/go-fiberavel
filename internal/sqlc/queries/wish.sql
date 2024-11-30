-- name: GetAllWish :many
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

-- name: DatatablesWish :many
SELECT id, name, wish_description
FROM wishes
WHERE (name ILIKE '%' || $1::text || '%' OR wish_description ILIKE '%' || $1::text || '%') AND deleted_at IS NULL
ORDER BY (case when $2 = 'name' and $3 = 'asc' then name end) ASC,
         (case when $2 = 'name' and $3 = 'desc' then name end) DESC,
         (case when $2 = 'wish_description' and $3 = 'asc' then wish_description end) ASC,
         (case when $2 = 'wish_description' and $3 = 'desc' then wish_description end) DESC,
         (case when $2 = '' then created_at end) DESC
LIMIT $4 OFFSET $5;

-- name: CountWish :one
SELECT COUNT(id)
FROM wishes
WHERE (name ILIKE '%' || $1::text || '%' OR wish_description ILIKE '%' || $1::text || '%') AND deleted_at IS NULL;
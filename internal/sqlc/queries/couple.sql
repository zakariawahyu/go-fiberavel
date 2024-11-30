-- name: GetAllCouple :many
SELECT id, couple_type, name, parent_description, father_name, mother_name, image, image_caption, instagram_url FROM couples WHERE deleted_at IS NULL ORDER BY created_at DESC;

-- name: GetCouple :one
SELECT id, couple_type, name, parent_description, father_name, mother_name, image, image_caption, instagram_url FROM couples WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: CreateCouple :one
INSERT INTO couples (
    couple_type, name, parent_description, father_name, mother_name, image, image_caption, instagram_url
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8
         )
    RETURNING *;

-- name: UpdateCouple :exec
UPDATE couples SET
    updated_at = NOW(),
    couple_type = $2,
    name = $3,
    parent_description = $4,
    father_name = $5,
    mother_name = $6,
    image = COALESCE($7, image),
    image_caption = $8,
    instagram_url = $9
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteCouple :exec
UPDATE couples SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL;

-- name: DatatablesCouple :many
SELECT id, couple_type, name
FROM couples
WHERE (couple_type ILIKE '%' || $1::text || '%' OR name ILIKE '%' || $1::text || '%') AND deleted_at IS NULL
ORDER BY (case when $2 = 'couple_type' and $3 = 'asc' then couple_type end) ASC,
         (case when $2 = 'couple_type' and $3 = 'desc' then couple_type end) DESC,
         (case when $2 = 'name' and $3 = 'asc' then name end) ASC,
         (case when $2 = 'name' and $3 = 'desc' then name end) DESC,
         (case when $2 = '' then created_at end) DESC
LIMIT $4 OFFSET $5;

-- name: CountCouple :one
SELECT COUNT(*)
FROM couples
WHERE (couple_type ILIKE '%' || $1::text || '%' OR name ILIKE '%' || $1::text || '%') AND deleted_at IS NULL;
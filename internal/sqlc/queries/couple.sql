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
    couple_type = $2,
    name = $3,
    parent_description = $4,
    father_name = $5,
    mother_name = $6,
    image = $7,
    image_caption = $8,
    instagram_url = $9
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteCouple :exec
UPDATE couples SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL;
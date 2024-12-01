-- name: GetAllGallery :many
SELECT id, image, image_caption FROM galleries WHERE deleted_at IS NULL;

-- name: GetGallery :one
SELECT id, image, image_caption FROM galleries WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: CreateGallery :one
INSERT INTO galleries (
    image, image_caption
) VALUES (
             $1, $2
         )
    RETURNING *;

-- name: UpdateGallery :exec
UPDATE galleries SET
    updated_at = NOW(),
    image = COALESCE($2, image),
    image_caption = $3
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteGallery :exec
UPDATE galleries SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL;

-- name: DatatablesGallery :many
SELECT id, image, image_caption
FROM galleries
WHERE (image_caption ILIKE '%' || $1::text || '%') AND deleted_at IS NULL
ORDER BY (case when $2 = 'image_caption' and $3 = 'asc' then image_caption end) ASC,
         (case when $2 = 'image_caption' and $3 = 'desc' then image_caption end) DESC,
         (case when $2 = '' then created_at end) DESC
LIMIT $4 OFFSET $5;

-- name: CountGallery :one
SELECT COUNT(id)
FROM galleries
WHERE (image_caption ILIKE '%' || $1::text || '%') AND deleted_at IS NULL;
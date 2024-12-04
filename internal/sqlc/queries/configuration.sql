-- name: GetConfigurationByType :one
SELECT id, type, title, description, image, image_caption, custom_data->'custom_data' as custom_data, is_active FROM configurations WHERE type = $1;

-- name: CreateConfiguration :exec
INSERT INTO configurations (
    type, title, description, image, image_caption, custom_data
) VALUES (
             $1, $2, $3, $4, $5, $6
         );

-- name: UpdateConfiguration :exec
UPDATE configurations SET
    title = $2,
    description = $3,
    image = COALESCE($4, image),
    image_caption = $5,
    custom_data = $6,
    updated_at = NOW()
WHERE type = $1;
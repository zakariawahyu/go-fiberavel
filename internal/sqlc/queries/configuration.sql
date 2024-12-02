-- name: GetConfigurationByType :one
SELECT id, type, title, description, image, image_caption, custom_data, is_active FROM configurations WHERE type = $1;

-- name: UpdateConfigurationCover :exec
UPDATE configurations
SET title = $2, description = $3, custom_data = $4
WHERE type = $1;
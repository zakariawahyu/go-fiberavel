-- name: GetConfigurationByType :one
SELECT id, type, title, description, image, image_caption, custom_data->'custom_data' as custom_data, is_active FROM configurations WHERE type = $1;

-- name: GetAllTypeConfigurations :many
SELECT id, type, is_active FROM configurations;

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

-- name: BulkUpdateIsActiveConfiguration :exec
UPDATE configurations
SET updated_at = NOW(),
    is_active = CASE
    WHEN type = ANY($1::text[]) THEN TRUE
    WHEN type != ANY($2::text[]) THEN FALSE
    ELSE is_active
    END

WHERE type = ANY($1::text[]) OR type != ANY($2::text[]);
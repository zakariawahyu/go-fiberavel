-- name: GetConfigurationByType :one
SELECT id, type, title, description, image, image_caption, custom_data->'custom_data' as custom_data, is_active FROM configurations WHERE type = $1;

-- name: UpdateConfigurationCover :exec
INSERT INTO configurations (type, title, description, custom_data, updated_at)
VALUES ($1, $2, $3, $4, NOW())
ON CONFLICT (type) DO UPDATE
SET
        title = EXCLUDED.title,
        description = EXCLUDED.description,
        custom_data = EXCLUDED.custom_data,
        updated_at = NOW();

-- name: UpdateConfigurationVenue :exec
INSERT INTO configurations (type, title, description, updated_at)
VALUES ($1, $2, $3, NOW())
ON CONFLICT (type) DO UPDATE
SET
        title = EXCLUDED.title,
        description = EXCLUDED.description,
        updated_at = NOW();

-- name: UpdateConfigurationGift :exec
INSERT INTO configurations (type, title, description, updated_at)
VALUES ($1, $2, $3, NOW())
ON CONFLICT (type) DO UPDATE
SET
        title = EXCLUDED.title,
        description = EXCLUDED.description,
        updated_at = NOW();

-- name: UpdateConfigurationWish :exec
INSERT INTO configurations (type, title, description, updated_at)
VALUES ($1, $2, $3, NOW())
ON CONFLICT (type) DO UPDATE
SET
        title = EXCLUDED.title,
        description = EXCLUDED.description,
        updated_at = NOW();
-- name: GetAllVenue :many
SELECT id, name, location, address, date_held, map FROM venues WHERE deleted_at IS NULL;

-- name: GetVenue :one
SELECT id, name, location, address, date_held, map FROM venues WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: CreateVenue :one
INSERT INTO venues (
    name, location, address, date_held, map
) VALUES (
             $1, $2, $3, $4, $5
         )
    RETURNING *;

-- name: UpdateVenue :exec
UPDATE venues SET
    updated_at = NOW(),
    name = $2,
    location = $3,
    address = $4,
    date_held = $5,
    map = $6
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteVenue :exec
UPDATE venues SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL;

-- name: DatatablesVenue :many
SELECT id, name, location, date_held
FROM venues
WHERE (name ILIKE '%' || $1::text || '%' OR location ILIKE '%' || $1::text || '%' OR date_held ILIKE '%' || $1::text || '%') AND deleted_at IS NULL
ORDER BY (case when $2 = 'name' and $3 = 'asc' then name end) ASC,
         (case when $2 = 'name' and $3 = 'desc' then name end) DESC,
         (case when $2 = 'location' and $3 = 'asc' then location end) ASC,
         (case when $2 = 'location' and $3 = 'desc' then location end) DESC,
         (case when $2 = 'date_held' and $3 = 'asc' then date_held end) ASC,
         (case when $2 = 'date_held' and $3 = 'desc' then date_held end) DESC,
         (case when $2 = '' then created_at end) DESC
LIMIT $4 OFFSET $5;

-- name: CountVenue :one
SELECT COUNT(id)
FROM venues
WHERE (name ILIKE '%' || $1::text || '%' OR location ILIKE '%' || $1::text || '%' OR date_held ILIKE '%' || $1::text || '%') AND deleted_at IS NULL;
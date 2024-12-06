-- name: GetAllGuest :many
SELECT id, name, slug, is_gift FROM guests WHERE deleted_at IS NULL;

-- name: GetGuest :one
SELECT id, name, slug, is_gift FROM guests WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: CreateGuest :exec
INSERT INTO guests (
    name, slug, is_gift
) VALUES (
     $1, $2, $3
 );

-- name: UpdateGuest :exec
UPDATE guests SET
    updated_at = NOW(),
    name = $2,
    slug = $3,
    is_gift = $4
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteGuest :exec
UPDATE guests SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL;

-- name: DatatablesGuest :many
SELECT id, name, slug, is_gift
FROM guests
WHERE (name ILIKE '%' || $1::text || '%' OR slug ILIKE '%' || $1::text || '%') AND deleted_at IS NULL
ORDER BY (case when $2 = 'name' and $3 = 'asc' then name end) ASC,
         (case when $2 = 'name' and $3 = 'desc' then name end) DESC,
         (case when $2 = 'slug' and $3 = 'asc' then slug end) ASC,
         (case when $2 = 'slug' and $3 = 'desc' then slug end) DESC,
         (case when $2 = 'is_gift' and $3 = 'asc' then is_gift end) ASC,
         (case when $2 = 'is_gift' and $3 = 'desc' then is_gift end) DESC,
         (case when $2 = '' then created_at end) DESC
LIMIT $4 OFFSET $5;

-- name: CountGuest :one
SELECT COUNT(id)
FROM guests
WHERE (name ILIKE '%' || $1::text || '%' OR slug ILIKE '%' || $1::text || '%') AND deleted_at IS NULL;
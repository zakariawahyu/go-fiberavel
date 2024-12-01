-- name: GetAllGift :many
SELECT id, bank, account_name, account_number FROM gifts WHERE deleted_at IS NULL;

-- name: GetGift :one
SELECT id, bank, account_name, account_number FROM gifts WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: CreateGift :one
INSERT INTO gifts (
    bank, account_name, account_number
) VALUES (
             $1, $2, $3
         )
    RETURNING *;

-- name: UpdateGift :exec
UPDATE gifts SET
    updated_at = NOW(),
    bank = $2,
    account_name = $3,
    account_number = $4
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteGift :exec
UPDATE gifts SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL;

-- name: DatatablesGift :many
SELECT id, bank, account_name, account_number
FROM gifts
WHERE (bank ILIKE '%' || $1::text || '%' OR account_name ILIKE '%' || $1::text || '%') AND deleted_at IS NULL
ORDER BY (case when $2 = 'bank' and $3 = 'asc' then bank end) ASC,
         (case when $2 = 'bank' and $3 = 'desc' then bank end) DESC,
         (case when $2 = 'account_name' and $3 = 'asc' then account_name end) ASC,
         (case when $2 = 'account_name' and $3 = 'desc' then account_name end) DESC,
         (case when $2 = '' then created_at end) DESC
LIMIT $4 OFFSET $5;

-- name: CountGift :one
SELECT COUNT(id)
FROM gifts
WHERE (bank ILIKE '%' || $1::text || '%' OR account_name ILIKE '%' || $1::text || '%') AND deleted_at IS NULL;
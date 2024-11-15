-- name: Login :one
SELECT username, password FROM admins WHERE username = $1 LIMIT 1;
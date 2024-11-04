-- name: CreateRsvp :one
INSERT INTO rsvp (
    name, phone_number, guest_amount, is_attend
) VALUES (
             $1, $2, $3, $4
         )
    RETURNING *;

-- name: CreateEntry :one
INSERT INTO entries (account_id, amount) values ($1, $2) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries where account_id = $1 order by id LIMIT $2 offset $3;
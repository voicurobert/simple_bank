-- name: CreateAccount :one
INSERT INTO accounts (owner, balance, currency) values ($1, $2, $3) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts order by id LIMIT $1 offset $2;

-- name: UpdateAccount :exec
UPDATE accounts SET balance = $2 WHERE id = $1 RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;
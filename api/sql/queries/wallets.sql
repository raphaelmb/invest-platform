-- name: CreateWallet :one
INSERT INTO wallet(id, created_at, updated_at)
VALUES($1, $2, $3)
RETURNING *;

-- name: GetAllWallets :many
SELECT * FROM wallet;
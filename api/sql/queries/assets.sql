-- name: CreateAsset :one
INSERT INTO assets(id, symbol, price, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAllAssets :many
SELECT * FROM assets;
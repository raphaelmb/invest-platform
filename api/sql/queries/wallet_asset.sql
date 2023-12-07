-- name: CreateWalletAsset :one
INSERT INTO wallet_asset(id, amount, created_at, updated_at, wallet_id, asset_id, shares)
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetAllWalletAsset :many
SELECT wa.*, a.symbol AS asset_symbol, a.price AS asset_price
FROM wallet_asset AS wa
JOIN assets as a ON wa.asset_id = a.id
WHERE wa.wallet_id = $1;

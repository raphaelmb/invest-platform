-- name: CreateOrder :one
INSERT INTO "order"(asset_id, wallet_id, shares, price, type, status, partial)
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
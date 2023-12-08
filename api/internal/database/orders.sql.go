// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: orders.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO "order"(asset_id, wallet_id, shares, price, type, status, partial)
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING id, type, status, shares, price, partial, created_at, updated_at, wallet_id, asset_id
`

type CreateOrderParams struct {
	AssetID  uuid.UUID
	WalletID uuid.UUID
	Shares   int32
	Price    string
	Type     OrderType
	Status   OrderStatus
	Partial  int32
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.AssetID,
		arg.WalletID,
		arg.Shares,
		arg.Price,
		arg.Type,
		arg.Status,
		arg.Partial,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Status,
		&i.Shares,
		&i.Price,
		&i.Partial,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.WalletID,
		&i.AssetID,
	)
	return i, err
}

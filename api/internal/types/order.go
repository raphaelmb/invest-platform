package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/raphaelmb/invest-platform/api/internal/database"
)

type Order struct {
	ID        uuid.UUID            `json:"id"`
	AssetID   uuid.UUID            `json:"asset_id"`
	WalletID  uuid.UUID            `json:"wallet_id"`
	Shares    int32                `json:"shares"`
	Price     string               `json:"price"`
	Type      database.OrderType   `json:"type"`
	Status    database.OrderStatus `json:"status"`
	Partial   int32                `json:"partial"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

func DatabaseOrderToOrder(dbOrder database.Order) Order {
	return Order{
		ID:        dbOrder.ID,
		AssetID:   dbOrder.AssetID,
		WalletID:  dbOrder.WalletID,
		Shares:    dbOrder.Shares,
		Price:     dbOrder.Price,
		Type:      dbOrder.Type,
		Status:    dbOrder.Status,
		Partial:   dbOrder.Partial,
		CreatedAt: dbOrder.CreatedAt,
		UpdatedAt: dbOrder.UpdatedAt,
	}
}

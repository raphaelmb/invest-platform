package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/raphaelmb/invest-platform/api/internal/database"
)

type WalletAsset struct {
	ID        uuid.UUID `json:"id"`
	Amount    string    `json:"amount"`
	Shares    string    `json:"shares"`
	WalletID  uuid.UUID `json:"wallet_id"`
	AssetID   uuid.UUID `json:"asset_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DatabaseWalletAssetToWalletAsset(dbWallet database.WalletAsset) WalletAsset {
	return WalletAsset{
		ID:        dbWallet.ID,
		CreatedAt: dbWallet.CreatedAt,
		UpdatedAt: dbWallet.UpdatedAt,
	}
}

//TODO: create mapper to get all wallet asset

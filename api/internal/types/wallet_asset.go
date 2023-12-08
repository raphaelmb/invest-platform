package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/raphaelmb/invest-platform/api/internal/database"
)

type WalletAsset struct {
	ID        uuid.UUID `json:"id"`
	Amount    string    `json:"amount"`
	Shares    int32     `json:"shares"`
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

type WalletAssetRow struct {
	ID          uuid.UUID `json:"id"`
	Amount      string    `json:"amount"`
	Shares      int32     `json:"shares"`
	WalletID    uuid.UUID `json:"wallet_id"`
	AssetID     uuid.UUID `json:"asset_id"`
	AssetSymbol string    `json:"asset_symbol"`
	AssetPrice  string    `json:"asset_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func DatabaseWalletAssetRowToWalletAssetRow(dbWalletAssetRow database.GetAllWalletAssetRow) WalletAssetRow {
	return WalletAssetRow{
		ID:          dbWalletAssetRow.ID,
		Amount:      dbWalletAssetRow.Amount,
		Shares:      dbWalletAssetRow.Shares,
		WalletID:    dbWalletAssetRow.WalletID,
		AssetID:     dbWalletAssetRow.AssetID,
		AssetSymbol: dbWalletAssetRow.AssetSymbol,
		AssetPrice:  dbWalletAssetRow.AssetPrice,
		CreatedAt:   dbWalletAssetRow.CreatedAt,
		UpdatedAt:   dbWalletAssetRow.UpdatedAt,
	}
}

func DatabaseWalletAssetsToWalletAssets(dbWalletAssets []database.GetAllWalletAssetRow) []WalletAssetRow {
	var walletAssets []WalletAssetRow
	for _, dbWalletAsset := range dbWalletAssets {
		walletAssets = append(walletAssets, DatabaseWalletAssetRowToWalletAssetRow(dbWalletAsset))
	}
	return walletAssets
}

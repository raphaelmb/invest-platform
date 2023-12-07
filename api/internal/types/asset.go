package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/raphaelmb/invest-platform/api/internal/database"
)

type Asset struct {
	ID        uuid.UUID `json:"id"`
	Symbol    string    `json:"symbol"`
	Price     string    `json:"string"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DatabaseAssetToAsset(dbAsset database.Asset) Asset {
	return Asset{
		ID:        dbAsset.ID,
		Symbol:    dbAsset.Symbol,
		Price:     dbAsset.Price,
		CreatedAt: dbAsset.CreatedAt,
		UpdatedAt: dbAsset.UpdatedAt,
	}
}

func DatabaseAssetsToAssets(dbAssets []database.Asset) []Asset {
	var assets []Asset
	for _, dbAsset := range dbAssets {
		assets = append(assets, DatabaseAssetToAsset(dbAsset))
	}
	return assets
}

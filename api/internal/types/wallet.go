package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/raphaelmb/invest-platform/api/internal/database"
)

type Wallet struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DatabaseWalletToWallet(dbWallet database.Wallet) Wallet {
	return Wallet{
		ID:        dbWallet.ID,
		CreatedAt: dbWallet.CreatedAt,
		UpdatedAt: dbWallet.UpdatedAt,
	}
}

func DatabaseWalletsToWallets(dbWallets []database.Wallet) []Wallet {
	var wallets []Wallet
	for _, dbWallet := range dbWallets {
		wallets = append(wallets, DatabaseWalletToWallet(dbWallet))
	}
	return wallets
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/raphaelmb/invest-platform/api/helper"
	"github.com/raphaelmb/invest-platform/api/internal/database"
	"github.com/raphaelmb/invest-platform/api/internal/types"
)

func (apiCfg *apiConfig) handlerCreateWallet(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID uuid.UUID `json:"id" validate:"required"`
	}
	var params parameters
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	err = helper.Validate(params)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Missing fields: %v", err))
		return
	}

	wallet, err := apiCfg.DB.CreateWallet(r.Context(), database.CreateWalletParams{
		ID:        params.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not create wallet: %v", err))
		return
	}

	helper.RespondWithJSON(w, http.StatusCreated, types.DatabaseWalletToWallet(wallet))
}

func (apiCfg *apiConfig) handlerGetAllWallets(w http.ResponseWriter, r *http.Request) {
	wallets, err := apiCfg.DB.GetAllWallets(r.Context())
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not retrieve wallets: %v", err))
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, types.DatabaseWalletsToWallets(wallets))
}

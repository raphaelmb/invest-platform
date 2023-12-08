package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/raphaelmb/invest-platform/api/helper"
	"github.com/raphaelmb/invest-platform/api/internal/database"
	"github.com/raphaelmb/invest-platform/api/internal/types"
)

func (apiCfg *apiConfig) handlerCreateWalletAsset(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "wallet_id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not parse id: %v", err))
		return
	}

	type parameters struct {
		AssetID uuid.UUID `json:"asset_id" validate:"required"`
		Shares  int32     `json:"shares" validate:"required"`
	}
	var params parameters
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	err = helper.Validate(params)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Missing fields: %v", err))
		return
	}

	walletAsset, err := apiCfg.DB.CreateWalletAsset(r.Context(), database.CreateWalletAssetParams{
		ID:        uuid.New(),
		AssetID:   params.AssetID,
		WalletID:  id,
		Shares:    params.Shares,
		Amount:    "0",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not create wallet asset: %v", err))
		return
	}

	helper.RespondWithJSON(w, http.StatusCreated, types.DatabaseWalletAssetToWalletAsset(walletAsset))
}

func (apiCfg *apiConfig) handlerGetAllWalletAsset(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "wallet_id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not parse id: %v", err))
		return
	}

	walletAssets, err := apiCfg.DB.GetAllWalletAsset(r.Context(), id)
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not retrieve wallet asset: %v", err))
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, types.DatabaseWalletAssetsToWalletAssets(walletAssets))
}

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

func (apiCfg *apiConfig) handlerCreateAsset(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID     uuid.UUID `json:"id" validate:"required"`
		Symbol string    `json:"symbol" validate:"required"`
		Price  string    `json:"price" validate:"required"`
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

	asset, err := apiCfg.DB.CreateAsset(r.Context(), database.CreateAssetParams{
		ID:        params.ID,
		Symbol:    params.Symbol,
		Price:     params.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not create asset: %v", err))
		return
	}

	helper.RespondWithJSON(w, http.StatusCreated, types.DatabaseAssetToAsset(asset))
}

func (apiCfg *apiConfig) handlerGetAllAssets(w http.ResponseWriter, r *http.Request) {
	assets, err := apiCfg.DB.GetAllAssets(r.Context())
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not retrieve assets: %v", err))
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, types.DatabaseAssetsToAssets(assets))
}

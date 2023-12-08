package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/raphaelmb/invest-platform/api/helper"
	"github.com/raphaelmb/invest-platform/api/internal/database"
	"github.com/raphaelmb/invest-platform/api/internal/types"
)

func (apiCfg *apiConfig) handlerCreateOrder(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID    uuid.UUID `json:"id" validate:"required"`
		Price string    `json:"price" validate:"required"`
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

	order, err := apiCfg.DB.CreateOrder(r.Context(), database.CreateOrderParams{})
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not create order: %v", err))
		return
	}

	helper.RespondWithJSON(w, http.StatusCreated, types.DatabaseOrderToOrder(order))
}

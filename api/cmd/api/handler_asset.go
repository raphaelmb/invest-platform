package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/raphaelmb/invest-platform/api/internal/database"
)

func (apiCfg *apiConfig) handlerCreateAsset(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}
	var params parameters
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := apiCfg.DB.CreateAsset(r.Context(), database.CreateAssetParams{
		ID:        uuid.New(),
		Symbol:    params.Symbol,
		Price:     params.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func (apiCfg *apiConfig) handlerGetAllAssets(w http.ResponseWriter, r *http.Request) {
	assets, err := apiCfg.DB.GetAllAssets(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(assets)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

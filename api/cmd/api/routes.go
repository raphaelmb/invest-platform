package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (apiCfg *apiConfig) routes() http.Handler {
	router := chi.NewRouter()

	router.Post("/assets", apiCfg.handlerCreateAsset)
	router.Get("/assets", apiCfg.handlerGetAllAssets)

	return router
}

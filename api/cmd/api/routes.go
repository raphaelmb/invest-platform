package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (apiCfg *apiConfig) routes() http.Handler {
	router := chi.NewRouter()

	router.Post("/assets", apiCfg.handlerCreateAsset)
	router.Get("/assets", apiCfg.handlerGetAllAssets)

	// TODO: refactor routes
	router.Get("/wallets", apiCfg.handlerGetAllWallets)
	router.Post("/wallets", apiCfg.handlerCreateWallet)

	router.Post("/wallets/{wallet_id}/assets", apiCfg.handlerCreateWalletAsset)
	router.Get("/wallets/{wallet_id}/assets", apiCfg.handlerGetAllWalletAsset)

	return router
}

package handlers

import (
	"rishabhsingh2305/spendings-app/controllers"

	"github.com/go-chi/chi/v5"
)

func WalletHandler(router chi.Router) {
	router.Get("/", controllers.CreateNewWalletHandler)
	router.Get("/wallets", controllers.GetWalletListHandler)
	router.Post("/wallet", controllers.CreateNewWalletFromRequest)
}

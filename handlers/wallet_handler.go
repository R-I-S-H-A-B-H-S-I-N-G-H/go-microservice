package handlers

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/controllers"

	"github.com/go-chi/chi/v5"
)

func WalletHandler(router chi.Router) {
	router.Get("/", controllers.CreateNewWalletHandler)
	router.Get("/wallets", controllers.GetWalletListHandler)
	router.Post("/wallet", controllers.CreateNewWalletFromRequest)
}

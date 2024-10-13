package handlers

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/controllers"

	"github.com/go-chi/chi/v5"
)

func EDHandler(router chi.Router) {
	router.Post("/", controllers.EDDecryptController)
}

package handlers

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/controllers"

	"github.com/go-chi/chi/v5"
)

func PixelHandler(router chi.Router) {
	router.Get("/{id}", controllers.PixelCaptureController)
	router.Get("/list", controllers.PixelListController)
	router.Post("/", controllers.PixelSaveController)
}

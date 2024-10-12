package handlers

import (
	"rishabhsingh2305/spendings-app/controllers"

	"github.com/go-chi/chi/v5"
)

func S3Handler(router chi.Router) {
	router.Post("/uptos3", controllers.PushDataToS3Controller)
}

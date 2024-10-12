package controllers

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/services"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/error_util"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/request_util"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func PixelCaptureController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Request method not allowed", http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "id")
	img_pixel := services.PixelCaptureService(id)

	err := request_util.ResponseToImage(&w, r, *img_pixel)
	error_util.Handle("Failed to send image", err)
	if err != nil {
		http.Error(w, "Failed to send image", http.StatusInternalServerError)
	}
}

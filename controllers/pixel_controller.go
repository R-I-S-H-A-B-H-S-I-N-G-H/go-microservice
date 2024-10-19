package controllers

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/services"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/error_util"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/request_util"
	"errors"
	"net/http"
	"strconv"

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

func PixelListController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Request method not allowed", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	page_size, err := strconv.Atoi(r.URL.Query().Get("size"))

	if err != nil {
		err := errors.New("Invalid page or size")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pixel_list, err := services.PixelListService(page, page_size)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	request_util.ResponseToJson(&w, r, pixel_list)
}

func PixelSaveController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Request method not allowed", http.StatusBadRequest)
		return
	}

	// Read the request body as json
	var requestData services.Pixel
	err := request_util.DecodeJson(r.Body, &requestData)
	error_util.Handle("Failed to decode JSON", err)

	// Save pixel
	pixel, err := services.PixelSaveService(requestData, "helixshadcn")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	request_util.ResponseToJson(&w, r, pixel)
}

package request_util

import (
	"encoding/json"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
)

func ResponseToJson(w *http.ResponseWriter, r *http.Request, data any) {
	(*w).Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(*w).Encode(data)
	if err != nil {
		log.Printf("Error encoding data: %v", err)
		http.Error(*w, "Failed to encode data", http.StatusInternalServerError)
	}
}

func ResponseToImage(w *http.ResponseWriter, r *http.Request, img image.RGBA) error {
	(*w).Header().Set("Content-Type", "image/png")
	return png.Encode((*w), &img)
}

func DecodeJson(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

package request_util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ResponseToJson(w* http.ResponseWriter, r* http.Request, data any) {
	(*w).Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(*w).Encode(data)
	if err != nil {
		log.Printf("Error encoding data: %v", err)
		http.Error(*w, "Failed to encode data", http.StatusInternalServerError)
	}
}

func DecodeJson(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
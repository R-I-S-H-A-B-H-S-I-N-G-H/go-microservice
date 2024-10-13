package controllers

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/encryption_util"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/error_util"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/request_util"
	"net/http"
)

func EDDecryptController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Request method not allowed", http.StatusBadRequest)
		return
	}

	// Read the request body as json
	var requestData map[string]string
	err := request_util.DecodeJson(r.Body, &requestData)
	error_util.Handle("Failed to decode JSON", err)

	ed := requestData["ed"]

	// Decrypt ED
	ed_decrypted, err := encryption_util.DecryptData(ed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request_util.ResponseToJson(&w, r, ed_decrypted)
}

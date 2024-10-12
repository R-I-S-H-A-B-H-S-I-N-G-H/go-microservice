package controllers

import (
	"net/http"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/services"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/request_util"
)

type S3Controller struct {
	WalletService *services.S3Service
}

type RequestData struct {
	FilePath string `json:"filePath"`
	FileData string `json:"fileData"`
}

func PushDataToS3Controller(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Request method not allowed", http.StatusBadRequest)
		return
	}

	// Read the request body as json
	var requestData RequestData
	if err := request_util.DecodeJson(r.Body, &requestData); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	// Push data to S3
	res, err := services.PushToS3(requestData.FilePath, requestData.FileData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	request_util.ResponseToJson(&w, r, res)
}

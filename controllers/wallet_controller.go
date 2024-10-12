package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/services"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/request_util"
)

type WalletController struct {
	WalletService *services.WalletService
}

func CreateNewWalletHandler(w http.ResponseWriter, r *http.Request) {
	wallet := services.GetNewWallet()
	request_util.ResponseToJson(&w, r, wallet)
}

func GetWalletListHandler(w http.ResponseWriter, r *http.Request) {
	walletList := services.GetWalletList()
	request_util.ResponseToJson(&w, r, walletList)
}

func CreateNewWalletFromRequest(w http.ResponseWriter, r *http.Request) {
// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close() // Ensure the body is closed after reading

	// Print the request body as a string
	fmt.Printf("Request Body: %s\n", body)

	// Create a new io.Reader from the byte slice
	r.Body = io.NopCloser(bytes.NewReader(body))

	// Decode the JSON into the Wallet struct
	var wallet services.Wallet
	if err := json.Unmarshal(body, &wallet); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	// Log the decoded wallet
	log.Printf("Decoded Wallet: %+v\n", wallet)

	// Respond to the client
	request_util.ResponseToJson(&w, r, wallet)
}

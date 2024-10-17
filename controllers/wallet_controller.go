package controllers

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/services"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/encryption_util"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/error_util"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/request_util"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
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

func SyncWalletToS3(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Request method not allowed", http.StatusBadRequest)
		return
	}

	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "https://spendings.pages.dev")  // Set this to your allowed origin
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allowed methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allowed headers
	w.Header().Set("Access-Control-Allow-Credentials", "true")                    // Allow credentials

	var cookie *http.Cookie
	var err error
	cookie, err = r.Cookie("user-id")

	if err != nil {
		new_encrypted_cookie, err := encryption_util.GenerateNewCookie()
		error_util.Handle("Failed to generate new cookie", err)
		cookie = &http.Cookie{
			Name:     "user-id",
			Value:    new_encrypted_cookie,
			Path:     "/",
			HttpOnly: true,
			Domain:   "https://spendings.pages.dev", // Correct domain
			SameSite: http.SameSiteNoneMode,
			Secure:   true,                                     // Set SameSite correctly
			Expires:  time.Now().Add(time.Hour * 24 * 30 * 12), // 12 months
		}
	}
	http.SetCookie(w, cookie)

	user_id, err := encryption_util.DecryptData(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Read the request body as json
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	s3url, err := services.SyncWalletToS3(user_id, string(body))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	request_util.ResponseToJson(&w, r, s3url)
}

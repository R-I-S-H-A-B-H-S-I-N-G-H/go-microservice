package controllers

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/services"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/error_util"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/request_util"
	"net/http"
)

type MailPayloadBody struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func SendMailController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Request method not allowed", http.StatusBadRequest)
		return
	}

	// Read the request body as json
	var requestData MailPayloadBody
	err := request_util.DecodeJson(r.Body, &requestData)
	error_util.Handle("Failed to decode JSON", err)

	// Send mail
	mail_send_err := services.SendMailService(requestData.From, requestData.To, requestData.Subject, requestData.Body)

	if mail_send_err != nil {
		http.Error(w, mail_send_err.Error(), http.StatusInternalServerError)
		return
	}

	request_util.ResponseToJson(&w, r, "Mail sent successfully")
}

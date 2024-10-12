package main

import (
	"fmt"
	"log"
	"net/http"
	"rishabhsingh2305/spendings-app/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)    // Logs every request
	router.Use(middleware.Recoverer) // Recovers from panic error
	router.Use(middleware.Heartbeat("/ping"))
	router.Route("/wallet", handlers.WalletHandler)
	router.Route("/s3", handlers.S3Handler)

	log.Println("Listening on port 3000")
	error := http.ListenAndServe(":3000", router)
	if error != nil {
		fmt.Println(error)
		log.Fatal(error)
	}
}

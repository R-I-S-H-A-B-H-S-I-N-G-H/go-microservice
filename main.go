package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rishabhsingh2305/spendings-app/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)    // Logs every request
	router.Use(middleware.Recoverer) // Recovers from panic error
	router.Use(middleware.Heartbeat("/ping"))
	router.Route("/wallet", handlers.WalletHandler)
	router.Route("/s3", handlers.S3Handler)

	PORT := ":" + os.Getenv("PORT")
	log.Println("Listening on port " + PORT)
	error := http.ListenAndServe(PORT, router)
	if error != nil {
		fmt.Println(error)
		log.Fatal(error)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/handlers"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/jobs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gopkg.in/robfig/cron.v2"
)

func main() {
	// Load .env file only in local development
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://jsserve.pages.dev", "http://localhost:*"}, // Use your allowed origins
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	router := chi.NewRouter()
	router.Use(c.Handler)
	router.Use(middleware.Logger)    // Logs every request
	router.Use(middleware.Recoverer) // Recovers from panic error
	router.Use(middleware.Heartbeat("/ping"))
	router.Route("/wallet", handlers.WalletHandler)
	router.Route("/s3", handlers.S3Handler)

	//jobs
	startJob()

	PORT := ":" + os.Getenv("PORT")
	log.Println("Listening on port " + PORT)
	error := http.ListenAndServe(PORT, router)
	if error != nil {
		fmt.Println(error)
		log.Fatal(error)
	}
}

func startJob() {
	c := cron.New()
	c.AddFunc("* */5 * * * *", jobs.PingMicroService)
	c.Start()
}

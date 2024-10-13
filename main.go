package main

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/handlers"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/jobs"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/error_util"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gopkg.in/robfig/cron.v2"
)

func main() {
	setupEnvVars()
	c := setupCors()

	router := chi.NewRouter()
	router.Use(c.Handler)
	router.Use(middleware.Logger) // Logs every request
	router.Use(middleware.Recoverer)

	setupRoutes(router)

	//jobs
	initJobs()

	PORT := ":" + os.Getenv("PORT")
	log.Println("Listening on port " + PORT)
	error := http.ListenAndServe(PORT, router)
	error_util.Handle("Error starting server", error)
}

func setupRoutes(router chi.Router) {
	// Recovers from panic error
	router.Use(middleware.Heartbeat("/ping"))
	router.Route("/wallet", handlers.WalletHandler)
	router.Route("/s3", handlers.S3Handler)
	router.Route("/mail", handlers.MailHandler)
	router.Route("/pixel", handlers.PixelHandler)
}

func setupEnvVars() {
	// Load .env file only in local development
	if os.Getenv("ENV") == "production" {
		return
	}
	err := godotenv.Load()
	error_util.Handle("Error loading .env file", err)
}

func setupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"https://jsserve.pages.dev", "http://localhost:*"}, // Use your allowed origins
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})
}

func initJobs() {
	c := cron.New()
	c.AddFunc("@every 10m", jobs.PingMicroService)
	c.Start()
}

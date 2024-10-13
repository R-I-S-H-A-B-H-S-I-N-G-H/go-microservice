package config

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/error_util"
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DatabaseClient *mongo.Client

func SetupDB() *mongo.Client {
	if DatabaseClient != nil {
		return DatabaseClient
	}

	db_connection_string := os.Getenv("DB_CONN")
	// Set client options
	clientOptions := options.Client().ApplyURI(db_connection_string) // Change the URI as needed

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	error_util.Handle("Failed to connect to MongoDB", err)

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	error_util.Handle("Failed to ping MongoDB", err)

	fmt.Println("Connected to MongoDB!")
	DatabaseClient = client
	return DatabaseClient
}

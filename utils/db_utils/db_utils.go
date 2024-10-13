package db_utils

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/config"
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/error_util"
	"context"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetMongoClient() *mongo.Client {
	mongoClient := config.DatabaseClient

	if mongoClient == nil {
		error_util.Handle("Failed to get mongo client", errors.New("failed to get mongo client"))
	}
	return mongoClient
}

func GetMongoDb() *mongo.Database {
	return GetMongoClient().Database(os.Getenv("DB_NAME"))
}

// CreateOne inserts a single document into the specified collection.
func CreateOne(collectionName string, document interface{}) error {

	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new collection
	collection := GetMongoDb().Collection(collectionName)

	// Insert the document
	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		return err // Return the error instead of terminating the program
	}

	log.Println("Document inserted successfully")
	return nil
}

// FindOne retrieves a single document from the specified collection based on the filter.
func FindOne(collectionName string, filter interface{}) (*mongo.SingleResult, error) {
	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := GetMongoDb().Collection(collectionName)

	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return result, result.Err() // Return the error
	}

	return result, nil
}

// UpdateOne updates a single document in the specified collection based on the filter.
func UpdateOne(collectionName string, filter interface{}, update interface{}) error {
	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := GetMongoDb().Collection(collectionName)

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err // Return the error
	}

	log.Println("Document updated successfully")
	return nil
}

// DeleteOne deletes a single document from the specified collection based on the filter.
func DeleteOne(collectionName string, filter interface{}) error {
	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := GetMongoDb().Collection(collectionName)

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err // Return the error
	}

	log.Println("Document deleted successfully")
	return nil
}

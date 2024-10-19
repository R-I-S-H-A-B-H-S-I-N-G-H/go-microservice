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
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
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

func CreateOne[T any](collectionName string, document *T) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := GetMongoDb().Collection(collectionName)

	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	log.Println("Document inserted successfully")
	return document, nil
}

// FindOne retrieves a single document from the specified collection based on the filter.
func FindOne[T any](collectionName string, filter interface{}) (*T, error) {
	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := GetMongoDb().Collection(collectionName)

	// Initialize a variable of type T to hold the result
	var result T

	// Find the document
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err // Return nil and the error
	}

	return &result, nil // Return the pointer to the result
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

func FindAllWithPagination(collectionName string, filter interface{}, page, limit int, sortField string, ascending bool) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := GetMongoDb().Collection(collectionName)

	// Create sort option
	sortOrder := 1
	if !ascending {
		sortOrder = -1
	}

	// Use bson.M for sort
	sort := bson.M{sortField: sortOrder}

	// Validate pagination values
	if page < 1 {
		page = 1 // Default to page 1
	}
	if limit < 1 {
		limit = 10 // Default limit
	}

	// Find with sort option and pagination
	cursor, err := collection.Find(ctx, filter, options.Find().SetSort(sort).SetLimit(int64(limit)).SetSkip(int64((page-1)*limit)))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

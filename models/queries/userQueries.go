package queries

import (
	"context"
	"log"
	"time"

	"github.com/dilippm92/taskmanager/config" // Import the config package
	"github.com/dilippm92/taskmanager/models" // Import the models package for the User type
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetUserCollection returns a handle to the users collection
func GetUserCollection() *mongo.Collection {
	if config.MongoClient == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	return config.MongoClient.Database("taskmanager").Collection("users")
}

// CreateUser inserts a new user document into the collection
func CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	collection := GetUserCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure that resources are cleaned up

	// Perform the InsertOne operation with the context
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Failed to insert user: %v", err)
		return nil, err
	}
	return result, nil
}

// FindUserByID retrieves a user document by ID
func FindUserByID(userID string) (models.User, error) {
	collection := GetUserCollection()
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		log.Printf("Failed to find user: %v", err)
		return models.User{}, err
	}
	return user, nil
}

// getuser by email
func FindUserByEmail(email string) (models.User, error) {
	collection := GetUserCollection()
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Printf("Failed to find user: %v", err)
		return models.User{}, err
	}
	return user, nil
}

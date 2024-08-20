package queries

import (
	"context"
	"errors"

	"log"
	"time"

	"github.com/dilippm92/taskmanager/config"
	"github.com/dilippm92/taskmanager/models"
	

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// get subtask collection returns a handle to the subtask collection
func GetSubTaskCollection() *mongo.Collection {
	if config.MongoClient == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	return config.MongoClient.Database("taskmanager").Collection("tasks")
}

func CreateSubTask(task models.SubTask) (*mongo.InsertOneResult, error) {
    // Get the sub-task collection
    collection := GetSubTaskCollection()
    if collection == nil {
        return nil, errors.New("failed to get sub-task collection")
    }

    // Create a context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Insert the sub-task into the collection
    result, err := collection.InsertOne(ctx, task)
    if err != nil {
        log.Printf("Failed to insert sub-task: %v", err)
        return nil, err
    }

    // Convert SprintID from string to ObjectID
    sprintID, err := primitive.ObjectIDFromHex(task.SprintID)
    if err != nil {
        log.Printf("Invalid SprintID: %v", err)
        return nil, err
    }

    // Update the sprint with the new sub-task ID
    err = UpdateSprintWithSubTask(sprintID, result.InsertedID)
    if err != nil {
        return nil, err
    }

    return result, nil
}

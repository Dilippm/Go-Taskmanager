package queries

import (
	"context"

	"log"

	"time"

	"github.com/dilippm92/taskmanager/config"
	"github.com/dilippm92/taskmanager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetSprintCollection returns a handle to the sprints collection
func GetSprintCollection() *mongo.Collection {
	if config.MongoClient == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	return config.MongoClient.Database("taskmanager").Collection("sprints")
}

// create a new sprint

func CreateSprint(sprint models.Sprint )(*mongo.InsertOneResult,error){
	collection:= GetSprintCollection()
	ctx,cancel:= context.WithTimeout(context.Background(),10* time.Second)
	defer cancel()
	// insert one operation with context
	result,err:= collection.InsertOne(ctx,sprint)
	if err!= nil{
		log.Printf("failed to insert sprint:%v",err)
		return nil,err
	}
return result,nil
}

// get all sprints for a user
func GetAllSprints(userIdStr string) ([]models.Sprint, error) {
	collection := GetSprintCollection()
	var sprints []models.Sprint
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find all sprints matching the userId
	cursor, err := collection.Find(ctx, bson.M{"userId": userIdStr})
	if err != nil {
		log.Printf("Failed to find sprints: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate through the cursor and decode each document into the sprints slice
	for cursor.Next(ctx) {
		var sprint models.Sprint
		if err := cursor.Decode(&sprint); err != nil {
			log.Printf("Failed to decode sprint: %v", err)
			return nil, err
		}
		sprints = append(sprints, sprint)
	}

	// Check for cursor error after iteration
	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return sprints, nil
}

// get single sprint by sprint id

func GetSingleSprintDetails(sprintId string)(models.Sprint,error){
	collection:= GetSprintCollection()
	id, err := primitive.ObjectIDFromHex(sprintId)
	if err != nil {
		return models.Sprint{}, err
	}
	var sprint models.Sprint
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&sprint)
	if err != nil {
		log.Printf("Failed to find user: %v", err)
		return models.Sprint{}, err
	}
	return sprint, nil
}
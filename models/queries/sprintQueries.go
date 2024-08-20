package queries

import (
	"context"

	"log"

	"time"
	"fmt"
"errors"
	"github.com/dilippm92/taskmanager/config"
	"github.com/dilippm92/taskmanager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	 "go.mongodb.org/mongo-driver/mongo/options"
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
	// Ensure the sub_tasks field is an array, defaulting to an empty array if not provided
	if sprint.SubTasks == nil {
		sprint.SubTasks = []primitive.ObjectID{}
	}

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

	// If the sprint contains SubTask IDs, populate them
	if len(sprint.SubTasks) > 0 {
		subTaskCollection := GetSubTaskCollection() // Assume you have this function
		cursor, err := subTaskCollection.Find(ctx, bson.M{"_id": bson.M{"$in": sprint.SubTasks}})
		if err != nil {
			log.Printf("Failed to find subtasks: %v", err)
			return models.Sprint{}, err
		}
		defer cursor.Close(ctx)

		var subTasks []models.SubTask
		for cursor.Next(ctx) {
			var subTask models.SubTask
			if err = cursor.Decode(&subTask); err != nil {
				log.Printf("Failed to decode subtask: %v", err)
				continue
			}
			subTasks = append(subTasks, subTask)
		}

		if err := cursor.Err(); err != nil {
			log.Printf("Cursor error: %v", err)
			return models.Sprint{}, err
		}

		// Assign the populated subTasks to the sprint
		sprint.PopulatedSubTasks = subTasks
	}

	return sprint, nil
}

// update a sprint by sprint id

func UpdateSprint(id primitive.ObjectID, sprint models.Sprint) (*mongo.UpdateResult, error) {
	collection:= GetSprintCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
	 // Create the update document
	 
	 update := bson.M{
        "$set": bson.M{
            "sprint_name": sprint.SprintName,
            "start_date":  sprint.StartDate,
            "end_date":    sprint.EndDate,
            "sub_tasks":   sprint.SubTasks,
            "priority":    sprint.Priority,
            "userId":      sprint.UserId,
			"status":sprint.Status,
        },
    }


    // Specify the filter and update options
    filter := bson.M{"_id": id}
    opts := options.Update().SetUpsert(false)

    // Perform the update
    result, err := collection.UpdateOne(ctx, filter, update, opts)
    if err != nil {
        return nil, err
    }

    return result, nil
	

}
// Delete a sprint by sprint id
func DeleteSprint(id primitive.ObjectID)(*mongo.DeleteResult, error){
	collection := GetSprintCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Specify the filter to match the document to delete
	filter := bson.M{"_id": id}

	// Perform the delete operation
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
// update sprint data subtask field with new task id
func UpdateSprintWithSubTask(sprintID primitive.ObjectID, subTaskID interface{}) error {
    // Get the sprint collection
    sprintCollection := GetSprintCollection()
    if sprintCollection == nil {
        return errors.New("failed to get sprint collection")
    }

    // Create a context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Prepare the update document to initialize sub_tasks as an array if it is null, then push the new sub-task ID
    update := bson.M{
       
        "$push": bson.M{
            "sub_tasks": subTaskID,
        },
    }

    // Prepare the filter to find the sprint by the provided sprint_id
    filter := bson.M{"_id": sprintID}

    // Perform the update
    sprResult, err := sprintCollection.UpdateOne(ctx, filter, update)
    if err != nil {
        log.Printf("Failed to update sprint: %v", err)
        return err
    }
    if sprResult.MatchedCount == 0 {
        log.Printf("No sprint found with the given SprintID")
        return fmt.Errorf("no sprint found with ID: %v", sprintID.Hex())
    }

    return nil
}

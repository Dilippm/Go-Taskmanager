package queries

import (
	"context"
	"errors"

	"log"
	"time"

	"github.com/dilippm92/taskmanager/config"
	"github.com/dilippm92/taskmanager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
     "go.mongodb.org/mongo-driver/mongo/options"
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

func GetTaskDetails(taskId string)(models.SubTask,error){
    collection:= GetSubTaskCollection()
    id,err:=primitive.ObjectIDFromHex(taskId)
if err!= nil{
    return models.SubTask{},err
}
var task models.SubTask
ctx,cancel:= context.WithTimeout(context.Background(),10*time.Second)
defer cancel()
err = collection.FindOne(ctx,bson.M{"_id":id}).Decode(&task)
if err != nil {
    log.Printf("Failed to find user: %v", err)
    return models.SubTask{}, err
}
return task, nil
}

func UpdateTask(id primitive.ObjectID, task models.SubTask)(*mongo.UpdateResult, error) {
    collection:= GetSubTaskCollection()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
	 // Create the update document
     update := bson.M{
        "$set": bson.M{
            "task_name": task.TaskName,
            "start_date":  task.StartDate,
            "end_date":    task.EndDate,
           "description": task.Description,
            "priority":    task.Priority,
            "status": task.Status,
           
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

func DeleteTask(id primitive.ObjectID)(*mongo.DeleteResult, error){
    collection:= GetSubTaskCollection()
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
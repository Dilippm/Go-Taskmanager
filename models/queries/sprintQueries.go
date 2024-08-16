package queries

import (
	"context"
	"log"
	
	"time"

	"github.com/dilippm92/taskmanager/config"
	"github.com/dilippm92/taskmanager/models"
	"go.mongodb.org/mongo-driver/mongo"
	
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
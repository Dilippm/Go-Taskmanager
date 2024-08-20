package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Priority represents the priority level of a task.
type Priority string

const (
	High   Priority = "High"
	Medium Priority = "Medium"
	Low    Priority = "Low"
)
// status represents current status of a task and sprint
type Status string
const (
	Todo Status ="Todo"
	Progress Status ="Progress"
	Completed Status="Completed"
)

// Sprint represents a sprint in the project management system.
type Sprint struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	SprintName  string             `bson:"sprint_name"`
	StartDate   time.Time          `bson:"start_date"`
	EndDate     time.Time          `bson:"end_date"`
	SubTasks    []primitive.ObjectID `bson:"sub_tasks"` // Array of SubTask IDs
	Priority    Priority           `bson:"priority"`
	Status    Status           `bson:"status"`
	UserId   string 			   `bson:"userId"`
	PopulatedSubTasks []SubTask          `json:"PopulatedSubTasks"`
}

// SubTask represents a subtask in the project management system.
type SubTask struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	TaskName    string             `bson:"task_name"`
	Description string             `bson:"description"`
	StartDate   time.Time          `bson:"start_date"`
	EndDate     time.Time          `bson:"end_date"`
	Priority    Priority           `bson:"priority"`
	Status    Status           `bson:"status"`
	SprintID    string `bson:"sprint_id"` // Reference to the Sprint ID
}

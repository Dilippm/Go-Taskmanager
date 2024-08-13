package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the structure of a user document in MongoDB
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}

package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Instructor struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Paths    []LearningPath     `bson:"paths" json:"paths"`
	Age      int                `bson:"age" json:"age"`
}
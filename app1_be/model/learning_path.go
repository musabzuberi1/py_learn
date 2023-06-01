package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LearningPath struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Instructor  primitive.ObjectID `bson:"instructor" json:"instructor"`
}


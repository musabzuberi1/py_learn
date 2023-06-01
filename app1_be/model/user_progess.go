package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompletedExercise struct {
	ExerciseID primitive.ObjectID `bson:"exercise_id" json:"exercise_id"`
	UserCode   string             `bson:"user_code" json:"user_code"`
}

type UserProgress struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	UserID          primitive.ObjectID   `bson:"user_id" json:"user_id"`
	LearningPathID  primitive.ObjectID   `bson:"learning_path_id" json:"learning_path_id"`
	CompletedExercises []CompletedExercise `bson:"completed_exercises" json:"completed_exercises"`
}


package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercise struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	LearningPathID primitive.ObjectID `bson:"learning_path_id" json:"learning_path_id"`
	Title          string             `bson:"title" json:"title"`
	Description    string             `bson:"description" json:"description"`
	VideoURL       string             `bson:"video_url" json:"video_url"`
	SampleCode     string             `bson:"sample_code" json:"sample_code"`
}
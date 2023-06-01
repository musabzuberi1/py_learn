package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const Points_Per_Exercise = 5
const level_thresh = 25

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Age		 int				`bson:"name" json:"name"`
	Level    int                `bson:"level" json:"level"`
	Points   int                `bson:"points" json:"points"`
}

func (u *User) AddPoints() {
	u.Points += Points_Per_Exercise
	if u.Points >= level_thresh {
		u.Level++
		u.Points -= level_thresh
	}
}

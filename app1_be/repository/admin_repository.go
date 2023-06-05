package repository

import (
	"context"
	"myapp/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepository struct {
	DB *mongo.Database
}

func (ar *AdminRepository) GetUser(ctx context.Context, id string) (model.User, error) {
	var user model.User
	err := ar.DB.Collection("users").FindOne(ctx, bson.M{"id": id}).Decode(&user)
	return user, err
}

func (ar *AdminRepository) GetInstructor(ctx context.Context, id string) (model.Instructor, error) {
	var instructor model.Instructor
	err := ar.DB.Collection("instructors").FindOne(ctx, bson.M{"id": id}).Decode(&instructor)
	return instructor, err
}

func (ar *AdminRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	cursor, err := ar.DB.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}
	err = cursor.All(ctx, &users)
	return users, err
}

func (ar *AdminRepository) GetAllInstructors(ctx context.Context) ([]model.Instructor, error) {
	var instructors []model.Instructor
	cursor, err := ar.DB.Collection("instructors").Find(ctx, bson.M{})
	if err != nil {
		return instructors, err
	}
	err = cursor.All(ctx, &instructors)
	return instructors, err
}

func (ar *AdminRepository) UpdateUser(ctx context.Context, id string, u model.User) error {
	_, err := ar.DB.Collection("users").UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": u})
	return err
}

func (ar *AdminRepository) UpdateInstructor(ctx context.Context, id string, i model.Instructor) error {
	_, err := ar.DB.Collection("instructors").UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": i})
	return err
}

func (ar *AdminRepository) DeleteUser(ctx context.Context, id string) error {
	_, err := ar.DB.Collection("users").DeleteOne(ctx, bson.M{"id": id})
	return err
}

func (ar *AdminRepository) DeleteInstructor(ctx context.Context, id string) error {
	_, err := ar.DB.Collection("instructors").DeleteOne(ctx, bson.M{"id": id})
	return err
}

// myapp/repository/user_repository.go
package repository

import (
	"errors"

	"myapp/model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type UserRepository struct {
	db         *mgo.Database
	collection *mgo.Collection
}

func NewUserRepository(db *mgo.Database) *UserRepository {
	collection := db.C("users")
	return &UserRepository{db: db, collection: collection}
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	err := ur.collection.Insert(user)
	return err
}

func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := ur.collection.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	err := ur.collection.FindId(bson.ObjectIdHex(id)).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) UpdateUser(id string, userUpdate *model.UserUpdate) (*model.User, error) {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	update := bson.M{"$set": bson.M{
		"email":    userUpdate.Email,
		"password": userUpdate.Password,
	}}

	err := ur.collection.Update(query, update)
	if err != nil {
		return nil, err
	}

	updatedUser, err := ur.FindByID(id)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (ur *UserRepository) Enroll(userID, learningPathID string) error {
	query := bson.M{"_id": bson.ObjectIdHex(userID)}
	update := bson.M{"$addToSet": bson.M{"enrolledLearningPaths": learningPathID}}

	err := ur.collection.Update(query, update)
	return err
}

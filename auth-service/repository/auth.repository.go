package repository

import (
	"auth/errors"
	"auth/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	cli *mongo.Client
}

func NewAuthRepository(cli *mongo.Client) (*AuthRepository, error) {
	return &AuthRepository{
		cli: cli,
	}, nil
}

func (a AuthRepository) SaveUser(user models.Citizen) (*models.Citizen, *errors.ErrorStruct) {
	userColection := a.cli.Database("auth").Collection("user")
	insertedUser, err := userColection.InsertOne(context.TODO(), user)
	if err != nil {
		err, status := errors.HandleInsertError(err, user)
		if status == -1 {
			status = 500
		}
		return nil, errors.NewError(err.Error(), status)
	}
	user.ID = insertedUser.InsertedID.(primitive.ObjectID)
	return &user, nil
}

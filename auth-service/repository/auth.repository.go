package repository

import (
	"auth/errors"
	"auth/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

func (a AuthRepository) FindUserByEmail(email string) (*models.Citizen, *errors.ErrorStruct) {
	userCollection := a.cli.Database("auth").Collection("user")
	var user models.Citizen
	err := userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Println("auth-db", err.Error())
		return nil, errors.NewError(
			"Bad credentials",
			401)
	}
	return &user, nil
}

package repositories

import (
	"context"
	"dorm-service/errors"
	"dorm-service/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApplicationsRepository struct {
	cli *mongo.Client
}

func NewApplicationsRepository(cli *mongo.Client) (*ApplicationsRepository, error) {
	return &ApplicationsRepository{
		cli: cli,
	}, nil
}

func (ar ApplicationsRepository) SaveNewDorm(application models.ApplicationForDorm) (*models.ApplicationForDorm, *errors.ErrorStruct) {
	applicationsCollection := ar.cli.Database("dorm").Collection("applications")
	insertedDorm, err := applicationsCollection.InsertOne(context.TODO(), application)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	application.ID = insertedDorm.InsertedID.(primitive.ObjectID)
	return &application, nil
}

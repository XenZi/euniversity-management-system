package repository

import (
	"context"
	"fakultet-service/errors"
	"fakultet-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UniversityRepository struct {
	cli *mongo.Client
}

func NewUniversityRepository(cli *mongo.Client) (*UniversityRepository, error) {
	return &UniversityRepository{
		cli: cli,
	}, nil
}

func (u UniversityRepository) SaveUniversity(university models.University) (*models.University, *errors.ErrorStruct) {
	universityCollection := u.cli.Database("university").Collection("university")
	insertResult, err := universityCollection.InsertOne(context.TODO(), university)
	if err != nil {
		err, status := errors.HandleUniversityInsertError(err, university)
		if status == -1 {
			status = 500
		}
		return nil, errors.NewError(err.Error(), status)
	}
	university.Id = insertResult.InsertedID.(primitive.ObjectID)
	return &university, nil

}

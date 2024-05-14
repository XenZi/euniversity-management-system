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
		return nil, errors.NewError(err.Error(), 500)
	}
	university.ID = insertResult.InsertedID.(primitive.ObjectID)
	return &university, nil

}

func (u UniversityRepository) SaveStudent(student models.Student) (*models.Student, *errors.ErrorStruct) {
	studentCollection := u.cli.Database("university").Collection("student")
	insertResult, err := studentCollection.InsertOne(context.TODO(), student)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	student.ID = insertResult.InsertedID.(primitive.ObjectID)
	return &student, nil
}

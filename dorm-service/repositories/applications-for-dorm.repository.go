package repositories

import (
	"context"
	"dorm-service/errors"
	"dorm-service/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

func (ar ApplicationsRepository) FindApplicationsByStudentID(pid string) ([]models.ApplicationForDorm, *errors.ErrorStruct) {
	applicationsCollection := ar.cli.Database("dorm").Collection("applications")
	filter := bson.M{
		"student.pid": pid,
	}
	cursor, err := applicationsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	var results []models.ApplicationForDorm
	for cursor.Next(context.TODO()) {
		var application models.ApplicationForDorm
		err := cursor.Decode(&application)
		if err != nil {
			log.Fatal(err)
			return nil, errors.NewError(err.Error(), 500)
		}
		results = append(results, application)
	}
	return results, nil
}

func (ar ApplicationsRepository) UpdateApplicationStatus(applicationID string, applicationStatus models.ApplicationStatus) (*models.ApplicationForDorm, *errors.ErrorStruct) {
	applicationsCollection := ar.cli.Database("dorm").Collection("applications")
	objectID, _ := primitive.ObjectIDFromHex(applicationID)
	filter := bson.D{{Key: "_id", Value: objectID}}
	update := bson.D{{
		Key: "$set", Value: bson.D{
			{
				Key: "applicationStatus", Value: applicationStatus,
			},
		},
	}}
	updateResult, err := applicationsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	if updateResult.ModifiedCount == 0 {
		return nil, errors.NewError("Dorm not found", 404)
	}
	dorm, errFromDormFinding := ar.FindApplicationById(applicationID)
	if errFromDormFinding != nil {
		return nil, errors.NewError("Dorm not found", 404)
	}
	return dorm, nil
}

func (ar ApplicationsRepository) FindApplicationById(id string) (*models.ApplicationForDorm, *errors.ErrorStruct) {
	applicationsCollection := ar.cli.Database("dorm").Collection("applications")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	var application models.ApplicationForDorm
	err = applicationsCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(application)
	if err != nil {
		return nil, errors.NewError(
			"Not found with following ID",
			401)
	}
	return &application, nil
}

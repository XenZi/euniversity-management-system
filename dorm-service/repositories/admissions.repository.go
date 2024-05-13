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

type AdmissionsRepository struct {
	cli *mongo.Client
}

func NewAdmissionsRepository(cli *mongo.Client) (*AdmissionsRepository, error) {
	return &AdmissionsRepository{
		cli: cli,
	}, nil
}

func (ar AdmissionsRepository) SaveNewAdmission(admissions models.DormitoryAdmissions) (*models.DormitoryAdmissions, *errors.ErrorStruct) {
	admissionsCollection := ar.cli.Database("dorm").Collection("admissions")
	insertedDorm, err := admissionsCollection.InsertOne(context.TODO(), admissions)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	log.Println(admissions)
	admissions.ID = insertedDorm.InsertedID.(primitive.ObjectID)
	return &admissions, nil
}

func (ar AdmissionsRepository) FindAdmissionById(id string) (*models.DormitoryAdmissions, *errors.ErrorStruct) {
	admissionsCollection := ar.cli.Database("dorm").Collection("admissions")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	var dorm models.DormitoryAdmissions
	err = admissionsCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&dorm)
	if err != nil {
		return nil, errors.NewError(
			"Not found with following ID",
			401)
	}
	return &dorm, nil
}

func (ar AdmissionsRepository) DeleteAdmissionById(id string) (*models.DormitoryAdmissions, *errors.ErrorStruct) {
	admission, err := ar.FindAdmissionById(id)
	if err != nil {
		return nil, err
	}
	addmissionCollection := ar.cli.Database("dorm").Collection("admissions")
	objectID, errFromCast := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewError(errFromCast.Error(), 500)
	}
	filter := bson.M{"_id": objectID}
	_, errFromDelete := addmissionCollection.DeleteOne(context.TODO(), filter)
	if errFromDelete != nil {
		return nil, errors.NewError(errFromDelete.Error(), 500)
	}
	return admission, nil
}

func (ar AdmissionsRepository) FindAdmissionByDormID(id string) (*models.DormitoryAdmissions, *errors.ErrorStruct) {
	admissionsCollection := ar.cli.Database("dorm").Collection("admissions")
	var dorm models.DormitoryAdmissions
	err := admissionsCollection.FindOne(context.TODO(), bson.M{"dormID": id}).Decode(&dorm)
	if err != nil {
		return nil, errors.NewError(
			"Not found with following ID",
			401)
	}
	return &dorm, nil
}

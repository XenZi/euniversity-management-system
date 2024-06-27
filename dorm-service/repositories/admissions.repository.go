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

func (ar AdmissionsRepository) FindAdmissions() ([]*models.DormitoryAdmissions, *errors.ErrorStruct) {
	admissionsCollection := ar.cli.Database("dorm").Collection("admissions")
	var admissions []*models.DormitoryAdmissions
	cursor, err := admissionsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var admission models.DormitoryAdmissions
		if err := cursor.Decode(&admission); err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		admissions = append(admissions, &admission)
	}

	if err != nil {
		return nil, errors.NewError(
			"Not found with following ID",
			401)
	}
	return admissions, nil
}

func (ar AdmissionsRepository) UpdateAdmission(admissions models.DormitoryAdmissions) (*models.DormitoryAdmissions, *errors.ErrorStruct) {
	admissionsCollection := ar.cli.Database("dorm").Collection("admissions")
	filter := bson.D{{Key: "_id", Value: admissions.ID}}
	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{Key: "start", Value: admissions.Start},
				{Key: "end", Value: admissions.End},
			},
		},
	}
	_, err := admissionsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	admission, errFromDormission := ar.FindAdmissionById(admissions.ID.Hex())
	if errFromDormission != nil {
		return nil, errFromDormission
	}
	return admission, nil
}

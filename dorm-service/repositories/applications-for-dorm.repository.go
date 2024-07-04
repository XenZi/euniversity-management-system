package repositories

import (
	"context"
	"dorm-service/errors"
	"dorm-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (ar ApplicationsRepository) FindApplicationByID(applicationID string) (*models.ApplicationForDorm, *errors.ErrorStruct) {
	applicationsCollection := ar.cli.Database("dorm").Collection("applications")
	objectID, err := primitive.ObjectIDFromHex(applicationID)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	var application models.ApplicationForDorm
	err = applicationsCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&application)
	if err != nil {
		return nil, errors.NewError(
			"Not found with following ID",
			404)
	}
	return &application, nil
}

func (ar ApplicationsRepository) FindAllApplicationsForOneAdmission(admissionID string) ([]*models.ApplicationForDorm, *errors.ErrorStruct) {
	applicationsCollection := ar.cli.Database("dorm").Collection("applications")
	var applications []*models.ApplicationForDorm
	filter := bson.M{
		"dormitoryAdmissionsID": admissionID,
	}
	cursor, err := applicationsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var application models.ApplicationForDorm
		if err := cursor.Decode(&application); err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		applications = append(applications, &application)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	return applications, nil
}

func (ar ApplicationsRepository) FindAllUnderReviewApplicationsForSpecifiedAdmission(admissionID string) ([]*models.ApplicationForDorm, *errors.ErrorStruct) {
	applicationsCollection := ar.cli.Database("dorm").Collection("applications")
	var applications []*models.ApplicationForDorm
	filter := bson.M{
		"dormitoryAdmissionsID": admissionID,
		"applicationStatus":     0,
	}
	cursor, err := applicationsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var application models.ApplicationForDorm
		if err := cursor.Decode(&application); err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		applications = append(applications, &application)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	return applications, nil
}

func (ar ApplicationsRepository) FindAllAcceptedApplicationsForSpecifiedAdmission(admissionID string) ([]*models.ApplicationForDorm, *errors.ErrorStruct) {
	applicationsCollection := ar.cli.Database("dorm").Collection("applications")
	var applications []*models.ApplicationForDorm

	filter := bson.M{
		"dormitoryAdmissionsID": admissionID,
		"applicationStatus":     1,
	}

	// Define the sorting criteria
	sort := bson.D{
		{Key: "student.studentUniversityData.budgetStatus", Value: -1},
		{Key: "student.studentUniversityData.espb", Value: -1},
		{Key: "student.studentUniversityData.semester", Value: -1},
	}

	cursor, err := applicationsCollection.Find(context.TODO(), filter, options.Find().SetSort(sort))
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var application models.ApplicationForDorm
		if err := cursor.Decode(&application); err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		applications = append(applications, &application)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	return applications, nil
}

func (ar ApplicationsRepository) UpdateApplication(app models.ApplicationForDorm) (*models.ApplicationForDorm, *errors.ErrorStruct) {
	applicationsCollection := ar.cli.Database("dorm").Collection("applications")
	filter := bson.D{{
		Key: "_id", Value: app.ID,
	}}
	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{
					Key:   "applicationType",
					Value: app.ApplicationType,
				},
				{
					Key:   "verifiedStudent",
					Value: app.VerifiedStudent,
				},
				{
					Key:   "healthInsurance",
					Value: app.HealthInsurance,
				},
				{
					Key:   "applicationStatus",
					Value: app.ApplicationStatus,
				},
			},
		},
	}
	_, err := applicationsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	return &app, nil
}

func (ar ApplicationsRepository) DeleteApplicationByID(id string) (*models.ApplicationForDorm, *errors.ErrorStruct) {
	applications, err := ar.FindApplicationByID(id)
	if err != nil {
		return nil, err
	}
	applicationsCollection := ar.cli.Database("dorm").Collection("applications")
	objectID, errFromCast := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewError(errFromCast.Error(), 500)
	}
	filter := bson.D{{
		Key: "_id", Value: objectID,
	}}
	_, errFromDelete := applicationsCollection.DeleteOne(context.TODO(), filter)
	if errFromDelete != nil {
		return nil, errors.NewError(errFromDelete.Error(), 500)
	}
	return applications, nil
}

func (ar ApplicationsRepository) FindAllApplicationsByUserPIN(personalIdentificationNumber string) ([]*models.ApplicationForDorm, *errors.ErrorStruct) {
	applicationsCollection := ar.cli.Database("dorm").Collection("applications")
	var applications []*models.ApplicationForDorm
	filter := bson.M{
		"student.personalIdentificationNumber": personalIdentificationNumber,
	}
	cursor, err := applicationsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var application models.ApplicationForDorm
		if err := cursor.Decode(&application); err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		applications = append(applications, &application)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	return applications, nil
}

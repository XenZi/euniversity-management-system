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

type DormRepository struct {
	cli *mongo.Client
}

func NewDormRepository(cli *mongo.Client) (*DormRepository, error) {
	return &DormRepository{
		cli: cli,
	}, nil
}

func (dr DormRepository) SaveNewDorm(dorm models.Dorm) (*models.Dorm, *errors.ErrorStruct) {
	dormCollection := dr.cli.Database("dorm").Collection("dorm")
	insertedDorm, err := dormCollection.InsertOne(context.TODO(), dorm)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	log.Println(insertedDorm)
	dorm.ID = insertedDorm.InsertedID.(primitive.ObjectID)
	return &dorm, nil
}

func (dr DormRepository) FindDormById(id string) (*models.Dorm, *errors.ErrorStruct) {
	dormCollection := dr.cli.Database("dorm").Collection("dorm")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	var dorm models.Dorm
	err = dormCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&dorm)
	if err != nil {
		return nil, errors.NewError(
			"Not found with following ID",
			401)
	}
	return &dorm, nil
}

func (dr DormRepository) DeleteDormById(id string) (*models.Dorm, *errors.ErrorStruct) {
	dorm, err := dr.FindDormById(id)
	if err != nil {
		return nil, err
	}
	dormCollection := dr.cli.Database("dorm").Collection("dorm")
	objectID, errFromCast := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewError(errFromCast.Error(), 500)
	}
	filter := bson.M{"_id": objectID}
	_, errFromDelete := dormCollection.DeleteOne(context.TODO(), filter)
	if errFromDelete != nil {
		return nil, errors.NewError(errFromDelete.Error(), 500)
	}
	return dorm, nil
}

func (dr DormRepository) UpdateDorm(id, name, location string) (*models.Dorm, *errors.ErrorStruct) {
	dormCollection := dr.cli.Database("dorm").Collection("dorm")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: name},
			{Key: "location", Value: location},
		}},
	}
	updateResult, err := dormCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	if updateResult.ModifiedCount == 0 {
		return nil, errors.NewError("Dorm not found", 404)
	}

	dorm, errFromDormFinding := dr.FindDormById(id)

	if err != nil {
		return nil, errFromDormFinding
	}
	return dorm, nil
}

func (dr DormRepository) FindAllDorms() ([]*models.Dorm, *errors.ErrorStruct) {
	dormCollection := dr.cli.Database("dorm").Collection("dorm")
	var dorms []*models.Dorm
	cursor, err := dormCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("OVDE ERR ", err)
		return nil, errors.NewError(err.Error(), 500)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var dorm models.Dorm
		if err := cursor.Decode(&dorm); err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		log.Println(dorm)
		dorms = append(dorms, &dorm)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	return dorms, nil
}

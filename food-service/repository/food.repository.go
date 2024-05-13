package repository

import (
	"context"
	"food/errors"
	"food/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoodRepository struct {
	cli *mongo.Client
}

func NewFoodRepository(cli *mongo.Client) (*FoodRepository, error) {
	return &FoodRepository{
		cli: cli,
	}, nil
}

func (f FoodRepository) SaveFoodCard(card models.FoodCard) (*models.FoodCard, *errors.ErrorStruct) {
	cardCollection := f.cli.Database("food-service").Collection("cards")
	card.StudentID = "Mock-User"
	card.MassRoomID = "Mock-Mass"
	insertedCard, err := cardCollection.InsertOne(context.TODO(), card)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	card.ID = insertedCard.InsertedID.(primitive.ObjectID)
	return &card, nil

}

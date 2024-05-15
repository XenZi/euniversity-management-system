package repository

import (
	"context"
	"food/errors"
	"food/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (f FoodRepository) GetAllFoodCards() ([]models.FoodCard, *errors.ErrorStruct) {
	cardCollection := f.cli.Database("food-service").Collection("cards")

	// Define filter to get all documents
	filter := bson.M{}

	// Find documents based on the filter
	cursor, err := cardCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	defer cursor.Close(context.TODO())

	var foodCards []models.FoodCard

	// Iterate over the cursor and decode each document into a FoodCard struct
	for cursor.Next(context.TODO()) {
		var card models.FoodCard
		if err := cursor.Decode(&card); err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		foodCards = append(foodCards, card)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	return foodCards, nil
}

func (f FoodRepository) SavePayment(payment models.Payment) (*models.Payment, *errors.ErrorStruct) {
	paymentCollection := f.cli.Database("food-service").Collection("payment")
	payment.FoodCardID = "Mock-FoodCard"
	insertedPayment, err := paymentCollection.InsertOne(context.TODO(), payment)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	payment.ID = insertedPayment.InsertedID.(primitive.ObjectID)
	return &payment, nil

}

func (f FoodRepository) PayForMeal(studentPIN string, price int) (*models.FoodCard, *errors.ErrorStruct) {
	// Get the collection
	cardCollection := f.cli.Database("food-service").Collection("cards")

	// Define the filter to find the card by studentID
	filter := bson.M{"student_pin": studentPIN}

	// Fetch the existing card
	existingCard := models.FoodCard{}
	err := cardCollection.FindOne(context.TODO(), filter).Decode(&existingCard)
	if err != nil {
		return nil, errors.NewError("No food card found for the provided student PIN", 404)
	}

	// Check if the balance is sufficient
	if existingCard.Balance < price {
		return nil, errors.NewError("Insufficient balance in the food card", 400)
	}

	// Define the update operation to decrement the food_points by the price
	update := bson.M{
		"$inc": bson.M{"balance": -price},
	}

	// Perform the update operation
	_, err = cardCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	// Fetch the updated card
	updatedCard := models.FoodCard{}
	err = cardCollection.FindOne(context.TODO(), filter).Decode(&updatedCard)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	return &updatedCard, nil
}

package repository

import (
	"context"
	"food/errors"
	"food/models"
	"log"

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

// MESS ROOM CRUD
func (f FoodRepository) CreateMessRoom(messRoom models.MessRoom) (*models.MessRoom, *errors.ErrorStruct) {

	messCollection := f.cli.Database("food-service").Collection("messes")
	insertedMess, err := messCollection.InsertOne(context.TODO(), messRoom)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	messRoom.ID = insertedMess.InsertedID.(primitive.ObjectID)

	return &messRoom, nil

}

func (f FoodRepository) GetAllMessRooms() ([]models.MessRoom, *errors.ErrorStruct) {
	messCollection := f.cli.Database("food-service").Collection("messes")
	filter := bson.M{}

	cursor, err := messCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	defer cursor.Close(context.TODO())

	var messRooms []models.MessRoom

	for cursor.Next(context.TODO()) {
		var messRoom models.MessRoom
		if err := cursor.Decode(&messRoom); err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		messRooms = append(messRooms, messRoom)
	}
	if err := cursor.Err(); err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	return messRooms, nil

}

func (f FoodRepository) RemoveMessRoom(id string) (bool, *errors.ErrorStruct) {
	messCollection := f.cli.Database("food-service").Collection("messes")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, errors.NewError("Invalid id format", 400)
	}
	filter := bson.M{"_id": objID}

	_, err = messCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, errors.NewError(err.Error(), 500)
	}
	return true, nil
}

func (f FoodRepository) FindMessById(id string) (*models.MessRoom, *errors.ErrorStruct) {

	messCollection := f.cli.Database("food-service").Collection("messes")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewError("Invalid id format", 400)
	}

	var mess models.MessRoom
	err = messCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&mess)
	if err != nil {
		return nil, errors.NewError("For this id mess not found", 401)
	}
	return &mess, nil

}

func (f FoodRepository) UpdateMessRoom(updatedMess models.MessRoomUpdate) (*models.MessRoom, *errors.ErrorStruct) {
	messCollection := f.cli.Database("food-service").Collection("messes")
	log.Println("Vrijednosti koje su stigle do repoa", updatedMess)
	objID, err := primitive.ObjectIDFromHex(updatedMess.ID)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: updatedMess.Name},
			{Key: "location", Value: updatedMess.Location},
			{Key: "capacity", Value: updatedMess.Capacity},
		}},
	}
	_, err = messCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	mess, err1 := f.FindMessById(updatedMess.ID)
	if err1 != nil {
		return nil, err1
	}
	return mess, nil

}

// FOOD CARD CRUD

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

func (f FoodRepository) RemoveFoodCard(id string) (bool, *errors.ErrorStruct) {

	cardCollection := f.cli.Database("food-service").Collection("cards")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, errors.NewError("Invalid id format", 400)
	}
	filter := bson.M{"_id": objID}

	_, err = cardCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, errors.NewError(err.Error(), 500)
	}
	return true, nil

}

// PAYMENT CRUD

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

func (f FoodRepository) SaveUsageStatistics(stats models.UsageStatistics) (*models.UsageStatistics, *errors.ErrorStruct) {
	statsCollection := f.cli.Database("food-service").Collection("statistics")
	insertedStats, err := statsCollection.InsertOne(context.TODO(), stats)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	stats.ID = insertedStats.InsertedID.(primitive.ObjectID)
	return &stats, nil
}

// SUPPLIER CRUD

func (f FoodRepository) SaveSupplier(supplier models.Supplier) (*models.Supplier, *errors.ErrorStruct) {
	supplierCollection := f.cli.Database("food-service").Collection("suppliers")
	insertedSupplier, err := supplierCollection.InsertOne(context.TODO(), supplier)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	supplier.ID = insertedSupplier.InsertedID.(primitive.ObjectID)
	return &supplier, nil
}

func (f FoodRepository) GetAllSuppliers() ([]models.Supplier, *errors.ErrorStruct) {
	supplierCollection := f.cli.Database("food-service").Collection("suppliers")

	filter := bson.M{}

	cursor, err := supplierCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	defer cursor.Close(context.TODO())

	var suppliers []models.Supplier

	for cursor.Next(context.TODO()) {
		var supplier models.Supplier
		if err := cursor.Decode(&supplier); err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		suppliers = append(suppliers, supplier)
	}
	if err := cursor.Err(); err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	return suppliers, nil
}

func (f FoodRepository) DeleteSupplier(id string) (bool, *errors.ErrorStruct) {
	supplierCollection := f.cli.Database("food-service").Collection("suppliers")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, errors.NewError("Invalid id format", 400)
	}
	filter := bson.M{"_id": objID}

	_, err = supplierCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, errors.NewError(err.Error(), 500)
	}
	return true, nil
}

func (f FoodRepository) GetSupplierById(id string) (*models.Supplier, *errors.ErrorStruct) {
	supplierCollection := f.cli.Database("food-service").Collection("suppliers")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewError("Invalid id format", 400)
	}

	var supplier models.Supplier
	err = supplierCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&supplier)
	if err != nil {
		return nil, errors.NewError("Supplier with this id does not exist", 401)
	}
	return &supplier, nil
}

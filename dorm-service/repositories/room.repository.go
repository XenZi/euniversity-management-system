package repositories

import (
	"context"
	"dorm-service/errors"
	"dorm-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepository struct {
	cli *mongo.Client
}

func NewRoomRepository(cli *mongo.Client) *RoomRepository {
	return &RoomRepository{
		cli: cli,
	}
}

func (rr RoomRepository) SaveNewRoom(room models.Room) (*models.Room, *errors.ErrorStruct) {
	roomsCollection := rr.cli.Database("dorm").Collection("rooms")
	insertedRoom, err := roomsCollection.InsertOne(context.TODO(), room)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	room.ID = insertedRoom.InsertedID.(primitive.ObjectID)
	return &room, nil
}

func (rr RoomRepository) GetAllRoomsForDorm(dormID string) ([]*models.Room, *errors.ErrorStruct) {
	roomsCollection := rr.cli.Database("dorm").Collection("rooms")
	var rooms []*models.Room

	filter := bson.M{"dormID": dormID}

	cursor, err := roomsCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var room models.Room
		if err := cursor.Decode(&room); err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		rooms = append(rooms, &room)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	return rooms, nil
}

func (rr RoomRepository) FindOneRoomByID(roomID string) (*models.Room, *errors.ErrorStruct) {
	roomsCollection := rr.cli.Database("dorm").Collection("rooms")
	objectID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	var room models.Room
	err = roomsCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&room)
	if err != nil {
		return nil, errors.NewError(
			"Not found with following ID",
			404)
	}
	return &room, nil
}
func (rr RoomRepository) UpdateRoom(roomID string, squareFoot float32, numberOfBeds int16, toalet models.ToaletType) (*models.Room, *errors.ErrorStruct) {
	roomsCollection := rr.cli.Database("dorm").Collection("rooms")
	objectID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{Key: "squareFoot", Value: squareFoot},
				{Key: "numberOfBeds", Value: numberOfBeds},
				{Key: "toalet", Value: toalet},
			},
		},
	}
	_, err = roomsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	room, errFromRoomFinding := rr.FindOneRoomByID(roomID)
	if err != nil {
		return nil, errFromRoomFinding
	}
	return room, nil
}

func (rr RoomRepository) DeleteRoom(roomID string) (*models.Room, *errors.ErrorStruct) {
	room, errFromRoomFinding := rr.FindOneRoomByID(roomID)
	if errFromRoomFinding != nil {
		return nil, errFromRoomFinding
	}
	roomsCollection := rr.cli.Database("dorm").Collection("rooms")
	objectID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	filter := bson.M{"_id": objectID}
	_, errFromDelete := roomsCollection.DeleteOne(context.TODO(), filter)
	if errFromDelete != nil {
		return nil, errors.NewError(errFromDelete.Error(), 500)
	}
	return room, nil
}

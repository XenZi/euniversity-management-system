package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Dorm         string             `json:"dormID" bson:"dormID"`
	SquareFoot   float32            `json:"squareFoot" bson:"squareFoot"`
	Toalet       ToaletType         `json:"toalet" bson:"toalet"`
	NumberOfBeds int16              `json:"numberOfBeds" bson:"numberOfBeds"`
	Students     []Student          `json:"students" bson:"students"`
}

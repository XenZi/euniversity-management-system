package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DormitoryAdmissions struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DormID string             `json:"dormID" bson:"dormID"`
	Start  string             `json:"start" bson:"start"`
	End    string             `json:"end" bson:"end"`
	Active bool               `json:"active" bson:"active"`
}

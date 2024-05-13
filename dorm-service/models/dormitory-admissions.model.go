package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DormitoryAdmissions struct {
	ID           primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Dorm         primitive.ObjectID   `json:"dorm" bson:"dorm"`
	Start        string               `json:"start" bson:"start"`
	End          string               `json:"end" bson:"end"`
	Applications []ApplicationForDorm `json:"applications"`
}

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DormPriceByCategory struct {
	ApplicationType ApplicationType `json:"applicationType"`
	Price           float32         `json:"price"`
}

type Dorm struct {
	ID       primitive.ObjectID    `json:"id" bson:"_id,omitempty"`
	Name     string                `json:"name" bson:"name"`
	Location string                `json:"location" bson:"location"`
	Prices   []DormPriceByCategory `json:"prices" bson:"prices"`
}

type DormDTO struct {
	ID       string                `json:"id"`
	Name     string                `json:"name"`
	Location string                `json:"location"`
	Prices   []DormPriceByCategory `json:"prices"`
}

package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID                           primitive.ObjectID `bson:"_id,omitempty"`
	FullName                     string             `json:"full_name" bson:"full_name"`
	Gender                       string             `json:"gender" bson:"gender"`
	IdentityCardNumber           string             `json:"identity_card_number" bson:"identity_card_number"`
	PersonalIdentificationNumber string             `json:"personal_id_number" bson:"personal_id_number"`
	BirthDate                    string             `json:"birth_date" bson:"birth_date"`
}

type FoodCard struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StudentID  string             `json:"student_id" bson:"student_id"`
	Expires    string             `json:"expires" bson:"expires"`
	UsedPoint  []string           `json:"used_point" bson:"used_point"`
	MassRoomID string             `json:"mass_room_id" bson:"mass_room_id"`
	Balance    int                `json:"balance" bson:"balance"`
}

type Payment struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FoodCardID       string             `json:"food_card_id" bson:"food_card_id"`
	CreditCardNumber string             `json:"credit_card_number" bson:"credit_card_number"`
	Name             string             `json:"name" bson:"name"`
	CVV              string             `json:"cvv" bson:"cvv"`
	Amount           int                `json:"amount" bson:"amount"`
}
type MessRoom struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name"`
	Location      string             `json:"location" bson:"location"`
	Capacity      int                `json:"capacity" bson:"capacity"`
	Rating        float64            `json:"rating" bson:"rating"`
	SupplierID    string             `json:"supplier_id" bson:"supplier_id"`
	MessRoomUsers []string           `json:"mess_room_users" bson:"mess_room_users"`
}
type Supplier struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Location    string             `json:"location" bson:"location"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	MassRoomID  string             `json:"mass_room_id" bson:"mass_room_id"`
}

type UsageStatistics struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	MassRoomID string             `json:"mass_room_id" bson:"mass_room_id"`
	StudentID  string             `json:"student_id" bson:"student_id"`
	UsedPoints int64              `json:"used_points" bson:"used_points"`
}

type EStatus int

const (
	Scholarship EStatus = iota
	SelfFinancing
)

func (f EStatus) String() string {
	switch f {
	case Scholarship:
		return "SCHOLARSHIP"
	case SelfFinancing:
		return "SELF FINANCING"
	default:
		return fmt.Sprintf("Unknown Form value: %d", int(f))
	}
}

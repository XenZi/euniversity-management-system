package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Department struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Schedule Schedule           `json:"schedule" bson:"schedule"`
}

type DepartmentDTO struct {
	ID       string   `json:"id" `
	Name     string   `json:"name"`
	Schedule Schedule `json:"schedule"`
}

type Schedule struct {
	Date map[string][]Slot `json:"date" bson:"date"`
}

type Slot struct {
	Time      string `json:"time" bson:"time"`
	DoctorID  string `json:"doctorID" bson:"doctorID"`
	PatientID string `json:"patientID" bson:"patientID"`
}

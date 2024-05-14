package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ApplicationForDorm struct {
	ID                    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DormitoryAdmissionsID string             `json:"dormitoryAdmissionsID" bson:"dormitoryAdmissionsID"`
	ApplicationType       ApplicationType    `json:"applicationType" bson:"applicationType"`
	VerifiedStudent       bool               `json:"verifiedStudent" bson:"verifiedStudent"`
	HealthInsurance       bool               `json:"healthInsurance" bson:"healthInsurance"`
	ApplicationStatus     ApplicationStatus  `json:"applicationStatus" bson:"applicationStatus"`
	Student               Student            `json:"student" bson:"student"`
}

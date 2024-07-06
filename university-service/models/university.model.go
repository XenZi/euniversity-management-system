package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type University struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Address string             `json:"address" bson:"address"`
}

type Citizen struct {
	ID                           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FullName                     string             `json:"fullName" bson:"username"`
	Gender                       string             `json:"gender" bson:"gender"`
	IdentityCardNumber           string             `json:"identityCardNumber" bson:"identityCardNumber"`
	Citizenship                  string             `json:"citizenship" bson:"citizenship"`
	PersonalIdentificationNumber string             `json:"personalIdentificationNumber" bson:"personalIdentificationNumber"`
	Residence                    Residence          `json:"residence" bson:"residence"`
	BirthData                    BirthData          `json:"birthData" bson:"birthData"`
	Email                        string             `json:"email" bson:"email"`
	Password                     string             `json:"password" bson:"password"`
	Roles                        []string           `json:"roles" bson:"roles"`
}

type Residence struct {
	Address                 string `json:"address" bson:"address"`
	PlaceOfResidence        string `json:"placeOfResidence" bson:"placeOfResidence"`
	MunicipalityOfResidence string `json:"municipalityOfResidence" bson:"municipalityOfResidence"`
	CountryOfResidence      string `json:"countryOfResidence" bson:"countryOfResidence"`
}

type BirthData struct {
	DateOfBirth         string `json:"dateOfBirth" bson:"dateOfBirth"`
	MunicipalityOfBirth string `json:"municapilityOfBirth" bson:"municapilityOfBirth"`
	CountryOfBirth      string `json:"countryOfBirth" bson:"countryOfBirth"`
}

type CitizenDTO struct {
	ID                           string    `json:"id"`
	FullName                     string    `json:"fullName"`
	Gender                       string    `json:"gender"`
	IdentityCardNumber           string    `json:"identityCardNumber"`
	Citizenship                  string    `json:"citizenship"`
	PersonalIdentificationNumber string    `json:"personalIdentificationNumber"`
	Residence                    Residence `json:"residence"`
	BirthData                    BirthData `json:"birthData"`
	Email                        string    `json:"email"`
	Roles                        []string  `json:"roles"`
}

type Student struct {
	ID                           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FullName                     string             `json:"fullName" bson:"username"`
	Gender                       string             `json:"gender" bson:"gender"`
	IdentityCardNumber           string             `json:"identityCardNumber" bson:"identityCardNumber"`
	Citizenship                  string             `json:"citizenship" bson:"citizenship"`
	PersonalIdentificationNumber string             `json:"personalIdentificationNumber" bson:"personalIdentificationNumber"`
	Residence                    Residence          `json:"residence" bson:"residence"`
	BirthData                    BirthData          `json:"birthData" bson:"birthData"`
	Email                        string             `json:"email" bson:"email"`
	Password                     string             `json:"password" bson:"password"`
	Roles                        []string           `json:"roles" bson:"roles"`
	BudgetStatus                 bool               `json:"budgetStatus" bson:"budgetStatus"`
	Espb                         int64              `json:"espb" bson:"espb"`
	Status                       StudentStatus      `json:"status" bson:"status"`
	University                   University         `json:"university" bson:"university"`
	Semester                     int64              `json:"semester" bson:"semester"`
}

type Professor struct {
	ID                           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FullName                     string             `json:"fullName" bson:"username"`
	Gender                       string             `json:"gender" bson:"gender"`
	IdentityCardNumber           string             `json:"identityCardNumber" bson:"identityCardNumber"`
	Citizenship                  string             `json:"citizenship" bson:"citizenship"`
	PersonalIdentificationNumber string             `json:"personalIdentificationNumber" bson:"personalIdentificationNumber"`
	Residence                    Residence          `json:"residence" bson:"residence"`
	BirthData                    BirthData          `json:"birthData" bson:"birthData"`
	Email                        string             `json:"email" bson:"email"`
	Password                     string             `json:"password" bson:"password"`
	Roles                        []string           `json:"roles" bson:"roles"`
	Status                       ProfessorStatus    `json:"status" bson:"status"`
	University                   University         `json:"university" bson:"university"`
}

type Scholarship struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Student Student            `json:"student" bson:"student"`
}
type ApplyForScholarship struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Student Student            `json:"student" bson:"student"`
}

type StudyingConfirmation struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	University   University         `json:"university" bson:"university"`
	Student      Student            `json:"student" bson:"student"`
	BudgetStatus bool               `json:"budgetStatus" bson:"budgetStatus"`
	Semester     int64              `json:"semester" bson:"semester"`
}

type StateExamApplication struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	University University         `json:"university" bson:"university"`
	Student    Student            `json:"student" bson:"student"`
}

type EntranceExam struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Citizen     CitizenDTO         `json:"citizen" bson:"citizen"`
	DateAndTime string             `json:"dateAndTime" bson:"dateAndTime"`
	University  University         `json:"university" bson:"university"`
}
type ExtendStatusApplication struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Citizen CitizenDTO         `json:"citizen" bson:"citizen"`
}

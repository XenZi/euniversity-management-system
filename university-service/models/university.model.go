package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	Citizen
	BudgetStatus string        `json:"budgetStatus" bson:"budgetStatus"`
	Espb         int64         `json:"espb" bson:"espb"`
	Status       StudentStatus `json:"status" bson:"status"`
}

type Professor struct {
	Citizen
	Status ProfessorStatus `json:"status" bson:"status"`
}

type Scholarship struct {
	Student    Student    `json:"student" bson:"student"`
	University University `json:"university" bson:"university"`
}

type StudyingConfirmation struct {
	University University `json:"university" bson:"university"`
	Student    Student    `json:"student" bson:"student"`
}

type StateExamApplication struct {
	University University `json:"university" bson:"university"`
	Student    Student    `json:"student" bson:"student"`
}

type EntranceExam struct {
	DateAndTime string `json:"dateAndTime" bson:"dateAndTime"`
	Address     string `json:"address" bson:"address"`
}

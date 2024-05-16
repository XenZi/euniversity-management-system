package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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

type LoginCitizenDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AddingRoles struct {
	PersonalIdentificationNumber string   `json:"pin"`
	Roles                        []string `json:"roles"`
}
type PersonFromClaims struct {
	Email string
	Roles []string
	PID   string
}

/*
	"name":  email,
	"roles": roles,
	"pid":   pid,
*/

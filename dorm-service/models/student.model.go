package models

type Student struct {
	ID                           string    `json:"id" bson:"_id,omitempty"`
	FullName                     string    `json:"fullName" bson:"username"`
	Gender                       string    `json:"gender" bson:"gender"`
	IdentityCardNumber           string    `json:"identityCardNumber" bson:"identityCardNumber"`
	Citizenship                  string    `json:"citizenship" bson:"citizenship"`
	PersonalIdentificationNumber string    `json:"personalIdentificationNumber" bson:"personalIdentificationNumber"`
	Residence                    Residence `json:"residence" bson:"residence"`
	BirthData                    BirthData `json:"birthData" bson:"birthData"`
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

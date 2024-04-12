package models

type Citizen struct {
	FullName                     string    `json:"fullName"`
	Gender                       string    `json:"gender"`
	IdentityCardNumber           string    `json:"identityCardNumber"`
	Citizenship                  string    `json:"citizenship"`
	PersonalIdentificationNumber string    `json:"personalIdentificationNumber"`
	Residence                    Residence `json:"residence"`
	BirthData                    BirthData `json:"birthData"`
	Email                        string    `json:"email"`
	Password                     string    `json:"password"`
	Roles                        []string  `json:"roles"`
}

type Residence struct {
	Address                 string `json:"address"`
	PlaceOfResidence        string `json:"placeOfResidence"`
	MunicipalityOfResidence string `json:"municipalityOfResidence"`
	CountryOfResidence      string `json:"countryOfResidence"`
}

type BirthData struct {
	DateOfBirth         string `json:"dateOfBirth"`
	MunicipalityOfBirth string `json:"municapilityOfBirth"`
	CountryOfBirth      string `json:"countryOfBirth"`
}

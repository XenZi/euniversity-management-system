package client

import (
	"dorm-service/errors"
	"dorm-service/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomCastStudent struct {
	ID                           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FullName                     string             `json:"fullName" bson:"username"`
	Gender                       string             `json:"gender" bson:"gender"`
	IdentityCardNumber           string             `json:"identityCardNumber" bson:"identityCardNumber"`
	Citizenship                  string             `json:"citizenship" bson:"citizenship"`
	PersonalIdentificationNumber string             `json:"personalIdentificationNumber" bson:"personalIdentificationNumber"`
	Residence                    models.Residence   `json:"residence" bson:"residence"`
	BirthData                    models.BirthData   `json:"birthData" bson:"birthData"`
	Email                        string             `json:"email" bson:"email"`
	Password                     string             `json:"password" bson:"password"`
	Roles                        []string           `json:"roles" bson:"roles"`
	BudgetStatus                 bool               `json:"budgetStatus" bson:"budgetStatus"`
	Espb                         int64              `json:"espb" bson:"espb"`
	Status                       int64              `json:"status" bson:"status"`
	University                   University         `json:"university" bson:"university"`
	Semester                     int64              `json:"semester" bson:"semester"`
}

type University struct {
	ID      string `json:"id" bson:"_id,omitempty"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
}

type UniversityClient struct {
	address string
	client  *http.Client
}

func NewUniversityClient(address string, client *http.Client) *UniversityClient {
	return &UniversityClient{
		address: address,
		client:  client,
	}
}

func (uc UniversityClient) VerifyUserIntegrityWithUniversity(personalIdentificationNumber string) (*models.StudentUniversityData, *errors.ErrorStruct) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/student/%s", uc.address, personalIdentificationNumber), nil)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, errors.NewError(err.Error(), 500)
	}
	resp, err := uc.client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, errors.NewError(err.Error(), 500)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		baseErrorResp := models.BaseErrorHttpResponse{}
		err := json.NewDecoder(resp.Body).Decode(&baseErrorResp)
		if err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		log.Println(baseErrorResp)
		return nil, errors.NewError("Causing an error because there is no document", 500)
	}
	baseHttpResponse := models.BaseHttpResponse{}
	err = json.NewDecoder(resp.Body).Decode(&baseHttpResponse)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	dataMap, ok := baseHttpResponse.Data.(map[string]interface{})
	if !ok {
		return nil, errors.NewError(err.Error(), 500)
	}

	dataJSON, err := json.Marshal(dataMap)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}

	var student CustomCastStudent
	err = json.Unmarshal(dataJSON, &student)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	studentUniversityData := models.StudentUniversityData{
		UniversityName:    student.University.Name,
		UniversityAddress: student.University.Address,
		BudgetStatus:      student.BudgetStatus,
		Espb:              int16(student.Espb),
		Semester:          int16(student.Semester),
	}
	return &studentUniversityData, nil
}

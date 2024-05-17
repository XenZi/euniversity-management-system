package services

import (
	"dorm-service/client"
	"dorm-service/errors"
	"dorm-service/models"
	"dorm-service/repositories"
	"log"
)

type ApplicationsService struct {
	applicationsRepository *repositories.ApplicationsRepository
	healthCareClient       *client.HealthCareClient
}

func NewApplicationsService(applicationsRepository *repositories.ApplicationsRepository, client *client.HealthCareClient) (*ApplicationsService, error) {
	return &ApplicationsService{
		applicationsRepository: applicationsRepository,
		healthCareClient:       client,
	}, nil
}

func (as ApplicationsService) CreateNewApplication(application models.ApplicationForDorm) (*models.ApplicationForDorm, *errors.ErrorStruct) {
	valueFromHealthCare, err := as.healthCareClient.GetUserHealthStatusConfirmation(application.Student.PersonalIdentificationNumber)
	if err != nil {
		log.Println(err.GetErrorMessage())
	}
	log.Println(valueFromHealthCare)
	application.HealthInsurance = valueFromHealthCare
	createdApplication, err := as.applicationsRepository.SaveNewDorm(application)
	if err != nil {
		return nil, err
	}
	return createdApplication, nil
}

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
	application.ApplicationStatus = models.Review

	createdApplication, err := as.applicationsRepository.SaveNewDorm(application)
	if err != nil {
		return nil, err
	}
	return createdApplication, nil
}

func (as ApplicationsService) FindApplicationsByUserPIN(userPIN string) ([]*models.ApplicationForDorm, *errors.ErrorStruct) {
	foundApplications, err := as.applicationsRepository.FindAllApplicationsByUserPIN(userPIN)
	if err != nil {
		return nil, err
	}
	return foundApplications, nil
}

func (as ApplicationsService) FindApplicationByID(applicationID string) (*models.ApplicationForDorm, *errors.ErrorStruct) {
	foundApplication, err := as.applicationsRepository.FindApplicationByID(applicationID)
	if err != nil {
		return nil, err
	}
	return foundApplication, nil
}

func (as ApplicationsService) FindAllAplicationsForOneAdmission(admissionID string) ([]*models.ApplicationForDorm, *errors.ErrorStruct) {
	foundApplications, err := as.applicationsRepository.FindAllApplicationsForOneAdmission(admissionID)
	if err != nil {
		return nil, err
	}
	return foundApplications, nil
}

func (as ApplicationsService) FindAllUnderReviewApplicationsForSpecifiedAdmission(admissionID string) ([]*models.ApplicationForDorm, *errors.ErrorStruct) {
	foundApplications, err := as.applicationsRepository.FindAllUnderReviewApplicationsForSpecifiedAdmission(admissionID)
	if err != nil {
		return nil, err
	}
	return foundApplications, nil
}

func (as ApplicationsService) FindAllAcceptedApplicationsForSpecifiedAdmission(admissionID string) ([]*models.ApplicationForDorm, *errors.ErrorStruct) {
	foundApplications, err := as.applicationsRepository.FindAllAcceptedApplicationsForSpecifiedAdmission(admissionID)
	if err != nil {
		return nil, err
	}
	return foundApplications, nil
}

func (as ApplicationsService) UpdateApplication(app models.ApplicationForDorm) (*models.ApplicationForDorm, *errors.ErrorStruct) {
	if !app.HealthInsurance {
		isHealthStatusConfirmed, err := as.healthCareClient.GetUserHealthStatusConfirmation(app.Student.PersonalIdentificationNumber)
		if err != nil {
			log.Println(err)
		}
		app.HealthInsurance = isHealthStatusConfirmed
	}
	// todo app.verifiedStudent = provera.
	if app.HealthInsurance && app.VerifiedStudent {
		app.ApplicationStatus = models.Accepted
	} else {
		app.ApplicationStatus = models.Pending
	}
	updatedApp, err := as.applicationsRepository.UpdateApplication(app)
	if err != nil {
		return nil, err
	}
	return updatedApp, nil
}

func (as ApplicationsService) DeleteApplicationByID(id string) (*models.ApplicationForDorm, *errors.ErrorStruct) {
	deletedApp, err := as.applicationsRepository.DeleteApplicationByID(id)
	if err != nil {
		return nil, err
	}
	return deletedApp, nil
}

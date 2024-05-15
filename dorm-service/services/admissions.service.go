package services

import (
	"dorm-service/errors"
	"dorm-service/models"
	"dorm-service/repositories"
)

type AdmissionsService struct {
	admissionsRepository *repositories.AdmissionsRepository
}

func NewAdmissionsServices(repository *repositories.AdmissionsRepository) (*AdmissionsService, error) {
	return &AdmissionsService{
		admissionsRepository: repository,
	}, nil
}

func (as AdmissionsService) CreateNewAdmission(admission models.DormitoryAdmissions) (*models.DormitoryAdmissions, *errors.ErrorStruct) {
	createdAdmission, err := as.admissionsRepository.SaveNewAdmission(admission)
	if err != nil {
		return nil, err
	}
	return createdAdmission, nil
}

func (as AdmissionsService) GetAdmissionById(id string) (*models.DormitoryAdmissions, *errors.ErrorStruct) {
	foundAdmission, err := as.admissionsRepository.FindAdmissionById(id)
	if err != nil {
		return nil, err
	}
	return foundAdmission, nil

}

func (as AdmissionsService) DeleteAdmissionById(id string) (*models.DormitoryAdmissions, *errors.ErrorStruct) {
	deletedAdmission, err := as.admissionsRepository.DeleteAdmissionById(id)
	if err != nil {
		return nil, err
	}
	return deletedAdmission, nil
}

func (as AdmissionsService) FindAdmissionByDormID(id string) (*models.DormitoryAdmissions, *errors.ErrorStruct) {
	foundAdmission, err := as.admissionsRepository.FindAdmissionByDormID(id)
	if err != nil {
		return nil, err
	}
	return foundAdmission, nil

}

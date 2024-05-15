package services

import (
	"dorm-service/errors"
	"dorm-service/models"
	"dorm-service/repositories"
)

type DormService struct {
	dormRepository *repositories.DormRepository
}

func NewDormService(dormRepository *repositories.DormRepository) (*DormService, error) {
	return &DormService{
		dormRepository: dormRepository,
	}, nil
}

func (ds DormService) CreateNewDorm(dorm models.Dorm) (*models.Dorm, *errors.ErrorStruct) {
	createdDorm, err := ds.dormRepository.SaveNewDorm(dorm)
	if err != nil {
		return nil, err
	}

	return createdDorm, nil
}

func (ds DormService) FindDormById(id string) (*models.Dorm, *errors.ErrorStruct) {
	dorm, err := ds.dormRepository.FindDormById(id)
	if err != nil {
		return nil, err
	}
	return dorm, nil
}

func (ds DormService) DeleteDormById(id string) (*models.Dorm, *errors.ErrorStruct) {
	dorm, err := ds.dormRepository.DeleteDormById(id)
	if err != nil {
		return nil, err
	}
	return dorm, nil
}

func (ds DormService) UpdateDormById(id, name, location string) (*models.Dorm, *errors.ErrorStruct) {
	dorm, err := ds.dormRepository.UpdateDorm(id, name, location)
	if err != nil {
		return nil, err
	}
	return dorm, nil
}

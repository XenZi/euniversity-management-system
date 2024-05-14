package services

import (
	"fakultet-service/errors"
	"fakultet-service/models"
	"fakultet-service/repository"
)

type UniversityService struct {
	UniversityRepository *repository.UniversityRepository
}

func NewUniversityService(universityRepository *repository.UniversityRepository) (*UniversityService, error) {
	return &UniversityService{
		UniversityRepository: universityRepository,
	}, nil
}

func (u UniversityService) CreateUniversity(university models.University) (*models.University, *errors.ErrorStruct) {
	addedUni, err := u.UniversityRepository.SaveUniversity(university)
	if err != nil {
		return nil, err
	}
	return addedUni, nil
}

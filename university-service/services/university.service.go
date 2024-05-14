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

func (u UniversityService) CreateStudent(student models.Student) (*models.Student, *errors.ErrorStruct) {
	student.Espb = 0
	addedStud, err := u.UniversityRepository.SaveStudent(student)
	if err != nil {
		return nil, err
	}
	return addedStud, nil
}

func (u UniversityService) FindStudentById(personalIdentificationNumber string) (*models.Student, *errors.ErrorStruct) {
	student, err := u.UniversityRepository.FindStudentById(personalIdentificationNumber)
	if err != nil {
		return nil, err
	}
	return student, nil
}
func (u UniversityService) CheckBudget(personalIdentificationNumber string) (bool, *errors.ErrorStruct) {
	student, err := u.UniversityRepository.FindStudentById(personalIdentificationNumber)
	if err != nil {
		return false, err
	}
	return student.BudgetStatus, nil
}

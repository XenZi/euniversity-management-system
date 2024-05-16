package services

import (
	"fakultet-service/client"
	"fakultet-service/errors"
	"fakultet-service/models"
	"fakultet-service/repository"
	"fmt"
)

type UniversityService struct {
	UniversityRepository *repository.UniversityRepository
	HealthCareClient     *client.HealthCareClient
}

func NewUniversityService(universityRepository *repository.UniversityRepository, client *client.HealthCareClient) (*UniversityService, error) {
	return &UniversityService{
		UniversityRepository: universityRepository,
		HealthCareClient:     client,
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
func (u UniversityService) ExtendStatus(personalIdentificationNumber string) (*models.Student, *errors.ErrorStruct) {
	student, err := u.UniversityRepository.FindStudentById(personalIdentificationNumber)
	fmt.Println(student)
	if err != nil {
		return nil, err
	}

	healthStatusConfirmed, err := u.HealthCareClient.GetUserHealthStatusConfirmation(personalIdentificationNumber)
	if err != nil {
		return nil, err
	}

	if !healthStatusConfirmed {
		return nil, errors.NewError("Health status confirmation failed", 400)
	}
	student.Semester += 1
	updatedStudent, err := u.UniversityRepository.UpdateStudent(*student)
	if err != nil {
		return nil, err
	}

	return updatedStudent, nil
}
func (u UniversityService) DeleteStudent(personalIdentificationNumber string) (*models.Student, *errors.ErrorStruct) {
	student, err := u.UniversityRepository.DeleteStudent(personalIdentificationNumber)
	if err != nil {
		return nil, err
	}
	return student, nil
}

package services

import (
	"fakultet-service/client"
	"fakultet-service/errors"
	"fakultet-service/models"
	"fakultet-service/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	student.Semester = 1
	student.Roles = append(student.Roles, "Citizen", "Student")
	student.Status = 0
	addedStud, err := u.UniversityRepository.SaveStudent(student)
	if err != nil {
		return nil, err
	}
	return addedStud, nil
}

func (u UniversityService) CreateProfessor(professor models.Professor) (*models.Professor, *errors.ErrorStruct) {
	addedProf, err := u.UniversityRepository.SaveProfessor(professor)
	if err != nil {
		return nil, err
	}
	return addedProf, nil
}

func (u UniversityService) CreateScholarship(scholarship models.Scholarship) (*models.Scholarship, *errors.ErrorStruct) {
	addedScholarship, err := u.UniversityRepository.SaveScholarship(scholarship)
	if err != nil {
		return nil, err
	}
	return addedScholarship, nil
}

func (u UniversityService) CreateApplicationForStateExam(application models.StateExamApplication) (*models.StateExamApplication, *errors.ErrorStruct) {
	addedApplication, err := u.UniversityRepository.SaveStateExamApplication(application)
	if err != nil {
		return nil, err
	}
	return addedApplication, nil
}

func (u UniversityService) CreateRandomEntranceExam(exam models.EntranceExam) (*models.EntranceExam, *errors.ErrorStruct) {
	addedExam, err := u.UniversityRepository.SaveEntranceExam(exam)
	if err != nil {
		return nil, err
	}
	return addedExam, nil
}

func (u UniversityService) CreateExtendStatusApplication(application models.ExtendStatusApplication) (*models.ExtendStatusApplication, *errors.ErrorStruct) {
	addedApplication, err := u.UniversityRepository.SaveExtendStatusApplication(application)
	if err != nil {
		return nil, err
	}
	return addedApplication, nil
}

func (u UniversityService) CreateScholarshipApplication(application models.ApplyForScholarship) (*models.ApplyForScholarship, *errors.ErrorStruct) {
	addedApplication, err := u.UniversityRepository.SaveScholarshipApplication(application)
	if err != nil {
		return nil, err
	}
	return addedApplication, nil
}

func (u UniversityService) FindStudentById(personalIdentificationNumber string) (*models.Student, *errors.ErrorStruct) {
	student, err := u.UniversityRepository.FindStudentById(personalIdentificationNumber)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (u UniversityService) FindProfessorById(personalIdentificationNumber string) (*models.Professor, *errors.ErrorStruct) {
	professor, err := u.UniversityRepository.FindProfessor(personalIdentificationNumber)
	if err != nil {
		return nil, err
	}
	return professor, nil
}
func (u UniversityService) FindUniversityById(id primitive.ObjectID) (*models.University, *errors.ErrorStruct) {
	university, err := u.UniversityRepository.FindUniversityById(id)
	if err != nil {
		return nil, err
	}
	return university, nil
}

func (u UniversityService) FindAllUniversities() ([]*models.University, *errors.ErrorStruct) {
	getAllUniversities, err := u.UniversityRepository.FindAllUniversities()
	if err != nil {
		return nil, err
	}
	return getAllUniversities, nil
}

func (u UniversityService) FindAllExams() ([]*models.EntranceExam, *errors.ErrorStruct) {
	getAllEntranceExams, err := u.UniversityRepository.FindAllEntranceExams()
	if err != nil {
		return nil, err
	}
	return getAllEntranceExams, nil
}

func (u UniversityService) FindAllExtendStatusApplications() ([]*models.ExtendStatusApplication, *errors.ErrorStruct) {
	getAllExtendStatusApplications, err := u.UniversityRepository.FindAllExtendStatusApplications()
	if err != nil {
		return nil, err
	}
	return getAllExtendStatusApplications, nil
}

func (u UniversityService) FindAllScholarshipApplications() ([]*models.ApplyForScholarship, *errors.ErrorStruct) {
	getAllScholarshipApplications, err := u.UniversityRepository.FindAllScholarshipApplications()
	if err != nil {
		return nil, err
	}
	return getAllScholarshipApplications, nil
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
	student.Espb += 20
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

func (u UniversityService) DeleteProfessor(personalIdentificationNumber string) (*models.Professor, *errors.ErrorStruct) {
	professor, err := u.UniversityRepository.DeleteProfessor(personalIdentificationNumber)
	if err != nil {
		return nil, err
	}
	return professor, nil
}

func (u UniversityService) DeleteScholarship(id primitive.ObjectID) (*models.Scholarship, *errors.ErrorStruct) {
	scholarship, err := u.UniversityRepository.DeleteScholarship(id)
	if err != nil {
		return nil, err
	}
	return scholarship, nil
}

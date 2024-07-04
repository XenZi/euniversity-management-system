package services

import (
	"dorm-service/errors"
	"dorm-service/models"
	"dorm-service/repositories"
	"fmt"
)

type AdmissionsService struct {
	admissionsRepository *repositories.AdmissionsRepository
	applicationsService  *ApplicationsService
	roomService          *RoomService
}

func NewAdmissionsServices(repository *repositories.AdmissionsRepository, appService *ApplicationsService, roomService *RoomService) (*AdmissionsService, error) {
	return &AdmissionsService{
		admissionsRepository: repository,
		applicationsService:  appService,
		roomService:          roomService,
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

func (as AdmissionsService) GetAllAdmissions() ([]*models.DormitoryAdmissions, *errors.ErrorStruct) {
	getAllAdmissions, err := as.admissionsRepository.FindAdmissions()
	if err != nil {
		return nil, err
	}
	return getAllAdmissions, nil
}

func (as AdmissionsService) UpdateAdmission(admission models.DormitoryAdmissions) (*models.DormitoryAdmissions, *errors.ErrorStruct) {
	updatedAdmission, err := as.admissionsRepository.UpdateAdmission(admission)
	if err != nil {
		return nil, err
	}
	return updatedAdmission, nil
}

func (as AdmissionsService) EndAdmission(admissionID string) (*models.DormitoryAdmissions, *errors.ErrorStruct) {
	foundAdmission, err := as.admissionsRepository.FindAdmissionById(admissionID)
	if err != nil {
		return nil, err
	}
	applications, err := as.applicationsService.FindAllAcceptedApplicationsForSpecifiedAdmission(admissionID)
	if err != nil {
		return nil, err
	}
	rooms, err := as.roomService.GetAllRoomsForDormID(foundAdmission.DormID)
	if err != nil {
		return nil, err
	}
	totalAvailablePlaces := 0
	for _, room := range rooms {
		totalAvailablePlaces += int(room.NumberOfBeds)
	}
	for i := range applications {
		placed := false
		for j := range rooms {
			if int16(len(rooms[j].Students)) < rooms[j].NumberOfBeds {
				rooms[j].Students = append(rooms[j].Students, applications[i].Student)

				_, err := as.roomService.AppendStudentToRoom(rooms[j].ID.Hex(), applications[i].Student)
				if err != nil {
					fmt.Printf("Failed to append student %s to room %s: %v\n", applications[i].Student.FullName, rooms[j].ID.Hex(), err)
					continue
				}

				placed = true
				break
			}
		}
		if !placed {
			fmt.Println("Not enough beds available for student:", applications[i].Student.FullName)
		}
	}
	return foundAdmission, nil
}

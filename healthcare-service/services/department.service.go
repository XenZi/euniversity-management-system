package services

import (
	"fmt"
	"healthcare/clients"
	"healthcare/errors"
	"healthcare/models"
	"healthcare/repository"

	"time"
)

type DepartmentService struct {
	DepartmentRepository *repository.DepartmentRepository
	HealthcareService    *HealthcareService
	DTOService           *DTOService
	universityClient     *clients.UniversityClient
}

func NewDepartmentService(repo *repository.DepartmentRepository, uniClient *clients.UniversityClient, serv *HealthcareService) (*DepartmentService, error) {
	return &DepartmentService{
		DepartmentRepository: repo,
		DTOService:           NewDTOService(),
		universityClient:     uniClient,
		HealthcareService:    serv,
	}, nil
}

func (d DepartmentService) CreateDepartment(name string) (*models.DepartmentDTO, *errors.ErrorStruct) {
	newDepartment := models.Department{
		Name:     name,
		Schedule: createSchedule(),
	}
	addedDept, err := d.DepartmentRepository.SaveDepartment(newDepartment)
	if err != nil {
		return nil, err
	}
	return d.DTOService.DeptToDeptDTO(*addedDept), nil
}

func (d DepartmentService) GetDepartmentByName(name string) (*models.DepartmentDTO, *errors.ErrorStruct) {
	foundDept, err := d.DepartmentRepository.GetDepartmentByName(name)
	if err != nil {
		return nil, err
	}
	return d.DTOService.DeptToDeptDTO(*foundDept), nil

}

func (d DepartmentService) AddDoctorToSlot(name, id, date, resTime string) (*models.DepartmentDTO, *errors.ErrorStruct) {
	foundDept, err := d.DepartmentRepository.GetDepartmentByName(name)
	if err != nil {
		return nil, err
	}
	for _, slot := range foundDept.Schedule.Date[date] {
		if slot.Time == resTime {
			slot.DoctorID = id
		}
	}
	updated, err := d.DepartmentRepository.UpdateDepartment(*foundDept)
	if err != nil {
		return nil, err
	}
	return d.DTOService.DeptToDeptDTO(*updated), nil
}

func (d DepartmentService) AddPatientToSlot(name, id, date, resTime, appType string) (*models.DepartmentDTO, *errors.ErrorStruct) {
	_, err := d.universityClient.CheckIfStudent(id)
	if err != nil {
		return nil, err
	}
	foundDept, err := d.DepartmentRepository.GetDepartmentByName(name)
	if err != nil {
		return nil, err
	}
	var doctorId string
	for _, slot := range foundDept.Schedule.Date[date] {
		if slot.Time == resTime {
			if slot.PatientID == "" {
				slot.PatientID = id
				doctorId = slot.DoctorID
			}
		}
	}

	dateTime := fmt.Sprintf("%s %s", date, resTime)
	appointment := models.AppointmentSchedule{
		DateOfIssue:     dateTime,
		PatientID:       id,
		DoctorID:        doctorId,
		AppointmentType: getTypeEnum(appType),
	}
	_, err = d.HealthcareService.CreateAppointment(appointment)
	updated, err := d.DepartmentRepository.UpdateDepartment(*foundDept)
	if err != nil {
		return nil, err
	}
	return d.DTOService.DeptToDeptDTO(*updated), nil
}

func (d DepartmentService) GetFreeSlots(name, date string) ([]string, *errors.ErrorStruct) {
	foundDept, err := d.DepartmentRepository.GetDepartmentByName(name)
	if err != nil {
		return nil, err
	}
	var freeSlots []string
	for _, slot := range foundDept.Schedule.Date[date] {
		if slot.DoctorID != "" && slot.PatientID == "" {
			freeSlots = append(freeSlots, slot.Time)
		}
	}
	return freeSlots, nil
}
func (d DepartmentService) GetFreeDoctorSlots(name, date string) ([]string, *errors.ErrorStruct) {
	foundDept, err := d.DepartmentRepository.GetDepartmentByName(name)
	if err != nil {
		return nil, err
	}
	var freeSlots []string
	for _, slot := range foundDept.Schedule.Date[date] {
		if slot.DoctorID == "" {
			freeSlots = append(freeSlots, slot.Time)
		}
	}
	return freeSlots, nil
}

func createSchedule() models.Schedule {
	currentDate := time.Now()
	endDate := currentDate.AddDate(0, 0, 5)
	var dateSlice []string
	for d := currentDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		formattedDate := d.Format("02.01.2006")
		dateSlice = append(dateSlice, formattedDate)
	}
	startTime := time.Date(2024, 6, 30, 8, 0, 0, 0, time.UTC)
	endTime := time.Date(2024, 6, 30, 13, 0, 0, 0, time.UTC)
	var times []string
	var slots []models.Slot
	for t := startTime; t.Before(endTime); t = t.Add(15 * time.Minute) {
		times = append(times, t.Format("15:04"))
	}
	for _, timeIns := range times {
		s := models.Slot{
			Time: timeIns,
		}
		slots = append(slots, s)
	}
	dates := make(map[string][]models.Slot)
	for _, date := range dateSlice {
		dates[date] = slots
	}
	schedule := models.Schedule{Date: dates}
	return schedule
}

func getTypeEnum(typeE string) models.EAppointmentType {
	switch typeE {
	case "INTERVENTION":
		return models.Intervention
	default:
		return models.Examination
	}
}

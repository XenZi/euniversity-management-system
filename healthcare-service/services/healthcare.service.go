package services

import (
	"healthcare/clients"
	"healthcare/errors"
	"healthcare/models"
	"healthcare/repository"
)

type HealthcareService struct {
	HealthcareRepository *repository.HealthcareRepository
	DtoServ              *DTOService
	universityClient     *clients.UniversityClient
}

func NewHealthcareService(healthcareRepository *repository.HealthcareRepository, uniClient *clients.UniversityClient) (*HealthcareService, error) {
	return &HealthcareService{
		HealthcareRepository: healthcareRepository,
		DtoServ:              NewDTOService(),
		universityClient:     uniClient,
	}, nil
}

func (h HealthcareService) CreateRecordForUser(patientID string) (*models.RecordDTO, *errors.ErrorStruct) {
	clientBool := true // needs student services communication
	if !clientBool {
		return nil, errors.NewError("patient is not a student", 405)
	}
	newRecord := models.Record{
		PatientID: patientID,
	}
	addedRecord, err := h.HealthcareRepository.SaveRecord(newRecord)
	if err != nil {
		return nil, err
	}
	return h.DtoServ.RecToRecDTO(*addedRecord), nil
}

func (h HealthcareService) CreateCertificateForUser(r models.CompletionReport) (*models.CertificateDTO, *errors.ErrorStruct) {
	_, err := h.universityClient.CheckIfStudent(r.PatientID)
	if err != nil {
		return nil, err
	}
	report := models.Report{
		Title:   r.Title,
		Content: r.Content,
	}
	insertedReport, err := h.HealthcareRepository.SaveReport(report)
	if err != nil {
		return nil, err
	}
	certificate := models.Certificate{
		DateOfIssue: r.DateOfIssue,
		PatientID:   r.PatientID,
		DoctorID:    r.DoctorID,
		Report:      *insertedReport,
	}
	addedCertificate, err := h.HealthcareRepository.SaveCertificate(certificate)
	if err != nil {
		return nil, err
	}
	return h.DtoServ.CertToCertDTO(*addedCertificate), nil
}

func (h HealthcareService) GetRecordForUser(id string) (*models.RecordDTO, *errors.ErrorStruct) {
	record, err := h.HealthcareRepository.GetRecordByPatientID(id)
	if err != nil {
		return nil, err
	}
	return h.DtoServ.RecToRecDTO(*record), nil
}

func (h HealthcareService) GetCertificateForUser(id string) (*models.CertificateDTO, *errors.ErrorStruct) {
	record, err := h.HealthcareRepository.GetRecordByPatientID(id)
	if err != nil {
		return nil, err
	}
	if record.Certificate != (models.Certificate{}) {
		return h.DtoServ.CertToCertDTO(record.Certificate), nil
	}
	return nil, errors.NewError("no certificate found", 418) // I'm a teapot
}

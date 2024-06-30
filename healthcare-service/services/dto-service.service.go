package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"healthcare/errors"
	"healthcare/models"
)

type DTOService struct {
}

func NewDTOService() *DTOService {
	return &DTOService{}
}

func (d DTOService) CertToCertDTO(c models.Certificate) *models.CertificateDTO {
	return &models.CertificateDTO{
		ID:          c.ID.Hex(),
		DateOfIssue: c.DateOfIssue,
		PatientID:   c.PatientID,
		DoctorID:    c.DoctorID,
		Report:      c.Report,
	}
}

func (d DTOService) CertDTOToCert(c models.CertificateDTO) (*models.Certificate, *errors.ErrorStruct) {
	id, err := primitive.ObjectIDFromHex(c.ID)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	return &models.Certificate{
		ID:          id,
		DateOfIssue: c.DateOfIssue,
		PatientID:   c.PatientID,
		DoctorID:    c.DoctorID,
		Report:      c.Report,
	}, nil
}

func (d DTOService) RecToRecDTO(r models.Record) *models.RecordDTO {
	return &models.RecordDTO{
		ID:            r.ID.Hex(),
		PatientID:     r.PatientID,
		Certificate:   r.Certificate,
		Prescriptions: r.Prescriptions,
		Referrals:     r.Referrals,
		Appointments:  r.Appointments,
	}
}

func (d DTOService) RecDTOToRec(r models.RecordDTO) (*models.Record, error) {
	id, err := primitive.ObjectIDFromHex(r.ID)
	if err != nil {
		return nil, err
	}
	return &models.Record{
		ID:            id,
		PatientID:     r.PatientID,
		Certificate:   r.Certificate,
		Prescriptions: r.Prescriptions,
		Referrals:     r.Referrals,
		Appointments:  r.Appointments,
	}, nil
}

func (d DTOService) RefToRefDTO(r models.Referral) *models.ReferralDTO {
	return &models.ReferralDTO{
		ID:          r.ID.Hex(),
		DateOfIssue: r.DateOfIssue,
		PatientID:   r.PatientID,
		DoctorID:    r.DoctorID,
	}
}

func (d DTOService) RefDTOToRef(r models.ReferralDTO) (*models.Referral, error) {
	id, err := primitive.ObjectIDFromHex(r.ID)
	if err != nil {
		return nil, err
	}
	return &models.Referral{
		ID:          id,
		DateOfIssue: r.DateOfIssue,
		PatientID:   r.PatientID,
		DoctorID:    r.DoctorID,
	}, nil
}

func (d DTOService) PresToPresDTO(r models.Prescription) *models.PrescriptionDTO {
	return &models.PrescriptionDTO{
		ID:          r.ID.Hex(),
		DateOfIssue: r.DateOfIssue,
		PatientID:   r.PatientID,
		DoctorID:    r.DoctorID,
		Drug:        r.Drug,
		Form:        r.Form,
		Dosage:      r.Dosage,
		Status:      r.Status,
	}
}

func (d DTOService) PresDTOToPres(r models.PrescriptionDTO) (*models.Prescription, error) {
	id, err := primitive.ObjectIDFromHex(r.ID)
	if err != nil {
		return nil, err
	}
	return &models.Prescription{
		ID:          id,
		DateOfIssue: r.DateOfIssue,
		PatientID:   r.PatientID,
		DoctorID:    r.DoctorID,
		Drug:        r.Drug,
		Form:        r.Form,
		Dosage:      r.Dosage,
		Status:      r.Status,
	}, nil
}

func (d DTOService) AppToAppDTO(r models.Appointment) *models.AppointmentDTO {
	return &models.AppointmentDTO{
		ID:                r.ID.Hex(),
		DateOfIssue:       r.DateOfIssue,
		PatientID:         r.PatientID,
		DoctorID:          r.DoctorID,
		AppointmentType:   r.AppointmentType,
		AppointmentStatus: r.AppointmentStatus,
		Report:            r.Report,
	}
}

func (d DTOService) AppDTOToApp(r models.AppointmentDTO) (*models.Appointment, error) {
	id, err := primitive.ObjectIDFromHex(r.ID)
	if err != nil {
		return nil, err
	}
	return &models.Appointment{
		ID:                id,
		DateOfIssue:       r.DateOfIssue,
		PatientID:         r.PatientID,
		DoctorID:          r.DoctorID,
		AppointmentType:   r.AppointmentType,
		AppointmentStatus: r.AppointmentStatus,
		Report:            r.Report,
	}, nil
}

func (d DTOService) DeptToDeptDTO(r models.Department) *models.DepartmentDTO {
	return &models.DepartmentDTO{
		ID:       r.ID.Hex(),
		Name:     r.Name,
		Schedule: r.Schedule,
	}
}

func (d DTOService) DeptDTOToDept(r models.DepartmentDTO) (*models.Department, error) {
	id, err := primitive.ObjectIDFromHex(r.ID)
	if err != nil {
		return nil, err
	}
	return &models.Department{
		ID:       id,
		Name:     r.Name,
		Schedule: r.Schedule,
	}, nil
}

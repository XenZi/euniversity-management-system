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
		//AppointmentStatus: c.AppointmentStatus,
		Report: c.Report,
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
		//AppointmentStatus: c.AppointmentStatus,
		Report: c.Report,
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

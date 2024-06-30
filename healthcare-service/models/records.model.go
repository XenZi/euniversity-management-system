package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Record struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	PatientID     string             `json:"patientID" bson:"patientID"`
	Certificate   Certificate        `json:"certificate" bson:"certificate"`
	Prescriptions []Prescription     `json:"prescriptions" bson:"prescriptions"`
	Referrals     []Referral         `json:"referrals" bson:"referrals"`
	Appointments  []Appointment      `json:"appointments" bson:"appointments"`
}

type RecordDTO struct {
	ID            string         `json:"id"`
	PatientID     string         `json:"patientID"`
	Certificate   Certificate    `json:"certificate"`
	Prescriptions []Prescription `json:"prescriptions"`
	Referrals     []Referral     `json:"referrals"`
	Appointments  []Appointment  `json:"appointments"`
}

type Certificate struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DateOfIssue string             `json:"dateOfIssue" bson:"dateOfIssue"`
	PatientID   string             `json:"patientID" bson:"patientID"`
	DoctorID    string             `json:"doctorID" bson:"doctorID"`
	Report      Report             `json:"report" bson:"report"`
}
type CertificateDTO struct {
	ID          string `json:"id"`
	DateOfIssue string `json:"dateOfIssue"`
	PatientID   string `json:"patientID"`
	DoctorID    string `json:"doctorID"`
	Report      Report `json:"report"`
}

type Prescription struct {
	ID          primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	DateOfIssue string              `json:"dateOfIssue" bson:"dateOfIssue"`
	PatientID   string              `json:"patientID" bson:"patientID"`
	DoctorID    string              `json:"doctorID" bson:"doctorID"`
	Drug        string              `json:"drug" bson:"drug"`
	Form        EForm               `json:"form" bson:"form"`
	Dosage      string              `json:"dosage" bson:"dosage"`
	Status      EPrescriptionStatus `json:"prescriptionStatus" bson:"prescriptionStatus"`
}

type PrescriptionDTO struct {
	ID          string              `json:"id"`
	DateOfIssue string              `json:"dateOfIssue" `
	PatientID   string              `json:"patientID"`
	DoctorID    string              `json:"doctorID" `
	Drug        string              `json:"drug"`
	Form        EForm               `json:"form"`
	Dosage      string              `json:"dosage"`
	Status      EPrescriptionStatus `json:"prescriptionStatus"`
}

type PrescriptionInfo struct {
	PatientID string `json:"patientID"`
	DoctorID  string `json:"doctorID" `
	Drug      string `json:"drug"`
	Form      string `json:"form"`
	Dosage    string `json:"dosage"`
	Status    string `json:"prescriptionStatus"`
}

type Referral struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DateOfIssue string             `json:"dateOfIssue" bson:"dateOfIssue"`
	PatientID   string             `json:"patientID" bson:"patientID"`
	DoctorID    string             `json:"doctorID" bson:"doctorID"`
}
type ReferralDTO struct {
	ID          string `json:"id"`
	DateOfIssue string `json:"dateOfIssue"`
	PatientID   string `json:"patientID"`
	DoctorID    string `json:"doctorID"`
}

type ReferralInfo struct {
	PatientID string `json:"patientID"`
	DoctorID  string `json:"doctorID"`
}

type Appointment struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DateOfIssue       string             `json:"dateOfIssue" bson:"dateOfIssue"`
	PatientID         string             `json:"patientID" bson:"patientID"`
	DoctorID          string             `json:"doctorID" bson:"doctorID"`
	AppointmentType   EAppointmentType   `json:"appointmentType" bson:"appointmentType"`
	AppointmentStatus EAppointmentStatus `json:"appointmentStatus" bson:"appointmentStatus"`
	Report            Report             `json:"report" bson:"report"`
}
type AppointmentDTO struct {
	ID                string             `json:"id"`
	DateOfIssue       string             `json:"dateOfIssue"`
	PatientID         string             `json:"patientID"`
	DoctorID          string             `json:"doctorID" `
	AppointmentType   EAppointmentType   `json:"appointmentType"`
	AppointmentStatus EAppointmentStatus `json:"appointmentStatus"`
	Report            Report             `json:"report"`
}

type Report struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Content     string             `json:"content" bson:"content"`
	DateOfIssue string             `json:"dateOfIssue" bson:"dateOfIssue"`
}

type CompletionReport struct {
	AppointmentID string `json:"AppointmentID"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	PatientID     string `json:"patientID"`
	DoctorID      string `json:"doctorID"`
}

type AppointmentSchedule struct {
	PatientID       string           `json:"patientID"`
	DoctorID        string           `json:"doctorID"`
	DateOfIssue     string           `json:"dateOfIssue"`
	AppointmentType EAppointmentType `json:"appointmentType"`
}

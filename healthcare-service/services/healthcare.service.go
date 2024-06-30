package services

import (
	"healthcare/clients"
	"healthcare/errors"
	"healthcare/models"
	"healthcare/repository"
	"time"
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

// RECORDS

func (h HealthcareService) CreateRecordForUser(patientID string) (*models.RecordDTO, *errors.ErrorStruct) {
	newRecord := models.Record{
		PatientID: patientID,
	}
	addedRecord, err := h.HealthcareRepository.SaveRecord(newRecord)
	if err != nil {
		return nil, err
	}
	return h.DtoServ.RecToRecDTO(*addedRecord), nil
}

func (h HealthcareService) GetRecordForUser(id string) (*models.RecordDTO, *errors.ErrorStruct) {
	record, err := h.HealthcareRepository.GetRecordByPatientID(id)
	if err != nil {
		return nil, err
	}
	return h.DtoServ.RecToRecDTO(*record), nil
}

func (h HealthcareService) GetAllRecords() ([]*models.RecordDTO, *errors.ErrorStruct) {
	recs, err := h.HealthcareRepository.GetAllRecords()
	if err != nil {
		return nil, err
	}
	var ret []*models.RecordDTO
	for _, ent := range recs {
		ret = append(ret, h.DtoServ.RecToRecDTO(*ent))
	}
	return ret, nil
}

// CERTIFICATES

func (h HealthcareService) CreateCertificate(r models.CompletionReport) (*models.CertificateDTO, *errors.ErrorStruct) {
	_, err := h.universityClient.CheckIfStudent(r.PatientID)
	if err != nil {
		return nil, err
	}
	report := models.Report{
		Title:       r.Title,
		Content:     r.Content,
		DateOfIssue: getNow(),
	}
	insertedReport, err := h.HealthcareRepository.SaveReport(report)
	if err != nil {
		return nil, err
	}
	certificate := models.Certificate{
		DateOfIssue: getNow(),
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

func (h HealthcareService) GetCertificateByPatientId(id string) (*models.CertificateDTO, *errors.ErrorStruct) {
	record, err := h.HealthcareRepository.GetRecordByPatientID(id)
	if err != nil {
		return nil, err
	}
	if record.Certificate != (models.Certificate{}) {
		return h.DtoServ.CertToCertDTO(record.Certificate), nil
	}
	return nil, errors.NewError("no certificate found", 418) // I'm a teapot
}

//APPOINTMENTS

// Already checked for student status
func (h HealthcareService) CreateAppointment(r models.AppointmentSchedule) (*models.AppointmentDTO, *errors.ErrorStruct) {
	appointment := models.Appointment{
		DateOfIssue:       r.DateOfIssue,
		PatientID:         r.PatientID,
		DoctorID:          r.DoctorID,
		AppointmentType:   r.AppointmentType,
		AppointmentStatus: models.Scheduled,
	}
	saved, err := h.HealthcareRepository.SaveAppointment(appointment)
	if err != nil {
		return nil, err
	}
	return h.DtoServ.AppToAppDTO(*saved), nil
}

func (h HealthcareService) UpdateAppointment(r models.CompletionReport) (*models.AppointmentDTO, *errors.ErrorStruct) {
	report := models.Report{
		Title:       r.Title,
		Content:     r.Content,
		DateOfIssue: getNow(),
	}
	insertedReport, err := h.HealthcareRepository.SaveReport(report)
	if err != nil {
		return nil, err
	}
	appointment, err := h.HealthcareRepository.GetAppointmentByID(r.AppointmentID)
	if err != nil {
		return nil, err
	}
	appointment.AppointmentStatus = models.Completed
	appointment.Report = *insertedReport
	appointment.DoctorID = r.DoctorID

	updatedApp, err := h.HealthcareRepository.UpdateAppointment(*appointment)
	if err != nil {
		return nil, err
	}
	return h.DtoServ.AppToAppDTO(*updatedApp), nil
}

func (h HealthcareService) GetAppointmentsByDoctorID(doctorID string) ([]*models.AppointmentDTO, *errors.ErrorStruct) {
	apps, err := h.HealthcareRepository.GetAllAppointmentsByDoctorID(doctorID)
	if err != nil {
		return nil, err
	}
	var appointments []*models.AppointmentDTO
	for _, app := range apps {
		appointments = append(appointments, h.DtoServ.AppToAppDTO(*app))
	}
	return appointments, nil
}

func (h HealthcareService) GetAppointmentsByPatientID(patientID string) ([]*models.AppointmentDTO, *errors.ErrorStruct) {
	reco, err := h.HealthcareRepository.GetRecordByPatientID(patientID)
	if err != nil {
		return nil, err
	}
	var appointments []*models.AppointmentDTO
	for _, app := range reco.Appointments {
		appointments = append(appointments, h.DtoServ.AppToAppDTO(app))
	}
	return appointments, nil
}

func (h HealthcareService) GetAppointmentById(id string) (*models.AppointmentDTO, *errors.ErrorStruct) {
	appo, err := h.HealthcareRepository.GetAppointmentByID(id)
	if err != nil {
		return nil, err
	}
	return h.DtoServ.AppToAppDTO(*appo), nil
}

// REFERALS

func (h HealthcareService) CreateReferral(r models.ReferralInfo) (*models.ReferralDTO, *errors.ErrorStruct) {
	_, err := h.universityClient.CheckIfStudent(r.PatientID)
	if err != nil {
		return nil, err
	}
	ref := models.Referral{
		DateOfIssue: getNow(),
		PatientID:   r.PatientID,
		DoctorID:    r.DoctorID,
	}
	insertedRef, err := h.HealthcareRepository.SaveReferral(ref)
	if err != nil {
		return nil, err
	}
	return h.DtoServ.RefToRefDTO(*insertedRef), nil
}

func (h HealthcareService) GetReferralsByPatientId(patientID string) ([]*models.ReferralDTO, *errors.ErrorStruct) {
	reco, err := h.HealthcareRepository.GetRecordByPatientID(patientID)
	if err != nil {
		return nil, err
	}
	var referrals []*models.ReferralDTO
	for _, app := range reco.Referrals {
		referrals = append(referrals, h.DtoServ.RefToRefDTO(app))
	}
	return referrals, nil
}

func (h HealthcareService) GetReferralById(id string) (*models.ReferralDTO, *errors.ErrorStruct) {
	refe, err := h.HealthcareRepository.GetReferralByID(id)
	if err != nil {
		return nil, err
	}
	return h.DtoServ.RefToRefDTO(*refe), nil
}

// PRESCRIPTIONS

func (h HealthcareService) CreatePrescription(r models.PrescriptionInfo) (*models.PrescriptionDTO, *errors.ErrorStruct) {
	prescription := models.Prescription{
		DateOfIssue: getNow(),
		PatientID:   r.PatientID,
		DoctorID:    r.DoctorID,
		Drug:        r.Drug,
		Form:        getFormEnum(r.Form),
		Dosage:      r.Dosage,
		Status:      getStatusEnum(r.Status),
	}
	insertedPres, err := h.HealthcareRepository.SavePrescription(prescription)
	if err != nil {
		return nil, err
	}
	return h.DtoServ.PresToPresDTO(*insertedPres), nil
}

func (h HealthcareService) GetPrescriptionsByPatientId(patientID string) ([]*models.PrescriptionDTO, *errors.ErrorStruct) {
	reco, err := h.HealthcareRepository.GetRecordByPatientID(patientID)
	if err != nil {
		return nil, err
	}
	var list []*models.PrescriptionDTO
	for _, app := range reco.Prescriptions {
		list = append(list, h.DtoServ.PresToPresDTO(app))
	}
	return list, nil
}

func (h HealthcareService) GetPrescriptionById(id string) (*models.PrescriptionDTO, *errors.ErrorStruct) {
	pres, err := h.HealthcareRepository.GetPrescriptionByID(id)
	if err != nil {
		return nil, err
	}
	return h.DtoServ.PresToPresDTO(*pres), nil
}

func (h HealthcareService) UpdatePrescription(id, status string) (*models.PrescriptionDTO, *errors.ErrorStruct) {
	pres, err := h.HealthcareRepository.GetPrescriptionByID(id)
	if err != nil {
		return nil, err
	}
	pres.Status = getStatusEnum(status)
	presc, err := h.HealthcareRepository.UpdatePrescription(*pres)
	if err != nil {
		return nil, err
	}
	return h.DtoServ.PresToPresDTO(*presc), nil
}

// Utils

func getNow() string {
	currentDate := time.Now()
	formattedDate := currentDate.Format("02.01.2006 15:04")
	return formattedDate
}

func getFormEnum(form string) models.EForm {
	switch form {
	case "SYRUP":
		return models.Syrup
	case "GEL":
		return models.Gel
	case "CAPSULE":
		return models.Capsule
	default:
		return models.Tablet
	}
}

func getStatusEnum(status string) models.EPrescriptionStatus {
	switch status {
	case "ISSUED":
		return models.Issued
	case "ISSUED_REPEAT":
		return models.IssuedRepeat
	case "CLAIMED_REPEAT":
		return models.ClaimedRepeat
	default:
		return models.Claimed
	}
}

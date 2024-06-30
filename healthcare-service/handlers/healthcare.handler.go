package handlers

import (
	"github.com/gorilla/mux"
	"healthcare/models"
	"healthcare/services"
	"healthcare/utils"
	"net/http"
)

type HealthcareHandler struct {
	HealthcareService *services.HealthcareService
	DepartmentService *services.DepartmentService
}

func NewHealthcareHandler(healthcareService *services.HealthcareService, deptService *services.DepartmentService) (*HealthcareHandler, error) {
	return &HealthcareHandler{
		HealthcareService: healthcareService,
		DepartmentService: deptService,
	}, nil
}

func (hh HealthcareHandler) Ping(rw http.ResponseWriter, h *http.Request) {
	utils.WriteResp(models.BaseHttpResponse{
		Data: "OK",
	}, 200, rw)
}

// RECORDS

func (hh HealthcareHandler) CreateRecordForUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	record, err := hh.HealthcareService.CreateRecordForUser(id)
	if err != nil { // may just end up skipping an error response
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/createRecord", rw)
		return
	}
	utils.WriteResp(record, 200, rw)
}

func (hh HealthcareHandler) GetRecordForUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	record, err := hh.HealthcareService.GetRecordForUser(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/getRecord", rw)
		return
	}
	utils.WriteResp(record, 200, rw)
}

// CERTIFICATES

func (hh HealthcareHandler) CreateCertificateForUser(rw http.ResponseWriter, h *http.Request) {
	var report models.CompletionReport
	if !utils.DecodeJSONFromRequest(h, rw, &report) {
		utils.WriteErrorResp("error while casting data into structure", 500, "api/healthcare/createCertificate", rw)
		return
	}
	response, err := hh.HealthcareService.CreateCertificate(report)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/createCertificate", rw)
		return
	}
	utils.WriteResp(response, 200, rw)
}

func (hh HealthcareHandler) GetCertificateForUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	certificate, err := hh.HealthcareService.GetCertificateByPatientId(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/getCertificate", rw)
		return
	}
	utils.WriteResp(certificate, 200, rw)
}

func (hh HealthcareHandler) CreateDepartment(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) GetDepartmentByName(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) GetFreeSlots(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) GetFreeDoctorSlots(rw http.ResponseWriter, r *http.Request) {

}
func (hh HealthcareHandler) AddDoctorToSchedule(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) AddPatientToSchedule(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) UpdateAppointment(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) GetAppointmentsByDoctorID(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) GetAppointmentsByPatientID(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) GetAppointmentById(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) CreateReferral(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) GetReferralsByPatientID(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) GetReferralById(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) CreatePrescription(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) GetPrescriptionsByPatientID(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) GetPrescriptionById(rw http.ResponseWriter, h *http.Request) {

}
func (hh HealthcareHandler) UpdatePrescription(rw http.ResponseWriter, h *http.Request) {

}

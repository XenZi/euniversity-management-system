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
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/record/{id}", rw)
		return
	}
	utils.WriteResp(record, 200, rw)
}

func (hh HealthcareHandler) GetAllRecords(rw http.ResponseWriter, h *http.Request) {
	resp, err := hh.HealthcareService.GetAllRecords()
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/records", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}

// CERTIFICATES

func (hh HealthcareHandler) CreateCertificateForUser(rw http.ResponseWriter, h *http.Request) {
	var report models.CompletionReport
	if !utils.DecodeJSONFromRequest(h, rw, &report) {
		utils.WriteErrorResp("error while casting data into structure", 500, "api/healthcare/record/{id}/createCertificate", rw)
		return
	}
	response, err := hh.HealthcareService.CreateCertificate(report)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/record/{id}/createCertificate", rw)
		return
	}
	utils.WriteResp(response, 200, rw)
}

func (hh HealthcareHandler) GetCertificateForUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	certificate, err := hh.HealthcareService.GetCertificateByPatientId(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/record/{id}/certificate", rw)
		return
	}
	utils.WriteResp(certificate, 200, rw)
}

// DEPARTMENT

func (hh HealthcareHandler) CreateDepartment(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	name := vars["name"]
	dept, err := hh.DepartmentService.CreateDepartment(name)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/department/create/{name}", rw)
		return
	}
	utils.WriteResp(dept, 200, rw)
}
func (hh HealthcareHandler) GetDepartmentByName(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	name := vars["name"]
	dept, err := hh.DepartmentService.GetDepartmentByName(name)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/department/{name}", rw)
		return
	}
	utils.WriteResp(dept, 200, rw)
}

// SCHEDULE

func (hh HealthcareHandler) GetFreeSlots(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	name := vars["name"]
	date := vars["date"]
	dept, err := hh.DepartmentService.GetFreeSlots(name, date)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/department/{name}/schedule/{date}", rw)
		return
	}
	utils.WriteResp(dept, 200, rw)
}
func (hh HealthcareHandler) GetFreeDoctorSlots(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	name := vars["name"]
	date := vars["date"]
	dept, err := hh.DepartmentService.GetFreeDoctorSlots(name, date)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/department/{name}/docSchedule/{date}", rw)
		return
	}
	utils.WriteResp(dept, 200, rw)
}
func (hh HealthcareHandler) AddDoctorToSchedule(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	name := vars["name"]
	date := vars["date"]
	var slotFill *models.SlotFill
	if !utils.DecodeJSONFromRequest(h, rw, &slotFill) {
		utils.WriteErrorResp("wrong structure", 500, "api/healthcare/department/{name}/docSchedule/{date}/", rw)
		return
	}
	dept, err := hh.DepartmentService.AddDoctorToSlot(name, slotFill.ID, date, slotFill.Time)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/department/{name}/docSchedule/{date}/", rw)
		return
	}
	utils.WriteResp(dept, 200, rw)
}
func (hh HealthcareHandler) AddPatientToSchedule(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	name := vars["name"]
	date := vars["date"]
	var slotFill *models.SlotFill
	if !utils.DecodeJSONFromRequest(h, rw, &slotFill) {
		utils.WriteErrorResp("wrong structure", 500, "api/healthcare/department/{name}/docSchedule/{date}/", rw)
		return
	}
	dept, err := hh.DepartmentService.AddPatientToSlot(name, slotFill.ID, date, slotFill.Time, slotFill.AppType)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/department/{name}/schedule/{date}/", rw)
		return
	}
	utils.WriteResp(dept, 200, rw)
}
func (hh HealthcareHandler) GetAppointmentsByDoctorID(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	resp, err := hh.HealthcareService.GetAppointmentsByDoctorID(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/appointments/{id}", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}
func (hh HealthcareHandler) GetAppointmentsByPatientID(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	resp, err := hh.HealthcareService.GetAppointmentsByPatientID(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/record/{id}/appointments", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}
func (hh HealthcareHandler) GetAppointmentById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["appId"]
	resp, err := hh.HealthcareService.GetAppointmentById(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/record/{id}/appointments/{appId}", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}
func (hh HealthcareHandler) UpdateAppointment(rw http.ResponseWriter, h *http.Request) {
	var report models.CompletionReport
	if !utils.DecodeJSONFromRequest(h, rw, &report) {
		utils.WriteErrorResp("error while casting data into structure", 500, "api/healthcare/record/{id}/appointments/{appId}", rw)
		return
	}
	response, err := hh.HealthcareService.UpdateAppointment(report)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/record/{id}/appointments/{appId}", rw)
		return
	}
	utils.WriteResp(response, 200, rw)
}
func (hh HealthcareHandler) GetReferralsByPatientID(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	resp, err := hh.HealthcareService.GetReferralsByPatientId(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/record/{id}/referrals", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}
func (hh HealthcareHandler) CreateReferral(rw http.ResponseWriter, h *http.Request) {
	var temp models.ReferralInfo
	if !utils.DecodeJSONFromRequest(h, rw, &temp) {
		utils.WriteErrorResp("error while casting data into structure", 500, "api/healthcare/record/{id}/referrals/createReferral", rw)
		return
	}
	response, err := hh.HealthcareService.CreateReferral(temp)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/record/{id}/referrals/createReferral", rw)
		return
	}
	utils.WriteResp(response, 200, rw)
}
func (hh HealthcareHandler) GetReferralById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["refId"]
	resp, err := hh.HealthcareService.GetReferralById(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/record/{id}/referrals/{refId}", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}
func (hh HealthcareHandler) GetPrescriptionsByPatientID(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	resp, err := hh.HealthcareService.GetPrescriptionsByPatientId(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/record/{id}/prescriptions", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}
func (hh HealthcareHandler) CreatePrescription(rw http.ResponseWriter, h *http.Request) {
	var temp models.PrescriptionInfo
	if !utils.DecodeJSONFromRequest(h, rw, &temp) {
		utils.WriteErrorResp("error while casting data into structure", 500, "api/healthcare/record/{id}/prescriptions/createPrescription", rw)
		return
	}
	response, err := hh.HealthcareService.CreatePrescription(temp)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/record/{id}/prescriptions/createPrescription", rw)
		return
	}
	utils.WriteResp(response, 200, rw)
}
func (hh HealthcareHandler) GetPrescriptionById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["presId"]
	resp, err := hh.HealthcareService.GetPrescriptionById(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/record/{id}/prescriptions/{presId}", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}
func (hh HealthcareHandler) UpdatePrescription(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["presId"]
	status := vars["status"]
	resp, err := hh.HealthcareService.UpdatePrescription(id, status)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/record/{id}/prescriptions/{presId}/{status}", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}

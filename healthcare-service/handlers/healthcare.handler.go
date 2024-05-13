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
}

func NewHealthcareHandler(healthcareService *services.HealthcareService) (*HealthcareHandler, error) {
	return &HealthcareHandler{
		HealthcareService: healthcareService,
	}, nil
}

func (hh HealthcareHandler) Ping(rw http.ResponseWriter, h *http.Request) {
	utils.WriteResp(models.BaseHttpResponse{
		Data: "OK",
	}, 200, rw)
}

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

func (hh HealthcareHandler) CreateCertificateForUser(rw http.ResponseWriter, h *http.Request) {
	var report models.CompletionReport
	if !utils.DecodeJSONFromRequest(h, rw, &report) {
		utils.WriteErrorResp("error while casting data into structure", 500, "api/healthcare/createCertificate", rw)
		return
	}
	response, err := hh.HealthcareService.CreateCertificateForUser(report)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/createCertificate", rw)
		return
	}
	utils.WriteResp(response, 200, rw)
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

func (hh HealthcareHandler) GetCertificateForUser(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	certificate, err := hh.HealthcareService.GetCertificateForUser(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/healthcare/getCertificate", rw)
		return
	}
	utils.WriteResp(certificate, 200, rw)
}

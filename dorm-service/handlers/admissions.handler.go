package handlers

import (
	"dorm-service/models"
	"dorm-service/services"
	"dorm-service/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AdmissionsHandler struct {
	admissionsService *services.AdmissionsService
}

func NewAdmissionsHandler(admissionsService *services.AdmissionsService) (*AdmissionsHandler, error) {
	return &AdmissionsHandler{
		admissionsService: admissionsService,
	}, nil
}

func (ah AdmissionsHandler) CreateNewAdmission(rw http.ResponseWriter, h *http.Request) {
	var admissions models.DormitoryAdmissions
	if !utils.DecodeJSONFromRequest(h, rw, &admissions) {
		utils.WriteErrorResp("Error neki", 500, "path", rw)
		return
	}
	ctx := h.Context()
	val := ctx.Value("user")
	log.Println(val)
	data, err := ah.admissionsService.CreateNewAdmission(admissions)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "path", rw)
		return
	}
	utils.WriteResp(data, 201, rw)
}

func (ah AdmissionsHandler) GetAdmissionByID(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
		return
	}
	data, err := ah.admissionsService.GetAdmissionById(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "path", rw)
		return
	}
	utils.WriteResp(data, 201, rw)
}

func (ah AdmissionsHandler) DeleteAdmissionById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
		return
	}

	data, err := ah.admissionsService.DeleteAdmissionById(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "path", rw)
		return
	}
	utils.WriteResp(data, 201, rw)
}

func (ah AdmissionsHandler) GetAdmissionByDormId(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
		return
	}
	data, err := ah.admissionsService.FindAdmissionByDormID(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "path", rw)
		return
	}
	utils.WriteResp(data, 201, rw)

}

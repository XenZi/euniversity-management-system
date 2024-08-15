package handlers

import (
	"dorm-service/models"
	"dorm-service/services"
	"dorm-service/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type DormApplicationHandler struct {
	ApplicationsService *services.ApplicationsService
}

func NewDormApplicationHandler(ApplicationsService *services.ApplicationsService) (*DormApplicationHandler, error) {
	return &DormApplicationHandler{
		ApplicationsService: ApplicationsService,
	}, nil
}

func (ah DormApplicationHandler) CreateNewApplication(rw http.ResponseWriter, h *http.Request) {
	var application models.ApplicationForDorm
	if !utils.DecodeJSONFromRequest(h, rw, &application) {
		utils.WriteErrorResp("Error neki", 500, "path", rw)
		return
	}
	log.Println(application)
	data, err := ah.ApplicationsService.CreateNewApplication(application)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "path", rw)
		return
	}
	utils.WriteResp(data, 201, rw)
}

func (ah DormApplicationHandler) FindAllApplicationsByUserPIN(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	pin := vars["pin"]
	if pin == "" {
		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
		return
	}
	data, err := ah.ApplicationsService.FindApplicationsByUserPIN(pin)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "applications/{pin}", rw)
		return
	}
	utils.WriteResp(data, 200, rw)
}

func (ah DormApplicationHandler) FindApplicationByID(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
		return
	}
	data, err := ah.ApplicationsService.FindApplicationByID(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "applications/{id}", rw)
		return
	}
	utils.WriteResp(data, 200, rw)
}

// func (ah DormApplicationHandler) FindAllApplicationsByID(rw http.ResponseWriter, h *http.Request) {
// 	vars := mux.Vars(h)
// 	id := vars["id"]
// 	if id == "" {
// 		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
// 		return
// 	}
// 	data, err := ah.ApplicationsService.FindApplicationByID(id)
// 	if err != nil {
// 		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "applications/{id}", rw)
// 		return
// 	}
// 	utils.WriteResp(data, 200, rw)
// }

func (ah DormApplicationHandler) FindAllApplicationsForOneAdmission(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
		return
	}
	data, err := ah.ApplicationsService.FindAllAplicationsForOneAdmission(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "applications/{id}", rw)
		return
	}
	utils.WriteResp(data, 200, rw)
}

func (ah DormApplicationHandler) FindAllUnderReviewApplicationsForSpecifiedAdmission(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
		return
	}
	data, err := ah.ApplicationsService.FindAllUnderReviewApplicationsForSpecifiedAdmission(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "applications/{id}", rw)
		return
	}
	utils.WriteResp(data, 200, rw)
}

func (ah DormApplicationHandler) FindAllAcceptedApplicationsForSpecifiedAdmission(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
		return
	}
	data, err := ah.ApplicationsService.FindAllAcceptedApplicationsForSpecifiedAdmission(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "applications/{id}", rw)
		return
	}
	utils.WriteResp(data, 200, rw)
}

func (ah DormApplicationHandler) UpdateApplication(rw http.ResponseWriter, h *http.Request) {
	var app models.ApplicationForDorm
	if !utils.DecodeJSONFromRequest(h, rw, &app) {
		utils.WriteErrorResp("Error while casting", 500, "path", rw)
		return
	}
	updatedApp, err := ah.ApplicationsService.UpdateApplication(app)
	if err != nil {
		utils.WriteErrorResp("Error neki", 500, "path", rw)
		return
	}
	utils.WriteResp(updatedApp, 201, rw)
}

func (ah DormApplicationHandler) DeleteApplicationByID(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
		return
	}
	data, err := ah.ApplicationsService.DeleteApplicationByID(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "applications/{id}", rw)
		return
	}
	utils.WriteResp(data, 200, rw)
}

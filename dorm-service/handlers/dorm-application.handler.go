package handlers

import (
	"dorm-service/models"
	"dorm-service/services"
	"dorm-service/utils"
	"log"
	"net/http"
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

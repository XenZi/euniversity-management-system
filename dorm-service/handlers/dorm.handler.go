package handlers

import (
	"dorm-service/models"
	"dorm-service/services"
	"dorm-service/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type DormHandler struct {
	dormService *services.DormService
}

func NewDormHandler(dormService *services.DormService) (*DormHandler, error) {
	return &DormHandler{
		dormService: dormService,
	}, nil
}

func (ah DormHandler) Ping(rw http.ResponseWriter, h *http.Request) {
	utils.WriteResp(models.BaseHttpResponse{
		Data: "OK",
	}, 200, rw)
}

func (ah DormHandler) CreateNewDorm(rw http.ResponseWriter, h *http.Request) {
	var dorm models.Dorm
	if !utils.DecodeJSONFromRequest(h, rw, &dorm) {
		utils.WriteErrorResp("Error neki", 500, "path", rw)
		return
	}
	createdDorm, err := ah.dormService.CreateNewDorm(dorm)
	if err != nil {
		utils.WriteErrorResp("Error neki", 500, "path", rw)
		return
	}
	utils.WriteResp(createdDorm, 201, rw)
}

func (ah DormHandler) FindDormById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
		return
	}
	dorm, err := ah.dormService.FindDormById(id)
	if err != nil {
		utils.WriteResp(err.GetErrorMessage(), err.GetErrorStatus(), rw)
		return
	}
	utils.WriteResp(dorm, 200, rw)
}

func (ah DormHandler) DeleteDormById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
		return
	}
	dorm, err := ah.dormService.DeleteDormById(id)
	if err != nil {
		utils.WriteResp(err.GetErrorMessage(), err.GetErrorStatus(), rw)
		return
	}
	utils.WriteResp(dorm, 200, rw)
}

func (ah DormHandler) UpdateDormById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
		return
	}
	var dormDTO models.DormDTO
	if !utils.DecodeJSONFromRequest(h, rw, &dormDTO) {
		utils.WriteErrorResp("Bad request", 400, "api/confirm-account", rw)
		return
	}
	dorm, err := ah.dormService.UpdateDormById(id, dormDTO.Name, dormDTO.Location)
	if err != nil {
		utils.WriteResp(err.GetErrorMessage(), err.GetErrorStatus(), rw)
		return
	}
	utils.WriteResp(dorm, 200, rw)
}

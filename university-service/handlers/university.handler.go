package handlers

import (
	"encoding/json"
	"fakultet-service/models"
	"fakultet-service/services"
	"fakultet-service/utils"
	"net/http"
)

type UniversityHandler struct {
	UniversityService *services.UniversityService
}

func NewUniversityHandler(universityService *services.UniversityService) (*UniversityHandler, error) {
	return &UniversityHandler{
		UniversityService: universityService,
	}, nil
}
func (uh UniversityHandler) Ping(rw http.ResponseWriter, h *http.Request) {
	utils.WriteResp(models.BaseHttpResponse{
		Data: "OK",
	}, 200, rw)
}

func (uh UniversityHandler) CreateUniversity(rw http.ResponseWriter, h *http.Request) {
	decoder := json.NewDecoder(h.Body)
	decoder.DisallowUnknownFields()
	var university models.University
	if err := decoder.Decode(&university); err != nil {
		utils.WriteResp(err.Error(), http.StatusBadRequest, rw)
		return
	}
	newUni, err := uh.UniversityService.CreateUniversity(university)
	if err != nil {
		utils.WriteResp(err.GetErrorMessage(), err.GetErrorStatus(), rw)
		return
	}
	utils.WriteResp(newUni, 200, rw)
}

func (uh UniversityHandler) CreateStudent(rw http.ResponseWriter, h *http.Request) {
	decoder := json.NewDecoder(h.Body)
	decoder.DisallowUnknownFields()
	var student models.Student
	if err := decoder.Decode(&student); err != nil {
		utils.WriteResp(err.Error(), http.StatusBadRequest, rw)
		return
	}
	newStudent, err := uh.UniversityService.CreateStudent(student)
	if err != nil {
		utils.WriteResp(err.GetErrorMessage(), err.GetErrorStatus(), rw)
		return
	}
	utils.WriteResp(newStudent, 200, rw)
}

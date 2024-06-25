package handlers

import (
	"encoding/json"
	"fakultet-service/models"
	"fakultet-service/services"
	"fakultet-service/utils"
	"github.com/gorilla/mux"
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

func (uh UniversityHandler) CreateProfessor(rw http.ResponseWriter, h *http.Request) {
	decoder := json.NewDecoder(h.Body)
	decoder.DisallowUnknownFields()
	var professor models.Professor
	if err := decoder.Decode(&professor); err != nil {
		utils.WriteResp(err.Error(), http.StatusBadRequest, rw)
		return
	}
	newProfessor, err := uh.UniversityService.CreateProfessor(professor)
	if err != nil {
		utils.WriteResp(err.GetErrorMessage(), err.GetErrorStatus(), rw)
		return
	}
	utils.WriteResp(newProfessor, 200, rw)
}

func (uh UniversityHandler) CreateScholarship(rw http.ResponseWriter, h *http.Request) {
	decoder := json.NewDecoder(h.Body)
	decoder.DisallowUnknownFields()
	var scholarship models.Scholarship
	if err := decoder.Decode(&scholarship); err != nil {
		utils.WriteResp(err.Error(), http.StatusBadRequest, rw)
		return
	}
	newScholarship, err := uh.UniversityService.CreateScholarship(scholarship)
	if err != nil {
		utils.WriteResp(err.GetErrorMessage(), err.GetErrorStatus(), rw)
		return
	}
	utils.WriteResp(newScholarship, 200, rw)
}

func (uh UniversityHandler) FindStudentById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteResp("Bad request", http.StatusNotFound, rw)
		return
	}
	student, err := uh.UniversityService.FindStudentById(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "/student", rw)
		return
	}
	utils.WriteResp(student, 200, rw)
}

func (uh UniversityHandler) FindProfessorById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteResp("Bad request", http.StatusNotFound, rw)
		return
	}
	professor, err := uh.UniversityService.FindProfessorById(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "/professor", rw)
		return
	}
	utils.WriteResp(professor, 200, rw)

}

func (uh UniversityHandler) CheckBudget(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteResp("Bad request", http.StatusNotFound, rw)
		return
	}
	resp, err := uh.UniversityService.CheckBudget(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "/student/budget", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}
func (uh UniversityHandler) ExtendStatus(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteResp("Bad request", http.StatusNotFound, rw)
		return
	}
	resp, err := uh.UniversityService.ExtendStatus(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "/student/status", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}
func (uh UniversityHandler) DeleteStudent(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteResp("Bad request", http.StatusNotFound, rw)
		return
	}
	resp, err := uh.UniversityService.DeleteStudent(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "/student", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}

func (uh UniversityHandler) DeleteProfessor(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteResp("Bad request", http.StatusNotFound, rw)
		return
	}
	resp, err := uh.UniversityService.DeleteProfessor(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "/professor", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}

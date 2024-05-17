package handlers

import (
	"auth/models"
	"auth/services"
	"auth/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	AuthService *services.AuthService
	JwtService  *services.JwtService
}

func NewAuthHandler(authService *services.AuthService, jwtService *services.JwtService) (*AuthHandler, error) {
	return &AuthHandler{
		AuthService: authService,
		JwtService:  jwtService,
	}, nil
}

func (ah AuthHandler) Ping(rw http.ResponseWriter, h *http.Request) {
	utils.WriteResp(models.BaseHttpResponse{
		Data: "OK",
	}, 200, rw)
}

func (ah AuthHandler) Register(rw http.ResponseWriter, h *http.Request) {
	var registerUser models.Citizen
	if !utils.DecodeJSONFromRequest(h, rw, &registerUser) {
		utils.WriteErrorResp("Error while casting data into structure", 500, "/api/auth/register", rw)
		return
	}
	response, err := ah.AuthService.CreateUser(registerUser)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/auth/register", rw)
		return
	}
	utils.WriteResp(response, 200, rw)
}

func (ah AuthHandler) Login(rw http.ResponseWriter, h *http.Request) {
	var loginUser models.LoginCitizenDTO
	if !utils.DecodeJSONFromRequest(h, rw, &loginUser) {
		utils.WriteErrorResp("Error while casting data into structure", 500, "/api/auth/register", rw)
		return
	}
	response, err := ah.AuthService.LoginUser(loginUser)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/auth/login", rw)
		return
	}
	utils.WriteResp(response, 200, rw)
}

func (ah AuthHandler) GetUserByPIN(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	resp, err := ah.AuthService.GetUserByPIN(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/auth/findOne", rw)
		return
	}
	utils.WriteResp(resp, 200, rw)
}

func (ah AuthHandler) AddRoles(rw http.ResponseWriter, h *http.Request) {
	var addingRoles models.AddingRoles
	decoder := json.NewDecoder(h.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&addingRoles); err != nil {
		utils.WriteErrorResp("error while casting data into structure", 500, "/api/auth/addRoles", rw)
		return
	}
	response, err := ah.AuthService.UpdateUserByPIN(addingRoles.PersonalIdentificationNumber, addingRoles.Roles)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/auth/addRoles", rw)
		return
	}
	utils.WriteResp(response, 200, rw)
}

func (ah AuthHandler) SwitchRole(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	pin := vars["id"]
	role := vars["role"]

	response, err := ah.AuthService.SwitchRoles(pin, role)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/auth/login", rw)
		return
	}
	utils.WriteResp(response, 200, rw)
}

func (ah AuthHandler) ValidateJWT(rw http.ResponseWriter, h *http.Request) {
	tokenString := utils.ExtractToken(h.Header.Get("Authorization"))
	log.Println(tokenString)
	response, err := ah.JwtService.ValidateToken(tokenString)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/auth/login", rw)
		return
	}
	log.Println(response)
	ctx := context.WithValue(h.Context(), "user", response)
	h = h.WithContext(ctx)
	utils.WriteResp(response, 200, rw)
}

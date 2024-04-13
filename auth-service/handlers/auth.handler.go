package handlers

import (
	"auth/models"
	"auth/services"
	"auth/utils"
	"net/http"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) (*AuthHandler, error) {
	return &AuthHandler{
		AuthService: authService,
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

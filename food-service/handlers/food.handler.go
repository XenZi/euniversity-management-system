package handlers

import (
	"food/models"
	"food/services"
	"food/utils"
	"net/http"
)

type FoodHandler struct {
	FoodService *services.FoodService
}

func NewFoodHandler(foodService *services.FoodService) (*FoodHandler, error) {
	return &FoodHandler{
		FoodService: foodService,
	}, nil
}

func (f FoodHandler) Ping(rw http.ResponseWriter, r *http.Request) {
	utils.WriteResp(models.BaseHttpResponse{
		Data: "pong",
	}, 200, rw)
}

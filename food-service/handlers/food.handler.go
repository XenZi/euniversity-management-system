package handlers

import (
	"food/models"
	"food/services"
	"food/utils"
	"github.com/gorilla/mux"
	"log"
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

func (f FoodHandler) Ping(rw http.ResponseWriter, h *http.Request) {
	utils.WriteResp(models.BaseHttpResponse{
		Data: "pong",
	}, 200, rw)
}

func (f FoodHandler) CreateFoodCard(rw http.ResponseWriter, h *http.Request) {
	var card models.FoodCard

	if !utils.DecodeJSONFromRequest(h, rw, &card) {
		utils.WriteErrorResp("Error while casting into structure", 500, "/api/food/createFoodCard", rw)
		return
	}
	response, err := f.FoodService.CreateFoodCardForUser(card)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/food/createFoodCard", rw)
		return
	}
	utils.WriteResp(response, 200, rw)

}

func (f FoodHandler) GetAllFoodCards(rw http.ResponseWriter, h *http.Request) {
	// Call the service method to get all food cards
	foodCards, err := f.FoodService.GetAllFoodCards()
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "/api/food/getAllFoodCards", rw)
		return
	}
	log.Println("Kartice za hranu su ", foodCards)
	// Encode the retrieved food cards into JSON format

	// Set the response content type to JSON
	rw.Header().Set("Content-Type", "application/json")

	utils.WriteResp(foodCards, 200, rw)
}

func (f FoodHandler) CreatePayment(rw http.ResponseWriter, h *http.Request) {
	var payment models.Payment

	if !utils.DecodeJSONFromRequest(h, rw, &payment) {
		utils.WriteErrorResp("Error while casting into structure", 500, "/api/food/createPayment", rw)
		return
	}
	response, err := f.FoodService.CreatePayment(payment)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/food/createPayment", rw)
		return
	}
	utils.WriteResp(response, 200, rw)

}
func (f FoodHandler) PayForMeal(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/food/payForMeal", rw)
	}
	card, err := f.FoodService.PayForMeal(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/food/payForMeal", rw)
	}
	utils.WriteResp(card, 200, rw)
}

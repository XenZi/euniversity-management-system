package handlers

import (
	"food/models"
	"food/services"
	"food/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

// MESS ROOM CRUD
func (f FoodHandler) CreateMessRoom(rw http.ResponseWriter, h *http.Request) {
	var messRoom models.MessRoom

	if !utils.DecodeJSONFromRequest(h, rw, &messRoom) {
		utils.WriteErrorResp("Error while casting into structure", 500, "/api/food/createMessRoom", rw)
		return
	}
	response, err := f.FoodService.CreateMessRoom(messRoom)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/food/createMessRoom", rw)
		return
	}
	utils.WriteResp(response, 200, rw)
}

func (f FoodHandler) GetAllMessRooms(rw http.ResponseWriter, h *http.Request) {
	messRooms, err := f.FoodService.GetAllMessRooms()
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/food/getAllMessRooms", rw)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	utils.WriteResp(messRooms, 200, rw)
}

func (f FoodHandler) DeleteMessRoom(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	massDeleted, err := f.FoodService.DeleteMessRoom(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/food/deleteMessRoom", rw)
	}
	utils.WriteResp(massDeleted, 200, rw)

}

func (f FoodHandler) UpdateMessRoom(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	var messUpdated models.MessRoomUpdate
	if !utils.DecodeJSONFromRequest(h, rw, &messUpdated) {
		utils.WriteErrorResp("Bad request", 400, "api/food/updateMess", rw)
		return
	}
	messUpdated.ID = id
	mess, err := f.FoodService.UpdateMessRoom(messUpdated)
	if err != nil {
		utils.WriteResp(err.GetErrorMessage(), err.GetErrorStatus(), rw)
		return
	}
	utils.WriteResp(mess, 200, rw)
}

// FOOD CARD CRUD

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

func (f FoodHandler) DeleteFoodCard(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	foodCardDeleted, err := f.FoodService.DeleteFoodCard(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/food/deleteMessRoom", rw)
	}
	utils.WriteResp(foodCardDeleted, 200, rw)

}

// PAYMENT CRUD

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

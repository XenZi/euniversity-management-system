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

// STUDENT CRUD

func (f FoodHandler) CreateStudent(rw http.ResponseWriter, h *http.Request) {
	var student models.Student

	if !utils.DecodeJSONFromRequest(h, rw, &student) {
		utils.WriteErrorResp("Error while casting into structure", 500, "/api/food/createStudent", rw)
		return
	}
	response, err := f.FoodService.CreateStudent(student)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/food/createStudent", rw)
		return
	}
	utils.WriteResp(response, 200, rw)
}

func (f FoodHandler) GetAllStudents(rw http.ResponseWriter, h *http.Request) {
	students, err := f.FoodService.GetAllStudents()
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/food/getAllStudents", rw)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	utils.WriteResp(students, 200, rw)
}


func (f FoodHandler) DeleteStudent(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	studentDeleted, err := f.FoodService.DeleteStudentById(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/food/deleteMessRoom", rw)
	}
	utils.WriteResp(studentDeleted, 200, rw)

}
func (f FoodHandler) GetStudentById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	studentFound, err := f.FoodService.FindStudentById(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/food/findStudentById", rw)
	}
	utils.WriteResp(studentFound, 200, rw)
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

func (f FoodHandler) UpdateMessRoomSupplier(rw http.ResponseWriter, h *http.Request) {
    vars := mux.Vars(h)
    id := vars["id"]
	supplier_id:=vars["supplierId"]
    

    mess, err := f.FoodService.UpdateMessRoomSupplier(id, supplier_id)
    if err != nil {
        utils.WriteErrorResp(err.GetErrorMessage(),err.GetErrorStatus(),"api/food/updateMess", rw)
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

// SUPPLIER CRUD

func (f FoodHandler) CreateSupplier(rw http.ResponseWriter, h *http.Request) {
	var supplier models.Supplier

	if !utils.DecodeJSONFromRequest(h, rw, &supplier) {
		utils.WriteErrorResp("Error while casting into structure", 500, "/api/food/createSupplier", rw)
		return
	}
	response, err := f.FoodService.CreateSupplier(supplier)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/food/createSupplier", rw)
		return
	}
	utils.WriteResp(response, 200, rw)

}

func (f FoodHandler) GetAllSuppliers(rw http.ResponseWriter, h *http.Request) {
	allSuppliers, err := f.FoodService.GetAllSuppliers()

	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "/api/food/getAllSuppliers", rw)
		return
	}

	utils.WriteResp(allSuppliers, 201, rw)

}

func (f FoodHandler) GetSupplierById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	supplier, err := f.FoodService.GetSupplierById(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "/api/food/getSupplierById", rw)
		return
	}
	utils.WriteResp(supplier, 201, rw)
}

func (f FoodHandler) DeleteSupplierById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	supplier, err := f.FoodService.DeleteSupllier(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "/api/food/getSupplierById", rw)
		return
	}
	utils.WriteResp(supplier, 201, rw)
}

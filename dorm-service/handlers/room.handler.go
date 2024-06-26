package handlers

import (
	"dorm-service/models"
	"dorm-service/services"
	"dorm-service/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type RoomHandler struct {
	RoomsService *services.RoomService
}

func NewRoomHandler(roomService *services.RoomService) (*RoomHandler, error) {
	return &RoomHandler{
		RoomsService: roomService,
	}, nil
}

func (rh RoomHandler) CreateRoom(rw http.ResponseWriter, h *http.Request) {
	var room models.Room
	if !utils.DecodeJSONFromRequest(h, rw, &room) {
		utils.WriteErrorResp("Error neki", 500, "path", rw)
		return
	}
	savedRoom, err := rh.RoomsService.CreateNewRoom(room)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "", rw)
		return
	}
	utils.WriteResp(savedRoom, 201, rw)
}

func (rh RoomHandler) GetAllRoomsByDormID(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/dorm/id/rooms", rw)
		return
	}
	rooms, err := rh.RoomsService.GetAllRoomsForDormID(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/dorm/id/rooms", rw)
		return
	}
	utils.WriteResp(rooms, 200, rw)
}

func (rh RoomHandler) GetRoomByID(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/dorm/room/id", rw)
		return
	}
	rooms, err := rh.RoomsService.GetRoomByID(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/dorm/room/{id}", rw)
		return
	}
	utils.WriteResp(rooms, 200, rw)
}

func (rh RoomHandler) UpdateRoom(rw http.ResponseWriter, h *http.Request) {
	var room models.Room
	if !utils.DecodeJSONFromRequest(h, rw, &room) {
		utils.WriteErrorResp("Error neki", 500, "path", rw)
		return
	}
	updatedRoom, err := rh.RoomsService.UpdateRoom(room)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/dorm/room/{id}", rw)
		return
	}
	utils.WriteResp(updatedRoom, 200, rw)
}

func (rh RoomHandler) DeleteRoom(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if id == "" {
		utils.WriteErrorResp("Bad request", 400, "api/dorm/room/id", rw)
		return
	}
	deletedRoom, err := rh.RoomsService.DeleteRoom(id)
	if err != nil {
		utils.WriteErrorResp(err.GetErrorMessage(), err.GetErrorStatus(), "api/dorm/room/id", rw)
		return
	}
	utils.WriteResp(deletedRoom, 200, rw)
}

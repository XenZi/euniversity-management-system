package services

import (
	"dorm-service/errors"
	"dorm-service/models"
	"dorm-service/repositories"
	"log"
)

type RoomService struct {
	RoomRepository *repositories.RoomRepository
}

func NewRoomService(roomRepository *repositories.RoomRepository) *RoomService {
	return &RoomService{
		RoomRepository: roomRepository,
	}
}

func (rs *RoomService) CreateNewRoom(room models.Room) (*models.Room, *errors.ErrorStruct) {
	createdRoom, err := rs.RoomRepository.SaveNewRoom(room)
	if err != nil {
		return nil, err
	}
	return createdRoom, nil
}

func (rs *RoomService) GetAllRoomsForDormID(id string) ([]*models.Room, *errors.ErrorStruct) {
	foundRooms, err := rs.RoomRepository.GetAllRoomsForDorm(id)
	if err != nil {
		return nil, err
	}
	return foundRooms, nil
}

func (rs *RoomService) GetRoomByID(id string) (*models.Room, *errors.ErrorStruct) {
	foundRoom, err := rs.RoomRepository.FindOneRoomByID(id)
	if err != nil {
		return nil, err
	}
	return foundRoom, nil
}

func (rs *RoomService) UpdateRoom(room models.Room) (*models.Room, *errors.ErrorStruct) {
	log.Println(room.ID.Hex())
	updatedRoom, err := rs.RoomRepository.UpdateRoom(room.ID.Hex(), room.SquareFoot, room.NumberOfBeds, room.Toalet)
	if err != nil {
		return nil, err
	}
	return updatedRoom, nil
}

func (rs *RoomService) DeleteRoom(id string) (*models.Room, *errors.ErrorStruct) {
	deletedRoom, err := rs.RoomRepository.DeleteRoom(id)
	if err != nil {
		return nil, err
	}
	return deletedRoom, nil
}

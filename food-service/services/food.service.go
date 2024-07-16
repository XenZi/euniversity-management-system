package services

import (
	"food/client"
	"food/errors"
	"food/models"
	"food/repository"
	"log"
)

type FoodService struct {
	FoodRepository   *repository.FoodRepository
	UniversityClient *client.UniversityClient
}

func NewFoodCardService(foodRepository *repository.FoodRepository, client *client.UniversityClient) (*FoodService, error) {
	return &FoodService{
		FoodRepository:   foodRepository,
		UniversityClient: client,
	}, nil
}

// MESS ROOM CRUD
func (fs FoodService) CreateMessRoom(messRoom models.MessRoom) (*models.MessRoom, *errors.ErrorStruct) {

	addedRecord, err := fs.FoodRepository.CreateMessRoom(messRoom)
	if err != nil {
		return nil, err
	}
	return addedRecord, nil

}

func (fs FoodService) GetAllMessRooms() ([]models.MessRoom, *errors.ErrorStruct) {
	messRooms, err := fs.FoodRepository.GetAllMessRooms()
	if err != nil {
		return nil, err
	}
	return messRooms, nil
}

func (fs FoodService) DeleteMessRoom(id string) (bool, *errors.ErrorStruct) {
	removedMass, err := fs.FoodRepository.RemoveMessRoom(id)
	if err != nil {
		return false, err
	}
	return removedMass, nil
}
func (fs FoodService) FindMessById(id string) (*models.MessRoom, *errors.ErrorStruct) {
	mess, err := fs.FoodRepository.FindMessById(id)
	if err != nil {
		return nil, err
	}
	return mess, nil
}
func (fs FoodService) UpdateMessRoom(updatedMess models.MessRoomUpdate) (*models.MessRoom, *errors.ErrorStruct) {
	mess, err := fs.FoodRepository.UpdateMessRoom(updatedMess)
	if err != nil {
		return nil, err
	}
	return mess, nil

}

// FOOD CARD CRUD

func (fs FoodService) CreateFoodCardForUser(foodCard models.FoodCard) (*models.FoodCard, *errors.ErrorStruct) {

	addedRecord, err := fs.FoodRepository.SaveFoodCard(foodCard)
	if err != nil {
		return nil, err
	}
	return addedRecord, nil
}

func (fs FoodService) GetAllFoodCards() ([]models.FoodCard, *errors.ErrorStruct) {
	// Call the repository method to fetch all food cards
	foodCards, err := fs.FoodRepository.GetAllFoodCards()
	if err != nil {
		return nil, err
	}

	// Return the retrieved food cards
	return foodCards, nil
}

func (fs FoodService) DeleteFoodCard(id string) (bool, *errors.ErrorStruct) {
	removedFoodCard, err := fs.FoodRepository.RemoveFoodCard(id)
	if err != nil {
		return false, err
	}
	return removedFoodCard, nil
}

// PAYMENT CRUD

func (fs FoodService) CreatePayment(payment models.Payment) (*models.Payment, *errors.ErrorStruct) {
	//	clientBool := true // needs student services communication
	//	if !clientBool {
	//		return nil, errors.NewError("patient is not a student", 405)
	//	}

	addedPayment, err := fs.FoodRepository.SavePayment(payment)
	if err != nil {
		return nil, err
	}
	return addedPayment, nil
}

func (fs FoodService) PayForMeal(studentPIN string) (*models.FoodCard, *errors.ErrorStruct) {
	//	clientBool := true // needs student services communication
	//	if !clientBool {
	//		return nil, errors.NewError("patient is not a student", 405)
	//	}
	isBudget, err := fs.UniversityClient.GetStudentStatus(studentPIN)
	if err != nil {
		log.Println(err.GetErrorMessage())
	}
	price := 4
	log.Println("Dal je dobro pokupilo", isBudget)
	if isBudget {
		price = 2
	}
	log.Println("Podaci koje proslijedjujemo", studentPIN, price)
	updatedCard, err := fs.FoodRepository.PayForMeal(studentPIN, price)
	if err != nil {
		return nil, err
	}
	return updatedCard, nil
}

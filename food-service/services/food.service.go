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

func (fs FoodService) CreateFoodCardForUser(foodCard models.FoodCard) (*models.FoodCard, *errors.ErrorStruct) {
	//	clientBool := true // needs student services communication
	//	if !clientBool {
	//		return nil, errors.NewError("patient is not a student", 405)
	//	}

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
	if isBudget == true {
		price = 2
	}
	updatedCard, err := fs.FoodRepository.PayForMeal(studentPIN, price)
	if err != nil {
		return nil, err
	}
	return updatedCard, nil
}

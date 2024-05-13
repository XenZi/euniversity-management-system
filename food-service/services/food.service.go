package services

import (
	"food/errors"
	"food/models"
	"food/repository"
)

type FoodService struct {
	FoodRepository *repository.FoodRepository
}

func NewFoodCardService(foodRepository *repository.FoodRepository) (*FoodService, error) {
	return &FoodService{
		FoodRepository: foodRepository,
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

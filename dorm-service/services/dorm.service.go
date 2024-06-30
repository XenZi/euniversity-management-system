package services

import (
	"dorm-service/errors"
	"dorm-service/models"
	"dorm-service/repositories"
	"log"
)

type DormService struct {
	dormRepository *repositories.DormRepository
}

func NewDormService(dormRepository *repositories.DormRepository) (*DormService, error) {
	return &DormService{
		dormRepository: dormRepository,
	}, nil
}

func (ds DormService) CreateNewDorm(dorm models.Dorm) (*models.Dorm, *errors.ErrorStruct) {
	createdDorm, err := ds.dormRepository.SaveNewDorm(dorm)
	if err != nil {
		return nil, err
	}

	return createdDorm, nil
}

func (ds DormService) FindDormById(id string) (*models.Dorm, *errors.ErrorStruct) {
	dorm, err := ds.dormRepository.FindDormById(id)
	if err != nil {
		return nil, err
	}
	return dorm, nil
}

func (ds DormService) DeleteDormById(id string) (*models.Dorm, *errors.ErrorStruct) {
	dorm, err := ds.dormRepository.DeleteDormById(id)
	if err != nil {
		return nil, err
	}
	return dorm, nil
}

func (ds DormService) UpdateDormById(castedDorm models.DormDTO) (*models.Dorm, *errors.ErrorStruct) {
	dorm, err := ds.dormRepository.UpdateDorm(castedDorm.ID, castedDorm.Name, castedDorm.Location, castedDorm.Prices)
	if err != nil {
		return nil, err
	}
	return dorm, nil
}

func (ds DormService) GetAllDorms() ([]*models.Dorm, *errors.ErrorStruct) {
	dorms, err := ds.dormRepository.FindAllDorms()
	log.Println(dorms)
	if err != nil {
		return nil, err
	}
	return dorms, nil
}

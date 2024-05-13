package handlers

import "dorm-service/services"

type DormApplicationHandler struct {
	admissionsService *services.AdmissionsService
}

func NewDormApplicationHandler(admissionsService *services.AdmissionsService) (*DormApplicationHandler, error) {
	return &DormApplicationHandler{
		admissionsService: admissionsService,
	}, nil
}

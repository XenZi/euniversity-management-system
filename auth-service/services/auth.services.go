package services

import "auth/repository"

type AuthService struct {
	AuthRepository *repository.AuthRepository
	JwtService     *JwtService
}

func NewAuthService(authRepository *repository.AuthRepository, jwtService *JwtService) (*AuthService, error) {
	return &AuthService{
		AuthRepository: authRepository,
		JwtService:     jwtService,
	}, nil
}

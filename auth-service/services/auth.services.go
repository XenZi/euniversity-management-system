package services

import (
	"auth/errors"
	"auth/models"
	"auth/repository"
)

type AuthService struct {
	AuthRepository  *repository.AuthRepository
	JwtService      *JwtService
	PasswordService *PasswordService
}

func NewAuthService(authRepository *repository.AuthRepository, jwtService *JwtService, passwordService *PasswordService) (*AuthService, error) {
	return &AuthService{
		AuthRepository:  authRepository,
		JwtService:      jwtService,
		PasswordService: passwordService,
	}, nil
}

func (a AuthService) CreateUser(registerCitizen models.Citizen) (*models.CitizenDTO, *errors.ErrorStruct) {
	hashedPassword, err := a.PasswordService.HashPassword(registerCitizen.Password)

	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	registerCitizen.Roles = append(registerCitizen.Roles, "Citizen")
	registerCitizen.Password = hashedPassword
	newUser, newErr := a.AuthRepository.SaveUser(registerCitizen)
	if newErr != nil {
		return nil, newErr
	}
	castedCitizen := a.convertCitizenToDTO(*newUser)
	return &castedCitizen, nil
}

func (a AuthService) LoginUser(loginUser models.LoginCitizenDTO) (*models.SuccessfullyLoggedUser, *errors.ErrorStruct) {
	user, err := a.AuthRepository.FindUserByEmail(loginUser.Email)
	if err != nil {
		return nil, err
	}
	isSamePassword := a.PasswordService.CheckPasswordHash(loginUser.Password, user.Password)
	if !isSamePassword {
		return nil, errors.NewError("Bad credentials", 401)
	}
	jwtToken, foundError := a.JwtService.CreateKey(user.Email, user.Roles, user.PersonalIdentificationNumber)
	if foundError != nil {
		return nil, errors.NewError(foundError.Error(), 500)
	}
	return &models.SuccessfullyLoggedUser{
		Token: *jwtToken,
		User:  a.convertCitizenToDTO(*user),
	}, nil
}

func (as AuthService) convertCitizenToDTO(c models.Citizen) models.CitizenDTO {
	return models.CitizenDTO{
		ID:                           c.ID.Hex(), // Convert ObjectID to string
		FullName:                     c.FullName,
		Gender:                       c.Gender,
		IdentityCardNumber:           c.IdentityCardNumber,
		Citizenship:                  c.Citizenship,
		PersonalIdentificationNumber: c.PersonalIdentificationNumber,
		Residence:                    c.Residence,
		BirthData:                    c.BirthData,
		Email:                        c.Email,
		Roles:                        c.Roles,
	}
}

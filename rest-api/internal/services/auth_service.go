package services

import (
	"memo-point-com/internal/models"
)

type AuthService struct {
	UserService *UserService
}

func (service *AuthService) Login(username, password string) (*models.User, error) {
	user, err := service.UserService.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}

	user, err = service.UserService.RefreshAuthToken(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *AuthService) Signup(username, email, password string) (*models.User, error) {
	err := service.UserService.CreateUser(username, email, password)
	if err != nil {
		return nil, err
	}

	user, err := service.UserService.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

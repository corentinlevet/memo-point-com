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

	if user.Password != password {
		return nil, err
	}

	return user, nil
}

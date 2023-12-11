package services

import (
	"errors"
	"memo-point-com/internal/database"
	"memo-point-com/internal/models"
)

type UserService struct{}

func (service *UserService) CreateUser(username, email, password string) error {
	toCheck, _ := service.GetUserByEmail(email)
	if toCheck != nil {
		return errors.New("Email already exists")
	}

	salt, err := database.GenerateSalt()
	if err != nil {
		return err
	}

	hashedPassword, err := database.HashPassword(password, salt)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Salt:     salt,
	}

	err = database.DbInstance.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (service *UserService) GetUsers() ([]models.User, error) {
	var users []models.User

	err := database.DbInstance.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := database.DbInstance.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserService) GetUserByUsernameAndPassword(username, password string) (*models.User, error) {
	users, err := service.GetUsers()

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			hashedPassword, err := database.HashPassword(password, user.Salt)
			if err != nil {
				return nil, err
			}

			if user.Password == hashedPassword {
				return &user, nil
			}
		}
	}

	return nil, nil
}

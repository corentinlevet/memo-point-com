package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenService struct{}

const secretKey = "eYIuY8Hworcb9nTMNssY9wefz2FmIcgMaOTGk3jdxcNrcHtJy6p8gg468HC2eVGa"

func (service *TokenService) GenerateTokenForUser(username, email string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (service *TokenService) ValidateToken(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}

package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type LoginServiceInterface interface {
	GenerateToken(email string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginService struct {
	secretKey string
	issure    string
}

func LoginServiceAuth() LoginServiceInterface {
	return &LoginService{
		secretKey: "secret",
		issure:    "Petar",
	}
}

func (service *LoginService) GenerateToken(email string) (string, error) {
	claims := &authCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (service *LoginService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("Invalid token")
		}
		return []byte(service.secretKey), nil
	})
}

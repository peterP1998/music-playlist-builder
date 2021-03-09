package service

import (
	"time"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type LoginServiceInterface interface {
	GenerateToken(email string, isUser bool) string
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

func (service *LoginService) GenerateToken(email string, isUser bool) string {
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
		panic(err)
	}
	return t
}

func (service *LoginService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("Invalid token")

		}
		return []byte(service.secretKey), nil
	})

}

package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type AppClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func SignToken(id string) (string, error) {
	claims := AppClaims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("hola"))
	if err != nil {
		return nil, err
	}
	return tokenString, nil
}

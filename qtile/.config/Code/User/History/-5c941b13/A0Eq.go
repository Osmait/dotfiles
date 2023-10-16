package helper

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func DecodeJwt(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := ctx.GetHeader("Authorization")
	fmt.Println(tokenString)
	if tokenString == "" {
		return nil, errors.New("Error No Token")

	}
	token, err := jwt.ParseWithClaims(tokenString, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("hola"), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

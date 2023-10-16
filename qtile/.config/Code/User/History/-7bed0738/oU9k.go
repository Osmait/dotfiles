package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var request loginRequest

		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

}

package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osmait/blog-go/internal/find"
	"github.com/osmait/blog-go/internal/helper"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(findUser find.FindUserService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var request loginRequest

		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := findUser.FindByEmail(ctx, request.Email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(request.Password)); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
			return
		}

		token, err := helper.SignToken(user.ID().String())
		ctx.JSON(http.StatusOK, token)

	}

}

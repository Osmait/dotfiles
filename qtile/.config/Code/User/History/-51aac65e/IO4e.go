package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osmait/blog-go/internal/creating"
)

type userRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(creatingUser creating.CreateUseService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var request userRequest

		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = creatingUser.Save(ctx, request.Email, request.Password)

		ctx.Status(http.StatusCreated)

	}
}

package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osmait/blog-go/internal/creating"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(creatingUser creating.CreateUseService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var request UserRequest

		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := creatingUser.Save(ctx, request.Email, request.Password)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		ctx.Status(http.StatusCreated)
	}

}

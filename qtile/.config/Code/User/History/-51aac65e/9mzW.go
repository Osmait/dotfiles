package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osmait/blog-go/internal/domain"
	"github.com/osmait/blog-go/internal/platfrom/storage/postgres"
)

type userRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(userRepository postgres.UserRepository) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var request userRequest

		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := domain.NewUser(request.Email, request.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}

	}
}

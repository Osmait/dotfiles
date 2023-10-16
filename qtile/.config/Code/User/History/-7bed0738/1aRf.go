package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osmait/blog-go/internal/find"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(findUser find.Find) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var request loginRequest

		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := findUser.FindUser(request.Email)

	}

}

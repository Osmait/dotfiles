package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osmait/blog-go/internal/creating"
	"github.com/osmait/blog-go/internal/helper"
)

type Post struct {
	Title  string `json:"title"`
	UserId string `json:"userId"`
}

func Createpost(creatingPost creating.CreatePostService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := helper.DecodeJwt(ctx)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
		}

		claims, ok := token.Claims.(*helper.AppClaims)

		if !ok || token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
			return

		}

		var request Post
		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = creatingPost.Save(ctx, request.Title, claims.Id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.Status(http.StatusCreated)

	}
}

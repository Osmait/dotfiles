package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osmait/blog-go/internal/creating"
)

type Post struct {
	Title  string `json:"title"`
	UserId string `json:"userId"`
}

func Createpost(creatingPost creating.CreatePostService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request Post
		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := creatingPost.Save(ctx, request.Title, request.UserId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.Status(http.StatusCreated)

	}
}

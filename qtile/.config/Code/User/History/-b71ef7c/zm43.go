package user

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/osmait/blog-go/internal/creating"
)

func TestCreateUser(t *testing.T) {
	createPost := creating.CreatePostService

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/user", CreateUser(createPost))

}

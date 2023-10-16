package user

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/osmait/blog-go/internal/creating"
	"github.com/osmait/blog-go/internal/platfrom/server/handler/post"
	"github.com/osmait/blog-go/internal/platfrom/storage/postgres"
	"github.com/stretchr/testify/assert"
)

const (
	host       = "localhost"
	port       = 3000
	dbUser     = "osmait"
	dbPassword = "admin123"
	dbName     = "my_store"
	dbHost     = "localhost"
	dbPort     = 5432
)

func TestCreateUser(t *testing.T) {
	postgresURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", postgresURI)
	assert.NoError(t, err)
	postRepository := postgres.NewPostRepository(db)
	creatingPostService := creating.NewCreatePostService(postRepository)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/user", post.Createpost(*creatingPostService))

}

package post

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/osmait/blog-go/internal/creating"

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

func TestCreatePost(t *testing.T) {

	postgresURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", postgresURI)
	assert.NoError(t, err)
	postRepository := postgres.NewPostRepository(db)
	creatingPostService := creating.NewCreatePostService(postRepository)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/posts", Createpost(*creatingPostService))
	postBody := map[string]string{"title": "Mi nuevo post", "userId": "c72dc854-e5f3-11ed-8b3b-00d8619ee79d"}

	jsonBody, _ := json.Marshal(postBody)

	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusCreated, resp.Code)

}

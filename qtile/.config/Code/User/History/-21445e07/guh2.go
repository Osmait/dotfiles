package user

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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

func TestUser(t *testing.T) {

	postgresURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", postgresURI)
	assert.NoError(t, err)
	userRepository := postgres.NewUserRepository(db)
	creatingUserService := creating.NewCreateUseService(userRepository)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/users", CreateUser(*creatingUserService))
	postBody := map[string]string{"email": "saulBurgos7@gmail", "password": "12345678"}

	jsonBody, _ := json.Marshal(postBody)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusCreated, resp.Code)

}

package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/health", CheckHandler())

	t.Run("it return 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/health", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.NotEqual(t, http.StatusOK, res.StatusCode)
	})

}

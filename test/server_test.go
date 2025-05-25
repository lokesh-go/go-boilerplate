package test

import (
	"net/http"
	"testing"

	"github.com/lokesh-go/go-api-microservice/internal/bootstrap"
	"github.com/stretchr/testify/assert"
)

// TestServer ...
func TestServer(t *testing.T) {
	// Start the server
	go func() {
		err := bootstrap.Initialize()
		if err != nil {
			assert.NoError(t, err, "app initialisation failed")
		}
	}()

	t.Run("should return public server 200 OK", func(t *testing.T) {
		res, err := http.Get("http://localhost/ping")
		assert.NoError(t, err, "public server failed")
		defer res.Body.Close()
		assert.Equal(t, http.StatusOK, res.StatusCode, "expected public server 200")
	})

	t.Run("should return internal server 200 OK", func(t *testing.T) {
		res, err := http.Get("http://localhost:8080/ping")
		assert.NoError(t, err, "internal server failed")
		defer res.Body.Close()
		assert.Equal(t, http.StatusOK, res.StatusCode, "expected internal server 200")
	})
}

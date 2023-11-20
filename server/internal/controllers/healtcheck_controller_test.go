package controllers_test

import (
	"gobooru/internal/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthCheckPing(t *testing.T) {
	healthCheckController := controllers.NewHealthCheckController()

	e := echo.New()
	req, err := http.NewRequest(
		http.MethodGet,
		"/ping",
		strings.NewReader("ping"),
	)
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	require.NoError(t, healthCheckController.Ping(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "pong", rec.Body.String())
}

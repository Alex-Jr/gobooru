package controllers_test

import (
	"gobooru/internal/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckPing(t *testing.T) {
	healthCheckController := controllers.NewHealthCheckController()

	e := echo.New()
	req, err := http.NewRequest(
		http.MethodGet,
		"/ping",
		strings.NewReader("ping"),
	)

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("The request could not be created because of: %v", err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, healthCheckController.Ping(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "pong", rec.Body.String())
	}
}

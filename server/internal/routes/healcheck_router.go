package routes

import (
	"gobooru/internal/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterHealthCheckRoutes(e *echo.Echo, c controllers.HealthCheckController) {
	e.GET("/ping", c.Ping)
}

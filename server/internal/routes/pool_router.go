package routes

import (
	"gobooru/internal/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterPoolRoutes(e *echo.Echo, poolController controllers.PoolController) {
	group := e.Group("/pool")

	group.POST("", poolController.CreatePool)
}

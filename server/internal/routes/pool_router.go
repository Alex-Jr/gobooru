package routes

import (
	"gobooru/internal/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterPoolRoutes(e *echo.Echo, poolController controllers.PoolController) {
	g := e.Group("/pool")

	g.GET("/:id", poolController.Fetch)
	g.GET("", poolController.List)
	g.POST("", poolController.Create)
}

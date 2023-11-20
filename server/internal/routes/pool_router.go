package routes

import (
	"gobooru/internal/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterPoolRoutes(e *echo.Echo, poolController controllers.PoolController) {
	g := e.Group("/pools")

	g.DELETE("/:id", poolController.Delete)
	g.GET("", poolController.List)
	g.GET("/:id", poolController.Fetch)
	g.PATCH("/:id", poolController.Update)
	g.POST("", poolController.Create)
}

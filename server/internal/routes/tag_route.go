package routes

import (
	"gobooru/internal/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterTagRoutes(e *echo.Echo, tagController controllers.TagController) {
	g := e.Group("/tags")

	g.GET("/:id", tagController.Fetch)
	g.GET("", tagController.List)
	g.DELETE("/:id", tagController.Delete)
}

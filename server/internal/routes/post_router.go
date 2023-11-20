package routes

import (
	"gobooru/internal/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterPostRoutes(e *echo.Echo, postController controllers.PostController) {
	g := e.Group("/posts")

	g.DELETE("/:id", postController.Delete)
	g.GET("/:id", postController.Fetch)
	g.POST("", postController.Create)
}

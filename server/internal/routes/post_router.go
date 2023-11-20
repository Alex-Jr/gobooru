package routes

import (
	"gobooru/internal/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterPostRoutes(e *echo.Echo, postController controllers.PostController) {
	g := e.Group("/posts")

	g.POST("", postController.Create)
}

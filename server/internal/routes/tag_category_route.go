package routes

import (
	"gobooru/internal/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterTagCategoryRoutes(e *echo.Echo, tagCategoryController controllers.TagCategoryController) {
	g := e.Group("/tag-categories")

	g.GET("", tagCategoryController.List)
}

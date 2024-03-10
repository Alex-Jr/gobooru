package routes

import (
	"gobooru/internal/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterPostRoutes(e *echo.Echo, postController controllers.PostController) {
	g := e.Group("/posts")

	g.DELETE("/:id", postController.Delete)
	g.GET("", postController.List)
	g.GET("/:id", postController.Fetch)
	g.PATCH("/:id", postController.Update)
	g.POST("", postController.Create)
	g.GET("/hash/:hash", postController.FetchByHash)
	g.POST("/:id/notes", postController.CreateNote)
}

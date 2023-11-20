package controllers

import (
	"gobooru/internal/dtos"
	"gobooru/internal/services"

	"github.com/labstack/echo/v4"
)

type PostController interface {
	Create(c echo.Context) error
}

type postController struct {
	postService services.PostService
}

type PostControllerConfig struct {
	PostService services.PostService
}

func NewPostController(c PostControllerConfig) PostController {
	return &postController{
		postService: c.PostService,
	}
}

func (c postController) Create(ctx echo.Context) error {
	dto := dtos.CreatePostDTO{}

	if err := ctx.Bind(&dto); err != nil {
		return err
	}

	response, err := c.postService.Create(
		ctx.Request().Context(),
		dto,
	)

	if err != nil {
		return err
	}

	return ctx.JSON(200, response)
}

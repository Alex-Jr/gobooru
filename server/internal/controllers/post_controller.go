package controllers

import (
	"fmt"
	"gobooru/internal/dtos"
	"gobooru/internal/services"

	"github.com/labstack/echo/v4"
)

type PostController interface {
	Create(c echo.Context) error
	Delete(c echo.Context) error
	Fetch(c echo.Context) error
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
		return fmt.Errorf("error binding dto: %w", err)
	}

	response, err := c.postService.Create(
		ctx.Request().Context(),
		dto,
	)

	if err != nil {
		return fmt.Errorf("postService.Create: %w", err)
	}

	return ctx.JSON(200, response)
}

func (c postController) Delete(ctx echo.Context) error {
	dto := dtos.DeletePostDTO{}

	if err := ctx.Bind(&dto); err != nil {
		return err
	}

	response, err := c.postService.Delete(
		ctx.Request().Context(),
		dto,
	)

	if err != nil {
		return err
	}

	return ctx.JSON(200, response)
}

func (c postController) Fetch(ctx echo.Context) error {
	dto := dtos.FetchPostDTO{}

	if err := ctx.Bind(&dto); err != nil {
		return err
	}

	response, err := c.postService.Fetch(
		ctx.Request().Context(),
		dto,
	)

	if err != nil {
		return err
	}

	return ctx.JSON(200, response)
}

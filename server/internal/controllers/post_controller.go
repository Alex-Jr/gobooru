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
	FetchByHash(c echo.Context) error
	List(c echo.Context) error
	Update(c echo.Context) error
	CreateNote(c echo.Context) error
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

func (cl postController) Create(c echo.Context) error {
	dto := dtos.CreatePostDTO{}

	if err := c.Bind(&dto); err != nil {
		return fmt.Errorf("c.Bind: %w", err)
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	dto.File = file

	response, err := cl.postService.Create(
		c.Request().Context(),
		dto,
	)

	if err != nil {
		return fmt.Errorf("postService.Create: %w", err)
	}

	return c.JSON(200, response)
}

func (cl postController) Delete(c echo.Context) error {
	dto := dtos.DeletePostDTO{}

	if err := c.Bind(&dto); err != nil {
		return fmt.Errorf("error binding dto: %w", err)
	}

	response, err := cl.postService.Delete(
		c.Request().Context(),
		dto,
	)

	if err != nil {
		return fmt.Errorf("postService.Delete: %w", err)
	}

	return c.JSON(200, response)
}

func (cl postController) Fetch(c echo.Context) error {
	dto := dtos.FetchPostDTO{}

	if err := c.Bind(&dto); err != nil {
		return fmt.Errorf("c.Bind: %w", err)
	}

	response, err := cl.postService.Fetch(
		c.Request().Context(),
		dto,
	)

	if err != nil {
		return fmt.Errorf("postService.Fetch: %w", err)
	}

	return c.JSON(200, response)
}

func (cl postController) FetchByHash(c echo.Context) error {
	dto := dtos.FetchPostByHashDTO{}

	if err := c.Bind(&dto); err != nil {
		return fmt.Errorf("c.Bind: %w", err)
	}

	response, err := cl.postService.FetchByHash(
		c.Request().Context(),
		dto,
	)

	if err != nil {
		return fmt.Errorf("postService.FetchByHash: %w", err)
	}

	return c.JSON(200, response)
}

func (cl postController) List(c echo.Context) error {
	dto := dtos.ListPostDTO{}

	if err := c.Bind(&dto); err != nil {
		return fmt.Errorf("c.Bind: %w", err)
	}

	response, err := cl.postService.List(
		c.Request().Context(),
		dto,
	)

	if err != nil {
		return fmt.Errorf("postService.List: %w", err)
	}

	return c.JSON(200, response)
}

func (cl postController) Update(c echo.Context) error {
	dto := dtos.UpdatePostDTO{}

	if err := c.Bind(&dto); err != nil {
		return fmt.Errorf("c.Bind: %w", err)
	}

	response, err := cl.postService.Update(
		c.Request().Context(),
		dto,
	)

	if err != nil {
		return fmt.Errorf("postService.Update: %w", err)
	}

	return c.JSON(200, response)
}

func (cl postController) CreateNote(c echo.Context) error {
	dto := dtos.CreatePostNoteDTO{}

	if err := c.Bind(&dto); err != nil {
		return fmt.Errorf("c.Bind: %w", err)
	}

	response, err := cl.postService.CreateNote(
		c.Request().Context(),
		dto,
	)

	if err != nil {
		return fmt.Errorf("postService.CreateNote: %w", err)
	}

	return c.JSON(200, response)
}

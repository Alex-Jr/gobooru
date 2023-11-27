package controllers

import (
	"fmt"
	"gobooru/internal/dtos"
	"gobooru/internal/services"

	"github.com/labstack/echo/v4"
)

type TagController interface {
	Fetch(c echo.Context) error
}

type tagController struct {
	tagService services.TagService
}

type TagControllerConfig struct {
	TagService services.TagService
}

func NewTagController(c TagControllerConfig) TagController {
	return &tagController{
		tagService: c.TagService,
	}
}

func (cl tagController) Fetch(c echo.Context) error {
	dto := dtos.FetchTagDTO{}

	err := c.Bind(&dto)
	if err != nil {
		return fmt.Errorf("c.Bind: %w", err)
	}

	response, err := cl.tagService.Fetch(
		c.Request().Context(),
		dto,
	)

	if err != nil {
		return fmt.Errorf("tagService.Fetch: %w", err)
	}

	return c.JSON(200, response)
}

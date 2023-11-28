package controllers

import (
	"fmt"
	"gobooru/internal/dtos"
	"gobooru/internal/services"

	"github.com/labstack/echo/v4"
)

type TagCategoryController interface {
	List(c echo.Context) error
}

type tagCategoryController struct {
	tagCategoryService services.TagCategoryService
}

type TagCategoryControllerConfig struct {
	TagCategoryService services.TagCategoryService
}

func NewTagCategoryController(c TagCategoryControllerConfig) TagCategoryController {
	return &tagCategoryController{
		tagCategoryService: c.TagCategoryService,
	}
}

func (cl tagCategoryController) List(c echo.Context) error {
	dto := dtos.ListTagCategoryDTO{}

	err := c.Bind(&dto)
	if err != nil {
		return fmt.Errorf("c.Bind: %w", err)
	}

	response, err := cl.tagCategoryService.List(
		c.Request().Context(),
		dto,
	)

	if err != nil {
		return fmt.Errorf("tagCategoryService.List: %w", err)
	}

	return c.JSON(200, response)
}

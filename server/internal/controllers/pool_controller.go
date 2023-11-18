package controllers

import (
	"context"
	"gobooru/internal/dtos"
	"gobooru/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PoolController interface {
	CreatePool(c echo.Context) error
}

type poolController struct {
	poolService services.PoolService
}

type PoolControllerConfig struct {
	PoolService services.PoolService
}

func NewPoolController(c PoolControllerConfig) PoolController {
	return &poolController{
		poolService: c.PoolService,
	}
}

func (p poolController) CreatePool(c echo.Context) error {
	var dto dtos.CreatePoolDTO

	err := c.Bind(&dto)
	if err != nil {
		return err
	}

	response, err := p.poolService.Create(context.TODO(), dto)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

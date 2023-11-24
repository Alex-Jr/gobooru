package controllers

import (
	"context"
	"fmt"
	"gobooru/internal/dtos"
	"gobooru/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PoolController interface {
	Create(c echo.Context) error
	Delete(c echo.Context) error
	Fetch(c echo.Context) error
	List(c echo.Context) error
	Update(c echo.Context) error
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

func (p poolController) Create(c echo.Context) error {
	var dto dtos.CreatePoolDTO

	err := c.Bind(&dto)
	if err != nil {
		return fmt.Errorf("c.Bind: %w", err)
	}

	response, err := p.poolService.Create(context.TODO(), dto)
	if err != nil {
		return fmt.Errorf("poolService.Create: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (p poolController) Fetch(c echo.Context) error {
	var dto dtos.FetchPoolDTO

	err := c.Bind(&dto)
	if err != nil {
		return fmt.Errorf("c.Bind: %w", err)
	}

	response, err := p.poolService.Fetch(context.TODO(), dto)
	if err != nil {
		return fmt.Errorf("poolService.Fetch: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (p poolController) Delete(c echo.Context) error {
	var dto dtos.DeletePoolDTO

	err := c.Bind(&dto)
	if err != nil {
		return fmt.Errorf("c.Bind: %w", err)
	}

	response, err := p.poolService.Delete(context.TODO(), dto)
	if err != nil {
		return fmt.Errorf("poolService.Delete: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (p poolController) List(c echo.Context) error {
	var dto dtos.ListPoolDTO

	err := c.Bind(&dto)
	if err != nil {
		return fmt.Errorf("c.Bind: %w", err)
	}

	response, err := p.poolService.List(context.TODO(), dto)
	if err != nil {
		return fmt.Errorf("poolService.List: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

func (p poolController) Update(c echo.Context) error {
	var dto dtos.UpdatePoolDTO

	err := c.Bind(&dto)
	if err != nil {
		return fmt.Errorf("c.Bind: %w", err)
	}

	response, err := p.poolService.Update(context.TODO(), dto)
	if err != nil {
		return fmt.Errorf("poolService.Update: %w", err)
	}

	return c.JSON(http.StatusOK, response)
}

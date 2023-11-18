package controllers

import "github.com/labstack/echo/v4"

type HealthCheckController interface {
	Ping(c echo.Context) error
}

type healthCheckController struct{}

func NewHealthCheckController() HealthCheckController {
	return &healthCheckController{}
}

func (h healthCheckController) Ping(c echo.Context) error {
	return c.String(200, "pong")
}

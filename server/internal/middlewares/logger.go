package middlewares

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type LoggerData struct {
	Level  string `json:"level"`
	Method string `json:"method"`
	IP     string `json:"ip"`
	URI    string `json:"uri"`
	Status int    `json:"status"`
	Date   string `json:"date"`
	Err    string `json:"err,omitempty"`
}

func NewLoggerMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			var b []byte

			if v.Error == nil {
				b, _ = json.Marshal(
					LoggerData{
						Level:  "INFO",
						Method: c.Request().Method,
						IP:     c.RealIP(),
						URI:    strings.Split(v.URI, "?")[0],
						Status: v.Status,
						Date:   v.StartTime.Format(time.RFC3339Nano),
					},
				)
			} else {
				b, _ = json.Marshal(
					LoggerData{
						Level:  "ERROR",
						Method: c.Request().Method,
						IP:     c.RealIP(),
						URI:    strings.Split(v.URI, "?")[0],
						Status: v.Status,
						Date:   v.StartTime.Format(time.RFC3339Nano),
						Err:    v.Error.Error(),
					},
				)
			}

			fmt.Print(string(b) + "\n")

			return nil
		},
	})
}

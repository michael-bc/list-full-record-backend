package http_server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthResponse struct {
	ServiceName string `json:"serviceName"`
	Version     string `json:"version"`
}

func HealthHandler(serviceName, version string) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, HealthResponse{
			ServiceName: serviceName,
			Version:     version,
		})
	}
}

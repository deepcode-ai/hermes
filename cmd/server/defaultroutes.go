package main

import (
	"net/http"

	"github.com/deepcode-ai/hermes/config"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	Version string `json:"version"`
	Status  string `json:"status"`
}

const StatusUp = "up"

func AddDefaultRoutes(cfg *config.AppConfig, e *echo.Echo) {
	e.GET("/healthz", func(c echo.Context) error {
		response := HealthHandler{
			Version: cfg.Version,
			Status:  StatusUp,
		}
		return c.JSON(http.StatusOK, response)
	})
}

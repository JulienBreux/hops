package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CheckResult struct {
	Name    string `json:"name"`
	Up      bool   `json:"up"`
	Latency int64  `json:"latency"`
}

type ServiceStatus struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Emoji       string        `json:"emoji"`
	Checks      []CheckResult `json:"checks"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Mock results for now
	e.GET("/api/health", func(c echo.Context) error {
		statuses := []ServiceStatus{
			{
				Name:        "API Gateway",
				Description: "Main entry point",
				Emoji:       "🚀",
				Checks: []CheckResult{
					{Name: "health", Up: true, Latency: 15},
				},
			},
			{
				Name:        "Auth Service",
				Description: "Authentication logic",
				Emoji:       "🔒",
				Checks: []CheckResult{
					{Name: "health", Up: true, Latency: 8},
				},
			},
		}
		return c.JSON(http.StatusOK, statuses)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

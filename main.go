package main

import (
        "log/slog"
        "net/http"
        "os"

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

        logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
        e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
                LogURI:     true,
                LogStatus:  true,
                LogMethod:  true,
                LogLatency: true,
                LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
                        logger.Info("request",
                                slog.String("method", v.Method),
                                slog.String("uri", v.URI),
                                slog.Int("status", v.Status),
                                slog.Duration("latency", v.Latency),
                        )
                        return nil
                },
        }))
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

	// Serve static files from the Svelte build directory and enable SPA fallback
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	        Root:   "web/build",
	        Index:  "index.html",
	        HTML5:  true,
	}))

	e.Logger.Fatal(e.Start(":8080"))}

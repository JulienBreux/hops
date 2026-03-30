package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/julienbreux/hops/internal/config"
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

func setupApp() *echo.Echo {
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

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		// We proceed with what we have (an empty config fallback is handled by the loader)
	}

	e.GET("/api/health", func(c echo.Context) error {
		var statuses []ServiceStatus
		if cfg != nil && cfg.Services != nil {
			for _, s := range cfg.Services {
				status := ServiceStatus{
					Name:        s.Name,
					Description: s.Description,
					Emoji:       s.Emoji,
					Checks:      []CheckResult{},
				}
				for _, chk := range s.Checks {
					status.Checks = append(status.Checks, CheckResult{
						Name:    chk.Name,
						Up:      true, // Mocking actual execution result for now
						Latency: 0,
					})
				}
				statuses = append(statuses, status)
			}
		}

		if statuses == nil {
			statuses = []ServiceStatus{}
		}

		return c.JSON(http.StatusOK, statuses)
	})

	// Serve static files from the Svelte build directory and enable SPA fallback
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "web/build",
		Index: "index.html",
		HTML5: true,
	}))

	return e
}

func main() {
	e := setupApp()
	e.Logger.Fatal(e.Start(":8080"))
}

package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

// Load reads the configuration from the file specified by the CONFIG_PATH
// environment variable. If the file is not found, it returns an empty configuration
// without failing.
func Load() (*Config, error) {
	cfg := &Config{
		Services: make([]Service, 0),
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		slog.Warn("CONFIG_PATH environment variable is empty, starting with no services")
		return cfg, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			slog.Warn("Configuration file not found, starting with no services", "path", configPath)
			return cfg, nil
		}
		// Return error for other read issues (e.g., permissions)
		slog.Error("Failed to read configuration file", "path", configPath, "error", err)
		return cfg, err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		slog.Error("Failed to parse configuration file", "path", configPath, "error", err)
		return cfg, err
	}

	slog.Info("Configuration loaded successfully", "services_count", len(cfg.Services))
	return cfg, nil
}

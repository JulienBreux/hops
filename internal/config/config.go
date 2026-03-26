package config

// Config represents the root configuration of the application.
type Config struct {
	Services []Service `yaml:"services"`
}

// Service represents a configured service to be monitored.
type Service struct {
	Name        string  `yaml:"name"`
	Description string  `yaml:"description"`
	Emoji       string  `yaml:"emoji"`
	Checks      []Check `yaml:"checks"`
}

// Check represents a specific connectivity check (HTTP or gRPC).
type Check struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`   // e.g., "http" or "grpc"
	Target string `yaml:"target"` // URL or host:port
}

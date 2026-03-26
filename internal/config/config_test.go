package config_test

import (
	"testing"

	"github.com/julienbreux/hops/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestConfigUnmarshalYAML(t *testing.T) {
	yamlData := []byte(`
services:
  - name: API Gateway
    description: Main entry point
    emoji: 🚀
    checks:
      - name: health
        type: http
        target: http://api:8080/health
  - name: Auth Service
    description: Authentication logic
    emoji: 🔒
    checks:
      - name: health
        type: grpc
        target: auth:50051
`)

	var cfg config.Config
	err := yaml.Unmarshal(yamlData, &cfg)

	require.NoError(t, err)
	require.Len(t, cfg.Services, 2)

	assert.Equal(t, "API Gateway", cfg.Services[0].Name)
	assert.Equal(t, "Main entry point", cfg.Services[0].Description)
	assert.Equal(t, "🚀", cfg.Services[0].Emoji)
	require.Len(t, cfg.Services[0].Checks, 1)
	assert.Equal(t, "health", cfg.Services[0].Checks[0].Name)
	assert.Equal(t, "http", cfg.Services[0].Checks[0].Type)
	assert.Equal(t, "http://api:8080/health", cfg.Services[0].Checks[0].Target)

	assert.Equal(t, "Auth Service", cfg.Services[1].Name)
	assert.Equal(t, "🔒", cfg.Services[1].Emoji)
	require.Len(t, cfg.Services[1].Checks, 1)
	assert.Equal(t, "grpc", cfg.Services[1].Checks[0].Type)
}

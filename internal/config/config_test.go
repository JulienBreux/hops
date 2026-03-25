package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	yamlContent := `
services:
  - name: test-service
    description: "A test service"
    emoji: "🔍"
    checks:
      - name: health
        type: http
        endpoint: http://localhost:8080/health
        method: GET
`
	tmpfile, err := os.CreateTemp("", "config.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write([]byte(yamlContent))
	assert.NoError(t, err)
	err = tmpfile.Close()
	assert.NoError(t, err)

	cfg, err := Parse(tmpfile.Name())
	assert.NoError(t, err)

	assert.Equal(t, 1, len(cfg.Services))
	assert.Equal(t, "test-service", cfg.Services[0].Name)
	assert.Equal(t, "A test service", cfg.Services[0].Description)
	assert.Equal(t, "🔍", cfg.Services[0].Emoji)
	assert.Equal(t, 1, len(cfg.Services[0].Checks))
	assert.Equal(t, "health", cfg.Services[0].Checks[0].Name)
	assert.Equal(t, "http", cfg.Services[0].Checks[0].Type)
	assert.Equal(t, "http://localhost:8080/health", cfg.Services[0].Checks[0].Endpoint)
	assert.Equal(t, "GET", cfg.Services[0].Checks[0].Method)
}

func TestParseConfigEnvFallback(t *testing.T) {
	os.Setenv("HOPS_SERVICE_NAME", "env-service")
	defer os.Unsetenv("HOPS_SERVICE_NAME")

	yamlContent := `
services:
  - name: "${HOPS_SERVICE_NAME}"
`
	tmpfile, err := os.CreateTemp("", "config.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write([]byte(yamlContent))
	assert.NoError(t, err)
	err = tmpfile.Close()
	assert.NoError(t, err)

	cfg, err := Parse(tmpfile.Name())
	assert.NoError(t, err)

	assert.Equal(t, 1, len(cfg.Services))
	assert.Equal(t, "env-service", cfg.Services[0].Name)
}

func TestParseConfigNonExistent(t *testing.T) {
	_, err := Parse("non-existent.yaml")
	assert.Error(t, err)
}

func TestParseConfigMalformed(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "malformed.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write([]byte("services: [malformed"))
	assert.NoError(t, err)
	err = tmpfile.Close()
	assert.NoError(t, err)

	_, err = Parse(tmpfile.Name())
	assert.Error(t, err)
}

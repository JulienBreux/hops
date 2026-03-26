package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/julienbreux/hops/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig_Success(t *testing.T) {
	// Setup temporary config file
	yamlData := []byte(`
services:
  - name: Test Service
    description: A test service
    emoji: 🧪
    checks:
      - name: http-check
        type: http
        target: http://localhost:8080
`)
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.yaml")
	err := os.WriteFile(configPath, yamlData, 0644)
	require.NoError(t, err)

	// Set environment variable
	t.Setenv("CONFIG_PATH", configPath)

	// Test loading
	cfg, err := config.Load()
	require.NoError(t, err)
	require.NotNil(t, cfg)
	require.Len(t, cfg.Services, 1)
	assert.Equal(t, "Test Service", cfg.Services[0].Name)
}

func TestLoadConfig_NoFileFallback(t *testing.T) {
	// Set an invalid path
	t.Setenv("CONFIG_PATH", "/path/that/does/not/exist.yaml")

	// It should NOT return an error, but an empty config
	cfg, err := config.Load()
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Empty(t, cfg.Services)
}

func TestLoadConfig_EmptyEnvVar(t *testing.T) {
	t.Setenv("CONFIG_PATH", "")

	cfg, err := config.Load()
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Empty(t, cfg.Services)
}

func TestLoadConfig_MalformedYAML(t *testing.T) {
	yamlData := []byte(`
services:
  - name: [malformed
`)
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.yaml")
	err := os.WriteFile(configPath, yamlData, 0644)
	require.NoError(t, err)

	t.Setenv("CONFIG_PATH", configPath)

	cfg, err := config.Load()
	require.Error(t, err)
	require.NotNil(t, cfg)
	assert.Empty(t, cfg.Services)
}

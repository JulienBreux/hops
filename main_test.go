package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIHealthWithConfig(t *testing.T) {
	// Create temporary configuration file
	yamlData := []byte(`
services:
  - name: Real Service
    description: A real configured service
    emoji: 🟢
    checks:
      - name: real-check
        type: http
        target: http://localhost:8080
`)
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.yaml")
	err := os.WriteFile(configPath, yamlData, 0644)
	require.NoError(t, err)

	// Set environment variable for configuration
	t.Setenv("CONFIG_PATH", configPath)

	// Initialize the app
	e := setupApp()

	// Perform request to /api/health
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Verify HTTP status
	require.Equal(t, http.StatusOK, rec.Code)

	// Verify response payload
	var statuses []ServiceStatus
	err = json.Unmarshal(rec.Body.Bytes(), &statuses)
	require.NoError(t, err)

	require.Len(t, statuses, 1)
	assert.Equal(t, "Real Service", statuses[0].Name)
	assert.Equal(t, "A real configured service", statuses[0].Description)
	assert.Equal(t, "🟢", statuses[0].Emoji)
	require.Len(t, statuses[0].Checks, 1)
	assert.Equal(t, "real-check", statuses[0].Checks[0].Name)
	// Up and Latency aren't fully implemented with real pings yet, but we check if they map.
}

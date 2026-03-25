package checker

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHTTPCheck(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	check := HTTPChecker{
		Name:     "test-http",
		URL:      ts.URL,
		Method:   "GET",
		Timeout:  1 * time.Second,
	}

	result, err := check.Check()
	assert.NoError(t, err)
	assert.True(t, result.Up)
	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.True(t, result.Latency > 0)
}

func TestHTTPCheckFailure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	check := HTTPChecker{
		Name:     "test-http-fail",
		URL:      ts.URL,
		Method:   "GET",
		Timeout:  1 * time.Second,
	}

	result, err := check.Check()
	assert.NoError(t, err)
	assert.False(t, result.Up)
	assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
}

func TestHTTPCheckTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	check := HTTPChecker{
		Name:     "test-http-timeout",
		URL:      ts.URL,
		Method:   "GET",
		Timeout:  100 * time.Millisecond,
	}

	result, err := check.Check()
	assert.Error(t, err)
	assert.False(t, result.Up)
}

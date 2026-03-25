package checker

import (
	"net/http"
	"time"
)

type Result struct {
	Up         bool
	StatusCode int
	Latency    time.Duration
}

type HTTPChecker struct {
	Name    string
	URL     string
	Method  string
	Timeout time.Duration
}

func (c *HTTPChecker) Check() (Result, error) {
	client := &http.Client{
		Timeout: c.Timeout,
	}

	req, err := http.NewRequest(c.Method, c.URL, nil)
	if err != nil {
		return Result{Up: false}, err
	}

	start := time.Now()
	resp, err := client.Do(req)
	latency := time.Since(start)

	if err != nil {
		return Result{Up: false, Latency: latency}, err
	}
	defer resp.Body.Close()

	up := resp.StatusCode >= 200 && resp.StatusCode < 300

	return Result{
		Up:         up,
		StatusCode: resp.StatusCode,
		Latency:    latency,
	}, nil
}

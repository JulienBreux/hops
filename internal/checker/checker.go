package checker

import (
	"context"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
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

type GRPCChecker struct {
	Name    string
	Address string
	Service string
	Timeout time.Duration
}

func (c *GRPCChecker) Check() (Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	start := time.Now()
	conn, err := grpc.DialContext(ctx, c.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return Result{Up: false, Latency: time.Since(start)}, err
	}
	defer conn.Close()

	client := grpc_health_v1.NewHealthClient(conn)
	resp, err := client.Check(ctx, &grpc_health_v1.HealthCheckRequest{
		Service: c.Service,
	})
	latency := time.Since(start)

	if err != nil {
		return Result{Up: false, Latency: latency}, err
	}

	up := resp.GetStatus() == grpc_health_v1.HealthCheckResponse_SERVING

	return Result{
		Up:      up,
		Latency: latency,
	}, nil
}

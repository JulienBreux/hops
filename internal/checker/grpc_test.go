package checker

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func TestGRPCCheck(t *testing.T) {
	lis, err := net.Listen("tcp", "localhost:0")
	assert.NoError(t, err)

	s := grpc.NewServer()
	hs := health.NewServer()
	hs.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(s, hs)

	go s.Serve(lis)
	defer s.Stop()

	check := GRPCChecker{
		Name:    "test-grpc",
		Address: lis.Addr().String(),
		Service: "",
		Timeout: 1 * time.Second,
	}

	result, err := check.Check()
	assert.NoError(t, err)
	assert.True(t, result.Up)
	assert.True(t, result.Latency > 0)
}

func TestGRPCCheckFailure(t *testing.T) {
	lis, err := net.Listen("tcp", "localhost:0")
	assert.NoError(t, err)

	s := grpc.NewServer()
	hs := health.NewServer()
	hs.SetServingStatus("", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
	grpc_health_v1.RegisterHealthServer(s, hs)

	go s.Serve(lis)
	defer s.Stop()

	check := GRPCChecker{
		Name:    "test-grpc-fail",
		Address: lis.Addr().String(),
		Service: "",
		Timeout: 1 * time.Second,
	}

	result, err := check.Check()
	assert.NoError(t, err)
	assert.False(t, result.Up)
}

func TestGRPCCheckTimeout(t *testing.T) {
	// No server running on this random port
	check := GRPCChecker{
		Name:    "test-grpc-timeout",
		Address: "localhost:54321",
		Service: "",
		Timeout: 100 * time.Millisecond,
	}

	result, err := check.Check()
	assert.Error(t, err)
	assert.False(t, result.Up)
}

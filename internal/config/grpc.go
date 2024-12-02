package config

import (
	"errors"
	"net"
	"os"
)

type GRPCConfig interface {
	GRPCAddress() string
}

type grpcConfig struct {
	host string
	port string
}

const grpcHost = "GRPC_HOST"
const grpcPort = "GRPC_PORT"

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHost)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}
	port := os.Getenv(grpcPort)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}
	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (c *grpcConfig) GRPCAddress() string {
	return net.JoinHostPort(c.host, c.port)
}

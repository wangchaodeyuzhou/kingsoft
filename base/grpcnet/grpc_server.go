package grpcnet

import (
	"google.golang.org/grpc"
)

func NewGrpcServer(c *GRPC) *grpc.Server {
	opts := []grpc.ServerOption{}
	opts = append(opts, grpc.ConnectionTimeout(c.Timeout))
	srv := grpc.NewServer(opts...)
	return srv
}

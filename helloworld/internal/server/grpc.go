package server

import (
	"helloworld/internal/conf"

	"google.golang.org/grpc"
)

// NewGRPCServer ...
func NewGRPCServer(c *conf.Server) *grpc.Server {
	var opts = []grpc.ServerOption{}
	if c != nil {
		if c.Timeout != nil {
			opts = append(opts, grpc.ConnectionTimeout(c.Timeout.AsDuration()))
		}
	}
	srv := grpc.NewServer(opts...)
	return srv
}

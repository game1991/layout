package server

import (
	"helloworld/internal/conf"

	"google.golang.org/grpc"
)

// NewGRPCServer ...
func NewGRPCServer(c *conf.Server) *grpc.Server {
	var opts = []grpc.ServerOption{}
	if c.GetGrpc() != nil {
		if c.Grpc.Network != "" {
			opts = append(opts)
		}
		if c.Grpc.Addr != "" {
			opts = append(opts, grpc.Address(c.Grpc.Addr))
		}
		if c.Grpc.Timeout != nil {
			opts = append(opts, grpc.ConnectionTimeout(c.Grpc.Timeout.AsDuration())
		}
	}
	srv := grpc.NewServer(opts...)
	return ssrv
}

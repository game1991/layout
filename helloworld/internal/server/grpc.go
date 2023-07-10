package server

import (
	"github.com/game1991/layout/helloworld/internal/conf"

	"google.golang.org/grpc"
)

// NewGRPCServer ...
func NewGRPCServer(c *conf.Server, unaryInts []grpc.UnaryServerInterceptor, streamInts []grpc.StreamServerInterceptor) *grpc.Server {
	var opts = []grpc.ServerOption{}
	if c != nil {
		if c.Timeout != nil {
			opts = append(opts, grpc.ConnectionTimeout(c.Timeout.AsDuration()))
		}
	}
	// 添加拦截器
	if len(unaryInts) > 0 {
		opts = append(opts, grpc.ChainUnaryInterceptor(unaryInts...))
	}
	if len(streamInts) > 0 {
		opts = append(opts, grpc.ChainStreamInterceptor(streamInts...))
	}
	srv := grpc.NewServer(opts...)
	return srv
}

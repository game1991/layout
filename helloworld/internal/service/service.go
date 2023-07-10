package service

import (
	"context"

	v1 "github.com/game1991/layout/helloworld/api/proto/v1"
)

// Service .
type Service struct {
	v1.UnimplementedHelloworldServiceServer
	gs *greeterSrv
	us *userSrv
}

// NewService .
func NewService(
	gs *greeterSrv,
	us *userSrv,
) *Service {
	return &Service{gs: gs, us: us}
}

// user service

// Login .
func (s *Service) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	return s.us.Login(ctx, req)
}

// User .
func (s *Service) User(ctx context.Context, req *v1.UserRequest) (*v1.UserInfo, error) {
	return s.us.User(ctx, req)
}

// UpdateInfo .
func (s *Service) UpdateInfo(ctx context.Context, req *v1.UpdateInfoRequest) (*v1.UpdateInfoResponse, error) {
	return s.us.UpdateInfo(ctx, req)
}

// Notify .
func (s *Service) Notify(ctx context.Context, req *v1.NotifyRequest) (*v1.NotifyResponse, error) {
	return s.us.Notify(ctx, req)
}

// greeter service

// SayHello .
func (s *Service) SayHello(ctx context.Context, req *v1.HelloworldRequset) (*v1.HelloworldReply, error) {
	return s.gs.SayHello(ctx, req)
}

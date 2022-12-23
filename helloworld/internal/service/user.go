package service

import (
	"context"
	v1 "helloworld/api/proto/v1"
	"helloworld/internal/repository"
)

// UserSrv ...
type UserSrv struct {
	v1.UnimplementedUserServiceServer
	ui repository.UserInter
}

// NewUserSrv ...
func NewUserSrv(ui repository.UserInter) *UserSrv {
	return &UserSrv{ui: ui}
}

// Login ...
func (us *UserSrv) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// TODO 检验账号密码
	return nil, nil
}

package service

import (
	"context"
	v1 "helloworld/api/proto/v1"
	"helloworld/internal/pkg/ecode"
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

// User ...
func (us *UserSrv) User(ctx context.Context, req *v1.UserRequest) (*v1.UserInfo, error) {
	// TODO 用户信息
	users, err := us.ui.FindByCondition(ctx, &repository.Condition{Name: "hello"})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, ecode.Fail(ecode.NotFound)
	}

	return convertToProtoUser(users[0]), nil
}

func convertToProtoUser(u *repository.User) (out *v1.UserInfo) {
	if u != nil {
		out = &v1.UserInfo{
			Id:       u.ID,
			UserName: u.Username,
			NickName: u.Nickname,
			Age:      u.Age,
			Gender:   v1.UserInfo_Gender(u.Gender),
		}
	}
	return
}

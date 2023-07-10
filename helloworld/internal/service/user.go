package service

import (
	"context"
	"time"

	v1 "github.com/game1991/layout/helloworld/api/proto/v1"
	"github.com/game1991/layout/helloworld/internal/conf"
	"github.com/game1991/layout/helloworld/internal/pkg/constant"
	"github.com/game1991/layout/helloworld/internal/pkg/ecode"
	"github.com/game1991/layout/helloworld/internal/repository"
	"github.com/gin-contrib/sessions"

	pContext "github.com/game1991/layout/helloworld/internal/pkg/context"
	iregexp "github.com/game1991/layout/helloworld/pkg/regexp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserSrv ...
type userSrv struct {
	ui          repository.UserInter
	sessionConf *conf.Session
}

// NewUserSrv ...
func NewUserSrv(ui repository.UserInter, sessionConf *conf.Session) *userSrv {
	return &userSrv{
		ui:          ui,
		sessionConf: sessionConf,
	}
}

// Login ...
func (us *userSrv) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// TODO 检验账号密码

	// 设置session
	// 从ctx读取gin.Context
	gCtx := pContext.GetGinCtx(ctx)
	if gCtx == nil {
		return nil, status.Errorf(codes.Internal, "ctx gin get err")
	}
	sess := sessions.DefaultMany(gCtx, us.sessionConf.GetSessionNameFromKey("user"))
	sess.Set(constant.USERKEY, &v1.UserInfo{
		Id:       1,
		UserName: "hello",
		NickName: "您好",
		Age:      18,
		Gender:   v1.UserInfo_FEMALE,
	})
	if err := sess.Save(); err != nil {
		return nil, err
	}
	return &v1.LoginResponse{
		LoginedAt: timestamppb.New(time.Now()),
	}, nil
}

// User ...
func (us *userSrv) User(ctx context.Context, req *v1.UserRequest) (*v1.UserInfo, error) {
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

// UpdateInfo .
func (us *userSrv) UpdateInfo(ctx context.Context, in *v1.UpdateInfoRequest) (*v1.UpdateInfoResponse, error) {
	if in.File == nil {
		return nil, status.Errorf(codes.InvalidArgument, "file is nil")
	}

	affeced, err := us.ui.Update(ctx, map[string]interface{}{"file": in.File})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &v1.UpdateInfoResponse{AffectedRows: affeced, IsSucceed: true}, nil
}

// Notify .
func (us *userSrv) Notify(ctx context.Context, in *v1.NotifyRequest) (*v1.NotifyResponse, error) {
	if in.Msg == nil {
		return nil, status.Errorf(codes.InvalidArgument, "msg is empty")
	}
	if in.GetEmail() == "" && in.GetPhone() == "" {
		return nil, ecode.Fail(ecode.BadRequest, "phone and email param is empty")
	}
	if in.GetPhone() != "" {
		// 校验手机号
		if !iregexp.ValidPhone(in.GetPhone()) {
			return nil, ecode.Fail(ecode.BadRequest, "phone regexp not pass")
		}
	}
	if in.GetEmail() != "" {
		// 校验邮箱
		if !iregexp.ValidEmail(in.GetEmail()) {
			return nil, ecode.Fail(ecode.BadRequest, "email regexp not pass")
		}
	}
	// TODO send msg
	return &v1.NotifyResponse{IsSend: true}, nil

}

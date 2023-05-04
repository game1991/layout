package service

import (
	"context"

	v1 "git.xq5.com/golang/helloworld/api/proto/v1"
	"git.xq5.com/golang/helloworld/internal/repository"
)

// GreeterSrv .
type greeterSrv struct {
	gi repository.GreeterInter
	ui repository.UserInter
}

// NewGreeterSrv .
func NewGreeterSrv(gi repository.GreeterInter, ui repository.UserInter) *greeterSrv {
	return &greeterSrv{gi: gi, ui: ui}
}

// SayHello .
func (gs *greeterSrv) SayHello(ctx context.Context, req *v1.HelloworldRequset) (*v1.HelloworldReply, error) {
	us, err := gs.ui.FindByCondition(ctx, &repository.Condition{Name: req.GetName()})
	if err != nil {
		return nil, err
	}

	response := &v1.HelloworldReply{}
	for _, item := range us {
		reply, err := gs.gi.Greeter(ctx, repository.CreateGreeter(item))
		if err != nil {
			return nil, err
		}
		response.Message += reply
	}
	return response, nil
}

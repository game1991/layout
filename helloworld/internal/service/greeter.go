package service

import (
	"context"
	v1 "helloworld/api/proto/v1"
	"helloworld/internal/repository"
)

// GreeterSrv .
type GreeterSrv struct {
	v1.UnimplementedGreeterServer
	gi repository.GreeterInter
	ui repository.UserInter
}

// NewGreeterSrv .
func NewGreeterSrv(gi repository.GreeterInter, ui repository.UserInter) *GreeterSrv {
	return &GreeterSrv{gi: gi, ui: ui}
}

// SayHello .
func (gs *GreeterSrv) SayHello(ctx context.Context, req *v1.HelloworldRequset) (*v1.HelloworldReply, error) {
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

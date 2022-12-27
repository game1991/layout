package controller

import (
	v1 "helloworld/api/proto/v1"
	"helloworld/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) login(ctx *gin.Context) {
	in := &v1.LoginRequest{}
	if err := ctx.Bind(in); err != nil {
		response.Fail(ctx, err)
		return
	}
	out, err := h.userSrv.Login(ctx, in)
	if err != nil {
		response.Fail(ctx, err)
		return
	}
	response.OK(ctx, out.LoginedAt.AsTime().UnixMilli())
}

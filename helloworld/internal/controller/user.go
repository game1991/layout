package controller

import (
	v1 "github.com/game1991/layout/helloworld/api/proto/v1"
	"github.com/game1991/layout/helloworld/internal/pkg/ecode"
	"github.com/game1991/layout/helloworld/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) create(ctx *gin.Context) {

}

func (h *Handler) getUser(ctx *gin.Context) {
	response.OK(ctx, "helloworld")
}

func (h *Handler) updateUserInfo(ctx *gin.Context) {

}

func (h *Handler) notify(ctx *gin.Context) {
	req := &v1.NotifyRequest{}
	if err := ctx.Bind(req); err != nil {
		response.Fail(ctx, ecode.Fail(ecode.BadRequest, err.Error()))
		return
	}
	resp, err := h.srv.Notify(ctx.Request.Context(), req)
	if err != nil {
		response.Fail(ctx, err)
		return
	}
	response.OK(ctx, resp.IsSend)
}

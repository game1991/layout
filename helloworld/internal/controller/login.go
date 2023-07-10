package controller

import (
	"context"
	"net/http"

	v1 "github.com/game1991/layout/helloworld/api/proto/v1"
	"github.com/game1991/layout/helloworld/internal/pkg/response"
	"github.com/game1991/layout/helloworld/pkg/log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) login(ctx *gin.Context) {
	in := &v1.LoginRequest{}
	if err := ctx.Bind(in); err != nil {
		response.Fail(ctx, err)
		return
	}
	out, err := h.srv.Login(ctx.MustGet("ctx").(context.Context), in)
	if err != nil {
		response.Fail(ctx, err)
		return
	}
	response.OK(ctx, out.LoginedAt.AsTime().UnixMilli())
}

func (h *Handler) logout(ctx *gin.Context) {
	log.Debugf("start to logout,clear session info")
	sessionname := h.sessionConf.GetSessionNameFromKey("user")
	sess := sessions.DefaultMany(ctx, sessionname)

	sess.Options(sessions.Options{MaxAge: -1})
	sess.Clear()
	if err := sess.Save(); err != nil {
		response.Fail(ctx, err)
		return
	}
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:   sessionname,
		MaxAge: -1,
	})
	response.OK(ctx, nil)
}

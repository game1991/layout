package controller

import (
	"context"
	"helloworld/internal/service"

	"github.com/gin-gonic/gin"
)

// Handler ...
type handler struct {
	userSrv *service.UserSrv
}

// NewHandler ...
func NewHandler(userSrv *service.UserSrv) *handler {
	return &handler{userSrv: userSrv}
}

// InstallHandler ...
func (h *handler) InstallHandler(ctx context.Context, g gin.IRouter) {
	g.POST("login", h.login)

	userG := g.Group("user")
	{
		userG.POST("", h.create)
		userG.GET("", h.getUser)
		userG.POST("updateInfo", h.updateUserInfo)
	}

}

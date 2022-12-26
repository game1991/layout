package controller

import (
	"context"
	"helloworld/internal/service"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
	userSrv *service.UserSrv
}

// NewHandler ...
func NewHandler(userSrv *service.UserSrv) *Handler {
	return &Handler{userSrv: userSrv}
}

// APIHandler ...
func (h *Handler) APIHandler(ctx context.Context, g gin.IRouter) {

	g.POST("login", h.login)

	userG := g.Group("user")
	{
		userG.POST("", h.create)
		userG.GET("", h.getUser)
		userG.POST("updateInfo", h.updateUserInfo)
	}

}

// SYSHandler ...
func (h *Handler) SYSHandler(ctx context.Context, g gin.IRouter) {
	g.GET("/swagger/", h.swaggerFile())
}

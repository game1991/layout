package controller

import (
	"context"
	"helloworld/internal/service"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
	srv *service.Service
}

// NewHandler ...
func NewHandler(srv *service.Service) *Handler {
	return &Handler{srv: srv}
}

// APIHandler ...
func (h *Handler) APIHandler(ctx context.Context, g gin.IRouter) {

	g.POST("login", h.login)

	userG := g.Group("user")
	{
		userG.POST("", h.create)
		userG.GET("", h.getUser)
		userG.POST("updateInfo", h.updateUserInfo)
		userG.POST("notify", h.notify)
	}

}

// SYSHandler ...
func (h *Handler) SYSHandler(ctx context.Context, g gin.IRouter) {
	g.GET("/swagger/*any", h.swaggerFile())
	g.GET("/swagger-ui/*any", h.swaggerUI())
}

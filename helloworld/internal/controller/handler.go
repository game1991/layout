package controller

import (
	"context"

	"github.com/game1991/layout/helloworld/internal/conf"
	"github.com/game1991/layout/helloworld/internal/middlerware"
	"github.com/game1991/layout/helloworld/internal/service"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
	srv         *service.Service
	sessionConf *conf.Session
}

// NewHandler ...
func NewHandler(srv *service.Service, sessionConf *conf.Session) *Handler {
	return &Handler{
		srv:         srv,
		sessionConf: sessionConf,
	}
}

// APIHandler ...
func (h *Handler) APIHandler(ctx context.Context, g gin.IRouter) {
	g.Use(middlerware.Context, middlerware.Session(h.sessionConf), middlerware.RequestID(), middlerware.Logg())
	g.POST("login", h.login)
	g.GET("logout", h.logout)

	loginCheck := g.Group("", middlerware.SessionAuth(h.sessionConf))
	userG := loginCheck.Group("user")
	{
		userG.POST("", h.create)
		userG.GET("", h.getUser)
		userG.POST("updateInfo", h.updateUserInfo)
		userG.POST("notify", h.notify)
	}

	module1 := loginCheck.Group("module1")
	{
		module1.Any("")
	}

	module2 := loginCheck.Group("module2")
	{
		module2.Any("")
	}
	//... ...
}

// SYSHandler ...
func (h *Handler) SYSHandler(ctx context.Context, g gin.IRouter) {
	g.GET("/swagger/*any", h.swaggerFile())
	g.GET("/swagger-ui/*any", h.swaggerUI())
}

package controller

import (
	"helloworld/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) create(ctx *gin.Context) {

}

func (h *Handler) getUser(ctx *gin.Context) {
	response.OK(ctx, "helloworld")
}

func (h *Handler) updateUserInfo(ctx *gin.Context) {

}

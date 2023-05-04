package context

import (
	"context"

	"github.com/gin-gonic/gin"
)

type GinContext struct{}

var GinCtx GinContext

// SetGinCtx 存储当前的gin.Context
func SetGinCtx(ctx context.Context, gc *gin.Context) context.Context {
	return context.WithValue(ctx, GinCtx, gc)
}

// GetGinCtx 获取context中的gin.Context
func GetGinCtx(ctx context.Context) *gin.Context {
	return ctx.Value(GinCtx).(*gin.Context)
}

package middlerware

import (
	pContext "github.com/game1991/layout/helloworld/internal/pkg/context"
	"github.com/gin-gonic/gin"
)

func Context(ctx *gin.Context) {
	warpCtx := pContext.SetGinCtx(ctx.Request.Context(), ctx)
	ctx.Set("ctx", warpCtx)
}

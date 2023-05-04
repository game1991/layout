package middlerware

import (
	pContext "git.xq5.com/golang/helloworld/internal/pkg/context"
	"github.com/gin-gonic/gin"
)

func Context(ctx *gin.Context) {
	warpCtx := pContext.SetGinCtx(ctx.Request.Context(), ctx)
	ctx.Set("ctx", warpCtx)
}

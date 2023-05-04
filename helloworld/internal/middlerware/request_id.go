package middlerware

import (
	"git.xq5.com/golang/helloworld/pkg/uuid"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestid.New(requestid.WithGenerator(func() string {
			return uuid.UUID22() // 使用短uuid方便查询
		}))
	}
}

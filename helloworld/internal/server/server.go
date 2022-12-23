package server

import (
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

// Serve in one port
func Serve(grpcServer *grpc.Server, ginServer *gin.Engine) (http.Handler, error) {
	// 监听端口并处理服务分流
	h2Handler := h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 判断协议是否为http/2 && 是grpc
		if r.ProtoMajor == 2 &&
			strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			// 按grpc方式来请求
			grpcServer.ServeHTTP(w, r)
		} else {
			// 当作普通api
			ginServer.ServeHTTP(w, r)
		}
	}), &http2.Server{})

	return h2Handler, nil
}

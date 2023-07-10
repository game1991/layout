package middlerware

import (
	"bytes"
	"io"
	"strings"

	"github.com/game1991/layout/helloworld/pkg/log"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var (
	FileReqType = "multipart/form-data"
)

func Logg() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Gin 上下文中获取 Request ID
		requestID := requestid.Get(c)
		if c.Request.Method == "GET" {
			queryParams := c.Request.URL.Query()
			log.Info("RequestRecord", "requestID", requestID, "method", c.Request.Method, "path", c.Request.URL.Path, "args", queryParams)
		} else {
			reqTypeStr := c.Request.Header.Get("Content-Type")
			reqType := strings.Split(reqTypeStr, "; ")
			if reqType[0] != FileReqType {
				jsonParams := make(map[string]interface{})
				buf := &bytes.Buffer{}
				tea := io.TeeReader(c.Request.Body, buf)
				body, err := io.ReadAll(tea)
				if err != nil {
					log.Panicf("read body err: %+v", err)
				}
				c.Request.Body = io.NopCloser(buf)
				if len(body) > 0 { // 如果有内容再进行解析
					if err := binding.JSON.BindBody(body, &jsonParams); err != nil && err != io.EOF {
						log.Panicf("json bind failed: %+v", err)
					}
					log.Info("RequestRecord", "requestID", requestID, "method", c.Request.Method, "path", c.Request.RequestURI, "args", jsonParams)
				} else {
					log.Info("RequestRecord", "requestID", requestID, "method", c.Request.Method, "path", c.Request.RequestURI)
				}
			}
		}
		c.Next()
	}
}

package controller

import (
	"net/http"
	"path"
	"strings"

	"git.xq5.com/golang/helloworld/pkg/log"

	"git.xq5.com/golang/helloworld/internal/conf"
	"git.xq5.com/golang/helloworld/pkg/swagger"
	swaggerUI "git.xq5.com/golang/helloworld/pkg/swagger-ui"

	"github.com/gin-gonic/gin"
)

/*
*
swaggerFileWithLocalPath: 提供对swagger.json文件的本地访问支持
*/
//nolint:unused
func (h *Handler) swaggerFileWithLocalPath() gin.HandlerFunc {
	// 通过配置判断是否要展示swagger
	if !conf.Bool("swagger.passed") {
		return func(c *gin.Context) {
			c.String(404, "")
		}
	}
	return gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			log.Debugf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}
		// 重写请求路径
		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		name := path.Join("api/swagger/proto", p)
		log.Infof("Serving swagger-file: %s", name)
		http.ServeFile(w, r, name)
	})
}

/*
*
swaggerFile: 提供对swagger.json文件的程序内的文件系统访问支持
*/
func (h *Handler) swaggerFile() gin.HandlerFunc {
	// 通过配置判断是否要展示swagger
	if !conf.Bool("swagger.passed") {
		return func(c *gin.Context) {
			c.String(404, "")
		}
	}
	// 使用bindata文件系统
	fileServer := http.FileServer(swagger.AssetFS())

	return gin.WrapF(
		func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasSuffix(r.URL.Path, "swagger.json") {
				log.Debugf("Not Found: %s", r.URL.Path)
				http.NotFound(w, r)
				return
			}
			// 重写请求路径
			p := strings.TrimPrefix(r.URL.Path, "/swagger/")
			name := path.Join("swagger/proto/", p)
			log.Infof("Serving swagger-file: %s", name)
			r.URL.Path = name
			r.URL.RawPath = name
			fileServer.ServeHTTP(w, r)
		})
}

/*
*
serveSwaggerUI: 提供UI支持
*/
func (h *Handler) swaggerUI() gin.HandlerFunc {
	// 通过配置判断是否要展示swagger
	if !conf.Bool("swagger.passed") {
		return func(c *gin.Context) {
			c.String(404, "")
		}
	}
	fileServer := http.FileServer(swaggerUI.AssetFS())
	// prefix := "/swagger-ui/"
	prefix := ""
	return gin.WrapH(http.StripPrefix(prefix, fileServer))
}

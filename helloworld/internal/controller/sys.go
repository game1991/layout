package controller

import (
	"helloworld/pkg/log"
	"net/http"
	"path"
	"strings"

	"helloworld/internal/conf"
	"helloworld/pkg/swagger"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
)

/*
*
swaggerFile: 提供对swagger.json文件的访问支持
*/
func (h *Handler) swaggerFile() gin.HandlerFunc {
	// 通过配置判断是否要展示swagger
	if conf.Bool("swagger.passed") == false {
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

		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		name := path.Join("api/openapi/proto/v1", p)
		log.Infof("Serving swagger-file: %s", name)
		http.ServeFile(w, r, name)
	})
}

/*
*
serveSwaggerUI: 提供UI支持
*/
func (h *Handler) swaggerUI() gin.HandlerFunc {
	// 通过配置判断是否要展示swagger
	if conf.Bool("swagger.passed") == false {
		return func(c *gin.Context) {
			c.String(404, "")
		}
	}
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	return gin.WrapH(http.StripPrefix(prefix, fileServer))
}

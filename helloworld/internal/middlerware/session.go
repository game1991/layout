package middlerware

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"strconv"

	v1 "git.xq5.com/golang/helloworld/api/proto/v1"
	"git.xq5.com/golang/helloworld/internal/conf"
	"git.xq5.com/golang/helloworld/internal/pkg/constant"
	"git.xq5.com/golang/helloworld/pkg/log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(&v1.UserInfo{})
}

// Session
func Session(c *conf.Session) gin.HandlerFunc {
	store, err := redis.NewStoreWithDB(int(c.MaxIdle), c.RedisConnectType, c.Host+":"+strconv.Itoa(int(c.Port)), c.Password, strconv.Itoa(int(c.DbNumber)), c.SessionSecret()...)
	if err != nil {
		log.Fatal("session init fail", "err info:", err)
	}
	store.Options(sessions.Options{
		Path:     c.Path,
		MaxAge:   int(c.MaxAge),
		HttpOnly: c.HttpOnly,
		SameSite: http.SameSite(c.SameSite),
	})

	return sessions.SessionsMany(c.SessionNames(), store)
}

// 主要用于鉴权和会话校验等中间件

// SessionAuth ...
func SessionAuth(c *conf.Session) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取用户的session
		sess := sessions.DefaultMany(ctx, c.GetSessionNameFromKey("user"))
		user, ok := sess.Get(constant.USERKEY).(*v1.UserInfo)
		if !ok || user == nil || user.Id == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err := sess.Save(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "sess save",
				"err": err.Error(),
			})
			return
		}
		ctx.Set(constant.USERKEY, user)
		log.Debug("登入验证,更新user信息",
			log.String("user:", fmt.Sprintf("%#v\n", user)),
		)
		ctx.Next()
	}
}

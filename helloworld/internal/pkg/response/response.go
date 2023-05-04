package response

import (
	"fmt"
	"net/http"
	"runtime"

	"git.xq5.com/golang/helloworld/internal/pkg/constant"
	"git.xq5.com/golang/helloworld/internal/pkg/ecode"

	"go.uber.org/zap"

	"git.xq5.com/golang/helloworld/pkg/log"
	"git.xq5.com/golang/helloworld/pkg/response"

	"github.com/gin-gonic/gin"
	// "go.uber.org/zap"
)

// Resp ... swagger response style
type Resp struct {
	Result bool        `json:"result" example:"true"`
	Code   int         `json:"code" example:"200"`
	Data   interface{} `json:"data"` //
}

// Pagination 分页参数
type Pagination struct {
	Total     int `json:"total" form:"total"`
	PageIndex int `json:"pageIndex" form:"pageIndex"`
	PageSize  int `json:"pageSize" form:"pageSize"`
}

// PaginationResult 分页结果
type PaginationResult struct {
	Total     int         `json:"total" form:"total"`         // 数据总数
	Pages     int         `json:"pages" form:"pages"`         // 总页数
	PageIndex int         `json:"pageIndex" form:"pageIndex"` // 当前页码
	PageSize  int         `json:"pageSize" form:"pageSize"`   // 每页数量
	List      interface{} `json:"list" form:"list"`           // 数据
}

// OK OK返回正确, data参数可选, body是否带有data字段由data参数决定
func OK(c *gin.Context, data interface{}) {

	resp := &response.Response{
		Result: true,
		Code:   ecode.OK,
		Data:   data,
	}
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

// Fail 返回错误
func Fail(c *gin.Context, err error, data ...interface{}) {
	result := &response.Response{}
	// 生成一个logger副本，便于定位controller层的错误位置
	// var traceCtx context.Context
	// v, ok := c.Get(constant.TraceCtxKey)
	// if ok {
	// 	traceCtx = v.(context.Context)
	// } else {
	// 	traceCtx = context.Background()
	// }
	// traceID := opentraceHandler.GetTraceID(traceCtx)
	// logger := log.DefaultLogger().WithFields(zap.Any(constant.REQUESTID, traceID))
	logger := log.DefaultLogger().WithOptions(zap.AddCallerSkip(0))
	if err != nil {
		// 对err进行包裹封装处理为标准response.Error
		err := ecode.WrapErrorWithCode(err)
		if err.Code() == ecode.Unauthorized || err.Code() == http.StatusUnauthorized {
			// logger.Error("用户未授权登录", constant.REQUESTID, traceID, logger.FieldErr(err))
			logger.Error("用户未授权登录", log.FieldErr(err))
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		result.Code = err.Code()

		if result.Code > 0 { // 业务code从正数开始，如果需要业务code则输出msg
			result.Message = err.Error()
		}
		logger.Error("response.Fail", log.FieldErr(err))
	}
	if len(data) > 0 {
		result.Err = data[0]
	}

	// 获取上层调用者PC，文件名，所在行
	var funName string
	var rPath string
	pc, codePath, codeLine, ok := runtime.Caller(1)
	if !ok {
		// 不ok，函数栈用尽了
		rPath = "-"
		funName = "-"
	} else {
		// 拼接文件名与所在行
		rPath = fmt.Sprintf("%s:%d", codePath, codeLine)
		// 根据PC获取函数名
		funName = runtime.FuncForPC(pc).Name()
	}

	lineKey := fmt.Sprintf("%s-%s", rPath, funName)

	c.Set(constant.LINEKEY, lineKey)
	// c.Set(constant.RESPONSEOBJ, result)

	c.AbortWithStatusJSON(http.StatusOK, result)
}

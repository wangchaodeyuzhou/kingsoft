package util

import (
	"git.kingsoft.go/intermediate/kqueue/util/code"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/exp/slog"
	"net/http"
)

// Response ...
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// APIResponse ....
func APIResponse(c *gin.Context, err *code.Err, data interface{}) {
	codeNo, message := code.DecodeErr(err)
	c.AddParam("result", gconv.String(codeNo))
	if data != nil {
		path := c.Request.URL.Path
		slog.Debug(path, "path", path, "response_body", data)
	}
	c.JSON(http.StatusOK, Response{
		Code: codeNo,
		Msg:  message,
		Data: data,
	})
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") // 请求头部
		if origin != "" {
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, "+
				"task-id,username, Business-Type")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

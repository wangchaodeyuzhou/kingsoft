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

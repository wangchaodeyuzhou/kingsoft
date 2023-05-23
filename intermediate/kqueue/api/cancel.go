package api

import (
	"git.kingsoft.go/intermediate/kqueue/manager"
	"git.kingsoft.go/intermediate/kqueue/request"
	"git.kingsoft.go/intermediate/kqueue/util"
	"git.kingsoft.go/intermediate/kqueue/util/code"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func CancelTaskToQueue(c *gin.Context) {
	req := request.CancelRequest{}
	if err := c.BindJSON(&req); err != nil {
		slog.Error("pares json fail", "err", err)
		util.APIResponse(c, code.ERROR, nil)
		return
	}

	if err := manager.Mgr.CancelTask(req.TaskId, req.WorkerId); err != nil {
		slog.Error("quit priority queue fail", "err", err)
		util.APIResponse(c, code.ERROR, nil)
		return
	}

	util.APIResponse(c, code.Success, nil)
}

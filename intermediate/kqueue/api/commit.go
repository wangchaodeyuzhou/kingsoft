package api

import (
	"git.kingsoft.go/intermediate/kqueue/manager"
	"git.kingsoft.go/intermediate/kqueue/task"
	"git.kingsoft.go/intermediate/kqueue/util"
	"git.kingsoft.go/intermediate/kqueue/util/code"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"math/rand"
	"strconv"
	"time"
)

type commitResponse struct {
	Rank     int           `json:"rank"`
	WaitTime time.Duration `json:"wait_time"`
}

func CommitTaskToQueue(c *gin.Context) {
	nextID, _ := util.NextID()
	taskId := strconv.FormatUint(nextID, 10)

	workerType := manager.Mgr.RobinNext()

	var t *task.Task
	header := c.GetHeader("Queue")
	if header == "FIFO" {
		// fifo
		t = task.NewTask(taskId, "wrc-fifo-"+taskId)
	} else {
		// priority-queue
		t = task.NewTask(taskId, "wrc-"+taskId, rand.Intn(10)+1)
	}

	// 轮询去提交到不同的业务组
	rank, waitTime, err := manager.Mgr.CommitTask(t, workerType)
	if err != nil {
		slog.Error("commit task fail", "err", err)
		util.APIResponse(c, code.ERROR, nil)
		return
	}

	cs := commitResponse{
		Rank:     rank,
		WaitTime: waitTime / time.Second,
	}

	util.APIResponse(c, code.Success, cs)
}

package manager

import (
	"fmt"
	"git.kingsoft.go/intermediate/kqueue/conf"
	"git.kingsoft.go/intermediate/kqueue/task"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"
	"sort"
	"testing"
)

func initConfig() {
	_, err := conf.LoadConfig()
	if err != nil {
		slog.Error("load configs fail", "err", err)
		return
	}
}

func TestFIFOQueue(t *testing.T) {
	initConfig()

	workerManager := NewWorkerManager()
	go workerManager.Run()
	for i := 1; i < 1000; i++ {
		curTask := task.NewTask("wrc"+gconv.String(i), "wrc-"+gconv.String(i))
		// 随机放到 worker1 或者 worker2 中去
		var workerId string
		if i < 500 {
			workerId = "worker1"
		} else {
			workerId = "worker2"
		}
		_, _, err := workerManager.CommitTask(curTask, workerId)
		assert.NoError(t, err)
	}

	err := workerManager.CancelTask("wrc99", "worker1")
	assert.NoError(t, err)

	err = workerManager.CancelTask("wrc599", "worker2")
	assert.NoError(t, err)
}

func TestPriorityQueue(t *testing.T) {
	initConfig()

	workerManager := NewWorkerManager()
	go workerManager.Run()
	for i := 1; i < 1000; i++ {
		var p int
		if i < 100 {
			p = 2
		} else if i < 200 {
			p = 1
		} else if i < 300 {
			p = 5
		} else {
			p = 3
		}
		curTask := task.NewTask("wrc"+gconv.String(i), "wrc-"+gconv.String(i), p)
		// 随机放到 worker1 或者 worker2 中去
		var workerId string
		if i < 500 {
			workerId = "worker1"
		} else {
			workerId = "worker2"
		}
		_, _, err := workerManager.CommitTask(curTask, workerId)
		assert.NoError(t, err)
	}

	err := workerManager.CancelTask("wrc990", "worker2")
	assert.NoError(t, err)

}

func TestSearchSort(t *testing.T) {
	a := []int{2, 3, 5, 7}
	p := 1
	search := sort.Search(len(a), func(i int) bool {
		return p < a[i]
	})
	fmt.Println(search)
	tail := make([]int, len(a)-search)
	copy(tail, a[search:])

	a = append(a[0:search], p)
	a = append(a, tail...)
	fmt.Println(a)
}

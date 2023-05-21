package manager

import (
	"errors"
	"fmt"
	"git.kingsoft.go/intermediate/kqueue/conf"
	"git.kingsoft.go/intermediate/kqueue/priority_queue"
	"git.kingsoft.go/intermediate/kqueue/task"
	"git.kingsoft.go/intermediate/kqueue/util"
	"git.kingsoft.go/intermediate/kqueue/worker"
	"golang.org/x/exp/slog"
	"math"
	"time"
)

type Manager struct {
	workers map[string][]*worker.Worker
}

func NewWorkerManager() *Manager {
	m := &Manager{
		workers: make(map[string][]*worker.Worker),
	}

	// load services to workers
	for workerId, workerNodes := range conf.GetConfigServices() {
		for _, node := range workerNodes {
			// todo 不同类型的队列表示
			// fifo 队列
			// q := fifo_queue.NewFIFOQueue(util.MaxQueueCapacity)

			// 优先级队列
			q := priority_queue.NewPriorityQueue(util.MaxQueueCapacity)
			m.workers[workerId] = append(m.workers[workerId], worker.NewWorker(node.Id, q))
		}
	}

	return m
}

func (m *Manager) Run() {
	slog.Debug("manager start to run")

	for _, workers := range m.workers {
		for _, w := range workers {
			w.StartWorker()
		}
	}

	slog.Debug("all worker are start")
}

// CommitTask 提交任务可进行一些平衡优化
func (m *Manager) CommitTask(task *task.Task, workerId string) (int, time.Duration, error) {
	leastIndex, err := m.findLenLeast(workerId)
	if err != nil || leastIndex == -1 {
		return -1, 0, errors.New("find queue least fail")
	}

	w := m.workers[workerId][leastIndex]
	rank, waitTime, err := w.Q.EnQueue(task)
	if err != nil {
		return -1, 0, errors.New("enter queue have fail")
	}

	slog.Info("commit task have success", "workerId", workerId, "taskId", task.TaskId, "rank", rank, "waitTime", waitTime)

	return rank, waitTime, nil
}

func (m *Manager) CancelTask(taskId, workerId string) error {
	slog.Debug("manage cancel task", "workerId", workerId, "taskId", taskId)

	workers, ok := m.workers[workerId]
	if !ok {
		slog.Error("workerId is not exist", "workerId", workerId)
		return errors.New(fmt.Sprintf("workerId %s not is exist", workerId))
	}

	for _, w := range workers {
		taskInfo, rank, _, err := w.Q.GetTaskInfo(taskId)
		if err == util.ErrTaskNotInQueue {
			continue
		} else {
			if taskInfo.Status == util.Queued && rank > 1 {
				if err = w.Q.CancelTask(taskId); err != nil {
					slog.Error("cancel task have fail", "taskId", taskId, "err", err)
					return err
				} else {
					slog.Info("manager cancel task success", "workerId", workerId, "taskId", taskId, "status",
						taskInfo.Status)
					return nil
				}
			}
		}

	}

	return nil
}

func (m *Manager) findLenLeast(workerId string) (int, error) {
	workers, ok := m.workers[workerId]
	if !ok {
		slog.Error("workerId is not exist", "workerId", workerId)
		return -1, errors.New(fmt.Sprintf("workerId %s not is exist", workerId))
	}

	minLen, index := math.MaxInt, -1
	for i, w := range workers {
		if w.GetQueueLen() < minLen {
			minLen = w.GetQueueLen()
			index = i
		}
	}

	return index, nil
}

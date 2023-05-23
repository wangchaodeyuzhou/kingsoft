package manager

import (
	"errors"
	"fmt"
	"git.kingsoft.go/intermediate/kqueue/conf"
	"git.kingsoft.go/intermediate/kqueue/fifo_queue"
	"git.kingsoft.go/intermediate/kqueue/priority_queue"
	"git.kingsoft.go/intermediate/kqueue/task"
	"git.kingsoft.go/intermediate/kqueue/util"
	"git.kingsoft.go/intermediate/kqueue/worker"
	"golang.org/x/exp/slog"
	"math"
	"time"
)

type Manager struct {
	Workers     map[string][]*worker.Worker
	index       int
	workerTypes []string
}

var Mgr *Manager

func NewWorkerManager() *Manager {
	m := &Manager{
		Workers:     make(map[string][]*worker.Worker),
		index:       0,
		workerTypes: make([]string, 0),
	}

	// load services to Workers
	for workerType, workerNodes := range conf.GetConfigServices() {
		for _, node := range workerNodes {
			// todo 不同类型的队列表示
			// fifo 队列
			q := fifo_queue.NewFIFOQueue(util.MaxQueueCapacity)

			// 优先级队列
			// q := priority_queue.NewPriorityQueue(util.MaxQueueCapacity)
			m.Workers[workerType] = append(m.Workers[workerType], worker.NewWorker(node.Id, q))
		}

		m.workerTypes = append(m.workerTypes, workerType)
	}

	return m
}

func (m *Manager) Run() {
	slog.Debug("manager start to run")

	for _, workers := range m.Workers {
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

	w := m.Workers[workerId][leastIndex]
	rank, waitTime, err := w.Q.EnQueue(task)
	if err != nil {
		return -1, 0, errors.New("enter queue have fail")
	}

	slog.Info("commit task have success", "workerId", workerId, "taskId", task.TaskId, "rank", rank, "waitTime", waitTime)

	return rank, waitTime, nil
}

func (m *Manager) CancelTask(taskId, workerId string) error {
	slog.Debug("manage cancel task", "workerId", workerId, "taskId", taskId)

	workers, ok := m.Workers[workerId]
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
					taskInfo.UpdateStatus(util.Cancel)
					slog.Info("manager cancel task success", "workerId", workerId, "taskId", taskId, "status",
						taskInfo.Status)

					// handleTasks
					switch w.Q.(type) {
					case *fifo_queue.FIFOQueue:
						q := w.Q.(*fifo_queue.FIFOQueue)
						q.HandledTasks.Cancel = append(q.HandledTasks.Cancel, taskInfo)
					case *priority_queue.PriorityQueue:
						q := w.Q.(*priority_queue.PriorityQueue)
						q.HandledTasks.Cancel = append(q.HandledTasks.Cancel, taskInfo)
					}

					return nil
				}
			}
		}

	}

	return nil
}

func (m *Manager) findLenLeast(workerId string) (int, error) {
	workers, ok := m.Workers[workerId]
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

// RobinNext 轮询
func (m *Manager) RobinNext() string {
	workerType := m.workerTypes[m.index]
	m.index = (m.index + 1) % len(m.workerTypes)
	return workerType
}

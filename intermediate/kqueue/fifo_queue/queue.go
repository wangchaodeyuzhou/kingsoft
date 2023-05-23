package fifo_queue

import (
	"errors"
	"fmt"
	"git.kingsoft.go/intermediate/kqueue/queue"
	"golang.org/x/exp/slog"
	"sync"
	"time"

	"git.kingsoft.go/intermediate/kqueue/task"
	"git.kingsoft.go/intermediate/kqueue/util"
)

// FIFOQueue 先入先出队列
type FIFOQueue struct {
	tasks      []*task.Task
	lock       sync.RWMutex
	capacity   int           // 队列容量
	noticeChan chan struct{} // 信道的事件通知方式

	// 处理过的任务
	HandledTasks *queue.HandleAllTasks
}

func NewFIFOQueue(capacity int) *FIFOQueue {
	return &FIFOQueue{
		tasks:        make([]*task.Task, 0, capacity),
		capacity:     capacity,
		noticeChan:   make(chan struct{}, capacity),
		HandledTasks: queue.NewHandleAllTasks(capacity),
	}
}

func (f *FIFOQueue) EnQueue(task *task.Task) (int, time.Duration, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	// 判断任务是否已经在队列中
	for _, t := range f.tasks {
		if t.TaskId == task.TaskId {
			return -1, 0, errors.New(fmt.Sprintf("task with ID %s already exists in queue", task.TaskId))
		}
	}

	// 判断是否已经到达队列容量上限
	if f.capacity <= len(f.tasks) {
		return -1, 0, errors.New("queue already full")
	}

	task.UpdateStatus(util.Queued)
	f.tasks = append(f.tasks, task)
	rank := len(f.tasks) + 1

	f.noticeChan <- struct{}{}

	return rank, f.GetWaitTime(rank), nil
}

func (f *FIFOQueue) DeQueue() (*task.Task, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if len(f.tasks) == 0 {
		return nil, errors.New("queue is empty")
	}

	// 取出第一个任务
	t := f.tasks[0]
	f.tasks = f.tasks[1:]
	return t, nil
}

func (f *FIFOQueue) Peek() (*task.Task, error) {
	f.lock.RLock()
	defer f.lock.RUnlock()

	if len(f.tasks) == 0 {
		return nil, errors.New("queue is empty")
	}

	return f.tasks[0], nil
}

func (f *FIFOQueue) CancelTask(taskId string) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	for i, t := range f.tasks {
		if t.TaskId == taskId {
			f.tasks = append(f.tasks[:i], f.tasks[i+1:]...)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("taskId %s not in queue", taskId))
}

// GetWaitTime 自定义需要等待的时间
func (f *FIFOQueue) GetWaitTime(rank int) time.Duration {
	return time.Duration(rank) * time.Second
}

func (f *FIFOQueue) NoticeQueue() <-chan struct{} {
	return f.noticeChan
}

func (f *FIFOQueue) GetLength() int {
	return len(f.tasks)
}

func (f *FIFOQueue) GetTaskInfo(taskId string) (*task.Task, int, time.Duration, error) {
	f.lock.RLock()
	defer f.lock.RUnlock()

	for i, t := range f.tasks {
		if t.TaskId == taskId {
			return t, i + 1, f.GetWaitTime(i + 1), nil
		}
	}

	slog.Error("task not exist", "taskId", taskId)
	return nil, -1, 0, util.ErrTaskNotInQueue
}

func (f *FIFOQueue) GetTasks() []*task.Task {
	return f.tasks
}

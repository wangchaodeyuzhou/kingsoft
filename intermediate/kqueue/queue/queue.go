package queue

import (
	"git.kingsoft.go/intermediate/kqueue/task"
	"time"
)

type Queue interface {
	// EnQueue 进入队列
	EnQueue(task *task.Task) (int, time.Duration, error)
	// DeQueue 从队列中移除第一个任务
	DeQueue() (*task.Task, error)
	// Peek 取队列第一个任务信息但是并不删除
	Peek() (*task.Task, error)
	// CancelTask 取消指定任务
	CancelTask(taskId string) error

	GetLength() int

	GetTaskInfo(taskId string) (*task.Task, int, time.Duration, error)
}

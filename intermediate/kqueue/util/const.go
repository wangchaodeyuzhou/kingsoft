package util

import "errors"

const (
	MaxQueueCapacity int = 1024
)

var (
	ErrTaskNotInQueue     = errors.New("task not found in queue")
	ErrPriorityQueueFull  = errors.New("priority queue are full")
	ErrPriorityQueueEmpty = errors.New("priority queue is empty")
)

// TaskStatus 任务状态
type TaskStatus int

const (
	Queued     TaskStatus = iota // 排队中
	Processing                   // 处理中
	Completed                    // 已完成
	Cancel                       // 取消
	Failed                       // 失败
)

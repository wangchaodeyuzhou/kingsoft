package queue

import "git.kingsoft.go/intermediate/kqueue/task"

type HandleAllTasks struct {
	Success []*task.Task
	Failed  []*task.Task
	Cancel  []*task.Task
}

func NewHandleAllTasks(capacity int) *HandleAllTasks {
	return &HandleAllTasks{
		Success: make([]*task.Task, 0, capacity),
		Failed:  make([]*task.Task, 0, capacity),
		Cancel:  make([]*task.Task, 0, capacity),
	}
}

package task

import "git.kingsoft.go/intermediate/kqueue/util"

// Task 工作单元
type Task struct {
	TaskId   string          // task 任务的唯一标识
	Status   util.TaskStatus // 任务的状态
	Payload  any             // 任务携带的信息
	Priority int             // 非负正整数 任务的优先级 越小的话优先级越高
}

func NewTask(taskId string, payload any, priority ...int) *Task {
	t := &Task{
		TaskId:  taskId,
		Payload: payload,
	}

	if len(priority) > 0 {
		t.Priority = priority[0]
	}

	return t
}

// UpdateStatus 更新任务状态
func (t *Task) UpdateStatus(status util.TaskStatus) {
	t.Status = status
}

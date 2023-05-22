package request

import (
	"errors"
	"git.kingsoft.go/intermediate/kqueue/fifo_queue"
	"git.kingsoft.go/intermediate/kqueue/manager"
	"git.kingsoft.go/intermediate/kqueue/priority_queue"
	"git.kingsoft.go/intermediate/kqueue/queue"
	"git.kingsoft.go/intermediate/kqueue/task"
	"golang.org/x/exp/slog"
)

type ManagerData struct {
	Workers map[string][]*WorkerInfo `json:"workers"`
}

type WorkerInfo struct {
	WorkerId  string `json:"worker_id"`
	QueueInfo any    `json:"queue_info"`
}

type FIFOQueue struct {
	Type  string       `json:"type"`
	Tasks []*task.Task `json:"tasks"`
}

type PriorityQueue struct {
	Type          string         `json:"type"`
	PriorityTasks [][]*task.Task `json:"priority_tasks"`
}

func convertTasks(q queue.Queue) (any, error) {
	var result any
	switch q.(type) {
	case *fifo_queue.FIFOQueue:
		tasks := q.(*fifo_queue.FIFOQueue).GetTasks()
		result = &FIFOQueue{
			Type:  "FIFO",
			Tasks: tasks,
		}
	case *priority_queue.PriorityQueue:
		priorityQueue := q.(*priority_queue.PriorityQueue)
		tmpData := make([][]*task.Task, 0, len(priorityQueue.TaskQueues))
		for _, taskQueue := range priorityQueue.TaskQueues {
			taskDatas := make([]*task.Task, 0, taskQueue.Tasks.Len())
			for i := taskQueue.Tasks.Front(); i != nil; i = i.Next() {
				taskDatas = append(taskDatas, i.Value.(*task.Task))
			}
			tmpData = append(tmpData, taskDatas)
		}
		result = &PriorityQueue{
			Type:          "Priority",
			PriorityTasks: tmpData,
		}
	default:
		return result, errors.New("queue info type is not exist")
	}

	return result, nil
}

func ConvertManagerData(m *manager.Manager) *ManagerData {
	data := make(map[string][]*WorkerInfo, len(m.Workers))
	for workerType, workers := range m.Workers {
		tmpWorkers := make([]*WorkerInfo, 0, len(workers))
		for _, worker := range workers {
			queueInfo, err := convertTasks(worker.Q)
			if err != nil {
				slog.Error("convert tasks queueInfo err", "err", err)
				continue
			}

			workerInfo := &WorkerInfo{
				WorkerId:  worker.WorkerId,
				QueueInfo: queueInfo,
			}
			tmpWorkers = append(tmpWorkers, workerInfo)
		}
		data[workerType] = tmpWorkers
	}

	return &ManagerData{Workers: data}
}

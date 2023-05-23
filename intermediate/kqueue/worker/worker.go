package worker

import (
	"git.kingsoft.go/intermediate/kqueue/fifo_queue"
	"git.kingsoft.go/intermediate/kqueue/priority_queue"
	"git.kingsoft.go/intermediate/kqueue/queue"
	"git.kingsoft.go/intermediate/kqueue/task"
	"git.kingsoft.go/intermediate/kqueue/util"
	"golang.org/x/exp/slog"
	"time"
)

// Worker 的职责是监听队列，并从队列中获取工作单元，执行工作单元的具体处理逻辑（简称 customer/worker）// 采用单队列-单消费者
type Worker struct {
	WorkerId string
	// 设计对应的 queue 模型
	Q queue.Queue
}

func NewWorker(workerId string, q queue.Queue) *Worker {
	return &Worker{
		WorkerId: workerId,
		Q:        q,
	}
}

func (w *Worker) StartWorker() {
	switch w.Q.(type) {
	case *fifo_queue.FIFOQueue:
		w.startFIFOWorker()
	case *priority_queue.PriorityQueue:
		w.startPriorityQueue()
	}
}

func (w *Worker) startFIFOWorker() {
	go func() {
		slog.Debug("start fifo queue worker begin")
		q := w.Q.(*fifo_queue.FIFOQueue)
		for {
			select {
			case <-q.NoticeQueue():
				peek, err := q.Peek()
				if err != nil || peek == nil {
					slog.Debug("peek queue fail")
					continue
				}

				// 处理任务
				w.processTask(peek)

				t, err := q.DeQueue()
				if err != nil {
					slog.Error("queue dequeue fail", "taskId", t.TaskId)
					continue
				}

				t.UpdateStatus(util.Completed)
				slog.Info("queue task have completed", "nodeId", w.WorkerId, "taskId", t.TaskId, "status", t.Status)
			}
		}
	}()
}

func (w *Worker) startPriorityQueue() {
	go func() {
		slog.Debug("start fifo queue worker begin")
		q := w.Q.(*priority_queue.PriorityQueue)
		for {
			select {
			case <-q.NoticeQueue():
				peek, err := q.DeQueue()
				if err != nil {
					// 处理失败
					q.HandledTasks.Failed = append(q.HandledTasks.Failed, peek)
					slog.Error("queue dequeue fail", "taskId", peek.TaskId)
					continue
				}

				// 处理任务
				w.processTask(peek)

				peek.UpdateStatus(util.Completed)
				slog.Info("queue task have completed", "nodeId", w.WorkerId, "taskId", peek.TaskId, "status", peek.Status)

				// 处理成功
				q.HandledTasks.Success = append(q.HandledTasks.Success, peek)
			}
		}
	}()
}

// 处理任务(消费任务)
func (w *Worker) processTask(task *task.Task) {
	task.UpdateStatus(util.Processing)
	time.Sleep(1 * time.Second) // 模拟处理过程
	slog.Info("handle you real task", "nodeId", w.WorkerId, "taskId", task.TaskId, "status", task.Status, "payload", task.Payload)
}

func (w *Worker) GetQueueLen() int {
	return w.Q.GetLength()
}

package priority_queue

import (
	"container/list"
	"git.kingsoft.go/intermediate/kqueue/task"
	"git.kingsoft.go/intermediate/kqueue/util"
	"sort"
	"sync"
	"time"
)

type PriorityQueue struct {
	TaskQueues  []*taskQueue // 优先级队列
	lock        sync.RWMutex
	capacity    int           // 优先级队列的容量
	PriorityIdx map[int]int   // 优先级索引映射 <priority,index>
	noticeChan  chan struct{} // 信道的事件通知方式
}

type taskQueue struct {
	Priority int        // 该队列的优先级
	Tasks    *list.List // 双向链表更加灵活 或者 []*task.Task
}

func NewPriorityQueue(capacity int) *PriorityQueue {
	return &PriorityQueue{
		capacity:    capacity,
		TaskQueues:  make([]*taskQueue, 0, capacity),
		PriorityIdx: make(map[int]int, capacity),
		noticeChan:  make(chan struct{}, capacity),
	}
}

func (p *PriorityQueue) EnQueue(task *task.Task) (int, time.Duration, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	idx, ok := p.PriorityIdx[task.Priority]
	if !ok {
		// 如果不存在该优先级的队列，则需要初始化一个队列，并返回该队列在切片中的索引位置
		idx = p.addTaskQueues(task.Priority)
	}

	// 队列已满
	if idx == -1 {
		return -1, 0, util.ErrPriorityQueueFull
	}

	taskQueues := p.TaskQueues[idx]
	taskQueues.Tasks.PushBack(task)

	rank := taskQueues.Tasks.Len()
	for i := 0; i < idx; i++ {
		rank += p.TaskQueues[i].Tasks.Len()
	}

	p.noticeChan <- struct{}{}

	return rank, p.GetWaitTime(rank), nil
}

func (p *PriorityQueue) addTaskQueues(priority int) int {
	n := len(p.TaskQueues)

	// 超过了多优先级队列的
	if n >= p.capacity {
		return -1
	}

	// 通过二分查找找到priority应插入的切片索引
	pos := sort.Search(n, func(i int) bool {
		return priority < p.TaskQueues[i].Priority
	})

	// 第一次
	if pos == n && pos == 0 {
		p.PriorityIdx[priority] = pos
	}

	// 更新映射表中优先级和切片索引的对应关系
	for i := pos; i < n; i++ {
		p.PriorityIdx[p.TaskQueues[i].Priority] = i
	}

	tail := make([]*taskQueue, n-pos)
	copy(tail, p.TaskQueues[pos:])

	p.TaskQueues = append(p.TaskQueues[:pos], &taskQueue{Tasks: list.New(), Priority: priority})
	p.TaskQueues = append(p.TaskQueues, tail...)

	return pos
}

// DeQueue 按照优先级队列的顺序执行任务
func (p *PriorityQueue) DeQueue() (*task.Task, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, queue := range p.TaskQueues {
		if queue.Tasks.Len() > 0 {
			taskElement := queue.Tasks.Front()
			queue.Tasks.Remove(taskElement)

			// 如何当前队列为空, 删除该映射
			if queue.Tasks.Len() == 0 {
				queue.Tasks.Init()
				delete(p.PriorityIdx, queue.Priority)
			}

			return taskElement.Value.(*task.Task), nil
		}
	}

	return nil, util.ErrPriorityQueueEmpty
}

func (p *PriorityQueue) Peek() (*task.Task, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	for _, queue := range p.TaskQueues {
		if queue.Tasks.Len() > 0 {
			taskElement := queue.Tasks.Front()

			return taskElement.Value.(*task.Task), nil
		}
	}

	return nil, util.ErrPriorityQueueEmpty
}

func (p *PriorityQueue) CancelTask(taskId string) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, queue := range p.TaskQueues {
		if queue.Tasks.Len() > 0 {
			for j := queue.Tasks.Front(); j != nil; j = j.Next() {
				t := j.Value.(*task.Task)
				if t.TaskId == taskId {
					queue.Tasks.Remove(j)
					return nil
				}
			}

		}
	}

	return util.ErrTaskNotInQueue
}

func (p *PriorityQueue) GetLength() int {
	return len(p.TaskQueues)
}

func (p *PriorityQueue) GetTaskInfo(taskId string) (*task.Task, int, time.Duration, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	rank := 0
	for _, queue := range p.TaskQueues {
		if queue.Tasks.Len() > 0 {
			for i := queue.Tasks.Front(); i != nil; i = i.Next() {
				rank++
				t := i.Value.(*task.Task)
				if t.TaskId == taskId {
					return t, rank, p.GetWaitTime(rank), nil
				}
			}
		}
	}
	return nil, -1, 0, util.ErrTaskNotInQueue
}

func (p *PriorityQueue) GetWaitTime(rank int) time.Duration {
	return time.Duration(rank) * time.Second
}

func (p *PriorityQueue) NoticeQueue() <-chan struct{} {
	return p.noticeChan
}

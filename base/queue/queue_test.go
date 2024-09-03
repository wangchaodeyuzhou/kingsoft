package queue

import (
	"fmt"
	"github.com/gammazero/deque"
	"sync"
	"testing"
	"time"
)

type ItemQueue struct {
	lock      sync.Mutex
	ops       deque.Deque[int]
	wake      chan struct{}
	done      chan struct{}
	isStopped bool
	isStarted bool
}

func NewItemQueue() *ItemQueue {
	return &ItemQueue{
		ops:  *deque.New[int](1024),
		wake: make(chan struct{}, 1),
		done: make(chan struct{}),
	}
}

func (q *ItemQueue) Start() {
	q.lock.Lock()
	if q.isStarted {
		q.lock.Unlock()
		return
	}

	q.isStarted = true
	q.lock.Unlock()
	go q.process()
}

func (q *ItemQueue) process() {
	defer close(q.done)

	for {
		<-q.wake
		for {
			q.lock.Lock()
			if q.isStopped {
				q.lock.Unlock()
				return
			}

			if q.ops.Len() == 0 {
				q.lock.Unlock()
				break
			}

			item := q.ops.PopFront()
			q.lock.Unlock()
			fmt.Println("Processing item ========》", item)
		}
	}
}

func (q *ItemQueue) Enqueue(item int) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.isStopped {
		return
	}

	q.ops.PushBack(item)
	if q.ops.Len() == 1 {
		select {
		case q.wake <- struct{}{}:
		default:
		}
	}
}

func TestQueueSync(t *testing.T) {
	queue := NewItemQueue()
	queue.Start()
	go func() {
		for i := 10; i < 20; i++ {
			queue.Enqueue(i)
		}
	}()
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
	}

	time.Sleep(time.Second * 2)
}

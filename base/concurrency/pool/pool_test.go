package pool

import (
	"fmt"
	"io"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	maxGoroutines   = 25
	pooledResources = 1
)

type dbConnection struct {
	ID int32
}

func (dbConnection *dbConnection) Close() error {
	fmt.Println("Close: connection", dbConnection.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	fmt.Println("New Connection", id)
	return &dbConnection{id}, nil
}

func TestPool_Acquire(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := New(createConnection, pooledResources)
	if err != nil {
		fmt.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	fmt.Println("ShutDown Program.")
	p.Close()
}

func performQueries(query int, p *pool) {
	conn, err := p.Acquire()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

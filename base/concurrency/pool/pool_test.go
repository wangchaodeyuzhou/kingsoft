package pool

import (
	"fmt"
	"io"
	"math/rand"
	"strconv"
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

var (
	pooll *sync.Pool
)

type Person struct {
	Name string
}

func initPool() {
	pooll = &sync.Pool{
		New: func() any {
			return new(Person)
		},
	}
}

func TestSyncPool(t *testing.T) {
	initPool()

	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		one := pooll.Get().(*Person)
		one.Name = "girl" + strconv.Itoa(i)
		pooll.Put(one)
	}

	fmt.Println(time.Since(startTime))

	fmt.Println("=========")
	startTime = time.Now()
	for i := 0; i < 10000; i++ {
		// 耗时创建这个结构体对象这里
		p := Person{Name: "Girl" + strconv.Itoa(i)}
		pooll.Put(p)
	}

	fmt.Println(time.Since(startTime))
}

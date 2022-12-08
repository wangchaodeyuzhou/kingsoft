package lock

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestMutexState(t *testing.T) {
	stateMap := make(map[int]int)

	lock := &sync.Mutex{}

	var (
		readOps   uint64
		writerOps uint64
	)

	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				lock.Lock()

				total += stateMap[key]

				lock.Unlock()

				atomic.AddUint64(&readOps, 1)

				runtime.Gosched()
			}
		}()
	}

	for w := 0; w < 100; w++ {
		go func() {
			key := rand.Intn(5)
			val := rand.Intn(100)

			lock.Lock()
			stateMap[key] = val
			lock.Unlock()

			atomic.AddUint64(&writerOps, 1)
		}()
	}

	time.Sleep(time.Second)

	readOPsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOpsFinal", readOPsFinal)

	writeOpsFinal := atomic.LoadUint64(&writerOps)
	fmt.Println("writerOpsFinal", writeOpsFinal)

	lock.Lock()
	defer lock.Unlock()
	fmt.Println("stateMap:", stateMap)
}

var wg = sync.WaitGroup{}
var sum int
var lock sync.Mutex

func Add() {

	lock.Lock()
	defer lock.Unlock()
	sum++
	wg.Done()
}

func TestSum(t *testing.T) {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go Add()
	}
	wg.Wait()
	fmt.Println("Sum: ", sum)
}

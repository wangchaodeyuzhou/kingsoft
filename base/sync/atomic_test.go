package sync

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var ops uint64

	for i := 0; i < 50000; i++ {
		go func() {
			atomic.AddUint64(&ops, 1)
			// 让其他 goroutine 执行

			runtime.Gosched()
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("ops: ", atomic.LoadUint64(&ops))
}

type person struct {
	name string
}

func TestLoadStore(t *testing.T) {
	p := &person{
		name: "Jsong",
	}
	g := atomic.Value{}
	g.Store(p)
	fmt.Println(g.Load().(*person).name)
	p1 := &person{name: "Jsong1"}
	g.Store(p1)
	fmt.Println(g.Load().(*person).name)
}

package work

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (name *namePrinter) Task() {
	fmt.Println(name.name)
	time.Sleep(time.Second)
}

func TestPool_Run(t *testing.T) {
	p := New(2)

	var wg sync.WaitGroup
	wg.Add(4 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	p.Shutdown()
}

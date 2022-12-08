package channal

import (
	"fmt"
	"testing"
)

func TestProcess(t *testing.T) {

	channels := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go Process(channels[i])
	}

	for i, ch := range channels {
		fmt.Println("i ", i, <-ch)
	}
}

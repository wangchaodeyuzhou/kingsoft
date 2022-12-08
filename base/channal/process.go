package channal

import "time"

func Process(ch chan int) {
	time.Sleep(time.Second)

	ch <- 1
}

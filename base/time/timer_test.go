package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer1 := time.NewTimer(3 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		for range timer2.C {
			fmt.Println("Timer 2 fired")
		}
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("time 2 is stop")
	}
	time.Sleep(2 * time.Second)
}

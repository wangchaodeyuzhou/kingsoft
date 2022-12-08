package channal

import (
	"fmt"
	"testing"
	"time"
)

func addNumToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}

func TestSelect(t *testing.T) {
	chan1 := make(chan int, 1)
	chan2 := make(chan int, 1)

	go addNumToChan(chan1)
	go addNumToChan(chan2)
	for {
		select {
		case e := <-chan1:
			fmt.Println("chan1 ", e)
		case e := <-chan2:
			fmt.Println("chan2 ", e)
		default:
			fmt.Println(" no data")
			time.Sleep(1 * time.Second)
		}
	}

}

func TestString(t *testing.T) {
	s := "csdcs"
	fmt.Println(s)
	for _, ss := range s {
		ss = 'f'
		fmt.Println(ss)
	}
	fmt.Println(s)
}

func TestDefer(t *testing.T) {
	println(defterFuncReturn())
}

func defterFuncReturn() (result int) {
	i := 1

	defer func() {
		result++
	}()
	return i
}

func TestGH(t *testing.T) {
	select {}
}

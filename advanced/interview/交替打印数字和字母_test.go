package interview

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母
// 效果如下 : 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

func TestChannelPrintAB(t *testing.T) {
	letter, number := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					wg.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				if i >= strings.Count(str, "") {
					i = 0
				}
				fmt.Print(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}(&wg)

	number <- true
	wg.Wait()
}

func TestChannelPrintABC(t *testing.T) {
	// Create two channels for communicating between the goroutines.
	letterChan := make(chan bool)
	numberChan := make(chan bool)

	// Create the goroutine that prints the numbers.
	go func() {
		i := 1
		for {
			// Wait for a message on the number channel.
			<-numberChan
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++

			// Send a message on the letter channel.
			letterChan <- true
		}
	}()

	// Create the goroutine that prints the letters.
	go func() {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			// Wait for a message on the letter channel.
			<-letterChan

			// Check if we have reached the end of the string.
			if i >= len(str)-1 {
				// We are done, so exit the goroutine.
				return
			}

			fmt.Print(str[i : i+1])
			i++

			// Check if we have reached the end of the string.
			if i >= len(str) {
				i = 0
			}

			fmt.Print(str[i : i+1])
			i++

			// Send a message on the number channel.
			numberChan <- true
		}
	}()

	// Start the execution by sending a message on the number channel.
	numberChan <- true
}

func TestChannelPrintABD(t *testing.T) {
	// Create a WaitGroup to track the goroutines.
	var wg sync.WaitGroup
	wg.Add(1)

	// Create two channels for communicating between the goroutines.
	letterChan := make(chan bool)
	numberChan := make(chan bool)

	// Create the goroutine that prints the numbers.
	go func() {
		i := 1
		for {
			// Wait for a message on the number channel.
			<-numberChan
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++

			// Send a message on the letter channel.
			letterChan <- true
		}
	}()

	// Create the goroutine that prints the letters.
	go func() {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			// Wait for a message on the letter channel.
			<-letterChan

			// Check if we have reached the end of the string.
			if i >= len(str)-1 {
				// We are done, so exit the goroutine.
				wg.Done()
				return
			}

			fmt.Print(str[i : i+1])
			i++

			// Check if we have reached the end of the string.
			if i >= len(str) {
				i = 0
			}

			fmt.Print(str[i : i+1])
			i++

			// Send a message on the number channel.
			numberChan <- true
		}
	}()

	// Start the execution by sending a message on the number channel.
	numberChan <- true

	// Wait for the goroutines to finish.
	wg.Wait()
}

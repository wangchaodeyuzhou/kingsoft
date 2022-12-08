package runner

import (
	"fmt"
	"os"
	"testing"
	"time"
)

const timeout = 4 * time.Second

func TestRunner_Add(t *testing.T) {
	fmt.Println("Starting worker")

	r := New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeout:
			fmt.Println("time out")
			os.Exit(1)
		case ErrInterrupt:
			fmt.Println("interrupt")
			os.Exit(2)
		}
	}

	fmt.Println("process end")
}

func createTask() func(int) {
	return func(id int) {
		fmt.Printf("Process - Task #%d.\n", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

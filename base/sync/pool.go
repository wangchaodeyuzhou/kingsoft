package sync

import (
	"fmt"
	"sync"
)

type Person struct {
	name string
}

var pool *sync.Pool

func initPool() {
	pool = &sync.Pool{
		New: func() any {
			fmt.Println("creat a new Person")
			return new(Person)
		},
	}
}

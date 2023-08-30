package sync

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	initPool()

	p := pool.Get().(*Person)
	fmt.Println("第一次重 pool 取出对象: ", p)

	p.name = "first"
	fmt.Printf("设置 p.name = %s\n", p.name)

	pool.Put(p)

	fmt.Println("pool 已经有一个对象了, 调用 Get :", pool.Get().(*Person))
	fmt.Println("pool 已经没有对象了, 调用 Get :", pool.Get().(*Person))
}

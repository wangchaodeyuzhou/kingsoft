package cgo

import (
	"fmt"
	"testing"
)

func TestZero(t *testing.T) {
	var a struct{}
	var b [0]int
	var c [1000]struct{}
	var d = make([]struct{}, 100)

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c[50])
	fmt.Printf("%p\n", &(d[50]))
}

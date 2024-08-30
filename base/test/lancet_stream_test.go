package test

import (
	"fmt"
	"github.com/duke-git/lancet/v2/stream"
	"testing"
)

func TestStreamFromSlice(t *testing.T) {
	s := stream.FromSlice([]int{1, 2, 3})

	data := s.ToSlice()

	fmt.Println(data)
}

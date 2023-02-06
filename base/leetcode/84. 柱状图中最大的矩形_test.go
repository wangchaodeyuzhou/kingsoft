package leetcode

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	// heights := []int{2, 1, 5, 6, 2, 3}
	// fmt.Println(largestRectangleAreaTimeOut(heights))

	heights := []int{6, 7, 5, 2, 4, 5, 9, 3}
	fmt.Println(largestRectangleArea(heights))
}

func largestRectangleAreaTimeOut(heights []int) int {
	ans, n := 0, len(heights)
	for i, h := range heights {
		left, right := i, i
		for left-1 >= 0 && heights[left-1] >= h {
			left--
		}

		for right+1 < n && heights[right+1] >= h {
			right++
		}

		ans = max(ans, (right-left+1)*h)
	}
	return ans
}

// 单调栈
func largestRectangleArea(heights []int) int {
	ans, n := 0, len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	var mono_stack []int
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}

	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

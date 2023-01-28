package leetcode

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	n := len(nums)
	maxIndex, end, steps := 0, 0, 0
	for i := 0; i < n-1; i++ {
		maxIndex = max(maxIndex, i+nums[i])
		if i == end {
			end = maxIndex
			steps++
		}
	}
	return steps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

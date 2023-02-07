package leetcode

import (
	"fmt"
	"testing"
)

func TestMaximalRectangle(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	maxArea := maximalRectangle(matrix)
	fmt.Println(maxArea)
}

func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	ans, m := 0, len(matrix[0])
	heights := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}
		ans = max(ans, largestAreas(heights))
	}

	return ans
}

// largestAreas leetcode 84
func largestAreas(heights []int) int {
	ans, n := 0, len(heights)
	if n == 0 {
		return ans
	}

	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}

	mono_stack := []int{}
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

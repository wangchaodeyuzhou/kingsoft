package leetcode

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	digits := []int{9, 8, 9, 9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}

	// 如果出现进位情况 就直接开始
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

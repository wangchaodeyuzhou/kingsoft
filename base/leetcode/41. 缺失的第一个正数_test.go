package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

// 请你找出其中没有出现的最小的正整数。
func TestFirstMissingPositive(t *testing.T) {
	nums := []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] <= nums[y]
	})

	minPositive := 1
	for _, v := range nums {
		if v < 1 {
			continue
		}
		if v == minPositive {
			minPositive++
		}
	}

	return minPositive
}

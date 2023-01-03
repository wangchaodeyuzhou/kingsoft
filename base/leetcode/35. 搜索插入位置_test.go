package leetcode

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums, target := []int{1, 3, 5, 8}, 7
	fmt.Println(searchInsert(nums, target))
	fmt.Println(searchTest(nums, target))
}

func searchInsert(nums []int, target int) int {
	return func([]int, int) int {
		l, r := 0, len(nums)
		for l < r {
			mid := l + (r-l)>>1
			if nums[mid] >= target {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return l
	}(nums, target)
}

func searchTest(nums []int, target int) int {
	return searchFunc(len(nums), func(i int) bool {
		return nums[i] >= target
	})
}

func searchFunc(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		h := l + (r-l)>>1
		if !f(h) {
			l = h + 1
		} else {
			r = h
		}
	}
	return l
}

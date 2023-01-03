package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestSearchRange(t *testing.T) {
	// nums, target := []int{5, 7, 7, 8, 8, 8, 8, 14}, 8
	// nums, target := []int{2, 2}, 3
	nums, target := []int{}, 0
	res := searchRange(nums, target)
	fmt.Println(res)
	fmt.Println(searchRange1(nums, target))
}

func searchRange1(nums []int, target int) []int {
	leftMost := sort.SearchInts(nums, target)
	if leftMost == len(nums) || nums[leftMost] != target {
		return []int{-1, -1}
	}

	rightMost := sort.SearchInts(nums, target+1) - 1
	return []int{leftMost, rightMost}
}

func TestSearchSlice(t *testing.T) {
	nums, target := []int{5, 7, 7, 8, 8, 8, 8, 14}, 8
	slice := SearchSlice(nums, target)
	fmt.Println(slice)
}

func SearchSlice(nums []int, target int) int {
	return search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := l + (r-l)>>1
		if !f(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func mostInt(tar int, nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] >= tar {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func searchRange(nums []int, target int) []int {
	l := mostInt(target, nums)

	if l >= len(nums) || nums[l] != target {
		return []int{-1, -1}
	}

	ll := mostInt(target+1, nums)
	return []int{l, ll - 1}
}

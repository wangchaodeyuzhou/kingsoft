package leetcode

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	// nums := []int{2, 5, 6, 0, 0, 1, 2}
	nums := []int{5, 1, 3}
	fmt.Println(search2(nums, 3))
}

func search2(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return nums[0] == target
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[l] && nums[mid] == nums[r] {
			l++
			r--
		} else if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}

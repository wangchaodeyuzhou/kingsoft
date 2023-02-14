package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}

// hash table
func longestConsecutive(nums []int) int {
	numsHash := map[int]bool{}

	// O(n)
	for _, v := range nums {
		numsHash[v] = true
	}

	countLongest := 0
	for key := range numsHash {
		if !numsHash[key-1] {
			currentNum := key
			currentCount := 1

			for numsHash[currentNum+1] {
				currentNum++
				currentCount++
			}

			if countLongest <= currentCount {
				countLongest = currentCount
			}
		}
	}

	return countLongest
}

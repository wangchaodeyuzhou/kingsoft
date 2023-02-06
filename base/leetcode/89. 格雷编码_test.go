package leetcode

import (
	"fmt"
	"testing"
)

func TestGrayCode(t *testing.T) {
	fmt.Println(grayCode(2))
}

func grayCode(n int) []int {
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = i>>1 ^ i
		fmt.Println(i>>1 ^ i)
	}
	return ans
}

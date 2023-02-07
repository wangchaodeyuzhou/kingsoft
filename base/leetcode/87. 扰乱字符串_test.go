package leetcode

import (
	"fmt"
	"testing"
)

func TestIsScramble(t *testing.T) {
	s1, s2 := "great", "rgeat"
	fmt.Println(isScramble(s1, s2))
}

func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	dp := make([][][]int8, n)
	for i := range dp {
		dp[i] = make([][]int8, n)
		for j := range dp[i] {
			dp[i][j] = make([]int8, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	// 表示第一个字符串是原始字符串从第 i1 开始， 第二个从 i2 开始 ， 长度为 length 的子串
	var dfs func(i1, i2, length int) int8
	dfs = func(i1, i2, length int) (res int8) {
		d := &dp[i1][i2][length]
		if *d != -1 {
			return *d
		}

		defer func() { *d = res }()

		// 前置条件校验
		x, y := s1[i1:length+i1], s2[i2:length+i2]
		if x == y {
			return 1
		}

		freq := [26]int{}
		for i, ch := range x {
			freq[ch-'a']++
			freq[y[i]-'a']--
		}

		for _, f := range freq {
			if f != 0 {
				return 0
			}
		}

		// 枚举分割位置
		for i := 1; i < length; i++ {
			// 不交换
			if dfs(i1, i2, i) == 1 && dfs(i1+i, i2+i, length-i) == 1 {
				return 1
			}

			// 交换
			if dfs(i1, i2+length-i, i) == 1 && dfs(i1+i, i2, length-i) == 1 {
				return 1
			}
		}
		return 0
	}
	return dfs(0, 0, n) == 1
}

func TestFreqq(t *testing.T) {
	freq := [26]int{1, 3, 54, 675, 76, 87}
	for _, v := range freq[:] {
		fmt.Println(v)
	}
}

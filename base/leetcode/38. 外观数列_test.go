package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	// n [1,30]
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	strSay := map[int]string{1: "1"}
	for i := 2; i <= n; i++ {
		preValue := strSay[i-1]
		start, pos := 0, 0
		curValue := ""
		for pos < len(preValue) {
			for pos < len(preValue) && preValue[pos] == preValue[start] {
				pos++
			}
			curValue += strconv.Itoa(pos-start) + string(preValue[start])
			start = pos
		}
		strSay[i] = curValue
	}

	return strSay[n]
}

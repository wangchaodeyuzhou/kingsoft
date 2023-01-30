package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

// 模拟
func TestAddBinary(t *testing.T) {
	fmt.Println(addBinary("101", "1"))
}

func addBinary(a string, b string) string {
	ans, carry := "", 0
	n := max(len(a), len(b))
	for i := 0; i < n; i++ {
		if i < len(a) {
			carry += int(a[len(a)-i-1] - '0')
		}

		if i < len(b) {
			carry += int(b[len(b)-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}

	if carry > 0 {
		ans = "1" + ans
	}

	return ans
}

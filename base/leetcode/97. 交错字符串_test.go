package leetcode

import (
	"fmt"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	s1, s2, s3 := "aabcc", "dbbca", "aadbbcbcac"
	fmt.Println(isInterleave(s1, s2, s3))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if n+m != t {
		return false
	}

	f := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]bool, m+1)
	}

	f[0][0] = true
	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			p := i + j - 1
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && (s1[i-1] == s3[p]))
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && (s2[j-1] == s3[p]))
			}
		}
	}
	return f[n][m]
}

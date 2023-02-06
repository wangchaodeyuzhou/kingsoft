package leetcode

import (
	"fmt"
	"testing"
)

func TestNumTrees(t *testing.T) {
	fmt.Println(numTrees(3))
}

func numTrees(n int) int {
	g := make([]int, n+1)
	g[0] = 1
	g[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}
	return g[n]
}

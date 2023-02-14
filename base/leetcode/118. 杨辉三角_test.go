package leetcode

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	g := generate(5)
	fmt.Printf("%+v", g)
}

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	for i := range ret {
		ret[i] = make([]int, i+1)
		ret[i][0] = 1
		ret[i][i] = 1
		for j := 1; j < i; j++ {
			ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
		}
	}
	return ret
}

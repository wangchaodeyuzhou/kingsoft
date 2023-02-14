package leetcode

import (
	"fmt"
	"testing"
)

func TestGetRow(t *testing.T) {
	g := getRow(3)
	fmt.Printf("%+v\n", g)
}

func getRow(rowIndex int) []int {
	ret := make([][]int, rowIndex+1)
	for i := range ret {
		ret[i] = make([]int, i+1)
		ret[i][0] = 1
		ret[i][i] = 1
		for j := 1; j < i; j++ {
			ret[i][j] = ret[i-1][j] + ret[i-1][j-1]
		}
	}

	return ret[rowIndex]
}

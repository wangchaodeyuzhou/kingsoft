package leetcode

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	rotate(matrix)
	fmt.Printf("%v", matrix)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			temp := matrix[i][j]
			// 第一个角 2,0 ==> 0, 0
			matrix[i][j] = matrix[n-j-1][i]
			// 第二个角 2,2 ==> 2, 0
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			// 第三个角 0,2 ==> 2,2
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			//	第四个角  0,0 ==> 0,2
			matrix[j][n-i-1] = temp
		}
	}
}

package leetcode

import (
	"fmt"
	"testing"
)

// 判断 9 x 9 有效数独
func TestIsValidSudoKu(t *testing.T) {

	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(isValidSudoku(board))
	fmt.Println(isValidSudoku2(board))
	fmt.Println("=============")
	board1 := [][]byte{
		{'.', '2', '1', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '6', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '7', '.', '.'},
		{'.', '.', '.', '.', '5', '.', '.', '.', '.'},
		{'.', '.', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'3', '.', '.', '.', '8', '.', '1', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '8'}}

	fmt.Println(isValidSudoku(board1))
	fmt.Println(isValidSudoku2(board1))
	fmt.Println("===========")
	board2 := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'}}
	fmt.Println(isValidSudoku(board2))
	fmt.Println(isValidSudoku2(board2))
}

func isValidSudoku(board [][]byte) bool {
	return hasRepeat(board)
}

// 是否有重复数字
func hasRepeat(board [][]byte) bool {
	// 横行检测
	for i := 0; i < len(board); i++ {
		magicNum := map[int]struct{}{}
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				num := int(board[i][j])
				if _, ok := magicNum[num]; ok {
					return false
				}
				magicNum[num] = struct{}{}
			}
		}
	}
	// 竖行检测
	for j := 0; j < len(board); j++ {
		magicNum := map[int]struct{}{}
		for i := 0; i < len(board[j]); i++ {
			if board[i][j] != '.' {
				num := int(board[i][j])
				if _, ok := magicNum[num]; ok {
					return false
				}
				magicNum[num] = struct{}{}
			}
		}
	}

	// 3 x 3 棋盘检测
	for i := 0; i < len(board); i++ {
		magicNum := map[int]struct{}{}
		for j := 0; j < len(board[i]); j++ {
			boxRow := i/3*3 + j/3
			boxRal := i%3*3 + j%3
			if board[boxRow][boxRal] != '.' {
				num := int(board[boxRow][boxRal])
				if _, ok := magicNum[num]; ok {
					return false
				}
				magicNum[num] = struct{}{}
			}
		}

	}

	return true
}

func isValidSudoku2(board [][]byte) bool {

	var (
		row [9][9]int
		ral [9][9]int
		box [3][3][9]int
	)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}

			index := board[i][j] - '1'
			row[i][index]++
			ral[j][index]++
			box[i/3][j/3][index]++

			if row[i][index] > 1 || ral[j][index] > 1 || box[i/3][j/3][index] > 1 {
				return false
			}
		}
	}

	return true
}

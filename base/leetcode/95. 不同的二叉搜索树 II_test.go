package leetcode

import (
	"fmt"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestGenerateTrees(t *testing.T) {
	for _, v := range generateTrees(3) {
		fmt.Printf("%+v\n", v)
	}
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTree := []*TreeNode{}

	for i := start; i <= end; i++ {
		leftTree := helper(start, i-1)
		rightTree := helper(i+1, end)
		for _, left := range leftTree {
			for _, right := range rightTree {
				root := &TreeNode{i, nil, nil}
				root.Left = left
				root.Right = right
				allTree = append(allTree, root)
			}
		}
	}
	return allTree
}

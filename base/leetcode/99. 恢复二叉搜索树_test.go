package leetcode

import (
	"fmt"
	"testing"
)

// 1,3,null,null,2
func TestRecoverTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Right: nil,
			Left:  &TreeNode{Val: 2},
		},
	}
	recoverTree(root)
	fmt.Println(root)
}

func recoverTree(root *TreeNode) {
	nums := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}

	inorder(root)
	x, y := FindTwoSwapped(nums)
	recoverTreeNode(root, 2, x, y)
}

func FindTwoSwapped(nums []int) (int, int) {
	index1, index2 := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			index2 = i + 1
			if index1 == -1 {
				index1 = i
			} else {
				break
			}
		}
	}
	return nums[index1], nums[index2]
}

func recoverTreeNode(node *TreeNode, count int, x, y int) {
	if node == nil {
		return
	}

	if node.Val == x || node.Val == y {
		if node.Val == x {
			node.Val = y
		} else {
			node.Val = x
		}

		count--
		if count == 0 {
			return
		}
	}

	recoverTreeNode(node.Right, count, x, y)
	recoverTreeNode(node.Left, count, x, y)
}

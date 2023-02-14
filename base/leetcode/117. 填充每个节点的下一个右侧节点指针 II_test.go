package leetcode

import (
	"fmt"
	"testing"
)

func TestConnect2(t *testing.T) {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{Val: 3,
			Right: &Node{Val: 7},
		},
	}

	fmt.Println(connect2(root))

}

// 这个不是完全二叉树
func connect2(root *Node) *Node {
	if root == nil {
		return root
	}

	for cur := root; cur != nil; {
		dummy := &Node{Val: 0}
		pre := dummy

		for cur != nil {
			if cur.Left != nil {
				pre.Next = cur.Left
				pre = pre.Next
			}

			if cur.Right != nil {
				pre.Next = cur.Right
				pre = pre.Next
			}
			cur = cur.Next
		}
		cur = dummy.Next
	}

	return root
}

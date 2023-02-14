package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	for leftMost := root; leftMost.Left != nil; leftMost = leftMost.Left {
		node := leftMost

		for node != nil {
			node.Left.Next = node.Right

			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
			node = node.Next
		}

	}

	return root
}

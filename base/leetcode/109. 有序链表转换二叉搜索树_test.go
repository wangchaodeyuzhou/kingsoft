package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

// 快慢指针
func getMedian(left, right *ListNode) *ListNode {
	fast, slow := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}

	mid := getMedian(left, right)
	root := &TreeNode{Val: mid.Val}
	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)

	return root
}

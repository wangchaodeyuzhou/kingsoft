package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cacheMap = make(map[*Node]*Node)

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if newNode, ok := cacheMap[node]; ok {
		return newNode
	}
	newNode := &Node{Val: node.Val}
	cacheMap[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList(head *Node) *Node {
	return deepCopy(head)
}

package leetcode

import "testing"

func TestCaseOne(t *testing.T) {

}

type NodeCloneGraph struct {
	Val       int
	Neighbors []*NodeCloneGraph
}

func cloneGraph(node *NodeCloneGraph) *NodeCloneGraph {
	// visited 记录当前节点是否被访问过
	visited := make(map[*NodeCloneGraph]*NodeCloneGraph)
	var called func(node *NodeCloneGraph) *NodeCloneGraph
	called = func(node *NodeCloneGraph) *NodeCloneGraph {
		// end condition
		if node == nil {
			return node
		}

		if _, ok := visited[node]; ok {
			return visited[node]
		}

		cloneNode := &NodeCloneGraph{
			Val:       node.Val,
			Neighbors: []*NodeCloneGraph{},
		}
		visited[node] = cloneNode
		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, called(n))
		}
		return cloneNode
	}
	return called(node)
}

package trie

type NodeTrie struct {
	children map[rune]*NodeTrie
	isWord   bool
}

func NewNodeTrie() *NodeTrie {
	return &NodeTrie{
		children: make(map[rune]*NodeTrie),
		isWord:   false,
	}
}

func insertWord(root *NodeTrie, word string) {
	node := root
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = NewNodeTrie()
		}
		node = node.children[char]
	}
	node.isWord = true
}

func FindWordsWithPrefix(root *NodeTrie, prefix string) []string {
	node := root
	for _, char := range prefix {
		if node.children[char] == nil {
			return nil
		}

		node = node.children[char]
	}
	return findWordsFromNode(node, prefix)
}

func findWordsFromNode(node *NodeTrie, prefix string) []string {
	words := make([]string, 0)
	if node.isWord {
		words = append(words, prefix)
	}

	for char, childNode := range node.children {
		suffixes := findWordsFromNode(childNode, prefix+string(char))
		words = append(words, suffixes...)
	}

	return words
}

func BuildTrie(words []string) *NodeTrie {
	root := NewNodeTrie()
	for _, word := range words {
		insertWord(root, word)
	}

	return root
}

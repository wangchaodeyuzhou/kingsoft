package trie

import (
	"fmt"
	"testing"
)

func TestBuildTrie(t *testing.T) {
	words := []string{"apple", "banana", "orange", "app", "orangeade", "grape", "appppp", "a", "ap"}
	root := BuildTrie(words)

	prefix := "appple"
	foundWords := FindWordsWithPrefix(root, prefix)
	fmt.Printf("Words with prefix '%s': %v\n", prefix, foundWords)
}

package binarytree

import (
	"fmt"
	"testing"
)

func TestTrieNode_Insert(t *testing.T) {
	tree :=NewTrieTree()

	tree.Insert([]byte("test"))

	fmt.Println( tree.Search([]byte("testtt")))

}
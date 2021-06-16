package binarytree

import "testing"

func complete(val interface{}, node *Node) int {
	v := val.(int)
	nodeV := val.(int)
	if v == nodeV {
		return 0
	} else if v < nodeV {
		return 1
	}
	return 2
}

func TestBinarySearchTree_Find(t *testing.T) {
	tree := NewBinarySearchTree(7, complete)
	tree.PreRootTraverse()
}

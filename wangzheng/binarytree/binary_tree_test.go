package binarytree

import "testing"

func TestTree_InRootTraverse(t *testing.T) {
	binaryTree := NewTree(33)
	binaryTree.root.left = NewNode(17)
	binaryTree.root.left.left = NewNode(13)
	binaryTree.root.left.left.right = NewNode(16)

	binaryTree.root.left.right = NewNode(18)
	binaryTree.root.left.right.right = NewNode(25)
	binaryTree.root.left.right.right.left = NewNode(19)
	binaryTree.root.left.right.right.right = NewNode(27)


	binaryTree.root.right = NewNode(50)
	binaryTree.InRootTraverse()
}

func TestTree_PreRootTraverse(t *testing.T) {
	binaryTree := NewTree(33)
	binaryTree.root.left = NewNode(17)
	binaryTree.root.left.left = NewNode(13)
	binaryTree.root.left.left.right = NewNode(16)

	binaryTree.root.left.right = NewNode(18)
	binaryTree.root.left.right.right = NewNode(25)
	binaryTree.root.left.right.right.left = NewNode(19)
	binaryTree.root.left.right.right.right = NewNode(27)


	binaryTree.root.right = NewNode(50)
	binaryTree.PreRootTraverse()
}
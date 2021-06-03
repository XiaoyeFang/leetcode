package binarytree

type Node struct {
	val   interface{}
	left  *Node
	Right *Node
}

type tree struct {
	root *Node
}

func NewTree(v interface{}) *tree {
	return &tree{&Node{val: v}}
}

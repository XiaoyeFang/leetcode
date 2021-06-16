package binarytree

import (
	"fmt"
	stack2 "github.com/XiaoyeFang/leetcode/wangzheng/stack"
)

type Node struct {
	val   interface{}
	left  *Node
	right *Node
	delete bool
}

func NewNode(v interface{}) *Node {
	return &Node{val: v}
}

type tree struct {
	root *Node
}

func NewTree(v interface{}) *tree {
	return &tree{&Node{val: v}}
}

func (t *tree) InRootTraverse() {
	var stack = stack2.NewLinkStack()
	p := t.root
	for !stack.IsEmpty() || p != nil {
		if p != nil {
			stack.Push(p)
			p = p.left
		} else {
			node := stack.Pop().(*Node)
			fmt.Println(node.val)
			p = node.right
		}
	}
}

func (t *tree) PreRootTraverse() {
	p := t.root
	s := stack2.NewLinkStack()

	for !s.IsEmpty() || nil != p {
		if nil != p {
			fmt.Printf("%+v ", p.val)
			s.Push(p)
			p = p.left
		} else {
			p = s.Pop().(*Node).right
		}
	}

	fmt.Println()
}

package stack

import "fmt"

type LinkNode struct {
	Value interface{}
	next  *LinkNode
}

type LinkStack struct {
	Top  int
	Node *LinkNode
}

func NewLinkStack() *LinkStack {

	return &LinkStack{
		Top:  0,
		Node: &LinkNode{},
	}
}

func (t *LinkStack) IsEmpty() bool {
	if t.Node == nil {
		return true
	}
	return false
}

func (t *LinkStack) Push(val interface{}) {
	if t.Top == 0 {
		t.Node = &LinkNode{Value: val, next: nil}
	} else {
		next := t.Node
		t.Node = &LinkNode{Value: val, next: next}
	}
	t.Top++
}

func (t *LinkStack) Pop() interface{} {
	if t.Top == 0 {
		return nil
	}
	top := t.Node.Value
	next := t.Node.next
	t.Node = next
	t.Top--
	return top
}

//func (t *LinkStack)Top()  interface{}{
//
//}

func (t *LinkStack) Flush() {
	t.Top = 0
	t.Node = &LinkNode{}
}

func (t *LinkStack) Print() {
	next := t.Node
	for next != nil {
		fmt.Println(next.Value)
		next = next.next
	}
}

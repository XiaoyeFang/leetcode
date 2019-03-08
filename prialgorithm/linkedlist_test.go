package prialgorithm

import "testing"

func TestReverseListNode(t *testing.T) {
	a := &ListNode{Val: 1}
	a.Next = &ListNode{Val: 2}
	//a.Next.Next=&ListNode{Val:3}
	//a.Next.Next.Next=&ListNode{Val:2}
	//a.Next.Next.Next.Next=&ListNode{Val:1}

	b := FindMiddleNode(a)
	t.Log(b)
}

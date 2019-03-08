package prialgorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

//请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。
//你就说这个题多缺德吧，都没有golang的语言选项
func deleteNode(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	fast := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	fast.Next = fast.Next.Next

	return head

}

/*给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
两种办法
一：扫描得到链表长度，计算出正数位置，在扫描删除节点
二：长短指针，长指针永远比短指针快n ，同时扫描当长到达队末时 删除短指针所指节点
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	fast, low := head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		low = low.Next
	}
	low.Next = low.Next.Next

	return head
}

//反转一个单链表。
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head

	mid := head.Next

	aft := head.Next.Next

	for {
		if mid == nil {
			break
		}
		aft = mid.Next
		mid.Next = pre
		pre = mid
		mid = aft

	}
	head.Next = nil
	return pre
}

//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res, ret := &ListNode{}, &ListNode{}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	p := l1
	q := l2

	if p.Val < q.Val {
		res = p
		p = p.Next
	} else {
		res = q
		q = q.Next
	}
	ret = res

	for {

		if p == nil || q == nil {
			break
		}

		if p.Val < q.Val {
			res.Next = p
			res = p
			p = p.Next
		} else {
			res.Next = q
			res = q
			q = q.Next
		}
	}
	for {
		if p == nil {
			break
		}

		res.Next = p
		res = p
		p = p.Next
	}

	for {
		if q == nil {
			break
		}
		res.Next = q
		res = q
		q = q.Next
	}
	return ret

}

//请判断一个链表是否为回文链表。

/*

先去找中结点,快慢指针  快指针的速度是慢指针的两倍

反转后半段

进行比较

*/
func IsPalindrome(head *ListNode) bool {

	if head == nil {
		return true
	}

	mid := FindMiddleNode(head)

	prev := ReverseListNode(mid)

	q := head

	for prev != nil {
		if prev.Val != q.Val {
			return false
		}
		prev = prev.Next
		q = q.Next
	}

	return true
}

//原地反转
func ReverseListNode(head *ListNode) *ListNode {
	prev := &ListNode{}

	for head != nil {
		temp := head.Next
		head.Next, prev = prev, head
		head = temp
	}

	return prev
}

//寻找链表中点
func FindMiddleNode(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	quick, slow := head, head

	for slow != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next

	}
	if quick != nil {
		slow = slow.Next
	}
	return slow
}

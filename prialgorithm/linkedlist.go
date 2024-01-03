package prialgorithm

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。
// 你就说这个题多缺德吧，都没有golang的语言选项
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

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
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

// 反转一个单链表。
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//pre := head
	//
	//mid := head.Next
	//
	//aft := head.Next.Next
	//
	//for {
	//	if mid == nil {
	//		break
	//	}
	//	aft = mid.Next
	//	mid.Next = pre
	//	pre = mid
	//	mid = aft
	//
	//}
	//head.Next = nil
	//头空新链   保奶 改奶 改头 后移
	var pre = &ListNode{}
	var curr = head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

/*
头插法来实现反转区域内节点

	1 2 3 4 5

从头开始遍历， 需要三个指针来标示位置
cur 标示待反转区的第一个元素 会变
next 永远指向cur的下一个元素 会随着cur变
pre 永远指向待反转区的前一个节点，也就是初始反转区域的前驱节点 始终不变
口诀
保奶
当奶奶奶
奶奶先奶
先奶奶
*/
func reverseBetweenTwo(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	dummy := &ListNode{Val: -1, Next: head}
	pre := dummy

	//找到第一个反转节点
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	//开始反转
	cur := pre.Next
	for i := left; i < right; i++ {
		//		next := curr.Next
		//		curr.Next = pre
		//		pre = curr
		//		curr = next
		next := cur.Next     // 保住next
		cur.Next = next.Next // 改变next
		dummy.print()
		fmt.Println("======================")
		next.Next = pre.Next
		dummy.print()
		fmt.Println("======================")
		pre.Next = next
		dummy.print()
		fmt.Println("======================")

	}

	return dummy.Next
}
func (head *ListNode) print() {
	pre := head
	for pre != nil {
		fmt.Print(" ", pre)
		pre = pre.Next
	}
}

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	// 第 1 步：从虚拟头节点走 left - 1 步，来到 left 节点的前一个节点
	// 建议写在 for 循环里，语义清晰
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 第 2 步：从 pre 再走 right - left + 1 步，来到 right 节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 第 3 步：切断出一个子链表（截取链表）
	leftNode := pre.Next
	curr := rightNode.Next

	// 注意：切断链接
	pre.Next = nil
	rightNode.Next = nil

	// 第 4 步：同第 206 题，反转链表的子区间
	reverseList(leftNode)

	// 第 5 步：接回到原来的链表中
	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next

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

// 原地反转
func ReverseListNode(head *ListNode) *ListNode {
	prev := &ListNode{}

	for head != nil {
		temp := head.Next
		head.Next, prev = prev, head
		head = temp
	}

	return prev
}

// 寻找链表中点
func FindMiddleNode(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	quick, slow := head, head

	for slow != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next

	}
	return slow
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	low, fast := head, head
	for low.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		low = low.Next
		fast = fast.Next.Next

		if low == fast {
			return true
		}

	}

	return false
}

func splitFileIndex(path string) string {
	paths := strings.Split(path, "/")
	if len(paths) < 7 {
		return ""
	}
	return strings.Join(paths[:6], "/") + "/file"
}

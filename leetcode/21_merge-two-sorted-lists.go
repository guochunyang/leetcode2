package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func _mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := (*ListNode)(nil) //
	tail := (*ListNode)(nil) // 结果链表的 尾结点
	for l1 != nil && l2 != nil {
		t := &ListNode{}
		if l1.Val < l2.Val {
			t.Val = l1.Val
			l1 = l1.Next
		} else if l1.Val == l2.Val {
			t.Val = l1.Val
			l1 = l1.Next
		} else {
			t.Val = l2.Val
			l2 = l2.Next
		}
		if head == nil {
			head = t
			tail = t
		} else {
			tail.Next = t
			tail = tail.Next
		}
	}

	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return head
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := (*ListNode)(nil)
	tail := (*ListNode)(nil)

	t := (*ListNode)(nil)
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			t = l1
			l1 = l1.Next
		} else if l1.Val == l2.Val {
			t = l1
			l1 = l1.Next
		} else {
			t = l2
			l2 = l2.Next
		}
		if head == nil {
			head = t
			tail = t
			head.Next = tail
		} else {
			tail.Next = t
			tail = t
		}
	}

	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return head
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
	}
}

func main() {
}

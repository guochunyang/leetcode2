package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	tail := head
	t := (*ListNode)(nil)

	// 从 head 开始算的 所以要 n-1
	for i := 0; i < n-1; i++ {
		tail = tail.Next
	}

	for tail.Next != nil {
		tail = tail.Next
		if t == nil {
			t = head
		} else {
			t = t.Next
		}
	}
	// 此时 tail 为 尾结点  t 为要删除的结点前面那个
	if t == nil {
		// 要删除的是 head
		head = head.Next
	} else {
		t.Next = t.Next.Next
	}

	return head
}

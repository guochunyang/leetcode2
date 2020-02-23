package main

func swapPairs(head *ListNode) *ListNode {

	// 保证有前面两个结点
	if head != nil && head.Next != nil {
		// 交换这俩结点
		t := head.Next // 这里是保存第二个结点
		head.Next = t.Next
		t.Next = head
		head = t
	} else {
		return head
	}

	p := head.Next // 第二个结点

	for p.Next != nil && p.Next.Next != nil {
		// swap p->next  & p->next->next
		first := p.Next
		second := p.Next.Next

		// 交换两个结点
		first.Next = second.Next
		second.Next = first

		p.Next = second
		p = first
	}

	return head
}

func main() {

}

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var result *ListNode = nil
	var current *ListNode = nil

	hasCarry := false

	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val
		if hasCarry {
			sum += 1
			hasCarry = false
		}

		if sum >= 10 {
			hasCarry = true
			sum = sum % 10
		}

		t := &ListNode{Val: sum, Next: nil}
		if result == nil {
			result = t
		}

		if current == nil {
			current = t
		} else {
			current.Next = t
			current = current.Next
		}

		// forward
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		if hasCarry {
			for l1 != nil {
				if hasCarry {
					val := l1.Val + 1
					hasCarry = false

					if val >= 10 {
						val = 0
						hasCarry = true
					}
					t := &ListNode{Val: val, Next: nil}
					current.Next = t
					current = current.Next
				} else {
					current.Next = l1
					current = l1
				}
				l1 = l1.Next
			}
		} else {
			// 直接把 l1 放在 current 的后面
			current.Next = l1
		}
	}

	if l2 != nil {
		if hasCarry {
			for l2 != nil {
				if hasCarry {
					val := l2.Val + 1
					hasCarry = false

					if val >= 10 {
						val = 0
						hasCarry = true
					}
					t := &ListNode{Val: val, Next: nil}
					current.Next = t
					current = current.Next
				} else {
					current.Next = l2
					current = l2
				}
				l2 = l2.Next
			}
		} else {
			current.Next = l2
		}
	}

	if hasCarry {
		current.Next = &ListNode{Val: 1, Next: nil}
	}

	return result
}

func main() {

}

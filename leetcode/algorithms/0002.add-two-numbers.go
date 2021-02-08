package algorithms

// time:O(n), space:O(n)
//执行用时: 8 ms 内存消耗: 4.9 MB
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, last *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		var total = carry
		if l1 != nil {
			total += l1.Val
		}
		if l2 != nil {
			total += l2.Val
		}

		node := &ListNode{total % 10, nil}
		carry = total / 10
		if last == nil {
			head = node
			last = node
		} else {
			last.Next = node
			last = last.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 && last != nil {
		last.Next = &ListNode{carry, nil}
	}

	return head
}

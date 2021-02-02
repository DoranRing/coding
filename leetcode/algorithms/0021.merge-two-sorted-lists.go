package algorithms

// 执行用时: 4 ms 内存消耗: 2.5 MB
func mergeTwoListsByLoop(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, cur *ListNode
	for l1 != nil || l2 != nil {
		var min *ListNode
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				min = l1
				l1 = l1.Next
			} else {
				min = l2
				l2 = l2.Next
			}

		} else if l1 != nil {
			min = l1
			l1 = l1.Next
		} else if l2 != nil {
			min = l2
			l2 = l2.Next
		}

		if cur == nil {
			head = min
			cur = min
		} else {
			cur.Next = min
			cur = cur.Next
		}
	}

	return head
}

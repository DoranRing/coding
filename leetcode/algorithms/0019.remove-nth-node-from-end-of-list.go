package algorithms

// 执行用时: 0 ms 内存消耗: 2.2 MB
func removeNthFromEndByRecursion(head *ListNode, n int) *ListNode {
	num := removeNthFromEndOperate(head, n)
	if num == n {
		return head.Next
	}
	return head
}

func removeNthFromEndOperate(node *ListNode, n int) int {
	if node.Next == nil {
		return 1
	}
	num := removeNthFromEndOperate(node.Next, n)
	if num == n {
		node.Next = node.Next.Next
	}
	return num + 1
}

//执行用时: 0 ms 内存消耗: 2.2 MB
func removeNthFromEndByTwoPoint(head *ListNode, n int) *ListNode {
	var slow *ListNode
	fast := head
	// 被删除节点的前置节点
	n++
	for fast != nil {
		fast = fast.Next
		n--
		if n <= 0 {
			if slow == nil {
				slow = head
			} else {
				slow = slow.Next
			}
		}
	}

	if slow == nil {
		return head.Next
	}
	slow.Next = slow.Next.Next
	return head
}

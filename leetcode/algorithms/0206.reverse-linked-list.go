package algorithms

// 执行用时: 0 ms 内存消耗: 3.1 MB
func reverseListByRecursion(head *ListNode) *ListNode {
	head, _ = reverseOperate(head)
	return head
}

func reverseOperate(node *ListNode) (*ListNode, *ListNode) {
	if node == nil || node.Next == nil {
		return node, node
	}
	head, pre := reverseOperate(node.Next)
	pre.Next = node
	node.Next = nil
	return head, node
}

// 执行用时: 4 ms 内存消耗: 2.5 MB
func reverseListByLoop(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	return newHead
}

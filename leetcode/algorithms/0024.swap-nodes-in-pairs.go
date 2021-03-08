package algorithms

//time:O(n), space:O(1)
// 执行用时: 0 ms 内存消耗: 2.1 MB
func swapPairs(head *ListNode) *ListNode {
	var resHead, resLast *ListNode = nil, nil
	for head != nil {
		var cur *ListNode
		if head.Next == nil {
			// last don't swap
			cur = head
			head = head.Next
		} else {
			// swap
			next, nextNext := head.Next, head.Next.Next
			head.Next = nextNext
			next.Next = head
			cur = next
			head = head.Next
		}

		if resLast == nil {
			// init
			resHead = cur
			resLast = cur
		} else {
			resLast.Next = cur
			resLast = resLast.Next
		}
		if resLast.Next != nil {
			resLast = resLast.Next
		}
	}

	return resHead
}

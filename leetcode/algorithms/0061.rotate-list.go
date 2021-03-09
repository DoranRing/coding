package algorithms

// time:O(n+k),space:O(1)
// 执行用时: 0 ms 内存消耗: 2.6 MB
func rotateRightByLoop(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	// 确认大小
	cur, size := head, 0
	var last *ListNode
	for cur != nil {
		size++
		if cur.Next == nil {
			last = cur
		}
		cur = cur.Next
	}
	k %= size
	if k == 0 {
		return head
	}

	// 确认转换点
	pre := head
	for i := 1; i < size-k; i++ {
		pre = pre.Next
	}
	res := pre.Next
	pre.Next = nil
	if last != nil {
		last.Next = head
	}
	return res
}

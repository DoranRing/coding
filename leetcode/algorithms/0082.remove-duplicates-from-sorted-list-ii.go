package algorithms

// time:O(n), space:O(1)
//执行用时: 0 ms 内存消耗: 3 MB
func deleteDuplicatesByTwoPoint(head *ListNode) *ListNode {
	cur := head
	var resHead, resLast, same *ListNode
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			same = cur
		} else if same == nil || same.Val != cur.Val {
			if resLast == nil {
				resHead = cur
				resLast = cur
			} else {
				resLast.Next.Val = cur.Val
				resLast = resLast.Next
			}
			same = nil
		}
		cur = cur.Next
	}

	if resLast != nil {
		resLast.Next = nil
	}
	return resHead
}

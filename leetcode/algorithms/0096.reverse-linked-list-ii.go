package algorithms

// time:O(right), space:O(1)
//执行用时: 0 ms 内存消耗: 2.1 MB
func reverseBetweenByHeadFirst(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	pre := &ListNode{Next: head}
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	var rangeHead *ListNode
	rangeLast, cur := pre.Next, pre.Next
	for i := 0; i <= right-left; i++ {
		next := cur.Next
		cur.Next = rangeHead
		rangeHead = cur
		cur = next
	}

	// link range last
	rangeLast.Next = cur
	pre.Next = rangeHead
	if left == 1 {
		return rangeHead
	}
	return head
}

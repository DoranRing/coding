package algorithms

// time:O(n), space:O(1)
// 执行用时: 0 ms 内存消耗: 2.4 MB
func partitionByLoop(head *ListNode, x int) *ListNode {
	var lessHead, lessLast, greatHead, greatLast *ListNode
	cur := head
	for cur != nil {
		if cur.Val < x {
			if lessLast == nil {
				lessHead = cur
				lessLast = cur
			} else {
				lessLast.Next = cur
				lessLast = lessLast.Next
			}
		} else {
			if greatLast == nil {
				greatHead = cur
				greatLast = cur
			} else {
				greatLast.Next = cur
				greatLast = greatLast.Next
			}
		}
		cur = cur.Next
	}

	if lessLast != nil {
		lessLast.Next = greatHead
	}
	if greatLast != nil {
		greatLast.Next = nil
	}

	if lessHead != nil {
		return lessHead
	}
	return greatHead
}

package algorithms

// time:O(n),space:O(1)
// 执行用时: 0 ms 内存消耗: 1.9 MB
func middleNodeByTwoPoint(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		}
	}
	return slow
}

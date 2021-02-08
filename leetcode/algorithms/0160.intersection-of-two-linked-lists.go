package algorithms

// 使用哈希表记录一个链表节点，然后遍历另外一个链表
// time:O(n),space:O(n)
// 执行用时: 36 ms 内存消耗: 7.7 MB
func getIntersectionNodeByHash(headA, headB *ListNode) *ListNode {
	hash := make(map[*ListNode]bool)
	for headA != nil {
		hash[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := hash[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// a + b = b + a, 如果两个链表有相交则依次遍历会有相遇的情况
// time:O(n),space:O(1)
// 执行用时: 32 ms 内存消耗: 7.2 MB
func getIntersectionNodeByTwoPoint(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	curAReset, curBReset := false, false
	for curA != nil && curB != nil {
		if curA == curB {
			return curA
		}

		if curA.Next == nil && !curAReset {
			curA = headB
			curAReset = true
		} else {
			curA = curA.Next
		}
		if curB.Next == nil && !curBReset {
			curB = headA
			curBReset = true
		} else {
			curB = curB.Next
		}
	}

	return nil
}

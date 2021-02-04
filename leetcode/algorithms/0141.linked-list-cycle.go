package algorithms

// 通过哈希表记录每个节点，遍历后续节点时检查哈希表其他节点的存在情况
// time:O(n), space:O(n)
//执行用时: 12 ms 内存消耗: 6.2 MB
func hasCycleByHash(head *ListNode) bool {
	hash := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		}
		hash[head] = true
		head = head.Next
	}

	return false
}

// 使用快慢两个指针遍历链表，如果快慢指针碰撞则表示链表是回环的，如果快指针没有后继节点则不是回环的
// time:O(n), space:O(1)
//执行用时: 8 ms 内存消耗: 4.3 MB
func hasCycleByTwoPoint(head *ListNode) bool {
	var fast, slow *ListNode
	for head != nil {
		// fast+1
		if fast == nil {
			fast = head
		} else {
			fast = fast.Next
		}

		// fast+1,slow+1
		if fast != nil && fast.Next != nil {
			fast = fast.Next
			if slow == nil {
				slow = head
			} else {
				slow = slow.Next
			}
		} else {
			// next is nil
			return false
		}

		// equal
		if slow == fast {
			return true
		}
	}

	return false
}

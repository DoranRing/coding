package algorithms

// 采用分治思想，把链表集合拆分成两条合并，再把合并后的结果两条一组进行合并，最终合并成一条结果
// time:O(n), space:O(1)
// 执行用时: 8 ms 内存消耗: 5.3 MB
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeKListsOperate(lists, 0, len(lists)-1)
}

func mergeKListsOperate(lists []*ListNode, l, r int) *ListNode {
	// 单个链表
	if l == r {
		return lists[l]
	}
	// 两个链表
	if l+1 == r {
		return mergeKListsForTwo(lists[l], lists[r])
	}
	// 分治
	mid := (l + r) / 2
	left := mergeKListsOperate(lists, l, mid)
	right := mergeKListsOperate(lists, mid+1, r)
	return mergeKListsForTwo(left, right)
}

func mergeKListsForTwo(list1, list2 *ListNode) *ListNode {
	var head, last *ListNode
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				head, last = mergeKListsForSet(head, last, list1)
				list1 = list1.Next
			} else {
				head, last = mergeKListsForSet(head, last, list2)
				list2 = list2.Next
			}
		} else if list1 != nil {
			head, last = mergeKListsForSet(head, last, list1)
			list1 = list1.Next
		} else if list2 != nil {
			head, last = mergeKListsForSet(head, last, list2)
			list2 = list2.Next
		}
	}

	return head
}

func mergeKListsForSet(head, last, node *ListNode) (*ListNode, *ListNode) {
	if last == nil {
		return node, node
	} else {
		last.Next = node
		return head, last.Next
	}
}

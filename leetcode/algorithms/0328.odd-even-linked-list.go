package algorithms

// 将奇数和偶数节点合并到一起
// 执行用时: 4 ms 内存消耗: 3.2 MB
func oddEvenListByMerge(head *ListNode) *ListNode {
	var oddHead, odd, evenHead, even *ListNode
	count := 1
	for head != nil {
		if count%2 != 0 {
			if odd == nil {
				oddHead = head
				odd = head
			} else {
				odd.Next = head
				odd = odd.Next
			}
		}
		if count%2 == 0 {
			if even == nil {
				evenHead = head
				even = head
			} else {
				even.Next = head
				even = even.Next
			}
		}
		count++
		head = head.Next
	}

	if odd != nil {
		odd.Next = evenHead
	}
	if even != nil {
		even.Next = nil
	}
	return oddHead
}

package algorithms

// 将链式表转换成顺序线性表
// time: O(n), space:O(n)
// 执行用时: 12 ms 内存消耗: 6.9 MB
func isPalindromeByCopy(head *ListNode) bool {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

// 通过递归反向遍历单链表
// time:O(n), space:O(1)
//执行用时: 24 ms 内存消耗: 11.8 MB
func isPalindromeByRecursion(head *ListNode) bool {
	if head == nil {
		return true
	}
	_, res := isPalindromeByRecursionOperate(head, head)
	return res
}

func isPalindromeByRecursionOperate(head, cur *ListNode) (*ListNode, bool) {
	if cur.Next == nil {
		return head.Next, cur.Val == head.Val
	}

	left, nextValid := isPalindromeByRecursionOperate(head, cur.Next)
	return left.Next, nextValid && cur.Val == left.Val
}

// 通过翻转单链表的后半部分，然后一次比对
// time:O(n), space(1)
// 执行用时: 20 ms 内存消耗: 5.5 MB
func isPalindromeByReverse(head *ListNode) bool {
	// 快慢指针划分两部分
	slow, count := isPalindromeScan(head)
	// 翻转后半部分
	reverse := isPalindromeReverse(slow)
	// 比对
	for i := 0; i < count; i++ {
		if head.Val != reverse.Val {
			return false
		}
		head = head.Next
		reverse = reverse.Next
	}

	return true
}

func isPalindromeScan(head *ListNode) (*ListNode, int) {
	// 快慢指针划分两部分
	var slow, fast *ListNode
	count := 0
	for true {
		// fast+1
		if fast == nil {
			fast = head
		} else {
			fast = fast.Next
		}
		// slow+1
		if slow == nil {
			slow = head
		} else {
			slow = slow.Next
		}

		if fast != nil && fast.Next != nil {
			fast = fast.Next
		} else {
			break
		}
		count++
	}
	return slow, count
}

func isPalindromeReverse(head *ListNode) *ListNode {
	newHead := head
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	return newHead
}

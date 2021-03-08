package algorithms

//执行用时: 0 ms 内存消耗: 2 MB
func getDecimalValue(head *ListNode) int {
	sum := 0
	for head != nil {
		sum = sum << 1
		sum += head.Val
		head = head.Next
	}
	return sum
}

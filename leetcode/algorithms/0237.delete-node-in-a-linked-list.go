package algorithms

//执行用时: 4 ms 内存消耗: 2.8 MB
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

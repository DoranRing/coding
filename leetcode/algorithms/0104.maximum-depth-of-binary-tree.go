package algorithms

// 通过中序遍历
// time:O(n), space:O(1)
//执行用时: 4 ms 内存消耗: 4.5 MB
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

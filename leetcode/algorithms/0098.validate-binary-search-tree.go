package algorithms

import "math"

// 通过中序遍历的方式遍历树，一次比较遍历过程中节点的值
// time:O(n), space:O(1)
//执行用时: 8 ms 内存消耗: 5.5 MB
var pre = 0

func isValidBST(root *TreeNode) bool {
	// leetcode need
	pre = math.MinInt64
	return isValidBSTOperate(root)
}

func isValidBSTOperate(node *TreeNode) bool {
	if node == nil {
		return true
	}
	res := isValidBSTOperate(node.Left) &&
		pre < node.Val &&
		isValidBSTOperate(node.Right)

	pre = node.Val
	return res
}

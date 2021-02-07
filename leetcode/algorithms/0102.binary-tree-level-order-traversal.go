package algorithms

// 使用队列记录每层节点，然后将每层节点的子节点入队列
// time:O(n), space:O(1)
//执行用时: 0 ms 内存消耗: 2.8 MB
func levelOrderByQueue(root *TreeNode) [][]int {
	res := make([][]int, 0, 0)
	queue := make([]*TreeNode, 0, 0)
	var first *TreeNode
	if root != nil {
		queue = append(queue, root)
		first = root
	}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top == first {
			res = append(res, []int{})
			first = nil
		}
		last := res[len(res)-1]
		last = append(last, top.Val)
		res[len(res)-1] = last

		if top.Left != nil {
			queue = append(queue, top.Left)
			if first == nil {
				first = top.Left
			}
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
			if first == nil {
				first = top.Right
			}
		}
	}

	return res
}

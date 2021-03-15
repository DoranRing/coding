package algorithms

// time:O(4*n), space:O(4*n)
// 执行用时: 0 ms 内存消耗: 3.5 MB
func copyRandomListByHash(head *Node) *Node {
	pos := make(map[*Node]int)
	for cur, idx := head, 0; cur != nil; idx, cur = idx+1, cur.Next {
		pos[cur] = idx
	}

	randomIdxes := make([]int, len(pos))
	for cur, idx := head, 0; cur != nil; idx, cur = idx+1, cur.Next {
		if cur.Random == nil {
			randomIdxes[idx] = -1
		} else {
			randomIdxes[idx] = pos[cur.Random]
		}
	}

	resIdxes := make([]*Node, len(pos))
	var resHead, resLast *Node
	for cur, idx := head, 0; cur != nil; idx, cur = idx+1, cur.Next {
		newNode := &Node{Val: cur.Val}
		if resLast == nil {
			resHead = newNode
			resLast = newNode
		} else {
			resLast.Next = newNode
			resLast = resLast.Next
		}
		resIdxes[idx] = newNode
	}

	for idx := 0; idx < len(pos); idx++ {
		randomIdx := randomIdxes[idx]
		if randomIdx != -1 {
			resIdxes[idx].Random = resIdxes[randomIdx]
		}
	}

	return resHead
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

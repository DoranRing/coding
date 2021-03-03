package algorithms

import "container/heap"

//time:O(n), space:O(k)
//执行用时: 40 ms 内存消耗: 6.8 MB
func kthSmallestByHeap(matrix [][]int, k int) int {
	intHeap := &MaxIntHeap{}
	heap.Init(intHeap)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			val := matrix[i][j]
			if intHeap.Len() == k {
				if intHeap.Top() < val {
					continue
				}
			}
			heap.Push(intHeap, val)
			if intHeap.Len() > k {
				heap.Pop(intHeap)
			}
		}
	}

	return heap.Pop(intHeap).(int)
}

package algorithms

import "container/heap"

//time:O(2N*logN),space:O(n)
// 执行用时: 0 ms 内存消耗: 2 MB
func lastStoneWeightByHeap(stones []int) int {
	intHeap := &MaxIntHeap{}
	heap.Init(intHeap)
	for _, stone := range stones {
		heap.Push(intHeap, stone)
	}

	for intHeap.Len() > 1 {
		x := heap.Pop(intHeap).(int)
		y := heap.Pop(intHeap).(int)
		heap.Push(intHeap, x-y)
	}
	return heap.Pop(intHeap).(int)
}

package algorithms

import "container/heap"

// time:O(n),space:O(n)
// 执行用时: 4 ms 内存消耗: 4.7 MB
func findKthLargestByHeap(nums []int, k int) int {
	intHeap := &MaxIntHeap{}
	heap.Init(intHeap)
	for _, num := range nums {
		heap.Push(intHeap, num)
	}

	for i := 1; i < k; i++ {
		heap.Pop(intHeap)
	}
	return heap.Pop(intHeap).(int)
}

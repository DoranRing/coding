package algorithms

import "container/heap"

// time:O(N*logN),space:O(n)
// 执行用时: 4 ms 内存消耗: 4.7 MB
func findKthLargestByMaxHeap(nums []int, k int) int {
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

// time:O(2N*logN),space:O(k)
//执行用时: 8 ms 内存消耗: 4.1 MB
func findKthLargestByMinHeap(nums []int, k int) int {
	intHeap := &MinIntHeap{}
	heap.Init(intHeap)
	for _, num := range nums {
		heap.Push(intHeap, num)
		if intHeap.Len() > k {
			heap.Pop(intHeap)
		}
	}

	return heap.Pop(intHeap).(int)
}

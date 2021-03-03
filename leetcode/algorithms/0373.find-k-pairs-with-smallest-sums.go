package algorithms

import "container/heap"

//time:O(n^2*logK),space:O(k)
// 执行用时: 4 ms 内存消耗: 4.7 MB
func kSmallestPairsByHeap(nums1 []int, nums2 []int, k int) [][]int {
	maxHeap := &MaxKSmallestPairsEntryHeap{}
	heap.Init(maxHeap)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum := nums1[i] + nums2[j]
			if maxHeap.Len() >= k && maxHeap.Top().sum < sum {
				continue
			}
			entry := &KSmallestPairsEntry{num1: nums1[i], num2: nums2[j], sum: sum}
			heap.Push(maxHeap, entry)
			if maxHeap.Len() > k {
				heap.Pop(maxHeap)
			}
		}
	}

	res := make([][]int, maxHeap.Len())
	for i := maxHeap.Len() - 1; i >= 0; i-- {
		top := heap.Pop(maxHeap).(*KSmallestPairsEntry)
		res[i] = []int{top.num1, top.num2}
	}

	return res
}

type KSmallestPairsEntry struct {
	num1 int
	num2 int
	sum  int
}

type MaxKSmallestPairsEntryHeap struct {
	arr []*KSmallestPairsEntry
}

func (m *MaxKSmallestPairsEntryHeap) Len() int {
	return len(m.arr)
}

func (m *MaxKSmallestPairsEntryHeap) Less(i, j int) bool {
	return m.arr[i].sum > m.arr[j].sum
}

func (m *MaxKSmallestPairsEntryHeap) Swap(i, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
}

func (m *MaxKSmallestPairsEntryHeap) Push(x interface{}) {
	m.arr = append(m.arr, x.(*KSmallestPairsEntry))
}

func (m *MaxKSmallestPairsEntryHeap) Pop() interface{} {
	entry := m.arr[len(m.arr)-1]
	m.arr = m.arr[:len(m.arr)-1]
	return entry
}

func (m *MaxKSmallestPairsEntryHeap) Top() *KSmallestPairsEntry {
	return m.arr[0]
}

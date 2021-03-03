package algorithms

import "container/heap"

//time:O(N*logK),space:O(k)
// 执行用时: 16 ms 内存消耗: 5.3 MB
func topKFrequentByHeap(nums []int, k int) []int {
	hash := make(map[int]int)
	for _, num := range nums {
		count := 0
		if exist, ok := hash[num]; ok {
			count = exist
		}
		count++
		hash[num] = count
	}

	intHeap := NewMinTopKFrequentEntryHeap()
	heap.Init(intHeap)
	for key, v := range hash {
		heap.Push(intHeap, &TopKFrequentEntry{key: key, value: v})
		// limit largest
		if intHeap.Len() > k {
			heap.Pop(intHeap)
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(intHeap).(*TopKFrequentEntry).key
	}
	return res
}

type TopKFrequentEntry struct {
	key   int
	value int
}

type MinTopKFrequentEntryHeap struct {
	arr []*TopKFrequentEntry
}

func NewMinTopKFrequentEntryHeap() *MinTopKFrequentEntryHeap {
	return &MinTopKFrequentEntryHeap{}
}

func (kl *MinTopKFrequentEntryHeap) Len() int {
	return len(kl.arr)
}

func (kl *MinTopKFrequentEntryHeap) Less(i, j int) bool {
	return kl.arr[i].value < kl.arr[j].value
}

func (kl *MinTopKFrequentEntryHeap) Swap(i, j int) {
	kl.arr[i], kl.arr[j] = kl.arr[j], kl.arr[i]
}

func (kl *MinTopKFrequentEntryHeap) Push(x interface{}) {
	kl.arr = append(kl.arr, x.(*TopKFrequentEntry))
}

func (kl *MinTopKFrequentEntryHeap) Pop() interface{} {
	val := kl.arr[kl.Len()-1]
	kl.arr = kl.arr[:kl.Len()-1]
	return val
}

func (kl *MinTopKFrequentEntryHeap) Top() *TopKFrequentEntry {
	return kl.arr[0]
}

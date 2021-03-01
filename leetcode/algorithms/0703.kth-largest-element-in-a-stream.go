package algorithms

import (
	"container/heap"
	"sort"
)

//time:O(N*logN),space:(k)
// 执行用时: 44 ms 内存消耗: 8.2 MB
type KthLargest struct {
	sort.IntSlice
	k int
}

func NewKthLargest(k int, nums []int) KthLargest {
	kthLargest := KthLargest{k: k}
	kthLargest.AddList(nums)
	return kthLargest
}

// Push implement heap.Interface
func (this *KthLargest) Push(x interface{}) {
	this.IntSlice = append(this.IntSlice, x.(int))
}

// Pop implement heap.Interface
func (this *KthLargest) Pop() interface{} {
	val := this.IntSlice[this.Len()-1]
	this.IntSlice = this.IntSlice[:this.Len()-1]
	return val
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}
	return this.IntSlice[0]
}

func (this *KthLargest) AddList(arr []int) {
	for _, val := range arr {
		this.Add(val)
	}
}

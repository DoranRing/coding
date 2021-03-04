package algorithms

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestMinIntHeap(t *testing.T) {
	arr := []int{3, 2, 1, 4, 5}
	intHeap := &MinIntHeap{}
	heap.Init(intHeap)
	for _, num := range arr {
		heap.Push(intHeap, num)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("pop: %d\n", heap.Pop(intHeap))
	}
}

func TestMaxIntHeap(t *testing.T) {
	arr := []int{3, 2, 1, 4, 5}
	intHeap := &MaxIntHeap{}
	heap.Init(intHeap)
	for _, num := range arr {
		heap.Push(intHeap, num)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("pop: %d\n", heap.Pop(intHeap))
	}
}

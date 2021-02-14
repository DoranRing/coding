package algorithms

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeap(t *testing.T) {
	arr := []int{3, 2, 1, 4, 5}
	intHeap := &IntHeap{}
	heap.Init(intHeap)
	for _, num := range arr {
		heap.Push(intHeap, num)
	}
	fmt.Printf("pop1: %d\n", heap.Pop(intHeap))
	fmt.Printf("pop2: %d\n", heap.Pop(intHeap))
	fmt.Printf("pop3: %d\n", heap.Pop(intHeap))
	fmt.Printf("pop4: %d\n", heap.Pop(intHeap))
	fmt.Printf("pop5: %d\n", heap.Pop(intHeap))
}

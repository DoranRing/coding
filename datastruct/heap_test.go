package datastruct

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewNonSortedArrayHeap(t *testing.T) {
	limit := 5
	heap := NewNonSortedArrayHeap(limit)
	heapTest(heap, limit)
}

// 155442 ns/op
func BenchmarkNewNonSortedArrayHeap(b *testing.B) {
	limit := 100
	for i := 0; i < b.N; i++ {
		heap := NewNonSortedArrayHeap(limit)
		heapBenchmarkTest(heap, limit, 1000, 1000*10)
	}
}

func TestNewSortedArrayHeap(t *testing.T) {
	limit := 5
	heap := NewSortedArrayHeap(limit)
	heapTest(heap, limit)
}

// 153443 ns/op
func BenchmarkNewSortedArrayHeap(b *testing.B) {
	limit := 100
	for i := 0; i < b.N; i++ {
		heap := NewNonSortedArrayHeap(limit)
		heapBenchmarkTest(heap, limit, 1000, 1000*10)
	}
}

func TestNewMaxBinaryHeap(t *testing.T) {
	limit := 20
	heap := NewMaxBinaryHeap()
	heapTest(heap, limit)
}

// 70241 ns/op
func BenchmarkNewMaxBinaryHeap(b *testing.B) {
	limit := 100
	for i := 0; i < b.N; i++ {
		heap := NewMaxBinaryHeap()
		heapBenchmarkTest(heap, limit, 1000, 1000*10)
	}
}

func heapTest(heap Heap, limit int) {
	nums := []int{3, 5, 7, 9, 2, 4, 6, 8, 1}
	for _, num := range nums {
		heap.Insert(num)
	}
	fmt.Printf("top:%d\n", heap.Top())
	for i := 0; i < limit; i++ {
		fmt.Printf("max:%d\n", heap.Pop())
	}
}

func heapBenchmarkTest(heap Heap, limit, num, max int) {
	nums := genIntNums(num, max)
	for i := 0; i < len(nums); i++ {
		heap.Insert(nums[i])
		if i%2 == 0 {
			heap.Pop()
		}
	}
	for i := 0; i < limit; i++ {
		heap.Pop()
	}
}

func genIntNums(len, max int) []int {
	res := make([]int, len, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		val := rand.Intn(max)
		res[i] = val
	}
	return res
}

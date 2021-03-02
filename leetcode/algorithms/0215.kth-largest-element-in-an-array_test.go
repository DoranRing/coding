package algorithms

import "testing"

func Test_findKthLargestByHeap(t *testing.T) {
	testFindKthLargest(t, findKthLargestByMaxHeap)
}

func Test_findKthLargestByMinHeap(t *testing.T) {
	testFindKthLargest(t, findKthLargestByMinHeap)
}

func testFindKthLargest(t *testing.T, fn func([]int, int) int) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{nums: []int{3, 2, 1, 5, 6, 4}, k: 2}, 5},
		{"2", args{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("actual %v, want %v", got, tt.want)
			}
		})
	}
}

package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		sortNums []int
		val      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{}, 5}, -1},
		{"2", args{[]int{1, 2, 3, 4, 5, 6}, 5}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchByBase(tt.args.sortNums, tt.args.val); got != tt.want {
				t.Errorf("BinarySearchByBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchByRangeFirst(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{}, 5}, -1},
		{"2", args{[]int{1, 2, 3, 4, 5, 6}, 5}, 4},
		{"3", args{[]int{1, 2, 3, 3, 3, 4, 5, 6}, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchByRangeFirst(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("BinarySearchByRangeFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

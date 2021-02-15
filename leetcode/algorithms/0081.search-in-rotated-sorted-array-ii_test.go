package algorithms

import "testing"

func Test_searchInRotatedSortedArrayIIByLoop(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 0}, true},
		{"2", args{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 3}, false},
		{"3", args{nums: []int{0, 0, 0, 0, 0, 0}, target: 0}, true},
		{"4", args{nums: []int{0, 0, 0, 0, 0, 0}, target: 10}, false},
		{"5", args{nums: []int{1, 0, 1, 1, 1}, target: 0}, true},
		{"6", args{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, target: 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInRotatedSortedArrayIIByLoop(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInRotatedSortedArrayIIByLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchInRotatedSortedArrayIIByBinarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{"1", args{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 0}, true},
		{"2", args{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 3}, false},
		{"3", args{nums: []int{0, 0, 0, 0, 0, 0}, target: 0}, true},
		{"4", args{nums: []int{0, 0, 0, 0, 0, 0}, target: 10}, false},
		{"5", args{nums: []int{1, 0, 1, 1, 1}, target: 0}, true},
		{"6", args{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, target: 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInRotatedSortedArrayIIByBinarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInRotatedSortedArrayIIByBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

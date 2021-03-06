package algorithms

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{[]int{}, []int{1}}, 1},
		{"2", args{[]int{2}, []int{}}, 2},
		{"3", args{[]int{0, 0}, []int{0, 0}}, 0},
		{"4", args{[]int{1, 2}, []int{3, 4}}, 2.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArraysByMergeTraversal(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArraysByMergeTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

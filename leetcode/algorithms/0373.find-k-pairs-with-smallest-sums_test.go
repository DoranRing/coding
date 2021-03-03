package algorithms

import (
	"reflect"
	"testing"
)

func Test_kSmallestPairsByHeap(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				nums1: []int{1, 7, 11},
				nums2: []int{2, 4, 6},
				k:     3,
			},
			[][]int{{1, 2}, {1, 4}, {1, 6}},
		},
		{
			"2",
			args{
				nums1: []int{1, 1, 2},
				nums2: []int{1, 2, 3},
				k:     2,
			},
			[][]int{{1, 1}, {1, 1}},
		},
		{
			"3",
			args{
				nums1: []int{1, 2},
				nums2: []int{3},
				k:     2,
			},
			[][]int{{1, 3}, {2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kSmallestPairsByHeap(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kSmallestPairsByHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

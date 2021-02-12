package algorithms

import (
	"reflect"
	"testing"
)

func Test_intersectionBySort(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 2, 1}, []int{2, 2}}, []int{2}},
		{"2", args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}}, []int{4, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectionBySort(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersectionBySort() = %v, want %v", got, tt.want)
			}
		})
	}
}

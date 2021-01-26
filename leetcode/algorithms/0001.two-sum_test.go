package algorithms

import (
	"reflect"
	"testing"
)

func Test_twoSumBySort(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumBySort(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumBySort() = %v, want %v", got, tt.want)
			}
		})
	}
}

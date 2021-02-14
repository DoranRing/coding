package algorithms

import (
	"reflect"
	"testing"
)

func Test_threeSumBySortAndTwoPoint(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[]int{}}, [][]int{}},
		{"2", args{[]int{0}}, [][]int{}},
		{"3", args{
			[]int{-1, 0, 1, 2, -1, -4}},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			"4",
			args{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}},
			[][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumBySortAndTwoPoint(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSumBySortAndTwoPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

package algorithms

import (
	"reflect"
	"testing"
)

func Test_mergeBySort(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1",
			args{
				intervals: [][]int{
					{1, 3},
					{2, 6},
					{8, 10},
					{15, 18},
				},
			},
			[][]int{
				{1, 6},
				{8, 10},
				{15, 18}},
		},
		{
			"2",
			args{
				intervals: [][]int{
					{1, 4},
					{4, 5},
				},
			},
			[][]int{
				{1, 5},
			},
		},
		{
			"3",
			args{
				intervals: [][]int{
					{1, 4},
					{0, 4},
				},
			},
			[][]int{
				{0, 4},
			},
		},
		{
			"4",
			args{
				intervals: [][]int{
					{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10},
				},
			},
			[][]int{
				{1, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeBySort(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeBySort() = %v, want %v", got, tt.want)
			}
		})
	}
}

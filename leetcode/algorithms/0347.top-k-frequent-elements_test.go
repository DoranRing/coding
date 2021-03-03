package algorithms

import (
	"reflect"
	"testing"
)

func Test_topKFrequentByHeap(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{nums: []int{1, 1, 1, 2, 2, 3}, k: 2}, []int{2, 1}},
		{"2", args{nums: []int{1}, k: 1}, []int{1}},
		{"3", args{nums: []int{1, 2}, k: 2}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequentByHeap(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequentByHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

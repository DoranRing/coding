package algorithms

import "testing"

func Test_arrayPairSumBySort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{nums: []int{1, 4, 3, 2}}, 4},
		{"2", args{nums: []int{6, 2, 6, 5, 1, 2}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayPairSumBySort(tt.args.nums); got != tt.want {
				t.Errorf("arrayPairSumBySort() = %v, want %v", got, tt.want)
			}
		})
	}
}

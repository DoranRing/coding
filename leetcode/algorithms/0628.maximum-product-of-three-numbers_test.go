package algorithms

import "testing"

func Test_maximumProductBySort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{nums: []int{1, 2, 3}}, 6},
		{"2", args{nums: []int{1, 2, 3, 4}}, 24},
		{"2", args{nums: []int{-1, -2, -3, -4}}, -6},
		{"2", args{nums: []int{-1, -2, -3, -4, 1}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProductBySort(tt.args.nums); got != tt.want {
				t.Errorf("maximumProductBySort() = %v, want %v", got, tt.want)
			}
		})
	}
}

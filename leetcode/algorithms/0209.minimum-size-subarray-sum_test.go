package algorithms

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{target: 7, nums: []int{2, 3, 1, 2, 4, 3}}, 2},
		{"2", args{target: 4, nums: []int{1, 4, 4}}, 1},
		{"3", args{target: 11, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLenBySlidingWindow(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLenBySlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

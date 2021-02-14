package algorithms

import "testing"

func Test_searchNums(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0}, 4},
		{"2", args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3}, -1},
		{"3", args{nums: []int{-1}, target: 0}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchNums(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

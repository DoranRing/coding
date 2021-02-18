package algorithms

import "testing"

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{nums: []int{2, 0, 2, 1, 1, 0}}},
		{"2", args{nums: []int{2, 0, 1}}},
		{"3", args{nums: []int{0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColorsByBucket(tt.args.nums)
		})
	}
}

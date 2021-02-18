package algorithms

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{nums: []int{1, 1, 1, 2, 2, 3}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicatesByTwoPoint(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicatesByTwoPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

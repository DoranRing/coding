package algorithms

import "testing"

func Test_removeElementByTwoPoint(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{nums: []int{3, 2, 2, 3}, val: 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElementByTwoPoint(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElementByTwoPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

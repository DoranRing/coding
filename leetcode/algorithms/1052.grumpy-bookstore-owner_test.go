package algorithms

import "testing"

func Test_maxSatisfiedBySlidingWindow(t *testing.T) {
	type args struct {
		customers []int
		grumpy    []int
		X         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{customers: []int{1, 0, 1, 2, 1, 1, 7, 5}, grumpy: []int{0, 1, 0, 1, 0, 1, 0, 1}, X: 3}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSatisfiedBySlidingWindow(tt.args.customers, tt.args.grumpy, tt.args.X); got != tt.want {
				t.Errorf("maxSatisfiedBySlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

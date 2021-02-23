package algorithms

import "testing"

func Test_longestOnesBySlidingWindow(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{A: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, K: 2}, 6},
		{"2", args{A: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, K: 3}, 10},
		{"3", args{A: []int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1}, K: 3}, 14},
		{"4", args{A: []int{0}, K: 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnesBySlidingWindow(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("longestOnesBySlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

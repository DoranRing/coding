package algorithms

import "testing"

func Test_maxScoreBySlidingWindow(t *testing.T) {
	type args struct {
		cardPoints []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{cardPoints: []int{1, 2, 3, 4, 5, 6, 1}, k: 3}, 12},
		{"2", args{cardPoints: []int{2, 2, 2}, k: 2}, 4},
		{"3", args{cardPoints: []int{9, 7, 7, 9, 7, 7, 9}, k: 7}, 55},
		{"4", args{cardPoints: []int{1, 1000, 1}, k: 1}, 1},
		{"4", args{cardPoints: []int{1, 79, 80, 1, 1, 1, 200, 1}, k: 3}, 202},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScoreBySlidingWindow(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("maxScoreBySlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

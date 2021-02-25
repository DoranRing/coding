package algorithms

import "testing"

func Test_equalSubstringBySlidingWindow(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{s: "abcd", t: "bcdf", maxCost: 3}, 3},
		{"2", args{s: "abcd", t: "cdef", maxCost: 3}, 1},
		{"3", args{s: "abcd", t: "acde", maxCost: 0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalSubstringBySlidingWindow(tt.args.s, tt.args.t, tt.args.maxCost); got != tt.want {
				t.Errorf("equalSubstringBySlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

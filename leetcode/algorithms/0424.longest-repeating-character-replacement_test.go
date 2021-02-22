package algorithms

import "testing"

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{s: "ABAB", k: 2}, 4},
		{"1", args{s: "AABABBA", k: 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacementBySlidingWindow(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacementBySlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

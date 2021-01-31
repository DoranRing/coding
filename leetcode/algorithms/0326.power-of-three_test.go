package algorithms

import (
	"math"
	"testing"
)

func Test_isPowerOfThreeByLoop(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{27}, true},
		{"2", args{1}, true},
		{"3", args{9}, true},
		{"4", args{45}, false},
		{"5", args{math.MaxInt32}, false},
		{"5", args{1162261467}, true},
		{"5", args{1162261468}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThreeByLoop(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThreeByLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

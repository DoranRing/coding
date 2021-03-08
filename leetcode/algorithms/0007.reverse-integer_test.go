package algorithms

import "testing"

func Test_reverseByMath(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{123}, 321},
		{"2", args{-123}, -321},
		{"3", args{120}, 21},
		{"4", args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseByMath(tt.args.x); got != tt.want {
				t.Errorf("reverseByMath() = %v, want %v", got, tt.want)
			}
		})
	}
}
